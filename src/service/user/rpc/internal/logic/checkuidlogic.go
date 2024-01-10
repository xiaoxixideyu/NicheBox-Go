package logic

import (
	"context"
	"errors"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CheckUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUidLogic {
	return &CheckUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUidLogic) CheckUid(in *user.CheckUidRequest) (*user.CheckUidResponse, error) {
	_, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user.CheckUidResponse{
				Exists: false,
			}, nil
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &user.CheckUidResponse{
		Exists: true,
	}, nil
}
