package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"nichebox/service/post/model/dto"
	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseUserViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseUserViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseUserViewLogic {
	return &IncreaseUserViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseUserViewLogic) IncreaseUserView(in *post.IncreaseUserViewRequest) (*post.IncreaseUserViewResponse, error) {
	// todo: 如果是游客的话，调用此RPC的visitorID应该换成ip地址
	err := l.svcCtx.PostCacheInterface.IncrUserView(l.ctx, in.PostID, in.VisitorID)
	if err != nil {
		fmt.Printf("redis error:%v\n", err)
		return nil, err
	}

	// 通过Kafka新增任务
	exists, err := l.svcCtx.PostCacheInterface.BloomCheckPostExists(l.ctx, in.PostID)
	if err != nil {
		fmt.Printf("redis2 error:%v\n", err)
	}

	if !exists {
		task := dto.UpdateUserViewTask{
			CreateDate: time.Now().Format(time.DateOnly),
			PostID:     in.PostID,
		}
		bytes, err := json.Marshal(task)
		if err != nil {
			fmt.Printf("json error:%v\n", err)
			return &post.IncreaseUserViewResponse{}, nil
		}
		name := l.svcCtx.KqUpdateUserViewPusherClient.Name()
		fmt.Printf("going kafka:%v\n name:%v", string(bytes), name)
		err = l.svcCtx.KqUpdateUserViewPusherClient.Push(string(bytes))

		if err != nil {
			fmt.Printf("kafka error:%v\n", err)
			return &post.IncreaseUserViewResponse{}, nil
		}
		l.svcCtx.PostCacheInterface.BloomAddPost(l.ctx, in.PostID)
	}
	return &post.IncreaseUserViewResponse{}, nil
}
