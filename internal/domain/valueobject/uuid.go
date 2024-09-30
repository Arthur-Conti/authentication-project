package valueobject

import "github.com/google/uuid"

type IdUUID struct {
	value string
}

func NewUUID(value string) *IdUUID {
	return &IdUUID{
		value: value,
	}
}

func CreateUUID() *IdUUID {
	id := uuid.New()
	return &IdUUID{
		value: id.String(),
	}
}

func (idUUID *IdUUID) GetValue() string {
	return idUUID.value
}