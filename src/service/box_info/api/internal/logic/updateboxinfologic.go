package logic

import (
	"context"
	"net/http"
	"strconv"

	"nichebox/service/box_info/api/internal/common"
	"nichebox/service/box_info/api/internal/svc"
	"nichebox/service/box_info/api/internal/types"
	"nichebox/service/box_info/rpc/pb/boxinfo"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type UpdateBoxInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBoxInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBoxInfoLogic {
	return &UpdateBoxInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBoxInfoLogic) UpdateBoxInfo(req *types.UpdateBoxInfoRequest) (resp *types.UpdateBoxInfoResponse, err error) {
	// Get uid
	uid, err := common.GetAndCheckUid(l.ctx, l.svcCtx.UserRpc)
	if err != nil {
		return nil, err
	}

	// Get bid
	bid, err := strconv.ParseInt(req.BoxId, 10, 64)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "box_id 无效")
	}

	// Check the role of user
	role, err := l.svcCtx.BoxUserRpc.GetRole(l.ctx, &boxuser.GetRoleRequest{
		Bid: bid,
		Uid: uid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "update box info 服务出错: 1")
	}
	if !role.Exist || role.Role != boxuser.UserRole_Owner {
		return nil, errors.New(http.StatusForbidden, "你不是拥有者")
	}

	_, err = l.svcCtx.BoxInfoRpc.UpdateBoxInfo(l.ctx, &boxinfo.UpdateBoxInfoRequest{
		Bid:          bid,
		Name:         req.Name,
		Introduction: req.Introduction,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "update box info 服务出错: 2")
	}

	return &types.UpdateBoxInfoResponse{}, nil
}
