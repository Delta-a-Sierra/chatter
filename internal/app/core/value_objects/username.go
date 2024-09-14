package valueobjects

import (
	"errors"
	"fmt"
	"regexp"
)

type Username string

const (
	maxUsernameCharLength = 20
	minUsernameCharLength = 6
)

var (
	ErrUsernameTooLong                   = fmt.Errorf("maxium char length of %d for username exceeded", maxUsernameCharLength)
	ErrUsernameTooShort                  = fmt.Errorf("minimum char length of %d for username not reached", minUsernameCharLength)
	ErrUsernameContainsSpecialCharacters = fmt.Errorf("username contains special characters")
)

func (u Username) ToString() string {
	return string(u)
}

func (u Username) Validate() error {
	var err error
	if len(u) < minUsernameCharLength {
		err = errors.Join(err, ErrUsernameTooShort)
	}
	if len(u) > maxUsernameCharLength {
		err = errors.Join(err, ErrUsernameTooLong)
	}
	specialCharExp := regexp.MustCompile("[^A-Za-z0-9_-]")
	if specialCharExp.MatchString(u.ToString()) {
		err = errors.Join(err, ErrUsernameContainsSpecialCharacters)
	}

	return err
}
