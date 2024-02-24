package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/long-connection/model/dto"
	"nichebox/service/long-connection/rpc/internal/svc"
)

type PushToUser struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushToUser(ctx context.Context, svcCtx *svc.ServiceContext) *PushToUser {
	return &PushToUser{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushToUser) Consume(key, value string) error {
	msg := dto.PushToUserMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal failed, err:", err)
		return err
	}
	err = l.svcCtx.SessionManager.PushToUser(msg.Uid, msg.Data)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Push to user failed, err:", err)
		return err
	}
	return nil
}
