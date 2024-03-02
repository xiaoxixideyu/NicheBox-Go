package model

import "context"

type RelationInterface interface {
	AddFollow(uid int64, fid int64) error
	RemoveFollow(uid int64, fid int64) error
	GetFollowers(uid int64, page, size int, order string) ([]*Relation, error)
	GetFollowings(uid int64, page, size int, order string) ([]*Relation, error)
	GetFollowerCount(uid int64) (int, error)
	GetFollowingCount(uid int64) (int, error)
	GetRelationship(uid int64, fid int64) (*Relation, error)
}

type RelationCacheInterface interface {
	GetRelationshipCtx(ctx context.Context, uid int64, fid int64) (*CacheRelationshipAttribute, error)
	GetAllRelationshipsCtx(ctx context.Context, uid int64) ([]*CacheRelationshipAttribute, error)
	RemoveRelationshipsCtx(ctx context.Context, uid int64) error
	BatchAddRelationshipsCtx(ctx context.Context, uid int64, attrs []*CacheRelationshipAttribute) error
	GetRelationCountCtx(ctx context.Context, uid int64) (followers, followings int, err error)
	DeleteRelationCountCtx(ctx context.Context, uid int64) error
	SetRelationCountCtx(ctx context.Context, uid int64, followers, followings int) error
	BloomAddRelationCtx(ctx context.Context, uid int64, fid int64) error
}
