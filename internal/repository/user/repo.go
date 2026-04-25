package user
import (
	"github.com/malfoit/SimpleProject/internal/repository"
	"github.com/malfoit/SimpleProject/internal/model"
)
type repo struct{
    users map[string]*model.User // ключ - email
    mu    sync.RWMutex
}
func NewRepository() repository.UserRepo {
	return &repo{
		users: make(map[string]*model.User),
	}
}