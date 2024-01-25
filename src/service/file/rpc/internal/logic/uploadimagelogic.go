package logic

import (
	"context"
	"io"
	"strconv"

	"nichebox/common/snowflake"
	"nichebox/service/file/model"
	"nichebox/service/file/rpc/common"
	"nichebox/service/file/rpc/internal/svc"
	"nichebox/service/file/rpc/pb/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImageLogic) UploadImage(stream file.File_UploadImageServer) error {
	var originData []byte
	var originName string

	// get image data
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return stream.SendAndClose(&file.UploadImageResponse{
				Code: file.UploadStatusCode_Failed,
			})
		}
		originData = append(originData, request.Content...)
		originName = request.FileName
	}

	webpId := snowflake.GenID()
	webpName := strconv.FormatInt(webpId, 10) + originName + ".webp"

	originId := snowflake.GenID()
	originName = strconv.FormatInt(originId, 10) + originName

	// image to webp
	webpBytes, err := common.ImageToWebp(originData, l.svcCtx.Config.Image.Quality)
	if err != nil {
		return stream.SendAndClose(&file.UploadImageResponse{
			Code: file.UploadStatusCode_Unknow,
		})
	}

	webpObject := "image/" + webpName
	originObject := "image/" + originName

	webpUrl, err := common.OSSUploadFile(l.svcCtx.OssClient, l.svcCtx.Config.AliyunOSS.BucketName, webpObject, webpBytes)
	if err != nil {
		return stream.SendAndClose(&file.UploadImageResponse{
			Code: file.UploadStatusCode_Unknow,
		})
	}

	originUrl, err := common.OSSUploadFile(l.svcCtx.OssClient, l.svcCtx.Config.AliyunOSS.BucketName, originObject, originData)
	if err != nil {
		return stream.SendAndClose(&file.UploadImageResponse{
			Code: file.UploadStatusCode_Unknow,
		})
	}

	webpModel := &model.Image{
		FileId: webpId,
		Url:    webpUrl,
	}
	if err := l.svcCtx.ImageFileInterface.CreateImage(webpModel); err != nil {
		return stream.SendAndClose(&file.UploadImageResponse{
			Code: file.UploadStatusCode_Unknow,
		})
	}

	originModel := &model.Image{
		FileId: originId,
		Url:    originUrl,
	}
	if err := l.svcCtx.ImageFileInterface.CreateImage(originModel); err != nil {
		return stream.SendAndClose(&file.UploadImageResponse{
			Code: file.UploadStatusCode_Unknow,
		})
	}

	return stream.SendAndClose(&file.UploadImageResponse{
		WebpId:   webpId,
		OriginId: originId,
		Code:     file.UploadStatusCode_Ok,
	})
}
