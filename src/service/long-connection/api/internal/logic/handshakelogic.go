package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/long-connection/rpc/pb/longConn"

	"nichebox/service/long-connection/api/internal/svc"
	"nichebox/service/long-connection/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandshakeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandshakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandshakeLogic {
	return &HandshakeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandshakeLogic) Handshake(req *types.HandShakeRequest) (resp *types.HandShakeResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	fmt.Printf("ua:%v, addr:%v\n", req.UserAgent, req.RemoteAddress)

	in := longConn.HandShakeRequest{
		Uid:       uid,
		UserAgent: req.UserAgent,
		Addr:      req.RemoteAddress,
	}

	out, err := l.svcCtx.LongConnRpc.HandShake(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Handshake error", err)
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
	}
	return &types.HandShakeResponse{Token: out.Token}, nil
}
