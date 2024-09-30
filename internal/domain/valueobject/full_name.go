package valueobject

import (
	"fmt"
	"strings"
)

type FullName struct {
	value string
}

func NewFullName(value string) (*FullName, error) {
	if !strings.Contains(value, " ") {
		return nil, fmt.Errorf("full name must contain at least two names separated by whitespace") 
	}
	return &FullName{value: value}, nil
}

func (fn *FullName) GetValue() string {
	return fn.value
}