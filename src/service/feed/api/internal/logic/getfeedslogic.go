package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/feed/rpc/pb/feed"

	"nichebox/service/feed/api/internal/svc"
	"nichebox/service/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFeedsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedsLogic {
	return &GetFeedsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFeedsLogic) GetFeeds(req *types.GetFeedsRequest) (resp *types.GetFeedsResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := feed.PullRequest{
		Uid:  uid,
		Page: int32(req.Page),
		Size: int32(req.Size),
	}
	out, err := l.svcCtx.FeedRpc.Pull(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get feeds failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	feeds := make([]*types.FeedMessage, 0, len(out.FeedMessages))
	for _, f := range out.FeedMessages {
		fm := types.FeedMessage{
			MessageID:   f.MessageID,
			MessageType: int(f.MessageType),
			AuthorID:    f.Author,
			PublishTime: f.CreateTime,
		}
		feeds = append(feeds, &fm)
	}

	return &types.GetFeedsResponse{Feeds: feeds}, nil
}
