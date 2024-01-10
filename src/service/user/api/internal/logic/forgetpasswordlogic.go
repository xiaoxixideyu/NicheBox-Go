package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/user/rpc/pb/user"
	"strings"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(req *types.ForgetPasswordRequest) (resp *types.ForgetPasswordResponse, err error) {
	in := user.ForgetPasswordRequest{
		Email:       req.Email,
		NewPassword: req.NewPassword,
		Code:        strings.ToUpper(req.Code),
	}
	_, err = l.svcCtx.UserRpc.ForgetPassword(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	return &types.ForgetPasswordResponse{}, nil
}
