package window_helper

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32                  = windows.NewLazySystemDLL("user32.dll")
	enumWindows             = user32.NewProc("EnumWindows")
	getWindowTextW          = user32.NewProc("GetWindowTextW")
	getWindowTextLengthW    = user32.NewProc("GetWindowTextLengthW")
	isWindowVisible         = user32.NewProc("IsWindowVisible")
	procSetWindowPos        = user32.NewProc("SetWindowPos")
	procSetForegroundWindow = user32.NewProc("SetForegroundWindow")
	procFindWindow          = user32.NewProc("FindWindowW")
)

func FindWindow(windowTitle string) (uintptr, error) {
	// Convert the window title to a UTF-16 pointer
	titlePtr, err := syscall.UTF16PtrFromString(windowTitle)
	if err != nil {
		return 0, err
	}

	hwnd, _, _ := procFindWindow.Call(0, uintptr(unsafe.Pointer(titlePtr)))

	return hwnd, nil
}

func ResizeWindow(hwnd uintptr, width, height int32) bool {
	// Constants for SetWindowPos
	const (
		SWP_NOMOVE   = 0x0002
		SWP_NOZORDER = 0x0004
	)

	// Call SetWindowPos to resize the window
	ret, _, _ := procSetWindowPos.Call(
		hwnd,
		0,
		0,
		0,
		uintptr(width),
		uintptr(height),
		SWP_NOMOVE|SWP_NOZORDER,
	)

	return ret != 0
}

func ActivateWindow(hwnd uintptr) bool {
	// Call SetForegroundWindow to activate the window
	ret, _, _ := procSetForegroundWindow.Call(hwnd)

	return ret != 0
}
