package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/common/cryptx"
	"nichebox/common/snowflake"
	"nichebox/service/user/model"
	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: check if code is valid

	uid := snowflake.GenID()

	// todo: randomUsername, default avatar
	userModel := model.User{
		Uid:      uid,
		Email:    in.Email,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Username: "default user",
	}

	err := l.svcCtx.UserInterface.CreateUser(&userModel)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return nil, status.Error(codes.AlreadyExists, "邮箱已存在")
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.RegisterResponse{Uid: userModel.Uid}, nil
}
