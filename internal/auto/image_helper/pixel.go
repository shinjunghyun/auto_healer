package image_helper

import (
	"fmt"
	"image"
)

func GetPixelColor(img image.Image, x, y int, tolerance float32) (color *uint32, err error) {
	// 이미지 경계 확인
	bounds := img.Bounds()
	if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
		return nil, fmt.Errorf("coordinates out of bounds")
	}

	// 특정 좌표의 색상 가져오기
	c := img.At(x, y)

	// 색상을 RGBA로 변환
	r, g, b, a := c.RGBA()

	// RGBA 값을 uint32로 병합
	pixelColor := (uint32(r>>8) << 24) | (uint32(g>>8) << 16) | (uint32(b>>8) << 8) | uint32(a>>8)

	return &pixelColor, nil
}
