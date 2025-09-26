package image_helper

import (
	"image"
	"image/color"
)

func IsColorMatch(img image.Image, x, y int, target color.Color, tolerance uint8) (bool, error) {
	// 이미지에서 해당 좌표의 색상 가져오기
	pixelColor := img.At(x, y)

	// RGBA 값으로 변환
	pr, pg, pb, pa := pixelColor.RGBA()
	tr, tg, tb, ta := target.RGBA()

	// 16비트 색상 값(0-65535)을 8비트(0-255)로 변환
	pr8 := uint8(pr >> 8)
	pg8 := uint8(pg >> 8)
	pb8 := uint8(pb >> 8)
	pa8 := uint8(pa >> 8)

	tr8 := uint8(tr >> 8)
	tg8 := uint8(tg >> 8)
	tb8 := uint8(tb >> 8)
	ta8 := uint8(ta >> 8)

	// 각 색상 채널의 차이 계산
	rDiff := absDiff(pr8, tr8)
	gDiff := absDiff(pg8, tg8)
	bDiff := absDiff(pb8, tb8)
	aDiff := absDiff(pa8, ta8)

	// 모든 채널이 허용치 이내인지 확인
	return rDiff <= tolerance && gDiff <= tolerance && bDiff <= tolerance && aDiff <= tolerance, nil
}

// absDiff는 두 uint8 값의 절대 차이를 반환합니다.
func absDiff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}
