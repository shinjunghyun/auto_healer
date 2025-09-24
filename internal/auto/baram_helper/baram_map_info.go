package baram_helper

import (
	"auto_healer/internal/auto/image_helper"
	"crypto/md5"
	"encoding/hex"
	"image"
	log "logger"
)

func GetMapImageHash() (string, error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return "", err
	}

	croppedImg, err := image_helper.CropImage(img, image.Rectangle{
		Min: image.Point{X: 272, Y: 0},
		Max: image.Point{X: 422, Y: 20},
	})
	if err != nil {
		return "", err
	}

	// Access raw pixel data from the cropped image
	bounds := croppedImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixelData := make([]byte, 0, width*height*4) // Assuming RGBA format

	// Calculate histogram for Otsu's method
	histogram := make([]int, 256)
	totalPixels := width * height
	grayscaleValues := make([]byte, totalPixels)

	// First pass: convert to grayscale and build histogram
	idx := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := croppedImg.At(x, y).RGBA()

			// Convert to grayscale using BT.709 standard weights
			// 0.2126*R + 0.7152*G + 0.0722*B
			gray := byte((0.2126*float64(r>>8) + 0.7152*float64(g>>8) + 0.0722*float64(b>>8)) / 1.0)
			grayscaleValues[idx] = gray
			histogram[gray]++
			idx++
		}
	}

	// Calculate Otsu threshold
	otsuThreshold := calculateOtsuThreshold(histogram, totalPixels)

	// Second pass: apply the threshold
	idx = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := grayscaleValues[idx]
			idx++

			// Apply threshold for binarization
			var binarized byte
			if gray > otsuThreshold {
				binarized = 255 // White
			} else {
				binarized = 0 // Black
			}

			// Append binarized pixel data (RGBA format)
			pixelData = append(pixelData, binarized, binarized, binarized, 255) // Alpha is fully opaque
		}
	}

	// Log the calculated Otsu threshold
	log.Info().Msgf("Otsu threshold calculated: %d", otsuThreshold)

	// Compute the MD5 hash of the binarized pixel data
	hash := md5.Sum(pixelData)
	hashStr := hex.EncodeToString(hash[:])

	return hashStr, nil
}

// calculateOtsuThreshold implements Otsu's method to find optimal threshold
func calculateOtsuThreshold(histogram []int, totalPixels int) byte {
	// Calculate sum and sum of squared values for histogram
	sum := 0
	for i := 0; i < 256; i++ {
		sum += i * histogram[i]
	}

	sumB := 0
	wB := 0
	wF := 0
	maxVariance := 0.0
	threshold := 0

	// For each possible threshold, calculate variance and find the maximum
	for t := 0; t < 256; t++ {
		wB += histogram[t] // Weight of background
		if wB == 0 {
			continue
		}

		wF = totalPixels - wB // Weight of foreground
		if wF == 0 {
			break
		}

		sumB += t * histogram[t]
		meanB := float64(sumB) / float64(wB)     // Mean of background
		meanF := float64(sum-sumB) / float64(wF) // Mean of foreground

		// Calculate between-class variance
		variance := float64(wB) * float64(wF) * (meanB - meanF) * (meanB - meanF)

		// Update threshold if new maximum found
		if variance > maxVariance {
			maxVariance = variance
			threshold = t
		}
	}

	return byte(threshold)
}
