package cronx

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/box-content/rpc/boxcontent"
	"nichebox/service/post/rpc/pb/post"
	"nichebox/service/task/rpc/internal/svc"
	"time"
)

type UpdateBoxContent struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	lastTime time.Time
	logx.Logger
}

func NewUpdateBoxContent(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBoxContent {
	startTime, _ := time.ParseInLocation(time.DateTime, time.DateTime, time.Local)
	return &UpdateBoxContent{
		ctx:      ctx,
		svcCtx:   svcCtx,
		lastTime: startTime,
		Logger:   logx.WithContext(ctx),
	}
}

func (l *UpdateBoxContent) AddUpdateBoxContentTask() {
	_, err := l.svcCtx.Cron.AddFunc("0/10 * * * * *", func() {
		from := l.lastTime.Format(time.DateTime)
		now := time.Now()
		to := now.Format(time.DateTime)

		inGet := post.GetModifiedPostsRequest{
			FromTime: from,
			ToTime:   to,
		}
		outGet, err := l.svcCtx.PostRpc.GetModifiedPosts(l.ctx, &inGet)
		if err != nil {
			l.Logger.Errorf("[Rpc][Cron] Get modified posts info error", err)
			return
		}
		// update new posts
		newPosts := make([]*boxcontent.ModifiedPostInfo, 0, 0)
		for _, p := range outGet.NewPosts {
			newPost := boxcontent.ModifiedPostInfo{
				PostID:    p.PostID,
				Time:      p.Time,
				BoxID:     p.BoxID,
				InfoCount: p.InfoCount,
			}
			newPosts = append(newPosts, &newPost)
		}
		inNew := boxcontent.UpdateNewPostsRequest{NewPosts: newPosts}
		_, err = l.svcCtx.BoxContentRpc.UpdateNewPosts(l.ctx, &inNew)
		if err != nil {
			l.Logger.Errorf("[Rpc][Cron] Update new posts info to box content error", err)
			return
		}
		// update deleted posts
		deletedPosts := make([]*boxcontent.ModifiedPostInfo, 0, 0)
		for _, p := range outGet.DeletedPosts {
			deletedPost := boxcontent.ModifiedPostInfo{
				PostID:    p.PostID,
				Time:      p.Time,
				BoxID:     p.BoxID,
				InfoCount: p.InfoCount,
			}
			deletedPosts = append(deletedPosts, &deletedPost)
		}
		inDeleted := boxcontent.UpdateDeletedPostsRequest{DeletedPosts: deletedPosts}
		_, err = l.svcCtx.BoxContentRpc.UpdateDeletedPosts(l.ctx, &inDeleted)
		if err != nil {
			l.Logger.Errorf("[Rpc][Cron] Update deleted posts info to box content error", err)
			return
		}

		l.lastTime = now
	})
	if err != nil {
		l.Logger.Errorf("[Cron] New update box content task error", err)
	}
}
