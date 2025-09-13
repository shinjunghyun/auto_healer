package auto

import (
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"time"

	"github.com/micmonay/keybd_event"
)

func AutoDebuf(ctx context.Context) {
	for {
		time.Sleep(20 * time.Millisecond)
		select {
		case <-ctx.Done():
			isDebufing = false
			log.Info().Msgf("auto debuf context is done")
			return

		default:
			if isSelfHealing {
				log.Debug().Msgf("currently self-healing, will skip debuf...")
				continue
			}
			isDebufing = true
			performDebuf()
		}
	}
}

func performDebuf() {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(10 * time.Millisecond)

	// 1
	simulator.SendKeyboardInput(keybd_event.VK_1)
	time.Sleep(10 * time.Millisecond)

	// up
	simulator.SendKeyboardInput(keybd_event.VK_UP)
	time.Sleep(50 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(50 * time.Millisecond)
}
