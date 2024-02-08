package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/common/biz"
	"nichebox/service/like/model"
	"nichebox/service/like/model/dto"
	"nichebox/service/like/rpc/internal/svc"
	"nichebox/service/like/rpc/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpLogic {
	return &ThumbsUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbsUpLogic) ThumbsUp(in *like.ThumbsUpRequest) (*like.ThumbsUpResponse, error) {
	likeModel := model.Like{
		Uid:       in.Uid,
		MessageID: in.MessageID,
		TypeID:    int(in.MessageType),
	}
	err := l.svcCtx.LikeInterface.CreateLikeAndUpdateLikeCountTX(&likeModel)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, status.Error(codes.AlreadyExists, "已经点赞过了")
		}

		return nil, err
	}

	// remove cache (cache aside)
	l.svcCtx.LikeCacheInterface.DeleteThumbsUpCountCtx(l.ctx, likeModel.MessageID, likeModel.TypeID)

	err = l.svcCtx.LikeCacheInterface.AddThumbsUpHistoryCtx(context.Background(), likeModel.MessageID, likeModel.TypeID, in.Uid)
	if err != nil {
		// err occurs, remove all redis data
		l.svcCtx.LikeCacheInterface.ClearAllThumbsUpHistoryCtx(context.Background(), likeModel.TypeID, in.Uid)
	}

	// push mq to update biz
	err = l.PushMQToUpdateBiz(int(in.MessageType), in.MessageID)
	if err != nil {
		// todo: save db
	}

	return &like.ThumbsUpResponse{}, nil
}

func (l *ThumbsUpLogic) PushMQToUpdateBiz(messageType int, messageID int64) error {
	if messageType == biz.MessageTypeComment {
		message := dto.UpdateCommentLikeCountMessage{CommentID: messageID, Delta: 1}
		bytes, err := json.Marshal(&message)
		if err != nil {
			l.Logger.Errorf("[Json][Producer] Json marshal error", err)
			return err
		}
		err = l.svcCtx.KqUpdateCommentLikeCountPusher.Push(string(bytes))
		if err != nil {
			l.Logger.Errorf("[Json][Producer] Kafka push error", err)
			return err
		}

	} else if messageType == biz.MessageTypePost {

	} else if messageType == biz.MessageTypeVideo {

	}

	return nil
}
