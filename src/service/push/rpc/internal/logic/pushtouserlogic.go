package logic

import (
	"context"
	"encoding/json"
	"nichebox/service/long-connection/model/dto"

	"nichebox/service/push/rpc/internal/svc"
	"nichebox/service/push/rpc/pb/push"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushToUserLogic {
	return &PushToUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushToUserLogic) PushToUser(in *push.PushToUserRequest) (*push.PushToUserResponse, error) {
	msg := dto.PushToUserMessage{
		Uid:  in.Uid,
		Data: in.Data,
	}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.KqPushToUserPusherClient.Push(string(bytes))
	if err != nil {
		// todo: 本地消息表
		return nil, err
	}

	return &push.PushToUserResponse{}, nil
}
