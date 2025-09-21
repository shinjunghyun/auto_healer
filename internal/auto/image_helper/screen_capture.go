package image_helper

import (
	"auto_healer/configs"
	"auto_healer/internal/auto/window_helper"
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/kbinani/screenshot"
)

const (
	screenCaptureCacheDuration = 500 * time.Millisecond
)

var (
	cachedImage  image.Image
	lastCachedAt time.Time
)

func CaptureBaramScreen() (image.Image, error) {
	if cachedImage != nil && time.Since(lastCachedAt) < screenCaptureCacheDuration {
		return cachedImage, nil
	}

	windowName := configs.BARAM_WINDOW_TITLE
	hwnd, err := window_helper.FindWindow(windowName)
	if err != nil {
		return nil, err
	}

	if hwnd == 0 {
		return nil, fmt.Errorf("window with title '%s' not found", windowName)
	}

	// fullImg, err := PreProcessingCaptureScreen(hwnd)
	fullImg, err := CaptureScreen(hwnd)
	if err != nil {
		return nil, err
	}

	// 바람 화면 영역 좌표
	x, y := 171, 28
	width, height := 1024, 768

	bounds := image.Rect(x, y, x+width, y+height)
	croppedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for cy := bounds.Min.Y; cy < bounds.Max.Y; cy++ {
		for cx := bounds.Min.X; cx < bounds.Max.X; cx++ {
			if cx < fullImg.Bounds().Max.X && cy < fullImg.Bounds().Max.Y {
				croppedImg.Set(cx-x, cy-y, fullImg.At(cx, cy))
			}
		}
	}

	cachedImage = croppedImg
	lastCachedAt = time.Now()

	return croppedImg, nil
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
			r, g, b, _ := img.At(x, y).RGBA()

			// RGB 평균값 계산 (16비트 값)
			avgColor := (r + g + b) / 3

			// color.RGBA는 8비트 값을 사용하므로, 완전한 흑백으로만 설정
			var binary uint8
			if avgColor > uint32(threshold)<<8 {
				binary = 255 // 완전한 흰색
			} else {
				binary = 0 // 완전한 검정
			}

			// alpha 값은 완전 불투명으로 설정
			processedImg.Set(x, y, color.RGBA{R: binary, G: binary, B: binary, A: 255})
		}
	}

	return processedImg, nil
}
