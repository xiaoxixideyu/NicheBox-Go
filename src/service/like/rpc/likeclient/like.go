// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package likeclient

import (
	"context"

	"nichebox/service/like/rpc/pb/like"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelThumbsResponse    = like.CancelThumbsResponse
	CancelThumbsUpRequest   = like.CancelThumbsUpRequest
	ThumbsUpCountRequest    = like.ThumbsUpCountRequest
	ThumbsUpCountResponse   = like.ThumbsUpCountResponse
	ThumbsUpHistoryRequest  = like.ThumbsUpHistoryRequest
	ThumbsUpHistoryResponse = like.ThumbsUpHistoryResponse
	ThumbsUpRequest         = like.ThumbsUpRequest
	ThumbsUpResponse        = like.ThumbsUpResponse

	Like interface {
		ThumbsUp(ctx context.Context, in *ThumbsUpRequest, opts ...grpc.CallOption) (*ThumbsUpResponse, error)
		CancelThumbsUp(ctx context.Context, in *CancelThumbsUpRequest, opts ...grpc.CallOption) (*CancelThumbsResponse, error)
		ThumbsUpCount(ctx context.Context, in *ThumbsUpCountRequest, opts ...grpc.CallOption) (*ThumbsUpCountResponse, error)
		ThumbsUpHistory(ctx context.Context, in *ThumbsUpHistoryRequest, opts ...grpc.CallOption) (*ThumbsUpHistoryResponse, error)
	}

	defaultLike struct {
		cli zrpc.Client
	}
)

func NewLike(cli zrpc.Client) Like {
	return &defaultLike{
		cli: cli,
	}
}

func (m *defaultLike) ThumbsUp(ctx context.Context, in *ThumbsUpRequest, opts ...grpc.CallOption) (*ThumbsUpResponse, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.ThumbsUp(ctx, in, opts...)
}

func (m *defaultLike) CancelThumbsUp(ctx context.Context, in *CancelThumbsUpRequest, opts ...grpc.CallOption) (*CancelThumbsResponse, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.CancelThumbsUp(ctx, in, opts...)
}

func (m *defaultLike) ThumbsUpCount(ctx context.Context, in *ThumbsUpCountRequest, opts ...grpc.CallOption) (*ThumbsUpCountResponse, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.ThumbsUpCount(ctx, in, opts...)
}

func (m *defaultLike) ThumbsUpHistory(ctx context.Context, in *ThumbsUpHistoryRequest, opts ...grpc.CallOption) (*ThumbsUpHistoryResponse, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.ThumbsUpHistory(ctx, in, opts...)
}
