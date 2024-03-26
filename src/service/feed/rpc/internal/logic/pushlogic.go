package logic

import (
	"context"
	"encoding/json"
	"nichebox/common/snowflake"
	"nichebox/service/feed/model"
	"nichebox/service/feed/model/dto"
	"time"

	"nichebox/service/feed/rpc/internal/svc"
	"nichebox/service/feed/rpc/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	return &PushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushLogic) Push(in *feed.PushRequest) (*feed.PushResponse, error) {
	publishTime, _ := time.ParseInLocation(time.DateTime, in.CreateTime, time.Local)
	id := snowflake.GenID()
	m := model.Feed{
		FeedID:      id,
		MessageID:   in.MessageID,
		TypeID:      int(in.MessageType),
		AuthorID:    in.Author,
		PublishTime: publishTime,
	}
	err := l.svcCtx.FeedInterface.AddFeed(&m)
	if err != nil {
		l.Logger.Errorf("[MySQL] Add feed failed, err:", err)
		return nil, err
	}
	// mq
	msg := dto.DeliverFeedToOutboxMessage{
		FeedID:      m.FeedID,
		AuthorID:    m.AuthorID,
		MessageID:   m.MessageID,
		MessageType: m.TypeID,
	}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		l.Logger.Errorf("[Json][Producer] Json marshal error", err)
	} else {
		err := l.svcCtx.KqDeliverFeedToOutboxPusherClient.Push(string(bytes))
		if err != nil {
			l.Logger.Errorf("[Json][Producer] Push failed, err:", err)
		}
	}

	return &feed.PushResponse{}, nil
}
