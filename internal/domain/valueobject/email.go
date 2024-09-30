package valueobject

import (
	"fmt"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(value) {
		return nil, fmt.Errorf("invalid email: %v", value)
	}
	return &Email{
		value: value,
	}, nil
}

func (email *Email) GetValue() string {
	return email.value
}
