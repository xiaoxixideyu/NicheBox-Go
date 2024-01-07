package common

import (
	"net/http"
	"nichebox/common/jwtx"
	"time"

	"github.com/zeromicro/x/errors"
)

func CreateTokenAndRefreshToken(uid, accessExpire, refreshExpire int64, accessSecret string) (string, string, error) {
	now := time.Now().Unix()
	accessToken, err := jwtx.GetToken(accessSecret, now, accessExpire, uid)
	if err != nil {
		return "", "", errors.New(http.StatusInternalServerError, "token 生成失败: 1")
	}
	refreshToken, err := jwtx.GetToken(accessSecret, now, refreshExpire, uid)
	if err != nil {
		return "", "", errors.New(http.StatusInternalServerError, "token 生成失败: 2")
	}

	return accessToken, refreshToken, nil
}
