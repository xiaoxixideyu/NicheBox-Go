package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/feed/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func (r RedisInterface) BatchDeliverFeedsToOutboxesCtx(ctx context.Context, messageID int64, messageType int, publishTime time.Time, followers []int64) error {
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		for _, id := range followers {
			key := KeyPrefixFeed + KeyOutbox + strconv.FormatInt(int64(messageType), 10) + Separator + strconv.FormatInt(id, 10)
			member := redis.Z{
				Score:  float64(publishTime.Unix()),
				Member: messageID,
			}
			pipeliner.ZAdd(ctx, key, &member)
		}
		_, err := pipeliner.Exec(ctx)
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.FeedCacheInterface, error) {
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
