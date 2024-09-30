package entity

import (
	"fmt"

	"github.com/Arthur-Conti/authentication-project/internal/domain/valueobject"
)

type User struct {
	id        valueobject.IdUUID
	firstName valueobject.Name
	lastName  valueobject.Name
	fullName  valueobject.FullName
	email     valueobject.Email
	password  valueobject.Password
	status    valueobject.Status
	isAdmin   bool
}

func NewUser(id string, firstName string, lastName string, fullName string, email string, status string, password string, isAdmin bool) (*User, error) {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error instantiating a user object: %v", err.Error())
	}
	newStatus, err := valueobject.NewStatus(status)
	if err != nil {
		return nil, fmt.Errorf("error instantiating a user object: %v", err.Error())
	}
	newPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error instantiating a user object: %v", err.Error())
	}
	newFullName, err := valueobject.NewFullName(fullName)
	if err != nil {
		return nil, fmt.Errorf("error instantiating a user object: %v", err.Error())
	}
	return &User{
		id:        *valueobject.NewUUID(id),
		firstName: *valueobject.NewName(firstName),
		lastName:  *valueobject.NewName(lastName),
		fullName:  *newFullName,
		email:     *newEmail,
		status:    *newStatus,
		password:  *newPassword,
		isAdmin:   isAdmin,
	}, nil
}

func CreateUser(firstName string, lastName string, fullName string, email string, password string, isAdmin bool) (*User, error) {
	id := valueobject.CreateUUID().GetValue()
	status := "active"
	return NewUser(
		id,
		firstName,
		lastName,
		fullName,
		email,
		status,
		password,
		isAdmin,
	)
}

func (user *User) GetID() string {
	return user.id.GetValue()
}

func (user *User) GetFirstName() string {
	return user.firstName.GetValue()
}

func (user *User) GetLastName() string {
	return user.lastName.GetValue()
}

func (user *User) GetFullName() string {
	return user.fullName.GetValue()
}

func (user *User) GetEmail() string {
	return user.email.GetValue()
}

func (user *User) GetPassword() string {
	return user.password.GetValue()
}

func (user *User) GetStatus() string {
	return user.status.GetValue()
}

func (user *User) GetIsadmin() bool {
	return user.isAdmin
}
