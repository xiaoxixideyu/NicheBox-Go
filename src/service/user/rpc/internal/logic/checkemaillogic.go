package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	_, err := l.svcCtx.UserInterface.GetUserByEmail(in.Email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return &user.CheckEmailResponse{Exists: false}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.CheckEmailResponse{Exists: true}, nil
}
