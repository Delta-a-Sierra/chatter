package valueobjects

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p *Password) Encrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(*p), 72)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: failed to generate encrypted password, error = %w", err)
	}
	*p = Password(password)
	return nil
}

func (p Password) Compare(password Password) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
	if err == nil {
		return true
	}
	switch {
	case errors.Is(err, bcrypt.ErrHashTooShort):
		return password == p
	default:
		return false
	}
}

const (
	maxCharLength = 22
	minCharLength = 8
)

var (
	ErrPasswordTooLong             = fmt.Errorf("maxium char length of %d for password exceeded", maxCharLength)
	ErrPasswordTooShort            = fmt.Errorf("minimum char length of %d for password not reached", maxCharLength)
	ErrPasswordDoesntContainUpper  = errors.New("password doesn't have at least 1 uppercase character")
	ErrPasswordDoesntContainLower  = errors.New("password doesn't have at least 1 lowercase character")
	ErrPasswordDoesntContainNumber = errors.New("password doesn't have at least 1 number character")
)

func (p Password) Validate() error {
	var err error
	if len(p) > maxCharLength {
		err = errors.Join(err, ErrPasswordTooLong)
	}
	if len(p) < minCharLength {
		err = errors.Join(err, ErrPasswordTooShort)
	}
	upprCaseExp := regexp.MustCompile("[A-Z]")
	if !upprCaseExp.MatchString(string(p)) {
		err = errors.Join(err, ErrPasswordDoesntContainUpper)
	}
	lowerCaseExp := regexp.MustCompile("[a-z]")
	if !lowerCaseExp.MatchString(string(p)) {
		err = errors.Join(err, ErrPasswordDoesntContainLower)
	}
	numExp := regexp.MustCompile("[0-9]")
	if !numExp.MatchString(string(p)) {
		err = errors.Join(err, ErrPasswordDoesntContainNumber)
	}
	return err
}
