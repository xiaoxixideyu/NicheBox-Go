package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/comment/rpc/pb/comment"

	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectInfoByMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubjectInfoByMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectInfoByMessageLogic {
	return &GetSubjectInfoByMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubjectInfoByMessageLogic) GetSubjectInfoByMessage(req *types.GetSubjectInfoByMessageRequest) (resp *types.GetSubjectInfoByMessageResponse, err error) {
	in := comment.GetSubjectByMessageRequest{
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
	}

	out, err := l.svcCtx.CommentRpc.GetSubjectByMessage(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get subject info error", err)
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.GetSubjectInfoByMessageResponse{
		SubjectID:    out.SubjectID,
		CommentCount: int(out.CommentCount),
	}, nil
}
