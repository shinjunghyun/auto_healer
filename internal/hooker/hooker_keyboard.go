package hooker

import (
	"auto_healer/internal/auto"
	"auto_healer/internal/hooker/input_event_handler"
	"context"
	"fmt"
	log "logger"
	"os"
	"unsafe"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"github.com/moutend/go-hook/pkg/win32"
)

func StartKeyboardHooker(callback func(input_event_handler.HandlerType, context.Context)) error {
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

		// F6: Auto Move
		case k.VKCode == types.VK_F6 && k.Message == types.WM_KEYDOWN:
			if auto.AutoMoveCtx == nil {
				auto.AutoMoveCtx, auto.AutoMoveCancel = context.WithCancelCause(context.Background())
				go callback(input_event_handler.HandlerTypeMove, auto.AutoMoveCtx)
			} else {
				auto.AutoMoveCancel(fmt.Errorf("canceled by user"))
				auto.AutoMoveCtx = nil
			}

		// F7: Auto Heal
		case k.VKCode == types.VK_F7 && k.Message == types.WM_KEYDOWN:
			if auto.AutoHealCtx == nil {
				auto.AutoHealCtx, auto.AutoHealCancel = context.WithCancelCause(context.Background())
				go callback(input_event_handler.HandlerTypeHeal, auto.AutoHealCtx)
			} else {
				auto.AutoHealCancel(fmt.Errorf("canceled by user"))
				auto.AutoHealCtx = nil
			}

		// F8: Auto Debuf
		case k.VKCode == types.VK_F8 && k.Message == types.WM_KEYDOWN:
			if auto.AutoDebufCtx == nil {
				auto.AutoDebufCtx, auto.AutoDebufCancel = context.WithCancelCause(context.Background())
				go callback(input_event_handler.HandlerTypeDebuf, auto.AutoDebufCtx)
			} else {
				auto.AutoDebufCancel(fmt.Errorf("canceled by user"))
				auto.AutoDebufCtx = nil
			}

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
			case types.VK_F6,
				types.VK_F7,
				types.VK_F8:
				return 1

			default:
				return win32.CallNextHookEx(0, code, wParam, lParam)
			}
		}
	}
}
