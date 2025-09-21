package baram_helper

import (
	"auto_healer/configs"
	opencv_proto "auto_healer/external/proto/opencv-proto"
	"auto_healer/internal/auto/image_helper"
	"auto_healer/internal/grpc_client"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	log "logger"
	"time"

	"github.com/ahmetb/go-linq"
	"golang.org/x/image/bmp"
)

func FindTabBoxPosition() (x, y int, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to capture baram screen: %v", err)
	}

	croppedImage, err := image_helper.CropImage(img, image.Rectangle{
		Min: image.Point{X: int(22), Y: int(16)},
		Max: image.Point{X: int(673), Y: int(594)},
	})

	// convert img to []byte
	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, croppedImage); err != nil {
		return 0, 0, fmt.Errorf("failed to encode image to BMP: %v", err)
	}
	imgBytes := buf.Bytes()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opencvServer := grpc_client.NewOpenCVServerClient(configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT)
	defer func() {
		if err := opencvServer.Close(); err != nil {
			log.Error().Err(err).Msgf("error at closing opencv grpc client from [%s:%d]", configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT)
		}
	}()

	err = opencvServer.Connect()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to connect to opencv grpc server [%s:%d]: %v", configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT, err)
	}

	res, err := opencvServer.FindTabBox(ctx, &opencv_proto.FindTabBoxRequest{Image: imgBytes})
	if err != nil {
		return 0, 0, fmt.Errorf("failed to find tab box position from opencv grpc server: %v", err)
	}

	found := res.GetFound()
	if !found {
		return 0, 0, fmt.Errorf("tab box not found in the image")
	}

	x = int(res.GetBox().GetX())
	y = int(res.GetBox().GetY())
	width := int(res.GetBox().GetWidth())
	height := int(res.GetBox().GetHeight())

	log.Trace().Msgf("tab box found at position (%d, %d) (%d, %d)", x, y, width, height)

	return x + width/2, y + height/2, nil
}

func GetHpMpPercent() (hpPercent, mpPercent float32, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to capture baram screen: %v", err)
	}

	cropped, err := image_helper.CropImage(img, image.Rectangle{
		Min: image.Point{X: int(BARAM_HP_MP_BOX_RECT.X), Y: int(BARAM_HP_MP_BOX_RECT.Y)},
		Max: image.Point{X: int(BARAM_HP_MP_BOX_RECT.X + BARAM_HP_MP_BOX_RECT.Width), Y: int(BARAM_HP_MP_BOX_RECT.Y + BARAM_HP_MP_BOX_RECT.Height)},
	})

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, cropped); err != nil {
		return 0, 0, fmt.Errorf("failed to encode image to PNG: %v", err)
	}
	croppedImgBytes := buf.Bytes()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opencvServer := grpc_client.NewOpenCVServerClient(configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT)
	defer func() {
		if err := opencvServer.Close(); err != nil {
			log.Error().Err(err).Msgf("error at closing opencv grpc client from [%s:%d]", configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT)
		}
	}()
	if err := opencvServer.Connect(); err != nil {
		return 0, 0, fmt.Errorf("failed to connect to opencv grpc server [%s:%d]: %v", configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT, err)
	} else if res, err := opencvServer.GetHpMpPercent(ctx, &opencv_proto.GetHpMpPercentRequest{CroppedImage: croppedImgBytes}); err != nil {
		return 0, 0, fmt.Errorf("failed to get hp/mp percent from opencv grpc server: %v", err)
	} else {
		hpPercent := res.GetHpPercent()
		mpPercent := res.GetMpPercent()

		log.Trace().Msgf("hp percent: %f, mp percent: %f", hpPercent, mpPercent)

		return hpPercent, mpPercent, nil
	}
}

func GetCoordinates() (x, y int, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to capture baram screen: %v", err)
	}

	// x 좌표 1의자리~100의자리
	for i := range 3 {
		pixelPointList := BARAM_COORDINATES_PIXELS[i]
		foundDigit := -1

		for digit := range 10 { // 0~9까지 숫자 확인
			log.Trace().Msgf("checking for digit %d at position %d", digit, i)
			if linq.From(pixelPointList[digit]).AllT(func(p PixelPoint) bool {
				log.Trace().Msgf("checking pixel at (%d, %d), expected color: %+v, actual color: %+v", p.X, p.Y, BARAM_COORDINATES_COLOR, img.At(p.X, p.Y))

				return img.At(p.X, p.Y) == BARAM_COORDINATES_COLOR
			}) {
				foundDigit = digit
				break
			}
		}

		// 자릿수에 맞게 x 값에 더하기 (i=0: 1의자리, i=1: 10의자리, i=2: 100의자리)
		multiplier := 1
		for range i {
			multiplier *= 10
		}
		x += foundDigit * multiplier
	}

	// y 좌표 1의자리~100의자리
	for i := range 3 {
		pixelPointList := BARAM_COORDINATES_PIXELS[i]
		foundDigit := -1

		for digit := range 10 { // 0~9까지 숫자 확인
			log.Trace().Msgf("checking for digit %d at position %d", digit, i)
			if linq.From(pixelPointList[digit]).AllT(func(p PixelPoint) bool {
				log.Trace().Msgf("checking pixel at (%d, %d), expected color: %+v, actual color: %+v", p.X, p.Y, BARAM_COORDINATES_COLOR, img.At(p.X, p.Y))

				return img.At(p.X+BARAM_COORDINATES_DISTANCE, p.Y) == BARAM_COORDINATES_COLOR
			}) {
				foundDigit = digit
				break
			}
		}

		// 자릿수에 맞게 y 값에 더하기 (i=0: 1의자리, i=1: 10의자리, i=2: 100의자리)
		multiplier := 1
		for range i {
			multiplier *= 10
		}
		y += foundDigit * multiplier
	}

	if x < 0 || y < 0 {
		return 0, 0, fmt.Errorf("failed to recognize coordinates, x: %d, y: %d", x, y)
	}

	return x, y, nil
}
