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
		Min: image.Point{X: 272, Y: 0},
		Max: image.Point{X: 422, Y: 15},
	})
	if err != nil {
		return "", err
	}

	// Access raw pixel data from the cropped image
	bounds := croppedImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixelData := make([]byte, 0, width*height*4) // Assuming RGBA format

	// Threshold for binarization
	const threshold = 128

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := croppedImg.At(x, y).RGBA()

			// Convert to grayscale using the average method
			gray := (r>>8 + g>>8 + b>>8) / 3

			// Apply threshold for binarization
			var binarized byte
			if gray > threshold {
				binarized = 255 // White
			} else {
				binarized = 0 // Black
			}

			// Append binarized pixel data (RGBA format)
			pixelData = append(pixelData, binarized, binarized, binarized, 255) // Alpha is fully opaque
		}
	}

	// Compute the MD5 hash of the binarized pixel data
	hash := md5.Sum(pixelData)
	hashStr := hex.EncodeToString(hash[:])

	return hashStr, nil
}
