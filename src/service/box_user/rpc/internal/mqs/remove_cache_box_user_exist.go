package mqs

import (
	"context"
	"encoding/json"
	"nichebox/service/box_user/model/dto"
	"nichebox/service/box_user/rpc/internal/svc"
)

type RemoveCacheBoxUser struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCacheBoxUser(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCacheBoxUser {
	return &RemoveCacheBoxUser{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCacheBoxUser) Consume(key, value string) error {
	msg := dto.RemoveCacheBoxUserMessagw{}

	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		return err
	}

	if err := l.svcCtx.BoxUserCaCheInterface.RemoveBoxUser(msg.Bid, msg.Uid); err != nil {
		return err
	}

	return nil
}
