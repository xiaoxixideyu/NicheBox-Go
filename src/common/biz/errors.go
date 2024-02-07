package biz

import "errors"

var (
	ErrRedisOutOfBounds = errors.New("page/size too large")
)
