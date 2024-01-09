package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	redisBiz "nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/pb/user"
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
	key := redisBiz.KeyPrefixUser + redisBiz.KeyCriticalCode + req.Email
	in := user.GetVerificationCodeRequest{
		Key: key,
	}
	out, err := l.svcCtx.UserRpc.GetVerificationCode(l.ctx, &in)
	
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

	return &types.CheckVerificationCodeCriticalUserInfoResponse{}, nil
}
