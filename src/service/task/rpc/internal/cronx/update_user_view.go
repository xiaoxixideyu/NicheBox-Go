package cronx

import (
	"context"
	"nichebox/service/task/rpc/internal/svc"
)

type UpdateUserView struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserView(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserView {
	return &UpdateUserView{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserView) AddUpdateUserViewTask() {
	l.svcCtx.Cron.AddFunc("CRON_TZ=Asia/Shanghai 30 02 * * *", func() {
		// todo: 分布式锁
		l.svcCtx.UpdateUserViewCond.Broadcast()
	})
}
