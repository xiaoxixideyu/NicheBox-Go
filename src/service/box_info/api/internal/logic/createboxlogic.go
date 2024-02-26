package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"nichebox/service/box_info/api/internal/svc"
	"nichebox/service/box_info/api/internal/types"
	"nichebox/service/box_info/rpc/pb/boxinfo"
	"nichebox/service/box_user/rpc/pb/boxuser"
	"nichebox/service/user/rpc/pb/user"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
)

type CreateBoxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBoxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBoxLogic {
	return &CreateBoxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBoxLogic) CreateBox(req *types.CreateBoxRequest) (resp *types.CreateBoxResponse, err error) {
	// get uid
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "鉴权无效")
	}

	// get bid
	bidRes, err := l.svcCtx.BoxInfoRpc.CreateBid(l.ctx, &boxinfo.CreateBidRequest{})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "生成bid失败")
	}

	// Verify that the uid is available
	userCheck, err := l.svcCtx.UserRpc.CheckUid(l.ctx, &user.CheckUidRequest{
		Uid: uid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "refreshtoken 服务出错: 1")
	}
	if !userCheck.Exists {
		return nil, errors.New(http.StatusUnauthorized, "无效身份")
	}

	// Get BoxInfoRpc BuildTarget
	boxInfoRpcBusiServer, err := l.svcCtx.Config.BoxInfoRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(http.StatusContinue, "盒子创建异常")
	}

	// Get BoxUserRpc BuildTarget
	boxUserRpcBusiServer, err := l.svcCtx.Config.BoxUserRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(http.StatusContinue, "盒子创建异常")
	}

	// etcd registration address of dtm service
	dtmServer := "etcd://etcd:2379/dtmservice"
	// create a gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// create a saga protocol transation
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).Add(boxInfoRpcBusiServer+"/boxinfoclient.BoxInfo/CreateBox", boxInfoRpcBusiServer+"/boxinfoclient.BoxInfo/CreateBoxRevert", &boxinfo.CreateBoxRequest{
		Bid:          bidRes.Bid,
		Name:         req.Name,
		Introduction: req.Introduction,
	}).Add(boxUserRpcBusiServer+"/boxuserclient.BoxUser/AddOwner", boxUserRpcBusiServer+"/boxuserclient.BoxUser/AddOwnerRevert", &boxuser.AddOwnerRequest{
		Uid: uid,
		Bid: bidRes.Bid,
	})

	// commit transation
	if err := saga.Submit(); err != nil {
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &types.CreateBoxResponse{
		BoxId: strconv.FormatInt(bidRes.Bid, 10),
	}, nil
}
