//go:build js && wasm

package bridge

import (
	"errors"
	"fmt"
	"syscall/js"
	"honnef.co/go/js/dom/v2"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// GoStringsToJSArray converts a Go string slice to a JavaScript array
func GoStringsToJSArray(goSlice []string) js.Value {
	if goSlice == nil {
		goSlice = []string{}
	}
	
	// Create JavaScript array
	jsArray := js.Global().Get("Array").New(len(goSlice))
	
	// Populate array
	for i, str := range goSlice {
		jsArray.SetIndex(i, js.ValueOf(str))
	}
	
	return jsArray
}

// GoMapToJSObject converts a Go map to a JavaScript object
func GoMapToJSObject(goMap map[string]interface{}) js.Value {
	if goMap == nil {
		goMap = make(map[string]interface{})
	}
	
	// Create JavaScript object
	jsObj := js.Global().Get("Object").New()
	
	// Populate object
	for key, value := range goMap {
		jsObj.Set(key, js.ValueOf(value))
	}
	
	return jsObj
}

// InitializeFlyonComponents initializes FlyonUI components using HSStaticMethods.autoInit()
func InitializeFlyonComponents(components []string) error {
	logutil.Logf("Initializing FlyonUI components: %v", components)
	
	// Check if HSStaticMethods is available
	hsStaticMethods := js.Global().Get("HSStaticMethods")
	if hsStaticMethods.IsUndefined() {
		logutil.Log("Warning: HSStaticMethods not available, FlyonUI JS may not be loaded")
		return errors.New("HSStaticMethods not available")
	}
	
	// Check if autoInit method exists
	autoInit := hsStaticMethods.Get("autoInit")
	if autoInit.IsUndefined() {
		logutil.Log("Warning: HSStaticMethods.autoInit not available")
		return errors.New("HSStaticMethods.autoInit not available")
	}
	
	// Convert components to JS array
	jsComponents := GoStringsToJSArray(components)
	
	// Call HSStaticMethods.autoInit with components
	try := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic during component initialization: %v", r)
			}
		}()
		
		autoInit.Invoke(jsComponents)
		return nil
	}
	
	if err := try(); err != nil {
		logutil.Logf("Error initializing components: %v", err)
		return err
	}
	
	logutil.Log("FlyonUI components initialized successfully")
	return nil
}

// InitializeDropdown initializes a specific dropdown component
func InitializeDropdown(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSDropdown", selector, options)
}

// InitializeModal initializes a specific modal component
func InitializeModal(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSModal", selector, options)
}

// InitializeTooltip initializes a specific tooltip component
func InitializeTooltip(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSTooltip", selector, options)
}

// InitializeAccordion initializes a specific accordion component
func InitializeAccordion(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSAccordion", selector, options)
}

// InitializeTabs initializes a specific tabs component
func InitializeTabs(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSTabs", selector, options)
}

// InitializeCarousel initializes a specific carousel component
func InitializeCarousel(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSCarousel", selector, options)
}

// InitializeCollapse initializes a specific collapse component
func InitializeCollapse(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSCollapse", selector, options)
}

// InitializeOffcanvas initializes a specific offcanvas component
func InitializeOffcanvas(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSOffcanvas", selector, options)
}

// InitializeScrollspy initializes a specific scrollspy component
func InitializeScrollspy(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSScrollspy", selector, options)
}

// InitializeSelect initializes a specific select component
func InitializeSelect(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSSelect", selector, options)
}

// InitializeTreeView initializes a specific tree view component (v1.2.0+)
func InitializeTreeView(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSTreeView", selector, options)
}

// InitializeDataTable initializes a specific data table component (v1.2.0+)
func InitializeDataTable(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSDataTable", selector, options)
}

// InitializeAdvancedRangeSlider initializes a specific advanced range slider component (v1.2.0+)
func InitializeAdvancedRangeSlider(selector string, options map[string]interface{}) error {
	return initializeSpecificComponent("HSAdvancedRangeSlider", selector, options)
}

// initializeSpecificComponent is a helper function to initialize any FlyonUI component
func initializeSpecificComponent(componentName, selector string, options map[string]interface{}) error {
	logutil.Logf("Initializing %s component with selector: %s", componentName, selector)
	
	// Check if the component class is available
	componentClass := js.Global().Get(componentName)
	if componentClass.IsUndefined() {
		logutil.Logf("Warning: %s not available, component may not be interactive", componentName)
		return fmt.Errorf("%s not available", componentName)
	}
	
	// Get DOM elements using honnef.co/go/js/dom/v2
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	
	if len(elements) == 0 {
		logutil.Logf("Warning: No elements found for selector: %s", selector)
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	// Convert options to JS object
	jsOptions := GoMapToJSObject(options)
	
	// Initialize component for each element
	try := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic during %s initialization: %v", componentName, r)
			}
		}()
		
		for _, element := range elements {
			// Create new component instance
			componentClass.New(element.Underlying(), jsOptions)
		}
		
		return nil
	}
	
	if err := try(); err != nil {
		logutil.Logf("Error initializing %s: %v", componentName, err)
		return err
	}
	
	logutil.Logf("%s component initialized successfully for %d elements", componentName, len(elements))
	return nil
}

// DestroyComponent destroys a specific component instance
func DestroyComponent(selector, componentType string) error {
	logutil.Logf("Destroying %s component with selector: %s", componentType, selector)
	
	// Get DOM elements
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	
	if len(elements) == 0 {
		logutil.Logf("Warning: No elements found for selector: %s", selector)
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	// Try to destroy component instances
	try := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic during component destruction: %v", r)
			}
		}()
		
		for _, element := range elements {
			// Try to get component instance and destroy it
			componentInstance := element.Underlying().Get("hs" + componentType)
			if !componentInstance.IsUndefined() {
				destroy := componentInstance.Get("destroy")
				if !destroy.IsUndefined() {
					destroy.Invoke()
				}
			}
		}
		
		return nil
	}
	
	if err := try(); err != nil {
		logutil.Logf("Error destroying %s: %v", componentType, err)
		return err
	}
	
	logutil.Logf("%s component destroyed successfully", componentType)
	return nil
}

// GetComponentInstance retrieves a component instance from a DOM element
func GetComponentInstance(selector, componentType string) (js.Value, error) {
	logutil.Logf("Getting %s component instance for selector: %s", componentType, selector)
	
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	
	if len(elements) == 0 {
		return js.Undefined(), fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	// Get the first element
	element := elements[0]
	componentInstance := element.Underlying().Get("hs" + componentType)
	
	if componentInstance.IsUndefined() {
		return js.Undefined(), fmt.Errorf("no %s instance found on element", componentType)
	}
	
	return componentInstance, nil
}

// AddEventListener adds an event listener to a DOM element using honnef.co/go/js/dom/v2
func AddEventListener(selector, eventType string, handler func(dom.Event)) error {
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	
	if len(elements) == 0 {
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	for _, element := range elements {
		element.AddEventListener(eventType, false, handler)
	}
	
	logutil.Logf("Event listener added for %s on %d elements", eventType, len(elements))
	return nil
}

// RemoveEventListener removes an event listener from elements matching the selector
// Note: This is a simplified implementation - in practice, you'd need to store
// the js.Func reference to properly remove the exact listener
func RemoveEventListener(selector, eventType string, handler func(dom.Event)) error {
	logutil.Logf("Removing %s event listener for selector: %s", eventType, selector)
	
	// Get DOM elements
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	
	if len(elements) == 0 {
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	// For proper event listener removal, we would need to store the js.Func
	// reference when adding the listener. This is a simplified implementation.
	for _, element := range elements {
		// Use syscall/js for removeEventListener since dom/v2 doesn't support
		// removing specific function references
		jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			// This is a placeholder - proper implementation would require
			// storing the original js.Func reference
			return nil
		})
		defer jsFunc.Release()
		
		element.Underlying().Call("removeEventListener", eventType, jsFunc)
	}
	
	logutil.Logf("Event listener removed for %s on %d elements", eventType, len(elements))
	return nil
}

// WaitForDOMReady waits for the DOM to be ready before executing a callback
func WaitForDOMReady(callback func()) {
	doc := dom.GetWindow().Document()
	
	// Access readyState as a property through the underlying JS object
	readyState := doc.Underlying().Get("readyState").String()
	
	if readyState == "loading" {
		// DOM is still loading, wait for DOMContentLoaded
		doc.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			callback()
		})
	} else {
		// DOM is already ready
		callback()
	}
}

// InitializeAllComponents initializes all FlyonUI components found in the DOM
func InitializeAllComponents() error {
	logutil.Log("Initializing all FlyonUI components")
	
	// List of all supported FlyonUI components
	allComponents := []string{
		"dropdown", "modal", "tooltip", "accordion", "tabs",
		"carousel", "collapse", "offcanvas", "scrollspy", "select",
		"tree-view", "data-table", "advanced-range-slider",
	}
	
	return InitializeFlyonComponents(allComponents)
}