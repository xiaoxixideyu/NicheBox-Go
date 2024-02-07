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

type GetSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectLogic {
	return &GetSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubjectLogic) GetSubject(in *comment.GetSubjectRequest) (*comment.GetSubjectResponse, error) {
	subjectInfoStr, err := l.svcCtx.CommentCacheInterface.GetSubjectInfoBySubjectIDCtx(l.ctx, in.SubjectID)
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
			return &comment.GetSubjectResponse{
				MessageID:    subjectInfo.MessageID,
				MessageType:  int32(subjectInfo.TypeID),
				CommentCount: int32(subjectInfo.CommentCount),
			}, nil
		}
	}

	subject, err := l.svcCtx.CommentInterface.GetSubjectBySubjectID(in.SubjectID)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get subject by id error", err)
		return nil, err
	}

	// rewrite cache
	l.svcCtx.CommentCacheInterface.SetSubjectInfoCtx(l.ctx, subject)

	return &comment.GetSubjectResponse{
		MessageID:    subject.MessageID,
		MessageType:  int32(subject.TypeID),
		CommentCount: int32(subject.CommentCount),
	}, nil
}
