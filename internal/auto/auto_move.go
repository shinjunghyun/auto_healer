package auto

import (
	"auto_healer/internal/auto/baram_helper"
	"context"
	log "logger"
	"time"

	"github.com/go-vgo/robotgo"
)

func AutoMove(ctx context.Context) {
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			log.Info().Msgf("auto move context is done")
			return

		default:
			performAutoMove()
		}
	}
}

func performAutoMove() {
	var err error

	log.Trace().Msgf("auto-move")

	targetX, targetY, err := baram_helper.FindTabBoxPosition()
	if err != nil {
		log.Error().Msgf("error at finding tab box position, will skip auto move: %s", err.Error())
		return
	}

	moveRightClick(int32(targetX), int32(targetY))
}

func moveRightClick(x, y int32) {
	mtx.Lock()
	defer mtx.Unlock()

	startX, startY, err := baram_helper.GetBaramWindowStartPosition()
	if err != nil {
		log.Error().Err(err).Msgf("error at getting baram window start position, will skip auto move")
		return
	}
	offsetX, offsetY := 170, 27

	robotgo.Move(startX+offsetX+int(x), startY+offsetY+int(y))
	robotgo.Click("right", false)
}
