package slog

import (
	"sync"
)

var (
	_globalMu sync.RWMutex
	_globalL  = NewNoopLogger()
)

// ReplaceGlobalLogger replaces global logger
func ReplaceGlobalLogger(logger *Logger) func() {
	_globalMu.Lock()
	prevLogger := _globalL
	_globalL = logger
	_globalMu.Unlock()
	return func() { ReplaceGlobalLogger(prevLogger) }
}

// L global logger
func L() *Logger {
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}
