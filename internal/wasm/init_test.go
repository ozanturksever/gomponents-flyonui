//go:build !js || !wasm

package wasm

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	// Test default values
	if config.Timeout != 30*time.Second {
		t.Errorf("Expected timeout 30s, got %v", config.Timeout)
	}

	if config.DOMReadyTimeout != 10*time.Second {
		t.Errorf("Expected DOM ready timeout 10s, got %v", config.DOMReadyTimeout)
	}

	if config.LibraryTimeout != 15*time.Second {
		t.Errorf("Expected library timeout 15s, got %v", config.LibraryTimeout)
	}

	if !config.AutoInitComponents {
		t.Error("Expected AutoInitComponents to be true")
	}

	if config.Components != nil {
		t.Error("Expected Components to be nil")
	}

	if config.MaxRetries != 3 {
		t.Errorf("Expected MaxRetries 3, got %d", config.MaxRetries)
	}

	if config.RetryDelay != 1*time.Second {
		t.Errorf("Expected RetryDelay 1s, got %v", config.RetryDelay)
	}
}

func TestInitializeWASM_MockEnvironment(t *testing.T) {
	config := DefaultConfig()
	result := InitializeWASM(config)

	if result == nil {
		t.Fatal("InitializeWASM returned nil result")
	}

	// In mock environment, initialization should fail
	if result.Success {
		t.Error("Expected initialization to fail in mock environment")
	}

	if result.DOMReady {
		t.Error("Expected DOMReady to be false in mock environment")
	}

	if result.LibraryLoaded {
		t.Error("Expected LibraryLoaded to be false in mock environment")
	}

	if result.ComponentsInit {
		t.Error("Expected ComponentsInit to be false in mock environment")
	}

	if result.Error == nil {
		t.Error("Expected error in mock environment")
	}

	if len(result.InitializedComps) != 0 {
		t.Error("Expected no initialized components in mock environment")
	}
}

func TestInitializeWASM_NilConfig(t *testing.T) {
	result := InitializeWASM(nil)

	if result == nil {
		t.Fatal("InitializeWASM returned nil result with nil config")
	}

	// Should use default config and still fail in mock environment
	if result.Success {
		t.Error("Expected initialization to fail in mock environment with nil config")
	}

	if result.Error == nil {
		t.Error("Expected error in mock environment with nil config")
	}
}

func TestInitializeWithCallback(t *testing.T) {
	config := DefaultConfig()
	callbackCalled := false
	var callbackResult *InitResult

	InitializeWithCallback(config, func(result *InitResult) {
		callbackCalled = true
		callbackResult = result
	})

	// Give some time for the callback to be called
	time.Sleep(100 * time.Millisecond)

	if !callbackCalled {
		t.Error("Callback was not called")
	}

	if callbackResult == nil {
		t.Error("Callback result was nil")
	}

	if callbackResult.Success {
		t.Error("Expected callback result to indicate failure in mock environment")
	}
}

func TestQuickInit(t *testing.T) {
	err := QuickInit()

	if err == nil {
		t.Error("Expected QuickInit to return error in mock environment")
	}
}

func TestInitWithComponents(t *testing.T) {
	components := []string{"HSDropdown", "HSModal"}
	timeout := 5 * time.Second

	err := InitWithComponents(components, timeout)

	if err == nil {
		t.Error("Expected InitWithComponents to return error in mock environment")
	}
}

func TestInitWithComponents_EmptyComponents(t *testing.T) {
	var components []string
	timeout := 5 * time.Second

	err := InitWithComponents(components, timeout)

	if err == nil {
		t.Error("Expected InitWithComponents to return error in mock environment")
	}
}

func TestInitWithComponents_ZeroTimeout(t *testing.T) {
	components := []string{"HSDropdown"}
	timeout := time.Duration(0)

	err := InitWithComponents(components, timeout)

	if err == nil {
		t.Error("Expected InitWithComponents to return error in mock environment")
	}
}

func TestInitResult_Structure(t *testing.T) {
	result := &InitResult{
		Success:          true,
		DOMReady:         true,
		LibraryLoaded:    true,
		ComponentsInit:   true,
		Error:            nil,
		InitializedComps: []string{"HSDropdown", "HSModal"},
		Duration:         5 * time.Second,
	}

	if !result.Success {
		t.Error("Expected Success to be true")
	}

	if !result.DOMReady {
		t.Error("Expected DOMReady to be true")
	}

	if !result.LibraryLoaded {
		t.Error("Expected LibraryLoaded to be true")
	}

	if !result.ComponentsInit {
		t.Error("Expected ComponentsInit to be true")
	}

	if result.Error != nil {
		t.Error("Expected Error to be nil")
	}

	if len(result.InitializedComps) != 2 {
		t.Errorf("Expected 2 initialized components, got %d", len(result.InitializedComps))
	}

	if result.Duration != 5*time.Second {
		t.Errorf("Expected duration 5s, got %v", result.Duration)
	}
}

func TestInitConfig_CustomValues(t *testing.T) {
	config := &InitConfig{
		Timeout:            60 * time.Second,
		DOMReadyTimeout:    20 * time.Second,
		LibraryTimeout:     30 * time.Second,
		AutoInitComponents: false,
		Components:         []string{"HSDropdown", "HSModal"},
		MaxRetries:         5,
		RetryDelay:         2 * time.Second,
	}

	if config.Timeout != 60*time.Second {
		t.Errorf("Expected timeout 60s, got %v", config.Timeout)
	}

	if config.DOMReadyTimeout != 20*time.Second {
		t.Errorf("Expected DOM ready timeout 20s, got %v", config.DOMReadyTimeout)
	}

	if config.LibraryTimeout != 30*time.Second {
		t.Errorf("Expected library timeout 30s, got %v", config.LibraryTimeout)
	}

	if config.AutoInitComponents {
		t.Error("Expected AutoInitComponents to be false")
	}

	if len(config.Components) != 2 {
		t.Errorf("Expected 2 components, got %d", len(config.Components))
	}

	if config.MaxRetries != 5 {
		t.Errorf("Expected MaxRetries 5, got %d", config.MaxRetries)
	}

	if config.RetryDelay != 2*time.Second {
		t.Errorf("Expected RetryDelay 2s, got %v", config.RetryDelay)
	}
}