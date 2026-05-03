package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) UpdatePassword(ctx context.Context, id, password, passwordConfirm string) error {
	if len(password) < 8 || len(password) > 72 {
		return errors.New("password must be between 8 and 72 characters")
	}
	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	if err = s.repo.UpdatePasswordHash(ctx, id, string(hash)); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
