package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"nichebox/service/post/model"

	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	postModel := model.Post{
		PostID: in.PostID,
	}
	err := l.svcCtx.PostInterface.DeletePost(&postModel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "帖子不存在")
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &post.DeletePostResponse{}, nil
}
