//go:build !js || !wasm

package bridge

import (
	"errors"
	"testing"
)

func TestNewGoStringsToJSArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int // expected length
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []string{"test"},
			expected: 1,
		},
		{
			name:     "multiple elements",
			input:    []string{"test1", "test2", "test3"},
			expected: 3,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock manager
			mockManager := NewMockBridgeManager()
			Initialize(mockManager)
			defer func() { defaultManager = nil }()
			
			result := NewGoStringsToJSArray(tt.input)
			
			if result == nil {
				t.Fatal("Expected non-nil result")
			}
			
			// Verify the array was created with correct length
			mockValue := result.(*MockJSValue)
			if !mockValue.isArray {
				t.Error("Expected result to be an array")
			}
			
			if len(mockValue.arrayData) != tt.expected {
				t.Errorf("Expected array length %d, got %d", tt.expected, len(mockValue.arrayData))
			}
			
			// Verify array contents
			for i, expectedStr := range tt.input {
				if mockValue.arrayData[i] != expectedStr {
					t.Errorf("Expected array[%d] = %s, got %v", i, expectedStr, mockValue.arrayData[i])
				}
			}
		})
	}
}

func TestNewGoMapToJSObject(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]interface{}
	}{
		{
			name:  "empty map",
			input: map[string]interface{}{},
		},
		{
			name:  "single entry",
			input: map[string]interface{}{"key1": "value1"},
		},
		{
			name: "multiple entries with mixed types",
			input: map[string]interface{}{
				"string": "value",
				"number": 42,
				"bool":   true,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock manager
			mockManager := NewMockBridgeManager()
			Initialize(mockManager)
			defer func() { defaultManager = nil }()
			
			result := NewGoMapToJSObject(tt.input)
			
			if result == nil {
				t.Fatal("Expected non-nil result")
			}
			
			// Verify the object was created
			mockValue := result.(*MockJSValue)
			if !mockValue.isObject {
				t.Error("Expected result to be an object")
			}
			
			// Verify object contents
			for key, expectedValue := range tt.input {
				actualValue := mockValue.objectData[key]
				if actualValue != expectedValue {
					t.Errorf("Expected object[%s] = %v, got %v", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestNewInitializeFlyonComponents(t *testing.T) {
	tests := []struct {
		name       string
		components []string
		expectErr  bool
	}{
		{
			name:       "empty components",
			components: []string{},
			expectErr:  false,
		},
		{
			name:       "single component",
			components: []string{"HSDropdown"},
			expectErr:  false,
		},
		{
			name:       "multiple components",
			components: []string{"HSDropdown", "HSModal", "HSTooltip"},
			expectErr:  false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock manager
			mockManager := NewMockBridgeManager()
			Initialize(mockManager)
			defer func() { defaultManager = nil }()
			
			err := NewInitializeFlyonComponents(tt.components)
			
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			
			// Verify the component bridge was called
			mockComponentBridge := mockManager.Component().(*MockComponentBridge)
			calls := mockComponentBridge.GetInitializeAllCalls()
			
			if len(calls) != 1 {
				t.Errorf("Expected 1 InitializeAll call, got %d", len(calls))
			}
			
			if len(calls) > 0 {
				actualComponents := calls[0]
				if len(actualComponents) != len(tt.components) {
					t.Errorf("Expected %d components, got %d", len(tt.components), len(actualComponents))
				}
				
				for i, expected := range tt.components {
					if i < len(actualComponents) && actualComponents[i] != expected {
						t.Errorf("Expected component[%d] = %s, got %s", i, expected, actualComponents[i])
					}
				}
			}
		})
	}
}

func TestNewInitializeSpecificComponent(t *testing.T) {
	tests := []struct {
		name          string
		componentName string
		selector      string
		options       map[string]interface{}
		setupError    error
		expectErr     bool
	}{
		{
			name:          "successful initialization",
			componentName: "HSDropdown",
			selector:      ".dropdown",
			options:       map[string]interface{}{"toggle": true},
			expectErr:     false,
		},
		{
			name:          "initialization with error",
			componentName: "HSModal",
			selector:      ".modal",
			options:       nil,
			setupError:    errors.New("component not available"),
			expectErr:     true,
		},
		{
			name:          "empty options",
			componentName: "HSTooltip",
			selector:      ".tooltip",
			options:       map[string]interface{}{},
			expectErr:     false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock manager
			mockManager := NewMockBridgeManager()
			mockComponentBridge := mockManager.Component().(*MockComponentBridge)
			
			if tt.setupError != nil {
				mockComponentBridge.SetInitError(tt.componentName, tt.setupError)
			}
			
			Initialize(mockManager)
			defer func() { defaultManager = nil }()
			
			err := NewInitializeSpecificComponent(tt.componentName, tt.selector, tt.options)
			
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			
			if !tt.expectErr {
				// Verify the component was initialized
				initializedComponents := mockComponentBridge.GetInitializedComponents()
				selectors, exists := initializedComponents[tt.componentName]
				
				if !exists {
					t.Errorf("Expected component %s to be initialized", tt.componentName)
				} else if len(selectors) != 1 || selectors[0] != tt.selector {
					t.Errorf("Expected selector %s to be initialized for component %s, got %v", tt.selector, tt.componentName, selectors)
				}
			}
		})
	}
}

func TestComponentSpecificInitializationFunctions(t *testing.T) {
	componentTests := []struct {
		name         string
		initFunc     func(string, map[string]interface{}) error
		expectedComp string
	}{
		{"NewInitializeDropdown", NewInitializeDropdown, "HSDropdown"},
		{"NewInitializeModal", NewInitializeModal, "HSModal"},
		{"NewInitializeTooltip", NewInitializeTooltip, "HSTooltip"},
		{"NewInitializeAccordion", NewInitializeAccordion, "HSAccordion"},
		{"NewInitializeTabs", NewInitializeTabs, "HSTabs"},
		{"NewInitializeCarousel", NewInitializeCarousel, "HSCarousel"},
		{"NewInitializeCollapse", NewInitializeCollapse, "HSCollapse"},
		{"NewInitializeOffcanvas", NewInitializeOffcanvas, "HSOffcanvas"},
		{"NewInitializeScrollspy", NewInitializeScrollspy, "HSScrollspy"},
		{"NewInitializeSelect", NewInitializeSelect, "HSSelect"},
		{"NewInitializeTreeView", NewInitializeTreeView, "HSTreeView"},
		{"NewInitializeDataTable", NewInitializeDataTable, "HSDataTable"},
		{"NewInitializeAdvancedRangeSlider", NewInitializeAdvancedRangeSlider, "HSAdvancedRangeSlider"},
	}
	
	for _, tt := range componentTests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock manager
			mockManager := NewMockBridgeManager()
			Initialize(mockManager)
			defer func() { defaultManager = nil }()
			
			selector := ".test-selector"
			options := map[string]interface{}{"test": true}
			
			err := tt.initFunc(selector, options)
			
			if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			
			// Verify the correct component was initialized
			mockComponentBridge := mockManager.Component().(*MockComponentBridge)
			initializedComponents := mockComponentBridge.GetInitializedComponents()
			selectors, exists := initializedComponents[tt.expectedComp]
			
			if !exists {
				t.Errorf("Expected component %s to be initialized", tt.expectedComp)
			} else if len(selectors) != 1 || selectors[0] != selector {
				t.Errorf("Expected selector %s to be initialized for component %s, got %v", selector, tt.expectedComp, selectors)
			}
		})
	}
}

func TestNewDestroyComponent(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	selector := ".test-component"
	componentType := "Dropdown"
	
	err := NewDestroyComponent(selector, componentType)
	
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	
	// Verify the component was destroyed
	mockComponentBridge := mockManager.Component().(*MockComponentBridge)
	destroyedComponents := mockComponentBridge.GetDestroyedComponents()
	selectors, exists := destroyedComponents[componentType]
	
	if !exists {
		t.Errorf("Expected component %s to be destroyed", componentType)
	} else if len(selectors) != 1 || selectors[0] != selector {
		t.Errorf("Expected selector %s to be destroyed for component %s, got %v", selector, componentType, selectors)
	}
}

func TestNewGetComponentInstance(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	mockComponentBridge := mockManager.Component().(*MockComponentBridge)
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	selector := ".test-component"
	componentType := "Dropdown"
	
	// First initialize a component to create an instance
	err := mockComponentBridge.InitializeComponent("HSDropdown", selector, nil)
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}
	
	instance, err := NewGetComponentInstance(selector, componentType)
	
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	
	if instance == nil {
		t.Error("Expected non-nil instance")
	}
}

func TestNewAddEventListener(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	selector := ".test-element"
	eventType := "click"
	handlerCalled := false
	
	handler := func(event DOMEvent) {
		handlerCalled = true
	}
	
	err := NewAddEventListener(selector, eventType, handler)
	
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	
	// Verify the event listener was added
	mockDOMBridge := mockManager.DOM().(*MockDOMBridge)
	listeners := mockDOMBridge.GetEventListeners(selector, eventType)
	
	if len(listeners) != 1 {
		t.Errorf("Expected 1 event listener, got %d", len(listeners))
	}
	
	// Test triggering the event
	mockDOMBridge.TriggerEvent(selector, eventType)
	
	if !handlerCalled {
		t.Error("Expected handler to be called")
	}
}

func TestNewRemoveEventListener(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	selector := ".test-element"
	eventType := "click"
	
	err := NewRemoveEventListener(selector, eventType)
	
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
}

func TestNewWaitForDOMReady(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	mockDOMBridge := mockManager.DOM().(*MockDOMBridge)
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	callbackCalled := false
	callback := func() {
		callbackCalled = true
	}
	
	// Test when DOM is already ready
	mockDOMBridge.SetReady(true)
	NewWaitForDOMReady(callback)
	
	if !callbackCalled {
		t.Error("Expected callback to be called immediately when DOM is ready")
	}
	
	// Test when DOM is not ready
	callbackCalled = false
	mockDOMBridge.SetReady(false)
	NewWaitForDOMReady(callback)
	
	if callbackCalled {
		t.Error("Expected callback not to be called when DOM is not ready")
	}
	
	// Simulate DOM becoming ready
	mockDOMBridge.SetReady(true)
	
	if !callbackCalled {
		t.Error("Expected callback to be called when DOM becomes ready")
	}
}

func TestNewInitializeAllComponents(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	err := NewInitializeAllComponents()
	
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	
	// Verify all components were initialized
	mockComponentBridge := mockManager.Component().(*MockComponentBridge)
	calls := mockComponentBridge.GetInitializeAllCalls()
	
	if len(calls) != 1 {
		t.Errorf("Expected 1 InitializeAll call, got %d", len(calls))
	}
	
	if len(calls) > 0 {
		expectedComponents := []string{
			"HSDropdown", "HSModal", "HSTooltip", "HSAccordion", "HSTabs",
			"HSCarousel", "HSCollapse", "HSOffcanvas", "HSScrollspy", "HSSelect",
			"HSTreeView", "HSDataTable", "HSAdvancedRangeSlider",
		}
		
		actualComponents := calls[0]
		if len(actualComponents) != len(expectedComponents) {
			t.Errorf("Expected %d components, got %d", len(expectedComponents), len(actualComponents))
		}
		
		for i, expected := range expectedComponents {
			if i < len(actualComponents) && actualComponents[i] != expected {
				t.Errorf("Expected component[%d] = %s, got %s", i, expected, actualComponents[i])
			}
		}
	}
}

func TestValidateSelector(t *testing.T) {
	tests := []struct {
		name      string
		selector  string
		expectErr bool
	}{
		{"valid selector", ".test", false},
		{"empty selector", "", true},
		{"valid ID selector", "#test", false},
		{"valid attribute selector", "[data-test]", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSelector(tt.selector)
			
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateComponentName(t *testing.T) {
	tests := []struct {
		name          string
		componentName string
		expectErr     bool
	}{
		{"valid component", "HSDropdown", false},
		{"empty component name", "", true},
		{"invalid component", "InvalidComponent", true},
		{"valid modal component", "HSModal", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateComponentName(tt.componentName)
			
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestQueryElements(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	mockDOMBridge := mockManager.DOM().(*MockDOMBridge)
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	// Add mock elements
	element1 := NewMockDOMElement("test1")
	element2 := NewMockDOMElement("test2")
	mockDOMBridge.AddMockElement(".test", element1)
	mockDOMBridge.AddMockElement(".test", element2)
	
	elements := QueryElements(".test")
	
	if len(elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elements))
	}
	
	// Test with invalid selector
	emptyElements := QueryElements("")
	if len(emptyElements) != 0 {
		t.Errorf("Expected 0 elements for empty selector, got %d", len(emptyElements))
	}
}

func TestGetElementByID(t *testing.T) {
	// Setup mock manager
	mockManager := NewMockBridgeManager()
	mockDOMBridge := mockManager.DOM().(*MockDOMBridge)
	Initialize(mockManager)
	defer func() { defaultManager = nil }()
	
	// Add mock element
	element := NewMockDOMElement("test-id")
	mockDOMBridge.AddMockElement("#test-id", element)
	
	result := GetElementByID("test-id")
	
	if result == nil {
		t.Error("Expected non-nil element")
	}
	
	// Test with empty ID
	emptyResult := GetElementByID("")
	if emptyResult != nil {
		t.Error("Expected nil element for empty ID")
	}
}

func TestBridgeError(t *testing.T) {
	cause := errors.New("underlying error")
	
	// Test with component
	err1 := NewBridgeError("initialize", ".selector", "HSDropdown", cause)
	expected1 := "bridge initialize failed for component HSDropdown with selector .selector: underlying error"
	if err1.Error() != expected1 {
		t.Errorf("Expected error message: %s, got: %s", expected1, err1.Error())
	}
	
	// Test without component
	err2 := NewBridgeError("query", ".selector", "", cause)
	expected2 := "bridge query failed for selector .selector: underlying error"
	if err2.Error() != expected2 {
		t.Errorf("Expected error message: %s, got: %s", expected2, err2.Error())
	}
}

func TestGetManagerPanic(t *testing.T) {
	// Ensure no manager is set
	defaultManager = nil
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when manager not initialized")
		}
	}()
	
	GetManager()
}