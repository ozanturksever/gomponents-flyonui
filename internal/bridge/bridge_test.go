//go:build js && wasm

package bridge

import (
	"strings"
	"testing"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// TestGoStringsToJSArray tests conversion of Go string slice to JavaScript array
func TestGoStringsToJSArray(t *testing.T) {
	t.Run("converts empty slice", func(t *testing.T) {
		goSlice := []string{}
		jsArray := GoStringsToJSArray(goSlice)
		
		if jsArray.Get("length").Int() != 0 {
			t.Errorf("Expected empty array, got length %d", jsArray.Get("length").Int())
		}
	})
	
	t.Run("converts single element slice", func(t *testing.T) {
		goSlice := []string{"dropdown"}
		jsArray := GoStringsToJSArray(goSlice)
		
		if jsArray.Get("length").Int() != 1 {
			t.Errorf("Expected array length 1, got %d", jsArray.Get("length").Int())
		}
		
		if jsArray.Index(0).String() != "dropdown" {
			t.Errorf("Expected 'dropdown', got '%s'", jsArray.Index(0).String())
		}
	})
	
	t.Run("converts multiple element slice", func(t *testing.T) {
		goSlice := []string{"dropdown", "modal", "tooltip"}
		jsArray := GoStringsToJSArray(goSlice)
		
		if jsArray.Get("length").Int() != 3 {
			t.Errorf("Expected array length 3, got %d", jsArray.Get("length").Int())
		}
		
		expected := []string{"dropdown", "modal", "tooltip"}
		for i, exp := range expected {
			if jsArray.Index(i).String() != exp {
				t.Errorf("Expected '%s' at index %d, got '%s'", exp, i, jsArray.Index(i).String())
			}
		}
	})
}

// TestGoMapToJSObject tests conversion of Go map to JavaScript object
func TestGoMapToJSObject(t *testing.T) {
	t.Run("converts empty map", func(t *testing.T) {
		goMap := map[string]interface{}{}
		jsObj := GoMapToJSObject(goMap)
		
		// Check that object exists
		if jsObj.IsUndefined() {
			t.Error("Expected valid JS object, got undefined")
		}
	})
	
	t.Run("converts map with string values", func(t *testing.T) {
		goMap := map[string]interface{}{
			"selector": ".dropdown",
			"trigger":  "click",
		}
		jsObj := GoMapToJSObject(goMap)
		
		if jsObj.Get("selector").String() != ".dropdown" {
			t.Errorf("Expected '.dropdown', got '%s'", jsObj.Get("selector").String())
		}
		
		if jsObj.Get("trigger").String() != "click" {
			t.Errorf("Expected 'click', got '%s'", jsObj.Get("trigger").String())
		}
	})
	
	t.Run("converts map with mixed value types", func(t *testing.T) {
		goMap := map[string]interface{}{
			"enabled": true,
			"delay":   500,
			"text":    "Hello",
		}
		jsObj := GoMapToJSObject(goMap)
		
		if !jsObj.Get("enabled").Bool() {
			t.Error("Expected true for 'enabled'")
		}
		
		if jsObj.Get("delay").Int() != 500 {
			t.Errorf("Expected 500 for 'delay', got %d", jsObj.Get("delay").Int())
		}
		
		if jsObj.Get("text").String() != "Hello" {
			t.Errorf("Expected 'Hello' for 'text', got '%s'", jsObj.Get("text").String())
		}
	})
}

// TestInitializeFlyonComponents tests FlyonUI component initialization
func TestInitializeFlyonComponents(t *testing.T) {
	t.Run("initializes with empty component list", func(t *testing.T) {
		components := []string{}
		err := InitializeFlyonComponents(components)
		
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for component initialization in test environment without FlyonUI JS")
		} else if !strings.Contains(err.Error(), "HSStaticMethods not available") {
			t.Errorf("Expected 'HSStaticMethods not available' error, got %v", err)
		}
	})
	
	t.Run("initializes with valid component list", func(t *testing.T) {
		components := []string{"dropdown", "modal", "tooltip"}
		err := InitializeFlyonComponents(components)
		
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for component initialization in test environment without FlyonUI JS")
		} else if !strings.Contains(err.Error(), "HSStaticMethods not available") {
			t.Errorf("Expected 'HSStaticMethods not available' error, got %v", err)
		}
	})
	
	t.Run("handles HSStaticMethods not available", func(t *testing.T) {
		// This test simulates when FlyonUI JS is not loaded
		components := []string{"dropdown"}
		err := InitializeFlyonComponents(components)
		
		// Should handle gracefully and log warning
		if err == nil {
			logutil.Log("Warning: HSStaticMethods not available, components may not be interactive")
		}
	})
}

// TestInitializeSpecificComponent tests individual component initialization
func TestInitializeSpecificComponent(t *testing.T) {
	t.Run("initializes dropdown component", func(t *testing.T) {
		options := map[string]interface{}{
			"trigger": "click",
			"delay":   100,
		}
		
		err := InitializeDropdown(".dropdown", options)
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for dropdown initialization in test environment without FlyonUI JS")
		} else if !strings.Contains(err.Error(), "HSDropdown not available") {
			t.Errorf("Expected 'HSDropdown not available' error, got %v", err)
		}
	})
	
	t.Run("initializes modal component", func(t *testing.T) {
		options := map[string]interface{}{
			"backdrop": true,
			"keyboard": true,
		}
		
		err := InitializeModal(".modal", options)
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for modal initialization in test environment without FlyonUI JS")
		} else if !strings.Contains(err.Error(), "HSModal not available") {
			t.Errorf("Expected 'HSModal not available' error, got %v", err)
		}
	})
	
	t.Run("initializes tooltip component", func(t *testing.T) {
		options := map[string]interface{}{
			"placement": "top",
			"delay":     200,
		}
		
		err := InitializeTooltip(".tooltip", options)
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for tooltip initialization in test environment without FlyonUI JS")
		} else if !strings.Contains(err.Error(), "HSTooltip not available") {
			t.Errorf("Expected 'HSTooltip not available' error, got %v", err)
		}
	})
}

// TestErrorHandling tests error handling in bridge functions
func TestErrorHandling(t *testing.T) {
	t.Run("handles nil slice gracefully", func(t *testing.T) {
		var nilSlice []string
		jsArray := GoStringsToJSArray(nilSlice)
		
		if jsArray.Get("length").Int() != 0 {
			t.Errorf("Expected empty array for nil slice, got length %d", jsArray.Get("length").Int())
		}
	})
	
	t.Run("handles nil map gracefully", func(t *testing.T) {
		var nilMap map[string]interface{}
		jsObj := GoMapToJSObject(nilMap)
		
		if jsObj.IsUndefined() {
			t.Error("Expected valid JS object for nil map, got undefined")
		}
	})
	
	t.Run("handles invalid component names", func(t *testing.T) {
		components := []string{"invalid-component", "another-invalid"}
		err := InitializeFlyonComponents(components)
		
		// Should not fail but may log warnings
		if err != nil {
			logutil.Logf("Warning: Some components may not be available: %v", err)
		}
	})
}

// TestComponentLifecycle tests component lifecycle management
func TestComponentLifecycle(t *testing.T) {
	t.Run("destroys component instance", func(t *testing.T) {
		// Test component destruction/cleanup
		err := DestroyComponent(".dropdown", "dropdown")
		// In test environment, we expect an error because no elements exist
		if err == nil {
			t.Error("Expected error for component destruction in test environment without DOM elements")
		}
	})
	
	t.Run("reinitializes destroyed component", func(t *testing.T) {
		// First destroy (will fail because no elements)
		_ = DestroyComponent(".dropdown", "dropdown")
		
		// Then reinitialize (will fail because no FlyonUI JS)
		options := map[string]interface{}{"trigger": "click"}
		err := InitializeDropdown(".dropdown", options)
		// In test environment without FlyonUI JS, we expect an error
		if err == nil {
			t.Error("Expected error for component reinitialization in test environment without FlyonUI JS")
		}
	})
}

// TestPerformance tests performance aspects of bridge functions
func TestPerformance(t *testing.T) {
	t.Run("handles large string arrays efficiently", func(t *testing.T) {
		// Create large slice
		largeSlice := make([]string, 1000)
		for i := range largeSlice {
			largeSlice[i] = "component-" + string(rune(i))
		}
		
		jsArray := GoStringsToJSArray(largeSlice)
		
		if jsArray.Get("length").Int() != 1000 {
			t.Errorf("Expected array length 1000, got %d", jsArray.Get("length").Int())
		}
	})
	
	t.Run("handles large maps efficiently", func(t *testing.T) {
		// Create large map
		largeMap := make(map[string]interface{})
		for i := 0; i < 100; i++ {
			largeMap["key"+string(rune(i))] = "value" + string(rune(i))
		}
		
		jsObj := GoMapToJSObject(largeMap)
		
		if jsObj.IsUndefined() {
			t.Error("Expected valid JS object for large map")
		}
	})
}