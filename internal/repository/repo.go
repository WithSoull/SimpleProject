package repository
import (
	"context"
)
type UserRepo interface {
	Create(ctx context.Context, name, email string, password string) error
}
