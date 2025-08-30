//go:build !js || !wasm

package mockdom

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// MockDOM represents a mock DOM implementation for testing
type MockDOM struct {
	mu        sync.RWMutex
	elements  map[string]*MockElement // for ID-based lookup
	allElements []*MockElement         // for storing all elements
	listeners map[string][]EventListener
	ready     bool
	readyCallbacks []func()
}

// MockElement represents a mock DOM element
type MockElement struct {
	ID         string
	TagName    string
	ClassName  string
	Attributes map[string]string
	Children   []*MockElement
	Parent     *MockElement
	TextContent string
	InnerHTML   string
	Style       map[string]string
	Dataset     map[string]string
	Listeners   map[string][]EventListener
	Component   interface{} // For storing component instances
}

// EventListener represents an event listener function
type EventListener struct {
	Type    string
	Handler func(*MockEvent)
	Once    bool
}

// MockEvent represents a mock DOM event
type MockEvent struct {
	Type           string
	Target         *MockElement
	CurrentTarget  *MockElement
	Bubbles        bool
	Cancelable     bool
	DefaultPrevented bool
	Data           map[string]interface{}
}

// NewMockDOM creates a new mock DOM instance
func NewMockDOM() *MockDOM {
	return &MockDOM{
		elements:       make(map[string]*MockElement),
		allElements:    make([]*MockElement, 0),
		listeners:      make(map[string][]EventListener),
		ready:          false,
		readyCallbacks: make([]func(), 0),
	}
}

// NewMockElement creates a new mock element
func NewMockElement(tagName, id string) *MockElement {
	return &MockElement{
		ID:         id,
		TagName:    strings.ToUpper(tagName),
		ClassName:  "",
		Attributes: make(map[string]string),
		Children:   make([]*MockElement, 0),
		Parent:     nil,
		TextContent: "",
		InnerHTML:   "",
		Style:       make(map[string]string),
		Dataset:     make(map[string]string),
		Listeners:   make(map[string][]EventListener),
	}
}

// SetReady marks the DOM as ready and triggers callbacks
func (dom *MockDOM) SetReady() {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	dom.ready = true
	for _, callback := range dom.readyCallbacks {
		callback()
	}
	dom.readyCallbacks = nil
}

// IsReady returns whether the DOM is ready
func (dom *MockDOM) IsReady() bool {
	dom.mu.RLock()
	defer dom.mu.RUnlock()
	return dom.ready
}

// OnReady adds a callback to be called when DOM is ready
func (dom *MockDOM) OnReady(callback func()) {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	if dom.ready {
		callback()
		return
	}
	
	dom.readyCallbacks = append(dom.readyCallbacks, callback)
}

// CreateElement creates a new element and adds it to the DOM
func (dom *MockDOM) CreateElement(tagName string) *MockElement {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	element := NewMockElement(tagName, "")
	return element
}

// GetElementByID retrieves an element by its ID
func (dom *MockDOM) GetElementByID(id string) *MockElement {
	dom.mu.RLock()
	defer dom.mu.RUnlock()
	
	return dom.elements[id]
}

// AddElement adds an element to the DOM
func (dom *MockDOM) AddElement(element *MockElement) {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	// Store all elements in the slice
	dom.allElements = append(dom.allElements, element)
	
	// Also store elements with IDs in the map for quick lookup
	if element.ID != "" {
		dom.elements[element.ID] = element
	}
}

// QuerySelector finds the first element matching the selector
func (dom *MockDOM) QuerySelector(selector string) *MockElement {
	dom.mu.RLock()
	defer dom.mu.RUnlock()
	
	// Simple selector parsing - supports ID, class, attribute, and tag selectors
	if strings.HasPrefix(selector, "#") {
		id := selector[1:]
		return dom.elements[id]
	}
	
	if strings.HasPrefix(selector, ".") {
		className := selector[1:]
		for _, element := range dom.allElements {
			if dom.hasClass(element, className) {
				return element
			}
		}
	}
	
	if strings.HasPrefix(selector, "[") && strings.HasSuffix(selector, "]") {
		// Attribute selector like [data-hs-dropdown]
		attrName := selector[1 : len(selector)-1]
		for _, element := range dom.allElements {
			if element.HasAttribute(attrName) {
				return element
			}
		}
	}
	
	// Tag selector
	tagName := strings.ToUpper(selector)
	for _, element := range dom.allElements {
		if element.TagName == tagName {
			return element
		}
	}
	
	return nil
}

// QuerySelectorAll finds all elements matching the selector
func (dom *MockDOM) QuerySelectorAll(selector string) []*MockElement {
	dom.mu.RLock()
	defer dom.mu.RUnlock()
	
	var results []*MockElement
	
	// Simple selector parsing
	if strings.HasPrefix(selector, "#") {
		id := selector[1:]
		if element := dom.elements[id]; element != nil {
			results = append(results, element)
		}
	} else if strings.HasPrefix(selector, ".") {
		className := selector[1:]
		for _, element := range dom.allElements {
			if dom.hasClass(element, className) {
				results = append(results, element)
			}
		}
	} else if strings.HasPrefix(selector, "[") && strings.HasSuffix(selector, "]") {
		// Attribute selector like [data-hs-dropdown]
		attrName := selector[1 : len(selector)-1]
		for _, element := range dom.allElements {
			if element.HasAttribute(attrName) {
				results = append(results, element)
			}
		}
	} else {
		// Tag selector
		tagName := strings.ToUpper(selector)
		for _, element := range dom.allElements {
			if element.TagName == tagName {
				results = append(results, element)
			}
		}
	}
	
	return results
}

// hasClass checks if an element has a specific class
func (dom *MockDOM) hasClass(element *MockElement, className string) bool {
	classes := strings.Fields(element.ClassName)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

// DispatchEvent dispatches an event on the DOM
func (dom *MockDOM) DispatchEvent(event *MockEvent) {
	dom.mu.RLock()
	listeners := dom.listeners[event.Type]
	dom.mu.RUnlock()
	
	for _, listener := range listeners {
		listener.Handler(event)
		if listener.Once {
			dom.RemoveEventListener(event.Type, listener.Handler)
		}
	}
}

// AddEventListener adds a global event listener
func (dom *MockDOM) AddEventListener(eventType string, handler func(*MockEvent)) {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	dom.listeners[eventType] = append(dom.listeners[eventType], EventListener{
		Type:    eventType,
		Handler: handler,
		Once:    false,
	})
}

// RemoveEventListener removes a global event listener
func (dom *MockDOM) RemoveEventListener(eventType string, handler func(*MockEvent)) {
	dom.mu.Lock()
	defer dom.mu.Unlock()
	
	listeners := dom.listeners[eventType]
	for i, listener := range listeners {
		// Note: This is a simplified comparison - in practice, you'd need
		// to store function references properly
		if listener.Type == eventType {
			dom.listeners[eventType] = append(listeners[:i], listeners[i+1:]...)
			break
		}
	}
}

// Element methods

// SetAttribute sets an attribute on the element
func (e *MockElement) SetAttribute(name, value string) {
	e.Attributes[name] = value
	if name == "id" {
		e.ID = value
	} else if name == "class" {
		e.ClassName = value
	}
}

// GetAttribute gets an attribute from the element
func (e *MockElement) GetAttribute(name string) string {
	return e.Attributes[name]
}

// HasAttribute checks if the element has an attribute
func (e *MockElement) HasAttribute(name string) bool {
	_, exists := e.Attributes[name]
	return exists
}

// RemoveAttribute removes an attribute from the element
func (e *MockElement) RemoveAttribute(name string) {
	delete(e.Attributes, name)
	if name == "id" {
		e.ID = ""
	} else if name == "class" {
		e.ClassName = ""
	}
}

// AddClass adds a CSS class to the element
func (e *MockElement) AddClass(className string) {
	classes := strings.Fields(e.ClassName)
	for _, class := range classes {
		if class == className {
			return // Already has the class
		}
	}
	classes = append(classes, className)
	e.ClassName = strings.Join(classes, " ")
	e.SetAttribute("class", e.ClassName)
}

// RemoveClass removes a CSS class from the element
func (e *MockElement) RemoveClass(className string) {
	classes := strings.Fields(e.ClassName)
	var newClasses []string
	for _, class := range classes {
		if class != className {
			newClasses = append(newClasses, class)
		}
	}
	e.ClassName = strings.Join(newClasses, " ")
	e.SetAttribute("class", e.ClassName)
}

// HasClass checks if the element has a specific CSS class
func (e *MockElement) HasClass(className string) bool {
	classes := strings.Fields(e.ClassName)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

// ToggleClass toggles a CSS class on the element
func (e *MockElement) ToggleClass(className string) {
	if e.HasClass(className) {
		e.RemoveClass(className)
	} else {
		e.AddClass(className)
	}
}

// AppendChild adds a child element
func (e *MockElement) AppendChild(child *MockElement) {
	child.Parent = e
	e.Children = append(e.Children, child)
}

// RemoveChild removes a child element
func (e *MockElement) RemoveChild(child *MockElement) error {
	for i, c := range e.Children {
		if c == child {
			child.Parent = nil
			e.Children = append(e.Children[:i], e.Children[i+1:]...)
			return nil
		}
	}
	return errors.New("child not found")
}

// QuerySelector finds the first descendant matching the selector
func (e *MockElement) QuerySelector(selector string) *MockElement {
	// Simple implementation - check children recursively
	for _, child := range e.Children {
		if e.matchesSelector(child, selector) {
			return child
		}
		if result := child.QuerySelector(selector); result != nil {
			return result
		}
	}
	return nil
}

// QuerySelectorAll finds all descendants matching the selector
func (e *MockElement) QuerySelectorAll(selector string) []*MockElement {
	var results []*MockElement
	for _, child := range e.Children {
		if e.matchesSelector(child, selector) {
			results = append(results, child)
		}
		results = append(results, child.QuerySelectorAll(selector)...)
	}
	return results
}

// matchesSelector checks if an element matches a selector
func (e *MockElement) matchesSelector(element *MockElement, selector string) bool {
	if strings.HasPrefix(selector, "#") {
		return element.ID == selector[1:]
	}
	if strings.HasPrefix(selector, ".") {
		return element.HasClass(selector[1:])
	}
	return element.TagName == strings.ToUpper(selector)
}

// AddEventListener adds an event listener to the element
func (e *MockElement) AddEventListener(eventType string, handler func(*MockEvent)) {
	e.Listeners[eventType] = append(e.Listeners[eventType], EventListener{
		Type:    eventType,
		Handler: handler,
		Once:    false,
	})
}

// RemoveEventListener removes an event listener from the element
func (e *MockElement) RemoveEventListener(eventType string, handler func(*MockEvent)) {
	listeners := e.Listeners[eventType]
	for i, listener := range listeners {
		if listener.Type == eventType {
			e.Listeners[eventType] = append(listeners[:i], listeners[i+1:]...)
			break
		}
	}
}

// DispatchEvent dispatches an event on the element
func (e *MockElement) DispatchEvent(event *MockEvent) {
	// Set target only if not already set (for initial dispatch)
	if event.Target == nil {
		event.Target = e
	}
	// Always set current target to the element currently handling the event
	event.CurrentTarget = e
	
	listeners := e.Listeners[event.Type]
	for _, listener := range listeners {
		listener.Handler(event)
		if listener.Once {
			e.RemoveEventListener(event.Type, listener.Handler)
		}
	}
	
	// Bubble up to parent if event bubbles
	if event.Bubbles && e.Parent != nil && !event.DefaultPrevented {
		e.Parent.DispatchEvent(event)
	}
}

// Click simulates a click event on the element
func (e *MockElement) Click() {
	event := &MockEvent{
		Type:       "click",
		Target:     e,
		Bubbles:    true,
		Cancelable: true,
		Data:       make(map[string]interface{}),
	}
	e.DispatchEvent(event)
}

// Focus simulates a focus event on the element
func (e *MockElement) Focus() {
	event := &MockEvent{
		Type:       "focus",
		Target:     e,
		Bubbles:    false,
		Cancelable: false,
		Data:       make(map[string]interface{}),
	}
	e.DispatchEvent(event)
}

// Blur simulates a blur event on the element
func (e *MockElement) Blur() {
	event := &MockEvent{
		Type:       "blur",
		Target:     e,
		Bubbles:    false,
		Cancelable: false,
		Data:       make(map[string]interface{}),
	}
	e.DispatchEvent(event)
}

// SetComponentInstance stores a component instance on the element
func (e *MockElement) SetComponentInstance(componentType string, instance interface{}) {
	if e.Component == nil {
		e.Component = make(map[string]interface{})
	}
	if compMap, ok := e.Component.(map[string]interface{}); ok {
		compMap[componentType] = instance
	} else {
		e.Component = map[string]interface{}{componentType: instance}
	}
}

// GetComponentInstance retrieves a component instance from the element
func (e *MockElement) GetComponentInstance(componentType string) interface{} {
	if e.Component == nil {
		return nil
	}
	if compMap, ok := e.Component.(map[string]interface{}); ok {
		return compMap[componentType]
	}
	return nil
}

// String returns a string representation of the element
func (e *MockElement) String() string {
	attrs := make([]string, 0, len(e.Attributes))
	for name, value := range e.Attributes {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, name, value))
	}
	
	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}
	
	return fmt.Sprintf("<%s%s>", strings.ToLower(e.TagName), attrStr)
}

// MatchesSelector checks if the element matches the given CSS selector
func (e *MockElement) MatchesSelector(selector string) bool {
	if selector == "" {
		return false
	}
	
	// ID selector
	if strings.HasPrefix(selector, "#") {
		return e.ID == selector[1:]
	}
	
	// Class selector
	if strings.HasPrefix(selector, ".") {
		className := selector[1:]
		return e.HasClass(className)
	}
	
	// Tag selector
	return strings.ToUpper(selector) == e.TagName
}

// NewMockEvent creates a new mock event
func NewMockEvent(eventType string, target *MockElement) *MockEvent {
	return &MockEvent{
		Type:       eventType,
		Target:     target,
		Bubbles:    true,
		Cancelable: true,
		Data:       make(map[string]interface{}),
	}
}

// PreventDefault prevents the default action of the event
func (e *MockEvent) PreventDefault() {
	e.DefaultPrevented = true
}

// StopPropagation stops the event from bubbling
func (e *MockEvent) StopPropagation() {
	e.Bubbles = false
}