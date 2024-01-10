package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/email/rpc/pb/email"
	"nichebox/service/user/api/internal/common"
	"nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/pb/user"
	"strings"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerificationCodeForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerificationCodeForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeForgetPasswordLogic {
	return &SendVerificationCodeForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerificationCodeForgetPasswordLogic) SendVerificationCodeForgetPassword(req *types.SendVerificationCodeForgetPasswordRequest) (resp *types.SendVerificationCodeForgetPasswordResponse, err error) {
	code := common.GenerateVerificationCode()
	inEmail := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        code,
		Type:        common.TYPEFORGETPASSWORD,
	}

	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &inEmail)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	inUser := user.SetVerificationCodeRequest{
		Key:        redis.KeyPrefixUser + redis.KeyForgetPasswordCode + req.Destination,
		Val:        strings.ToUpper(code),
		Expiration: common.VERIFICATIONCODEEXPIRATION,
	}

	_, err = l.svcCtx.UserRpc.SetVerificationCode(context.Background(), &inUser)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SendVerificationCodeForgetPasswordResponse{}, nil
}
