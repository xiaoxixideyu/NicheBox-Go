package logic

import (
	"context"
	"database/sql"

	"nichebox/service/box_info/model"
	"nichebox/service/box_info/rpc/internal/svc"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateBoxRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBoxRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBoxRevertLogic {
	return &CreateBoxRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBoxRevertLogic) CreateBoxRevert(in *boxinfo.CreateBoxRequest) (*boxinfo.CreateBoxResponse, error) {
	box := &model.Box{
		Bid: in.Bid,
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
		err = l.svcCtx.BoxInterface.RemoveBoxByTx(box, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &boxinfo.CreateBoxResponse{}, nil
}
