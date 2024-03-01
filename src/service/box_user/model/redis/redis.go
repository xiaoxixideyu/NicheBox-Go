package redis

import (
	"encoding/json"
	"nichebox/service/box_user/model"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisInterface struct {
	rds *redis.Redis
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.BoxUserCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	r := &RedisInterface{
		rds: redis.MustNewRedis(conf),
	}
	return r, nil
}

func (r *RedisInterface) SetBoxUser(boxUser *model.BoxUserCache, expireTime int) error {
	bid := strconv.FormatInt(boxUser.Bid, 10)
	uid := strconv.FormatInt(boxUser.Uid, 10)

	key := KeyBoxUser + "." + bid + "." + uid
	val := &model.BoxUserCacheVal{
		Exist: boxUser.Exist,
		Role:  boxUser.Role,
	}
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if err := r.rds.Set(key, string(bytes)); err != nil {
		return err
	}

	if expireTime != -1 {
		if err := r.rds.Expire(key, expireTime); err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisInterface) GetBoxUser(bid, uid int64, expireTime int) (*model.BoxUserCache, error) {
	bidStr := strconv.FormatInt(bid, 10)
	uidStr := strconv.FormatInt(uid, 10)

	key := KeyBoxUser + "." + bidStr + "." + uidStr
	valStr, err := r.rds.Get(key)
	if err != nil {
		return nil, err
	}

	val := &model.BoxUserCacheVal{}
	if err := json.Unmarshal([]byte(valStr), val); err != nil {
		return nil, err
	}

	if expireTime != -1 {
		if err := r.rds.Expire(key, expireTime); err != nil {
			return nil, err
		}
	}

	return &model.BoxUserCache{
		Bid:   bid,
		Uid:   uid,
		Exist: val.Exist,
		Role:  val.Role,
	}, nil
}

func (r *RedisInterface) RemoveBoxUser(bid, uid int64) error {
	bidStr := strconv.FormatInt(bid, 10)
	uidStr := strconv.FormatInt(uid, 10)

	key := KeyBoxUser + "." + bidStr + "." + uidStr
	if _, err := r.rds.Del(key); err != nil {
		return err
	}
	return nil
}
