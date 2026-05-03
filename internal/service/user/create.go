package user

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/malfoit/SimpleProject/internal/model"
	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) Create(ctx context.Context, name, email, password, passwordConfirm string) (string, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if len(name) < 3 || len(name) > 50 {
		return "", errors.New("name must be between 3 and 50 characters")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return "", errors.New("invalid email format")
	}
	if len(password) < 8 || len(password) > 72 {
		return "", errors.New("password must be between 8 and 72 characters")
	}
	if password != passwordConfirm {
		return "", errors.New("passwords do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	u := &model.User{
		UserInfo:     model.UserInfo{Name: name, Email: email},
		PasswordHash: string(hash),
	}
	if err = s.repo.Create(ctx, u); err != nil {
		if errors.Is(err, userRepo.ErrAlreadyExists) {
			return "", errors.New("user with this email already exists")
		}
		return "", err
	}
	return u.ID, nil
}
