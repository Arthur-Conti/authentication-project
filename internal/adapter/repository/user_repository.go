package repository

import (
	"github.com/Arthur-Conti/authentication-project/internal/domain/entity"
	"github.com/Arthur-Conti/authentication-project/internal/domain/models"
	"github.com/Arthur-Conti/authentication-project/internal/infra/database"
)

type userRepository struct {
	database database.DatabaseInterface
}

func NewUserRepository(database database.DatabaseInterface) *userRepository {
	return &userRepository{database: database}
}

func (repository *userRepository) CreateUser(user *entity.User) (string, error) {
	conn := repository.database.GetConnection()
	query := "INSERT INTO users (id, first_name, last_name, full_name, email, password, status, is_admin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := conn.Exec(
		repository.database.GetContext(),
		query,
		user.GetID(), user.GetFirstName(), user.GetLastName(), user.GetFullName(), user.GetEmail(), user.GetPassword(), user.GetStatus(), user.GetIsadmin(),
	)
	if err != nil {
		return "", err
	}
	return user.GetID(), nil
}

func (repository *userRepository) GetUserByID(id string) (*entity.User, error) {
	conn := repository.database.GetConnection()
	query := "SELECT * FROM users WHERE id=$1"
	var user models.User
	err := conn.QueryRow(
		repository.database.GetContext(),
		query,
		id,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.FullName, &user.Email, &user.Password, &user.Status, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(
		user.ID,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.Email,
		user.Status,
		user.Password,
		user.IsAdmin,
	)
}

func (repository *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	conn := repository.database.GetConnection()
	query := "SELECT * FROM users WHERE email=$1"
	var user models.User
	err := conn.QueryRow(
		repository.database.GetContext(),
		query,
		email,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.FullName, &user.Email, &user.Password, &user.Status, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(
		user.ID,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.Email,
		user.Status,
		user.Password,
		user.IsAdmin,
	)
}
