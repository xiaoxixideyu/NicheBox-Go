package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/common/biz"
	"nichebox/service/box-content/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func (r *RedisInterface) UpdateNewPostsCtx(ctx context.Context, boxID int64, infos []*model.ModifiedPostInfo) error {
	key := KeyPrefixBoxContent + KeyRank + KeyCreateTime + strconv.FormatInt(boxID, 10) + Separator + strconv.FormatInt(biz.MessageTypePost, 10)
	zs := make([]*redis.Z, 0, len(infos))
	for _, info := range infos {
		z := redis.Z{
			Score:  float64(info.Time.Unix()),
			Member: info.MessageID,
		}
		zs = append(zs, &z)
	}
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.ZAdd(ctx, key, zs...)
		_, err := pipeliner.Exec(ctx)
		return err
	})
	return err
}

func (r *RedisInterface) UpdateDeletedPostsCtx(ctx context.Context, boxID int64, infos []*model.ModifiedPostInfo) error {
	key := KeyPrefixBoxContent + KeyRank + KeyCreateTime + strconv.FormatInt(boxID, 10) + Separator + strconv.FormatInt(biz.MessageTypePost, 10)
	ids := make([]interface{}, 0, len(infos))
	for _, info := range infos {
		ids = append(ids, info.MessageID)
	}

	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.ZRem(ctx, key, ids...)
		_, err := pipeliner.Exec(ctx)
		return err
	})
	return err
}

func (r *RedisInterface) GetPostIDsCtx(ctx context.Context, boxID int64, page, size int, order string) ([]redis.Pair, error) {
	var err error
	var indexes []redis.Pair
	var key string

	start := (page - 1) * size
	stop := start + size - 1

	if order == biz.OrderByCreateTimeAsc {
		key = KeyPrefixBoxContent + KeyRank + KeyCreateTime + strconv.FormatInt(boxID, 10) + Separator + strconv.FormatInt(boxID, 10)
		indexes, err = r.rds.ZrangeWithScoresCtx(ctx, key, int64(start), int64(stop))

	} else if order == biz.OrderByCreateTimeDesc {
		key = KeyPrefixBoxContent + KeyRank + KeyCreateTime + strconv.FormatInt(boxID, 10) + Separator + strconv.FormatInt(boxID, 10)
		indexes, err = r.rds.ZrevrangeWithScoresCtx(ctx, key, int64(start), int64(stop))

		//} else if order == biz.OrderByView {
		//	key = KeyPrefixBoxContent + KeyRank + KeyViews + strconv.FormatInt(boxID, 10) + Separator + strconv.FormatInt(boxID, 10)
		//	indexes, err = r.rds.ZrevrangeWithScoresCtx(ctx, key, int64(start), int64(stop))
		//
	} else {
		return nil, biz.ErrRedisUnknownOrder
	}

	if err != nil {
		return nil, err
	}
	if len(indexes) == 0 {
		// check if page/size out of size or redis cache expired
		card, err := r.rds.ZcardCtx(ctx, key)
		if err != nil {
			return nil, err
		}
		if card == 0 {
			// cache expired or no data in this subject
			return nil, redis.Nil
		} else {
			// page/size out of size
			return nil, biz.ErrRedisOutOfBounds
		}
	}
	return indexes, nil
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.BoxContentCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	rds := redis.MustNewRedis(conf)
	r := &RedisInterface{
		rds: rds,
	}
	return r, nil
}
