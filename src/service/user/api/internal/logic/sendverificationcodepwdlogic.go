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

type SendVerificationCodePWDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerificationCodePWDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodePWDLogic {
	return &SendVerificationCodePWDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerificationCodePWDLogic) SendVerificationCodePWD(req *types.SendVerificationCodePWDRequest) (resp *types.SendVerificationCodePWDResponse, err error) {
	// todo: 测试发现send过程很可能会比较耗时，应该投放到消息队列异步进行

	code := common.GenerateVerificationCode()
	inEmail := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        code,
		Type:        common.TYPEPWD,
	}

	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &inEmail)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	inUser := user.SetVerificationCodeRequest{
		Key:        redis.KeyPrefixUser + redis.KeyPWDCode + req.Destination,
		Val:        strings.ToUpper(code),
		Expiration: common.VERIFICATIONCODEEXPIRATION,
	}

	_, err = l.svcCtx.UserRpc.SetVerificationCode(context.Background(), &inUser)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SendVerificationCodePWDResponse{}, nil
}
