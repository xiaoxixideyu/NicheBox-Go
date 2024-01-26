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

type ThumbsUpHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpHistoryLogic {
	return &ThumbsUpHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpHistoryLogic) ThumbsUpHistory(req *types.ThumbsUpHistoryRequest) (resp *types.ThumbsUpHistoryResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := like.ThumbsUpHistoryRequest{
		Uid:         uid,
		MessageType: int32(req.MessageType),
		Page:        int32(req.Page),
		Size:        int32(req.Size),
	}

	out, err := l.svcCtx.LikeRpc.ThumbsUpHistory(l.ctx, &in)
	if err != nil {

	}

	return &types.ThumbsUpHistoryResponse{MessageIDs: out.MessageIDs}, nil
}
