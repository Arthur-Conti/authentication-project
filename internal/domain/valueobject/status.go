package valueobject

import "fmt"

type Status struct {
	value string
}

func NewStatus(value string) (*Status, error) {
	allowedStatus := []string{"active", "inactive"}
	for _, item := range allowedStatus {
		if value == item {
			return &Status{
				value: value,
			}, nil
		}
	}
	return nil, fmt.Errorf("status must be either active or inactive")
}

func (status *Status) GetValue() string {
	return status.value
}