package baram_helper

import (
	"auto_healer/internal/auto/image_helper"
	"fmt"
	"image"
)

func readNumber(img image.Image, baseX, baseY, offset, digits int) (uint64, error) {
	number := uint64(0)
	for d := range digits { // 3자리 숫자
		digitX := baseX + d*offset
		found := false

		// 숫자 0~9를 탐색
		for i := range 10 {
			allMatch := true
			for _, targetPixel := range CoordsNumberPixlesMap[int8(i)] {
				color, err := image_helper.GetPixelColor(img, digitX+int(targetPixel.X), baseY+int(targetPixel.Y))
				if err != nil {
					return 0, err
				}

				if color != targetPixel.color {
					allMatch = false
					break
				}
			}

			if allMatch {
				number = number*10 + uint64(i) // 자릿수를 반영하여 숫자를 구성
				found = true
				break
			}
		}

		if !found {
			return 0, fmt.Errorf("failed to recognize digit at (%d, %d)", digitX, baseY)
		}
	}
	return number, nil
}

func GetCoordinates() (coordX, coordY int32, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, err
	}

	xBaseX, xBaseY := 100, 100 // FIXME: 샘플, X좌표의 기준 좌표
	yBaseX, yBaseY := 200, 100 // FIXME: 샘플, Y좌표의 기준 좌표
	offset := 20               // FIXME: 샘플, 숫자 간격

	// X좌표 읽기
	coordX64, err := readNumber(img, xBaseX, xBaseY, offset, 3)
	if err != nil {
		return 0, 0, err
	}
	coordX = int32(coordX64)

	// Y좌표 읽기
	coordY64, err := readNumber(img, yBaseX, yBaseY, offset, 3)
	if err != nil {
		return 0, 0, err
	}
	coordY = int32(coordY64)

	return coordX, coordY, nil
}

func GetHpMpExp() (hp, mp uint32, exp uint64, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, 0, err
	}

	hpBaseX, hpBaseY := 100, 100   // FIXME: 샘플, HP의 기준 좌표
	mpBaseX, mpBaseY := 100, 200   // FIXME: 샘플, MP의 기준 좌표
	expBaseX, expBaseY := 100, 200 // FIXME: 샘플, EXP의 기준 좌표
	offset := 20                   // FIXME: 샘플, 숫자 간격

	// HP 읽기
	hp64, err := readNumber(img, hpBaseX, hpBaseY, offset, 7)
	if err != nil {
		return 0, 0, 0, err
	}
	hp = uint32(hp64)

	// MP 읽기
	mp64, err := readNumber(img, mpBaseX, mpBaseY, offset, 7)
	if err != nil {
		return 0, 0, 0, err
	}
	mp = uint32(mp64)

	// EXP 읽기
	exp, err = readNumber(img, expBaseX, expBaseY, offset, 10)
	if err != nil {
		return 0, 0, 0, err
	}

	return hp, mp, exp, nil
}
