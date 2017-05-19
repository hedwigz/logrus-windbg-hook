package windbghook

import (
	"github.com/Sirupsen/logrus"
)

type WinDebuggerHook struct{}

func NewWinDebuggerHook() *WinDebuggerHook {
	return &WinDebuggerHook{}
}

func (w *WinDebuggerHook) Levels() []logrus.Level {
	return nil
}

func (w *WinDebuggerHook) Fire(e *logrus.Entry) error {
	return nil
}
