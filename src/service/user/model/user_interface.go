package model

type UserInterface interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}
