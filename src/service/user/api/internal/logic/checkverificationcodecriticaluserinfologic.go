package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/user/api/internal/common"
	redisBiz "nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/pb/user"
	"strconv"
	"strings"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckVerificationCodeCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckVerificationCodeCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVerificationCodeCriticalUserInfoLogic {
	return &CheckVerificationCodeCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckVerificationCodeCriticalUserInfoLogic) CheckVerificationCodeCriticalUserInfo(req *types.CheckVerificationCodeCriticalUserInfoRequest) (resp *types.CheckVerificationCodeCriticalUserInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	key := redisBiz.KeyPrefixUser + redisBiz.KeyCriticalCode + req.Email
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

	// delete original verification code
	inRemove := user.RemoveVerificationCodeRequest{
		Key: key,
	}
	_, _ = l.svcCtx.UserRpc.RemoveVerificationCode(l.ctx, &inRemove)

	// set critical token
	tokenKey := redisBiz.KeyPrefixUser + redisBiz.KeyCriticalToken + strconv.FormatInt(uid, 10)
	tokenCode := common.GenerateVerificationCode()
	inSet := user.SetVerificationCodeRequest{
		Key:        tokenKey,
		Val:        tokenCode,
		Expiration: common.VERIFICATIONCODEEXPIRATION,
	}
	_, err = l.svcCtx.UserRpc.SetVerificationCode(l.ctx, &inSet)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误: 1")
	}

	return &types.CheckVerificationCodeCriticalUserInfoResponse{VerificationCode: tokenCode}, nil
}
