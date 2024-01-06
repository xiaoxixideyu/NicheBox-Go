package model

type UserInterface interface {
	GetUserByEmail(email string) (*User, error)
}
