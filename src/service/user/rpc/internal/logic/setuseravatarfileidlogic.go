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

type SetUserAvatarFileIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserAvatarFileIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAvatarFileIdLogic {
	return &SetUserAvatarFileIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserAvatarFileIdLogic) SetUserAvatarFileId(in *user.SetUserAvatarFileIdRequest) (*user.SetUserAvatarFileIdResponse, error) {
	userModel, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	userModel.AvatarOriginId = in.OriginId
	userModel.AvatarWebpId = in.WebpId

	err = l.svcCtx.UserInterface.UpdateUserTX(userModel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.SetUserAvatarFileIdResponse{}, nil
}
