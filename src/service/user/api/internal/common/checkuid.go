package common

import (
	"context"
	"encoding/json"
	"net/http"
	"nichebox/service/user/rpc/pb/user"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/x/errors"
)

func CheckUid(ctx context.Context, userRpc userclient.User) error {
	uid, err := ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return errors.New(http.StatusUnauthorized, "uid无效")
	}

	userCheck, err := userRpc.CheckUid(ctx, &user.CheckUidRequest{
		Uid: uid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError, "UploadAvatar 服务出错: 1")
	}
	if !userCheck.Exists {
		return errors.New(http.StatusUnauthorized, "无效身份")
	}
	return nil
}
