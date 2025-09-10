package simulator

import (
	log "logger"

	"github.com/micmonay/keybd_event"
)

var (
	kb keybd_event.KeyBonding
)

func init() {
	var err error
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		log.Error().Msgf("error at new key binding: %s", err.Error())
	}
}

func SendKeyboardInput(key int) (err error) {
	log.Debug().Msgf("simulating keyCode [0x%02X]...", key)
	kb.SetKeys(key)
	return kb.Launching()
}
