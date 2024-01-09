package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/email/api/internal/svc"
	"nichebox/service/email/api/internal/types"
	"nichebox/service/email/rpc/pb/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerificationCodeRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerificationCodeRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeRegisterLogic {
	return &SendVerificationCodeRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerificationCodeRegisterLogic) SendVerificationCodeRegister(req *types.SendVerificationCodeRegisterRequest) (resp *types.SendVerificationCodeRegisterResponse, err error) {
	// todo: 测试发现send过程很可能会比较耗时，应该投放到消息队列异步进行

	in := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        generateVerificationCode(),
		Type:        TYPEREGISTER,
	}

	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SendVerificationCodeRegisterResponse{}, nil
}
