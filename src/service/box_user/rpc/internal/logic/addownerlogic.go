package logic

import (
	"context"
	"database/sql"

	"nichebox/service/box_user/model"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddOwnerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOwnerLogic {
	return &AddOwnerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOwnerLogic) AddOwner(in *boxuser.AddOwnerRequest) (*boxuser.AddOwnerRequest, error) {
	boxUser := &model.BoxUser{
		Bid:  in.Bid,
		Uid:  in.Uid,
		Role: int(boxuser.UserRole_Owner),
	}

	// get subtransaction barrier object
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// enable subtransaction barrier
	tx := l.svcCtx.BoxUserInterface.GetTx()
	sourceTx := tx.Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(sourceTx, func(tx1 *sql.Tx) error {
		exists, err := l.svcCtx.BoxUserInterface.IsBoxExistsByTx(boxUser, tx)
		if err != nil {
			return err
		}
		if exists {
			return dtmcli.ErrFailure
		}
		err = l.svcCtx.BoxUserInterface.AddBoxUserByTx(boxUser, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	l.svcCtx.BoxUserCaCheInterface.SetBoxUser(&model.BoxUserCache{
		Bid:   in.Bid,
		Uid:   in.Uid,
		Exist: true,
		Role:  int(boxuser.UserRole_Owner),
	}, l.svcCtx.Config.CacheExpire.BoxUserExist)

	return &boxuser.AddOwnerRequest{}, nil
}
