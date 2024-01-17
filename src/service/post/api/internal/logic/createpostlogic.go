package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/post/rpc/pb/post"

	"nichebox/service/post/api/internal/svc"
	"nichebox/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := post.CreatePostRequest{
		AuthorID: uid,
		BoxID:    req.BoxID,
		Title:    req.Title,
		Content:  req.Content,
		Photos:   req.Photos,
		Cover:    req.Cover,
	}

	out, err := l.svcCtx.PostRpc.CreatePost(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.CreatePostResponse{PostID: out.PostID}, nil
}
