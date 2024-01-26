package model

import "context"

type TaskInterface interface {
	UpdatePostUserView(postID int64, userView int64) error
}

type TaskCacheInterface interface {
	GetUserView(ctx context.Context, postID int64) (int64, error)
}
