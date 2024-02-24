package biz

import "errors"

var (
	ErrRedisOutOfBounds  = errors.New("page/size too large")
	ErrRedisUnknownOrder = errors.New("unknown order")
)

var (
	ErrConnectionNotEnoughBytes = errors.New("connection not enough bytes")
	ErrConnectionReadTimeout    = errors.New("connection i/o timeout")
	ErrConnectionNotFound       = errors.New("connection not found")
)
