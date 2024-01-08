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

type CheckEmailExistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckEmailExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailExistsLogic {
	return &CheckEmailExistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckEmailExistsLogic) CheckEmailExists(req *types.CheckEmailExistsReqeust) (resp *types.CheckEmailExistsResponse, err error) {
	in := user.CheckEmailRequest{Email: req.Email}
	out, err := l.svcCtx.UserRpc.CheckEmail(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	return &types.CheckEmailExistsResponse{Exist: out.Exists}, nil
}
