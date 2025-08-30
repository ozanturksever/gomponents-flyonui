//go:build !js || !wasm

package wasm

import (
	"errors"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// InitConfig holds configuration for WASM initialization
type InitConfig struct {
	// Timeout for overall initialization
	Timeout time.Duration
	// Timeout for DOM ready check
	DOMReadyTimeout time.Duration
	// Timeout for FlyonUI library loading
	LibraryTimeout time.Duration
	// Whether to initialize all components automatically
	AutoInitComponents bool
	// Specific components to initialize (if AutoInitComponents is false)
	Components []string
	// Retry configuration
	MaxRetries int
	RetryDelay time.Duration
}

// DefaultConfig returns a default initialization configuration
func DefaultConfig() *InitConfig {
	return &InitConfig{
		Timeout:            30 * time.Second,
		DOMReadyTimeout:    10 * time.Second,
		LibraryTimeout:     15 * time.Second,
		AutoInitComponents: true,
		Components:         nil,
		MaxRetries:         3,
		RetryDelay:         1 * time.Second,
	}
}

// InitResult holds the result of WASM initialization
type InitResult struct {
	Success           bool
	DOMReady          bool
	LibraryLoaded     bool
	ComponentsInit    bool
	Error             error
	InitializedComps  []string
	Duration          time.Duration
}

// InitializeWASM performs comprehensive WASM initialization with timeouts and error handling
// This is a mock implementation for non-WASM environments
func InitializeWASM(config *InitConfig) *InitResult {
	logutil.Log("Mock WASM initialization (non-WASM environment)")
	
	return &InitResult{
		Success:          false,
		DOMReady:         false,
		LibraryLoaded:    false,
		ComponentsInit:   false,
		Error:            errors.New("WASM initialization not available in non-WASM environment"),
		InitializedComps: []string{},
		Duration:         0,
	}
}

// InitializeWithCallback initializes WASM and calls a callback with the result
// This is a mock implementation for non-WASM environments
func InitializeWithCallback(config *InitConfig, callback func(*InitResult)) {
	result := InitializeWASM(config)
	callback(result)
}

// QuickInit performs a quick initialization with default settings
// This is a mock implementation for non-WASM environments
func QuickInit() error {
	return errors.New("WASM initialization not available in non-WASM environment")
}

// InitWithComponents initializes WASM with specific components
// This is a mock implementation for non-WASM environments
func InitWithComponents(components []string, timeout time.Duration) error {
	return errors.New("WASM initialization not available in non-WASM environment")
}