package logic

import (
	"context"

	"nichebox/service/task/rpc/internal/svc"
	"nichebox/service/task/rpc/pb/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *task.CreateTaskRequest) (*task.CreateTaskResponse, error) {
	// todo: add your logic here and delete this line

	return &task.CreateTaskResponse{}, nil
}
