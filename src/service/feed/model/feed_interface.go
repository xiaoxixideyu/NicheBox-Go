package model

import (
	"context"
	"time"
)

type FeedInterface interface {
	AddFeed(feed *Feed) error
	GetFeeds(followings []int64, page, size int) ([]*Feed, error)
}

type FeedCacheInterface interface {
	BatchDeliverFeedsToOutboxesCtx(ctx context.Context, MessageID int64, MessageType int, publishTime time.Time, followers []int64) error
}
