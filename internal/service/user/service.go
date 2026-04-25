package user
import(
	"github.com/malfoit/SimpleProject/internal/repository"
	"github.com/malfoit/SimpleProject/internal/model"
	"github.com/malfoit/SimpleProject/internal/service"
)
type service struct {
	repo repository.UserRepo
}
func NewService(repo repository.UserRepo) service.UserService {
	return &service{repo: repo}
}
