package model

type UserInterface interface {
	GerUserByUid(uid int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUserTX(user *User) error
	CreateUser(user *User) error
}
