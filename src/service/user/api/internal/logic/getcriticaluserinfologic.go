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

type GetCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCriticalUserInfoLogic {
	return &GetCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCriticalUserInfoLogic) GetCriticalUserInfo(req *types.GetCriticalUserInfoRequest) (resp *types.GetCriticalUserInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := user.GetCriticalUserInfoRequest{Uid: uid}
	out, err := l.svcCtx.UserRpc.GetCriticalUserInfo(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.GetCriticalUserInfoResponse{Email: out.Email, Telephone: out.Telephone}, nil
}
