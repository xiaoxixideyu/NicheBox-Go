package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/email/rpc/pb/email"
	"nichebox/service/user/api/internal/common"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/pb/user"
)

type SendVerificationCodeCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerificationCodeCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeCriticalUserInfoLogic {
	return &SendVerificationCodeCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerificationCodeCriticalUserInfoLogic) SendVerificationCodeCriticalUserInfo(req *types.SendVerificationCodeCriticalUserInfoRequest) (resp *types.SendVerificationCodeCriticalUserInfoResponse, err error) {
	code := common.GenerateVerificationCode()

	inUser := user.SetVerificationCodeRequest{
		Key:        redis.KeyPrefixUser + redis.KeyCriticalCode + req.Destination,
		Val:        code,
		Expiration: common.VERIFICATIONCODEEXPIRATION,
	}

	_, err = l.svcCtx.UserRpc.SetVerificationCode(context.Background(), &inUser)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	inEmail := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        code,
		Type:        common.TYPECRITICAL,
	}
	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &inEmail)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SendVerificationCodeCriticalUserInfoResponse{}, nil
}
