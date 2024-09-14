package valueobjects_test

import (
	"testing"

	valueobjects "github.com/Delta-a-Sierra/chatter/internal/app/core/value_objects"
	"github.com/stretchr/testify/assert"
)

func Test_Username_Validate(t *testing.T) {
	type testcase struct {
		err      error
		username string
	}
	tests := map[string]testcase{
		"happy - test valid username returns no error": {
			username: "bobbyruss",
			err:      nil,
		},
		"sad - test username with special character ! returns ErrUsernameContainsSpecialCharacters": {
			username: "bobby!russ",
			err:      valueobjects.ErrUsernameContainsSpecialCharacters,
		},
		"sad - test username with special character . returns ErrUsernameContainsSpecialCharacters": {
			username: "bobby.russ",
			err:      valueobjects.ErrUsernameContainsSpecialCharacters,
		},
		"sad - test username with special character { returns ErrUsernameContainsSpecialCharacters": {
			username: "bobby{russ",
			err:      valueobjects.ErrUsernameContainsSpecialCharacters,
		},
		"sad - test username with special character } returns ErrUsernameContainsSpecialCharacters": {
			username: "bobby}russ",
			err:      valueobjects.ErrUsernameContainsSpecialCharacters,
		},
		"sad - test username with special character \\ returns ErrUsernameContainsSpecialCharacters": {
			username: "bobby\\russ",
			err:      valueobjects.ErrUsernameContainsSpecialCharacters,
		},
		"sad - test username with allowed speicla character - returns no error": {
			username: "bobby-russ",
			err:      nil,
		},
		"happy - test username with allowed speicla character _ returns no error": {
			username: "bobby_russ",
			err:      nil,
		},
		"sad - test username with too many ccharacters returns ErrUsernameTooLong": {
			username: "bobby-russell-the-first-great-ganderer",
			err:      valueobjects.ErrUsernameTooLong,
		},

		"sad - test username with too little ccharacters returns ErrUsernameTooShort": {
			username: "bo",
			err:      valueobjects.ErrUsernameTooShort,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.ErrorIs(t, valueobjects.Username(test.username).Validate(), test.err)
		})
	}
}

func Test_Username_ToString(t *testing.T) {
	assert.Equal(t, "", valueobjects.Username("").ToString())

	assert.Equal(t, "username", valueobjects.Username("username").ToString())
}
