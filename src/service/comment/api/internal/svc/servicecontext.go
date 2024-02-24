package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/comment/api/internal/config"
	"nichebox/service/comment/rpc/commentclient"
	"nichebox/service/comment/rpc/pb/comment"
	"nichebox/service/post/rpc/pb/post"
	"nichebox/service/post/rpc/postclient"
	"nichebox/service/push/rpc/pb/push"
	"nichebox/service/push/rpc/pushclient"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc comment.CommentClient
	PostRpc    post.PostClient
	PushRpc    push.PushClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
		PostRpc:    postclient.NewPost(zrpc.MustNewClient(c.PostRpc)),
		PushRpc:    pushclient.NewPush(zrpc.MustNewClient(c.PushRpc)),
	}
}
