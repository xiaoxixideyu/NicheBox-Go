package logic

import (
	"context"
	"nichebox/service/like/model"

	"nichebox/service/like/rpc/internal/svc"
	"nichebox/service/like/rpc/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelThumbsUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelThumbsUpLogic {
	return &CancelThumbsUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelThumbsUpLogic) CancelThumbsUp(in *like.CancelThumbsUpRequest) (*like.CancelThumbsResponse, error) {
	likeModel := model.Like{
		MessageID: in.MessageID,
		Uid:       in.Uid,
		TypeID:    int(in.MessageType),
	}
	err := l.svcCtx.LikeInterface.DeleteLikeAndUpdateLikeCountTX(&likeModel)
	if err != nil {
		return nil, err
	}

	// remove cache (cache aside)
	l.svcCtx.LikeCacheInterface.DeleteThumbsUpCountCtx(l.ctx, likeModel.MessageID, likeModel.TypeID)

	l.svcCtx.LikeCacheInterface.RemoveThumbsUpHistoryCtx(context.Background(), likeModel.MessageID, likeModel.TypeID, in.Uid)

	return &like.CancelThumbsResponse{}, nil
}
