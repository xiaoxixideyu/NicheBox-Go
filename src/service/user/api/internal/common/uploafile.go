package common

import (
	"io"
	"mime/multipart"
	"net/http"
	"nichebox/service/file/rpc/pb/file"

	"github.com/zeromicro/x/errors"
)

func UploadImage(origin multipart.File, fileName string, stream file.File_UploadImageClient) (int64, int64, error) {
	buf := make([]byte, 1024)
	for {
		num, err := origin.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}

		if err := stream.Send(&file.UploadImageRequest{FileName: fileName, Content: buf[:num]}); err != nil {
			return 0, 0, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return 0, 0, err
	}

	if res.Code != file.UploadStatusCode_Ok {
		return 0, 0, errors.New(http.StatusInternalServerError, "上传失败")
	}

	return res.OriginId, res.WebpId, nil
}
