package tcp_packet

type InputType uint8

const (
	KEY_LEFT InputType = iota
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

func (it InputType) String() string {
	switch it {
	case KEY_LEFT:
		return "KEY_LEFT"
	case KEY_DOWN:
		return "KEY_DOWN"
	case KEY_RIGHT:
		return "KEY_RIGHT"
	case KEY_UP:
		return "KEY_UP"
	case KEY_SPACE:
		return "KEY_SPACE"
	case NUMBER_0:
		return "NUMBER_0"
	case NUMBER_1:
		return "NUMBER_1"
	case NUMBER_2:
		return "NUMBER_2"
	case NUMBER_3:
		return "NUMBER_3"
	case NUMBER_4:
		return "NUMBER_4"
	case NUMBER_5:
		return "NUMBER_5"
	case NUMBER_6:
		return "NUMBER_6"
	case NUMBER_7:
		return "NUMBER_7"
	case NUMBER_8:
		return "NUMBER_8"
	case NUMBER_9:
		return "NUMBER_9"
	case KEY_ESC:
		return "KEY_ESC"
	case KEY_TAB:
		return "KEY_TAB"
	case KEY_ENTER:
		return "KEY_ENTER"
	default:
		return "UNKNOWN"
	}
}

type PacketPressed struct {
	PacketBase
	InputData InputType
}
