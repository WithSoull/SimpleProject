package user

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) Update(ctx context.Context, id string, name, email *string) error {
	if name != nil {
		trimmed := strings.TrimSpace(*name)
		if len(trimmed) < 3 || len(trimmed) > 50 {
			return errors.New("name must be between 3 and 50 characters")
		}
		*name = trimmed
	}
	if email != nil {
		trimmed := strings.TrimSpace(*email)
		if _, err := mail.ParseAddress(trimmed); err != nil {
			return errors.New("invalid email format")
		}
		*email = trimmed
	}

	if err := s.repo.Update(ctx, id, name, email); err != nil {
		switch {
		case errors.Is(err, userRepo.ErrNotFound):
			return errors.New("user not found")
		case errors.Is(err, userRepo.ErrAlreadyExists):
			return errors.New("email already taken")
		}
		return err
	}
	return nil
}
