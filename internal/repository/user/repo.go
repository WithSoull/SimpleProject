package user

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/malfoit/SimpleProject/internal/model"
	"github.com/malfoit/SimpleProject/internal/repository"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
)

type repo struct {
	mu       sync.RWMutex
	byID     map[string]*model.User
	emailIdx map[string]string // email -> id
}

func NewRepository() repository.UserRepo {
	return &repo{
		byID:     make(map[string]*model.User),
		emailIdx: make(map[string]string),
	}
}

func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (r *repo) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.emailIdx[user.UserInfo.Email]; exists {
		return ErrAlreadyExists
	}

	user.ID = newID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	cp := *user
	r.byID[user.ID] = &cp
	r.emailIdx[user.UserInfo.Email] = user.ID
	return nil
}

func (r *repo) GetByID(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, exists := r.byID[id]
	if !exists {
		return nil, ErrNotFound
	}
	cp := *u
	return &cp, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.emailIdx[email]
	if !exists {
		return nil, ErrNotFound
	}
	cp := *r.byID[id]
	return &cp, nil
}

func (r *repo) Update(ctx context.Context, id string, name, email *string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, exists := r.byID[id]
	if !exists {
		return ErrNotFound
	}

	if email != nil && *email != u.UserInfo.Email {
		if _, taken := r.emailIdx[*email]; taken {
			return ErrAlreadyExists
		}
		delete(r.emailIdx, u.UserInfo.Email)
		u.UserInfo.Email = *email
		r.emailIdx[*email] = id
	}
	if name != nil {
		u.UserInfo.Name = *name
	}
	u.UpdatedAt = time.Now()
	return nil
}

func (r *repo) UpdatePasswordHash(ctx context.Context, id, passwordHash string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, exists := r.byID[id]
	if !exists {
		return ErrNotFound
	}
	u.PasswordHash = passwordHash
	u.UpdatedAt = time.Now()
	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, exists := r.byID[id]
	if !exists {
		return ErrNotFound
	}
	delete(r.emailIdx, u.UserInfo.Email)
	delete(r.byID, id)
	return nil
}
