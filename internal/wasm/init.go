//go:build js && wasm

package wasm

import (
	"context"
	"errors"
	"fmt"
	"syscall/js"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
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
func InitializeWASM(config *InitConfig) *InitResult {
	if config == nil {
		config = DefaultConfig()
	}

	logutil.Log("Starting WASM initialization")
	startTime := time.Now()

	result := &InitResult{
		Success:          false,
		DOMReady:         false,
		LibraryLoaded:    false,
		ComponentsInit:   false,
		InitializedComps: []string{},
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	// Step 1: Wait for DOM ready
	logutil.Log("Waiting for DOM ready...")
	if err := waitForDOMReadyWithTimeout(ctx, config.DOMReadyTimeout); err != nil {
		result.Error = fmt.Errorf("DOM ready timeout: %w", err)
		result.Duration = time.Since(startTime)
		return result
	}
	result.DOMReady = true
	logutil.Log("DOM is ready")

	// Step 2: Wait for FlyonUI library
	logutil.Log("Waiting for FlyonUI library...")
	if err := waitForLibraryWithTimeout(ctx, config.LibraryTimeout, config.MaxRetries, config.RetryDelay); err != nil {
		result.Error = fmt.Errorf("FlyonUI library timeout: %w", err)
		result.Duration = time.Since(startTime)
		return result
	}
	result.LibraryLoaded = true
	logutil.Log("FlyonUI library loaded")

	// Step 3: Initialize bridge manager
	logutil.Log("Initializing bridge manager...")
	bridge.InitializeManager()

	// Step 4: Initialize components
	logutil.Log("Initializing components...")
	if config.AutoInitComponents {
		if err := bridge.InitializeAllComponents(); err != nil {
			result.Error = fmt.Errorf("component initialization failed: %w", err)
			result.Duration = time.Since(startTime)
			return result
		}
		result.InitializedComps = []string{"all"}
	} else if len(config.Components) > 0 {
		if err := bridge.InitializeFlyonComponents(config.Components); err != nil {
			result.Error = fmt.Errorf("specific component initialization failed: %w", err)
			result.Duration = time.Since(startTime)
			return result
		}
		result.InitializedComps = config.Components
	}
	result.ComponentsInit = true

	// Success!
	result.Success = true
	result.Duration = time.Since(startTime)
	logutil.Logf("WASM initialization completed successfully in %v", result.Duration)

	return result
}

// waitForDOMReadyWithTimeout waits for DOM to be ready with a timeout
func waitForDOMReadyWithTimeout(ctx context.Context, timeout time.Duration) error {
	domCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := make(chan struct{})
	var initErr error

	// Check DOM ready state
	bridge.WaitForDOMReady(func() {
		select {
		case <-domCtx.Done():
			// Context already cancelled
			return
		default:
			close(done)
		}
	})

	select {
	case <-done:
		return nil
	case <-domCtx.Done():
		if errors.Is(domCtx.Err(), context.DeadlineExceeded) {
			return errors.New("DOM ready timeout exceeded")
		}
		return domCtx.Err()
	}
}

// waitForLibraryWithTimeout waits for FlyonUI library to be available with retries
func waitForLibraryWithTimeout(ctx context.Context, timeout time.Duration, maxRetries int, retryDelay time.Duration) error {
	libCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for attempt := 0; attempt <= maxRetries; attempt++ {
		select {
		case <-libCtx.Done():
			if errors.Is(libCtx.Err(), context.DeadlineExceeded) {
				return errors.New("FlyonUI library timeout exceeded")
			}
			return libCtx.Err()
		default:
		}

		// Check if HSStaticMethods is available
		if isLibraryAvailable() {
			return nil
		}

		if attempt < maxRetries {
			logutil.Logf("FlyonUI library not ready, retrying in %v (attempt %d/%d)", retryDelay, attempt+1, maxRetries+1)
			time.Sleep(retryDelay)
		}
	}

	return errors.New("FlyonUI library not available after all retries")
}

// isLibraryAvailable checks if FlyonUI library is loaded and ready
func isLibraryAvailable() bool {
	// Check if HSStaticMethods is available
	hsStaticMethods := js.Global().Get("HSStaticMethods")
	if hsStaticMethods.IsUndefined() {
		return false
	}

	// Check if autoInit method exists
	autoInit := hsStaticMethods.Get("autoInit")
	if autoInit.IsUndefined() {
		return false
	}

	// Additional checks for common component classes
	commonComponents := []string{"HSDropdown", "HSModal", "HSTooltip"}
	for _, comp := range commonComponents {
		if js.Global().Get(comp).IsUndefined() {
			return false
		}
	}

	return true
}

// InitializeWithCallback initializes WASM and calls a callback with the result
func InitializeWithCallback(config *InitConfig, callback func(*InitResult)) {
	go func() {
		result := InitializeWASM(config)
		callback(result)
	}()
}

// QuickInit performs a quick initialization with default settings
func QuickInit() error {
	result := InitializeWASM(DefaultConfig())
	if !result.Success {
		return result.Error
	}
	return nil
}

// InitWithComponents initializes WASM with specific components
func InitWithComponents(components []string, timeout time.Duration) error {
	config := DefaultConfig()
	config.AutoInitComponents = false
	config.Components = components
	if timeout > 0 {
		config.Timeout = timeout
	}

	result := InitializeWASM(config)
	if !result.Success {
		return result.Error
	}
	return nil
}