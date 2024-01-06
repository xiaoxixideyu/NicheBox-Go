package logic

import (
	"context"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailLogic {
	return &CheckEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckEmailLogic) CheckEmail(in *user.CheckEmailRequest) (*user.CheckEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &user.CheckEmailResponse{}, nil
}
