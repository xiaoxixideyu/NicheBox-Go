package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type BoxContentCacheInterface interface {
	GetPostIDsCtx(ctx context.Context, boxID int64, page, size int, order string) ([]redis.Pair, error)
	UpdateNewPostsCtx(ctx context.Context, boxID int64, infos []*ModifiedPostInfo) error
	UpdateDeletedPostsCtx(ctx context.Context, boxID int64, infos []*ModifiedPostInfo) error
}
