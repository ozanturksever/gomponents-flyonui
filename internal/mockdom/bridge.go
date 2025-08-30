//go:build !js || !wasm

package mockdom

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
)

// MockBridge implements the bridge interfaces using mock DOM
type MockBridge struct {
	dom *MockDOM
	mu  sync.RWMutex
	
	// Component tracking
	components map[string]map[string]interface{} // componentType -> elementID -> instance
	
	// Event listener tracking
	eventListeners map[string][]func()
	
	// Global JS object simulation
	global map[string]interface{}
	
	// HSStaticMethods simulation
	hsStaticMethods map[string]interface{}
}

// MockJSValue implements bridge.JSValue using mock data
type MockJSValue struct {
	value interface{}
	props map[string]interface{}
}

// MockDOMElement implements bridge.DOMElement using mock DOM
type MockDOMElement struct {
	element *MockElement
	bridge  *MockBridge
}

// NewMockBridge creates a new mock bridge with a mock DOM
func NewMockBridge() *MockBridge {
	dom := NewMockDOM()
	
	bridge := &MockBridge{
		dom:        dom,
		components: make(map[string]map[string]interface{}),
		eventListeners: make(map[string][]func()),
		global:     make(map[string]interface{}),
		hsStaticMethods: make(map[string]interface{}),
	}
	
	// Set up default HSStaticMethods
	bridge.setupHSStaticMethods()
	
	return bridge
}

// setupHSStaticMethods sets up mock HSStaticMethods for FlyonUI components
func (b *MockBridge) setupHSStaticMethods() {
	autoInit := func(components []string, options map[string]interface{}) {
		// Mock autoInit - just track that it was called
		for _, component := range components {
			b.initializeComponent(component, options)
		}
	}
	
	b.hsStaticMethods["autoInit"] = autoInit
	
	// Add individual component initializers
	componentTypes := []string{
		"HSDropdown", "HSModal", "HSCollapse", "HSTabs", "HSAccordion",
		"HSCarousel", "HSTooltip", "HSPopover", "HSOffcanvas", "HSScrollspy",
		"HSSelect", "HSInputNumber", "HSFileUpload", "HSDatepicker",
		"HSTimepicker", "HSColorpicker", "HSRange", "HSTogglePassword",
		"HSPinInput", "HSCopyMarkup", "HSComboBox", "HSDataTable",
		"HSThemeSwitch", "HSToggleCount", "HSRemoveElement", "HSStepForm",
	}
	
	for _, componentType := range componentTypes {
		b.hsStaticMethods[componentType] = map[string]interface{}{
			"autoInit": func(selector string, options map[string]interface{}) {
				b.initializeSpecificComponent(componentType, selector, options)
			},
			"getInstance": func(element interface{}) interface{} {
				if mockEl, ok := element.(*MockDOMElement); ok {
					return mockEl.element.GetComponentInstance(componentType)
				}
				return nil
			},
		}
	}
	
	b.global["HSStaticMethods"] = b.hsStaticMethods
}

// initializeComponent initializes a component type globally
func (b *MockBridge) initializeComponent(componentType string, options map[string]interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.components[componentType] == nil {
		b.components[componentType] = make(map[string]interface{})
	}
	
	// Find all elements that could be this component type
	selector := b.getComponentSelector(componentType)
	elements := b.dom.QuerySelectorAll(selector)
	
	for _, element := range elements {
		instance := b.createComponentInstance(componentType, element, options)
		element.SetComponentInstance(componentType, instance)
		if element.ID != "" {
			b.components[componentType][element.ID] = instance
		}
	}
}

// initializeSpecificComponent initializes components matching a selector
func (b *MockBridge) initializeSpecificComponent(componentType, selector string, options map[string]interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.components[componentType] == nil {
		b.components[componentType] = make(map[string]interface{})
	}
	
	elements := b.dom.QuerySelectorAll(selector)
	for _, element := range elements {
		instance := b.createComponentInstance(componentType, element, options)
		element.SetComponentInstance(componentType, instance)
		if element.ID != "" {
			b.components[componentType][element.ID] = instance
		}
	}
}

// getComponentSelector returns the default selector for a component type
func (b *MockBridge) getComponentSelector(componentType string) string {
	selectors := map[string]string{
		"HSDropdown":       "[data-hs-dropdown]",
		"HSModal":          "[data-hs-overlay]",
		"HSCollapse":       "[data-hs-collapse]",
		"HSTabs":           "[data-hs-tab]",
		"HSAccordion":      "[data-hs-accordion]",
		"HSCarousel":       "[data-hs-carousel]",
		"HSTooltip":        "[data-hs-tooltip]",
		"HSPopover":        "[data-hs-popover]",
		"HSOffcanvas":      "[data-hs-offcanvas]",
		"HSScrollspy":      "[data-hs-scrollspy]",
		"HSSelect":         "[data-hs-select]",
		"HSInputNumber":    "[data-hs-input-number]",
		"HSFileUpload":     "[data-hs-file-upload]",
		"HSDatepicker":     "[data-hs-datepicker]",
		"HSTimepicker":     "[data-hs-timepicker]",
		"HSColorpicker":    "[data-hs-color-picker]",
		"HSRange":          "[data-hs-range]",
		"HSTogglePassword": "[data-hs-toggle-password]",
		"HSPinInput":       "[data-hs-pin-input]",
		"HSCopyMarkup":     "[data-hs-copy-markup]",
		"HSComboBox":       "[data-hs-combo-box]",
		"HSDataTable":      "[data-hs-datatable]",
		"HSThemeSwitch":    "[data-hs-theme-switch]",
		"HSToggleCount":    "[data-hs-toggle-count]",
		"HSRemoveElement":  "[data-hs-remove-element]",
		"HSStepForm":       "[data-hs-stepper]",
	}
	
	if selector, exists := selectors[componentType]; exists {
		return selector
	}
	return fmt.Sprintf("[data-hs-%s]", strings.ToLower(componentType[2:]))
}

// createComponentInstance creates a mock component instance
func (b *MockBridge) createComponentInstance(componentType string, element *MockElement, options map[string]interface{}) interface{} {
	// Return the options map directly as expected by tests
	if options == nil {
		options = make(map[string]interface{})
	}
	return options
}

// MockComponentInstance represents a mock FlyonUI component instance
type MockComponentInstance struct {
	ComponentType string
	Element       *MockElement
	Options       map[string]interface{}
	Initialized   bool
	Destroyed     bool
	EventHandlers map[string][]func()
}

// Destroy destroys the component instance
func (c *MockComponentInstance) Destroy() {
	c.Destroyed = true
	c.Initialized = false
	c.EventHandlers = nil
}

// Show shows the component (for components that support it)
func (c *MockComponentInstance) Show() {
	if c.Destroyed {
		return
	}
	c.Element.AddClass("show")
	c.Element.SetAttribute("aria-hidden", "false")
}

// Hide hides the component (for components that support it)
func (c *MockComponentInstance) Hide() {
	if c.Destroyed {
		return
	}
	c.Element.RemoveClass("show")
	c.Element.SetAttribute("aria-hidden", "true")
}

// Toggle toggles the component visibility
func (c *MockComponentInstance) Toggle() {
	if c.Destroyed {
		return
	}
	if c.Element.HasClass("show") {
		c.Hide()
	} else {
		c.Show()
	}
}

// Bridge interface implementations

// ConvertToJSValue converts a Go value to a mock JS value
func (b *MockBridge) ConvertToJSValue(value interface{}) bridge.JSValue {
	return &MockJSValue{
		value: value,
		props: make(map[string]interface{}),
	}
}

// ConvertStringSliceToJSArray converts a string slice to a mock JS array
func (b *MockBridge) ConvertStringSliceToJSArray(slice []string) bridge.JSValue {
	// Convert []string to []interface{} to match test expectations
	array := make([]interface{}, len(slice))
	for i, s := range slice {
		array[i] = s
	}
	return &MockJSValue{
		value: array,
		props: make(map[string]interface{}),
	}
}

// ConvertMapToJSObject converts a map to a mock JS object
func (b *MockBridge) ConvertMapToJSObject(m map[string]interface{}) bridge.JSValue {
	return &MockJSValue{
		value: m,
		props: m,
	}
}

// InitializeFlyonComponents initializes FlyonUI components
func (b *MockBridge) InitializeFlyonComponents(components []string) error {
	if !b.dom.IsReady() {
		return errors.New("DOM not ready")
	}
	
	// If empty list, initialize all known components
	if len(components) == 0 {
		components = []string{"HSDropdown", "HSModal", "HSCollapse", "HSTabs", "HSAccordion", "HSCarousel", "HSTooltip", "HSPopover"}
	}
	
	// Initialize each component type
	for _, componentType := range components {
		b.initializeComponent(componentType, make(map[string]interface{}))
	}
	
	return nil
}

// InitializeDropdown initializes dropdown components
func (b *MockBridge) InitializeDropdown(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSDropdown", selector, options)
}

// InitializeModal initializes modal components
func (b *MockBridge) InitializeModal(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSModal", selector, options)
}

// InitializeCollapse initializes collapse components
func (b *MockBridge) InitializeCollapse(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSCollapse", selector, options)
}

// InitializeTabs initializes tabs components
func (b *MockBridge) InitializeTabs(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSTabs", selector, options)
}

// InitializeAccordion initializes accordion components
func (b *MockBridge) InitializeAccordion(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSAccordion", selector, options)
}

// InitializeCarousel initializes carousel components
func (b *MockBridge) InitializeCarousel(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSCarousel", selector, options)
}

// InitializeTooltip initializes tooltip components
func (b *MockBridge) InitializeTooltip(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSTooltip", selector, options)
}

// InitializePopover initializes popover components
func (b *MockBridge) InitializePopover(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSPopover", selector, options)
}

// InitializeOffcanvas initializes offcanvas components
func (b *MockBridge) InitializeOffcanvas(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSOffcanvas", selector, options)
}

// InitializeScrollspy initializes scrollspy components
func (b *MockBridge) InitializeScrollspy(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSScrollspy", selector, options)
}

// InitializeSelect initializes select components
func (b *MockBridge) InitializeSelect(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSSelect", selector, options)
}

// InitializeInputNumber initializes input number components
func (b *MockBridge) InitializeInputNumber(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSInputNumber", selector, options)
}

// InitializeFileUpload initializes file upload components
func (b *MockBridge) InitializeFileUpload(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSFileUpload", selector, options)
}

// InitializeDatepicker initializes datepicker components
func (b *MockBridge) InitializeDatepicker(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSDatepicker", selector, options)
}

// InitializeTimepicker initializes timepicker components
func (b *MockBridge) InitializeTimepicker(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSTimepicker", selector, options)
}

// InitializeColorpicker initializes colorpicker components
func (b *MockBridge) InitializeColorpicker(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSColorpicker", selector, options)
}

// InitializeRange initializes range components
func (b *MockBridge) InitializeRange(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSRange", selector, options)
}

// InitializeTogglePassword initializes toggle password components
func (b *MockBridge) InitializeTogglePassword(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSTogglePassword", selector, options)
}

// InitializePinInput initializes pin input components
func (b *MockBridge) InitializePinInput(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSPinInput", selector, options)
}

// InitializeCopyMarkup initializes copy markup components
func (b *MockBridge) InitializeCopyMarkup(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSCopyMarkup", selector, options)
}

// InitializeComboBox initializes combo box components
func (b *MockBridge) InitializeComboBox(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSComboBox", selector, options)
}

// InitializeDataTable initializes data table components
func (b *MockBridge) InitializeDataTable(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSDataTable", selector, options)
}

// InitializeThemeSwitch initializes theme switch components
func (b *MockBridge) InitializeThemeSwitch(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSThemeSwitch", selector, options)
}

// InitializeToggleCount initializes toggle count components
func (b *MockBridge) InitializeToggleCount(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSToggleCount", selector, options)
}

// InitializeRemoveElement initializes remove element components
func (b *MockBridge) InitializeRemoveElement(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSRemoveElement", selector, options)
}

// InitializeStepForm initializes step form components
func (b *MockBridge) InitializeStepForm(selector string, options map[string]interface{}) error {
	return b.initializeComponentBySelector("HSStepForm", selector, options)
}

// initializeComponentBySelector is a helper for component initialization
func (b *MockBridge) initializeComponentBySelector(componentType, selector string, options map[string]interface{}) error {
	if !b.dom.IsReady() {
		return errors.New("DOM not ready")
	}
	
	b.initializeSpecificComponent(componentType, selector, options)
	return nil
}

// DestroyComponent destroys a component instance
func (b *MockBridge) DestroyComponent(element bridge.DOMElement, componentType string) error {
	mockEl, ok := element.(*MockDOMElement)
	if !ok {
		return errors.New("invalid element type")
	}
	
	instance := mockEl.element.GetComponentInstance(componentType)
	if instance == nil {
		return errors.New("component instance not found")
	}
	
	// Clear the component instance from the element
	mockEl.element.SetComponentInstance(componentType, nil)
	
	// Remove from tracking
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.components[componentType] != nil && mockEl.element.ID != "" {
		delete(b.components[componentType], mockEl.element.ID)
	}
	
	return nil
}

// GetComponentInstance retrieves a component instance
func (b *MockBridge) GetComponentInstance(element bridge.DOMElement, componentType string) (bridge.JSValue, error) {
	mockEl, ok := element.(*MockDOMElement)
	if !ok {
		return nil, errors.New("invalid element type")
	}
	
	instance := mockEl.element.GetComponentInstance(componentType)
	if instance == nil {
		return nil, errors.New("component instance not found")
	}
	
	return b.ConvertToJSValue(instance), nil
}

// AddEventListener adds an event listener
func (b *MockBridge) AddEventListener(element bridge.DOMElement, eventType string, handler func()) error {
	mockEl, ok := element.(*MockDOMElement)
	if !ok {
		return errors.New("invalid element type")
	}
	
	mockEl.element.AddEventListener(eventType, func(event *MockEvent) {
		handler()
	})
	
	// Track the event listener in the bridge
	b.mu.Lock()
	defer b.mu.Unlock()
	
	key := fmt.Sprintf("%s:%s", mockEl.element.ID, eventType)
	b.eventListeners[key] = append(b.eventListeners[key], handler)
	
	return nil
}

// RemoveEventListener removes an event listener
func (b *MockBridge) RemoveEventListener(element bridge.DOMElement, eventType string, handler func()) error {
	mockEl, ok := element.(*MockDOMElement)
	if !ok {
		return errors.New("invalid element type")
	}
	
	// Remove from bridge tracking
	b.mu.Lock()
	defer b.mu.Unlock()
	
	key := fmt.Sprintf("%s:%s", mockEl.element.ID, eventType)
	if _, exists := b.eventListeners[key]; exists {
		// For simplicity, remove all handlers for this event type
		// In a real implementation, you'd need to match the specific handler
		delete(b.eventListeners, key)
		
		// Also remove from the element
		mockEl.element.RemoveEventListener(eventType, func(event *MockEvent) {
			handler()
		})
	}
	
	return nil
}

// WaitForDOMReady waits for DOM to be ready and executes callback
func (b *MockBridge) WaitForDOMReady(callback func()) {
	b.dom.OnReady(callback)
}

// InitializeAllComponents initializes all FlyonUI components
func (b *MockBridge) InitializeAllComponents() error {
	components := []string{
		"HSDropdown", "HSModal", "HSCollapse", "HSTabs", "HSAccordion",
		"HSCarousel", "HSTooltip", "HSPopover", "HSOffcanvas", "HSScrollspy",
		"HSSelect", "HSInputNumber", "HSFileUpload", "HSDatepicker",
		"HSTimepicker", "HSColorpicker", "HSRange", "HSTogglePassword",
		"HSPinInput", "HSCopyMarkup", "HSComboBox", "HSDataTable",
		"HSThemeSwitch", "HSToggleCount", "HSRemoveElement", "HSStepForm",
	}
	
	return b.InitializeFlyonComponents(components)
}

// QuerySelector finds an element by selector
func (b *MockBridge) QuerySelector(selector string) bridge.DOMElement {
	element := b.dom.QuerySelector(selector)
	if element == nil {
		return nil
	}
	return &MockDOMElement{element: element, bridge: b}
}

// QuerySelectorAll finds all elements by selector
func (b *MockBridge) QuerySelectorAll(selector string) []bridge.DOMElement {
	elements := b.dom.QuerySelectorAll(selector)
	result := make([]bridge.DOMElement, len(elements))
	for i, element := range elements {
		result[i] = &MockDOMElement{element: element, bridge: b}
	}
	return result
}

// GetDOM returns the underlying mock DOM for testing
func (b *MockBridge) GetDOM() *MockDOM {
	return b.dom
}

// SetDOMReady marks the DOM as ready
func (b *MockBridge) SetDOMReady() {
	b.dom.SetReady()
}

// MockJSValue implementations

// Get gets a property from the JS value
func (v *MockJSValue) Get(key string) bridge.JSValue {
	if v.props == nil {
		v.props = make(map[string]interface{})
	}
	
	if val, exists := v.props[key]; exists {
		return &MockJSValue{value: val, props: make(map[string]interface{})}
	}
	
	// Handle special cases for common JS properties
	switch key {
	 case "length":
		if slice, ok := v.value.([]string); ok {
			return &MockJSValue{value: len(slice), props: make(map[string]interface{})}
		}
		if arr, ok := v.value.([]interface{}); ok {
			return &MockJSValue{value: len(arr), props: make(map[string]interface{})}
		}
	}
	
	return &MockJSValue{value: nil, props: make(map[string]interface{})}
}

// Index gets an element from an array-like JS value
func (v *MockJSValue) Index(i int) bridge.JSValue {
	if slice, ok := v.value.([]string); ok {
		if i >= 0 && i < len(slice) {
			return &MockJSValue{value: slice[i], props: make(map[string]interface{})}
		}
	}
	if arr, ok := v.value.([]interface{}); ok {
		if i >= 0 && i < len(arr) {
			return &MockJSValue{value: arr[i], props: make(map[string]interface{})}
		}
	}
	return &MockJSValue{value: nil, props: make(map[string]interface{})}
}

// SetIndex sets an element in an array-like JS value
func (v *MockJSValue) SetIndex(i int, value interface{}) {
	if arr, ok := v.value.([]interface{}); ok {
		if i >= 0 && i < len(arr) {
			arr[i] = value
		}
	}
}

// Set sets a property on the JS value
func (v *MockJSValue) Set(key string, value interface{}) {
	if v.props == nil {
		v.props = make(map[string]interface{})
	}
	v.props[key] = value
}

// Call calls the JS value as a function
func (v *MockJSValue) Call(method string, args ...interface{}) bridge.JSValue {
	// Mock function calls - return a mock result
	return &MockJSValue{value: "mock_result", props: make(map[string]interface{})}
}

// New creates a new instance using the JS value as a constructor
func (v *MockJSValue) New(args ...interface{}) bridge.JSValue {
	return &MockJSValue{value: fmt.Sprintf("new instance with %d args", len(args)), props: make(map[string]interface{})}
}

// Invoke invokes the JS value as a function
func (v *MockJSValue) Invoke(args ...interface{}) bridge.JSValue {
	// Mock function invocation
	if fn, ok := v.value.(func(...interface{}) interface{}); ok {
		result := fn(args...)
		return &MockJSValue{value: result, props: make(map[string]interface{})}
	}
	return &MockJSValue{value: "mock_invoke_result", props: make(map[string]interface{})}
}

// IsUndefined checks if the value is undefined
func (v *MockJSValue) IsUndefined() bool {
	return v.value == nil
}

// IsNull checks if the value is null
func (v *MockJSValue) IsNull() bool {
	return v.value == nil
}

// Bool returns the boolean value
func (v *MockJSValue) Bool() bool {
	if b, ok := v.value.(bool); ok {
		return b
	}
	return false
}

// Int returns the integer value
func (v *MockJSValue) Int() int {
	if i, ok := v.value.(int); ok {
		return i
	}
	if f, ok := v.value.(float64); ok {
		return int(f)
	}
	return 0
}

// Float returns the float value
func (v *MockJSValue) Float() float64 {
	if f, ok := v.value.(float64); ok {
		return f
	}
	if i, ok := v.value.(int); ok {
		return float64(i)
	}
	return 0.0
}

// String returns the string value
func (v *MockJSValue) String() string {
	if s, ok := v.value.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v.value)
}

// Truthy returns whether the value is truthy
func (v *MockJSValue) Truthy() bool {
	if v.value == nil {
		return false
	}
	if b, ok := v.value.(bool); ok {
		return b
	}
	if s, ok := v.value.(string); ok {
		return s != ""
	}
	if i, ok := v.value.(int); ok {
		return i != 0
	}
	if f, ok := v.value.(float64); ok {
		return f != 0.0
	}
	return true
}

// Underlying returns the underlying value
func (v *MockJSValue) Underlying() interface{} {
	return v.value
}

// MockDOMElement implementations

// GetAttribute gets an attribute from the element
func (e *MockDOMElement) GetAttribute(name string) string {
	return e.element.GetAttribute(name)
}

// SetAttribute sets an attribute on the element
func (e *MockDOMElement) SetAttribute(name, value string) {
	e.element.SetAttribute(name, value)
}

// HasAttribute checks if the element has an attribute
func (e *MockDOMElement) HasAttribute(name string) bool {
	return e.element.HasAttribute(name)
}

// RemoveAttribute removes an attribute from the element
func (e *MockDOMElement) RemoveAttribute(name string) {
	e.element.RemoveAttribute(name)
}

// GetID returns the element's ID
func (e *MockDOMElement) GetID() string {
	return e.element.ID
}

// SetID sets the element's ID
func (e *MockDOMElement) SetID(id string) {
	e.element.ID = id
	e.element.SetAttribute("id", id)
	// Add to DOM's element map
	e.bridge.dom.AddElement(e.element)
}

// GetClassName returns the element's class name
func (e *MockDOMElement) GetClassName() string {
	return e.element.ClassName
}

// SetClassName sets the element's class name
func (e *MockDOMElement) SetClassName(className string) {
	e.element.ClassName = className
	e.element.SetAttribute("class", className)
}

// AddClass adds a CSS class
func (e *MockDOMElement) AddClass(className string) {
	e.element.AddClass(className)
}

// RemoveClass removes a CSS class
func (e *MockDOMElement) RemoveClass(className string) {
	e.element.RemoveClass(className)
}

// HasClass checks if the element has a CSS class
func (e *MockDOMElement) HasClass(className string) bool {
	return e.element.HasClass(className)
}

// ToggleClass toggles a CSS class
func (e *MockDOMElement) ToggleClass(className string) {
	e.element.ToggleClass(className)
}

// GetTextContent returns the element's text content
func (e *MockDOMElement) GetTextContent() string {
	return e.element.TextContent
}

// SetTextContent sets the element's text content
func (e *MockDOMElement) SetTextContent(text string) {
	e.element.TextContent = text
}

// GetInnerHTML returns the element's inner HTML
func (e *MockDOMElement) GetInnerHTML() string {
	return e.element.InnerHTML
}

// SetInnerHTML sets the element's inner HTML
func (e *MockDOMElement) SetInnerHTML(html string) {
	e.element.InnerHTML = html
}

// QuerySelector finds a descendant element
func (e *MockDOMElement) QuerySelector(selector string) bridge.DOMElement {
	element := e.element.QuerySelector(selector)
	if element == nil {
		return nil
	}
	return &MockDOMElement{element: element, bridge: e.bridge}
}

// QuerySelectorAll finds all descendant elements
func (e *MockDOMElement) QuerySelectorAll(selector string) []bridge.DOMElement {
	elements := e.element.QuerySelectorAll(selector)
	result := make([]bridge.DOMElement, len(elements))
	for i, element := range elements {
		result[i] = &MockDOMElement{element: element, bridge: e.bridge}
	}
	return result
}

// AppendChild appends a child element
func (e *MockDOMElement) AppendChild(child bridge.DOMElement) {
	if mockChild, ok := child.(*MockDOMElement); ok {
		e.element.AppendChild(mockChild.element)
	}
}

// RemoveChild removes a child element
func (e *MockDOMElement) RemoveChild(child bridge.DOMElement) error {
	if mockChild, ok := child.(*MockDOMElement); ok {
		return e.element.RemoveChild(mockChild.element)
	}
	return errors.New("invalid child element type")
}

// Click simulates a click on the element
func (e *MockDOMElement) Click() {
	e.element.Click()
}

// Focus simulates focusing the element
func (e *MockDOMElement) Focus() {
	e.element.Focus()
}

// Blur simulates blurring the element
func (e *MockDOMElement) Blur() {
	e.element.Blur()
}

// Underlying returns the underlying element
func (e *MockDOMElement) Underlying() interface{} {
	return e.element
}

// AddEventListener adds an event listener
func (e *MockDOMElement) AddEventListener(eventType string, handler func()) {
	e.element.AddEventListener(eventType, func(event *MockEvent) {
		handler()
	})
}

// RemoveEventListener removes an event listener
func (e *MockDOMElement) RemoveEventListener(eventType string, handler func()) {
	// Simplified implementation
	e.element.RemoveEventListener(eventType, func(event *MockEvent) {
		handler()
	})
}

// DispatchEvent dispatches an event
func (e *MockDOMElement) DispatchEvent(eventType string, data map[string]interface{}) {
	event := NewMockEvent(eventType, e.element)
	if data != nil {
		event.Data = data
	}
	e.element.DispatchEvent(event)
}

// Closest finds the closest ancestor element matching the selector
func (e *MockDOMElement) Closest(selector string) bridge.DOMElement {
	current := e.element.Parent
	for current != nil {
		if current.MatchesSelector(selector) {
			return &MockDOMElement{element: current, bridge: e.bridge}
		}
		current = current.Parent
	}
	return nil
}

// GetParent returns the parent element
func (e *MockDOMElement) GetParent() bridge.DOMElement {
	if e.element.Parent != nil {
		return &MockDOMElement{element: e.element.Parent, bridge: e.bridge}
	}
	return nil
}

// GetMockElement returns the underlying mock element for testing
func (e *MockDOMElement) GetMockElement() *MockElement {
	return e.element
}