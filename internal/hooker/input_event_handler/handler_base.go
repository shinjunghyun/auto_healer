package input_event_handler

import (
	"auto_healer/internal/auto"
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"tcp_packet"
	"time"
)

func HandleInputEvent(ht HandlerType, ctx context.Context) {
	log.Debug().Msgf("start handling input event...")

	switch ht {
	case HandlerTypeMove:
		log.Info().Msgf("handling moving event...")
		auto.AutoMove(ctx)

	case HandlerTypeHeal:
		log.Info().Msgf("handling healing event...")
		auto.AutoHeal(ctx)

	case HandlerTypeDebuf:
		log.Info().Msgf("handling debuf event...")
		auto.AutoDebuf(ctx)

	default:
		log.Error().Msgf("unknown handler type: %s", ht.String())
	}

	// 격수 분혼
	key := tcp_packet.NUMBER_0
	log.Debug().Msgf("trying to simulate {%s} key...", key.String())
	if err := simulator.SendKeyboardInput(int(key)); err != nil {
		log.Error().Msgf("error at simulating key {%s}: %s", key.String(), err.Error())
		return
	}
	time.Sleep(300 * time.Millisecond)
}
