package hooker

import (
	log "logger"
	"math/rand"
	"os"
	"time"
	"unsafe"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"github.com/moutend/go-hook/pkg/win32"
)

func StartKeyboardHooker(callback func()) error {
	log.Info().Msgf("start keyboard hooking...")

	keyChan := make(chan types.KeyboardEvent, 1)
	if err := keyboard.Install(ignoreKeyboardHookHandler, keyChan); err != nil {
		return err
	}
	defer func() {
		if err := keyboard.Uninstall(); err != nil {
			log.Error().Msgf("error at uninstalling keyboard hook: %s", err.Error())
		} else {
			log.Info().Msgf("keyboard hooking has been finished")
		}
	}()

	ctrlState := false
	qState := false

	for k := range keyChan {
		switch {
		case k.VKCode == types.VK_LCONTROL && k.Message == types.WM_KEYUP:
			ctrlState = false

		case k.VKCode == types.VK_Q && k.Message == types.WM_KEYUP:
			qState = false

		case k.VKCode == types.VK_Q && k.Message == types.WM_KEYDOWN:
			qState = true
			if ctrlState && qState {
				os.Exit(0)
			}

		case k.VKCode == types.VK_F7 && k.Message == types.WM_KEYDOWN:
			callback()

		case k.VKCode == types.VK_F8 && k.Message == types.WM_KEYDOWN:
			randObj := rand.New(rand.NewSource(time.Now().UnixNano()))
			go func() {
				for {
					callback()
					sleepSecond := 180 + randObj.Intn(120)
					log.Info().Msgf("next casting will be start at [%s]", time.Now().Add(time.Duration(sleepSecond)*time.Second).Format(time.RFC3339))
					time.Sleep(time.Duration(sleepSecond) * time.Second)
				}
			}()

		case k.VKCode == types.VK_LCONTROL && k.Message == types.WM_KEYDOWN:
			ctrlState = true
			switch {
			case qState:
				os.Exit(0)
			}
		}
	}

	return nil
}

func ignoreKeyboardHookHandler(c chan<- types.KeyboardEvent) types.HOOKPROC {
	return func(code int32, wParam, lParam uintptr) uintptr {
		if lParam == 0 {
			return win32.CallNextHookEx(0, code, wParam, lParam)
		} else {
			c <- types.KeyboardEvent{
				Message:         types.Message(wParam),
				KBDLLHOOKSTRUCT: *(*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
			}

			switch (*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)).VKCode {
			case types.VK_BROWSER_SEARCH: // 이런식으로 쓰자
				fallthrough
			default:
				return win32.CallNextHookEx(0, code, wParam, lParam)
			}
		}
	}
}
