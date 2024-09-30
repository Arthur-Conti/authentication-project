package ports

import (
	"github.com/Arthur-Conti/authentication-project/internal/domain/entity"
)

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (string, error)
	GetUserByID(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}