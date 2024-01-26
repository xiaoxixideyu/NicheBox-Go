// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package server

import (
	"context"

	"nichebox/service/post/rpc/internal/logic"
	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"
)

type PostServer struct {
	svcCtx *svc.ServiceContext
	post.UnimplementedPostServer
}

func NewPostServer(svcCtx *svc.ServiceContext) *PostServer {
	return &PostServer{
		svcCtx: svcCtx,
	}
}

func (s *PostServer) CreatePost(ctx context.Context, in *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	l := logic.NewCreatePostLogic(ctx, s.svcCtx)
	return l.CreatePost(in)
}

func (s *PostServer) DeletePost(ctx context.Context, in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	l := logic.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}

func (s *PostServer) GetPostDetail(ctx context.Context, in *post.GetPostDetailRequest) (*post.GetPostDetailResponse, error) {
	l := logic.NewGetPostDetailLogic(ctx, s.svcCtx)
	return l.GetPostDetail(in)
}

func (s *PostServer) IncreaseUserView(ctx context.Context, in *post.IncreaseUserViewRequest) (*post.IncreaseUserViewResponse, error) {
	l := logic.NewIncreaseUserViewLogic(ctx, s.svcCtx)
	return l.IncreaseUserView(in)
}
