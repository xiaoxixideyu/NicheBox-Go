package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/like/rpc/pb/like"

	"nichebox/service/like/api/internal/svc"
	"nichebox/service/like/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpLogic {
	return &ThumbsUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpLogic) ThumbsUp(req *types.ThumbsUpRequest) (resp *types.ThumbsUpResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := like.ThumbsUpRequest{
		Uid:         uid,
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
	}
	_, err = l.svcCtx.LikeRpc.ThumbsUp(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.AlreadyExists {
				return nil, errors.New(http.StatusBadRequest, "您已经点赞过了哦~")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	return &types.ThumbsUpResponse{}, nil
}
