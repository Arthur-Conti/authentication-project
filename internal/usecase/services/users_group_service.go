package services

import (
	"fmt"

	"github.com/Arthur-Conti/authentication-project/internal/adapter/controllers"
	"github.com/Arthur-Conti/authentication-project/internal/domain/entity"
	"github.com/Arthur-Conti/authentication-project/internal/usecase/ports"
)

type UsersGroupService struct {
	usersGroupRepository ports.UsersGroupRepositoryInterface
	userRepository       ports.UserRepositoryInterface
}

func NewUsersGroupService(usersGroupRepository ports.UsersGroupRepositoryInterface, userRepository ports.UserRepositoryInterface) *UsersGroupService {
	return &UsersGroupService{
		usersGroupRepository: usersGroupRepository,
		userRepository:       userRepository,
	}
}

func (service *UsersGroupService) CreateUsersGroup(userID string, request controllers.CreateUsersGroupRequest) (*UsersGroupCreatedDTO, error) {
	user, err := service.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if !user.GetIsadmin() {
		return nil, fmt.Errorf("user must be admin to create users groups")
	}
	if user.GetStatus() == "inactive" {
		return nil, fmt.Errorf("user must be active to create users groups")
	}
	usersGroup, err := entity.CreateUsersGroup(
		request.Name,
		request.Description,
	)
	if err != nil {
		return nil, err
	}
	usersGroupID, err := service.usersGroupRepository.CreateUsersGroup(usersGroup)
	if err != nil {
		return nil, err
	}
	return &UsersGroupCreatedDTO{id: usersGroupID}, nil
}

func (service *UsersGroupService) GetUsersGroupByID(groupID string) (*entity.UsersGroup, error) {
	return service.usersGroupRepository.GetUsersGroupByID(groupID)
}

func (service *UsersGroupService) AddUserToGroup(userID, groupID string) (*AddUserToGroupDTO, error) {
	user, err := service.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if !user.GetIsadmin() {
		return nil, fmt.Errorf("user must be admin to add users to groups")
	}
	if user.GetStatus() == "inactive" {
		return nil, fmt.Errorf("user must be active to be added into a group")
	}
	usersGroup, err := service.usersGroupRepository.GetUsersGroupByID(groupID)
	if err != nil {
		return nil, err
	}
	if usersGroup.GetStatus() == "inactive" {
		return nil, fmt.Errorf("group must be active recive new users")
	}
	service.usersGroupRepository.AddUserToGroup(*user, *usersGroup)
	return &AddUserToGroupDTO{id: groupID}, nil
}

type UsersGroupCreatedDTO struct {
	id string
}

type AddUserToGroupDTO struct {
	id string
}
