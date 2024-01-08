package logic

import (
	"context"

	"nichebox/service/email/rpc/internal/svc"
	"nichebox/service/email/rpc/pb/email"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	Mail   = 0b00000001
	Wechat = 0b00000010

	ServerMail = "detectionproject@163.com"
	ServerPwd  = "TQCKFVKLGZICJXSQ"
	Host       = "smtp.163.com"
	Port       = 25
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *email.SendEmailRequest) (*email.SendEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &email.SendEmailResponse{}, nil
}
