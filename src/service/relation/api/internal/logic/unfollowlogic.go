package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/relation/rpc/pb/relation"

	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfollowLogic) Unfollow(req *types.UnfollowRequest) (resp *types.UnfollowResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	if uid == req.Fid {
		return nil, errors.New(http.StatusBadRequest, "无效操作")
	}

	in := relation.UnfollowRequest{
		Uid: uid,
		Fid: req.Fid,
	}
	_, err = l.svcCtx.RelationRpc.Unfollow(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "您尚未关注此用户")
			}
		}
		l.Logger.Errorf("[RPC] Follow failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	return &types.UnfollowResponse{}, nil
}
