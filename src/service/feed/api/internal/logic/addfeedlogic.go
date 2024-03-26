package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/feed/rpc/pb/feed"

	"nichebox/service/feed/api/internal/svc"
	"nichebox/service/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFeedLogic {
	return &AddFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFeedLogic) AddFeed(req *types.AddFeedRequest) (resp *types.AddFeedResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := feed.PushRequest{
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
		Author:      uid,
		CreateTime:  req.PublishTime,
	}
	_, err = l.svcCtx.FeedRpc.Push(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Add feed failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.AddFeedResponse{}, nil
}
