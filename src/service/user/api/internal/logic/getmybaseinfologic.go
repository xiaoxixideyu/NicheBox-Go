package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetMyBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyBaseInfoLogic {
	return &GetMyBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyBaseInfoLogic) GetMyBaseInfo(req *types.GetMyBaseInfoRequest) (resp *types.GetMyBaseInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	res, err := l.svcCtx.UserRpc.GetUserBaseInfo(l.ctx, &user.GetUserBaseInfoRequest{
		Uid: uid,
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

	return &types.GetMyBaseInfoResponse{
		Uid:          strconv.FormatInt(uid, 10),
		UserName:     res.Username,
		Introduction: res.Introduction,
	}, nil
}
