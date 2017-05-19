package windbghook

import (
	"syscall"
	"unsafe"

	"github.com/Sirupsen/logrus"
)

type WinDebuggerHook struct {
	kernel32 *syscall.DLL
}

func NewWinDebuggerHook() *WinDebuggerHook {
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		panic(err)
	}
	return &WinDebuggerHook{
		kernel32: dll,
	}
}

func (w *WinDebuggerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
	}
}

func (w *WinDebuggerHook) Fire(e *logrus.Entry) error {
	log, err := e.String()
	if err != nil {
		return err
	}
	p := syscall.StringToUTF16Ptr(log)
	w.kernel32.MustFindProc("OutputDebugStringW").Call(uintptr(unsafe.Pointer(p)))
	return nil
}
