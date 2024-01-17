package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/post/rpc/pb/post"

	"nichebox/service/post/api/internal/svc"
	"nichebox/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	OperatorAuthor = "author"
	OperatorAdmin  = "admin"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.DeletePostResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	inGetPost := post.GetPostDetailRequest{PostID: req.PostID}
	outGetPost, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &inGetPost)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "该帖子不存在")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	if req.Operator == OperatorAdmin {
		// todo: 根据operator和uid判断，如果operator是admin，判断uid是否是当前post的管理员

	} else if req.Operator == OperatorAuthor {
		if outGetPost.AuthorID != uid {
			return nil, errors.New(http.StatusBadRequest, "你没有删除的权限")
		}
	}

	in := post.DeletePostRequest{PostID: req.PostID, Operator: req.Operator}
	_, err = l.svcCtx.PostRpc.DeletePost(l.ctx, &in)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "该晒图不存在")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}
	// todo: 根据operator调用通知服务

	return &types.DeletePostResponse{}, nil
}
