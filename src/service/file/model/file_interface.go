package model

type ImageFileInterface interface {
	CreateImage(image *Image) error
	GetImageByFileId(fileId int64) (*Image, error)
}
