package auto

import (
	"auto_healer/internal/auto/baram_helper"
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/micmonay/keybd_event"
)

func AutoMove(ctx context.Context) {
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case <-ctx.Done():
			log.Info().Msgf("auto move context is done")
			return

		default:
			if isSelfHealing {
				log.Debug().Msgf("currently self-healing, will skip moving...")
				continue
			}
			if isDebufing {
				log.Debug().Msgf("currently debufing, will skip moving...")
				continue
			}
			performAutoMove()
		}
	}
}

func performAutoMove() {
	mtx.Lock()
	defer mtx.Unlock()
	log.Trace().Msgf("auto-move")

	targetX, targetY, err := baram_helper.FindTabBoxPosition()

	if err != nil {
		if err.Error() == "tab box not found in the image" {
			log.Info().Msgf("tab box not found in the image, will tab to find party member")

			// esc
			simulator.SendKeyboardInput(keybd_event.VK_ESC)
			time.Sleep(50 * time.Millisecond)

			for range 2 {
				// tab
				simulator.SendKeyboardInput(keybd_event.VK_TAB)
				time.Sleep(100 * time.Millisecond)
			}
			targetX, targetY, err = baram_helper.FindTabBoxPosition()
			if err != nil {
				log.Err(err).Msgf("error at finding tab box position after tabbing, will skip auto move")
				return
			}
		} else {
			log.Error().Msgf("error at finding tab box position, will skip auto move: %s", err.Error())
			return
		}
	}

	if targetX > 0 && targetY > 0 {
		moveRightClick(int32(targetX), int32(targetY))
	}
}

func moveRightClick(x, y int32) {
	startX, startY, err := baram_helper.GetBaramWindowStartPosition()
	if err != nil {
		log.Error().Err(err).Msgf("error at getting baram window start position, will skip auto move")
		return
	}
	offsetX, offsetY := 185, 27
	targetX, targetY := startX+offsetX+int(x), startY+offsetY+int(y)

	robotgo.Move(targetX, targetY)
	robotgo.Click("right")
}
