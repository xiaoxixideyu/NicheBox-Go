package model

import (
	"context"
	"nichebox/service/post/model/dto"
	"time"
)

type PostInterface interface {
	CreatePost(post *Post) error
	DeletePost(post *Post) error
	GetPostByID(postID int64) (*Post, error)
	GetModifiedPosts(from time.Time, to time.Time) ([]*dto.NewPostInfo, []*dto.DeletedPostInfo, error)
}

type PostCacheInterface interface {
	GetUserView(ctx context.Context, postID int64) (int64, error)
	IncrUserView(ctx context.Context, postID int64, visitorID int64) error
	BloomCheckPostExists(ctx context.Context, postID int64) (bool, error)
	BloomAddPost(ctx context.Context, postID int64) error
}
