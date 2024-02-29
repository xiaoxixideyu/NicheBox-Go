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

type GetUserAvatarFileIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAvatarFileIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAvatarFileIdLogic {
	return &GetUserAvatarFileIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAvatarFileIdLogic) GetUserAvatarFileId(in *user.GetUserAvatarFileIdRequest) (*user.GetUserAvatarFileIdResponse, error) {
	userModel, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.GetUserAvatarFileIdResponse{
		OriginId: userModel.AvatarOriginId,
		WebpId:   userModel.AvatarWebpId,
	}, nil
}
