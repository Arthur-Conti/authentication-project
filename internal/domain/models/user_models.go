package models

type User struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Email     string
	Password  string
	Status    string
	IsAdmin   bool
}
