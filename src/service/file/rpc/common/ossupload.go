package common

import (
	"bytes"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func OSSUploadFile(ossClient *oss.Client, bucketName, objectName string, data []byte) (string, error) {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		log.Printf("failed to open bucket, bucket name: %v, err: %v\n", bucketName, err)
		return "", err
	}

	err = bucket.PutObject(objectName, bytes.NewReader(data))
	if err != nil {
		log.Printf("failed to put image, err: %v\n", err)
		return "", err
	}

	return getPublicReadUrl(bucketName, ossClient.Config.Endpoint, objectName), nil
}

func getPublicReadUrl(bucketName, endPoint, objectName string) string {
	url := "https://" + bucketName + "." + endPoint + "/" + objectName
	return url
}
