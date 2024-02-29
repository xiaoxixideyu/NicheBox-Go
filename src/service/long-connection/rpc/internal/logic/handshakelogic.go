package logic

import (
	"context"
	"nichebox/service/long-connection/rpc/internal/routes"

	"nichebox/service/long-connection/rpc/internal/svc"
	"nichebox/service/long-connection/rpc/pb/longConn"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandShakeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandShakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandShakeLogic {
	return &HandShakeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandShakeLogic) HandShake(in *longConn.HandShakeRequest) (*longConn.HandShakeResponse, error) {
	// todo: delete
	in.Addr = "127.0.0.1"
	routes.GetRouter().RegisterAddr(in.Addr, in.Uid, in.UserAgent)
	// todo: token
	token := "token"
	// todo: server addr
	serverAddr := "127.0.0.1:10031"
	// todo: secret
	secret := "secret"

	return &longConn.HandShakeResponse{Token: token, ServerAddr: serverAddr, Secret: secret}, nil
}
