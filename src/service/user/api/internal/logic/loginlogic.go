package logic

import (
	"context"
	"net/http"
	"time"

	"nichebox/common/jwtx"
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
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Uid)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 2")
	}
	refreshToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, refreshExpire, res.Uid)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 3")
	}

	return &types.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
