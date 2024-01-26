package logic

import (
	"context"
	"nichebox/common/snowflake"
	"nichebox/service/post/model"
	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

const Separator = ";"

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	postID := snowflake.GenID()

	postModel := model.Post{
		PostID:   postID,
		AuthorID: in.AuthorID,
		BoxID:    in.BoxID,
		Title:    in.Title,
		Content:  in.Content,
		Photos:   transformPhotoArrayToString(in.Photos),
		Cover:    in.Cover,
	}
	err := l.svcCtx.PostInterface.CreatePost(&postModel)
	if err != nil {
		return nil, err
	}

	return &post.CreatePostResponse{PostID: postID}, nil
}

func transformPhotoArrayToString(photos []string) string {
	var photoStr string
	for idx, p := range photos {
		photoStr += p
		if idx != len(photos)-1 {
			photoStr += Separator
		}
	}
	return photoStr
}
