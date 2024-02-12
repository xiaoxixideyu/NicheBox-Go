package logic

import (
	"context"
	"time"

	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetModifiedPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetModifiedPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetModifiedPostsLogic {
	return &GetModifiedPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetModifiedPostsLogic) GetModifiedPosts(in *post.GetModifiedPostsRequest) (*post.GetModifiedPostsResponse, error) {

	from, err := time.ParseInLocation(time.DateTime, in.FromTime, time.Local)
	if err != nil {
		return nil, err
	}
	to, err := time.ParseInLocation(time.DateTime, in.ToTime, time.Local)
	if err != nil {
		return nil, err
	}
	newInfos, deletedInfos, err := l.svcCtx.PostInterface.GetModifiedPosts(from, to)
	if err != nil {
		return nil, err
	}

	rpcNewInfos := make([]*post.ModifiedPostInfo, 0, len(newInfos))
	for _, i := range newInfos {
		rpcInfo := post.ModifiedPostInfo{
			PostID:    i.PostID,
			Time:      i.CreatedAt.Format(time.DateTime),
			BoxID:     i.BoxID,
			InfoCount: int32(i.Count),
		}
		rpcNewInfos = append(rpcNewInfos, &rpcInfo)
	}

	rpcDeletedInfos := make([]*post.ModifiedPostInfo, 0, len(deletedInfos))
	for _, i := range deletedInfos {
		rpcInfo := post.ModifiedPostInfo{
			PostID:    i.PostID,
			Time:      i.DeletedAt.Format(time.DateTime),
			BoxID:     i.BoxID,
			InfoCount: int32(i.Count),
		}
		rpcDeletedInfos = append(rpcDeletedInfos, &rpcInfo)
	}
	return &post.GetModifiedPostsResponse{NewPosts: rpcNewInfos, DeletedPosts: rpcDeletedInfos}, nil
}
