package common

import (
	"encoding/json"
	"log"
	"nichebox/service/box_user/model/dto"
	"nichebox/service/box_user/rpc/internal/svc"
)

func PushRemoveCacheBoxUserExist(bid, uid int64, svcCtx *svc.ServiceContext) error {
	msg := &dto.RemoveCacheBoxUserMessagw{
		Bid: bid,
		Uid: uid,
	}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		log.Printf("[Json][Producer] Json marshal error: %v\n", err)
		return err
	}
	err = svcCtx.KqRemoveCacheBoxUserPusherClient.Push(string(bytes))
	if err != nil {
		log.Printf("[Kafka][Producer] MQ push error: %v\n", err)
		return err
	}
	return nil
}
