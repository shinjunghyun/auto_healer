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
	ClientMinHP = 75_000
	ClientMaxHP = 150_000
	ClientMinMp = 30_000

	ServerMinHP = 1_400_000
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
				var err error

				// 클라이언트의 체마 갱신
				ClientBaramInfoData.HP, ClientBaramInfoData.MP, ClientBaramInfoData.Exp, err = baram_helper.GetHpMpExp()
				if err != nil {
					log.Error().Msgf("error at retrieving HpMpExp, will skip auto heal: %s", err.Error())
					continue
				}

				performAutoHeal(ServerBaramInfoData.PacketBaramInfo, ClientBaramInfoData)
			}
		}
	}
}

func performAutoHeal(ServerCharacter, ClientCharacter tcp_packet.PacketBaramInfo) {
	var err error

	log.Debug().Msgf("auto-heal: server [%d/%d] client [%d/%d]", ServerCharacter.HP, ServerCharacter.MP, ClientCharacter.HP, ClientCharacter.MP)

	// 마나 충전 확인
	if ClientCharacter.MP < ClientMinMp {
		log.Debug().Msgf("charging mana... [%d, %d]", ClientCharacter.MP, ClientMinMp)
		ChargeMP()
		return
	}

	// 자기 체력 확인
	if isSelfHealing || ClientCharacter.HP < ClientMinHP {
		log.Debug().Msgf("self healing... [%d, %d, %d]", ClientMinHP, ClientCharacter.HP, ClientMaxHP)
		isSelfHealing = true
		SelfHeal()

		// 클라이언트의 체마 갱신
		ClientBaramInfoData.HP, ClientBaramInfoData.MP, ClientBaramInfoData.Exp, err = baram_helper.GetHpMpExp()
		if err != nil {
			log.Error().Msgf("error at retrieving HpMpExp, will skip auto heal: %s", err.Error())
			return
		}

		if ClientCharacter.HP >= ClientMaxHP {
			isSelfHealing = false
		}
		return
	}

	// 상대 체력 확인
	if ServerCharacter.HP < ServerMinHP {
		log.Debug().Msgf("party healing... [%d, %d]", ServerCharacter.HP, ServerMinHP)
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
