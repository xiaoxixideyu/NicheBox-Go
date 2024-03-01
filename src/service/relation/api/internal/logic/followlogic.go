package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/relation/rpc/pb/relation"

	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowRequest) (resp *types.FollowResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	if uid == req.Fid {
		return nil, errors.New(http.StatusBadRequest, "您不能关注自己")
	}

	in := relation.FollowRequest{
		Uid: uid,
		Fid: req.Fid,
	}
	_, err = l.svcCtx.RelationRpc.Follow(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Follow failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	return &types.FollowResponse{}, nil
}
