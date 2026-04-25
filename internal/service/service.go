package service 
import (
	"context"
	"github.com/malfoit/SimpleProject/internal/model"
)
type UserService interface {
	Create(ctx context.Context, user model.User) (email string, err error)
}