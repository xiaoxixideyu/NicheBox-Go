package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/feed/model/dto"
	"nichebox/service/feed/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"
)

type DeliverFeedToOutbox struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverFeedToOutbox(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverFeedToOutbox {
	return &DeliverFeedToOutbox{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeliverFeedToOutbox) Consume(key, value string) error {
	msg := dto.DeliverFeedToOutboxMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal failed, err:", err)
		return err
	}

	in := relation.GetFollowersRequest{
		Uid:  msg.AuthorID,
		Page: 1,
		Size: -1,
	}
	out, err := l.svcCtx.RelationRpc.GetFollowers(l.ctx, &in)
	ids := make([]int64, 0, len(out.Followers))
	for _, f := range out.Followers {
		ids = append(ids, f.Fid)
	}

	return nil
}
