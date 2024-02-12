package cronx

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/task/rpc/internal/svc"
)

type UpdateUserView struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserView(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserView {
	return &UpdateUserView{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserView) AddUpdateUserViewTask() {
	_, err := l.svcCtx.Cron.AddFunc("CRON_TZ=Asia/Shanghai 0 30 2 ? * *", func() {
		// todo: 分布式锁
		l.svcCtx.UpdateUserViewCond.Broadcast()
	})
	if err != nil {
		l.Logger.Errorf("[Cron] New update uv task error", err)
	}
}
