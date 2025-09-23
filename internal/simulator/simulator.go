package simulator

import (
	log "logger"
	"time"

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
	log.Trace().Msgf("simulating keyCode [0x%02X]...", key)
	kb.SetKeys(key)
	return kb.Launching()
}

func HoldKeyForMilliseconds(key int, durationMilliSeconds int) (err error) {
	log.Trace().Msgf("simulating keyCode [0x%02X] hold for %d ms...", key, durationMilliSeconds)
	kb.SetKeys(key)
	if err = kb.Press(); err != nil {
		return err
	}
	time.Sleep(time.Duration(durationMilliSeconds) * time.Millisecond)
	if err = kb.Release(); err != nil {
		return err
	}
	return nil
}
