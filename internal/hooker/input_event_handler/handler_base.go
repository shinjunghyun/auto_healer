package input_event_handler

import (
	"auto_healer/internal/auto"
	"context"
	log "logger"
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
}
