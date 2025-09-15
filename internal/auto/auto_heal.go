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
	ClientMinHpPercent, ClientMaxHpPercent = 0.15, 0.2
	ClientMinMpPercent                     = 0.25

	ServerMinHpPercent = 1.0

	backHoCooldown     = 3 * time.Second
	backHoChumCooldown = 3 * time.Second
)

var (
	isSelfHealing = false
	isDebuffing   = false

	baekHoUsedAt     = time.Time{}
	baekHoChumUsedAt = time.Time{}
)

func AutoHeal(ctx context.Context) {
	for {
		time.Sleep(90 * time.Millisecond)
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
				if time.Since(ClientBaramInfoData.LastUpdatedAt) > 1*time.Second {
					var err error
					ClientBaramInfoData.HpPercent, ClientBaramInfoData.MpPercent, err = baram_helper.GetHpMpPercent()
					if err != nil {
						log.Error().Msgf("error at retrieving HpMpPercent, will skip auto heal: %s", err.Error())
						continue
					}
					ClientBaramInfoData.LastUpdatedAt = time.Now()
				}

				if !isSelfHealing && isDebuffing {
					log.Debug().Msgf("currently debuffing, will skip auto heal...")
					continue
				}
				performAutoHeal(ServerBaramInfoData.PacketBaramInfo, ClientBaramInfoData.PacketBaramInfo)
			}
		}
	}
}

func performAutoHeal(ServerCharacter, ClientCharacter tcp_packet.PacketBaramInfo) {
	log.Debug().Msgf("auto-heal: server [%.f%%/%1.f%%] client [%.f%%/%.f%%]", ServerCharacter.HpPercent*100, ServerCharacter.MpPercent*100, ClientCharacter.HpPercent*100, ClientCharacter.MpPercent*100)

	// 마나 충전 확인
	if ClientCharacter.MpPercent < ClientMinMpPercent {
		log.Debug().Msgf("charging mana... [%.3f%%, %.3f%%]", ClientCharacter.MpPercent*100, ClientMinMpPercent*100)
		ChargeMP()
		return
	}

	// 자기 체력 확인
	if isSelfHealing || ClientCharacter.HpPercent < ClientMinHpPercent {
		log.Debug().Msgf("self healing... [%.3f%%, %.3f%%, %.3f%%]", ClientMinHpPercent*100, ClientCharacter.HpPercent*100, ClientMaxHpPercent*100)
		isSelfHealing = true
		SelfHeal(ClientCharacter.HpPercent)

		if ClientCharacter.HpPercent >= ClientMaxHpPercent {
			isSelfHealing = false
		}
		return
	}

	// 상대 체력 확인
	if ServerCharacter.HpPercent < ServerMinHpPercent {
		log.Debug().Msgf("party healing... [%.3f%%, %.3f%%]", ServerCharacter.HpPercent*100, ServerMinHpPercent*100)
		PartyHeal(ServerCharacter.HpPercent)
		return
	}
}

func SelfHeal(hp float32) {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(10 * time.Millisecond)

	if hp == 0 {
		// 2
		simulator.SendKeyboardInput(keybd_event.VK_2)
		time.Sleep(10 * time.Millisecond)

		// home
		simulator.SendKeyboardInput(keybd_event.VK_HOME)
		time.Sleep(10 * time.Millisecond)

		// enter
		simulator.SendKeyboardInput(keybd_event.VK_ENTER)
		time.Sleep(100 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)
	time.Sleep(10 * time.Millisecond)

	// home
	simulator.SendKeyboardInput(keybd_event.VK_HOME)
	time.Sleep(10 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(10 * time.Millisecond)
}

func PartyHeal(hp float32) {
	mtx.Lock()
	defer mtx.Unlock()

	// tab box 확인
	if time.Since(lastTabBoxCheckAt) < TabBoxCheckInterval {
		log.Trace().Msgf("skipping tab box check, last checked at [%s] ago", time.Since(lastTabBoxCheckAt).String())
	} else {
		lastTabBoxCheckAt = time.Now()
		if _, _, err := baram_helper.FindTabBoxPosition(); err != nil {
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

			} else {
				log.Error().Msgf("error at finding tab box position, will skip party heal: %s", err.Error())
				return
			}
		}
	}

	if hp == 0 {
		// 2
		simulator.SendKeyboardInput(keybd_event.VK_2)
		time.Sleep(10 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)

	if hp < 0.5 && time.Since(baekHoUsedAt) > backHoCooldown { // 백호의희원 사용
		simulator.SendKeyboardInput(keybd_event.VK_5)
		baekHoUsedAt = time.Now()
	} else if hp < 0.5 && time.Since(baekHoChumUsedAt) > backHoChumCooldown { // 백호의희원'첨 사용
		simulator.SendKeyboardInput(keybd_event.VK_6)
		baekHoChumUsedAt = time.Now()
	}
}

func ChargeMP() {
	mtx.Lock()
	defer mtx.Unlock()

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(50 * time.Millisecond)

	for range 2 {
		// 4
		simulator.SendKeyboardInput(keybd_event.VK_4)
		time.Sleep(50 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)
	time.Sleep(50 * time.Millisecond)

	// home
	simulator.SendKeyboardInput(keybd_event.VK_HOME)
	time.Sleep(10 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(90 * time.Millisecond)

	// 3
	simulator.SendKeyboardInput(keybd_event.VK_3)
	time.Sleep(50 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(200 * time.Millisecond)
}
