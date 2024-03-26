package logic

import (
	"context"
	"nichebox/service/relation/rpc/pb/relation"
	"time"

	"nichebox/service/feed/rpc/internal/svc"
	"nichebox/service/feed/rpc/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullLogic) Pull(in *feed.PullRequest) (*feed.PullResponse, error) {
	inF := relation.GetFollowingsRequest{
		Uid:  in.Uid,
		Page: 1,
		Size: -1,
	}
	outF, err := l.svcCtx.RelationRpc.GetFollowings(l.ctx, &inF)
	if err != nil {
		l.Logger.Errorf("[RPC] Get followings failed, err:", err)
		return nil, err
	}
	followings := make([]int64, 0, len(outF.Followings))
	for _, f := range outF.Followings {
		followings = append(followings, f.Fid)
	}
	feeds, err := l.svcCtx.FeedInterface.GetFeeds(followings, int(in.Page), int(in.Size))
	if err != nil {
		l.Logger.Errorf("[MySQL] Get feeds failed, err:", err)
		return nil, err
	}
	fs := make([]*feed.FeedMessage, 0, len(feeds))
	for _, fee := range feeds {
		f := feed.FeedMessage{
			MessageID:   fee.MessageID,
			MessageType: int32(fee.TypeID),
			Author:      fee.AuthorID,
			CreateTime:  fee.PublishTime.Format(time.DateTime),
		}
		fs = append(fs, &f)
	}

	return &feed.PullResponse{FeedMessages: fs}, nil
}
