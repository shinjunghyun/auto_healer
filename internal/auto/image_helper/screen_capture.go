package image_helper

import (
	"auto_healer/configs"
	"auto_healer/internal/auto/window_helper"
	"fmt"
	"image"
	"image/color"

	"github.com/kbinani/screenshot"
)

func CaptureBaramScreen() (image.Image, error) {
	windowName := configs.BARAM_WINDOW_TITLE
	hwnd, err := window_helper.FindWindow(windowName)
	if err != nil {
		return nil, err
	}

	if hwnd == 0 {
		return nil, fmt.Errorf("window with title '%s' not found", windowName)
	}

	return PreProcessingCaptureScreen(hwnd)
}

func CaptureScreen(hwnd uintptr) (image.Image, error) {
	bounds := window_helper.GetClientBounds(hwnd)
	if bounds == nil {
		return nil, fmt.Errorf("failed to get window bounds")
	}

	img, err := screenshot.CaptureRect(image.Rect(
		int(bounds.Left),
		int(bounds.Top),
		int(bounds.Right),
		int(bounds.Bottom),
	))
	if err != nil {
		return nil, fmt.Errorf("failed to capture screen: %v", err)
	}
	return img, nil
}

func PreProcessingCaptureScreen(hwnd uintptr) (image.Image, error) {
	// 캡처된 이미지를 가져옵니다.
	img, err := CaptureScreen(hwnd)
	if err != nil {
		return nil, err
	}

	// RGBA 이미지를 가져옵니다.
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	processedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// 전처리: RGB 평균값 기반 이진화
	threshold := uint8(4) // 임계값
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 원본 픽셀 색상 가져오기
			r, g, b, a := img.At(x, y).RGBA()

			// RGB 평균값을 바로 이진화
			avgColor := uint8((r + g + b) / 3 >> 8)
			binary := uint8(0)
			if avgColor > threshold {
				binary = 255
			}

			// 새로운 픽셀 설정
			processedImg.Set(x, y, color.RGBA{R: binary, G: binary, B: binary, A: uint8(a >> 8)})
		}
	}

	return processedImg, nil
}
