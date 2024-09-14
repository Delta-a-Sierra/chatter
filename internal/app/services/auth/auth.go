package auth

import valueobjects "github.com/Delta-a-Sierra/chatter/internal/app/core/value_objects"

type Service struct{}

type credentials struct {
	password valueobjects.Password
}

func (s *Service) Login(username, password string) (int, error) {
	return 1, nil
}
