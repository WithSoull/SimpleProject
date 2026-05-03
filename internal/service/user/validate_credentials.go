package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) ValidateCredentials(ctx context.Context, email, password string) (string, bool, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return "", false, nil
		}
		return "", false, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", false, nil
	}
	return u.ID, true, nil
}
