package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	redisBiz "nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/pb/user"
	"strconv"
	"strings"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCriticalUserInfoLogic {
	return &SetCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCriticalUserInfoLogic) SetCriticalUserInfo(req *types.SetCriticalUserInfoRequest) (resp *types.SetCriticalUserInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	key := redisBiz.KeyPrefixUser + redisBiz.KeyCriticalToken + strconv.FormatInt(uid, 10)
	inGet := user.GetVerificationCodeRequest{
		Key: key,
	}
	out, err := l.svcCtx.UserRpc.GetVerificationCode(l.ctx, &inGet)

	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "验证码过期")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	if strings.ToUpper(req.Code) != out.Val {
		return nil, errors.New(http.StatusBadRequest, "验证码错误")
	}

	in := user.SetCriticalUserInfoRequest{
		Uid:      uid,
		Password: req.Password,
	}
	_, err = l.svcCtx.UserRpc.SetCriticalUserInfo(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	// delete original verification token
	inRemove := user.RemoveVerificationCodeRequest{
		Key: key,
	}
	_, _ = l.svcCtx.UserRpc.RemoveVerificationCode(l.ctx, &inRemove)

	return &types.SetCriticalUserInfoResponse{}, nil
}
