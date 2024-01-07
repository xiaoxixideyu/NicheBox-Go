package logic

import (
	"context"
	"net/http"

	"nichebox/service/user/api/internal/common"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReqeust) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetUidByEmailAndPwd(l.ctx, &user.GetUidByEmailAndPwdRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if ok {
			code := grpcStatus.Code()
			if code == codes.NotFound || code == codes.Unauthenticated {
				return nil, errors.New(http.StatusBadRequest, "邮箱或密码错误")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "login 服务出错: 1")
	}

	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, refreshToken, err := common.CreateTokenAndRefreshToken(res.Uid, accessExpire, refreshExpire, accessSecret)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
