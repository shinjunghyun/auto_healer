package auto

import (
	"auto_healer/internal/auto/baram_helper"
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"tcp_packet"
	"time"

	"github.com/micmonay/keybd_event"
)

const (
	ClientMinHpPercent = 12.5
	ClientMaxHpPercent = 25.0
	ClientMinMpPercent = 5.0

	ServerMinHpPercent = 100
)

var (
	isSelfHealing = false
)

func AutoHeal(ctx context.Context) {
	for {
		time.Sleep(50 * time.Millisecond)
		select {
		case <-ctx.Done():
			log.Info().Msgf("auto heal context is done")
			return

		default:
			if time.Since(ServerBaramInfoData.LastUpdatedAt) > 1*time.Second {
				log.Error().Msgf("received baram info is too old to use [%s] ago", time.Since(ServerBaramInfoData.LastUpdatedAt).String())
				continue
			} else {
				// 클라이언트의 체마 갱신
				var err error
				ClientBaramInfoData.HpPercent, ClientBaramInfoData.MpPercent, err = baram_helper.GetHpMpPercent()
				if err != nil {
					log.Error().Msgf("error at retrieving HpMpPercent, will skip auto heal: %s", err.Error())
					continue
				}

				performAutoHeal(ServerBaramInfoData.PacketBaramInfo, ClientBaramInfoData)
			}
		}
	}
}

func performAutoHeal(ServerCharacter, ClientCharacter tcp_packet.PacketBaramInfo) {
	var err error

	log.Debug().Msgf("auto-heal: server [%1.f%%/%1.f%%] client [%1.f%%/%1.f%%]", ServerCharacter.HpPercent, ServerCharacter.MpPercent, ClientCharacter.HpPercent, ClientCharacter.MpPercent)

	// 마나 충전 확인
	if ClientCharacter.MpPercent < ClientMinMpPercent {
		log.Debug().Msgf("charging mana... [%.1f%%, %.1f%%]", ClientCharacter.MpPercent, ClientMinMpPercent)
		ChargeMP()
		return
	}

	// 자기 체력 확인
	if isSelfHealing || ClientCharacter.HpPercent < ClientMinHpPercent {
		log.Debug().Msgf("self healing... [%.1f%%, %.1f%%, %.1f%%]", ClientMinHpPercent, ClientCharacter.HpPercent, ClientMaxHpPercent)
		isSelfHealing = true
		SelfHeal()

		// 클라이언트의 체마 갱신
		ClientBaramInfoData.HpPercent, ClientBaramInfoData.MpPercent, err = baram_helper.GetHpMpPercent()
		if err != nil {
			log.Error().Msgf("error at retrieving HpMpPercent, will skip auto heal: %s", err.Error())
			return
		}

		if ClientCharacter.HpPercent >= ClientMaxHpPercent {
			isSelfHealing = false
		}
		return
	}

	// 상대 체력 확인
	if ServerCharacter.HpPercent < ServerMinHpPercent {
		log.Debug().Msgf("party healing... [%.1f, %.1f]", ServerCharacter.HpPercent, ServerMinHpPercent)
		PartyHeal()
		return
	}
}

func SelfHeal() {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(50 * time.Millisecond)

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)
	time.Sleep(50 * time.Millisecond)

	// home
	simulator.SendKeyboardInput(keybd_event.VK_HOME)
	time.Sleep(50 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(50 * time.Millisecond)
}

func PartyHeal() {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(50 * time.Millisecond)

	for range 2 {
		// tab
		simulator.SendKeyboardInput(keybd_event.VK_TAB)
		time.Sleep(50 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)
	time.Sleep(50 * time.Millisecond)
}

func ChargeMP() {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(50 * time.Millisecond)

	// 4
	simulator.SendKeyboardInput(keybd_event.VK_4)
	time.Sleep(50 * time.Millisecond)
}
