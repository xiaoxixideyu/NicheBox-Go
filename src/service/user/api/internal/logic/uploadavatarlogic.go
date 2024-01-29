package logic

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"nichebox/service/user/api/internal/common"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadAvatarLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadAvatarLogic) UploadAvatar() (resp *types.UploadAvatarResponse, err error) {
	if err := common.CheckUid(l.ctx, l.svcCtx.UserRpc); err != nil {
		return nil, err
	}

	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	l.r.ParseMultipartForm(l.svcCtx.Config.File.MaxMemory)
	file, fileHeader, err := l.r.FormFile("avatar")
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "File 出错")
	}
	log.Println("file name:", fileHeader.Size)
	defer file.Close()

	stream, err := l.svcCtx.FileRpc.UploadImage(l.ctx)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "UploadAvatar 服务出错: 2")
	}

	originId, webpId, err := common.UploadImage(file, strconv.FormatInt(uid, 10)+"avatar", stream)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserRpc.SetUserAvatarFileId(l.ctx, &userclient.SetUserAvatarFileIdRequest{
		Uid:      uid,
		OriginId: originId,
		WebpId:   webpId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UploadAvatarResponse{}, nil
}
