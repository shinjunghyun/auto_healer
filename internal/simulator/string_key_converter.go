package simulator

import "github.com/micmonay/keybd_event"

var (
	StringKeyToKeyCode = map[string]int{
		"1": keybd_event.VK_1,
		"2": keybd_event.VK_2,
		"3": keybd_event.VK_3,
		"4": keybd_event.VK_4,
		"5": keybd_event.VK_5,
		"6": keybd_event.VK_6,
		"7": keybd_event.VK_7,
		"8": keybd_event.VK_8,
		"9": keybd_event.VK_9,
		"0": keybd_event.VK_0,
	}
)
