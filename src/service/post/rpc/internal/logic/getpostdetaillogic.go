package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"

	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *post.GetPostDetailRequest) (*post.GetPostDetailResponse, error) {
	postModel, err := l.svcCtx.PostInterface.GetPostByID(in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	out := post.GetPostDetailResponse{
		AuthorID: postModel.AuthorID,
		BoxID:    postModel.BoxID,
		Title:    postModel.Title,
		Content:  postModel.Content,
		Photos:   transformPhotoStringToArray(postModel.Photos),
		Cover:    postModel.Cover,
		UserView: int32(postModel.UserView),
	}

	uv, err := l.svcCtx.PostCacheInterface.GetUserView(l.ctx, in.PostID)
	// 如果redis出错，用mysql数据兜底即可
	if err != nil {
		return &out, nil
	}

	out.UserView = int32(uv)

	return &out, nil
}

func transformPhotoStringToArray(photos string) []string {
	photoArray := strings.Split(photos, Separator)
	return photoArray
}
