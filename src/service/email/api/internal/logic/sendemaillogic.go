package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"math/rand"
	"net/http"
	"nichebox/service/email/rpc/pb/email"
	"time"

	"nichebox/service/email/api/internal/svc"
	"nichebox/service/email/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	VERIFICATIONCODELENGTH = 6
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendVerificationCode(req *types.SendVerificationCodeRequest) (resp *types.SendVerificationCodeResponse, err error) {
	// todo: 测试发现send过程很可能会比较耗时，应该投放到消息队列异步进行

	in := email.SendVerificationCodeRequest{
		Destination: req.Destination,
		Code:        generateVerificationCode(),
	}

	_, err = l.svcCtx.EmailRpc.SendVerificationCode(l.ctx, &in)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.SendVerificationCodeResponse{}, nil
}

func generateVerificationCode() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < VERIFICATIONCODELENGTH; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
