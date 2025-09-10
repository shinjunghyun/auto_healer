package simulator

import (
	"fmt"
	log "logger"
	"tcp_packet"

	"github.com/micmonay/keybd_event"
)

var (
	keyConvertingTable map[tcp_packet.InputType]int = map[tcp_packet.InputType]int{
		tcp_packet.KEY_LEFT:  keybd_event.VK_LEFT,
		tcp_packet.KEY_DOWN:  keybd_event.VK_DOWN,
		tcp_packet.KEY_RIGHT: keybd_event.VK_RIGHT,
		tcp_packet.KEY_UP:    keybd_event.VK_UP,
		tcp_packet.KEY_SPACE: keybd_event.VK_SPACE,
		tcp_packet.NUMBER_0:  keybd_event.VK_0,
		tcp_packet.NUMBER_1:  keybd_event.VK_1,
		tcp_packet.NUMBER_2:  keybd_event.VK_2,
		tcp_packet.NUMBER_3:  keybd_event.VK_3,
		tcp_packet.NUMBER_4:  keybd_event.VK_4,
		tcp_packet.NUMBER_5:  keybd_event.VK_5,
		tcp_packet.NUMBER_6:  keybd_event.VK_6,
		tcp_packet.NUMBER_7:  keybd_event.VK_7,
		tcp_packet.NUMBER_8:  keybd_event.VK_8,
		tcp_packet.NUMBER_9:  keybd_event.VK_9,
		tcp_packet.KEY_ESC:   keybd_event.VK_ESC,
		tcp_packet.KEY_TAB:   keybd_event.VK_TAB,
		tcp_packet.KEY_ENTER: keybd_event.VK_ENTER,
	}

	kb keybd_event.KeyBonding
)

func init() {
	var err error
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		log.Error().Msgf("error at new key binding: %s", err.Error())
	}
}

func SendKeyboardInput(inputData tcp_packet.InputType) (err error) {
	vk, ok := keyConvertingTable[inputData]
	if !ok {
		return fmt.Errorf("there is no mapping data in converting table")
	}

	log.Debug().Msgf("simulating keyCode [0x%02X]...", vk)
	kb.SetKeys(vk)
	return kb.Launching()
}
