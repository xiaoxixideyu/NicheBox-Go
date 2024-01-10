package model

import "context"

type UserInterface interface {
	GetUserByUid(uid int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUserTX(user *User) error
	CreateUser(user *User) error
	UpdatePasswordByEmail(email, password string) error
}

type UserCacheInterface interface {
	GetVerificationCode(ctx context.Context, key string) (string, error)
	SetVerificationCode(ctx context.Context, key, code string, expiration int) error
	RemoveVerificationCode(ctx context.Context, key string) error
}
