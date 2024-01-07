package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"nichebox/service/user/api/internal/common"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReqeust) (resp *types.RefreshTokenResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "鉴权无效")
	}

	userCheck, err := l.svcCtx.UserRpc.CheckUid(l.ctx, &user.CheckUidRequest{
		Uid: uid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "refreshtoken 服务出错: 1")
	}

	if !userCheck.Exists {
		return nil, errors.New(http.StatusUnauthorized, "无效身份")
	}

	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, refreshToken, err := common.CreateTokenAndRefreshToken(uid, accessExpire, refreshExpire, accessSecret)
	if err != nil {
		return nil, err
	}

	return &types.RefreshTokenResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
