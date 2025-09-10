package tcp_packet

type InputType uint8

const (
	KEY_LEFT = iota
	KEY_DOWN
	KEY_RIGHT
	KEY_UP
	KEY_SPACE
	NUMBER_0
	NUMBER_1
	NUMBER_2
	NUMBER_3
	NUMBER_4
	NUMBER_5
	NUMBER_6
	NUMBER_7
	NUMBER_8
	NUMBER_9
	KEY_ESC
	KEY_TAB
	KEY_ENTER
)

type PacketPressed struct {
	PacketBase
	InputData InputType
}
