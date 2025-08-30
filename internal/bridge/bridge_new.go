//go:build !js || !wasm

package bridge

import (
	"errors"
	"fmt"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// Global bridge manager instance
var defaultManager BridgeManager

// Initialize sets the global bridge manager
func Initialize(manager BridgeManager) {
	defaultManager = manager
}

// GetManager returns the global bridge manager
func GetManager() BridgeManager {
	if defaultManager == nil {
		panic("bridge manager not initialized - call bridge.Initialize() first")
	}
	return defaultManager
}

// NewGoStringsToJSArray converts Go string slice to JS array using the bridge
func NewGoStringsToJSArray(strings []string) JSValue {
	if len(strings) == 0 {
		return GetManager().JS().CreateArray(0)
	}
	
	jsArray := GetManager().JS().CreateArray(len(strings))
	for i, str := range strings {
		jsArray.SetIndex(i, str)
	}
	return jsArray
}

// NewGoMapToJSObject converts Go map to JS object using the bridge
func NewGoMapToJSObject(goMap map[string]interface{}) JSValue {
	if len(goMap) == 0 {
		return GetManager().JS().CreateObject()
	}
	
	jsObject := GetManager().JS().CreateObject()
	for key, value := range goMap {
		jsObject.Set(key, value)
	}
	return jsObject
}

// NewInitializeFlyonComponents initializes FlyonUI components using the bridge
func NewInitializeFlyonComponents(components []string) error {
	logutil.Logf("Initializing FlyonUI components: %v", components)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during FlyonUI initialization: %v", r)
		}
	}()
	
	return GetManager().Component().InitializeAll(components)
}

// NewInitializeSpecificComponent initializes a specific component type using the bridge
func NewInitializeSpecificComponent(componentName, selector string, options map[string]interface{}) error {
	logutil.Logf("Initializing %s component with selector: %s", componentName, selector)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during %s initialization: %v", componentName, r)
		}
	}()
	
	return GetManager().Component().InitializeComponent(componentName, selector, options)
}

// Component-specific initialization functions
func NewInitializeDropdown(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSDropdown", selector, options)
}

func NewInitializeModal(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSModal", selector, options)
}

func NewInitializeTooltip(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSTooltip", selector, options)
}

func NewInitializeAccordion(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSAccordion", selector, options)
}

func NewInitializeTabs(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSTabs", selector, options)
}

func NewInitializeCarousel(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSCarousel", selector, options)
}

func NewInitializeCollapse(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSCollapse", selector, options)
}

func NewInitializeOffcanvas(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSOffcanvas", selector, options)
}

func NewInitializeScrollspy(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSScrollspy", selector, options)
}

func NewInitializeSelect(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSSelect", selector, options)
}

func NewInitializeTreeView(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSTreeView", selector, options)
}

func NewInitializeDataTable(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSDataTable", selector, options)
}

func NewInitializeAdvancedRangeSlider(selector string, options map[string]interface{}) error {
	return NewInitializeSpecificComponent("HSAdvancedRangeSlider", selector, options)
}

// NewDestroyComponent destroys a component instance using the bridge
func NewDestroyComponent(selector, componentType string) error {
	logutil.Logf("Destroying %s component with selector: %s", componentType, selector)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during %s destruction: %v", componentType, r)
		}
	}()
	
	return GetManager().Component().DestroyComponent(selector, componentType)
}

// NewGetComponentInstance retrieves a component instance using the bridge
func NewGetComponentInstance(selector, componentType string) (JSValue, error) {
	logutil.Logf("Getting %s component instance for selector: %s", componentType, selector)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during %s instance retrieval: %v", componentType, r)
		}
	}()
	
	return GetManager().Component().GetComponentInstance(selector, componentType)
}

// NewAddEventListener adds an event listener using the bridge
func NewAddEventListener(selector, eventType string, handler func(DOMEvent)) error {
	logutil.Logf("Adding %s event listener for selector: %s", eventType, selector)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during event listener addition: %v", r)
		}
	}()
	
	return GetManager().DOM().AddEventListener(selector, eventType, handler)
}

// NewRemoveEventListener removes an event listener using the bridge
func NewRemoveEventListener(selector, eventType string) error {
	logutil.Logf("Removing %s event listener for selector: %s", eventType, selector)
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during event listener removal: %v", r)
		}
	}()
	
	return GetManager().DOM().RemoveEventListener(selector, eventType)
}

// NewWaitForDOMReady waits for DOM to be ready using the bridge
func NewWaitForDOMReady(callback func()) {
	logutil.Log("Waiting for DOM to be ready")
	
	defer func() {
		if r := recover(); r != nil {
			logutil.Logf("Panic during DOM ready wait: %v", r)
		}
	}()
	
	GetManager().DOM().WaitForReady(callback)
}

// NewInitializeAllComponents initializes all supported FlyonUI components using the bridge
func NewInitializeAllComponents() error {
	components := []string{
		"HSDropdown",
		"HSModal",
		"HSTooltip",
		"HSAccordion",
		"HSTabs",
		"HSCarousel",
		"HSCollapse",
		"HSOffcanvas",
		"HSScrollspy",
		"HSSelect",
		"HSTreeView",
		"HSDataTable",
		"HSAdvancedRangeSlider",
	}
	
	return NewInitializeFlyonComponents(components)
}

// Validation functions
func ValidateSelector(selector string) error {
	if selector == "" {
		return errors.New("selector cannot be empty")
	}
	return nil
}

func ValidateComponentName(componentName string) error {
	if componentName == "" {
		return errors.New("component name cannot be empty")
	}
	
	validComponents := map[string]bool{
		"HSDropdown":             true,
		"HSModal":                true,
		"HSTooltip":              true,
		"HSAccordion":            true,
		"HSTabs":                 true,
		"HSCarousel":             true,
		"HSCollapse":             true,
		"HSOffcanvas":            true,
		"HSScrollspy":            true,
		"HSSelect":               true,
		"HSTreeView":             true,
		"HSDataTable":            true,
		"HSAdvancedRangeSlider":  true,
	}
	
	if !validComponents[componentName] {
		return fmt.Errorf("unsupported component: %s", componentName)
	}
	
	return nil
}

// Helper functions for common operations
func QueryElements(selector string) []DOMElement {
	if err := ValidateSelector(selector); err != nil {
		logutil.Logf("Invalid selector: %v", err)
		return []DOMElement{}
	}
	
	return GetManager().DOM().QuerySelectorAll(selector)
}

func GetElementByID(id string) DOMElement {
	if id == "" {
		logutil.Log("Element ID cannot be empty")
		return nil
	}
	
	return GetManager().DOM().GetElementByID(id)
}

// Error handling utilities
type BridgeError struct {
	Operation string
	Selector  string
	Component string
	Cause     error
}

func (e *BridgeError) Error() string {
	if e.Component != "" {
		return fmt.Sprintf("bridge %s failed for component %s with selector %s: %v", e.Operation, e.Component, e.Selector, e.Cause)
	}
	return fmt.Sprintf("bridge %s failed for selector %s: %v", e.Operation, e.Selector, e.Cause)
}

func NewBridgeError(operation, selector, component string, cause error) *BridgeError {
	return &BridgeError{
		Operation: operation,
		Selector:  selector,
		Component: component,
		Cause:     cause,
	}
}