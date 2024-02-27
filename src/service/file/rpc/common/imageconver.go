package common

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/chai2010/webp"
)

func ImageToWebp(data []byte, quality float32) ([]byte, error) {
	// decode image
	img, format, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		log.Printf("failed to decode image, bytes len: %v, format: %v,  err: %v\n", len(data), format, err)
		return nil, err
	}

	// convert to webp
	webpBytes, err := webp.EncodeRGBA(img, quality)

	if err != nil {
		log.Printf("image failed to convert to webp, err: %v\n", err)
		return nil, err
	}

	return webpBytes, nil
}
