package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/user/rpc/pb/user"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

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
		Code:     req.Code,
	}
	out, err := l.svcCtx.UserRpc.Register(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.AlreadyExists {
				return nil, errors.New(http.StatusBadRequest, "此邮箱已被注册")
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
		// todo: 注册成功但是登录失败， 注册的resp增加一个标识注册是否成功的字段，此时不帮忙自动登录了，单单响应注册成功
	}

	return &types.RegisterResponse{
		Token:        "",
		RefreshToken: "",
	}
}
