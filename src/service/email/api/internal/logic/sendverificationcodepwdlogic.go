package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/email/rpc/pb/email"

	"nichebox/service/email/api/internal/svc"
	"nichebox/service/email/api/internal/types"

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

	in := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        generateVerificationCode(),
		Type:        TYPEPWD,
	}

	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}
	return &types.SendVerificationCodePWDResponse{}, nil
}
