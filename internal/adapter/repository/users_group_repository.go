package repository

import (
	"github.com/Arthur-Conti/authentication-project/internal/domain/entity"
	"github.com/Arthur-Conti/authentication-project/internal/domain/models"
	"github.com/Arthur-Conti/authentication-project/internal/infra/database"
)

type usersGroupRepository struct {
	database database.DatabaseInterface
}

func NewUsersGroupRepository(database database.DatabaseInterface) *usersGroupRepository {
	return &usersGroupRepository{database: database}
}

func (repository *usersGroupRepository) CreateUsersGroup(usersGroup *entity.UsersGroup) (string, error) {
	conn := repository.database.GetConnection()
	query := "INSERT INTO users_groups (id, name, description, status) VALUES ($1, $2, $3, $4)"
	_, err := conn.Exec(
		repository.database.GetContext(),
		query,
		usersGroup.GetID(), usersGroup.GetName(), usersGroup.GetDescription(), usersGroup.GetStatus(),
	)
	if err != nil {
		return "", err
	}
	return usersGroup.GetID(), nil
}

func (repository *usersGroupRepository) GetUsersGroupByID(id string) (*entity.UsersGroup, error) {
	conn := repository.database.GetConnection()
	query := "SELECT * FROM users_groups where id=$1"
	var usersGroup models.UsersGroup
	err := conn.QueryRow(
		repository.database.GetContext(),
		query,
		id,
	).Scan(&usersGroup.ID, &usersGroup.Name, &usersGroup.Description, &usersGroup.Status)
	if err != nil {
		return nil, err
	}
	return entity.NewUsersGroup(
		usersGroup.ID,
		usersGroup.Name,
		usersGroup.Description,
		usersGroup.Status,
	)
}

func (repository *usersGroupRepository) AddUserToGroup(user entity.User, group entity.UsersGroup) error {
	conn := repository.database.GetConnection()
	query := "INSERT INTO usersXgroups (user_id, group_id) VALUES ($1, $2)"
	_, err := conn.Exec(
		repository.database.GetContext(),
		query,
		user.GetID(), group.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}
