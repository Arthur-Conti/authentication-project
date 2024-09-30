package valueobject

import "strings"

type Name struct {
	value string
}

func NewName(value string) *Name {
	cleanValue := strings.TrimSpace(value)
	finalValue := strings.ToLower(cleanValue)
	return &Name{value: finalValue}
}

func (n *Name) GetValue() string {
	return n.value
}