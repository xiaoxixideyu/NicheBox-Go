package logic

import (
	"context"
	"encoding/json"
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
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := user.SetCriticalUserInfoRequest{
		Uid:      uid,
		Password: req.Password,
	}
	_, err = l.svcCtx.UserRpc.SetCriticalUserInfo(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SetCriticalUserInfoResponse{}, nil
}
