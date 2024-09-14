package valueobjects_test

import (
	"errors"
	"testing"

	valueobjects "github.com/Delta-a-Sierra/chatter/internal/app/core/value_objects"
	"github.com/stretchr/testify/assert"
)

func Test_Password_Validate(t *testing.T) {
	type testcase struct {
		Err      error
		Password valueobjects.Password
	}
	tests := map[string]testcase{
		"sad - test short password returns length error": {
			Err:      valueobjects.ErrPasswordTooShort,
			Password: "a1A",
		},
		"sad - test empty password returns length error": {
			Err:      valueobjects.ErrPasswordTooShort,
			Password: "",
		},
		"happy - test valid password returns nil": {
			Err:      nil,
			Password: "Passw0rd!",
		},
		"sad - test password with no upper returns ErrPasswordDoesntContainUpper": {
			Err:      valueobjects.ErrPasswordDoesntContainUpper,
			Password: "passw0rd!",
		},
		"sad - test valid password with no lowercase returns ErrPasswordDoesntContainLower": {
			Err:      valueobjects.ErrPasswordDoesntContainLower,
			Password: "PASSW0RD!",
		},
		"sad - test password with no number returns ErrPasswordDoesntContainNumber": {
			Err:      valueobjects.ErrPasswordDoesntContainNumber,
			Password: "Password!",
		},
		"sad - test password with with too many chars returns ErrPasswordTooLong": {
			Err:      valueobjects.ErrPasswordTooLong,
			Password: "Password!123456789123456789",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Log(test.Password.Validate())
			assert.ErrorIs(t, test.Password.Validate(), test.Err)
		})
	}
}

func Test_Password_Validate_multi(t *testing.T) {
	type testcase struct {
		Err      error
		Password valueobjects.Password
	}
	tests := map[string]testcase{
		"sad - test all errors are combined and output with short passwords": {
			Err: errors.Join(
				valueobjects.ErrPasswordTooShort,
				valueobjects.ErrPasswordDoesntContainUpper,
				valueobjects.ErrPasswordDoesntContainLower,
				valueobjects.ErrPasswordDoesntContainNumber,
			),
			Password: "",
		},
		"sad - test all errors are combined and ouput with long passwords": {
			Err: errors.Join(
				valueobjects.ErrPasswordTooLong,
				valueobjects.ErrPasswordDoesntContainUpper,
				valueobjects.ErrPasswordDoesntContainLower,
				valueobjects.ErrPasswordDoesntContainNumber,
			),
			Password: "..................................",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Log(test.Password.Validate())
			assert.ErrorContains(t, test.Password.Validate(), test.Err.Error())
		})
	}
}

func Test_Password_Encrypt(t *testing.T) {
}
