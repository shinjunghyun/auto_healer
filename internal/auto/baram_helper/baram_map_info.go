package baram_helper

import (
	"auto_healer/internal/auto/image_helper"
	"crypto/md5"
	"encoding/hex"
	"image"
)

func GetMapImageHash() (string, error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return "", err
	}

	croppedImg, err := image_helper.CropImage(img, image.Rectangle{
		Min: image.Point{X: 10, Y: 10},  // TODO: verify these values
		Max: image.Point{X: 110, Y: 60}, // TODO: verify these values
	})
	if err != nil {
		return "", err
	}

	// Access raw pixel data from the cropped image
	bounds := croppedImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixelData := make([]byte, 0, width*height*4) // Assuming RGBA format

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := croppedImg.At(x, y).RGBA()
			pixelData = append(pixelData, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8))
		}
	}

	// Compute the MD5 hash of the raw pixel data
	hash := md5.Sum(pixelData)
	hashStr := hex.EncodeToString(hash[:])

	return hashStr, nil
}
