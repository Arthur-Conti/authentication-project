package controllers

type CreateUserRequest struct {
	FirstName string
	LastName  string
	FullName  string
	Email     string
	Password  string
	IsAdmin   bool
}

type CreateUsersGroupRequest struct {
	Name        string
	Description string
}
