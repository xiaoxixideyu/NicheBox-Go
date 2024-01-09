package logic

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"

	"nichebox/service/email/rpc/internal/svc"
	"nichebox/service/email/rpc/pb/email"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	SUBJECT                    = "【小众盒】验证码"
	VERIFICATIONCODEEXPIRATION = 60 * 5

	TYPEREGISTER = "register"
	TYPEPWD      = "pwd"
	TYPECRITICAL = "critical"
)

type SendVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeLogic {
	return &SendVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendVerificationCodeLogic) SendVerificationCode(in *email.SendVerificationCodeRequest) (*email.SendVerificationCodeResponse, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", l.svcCtx.Config.ServerMail.Address)
	m.SetHeader("To", in.Destination)

	m.SetHeader("Subject", SUBJECT)

	m.SetBody("text/html", generateBody(in.Code, in.Type))

	dialer := gomail.NewDialer(
		l.svcCtx.Config.ServerMail.Host,
		l.svcCtx.Config.ServerMail.Port,
		l.svcCtx.Config.ServerMail.Address,
		l.svcCtx.Config.ServerMail.Password)
	err := dialer.DialAndSend(m)
	if err != nil {
		return nil, err
	}

	return &email.SendVerificationCodeResponse{}, nil
}

func generateBody(verificationCode string, bodyType string) string {
	var body string
	if bodyType == TYPEREGISTER {
		header := "欢迎加入小众盒，您的验证码如下"
		code := "<h2>" + verificationCode + "</h2>"
		expirationNotification := "请在5分钟内进行验证哦"
		body = fmt.Sprintf("%s<br>%s<br>%s", header, code, expirationNotification)

	} else if bodyType == TYPEPWD {
		header := "您正在尝试重设密码，验证码如下"
		code := "<h2>" + verificationCode + "</h2>"
		expirationNotification := "请在5分钟内进行验证"
		body = fmt.Sprintf("%s<br>%s<br>%s", header, code, expirationNotification)

	} else if bodyType == TYPECRITICAL {
		header := "您正在尝试修改重要信息，验证码如下"
		code := "<h2>" + verificationCode + "</h2>"
		expirationNotification := "请在5分钟内进行验证"
		body = fmt.Sprintf("%s<br>%s<br>%s", header, code, expirationNotification)

	}
	return body
}
