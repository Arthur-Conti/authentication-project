package services

import (
	"fmt"
	"reflect"

	"github.com/Arthur-Conti/authentication-project/internal/adapter/controllers"
	"github.com/Arthur-Conti/authentication-project/internal/domain/entity"
	"github.com/Arthur-Conti/authentication-project/internal/usecase/ports"
)

type UserService struct {
	userRepository ports.UserRepositoryInterface
}

func NewUserService(userRepository ports.UserRepositoryInterface) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) CreateUser(request controllers.CreateUserRequest) (*UserCreatedDTO, error) {
	user, err := entity.CreateUser(
		request.FirstName,
		request.LastName,
		request.FullName,
		request.Email,
		request.Password,
		request.IsAdmin,
	)
	if err != nil {
		return nil, err
	}
	userCheck, _ := service.userRepository.GetUserByEmail(user.GetEmail())
	emptyUser := reflect.New(reflect.TypeOf(userCheck)).Elem()
	if !reflect.DeepEqual(userCheck, emptyUser.Interface()) {
		return nil, fmt.Errorf("user already exists")
	}
	userID, err := service.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &UserCreatedDTO{id: userID}, nil
}

func (service *UserService) GetUserByID(id string) (*entity.User, error) {
	return service.userRepository.GetUserByID(id)
}

func (service *UserService) GetUserByEmail(email string) (*entity.User, error) {
	return service.userRepository.GetUserByEmail(email)
}

type UserCreatedDTO struct {
	id string
}
