package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/comment/model"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectByMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectByMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectByMessageLogic {
	return &GetSubjectByMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubjectByMessageLogic) GetSubjectByMessage(in *comment.GetSubjectByMessageRequest) (*comment.GetSubjectByMessageResponse, error) {
	subjectInfoStr, err := l.svcCtx.CommentCacheInterface.GetSubjectInfoByMessageCtx(l.ctx, in.MessageID, int(in.MessageType))
	needQueryDB := false
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get subject info by message error", err)
		}
		needQueryDB = true
	}
	if !needQueryDB {
		subjectInfo := model.Subject{}
		err := json.Unmarshal([]byte(subjectInfoStr), &subjectInfo)
		if err != nil {
			needQueryDB = true
		} else {
			return &comment.GetSubjectByMessageResponse{SubjectID: int64(subjectInfo.ID), CommentCount: int32(subjectInfo.CommentCount)}, nil
		}
	}

	subject := model.Subject{
		TypeID:    int(in.MessageType),
		MessageID: in.MessageID,
	}
	err = l.svcCtx.CommentInterface.FirstOrCreateSubject(&subject)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get subject from db error", err)
		return nil, err
	}
	// rewrite cache
	l.svcCtx.CommentCacheInterface.SetSubjectInfoCtx(l.ctx, &subject)
	return &comment.GetSubjectByMessageResponse{SubjectID: int64(subject.ID), CommentCount: int32(subject.CommentCount)}, nil
}
