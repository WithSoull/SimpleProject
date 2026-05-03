package user

import (
	"context"
	"errors"

	userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
)

func (s *userService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
