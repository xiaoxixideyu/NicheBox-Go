package logic

import (
	"context"
	"net/http"
	"strconv"

	"nichebox/service/file/rpc/fileclient"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GetAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAvatarLogic {
	return &GetAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAvatarLogic) GetAvatar(req *types.GetAvatarRequest) (resp *types.GetAvatarResponse, err error) {
	uid, err := strconv.ParseInt(req.Uid, 10, 64)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "uid 不正确")
	}
	fileIdRes, err := l.svcCtx.UserRpc.GetUserAvatarFileId(l.ctx, &userclient.GetUserAvatarFileIdRequest{
		Uid: uid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "GetAvatar 服务出错: 1")
	}

	if fileIdRes.OriginId == 0 {
		return &types.GetAvatarResponse{
			OriginUrl: "null",
			WebpUrl:   "null",
		}, nil
	}

	originUrlRes, err := l.svcCtx.FileRpc.GetFileUrl(l.ctx, &fileclient.GetFileUrlRequest{
		FileId: fileIdRes.OriginId,
	})
	if err != nil {
		return nil, errors.New(http.StatusNotFound, "GetAvatar 服务出错: 2")
	}

	webpUrlRes, err := l.svcCtx.FileRpc.GetFileUrl(l.ctx, &fileclient.GetFileUrlRequest{
		FileId: fileIdRes.WebpId,
	})
	if err != nil {
		return nil, errors.New(http.StatusNotFound, "GetAvatar 服务出错: 3")
	}

	return &types.GetAvatarResponse{
		OriginUrl: originUrlRes.FileUrl,
		WebpUrl:   webpUrlRes.FileUrl,
	}, nil
}
