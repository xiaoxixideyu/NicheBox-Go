package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/service/like/rpc/pb/like"

	"nichebox/service/like/api/internal/svc"
	"nichebox/service/like/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpCountLogic {
	return &ThumbsUpCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpCountLogic) ThumbsUpCount(req *types.ThumbsUpCountRequest) (resp *types.ThumbsUpCountResponse, err error) {
	in := like.ThumbsUpCountRequest{
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
	}
	out, err := l.svcCtx.LikeRpc.ThumbsUpCount(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return &types.ThumbsUpCountResponse{Count: 0}, nil
			}
		}
		return nil, err
	}

	return &types.ThumbsUpCountResponse{Count: int(out.Count)}, nil
}
