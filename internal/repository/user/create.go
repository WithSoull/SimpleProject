package user 
import (
	"context"
    "github.com/malfoit/SimpleProject/internal/model"
)
func (r *repo) Create(ctx context.Context, name, email string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    user := &model.User{
        Name:     name,
        Email:    email,
        Password: password,
    }
    
    r.users[email] = user
	return nil
}