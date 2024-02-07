package logic

import (
	"context"
	"math"
	"nichebox/service/like/model/redis"
	"strconv"

	"nichebox/service/like/rpc/internal/svc"
	"nichebox/service/like/rpc/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbsUpHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpHistoryLogic {
	return &ThumbsUpHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbsUpHistoryLogic) ThumbsUpHistory(in *like.ThumbsUpHistoryRequest) (*like.ThumbsUpHistoryResponse, error) {
	start := (in.Page - 1) * in.Size
	stop := start + in.Size - 1
	ids, err := l.svcCtx.LikeCacheInterface.GetThumbsUpHistoryCtx(l.ctx, int(in.MessageType), in.Uid, int(start), int(stop))
	var sizeFromDB int
	needRewriteCache := false
	if err != nil || len(ids) == 0 {
		// need to query all data from DB
		sizeFromDB = int(in.Size)
	} else {
		sizeFromDB = int(in.Size) - len(ids)
		// if redis data is empty, need to rewrite cache
		if sizeFromDB == int(in.Size) {
			needRewriteCache = true
			sizeFromDB = int(math.Max(float64(sizeFromDB), redis.HistoryListLength))
		}
	}
	messageIDs := make([]int64, 0, len(ids))
	// redis result convert string to int64
	for _, id := range ids {
		messageID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		messageIDs = append(messageIDs, messageID)
	}
	// if redis can provide all data, just return then
	if sizeFromDB == 0 {
		return &like.ThumbsUpHistoryResponse{MessageIDs: messageIDs}, nil
	}

	// query DB
	startFromDB := int(stop) - sizeFromDB + 1
	likesFromDB, err := l.svcCtx.LikeInterface.GetLikeByUpdateDateDesc(int(in.MessageType), in.Uid, sizeFromDB, startFromDB)
	if err != nil {
		return nil, err
	}

	for _, l := range likesFromDB {
		messageIDs = append(messageIDs, l.MessageID)
	}

	go func() {
		// rewrite cache (only when redis data is empty but db data is not empty)
		if needRewriteCache {
			cap := math.Min(redis.HistoryListLength, float64(len(likesFromDB)))
			likesToRewrite := likesFromDB[0:int(cap)]
			err := l.svcCtx.LikeCacheInterface.BatchAddThumbsUpHistoryCtx(context.Background(), likesToRewrite)
			if err != nil {
				l.svcCtx.LikeCacheInterface.ClearAllThumbsUpHistoryCtx(context.Background(), int(in.MessageType), in.Uid)
			}
		}
	}()

	return &like.ThumbsUpHistoryResponse{MessageIDs: messageIDs}, nil
}
