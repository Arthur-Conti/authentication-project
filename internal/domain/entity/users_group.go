package entity

import (
	"fmt"

	"github.com/Arthur-Conti/authentication-project/internal/domain/valueobject"
)

type UsersGroup struct {
	id valueobject.IdUUID
	name valueobject.Name
	description valueobject.Description
	status valueobject.Status
}

func NewUsersGroup(id, name, description, status string) (*UsersGroup, error) {
	newStatus, err := valueobject.NewStatus(status)
	if err != nil {
		return nil, fmt.Errorf("error instanciating users group: %v", err.Error())
	} 
	return &UsersGroup{
		id: *valueobject.NewUUID(id),
		name: *valueobject.NewName(name),
		description: *valueobject.NewDescription(description),
		status: *newStatus,
	}, nil
}

func CreateUsersGroup(name, description string) (*UsersGroup, error) {
	id := valueobject.CreateUUID().GetValue()
	status := "active"
	return NewUsersGroup(
		id,
		name,
		description,
		status,
	)
}

func (usersGroup *UsersGroup) GetID() string {
	return usersGroup.id.GetValue()
}

func (usersGroup *UsersGroup) GetName() string {
	return usersGroup.name.GetValue()
}

func (usersGroup *UsersGroup) GetDescription() string {
	return usersGroup.description.GetValue()
}

func (usersGroup *UsersGroup) GetStatus() string {
	return usersGroup.status.GetValue()
}
