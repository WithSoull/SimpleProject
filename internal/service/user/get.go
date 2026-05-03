package user

import (
	"context"
	"errors"

	"github.com/malfoit/SimpleProject/internal/model"
	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) Get(ctx context.Context, id string) (*model.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}
