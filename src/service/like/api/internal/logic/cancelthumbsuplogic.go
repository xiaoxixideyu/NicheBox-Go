package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/like/rpc/pb/like"

	"nichebox/service/like/api/internal/svc"
	"nichebox/service/like/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelThumbsUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelThumbsUpLogic {
	return &CancelThumbsUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelThumbsUpLogic) CancelThumbsUp(req *types.CancelThumbsUpRequest) (resp *types.CancelThumbsUpResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := like.CancelThumbsUpRequest{
		Uid:         uid,
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
	}
	_, err = l.svcCtx.LikeRpc.CancelThumbsUp(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误-1")
	}

	return &types.CancelThumbsUpResponse{}, nil
}
