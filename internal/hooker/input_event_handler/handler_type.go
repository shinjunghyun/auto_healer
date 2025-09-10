package input_event_handler

type HandlerType int

const (
	HandlerTypeMove HandlerType = iota
	HandlerTypeHeal
	HandlerTypeDebuf
)

func (ht HandlerType) String() string {
	switch ht {
	case HandlerTypeMove:
		return "HandlerTypeMove"

	case HandlerTypeHeal:
		return "HandlerTypeHeal"

	case HandlerTypeDebuf:
		return "HandlerTypeDebuf"

	default:
		return "Unknown"
	}
}
