//go:build js && wasm

package bridge

import (
	"syscall/js"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// Initialize the global manager on package load
func init() {
	InitializeManager()
}

// GoStringsToJSArray converts a Go string slice to a JavaScript array
func GoStringsToJSArray(goSlice []string) js.Value {
	return GetManager().GoStringsToJSArray(goSlice).Underlying().(js.Value)
}

// GoMapToJSObject converts a Go map to a JavaScript object
func GoMapToJSObject(goMap map[string]interface{}) js.Value {
	return GetManager().GoMapToJSObject(goMap).Underlying().(js.Value)
}

// InitializeFlyonComponents initializes FlyonUI components using HSStaticMethods.autoInit()
func InitializeFlyonComponents(components []string) error {
	return GetManager().InitializeFlyonComponents(components)
}

// InitializeDropdown initializes a specific dropdown component
func InitializeDropdown(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSDropdown", selector, options)
}

// InitializeModal initializes a specific modal component
func InitializeModal(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSModal", selector, options)
}

// InitializeTooltip initializes a specific tooltip component
func InitializeTooltip(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSTooltip", selector, options)
}

// InitializeAccordion initializes a specific accordion component
func InitializeAccordion(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSAccordion", selector, options)
}

// InitializeTabs initializes a specific tabs component
func InitializeTabs(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSTabs", selector, options)
}

// InitializeCarousel initializes a specific carousel component
func InitializeCarousel(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSCarousel", selector, options)
}

// InitializeCollapse initializes a specific collapse component
func InitializeCollapse(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSCollapse", selector, options)
}

// InitializeOffcanvas initializes a specific offcanvas component
func InitializeOffcanvas(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSOffcanvas", selector, options)
}

// InitializeScrollspy initializes a specific scrollspy component
func InitializeScrollspy(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSScrollspy", selector, options)
}

// InitializeSelect initializes a specific select component
func InitializeSelect(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSSelect", selector, options)
}

// InitializeTreeView initializes a specific tree view component (v1.2.0+)
func InitializeTreeView(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSTreeView", selector, options)
}

// InitializeDataTable initializes a specific data table component (v1.2.0+)
func InitializeDataTable(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSDataTable", selector, options)
}

// InitializeAdvancedRangeSlider initializes a specific advanced range slider component (v1.2.0+)
func InitializeAdvancedRangeSlider(selector string, options map[string]interface{}) error {
	return GetManager().InitializeSpecificComponent("HSAdvancedRangeSlider", selector, options)
}

// DestroyComponent destroys a specific component instance
func DestroyComponent(selector, componentType string) error {
	return GetManager().DestroyComponent(selector, componentType)
}

// GetComponentInstance retrieves a component instance from a DOM element
func GetComponentInstance(selector, componentType string) (js.Value, error) {
	instance, err := GetManager().GetComponentInstance(selector, componentType)
	if err != nil {
		return js.Undefined(), err
	}
	return instance.Underlying().(js.Value), nil
}

// AddEventListener adds an event listener to a DOM element
func AddEventListener(selector, eventType string, handler func(event DOMEvent)) error {
	return GetManager().AddEventListener(selector, eventType, handler)
}

// RemoveEventListener removes an event listener from elements matching the selector
func RemoveEventListener(selector, eventType string, handler func(event DOMEvent)) error {
	return GetManager().RemoveEventListener(selector, eventType, handler)
}

// WaitForDOMReady waits for the DOM to be ready before executing a callback
func WaitForDOMReady(callback func()) {
	GetManager().WaitForDOMReady(callback)
}

// InitializeAllComponents initializes all FlyonUI components found in the DOM
func InitializeAllComponents() error {
	logutil.Log("Initializing all FlyonUI components")
	return GetManager().InitializeAllComponents()
}