package baram_helper

import (
	"auto_healer/configs"
	"auto_healer/internal/auto/window_helper"
	"fmt"
)

func GetBaramWindowStartPosition() (int, int, error) {
	windowName := configs.BARAM_WINDOW_TITLE
	hwnd, err := window_helper.FindWindow(windowName)
	if err != nil {
		return 0, 0, err
	}

	if hwnd == 0 {
		return 0, 0, fmt.Errorf("window with title '%s' not found", windowName)
	}

	rect := window_helper.GetClientBounds(hwnd)
	if rect == nil {
		return 0, 0, fmt.Errorf("failed to get window bounds")
	}
	return int(rect.Left), int(rect.Top), nil
}
