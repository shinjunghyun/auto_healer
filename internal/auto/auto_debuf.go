package auto

import (
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"time"

	"github.com/micmonay/keybd_event"
)

func AutoDebuff(ctx context.Context) {
	time.Sleep(500 * time.Millisecond)
	performGuiYum()
	for {
		time.Sleep(20 * time.Millisecond)
		select {
		case <-ctx.Done():
			isDebuffing = false
			log.Info().Msgf("auto debuff context is done")
			return

		default:
			if isSelfHealing {
				log.Debug().Msgf("currently self-healing, will skip debuff...")
				continue
			}
			isDebuffing = true
			performDebuff()
		}
	}
}

func performGuiYum() {
	mtx.Lock()
	defer mtx.Unlock()

	hotkeys := ServerConfigInstance.CastingHotkeys

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(10 * time.Millisecond)

	// 1
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.GuiYum])
}

func performDebuff() {
	mtx.Lock()
	defer mtx.Unlock()

	hotkeys := ServerConfigInstance.CastingHotkeys

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(10 * time.Millisecond)

	// 1
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.HonMa])
	time.Sleep(10 * time.Millisecond)

	// up
	simulator.SendKeyboardInput(keybd_event.VK_UP)
	time.Sleep(50 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(50 * time.Millisecond)
}
