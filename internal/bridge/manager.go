//go:build js && wasm

package bridge

import (
	"sync"
)

// Global manager instance
var (
	globalManager BridgeManager
	managerOnce   sync.Once
)

// InitializeManager initializes the global bridge manager with real implementations
func InitializeManager() {
	managerOnce.Do(func() {
		globalManager = NewRealBridgeManager()
	})
}

// GetManager returns the global bridge manager instance
func GetManager() BridgeManager {
	if globalManager == nil {
		InitializeManager()
	}
	return globalManager
}

// SetManager sets a custom bridge manager (useful for testing)
func SetManager(manager BridgeManager) {
	globalManager = manager
}

// ResetManager resets the global manager (useful for testing)
func ResetManager() {
	globalManager = nil
	managerOnce = sync.Once{}
}