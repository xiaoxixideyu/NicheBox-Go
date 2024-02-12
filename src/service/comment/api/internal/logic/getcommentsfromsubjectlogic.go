package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/common/biz"
	"nichebox/service/comment/rpc/pb/comment"

	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsFromSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsFromSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsFromSubjectLogic {
	return &GetCommentsFromSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsFromSubjectLogic) GetCommentsFromSubject(req *types.GetCommentsFromSubjectRequest) (resp *types.GetCommentsFromSubjectResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	if !biz.CheckIfCommentOrderValid(req.Order) {
		return nil, errors.New(http.StatusBadRequest, "order invalid")
	}

	inGetComment := comment.GetCommentsFromSubjectRequest{
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
		Page:        int32(req.Page),
		Size:        int32(req.Size),
		Order:       req.Order,
		Uid:         uid,
	}

	out, err := l.svcCtx.CommentRpc.GetCommentsFromSubject(l.ctx, &inGetComment)
	if err != nil {
		l.Logger.Errorf("[RPC] Get comments from subject error", err)
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
	}

	comments := make([]*types.CommentInfo, 0, len(out.Comments))
	for _, c := range out.Comments {
		info := types.CommentInfo{
			CommentID:          c.CommentID,
			SubjectID:          c.SubjectID,
			RootID:             c.RootID,
			ParentID:           c.ParentID,
			DialogID:           c.DialogID,
			OwnerID:            c.OwnerID,
			LikeCount:          int(c.LikeCount),
			ThumbsUp:           false,
			Floor:              int(c.Floor),
			CreateTime:         c.CreateTime,
			InnerFloorCount:    int(c.InnerFloorCount),
			InnerFloorComments: nil,
			Content:            c.Content,
		}
		// inner floor
		if c.InnerFloorComments != nil {
			info.InnerFloorComments = make([]*types.CommentInfo, 0, len(c.InnerFloorComments))
			for _, innerC := range c.InnerFloorComments {
				innerInfo := types.CommentInfo{
					CommentID:          innerC.CommentID,
					SubjectID:          innerC.SubjectID,
					RootID:             innerC.RootID,
					ParentID:           innerC.ParentID,
					DialogID:           innerC.DialogID,
					OwnerID:            innerC.OwnerID,
					LikeCount:          int(innerC.LikeCount),
					ThumbsUp:           false,
					Floor:              int(innerC.Floor),
					CreateTime:         innerC.CreateTime,
					InnerFloorCount:    int(innerC.InnerFloorCount),
					InnerFloorComments: nil,
					Content:            innerC.Content,
				}
				info.InnerFloorComments = append(info.InnerFloorComments, &innerInfo)
			}
		}

		comments = append(comments, &info)
	}

	return &types.GetCommentsFromSubjectResponse{Comments: comments}, nil
}
