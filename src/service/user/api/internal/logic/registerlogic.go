package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/pb/user"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	in := user.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Code:     strings.ToUpper(req.Code),
	}

	_, err = l.svcCtx.UserRpc.Register(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.AlreadyExists {
				return nil, errors.New(http.StatusBadRequest, "此邮箱已被注册")
			} else if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "验证码错误或过期")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	loginLogic := NewLoginLogic(l.ctx, l.svcCtx)
	loginReq := types.LoginReqeust{
		Email:    req.Email,
		Password: req.Password,
	}
	loginResp, err := loginLogic.Login(&loginReq)
	if err != nil {
		// Login failed but dont response error
		return &types.RegisterResponse{
			LoginSuccess: false,
		}, nil
	}

	return &types.RegisterResponse{
		LoginSuccess: true,
		Token:        loginResp.Token,
		RefreshToken: loginResp.RefreshToken,
	}, nil
}
