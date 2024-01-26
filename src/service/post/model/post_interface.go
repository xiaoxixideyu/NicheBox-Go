package model

import "context"

type PostInterface interface {
	CreatePost(post *Post) error
	DeletePost(post *Post) error
	GetPostByID(postID int64) (*Post, error)
}

type PostCacheInterface interface {
	GetUserView(ctx context.Context, postID int64) (int64, error)
	IncrUserView(ctx context.Context, postID int64, visitorID int64) error
	BloomCheckPostExists(ctx context.Context, postID int64) (bool, error)
	BloomAddPost(ctx context.Context, postID int64) error
}
