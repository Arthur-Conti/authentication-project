package tests

import (
	"math/rand"
	"fmt"
	"testing"

	"github.com/Arthur-Conti/authentication-project/internal/adapter/controllers"
	"github.com/Arthur-Conti/authentication-project/internal/adapter/repository"
	"github.com/Arthur-Conti/authentication-project/internal/infra/database"
	"github.com/Arthur-Conti/authentication-project/internal/usecase/services"
)

var (
	admin_active_user = "041c689d-9a35-4020-8759-082c4fdbe251"
	not_admin_active_user = "13b98016-239d-4b30-a576-be4296abc7c3"
	admin_inactive_user = "09c2dd87-e885-455f-ab5b-25ee513bd2a9"
)

func TestCreateUsersGroup(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	usersGroupRepository := repository.NewUsersGroupRepository(database)
	userRepository := repository.NewUserRepository(database)
	service := services.NewUsersGroupService(usersGroupRepository, userRepository)
	t.Run("must create a new users group", func(t *testing.T) {
		userID := admin_active_user
		request := controllers.CreateUsersGroupRequest{
			Name:        fmt.Sprintf("test group %v", rand.Int()),
			Description: "a group to test",
		}
		_, err := service.CreateUsersGroup(userID, request)
		if err != nil {
			t.Errorf("error creating users group: %v", err.Error())
		}
	})
	t.Run("must not create a group if user is not admin", func(t *testing.T) {
		userID := not_admin_active_user
		request := controllers.CreateUsersGroupRequest{
			Name: fmt.Sprintf("test group %v", rand.Int()),
			Description: "a group to test",
		}
		_, err := service.CreateUsersGroup(userID, request)
		if err == nil {
			t.Errorf("must not create a users group if the user is not admin")
		}
	})
	t.Run("must not create a group if the user is inactive", func(t *testing.T) {
		userID := admin_inactive_user
		request := controllers.CreateUsersGroupRequest{
			Name: fmt.Sprintf("test group %v", rand.Int()),
			Description: "a group to test",
		}
		_, err := service.CreateUsersGroup(userID, request)
		if err == nil {
			t.Errorf("must not create a users group if the user is inactive")
		}
	})
}

func TestGetUsersGroupByID(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	usersGroupRepository := repository.NewUsersGroupRepository(database)
	userRepository := repository.NewUserRepository(database)
	service := services.NewUsersGroupService(usersGroupRepository, userRepository)
	t.Run("must return a users group by the given id", func(t *testing.T) {
		groupID := "71952782-a491-4ce7-b5de-3df94613b01b"
		group, err := service.GetUsersGroupByID(groupID)
		if err  != nil {
			t.Errorf("error getting users group by id: %v", err.Error())
		}
		if group.GetID() != groupID {
			t.Errorf("ids do not match")
		}
	})
}

func TestAddUserToGroup(t *testing.T) {
	database := database.NewDatabase()
	database.Connect()
	defer database.Close()
	usersGroupRepository := repository.NewUsersGroupRepository(database)
	userRepository := repository.NewUserRepository(database)
	service := services.NewUsersGroupService(usersGroupRepository, userRepository)
	t.Run("must add a user to a group", func(t *testing.T) {
		userID := admin_active_user
		groupID := "71952782-a491-4ce7-b5de-3df94613b01b"
		_, err := service.AddUserToGroup(userID, groupID)
		if err != nil {
			t.Errorf("error adding user to a group: %v", err.Error())
		}
	})
	t.Run("must not add a user to a group if user is not admin", func(t *testing.T) {
		userID := not_admin_active_user
		groupID := "71952782-a491-4ce7-b5de-3df94613b01b"
		_, err := service.AddUserToGroup(userID, groupID)
		if err == nil {
			t.Errorf("must not add a user to a group if user is not admin")
		}
	})
	t.Run("must not add a user to a group if user is inactive", func(t *testing.T) {
		userID := admin_inactive_user
		groupID := "09c2dd87-e885-455f-ab5b-25ee513bd2a9"
		_, err := service.AddUserToGroup(userID, groupID)
		if err == nil {
			t.Errorf("must not add a user to a group if user is inactive")
		}
	})
	t.Run("must not add a user to a group if group is inactive", func(t *testing.T) {
		userID := admin_active_user
		groupID := "16025229-d365-407c-9c83-4ff184b1dcc7"
		_, err := service.AddUserToGroup(userID, groupID)
		if err == nil {
			t.Errorf("must not add a user to a group if group is inactive")
		}
	})
}
