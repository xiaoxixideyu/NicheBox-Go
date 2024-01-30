package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"nichebox/service/like/model"
	"strconv"

	"nichebox/service/like/rpc/internal/svc"
	"nichebox/service/like/rpc/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbsUpCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpCountLogic {
	return &ThumbsUpCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbsUpCountLogic) ThumbsUpCount(in *like.ThumbsUpCountRequest) (*like.ThumbsUpCountResponse, error) {
	countStr, err := l.svcCtx.LikeCacheInterface.GetThumbsUpCountCtx(l.ctx, in.MessageID, uint8(in.MessageType))
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get thumbs up count error", err)
		}
		// cache expired
		likeCountModel := model.LikeCount{
			TypeID:    uint8(in.MessageType),
			MessageID: in.MessageID,
		}
		err := l.svcCtx.LikeInterface.GetLikeCount(&likeCountModel)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// rewrite cache
			l.svcCtx.LikeCacheInterface.SetThumbsUpCountCtx(l.ctx, in.MessageID, uint8(in.MessageType), 0)
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if err != nil {
			return nil, err
		}
		// rewrite cache and return count from mysql
		l.svcCtx.LikeCacheInterface.SetThumbsUpCountCtx(l.ctx, in.MessageID, uint8(in.MessageType), likeCountModel.Count)
		return &like.ThumbsUpCountResponse{Count: int32(likeCountModel.Count)}, nil
	}

	// cache valid, return count from cache
	count, _ := strconv.Atoi(countStr)
	return &like.ThumbsUpCountResponse{Count: int32(count)}, nil
}
