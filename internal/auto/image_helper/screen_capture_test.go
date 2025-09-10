package image_helper

import (
	"auto_healer/internal/auto/window_helper"
	"image/png"
	"os"
	"testing"

	"golang.org/x/sys/windows"
)

func TestCaptureScreen(t *testing.T) {
	user32 := windows.NewLazySystemDLL("user32.dll")
	isWindowVisible := user32.NewProc("IsWindowVisible")

	// Find the window with the title "PingInfoView"
	windowName := "PingInfoView"
	hwnd, err := window_helper.FindWindow(windowName)
	if err != nil {
		t.Errorf("Error finding window: %v\n", err)
		return
	}

	if hwnd == 0 {
		t.Errorf("Window with title '%s' not found.\n", windowName)
		return
	}

	// Check if the window is visible
	visible, _, _ := isWindowVisible.Call(hwnd)
	if visible == 0 {
		t.Errorf("Window with title '%s' is not visible.\n", windowName)
		return
	}

	// Capture the screen of the found window
	img, err := CaptureScreen(hwnd)
	if err != nil {
		t.Errorf("Failed to capture screen: %v\n", err)
		return
	}

	// Save the captured image to a file
	file, err := os.Create("./capture.png")
	if err != nil {
		t.Errorf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Errorf("Failed to save image: %v\n", err)
		return
	}

	t.Logf("Screen captured and saved to ./capture.png")
}
