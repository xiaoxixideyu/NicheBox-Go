package model

import "context"

type EmailRedisInterface interface {
	SetVerificationCode(ctx context.Context, destination, code string, expiration int) error
}
