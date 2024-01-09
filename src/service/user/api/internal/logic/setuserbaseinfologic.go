package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetUserBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserBaseInfoLogic {
	return &SetUserBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserBaseInfoLogic) SetUserBaseInfo(req *types.SetUserBaseInfoRequest) (resp *types.SetUserBaseInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	_, err = l.svcCtx.UserRpc.SetUserBaseInfo(l.ctx, &userclient.SetUserBaseInfoRequest{
		Uid:          uid,
		Username:     req.Username,
		Introduction: req.Introduction,
	})
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if ok {
			code := grpcStatus.Code()
			if code == codes.NotFound {
				return nil, errors.New(http.StatusNotFound, "用户不存在")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "SetUserBaseInfo 服务出错: 1")
	}

	return &types.SetUserBaseInfoResponse{}, nil
}
