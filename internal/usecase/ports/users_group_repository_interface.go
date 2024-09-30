package ports

import "github.com/Arthur-Conti/authentication-project/internal/domain/entity"

type UsersGroupRepositoryInterface interface {
	CreateUsersGroup(usersGroup *entity.UsersGroup) (string, error)
	GetUsersGroupByID(groupID string) (*entity.UsersGroup, error)
	AddUserToGroup(user entity.User, group entity.UsersGroup) error
}