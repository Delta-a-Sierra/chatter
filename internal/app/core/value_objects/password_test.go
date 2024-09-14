package valueobjects_test

import (
	"errors"
	"testing"

	valueobjects "github.com/Delta-a-Sierra/chatter/internal/app/core/value_objects"
	"github.com/stretchr/testify/assert"
)

func Test_Password_Validate(t *testing.T) {
	type testcase struct {
		err      error
		password valueobjects.Password
	}
	tests := map[string]testcase{
		"sad - test short password returns length error": {
			err:      valueobjects.ErrPasswordTooShort,
			password: "a1A",
		},
		"sad - test empty password returns length error": {
			err:      valueobjects.ErrPasswordTooShort,
			password: "",
		},
		"happy - test valid password returns nil": {
			err:      nil,
			password: "Passw0rd!",
		},
		"sad - test password with no upper returns ErrPasswordDoesntContainUpper": {
			err:      valueobjects.ErrPasswordDoesntContainUpper,
			password: "passw0rd!",
		},
		"sad - test valid password with no lowercase returns ErrPasswordDoesntContainLower": {
			err:      valueobjects.ErrPasswordDoesntContainLower,
			password: "PASSW0RD!",
		},
		"sad - test password with no number returns ErrPasswordDoesntContainNumber": {
			err:      valueobjects.ErrPasswordDoesntContainNumber,
			password: "Password!",
		},
		"sad - test password with with too many chars returns ErrPasswordTooLong": {
			err:      valueobjects.ErrPasswordTooLong,
			password: "Password!123456789123456789",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Log(test.password.Validate())
			assert.ErrorIs(t, test.password.Validate(), test.err)
		})
	}
}

func Test_Password_Validate_multi(t *testing.T) {
	type testcase struct {
		err      error
		password valueobjects.Password
	}
	tests := map[string]testcase{
		"sad - test all errors are combined and output with short passwords": {
			err: errors.Join(
				valueobjects.ErrPasswordTooShort,
				valueobjects.ErrPasswordDoesntContainUpper,
				valueobjects.ErrPasswordDoesntContainLower,
				valueobjects.ErrPasswordDoesntContainNumber,
			),
			password: "",
		},
		"sad - test all errors are combined and ouput with long passwords": {
			err: errors.Join(
				valueobjects.ErrPasswordTooLong,
				valueobjects.ErrPasswordDoesntContainUpper,
				valueobjects.ErrPasswordDoesntContainLower,
				valueobjects.ErrPasswordDoesntContainNumber,
			),
			password: "..................................",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Log(test.password.Validate())
			assert.ErrorContains(t, test.password.Validate(), test.err.Error())
		})
	}
}

func Test_Password_ToString(t *testing.T) {
	assert.Equal(t, "", valueobjects.Password("").ToString())

	assert.Equal(t, "Passw0rdx!", valueobjects.Password("Passw0rdx!").ToString())
}

func Test_Password_Encrypt(t *testing.T) {
	type testcase struct {
		err      error
		password valueobjects.Password
	}
	tests := map[string]testcase{
		"sad - test that invalid short password returns correct validation error": {
			password: "",
			err:      valueobjects.ErrPasswordTooShort,
		},
		"sad - test that invalid long password returns correct validation error": {
			password: "Passw0rd1xxxxxxxxxxxxxxxxxxxxxxx",
			err:      valueobjects.ErrPasswordTooLong,
		},
		"sad - test that invalid no upper password returns correct validation error": {
			password: "passw0rd1!",
			err:      valueobjects.ErrPasswordDoesntContainUpper,
		},
		"sad - test that invalid no lower password contains correct validation error": {
			password: "PASSW0RD1!",
			err:      valueobjects.ErrPasswordDoesntContainLower,
		},
		"sad - test that invalid no number password contains correct validation error": {
			password: "Passwordx!",
			err:      valueobjects.ErrPasswordDoesntContainNumber,
		},
		"happy - test that valid password is Encrypted successfully": {
			password: "Passw0rd1x",
			err:      nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			prevPassword := test.password
			err := test.password.Encrypt()
			assert.ErrorIs(t, err, test.err)
			if err == nil {
				assert.NotEqual(t, prevPassword, test.password)
			}
		})
	}
}

func Test_Password_Compare(t *testing.T) {
	type testcase struct {
		password valueobjects.Password
		encrypt  bool
		isMatch  bool
	}
	tests := map[string]testcase{
		"happy - test plain passwords return correct true false values": {
			password: "password",
			isMatch:  true,
		},
		"happy - test encrypted passwords return correct true false values  ": {
			password: "Passw0rd1x",
			isMatch:  true,
			encrypt:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			prev := test.password
			if test.encrypt {
				if err := test.password.Encrypt(); err != nil {
					t.Fatal("failed encrypt password", err)
				}
			}
			assert.Equal(t, test.isMatch, test.password.Compare(prev))
			assert.NotEqual(t, test.isMatch, test.password.Compare(prev+"xyhz"))
		})
	}
}
