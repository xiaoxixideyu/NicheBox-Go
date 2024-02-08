package biz

import "errors"

var (
	ErrRedisOutOfBounds  = errors.New("page/size too large")
	ErrRedisUnknownOrder = errors.New("unknown order")
)
