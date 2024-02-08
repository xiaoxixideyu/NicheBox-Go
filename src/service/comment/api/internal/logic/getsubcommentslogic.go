package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCommentsLogic {
	return &GetSubCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubCommentsLogic) GetSubComments(req *types.GetSubCommentsRequest) (resp *types.GetSubCommentsResponse, err error) {
	// todo: uid check if thumbs up
	_, err = l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := comment.GetSubCommentsRequest{
		RootID: req.RootID,
		Page:   int32(req.Page),
		Size:   int32(req.Size),
	}

	out, err := l.svcCtx.CommentRpc.GetSubComments(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[Rpc] Get sub comments error", err)
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.OutOfRange {
				return nil, errors.New(http.StatusBadRequest, err.Error())
			}
		}
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	comments := make([]*types.CommentInfo, 0, len(out.SubComments))
	for _, c := range out.SubComments {
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
		comments = append(comments, &info)
	}

	return &types.GetSubCommentsResponse{SubComments: comments}, nil
}
