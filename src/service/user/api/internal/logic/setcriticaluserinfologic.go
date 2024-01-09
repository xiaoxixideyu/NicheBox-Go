package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/user/rpc/pb/user"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCriticalUserInfoLogic {
	return &SetCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCriticalUserInfoLogic) SetCriticalUserInfo(req *types.SetCriticalUserInfoRequest) (resp *types.SetCriticalUserInfoResponse, err error) {
	in := user.SetCriticalUserInfoRequest{
		Password: req.Password,
	}
	_, err = l.svcCtx.UserRpc.SetCriticalUserInfo(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SetCriticalUserInfoResponse{}, nil
}
