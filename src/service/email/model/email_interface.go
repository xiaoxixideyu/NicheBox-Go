package model

import "context"

type EmailRedisInterface interface {
	SetVerificationCodeRegister(ctx context.Context, destination, code string, expiration int) error
	SetVerificationCodePWD(ctx context.Context, destination, code string, expiration int) error
}
