package logic

import (
	"context"
	"database/sql"

	"nichebox/service/box_info/model"
	"nichebox/service/box_info/rpc/internal/svc"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
)

type CreateBoxLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBoxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBoxLogic {
	return &CreateBoxLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBoxLogic) CreateBox(in *boxinfo.CreateBoxRequest) (*boxinfo.CreateBoxResponse, error) {
	box := &model.Box{
		Bid:          in.Bid,
		Name:         in.Name,
		Introduction: in.Introduction,
	}

	// get subtransaction barrier object
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// enable subtransaction barrier
	tx := l.svcCtx.BoxInterface.GetTx()
	sourceTx := tx.Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(sourceTx, func(tx1 *sql.Tx) error {
		exists, err := l.svcCtx.BoxInterface.IsBoxExistsByTx(box, tx)
		if err != nil {
			return err
		}
		if exists {
			return dtmcli.ErrFailure
		}
		err = l.svcCtx.BoxInterface.CreateBoxByTx(box, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Internal, dtmcli.ResultFailure)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &boxinfo.CreateBoxResponse{}, nil
}
