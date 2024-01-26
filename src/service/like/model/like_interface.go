package model

import "context"

type LikeInterface interface {
	CreateLikeAndUpdateLikeCountTX(likeModel *Like) error
	DeleteLikeAndUpdateLikeCountTX(likeModel *Like) error
	GetLikeCount(likeCountModel *LikeCount) error
	GetLikeByUpdateDateDesc(typeID uint8, uid int64, limit int, offset int) ([]*Like, error)
}

type LikeCacheInterface interface {
	SetThumbsUpCountCtx(ctx context.Context, messageID int64, typeID uint8, count int) error
	GetThumbsUpCountCtx(ctx context.Context, messageID int64, typeID uint8) (string, error)
	DeleteThumbsUpCountCtx(ctx context.Context, messageID int64, typeID uint8) (int, error)
	AddThumbsUpHistoryCtx(ctx context.Context, messageID int64, typeID uint8, uid int64) error
	BatchAddThumbsUpHistoryCtx(ctx context.Context, likes []*Like) error
	RemoveThumbsUpHistoryCtx(ctx context.Context, messageID int64, typeID uint8, uid int64) error
	ClearAllThumbsUpHistoryCtx(ctx context.Context, typeID uint8, uid int64) error
	GetThumbsUpHistoryCtx(ctx context.Context, typeID uint8, uid int64, start int, stop int) ([]string, error)
}
