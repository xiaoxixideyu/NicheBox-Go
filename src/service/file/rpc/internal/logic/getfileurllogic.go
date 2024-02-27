package logic

import (
	"context"
	"errors"

	"nichebox/service/file/rpc/internal/svc"
	"nichebox/service/file/rpc/pb/file"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetFileUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUrlLogic {
	return &GetFileUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileUrlLogic) GetFileUrl(in *file.GetFileUrlRequest) (*file.GetFileUrlResponse, error) {
	imageModel, err := l.svcCtx.ImageFileInterface.GetImageByFileId(in.FileId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "资源不存在")
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &file.GetFileUrlResponse{
		FileUrl: imageModel.Url,
	}, nil
}
