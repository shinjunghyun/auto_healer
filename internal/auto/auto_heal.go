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
	if ClientCharacter.MpPercent < ServerConfigInstance.HpMpControl.ClientMinMpPercent {
		log.Debug().Msgf("charging mana... [%.3f%%, %.3f%%]", ClientCharacter.MpPercent*100, ServerConfigInstance.HpMpControl.ClientMinMpPercent*100)
		chargeMP()
		return
	}

	// 자기 체력 확인
	if isSelfHealing || ClientCharacter.HpPercent < ServerConfigInstance.HpMpControl.ClientMinHpPercent {
		log.Debug().Msgf("self healing... [%.3f%%, %.3f%%, %.3f%%]", ServerConfigInstance.HpMpControl.ClientMinHpPercent*100, ClientCharacter.HpPercent*100, ServerConfigInstance.HpMpControl.ClientMaxHpPercent*100)
		isSelfHealing = true
		selfHeal(ClientCharacter.HpPercent)

		if ClientCharacter.HpPercent >= ServerConfigInstance.HpMpControl.ClientMaxHpPercent {
			isSelfHealing = false
		}
		return
	}

	// 상대 체력 확인
	if ServerCharacter.HpPercent < ServerConfigInstance.HpMpControl.ServerMinHpPercent {
		log.Debug().Msgf("party healing... [%.3f%%, %.3f%%]", ServerCharacter.HpPercent*100, ServerConfigInstance.HpMpControl.ServerMinHpPercent*100)
		partyHeal(ServerCharacter.HpPercent)
		return
	}
}

func selfHeal(hp float32) {
	mtx.Lock()
	defer mtx.Unlock()

	hotkeys := ServerConfigInstance.CastingHotkeys

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(10 * time.Millisecond)

	if hp == 0 {
		// 2
		simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.BooHwal])
		time.Sleep(10 * time.Millisecond)

		// home
		simulator.SendKeyboardInput(keybd_event.VK_HOME)
		time.Sleep(10 * time.Millisecond)

		// enter
		simulator.SendKeyboardInput(keybd_event.VK_ENTER)
		time.Sleep(100 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.KiWon])
	time.Sleep(10 * time.Millisecond)

	// home
	simulator.SendKeyboardInput(keybd_event.VK_HOME)
	time.Sleep(10 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(10 * time.Millisecond)
}

func partyHeal(hp float32) {
	mtx.Lock()
	defer mtx.Unlock()

	hotkeys := ServerConfigInstance.CastingHotkeys

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
		simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.BooHwal])
		time.Sleep(10 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.KiWon])

	if hp < 0.5 && time.Since(baekHoUsedAt) > time.Duration(ServerConfigInstance.CastingConfig.BaekHoCooldownMilliseconds)*time.Millisecond { // 백호의희원 사용
		simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.BaekHo])
		baekHoUsedAt = time.Now()
	} else if hp < 0.5 && time.Since(baekHoChumUsedAt) > time.Duration(ServerConfigInstance.CastingConfig.BaekHoChumCooldownMilliseconds)*time.Millisecond { // 백호의희원'첨 사용
		simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.BaekHoChum])
		baekHoChumUsedAt = time.Now()
	}
}

func chargeMP() {
	mtx.Lock()
	defer mtx.Unlock()

	hotkeys := ServerConfigInstance.CastingHotkeys

	// esc
	simulator.SendKeyboardInput(keybd_event.VK_ESC)
	time.Sleep(50 * time.Millisecond)

	for range 2 {
		// 4
		simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.GongRyuk])
		time.Sleep(50 * time.Millisecond)
	}

	// 3
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.KiWon])
	time.Sleep(50 * time.Millisecond)

	// home
	simulator.SendKeyboardInput(keybd_event.VK_HOME)
	time.Sleep(10 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(90 * time.Millisecond)

	// 3
	simulator.SendKeyboardInput(simulator.StringKeyToKeyCode[hotkeys.KiWon])
	time.Sleep(50 * time.Millisecond)

	// enter
	simulator.SendKeyboardInput(keybd_event.VK_ENTER)
	time.Sleep(100 * time.Millisecond)
}
