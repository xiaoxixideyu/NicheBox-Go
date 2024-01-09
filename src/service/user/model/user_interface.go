package model

import "context"

type UserInterface interface {
	GerUserByUid(uid int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}

type UserRedisInterface interface {
	GetVerificationCode(ctx context.Context, key string) (string, error)
	SetVerificationCode(ctx context.Context, key, code string, expiration int) error
	RemoveVerificationCode(ctx context.Context, key string) error
}
