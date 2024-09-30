package valueobject

import "fmt"

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if len(value) < 8 {
		return nil, fmt.Errorf("password must contain more then 8 characteres")
	}
	return &Password{
		value: value,
	}, nil
}

func (password *Password) GetValue() string {
	return password.value
}