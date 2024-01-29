package mqs

import (
	"context"
	"encoding/json"
	"nichebox/service/task/model/dto"
	"nichebox/service/task/rpc/internal/svc"
	"time"
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

func (l *UpdateUserView) Consume(key, val string) error {
	taskModel := dto.UpdateUserViewTask{}
	err := json.Unmarshal([]byte(val), &taskModel)
	if err != nil {
		return err
	}

	createDate, err := time.Parse(time.DateOnly, taskModel.CreateDate)
	if err != nil {
		return err
	}
	nowDate, err := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	if err != nil {
		return err
	}

	if !nowDate.After(createDate) {
		l.svcCtx.UpdateUserViewCond.Wait()
	}

	// todo: batch update
	uv, err := l.svcCtx.TaskCacheInterface.GetUserView(l.ctx, taskModel.PostID)
	if err != nil {
		return err
	}
	err = l.svcCtx.TaskInterface.UpdatePostUserView(taskModel.PostID, uv)
	if err != nil {
		return err
	}

	return nil
}
