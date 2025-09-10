package input_event_handler

import (
	"auto_healer/internal/simulator"
	"auto_healer/internal/tcp_client/tcp_handler"
	log "logger"
	"tcp_packet"
	"time"
)

var moving bool = true

func HandleInputEvent() {
	log.Debug().Msgf("detected ctrl+e input event")

	// 도사 esc
	log.Debug().Msgf("trying to send {esc} key to the server...")
	if err := SendInputToServer(tcp_packet.KEY_ESC); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
	}
	time.Sleep(200 * time.Millisecond)

	// 도사 탭고정
	for range 2 {
		log.Debug().Msgf("trying to send {tab} key to the server...")
		if err := SendInputToServer(tcp_packet.KEY_TAB); err != nil {
			log.Error().Msgf("error at sending input to the server: %s", err.Error())
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 도사 부활
	log.Debug().Msgf("trying to send {2} key to the server...")
	if err := SendInputToServer(tcp_packet.NUMBER_2); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 n초동안 자힐
	startTime := time.Now()
	const healDuration = 5 * time.Second
	for {
		log.Debug().Msgf("trying to send {3} key to the server...")
		if err := SendInputToServer(tcp_packet.NUMBER_3); err != nil {
			log.Error().Msgf("error at sending input to the server: %s", err.Error())
			return
		}
		time.Sleep(100 * time.Millisecond)

		if time.Since(startTime) >= healDuration {
			break
		}
	}

	// 도사 esc
	log.Debug().Msgf("trying to send {esc} key to the server...")
	if err := SendInputToServer(tcp_packet.KEY_ESC); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
	}
	time.Sleep(500 * time.Millisecond)

	// 격수 esc
	log.Debug().Msgf("trying to simulate {esc} key...")
	if err := simulator.SendKeyboardInput(tcp_packet.KEY_ESC); err != nil {
		log.Error().Msgf("error at simulating key: %s", err.Error())
		return
	}
	time.Sleep(250 * time.Millisecond)

	// 격수 스페이스바
	log.Debug().Msgf("trying to simulate {space} key...")
	if err := simulator.SendKeyboardInput(tcp_packet.KEY_SPACE); err != nil {
		log.Error().Msgf("error at simulating key: %s", err.Error())
		return
	}
	time.Sleep(250 * time.Millisecond)

	// 도사 부활
	log.Debug().Msgf("trying to send {2} key to the server...")
	if err := SendInputToServer(tcp_packet.NUMBER_2); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 엔터
	log.Debug().Msgf("trying to send {enter} key to the server...")
	if err := SendInputToServer(tcp_packet.KEY_ENTER); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 희원
	log.Debug().Msgf("trying to send {5} key to the server...")
	if err := SendInputToServer(tcp_packet.NUMBER_5); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 엔터
	log.Debug().Msgf("trying to send {enter} key to the server...")
	if err := SendInputToServer(tcp_packet.KEY_ENTER); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 희원첨
	log.Debug().Msgf("trying to send {6} key to the server...")
	if err := SendInputToServer(tcp_packet.NUMBER_6); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)

	// 도사 파력
	log.Debug().Msgf("trying to send {0} key to the server...")
	if err := SendInputToServer(tcp_packet.NUMBER_0); err != nil {
		log.Error().Msgf("error at sending input to the server: %s", err.Error())
		return
	}
	time.Sleep(1100 * time.Millisecond)

	// 격수 분혼
	log.Debug().Msgf("trying to simulate {0} key...")
	if err := simulator.SendKeyboardInput(tcp_packet.NUMBER_0); err != nil {
		log.Error().Msgf("error at simulating key: %s", err.Error())
		return
	}
	time.Sleep(300 * time.Millisecond)

	// 무빙
	if moving {
		log.Debug().Msgf("trying to send {right} key to the server...")
		if err := SendInputToServer(tcp_packet.KEY_RIGHT); err != nil {
			log.Error().Msgf("error at sending input to the server: %s", err.Error())
			return
		}
		log.Debug().Msgf("trying to simulate {left} key...")
		if err := simulator.SendKeyboardInput(tcp_packet.KEY_LEFT); err != nil {
			log.Error().Msgf("error at simulating key: %s", err.Error())
			return
		}
	} else {
		log.Debug().Msgf("trying to send {left} key to the server...")
		if err := SendInputToServer(tcp_packet.KEY_LEFT); err != nil {
			log.Error().Msgf("error at sending input to the server: %s", err.Error())
			return
		}
		log.Debug().Msgf("trying to simulate {right} key...")
		if err := simulator.SendKeyboardInput(tcp_packet.KEY_RIGHT); err != nil {
			log.Error().Msgf("error at simulating key: %s", err.Error())
			return
		}
	}
	moving = !moving

	// 도사 공력 증강
	for range 8 {
		log.Debug().Msgf("trying to send {4} key to the server...")
		if err := SendInputToServer(tcp_packet.NUMBER_4); err != nil {
			log.Error().Msgf("error at sending input to the server: %s", err.Error())
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func SendInputToServer(input tcp_packet.InputType) error {
	return tcp_handler.SendPacket(&tcp_packet.PacketPressed{
		PacketBase: tcp_packet.PacketBase{
			PacketType: tcp_packet.PacketTypePressed,
		},
		InputData: input,
	})
}
