package tests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Arthur-Conti/authentication-project/internal/adapter/controllers"
	"github.com/Arthur-Conti/authentication-project/internal/adapter/repository"
	"github.com/Arthur-Conti/authentication-project/internal/infra/database"
	"github.com/Arthur-Conti/authentication-project/internal/usecase/services"
)

func TestCreateUserService(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	repository := repository.NewUserRepository(database)
	service := services.NewUserService(repository)
	t.Run("must create a new user", func(t *testing.T) {
		request := controllers.CreateUserRequest{
			FirstName: "arthur",
			LastName:  "silva",
			FullName:  "arthur silva",
			Email:     fmt.Sprintf("asilva%v@teste.com", rand.Int()),
			Password:  "12345678",
			IsAdmin:   true,
		}
		_, err := service.CreateUser(request)
		if err != nil {
			t.Errorf("error: %v", err.Error())
		}
	})
	t.Run("must not create a user with a email that already exists", func(t *testing.T) {
		request := controllers.CreateUserRequest{
			FirstName: "arthur",
			LastName:  "silva",
			FullName:  "arthur silva",
			Email:     "asilva@teste.com",
			Password:  "12345678",
			IsAdmin:   true,
		}
		_, err := service.CreateUser(request)
		if err == nil {
			t.Errorf("no errors found, user created with an wrong information")
		}
	})
	t.Run("must not create a user with a wrong full name", func(t *testing.T) {
		request := controllers.CreateUserRequest{
			FirstName: "arthur",
			LastName:  "silva",
			FullName:  "arthur",
			Email:     "asilva@teste.com",
			Password:  "12345678",
			IsAdmin:   true,
		}
		_, err := service.CreateUser(request)
		if err == nil {
			t.Errorf("no errors found, user created with an wrong information")
		}
	})
	t.Run("must not create a user with a wrong password", func(t *testing.T) {
		request := controllers.CreateUserRequest{
			FirstName: "arthur",
			LastName:  "silva",
			FullName:  "arthur silva",
			Email:     "asilva@teste.com",
			Password:  "12345",
			IsAdmin:   true,
		}
		_, err := service.CreateUser(request)
		if err == nil {
			t.Errorf("no errors found, user created with an wrong information")
		}
	})
}

func TestGetUserByID(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	repository := repository.NewUserRepository(database)
	service := services.NewUserService(repository)
	t.Run("must return a user with the given id", func(t *testing.T) {
		id := "041c689d-9a35-4020-8759-082c4fdbe251"
		user, err := service.GetUserByID(id)
		if err != nil {
			t.Errorf("error: %v", err.Error())
		}
		if user.GetID() != id {
			t.Error("user id do not match")
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	repository := repository.NewUserRepository(database)
	service := services.NewUserService(repository)
	t.Run("must return a user with the given email", func(t *testing.T) {
		email := "asilva@teste.com"
		user, err := service.GetUserByEmail(email)
		if err != nil {
			t.Errorf("error: %v", err.Error())
		}
		if user.GetEmail() != email {
			t.Error("user email do not match")
		}
	})
}
