package baram_helper

import (
	"auto_healer/configs"
	opencv_proto "auto_healer/external/proto/opencv-proto"
	"auto_healer/internal/auto/image_helper"
	"auto_healer/internal/grpc_client"
	"bytes"
	"context"
	"fmt"
	"image/png"
	log "logger"
	"time"
)

func FindTabBoxPosition() (x, y int, err error) {
	img, err := image_helper.CaptureBaramScreen()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to capture baram screen: %v", err)
	}

	// convert img to []byte
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return 0, 0, fmt.Errorf("failed to encode image to PNG: %v", err)
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
	if err := opencvServer.Connect(); err != nil {
		return 0, 0, fmt.Errorf("failed to connect to opencv grpc server [%s:%d]: %v", configs.OPENCV_SERVER_HOST, configs.OPENCV_SERVER_PORT, err)
	} else if res, err := opencvServer.FindTabBox(ctx, &opencv_proto.FindTabBoxRequest{Image: imgBytes}); err != nil {
		return 0, 0, fmt.Errorf("failed to find tab box position from opencv grpc server: %v", err)
	} else if found := res.GetFound(); !found {
		return 0, 0, fmt.Errorf("tab box not found in the image")
	} else {
		x := int(res.GetBox().GetX())
		y := int(res.GetBox().GetY())

		log.Trace().Msgf("tab box found at position (%d, %d)", x, y)

		return x, y, nil
	}
}

func GetHpMpPercent() (hpPercent, mpPercent float32, err error) {
	return 0, 0, fmt.Errorf("not implemented yet")
}
