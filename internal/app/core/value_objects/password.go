package valueobjects

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p *Password) Encrypt() error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("Password.Validate: %w", err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(*p), 12)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
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
	maxPasswordCharLength = 22
	minPasswordCharLength = 8
)

var (
	ErrPasswordTooLong             = fmt.Errorf("maxium char length of %d for password exceeded", maxPasswordCharLength)
	ErrPasswordTooShort            = fmt.Errorf("minimum char length of %d for password not reached", minPasswordCharLength)
	ErrPasswordDoesntContainUpper  = errors.New("password doesn't have at least 1 uppercase character")
	ErrPasswordDoesntContainLower  = errors.New("password doesn't have at least 1 lowercase character")
	ErrPasswordDoesntContainNumber = errors.New("password doesn't have at least 1 number character")
)

func (p Password) ToString() string {
	return string(p)
}

func (p Password) Validate() error {
	var err error
	if len(p) > maxPasswordCharLength {
		err = errors.Join(err, ErrPasswordTooLong)
	}
	if len(p) < minPasswordCharLength {
		err = errors.Join(err, ErrPasswordTooShort)
	}
	upprCaseExp := regexp.MustCompile("[A-Z]")
	if !upprCaseExp.MatchString(p.ToString()) {
		err = errors.Join(err, ErrPasswordDoesntContainUpper)
	}
	lowerCaseExp := regexp.MustCompile("[a-z]")
	if !lowerCaseExp.MatchString(p.ToString()) {
		err = errors.Join(err, ErrPasswordDoesntContainLower)
	}
	numExp := regexp.MustCompile("[0-9]")
	if !numExp.MatchString(p.ToString()) {
		err = errors.Join(err, ErrPasswordDoesntContainNumber)
	}
	return err
}
