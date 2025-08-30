//go:build js && wasm

package bridge

import (
	"errors"
	"fmt"
	"syscall/js"
	"honnef.co/go/js/dom/v2"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// RealJSBridge implements JSBridge using actual syscall/js
type RealJSBridge struct{}

func NewRealJSBridge() *RealJSBridge {
	return &RealJSBridge{}
}

func (b *RealJSBridge) GetGlobal(name string) JSValue {
	return &RealJSValue{value: js.Global().Get(name)}
}

func (b *RealJSBridge) CreateArray(length int) JSValue {
	return &RealJSValue{value: js.Global().Get("Array").New(length)}
}

func (b *RealJSBridge) CreateObject() JSValue {
	return &RealJSValue{value: js.Global().Get("Object").New()}
}

func (b *RealJSBridge) ValueOf(v interface{}) JSValue {
	return &RealJSValue{value: js.ValueOf(v)}
}

// RealJSValue implements JSValue using actual js.Value
type RealJSValue struct {
	value js.Value
}

func (v *RealJSValue) IsUndefined() bool {
	return v.value.IsUndefined()
}

func (v *RealJSValue) String() string {
	return v.value.String()
}

func (v *RealJSValue) Int() int {
	return v.value.Int()
}

func (v *RealJSValue) Bool() bool {
	return v.value.Bool()
}

func (v *RealJSValue) Get(key string) JSValue {
	return &RealJSValue{value: v.value.Get(key)}
}

func (v *RealJSValue) Set(key string, value interface{}) {
	v.value.Set(key, js.ValueOf(value))
}

func (v *RealJSValue) Index(i int) JSValue {
	return &RealJSValue{value: v.value.Index(i)}
}

func (v *RealJSValue) SetIndex(i int, value interface{}) {
	v.value.SetIndex(i, js.ValueOf(value))
}

func (v *RealJSValue) Invoke(args ...interface{}) JSValue {
	return &RealJSValue{value: v.value.Invoke(args...)}
}

func (v *RealJSValue) New(args ...interface{}) JSValue {
	return &RealJSValue{value: v.value.New(args...)}
}

func (v *RealJSValue) Call(method string, args ...interface{}) JSValue {
	return &RealJSValue{value: v.value.Call(method, args...)}
}

func (v *RealJSValue) Underlying() interface{} {
	return v.value
}

// RealDOMBridge implements DOMBridge using actual honnef.co/go/js/dom/v2
type RealDOMBridge struct{}

func NewRealDOMBridge() *RealDOMBridge {
	return &RealDOMBridge{}
}

func (b *RealDOMBridge) QuerySelectorAll(selector string) []DOMElement {
	doc := dom.GetWindow().Document()
	elements := doc.QuerySelectorAll(selector)
	result := make([]DOMElement, len(elements))
	for i, elem := range elements {
		result[i] = &RealDOMElement{element: elem}
	}
	return result
}

func (b *RealDOMBridge) GetElementByID(id string) DOMElement {
	doc := dom.GetWindow().Document()
	elem := doc.GetElementByID(id)
	if elem == nil {
		return nil
	}
	return &RealDOMElement{element: elem}
}

func (b *RealDOMBridge) GetDocument() DOMDocument {
	return &RealDOMDocument{doc: dom.GetWindow().Document()}
}

func (b *RealDOMBridge) AddEventListener(selector, eventType string, handler func(DOMEvent)) error {
	elements := b.QuerySelectorAll(selector)
	if len(elements) == 0 {
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	for _, elem := range elements {
		realElem := elem.(*RealDOMElement)
		realElem.element.AddEventListener(eventType, false, func(event dom.Event) {
			handler(&RealDOMEvent{event: event})
		})
	}
	
	logutil.Logf("Event listener added for %s on %d elements", eventType, len(elements))
	return nil
}

func (b *RealDOMBridge) RemoveEventListener(selector, eventType string) error {
	// Note: This is a simplified implementation
	// In practice, we'd need to store the js.Func references
	logutil.Logf("Event listener removal requested for %s on %s", eventType, selector)
	return nil
}

func (b *RealDOMBridge) WaitForReady(callback func()) {
	doc := dom.GetWindow().Document()
	readyState := doc.Underlying().Get("readyState").String()
	
	if readyState == "loading" {
		doc.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			callback()
		})
	} else {
		callback()
	}
}

// RealDOMElement implements DOMElement
type RealDOMElement struct {
	element dom.Element
}

func (e *RealDOMElement) GetAttribute(name string) string {
	return e.element.GetAttribute(name)
}

func (e *RealDOMElement) SetAttribute(name, value string) {
	e.element.SetAttribute(name, value)
}

func (e *RealDOMElement) AddClass(class string) {
	e.element.Class().Add(class)
}

func (e *RealDOMElement) RemoveClass(class string) {
	e.element.Class().Remove(class)
}

func (e *RealDOMElement) HasClass(class string) bool {
	return e.element.Class().Contains(class)
}

func (e *RealDOMElement) GetParent() DOMElement {
	parent := e.element.ParentElement()
	if parent == nil {
		return nil
	}
	return &RealDOMElement{element: parent}
}

func (e *RealDOMElement) QuerySelector(selector string) DOMElement {
	elem := e.element.QuerySelector(selector)
	if elem == nil {
		return nil
	}
	return &RealDOMElement{element: elem}
}

func (e *RealDOMElement) Closest(selector string) DOMElement {
	elem := e.element.Closest(selector)
	if elem == nil {
		return nil
	}
	return &RealDOMElement{element: elem}
}

func (e *RealDOMElement) Underlying() interface{} {
	return e.element.Underlying()
}

// RealDOMDocument implements DOMDocument
type RealDOMDocument struct {
	doc dom.Document
}

func (d *RealDOMDocument) QuerySelectorAll(selector string) []DOMElement {
	elements := d.doc.QuerySelectorAll(selector)
	result := make([]DOMElement, len(elements))
	for i, elem := range elements {
		result[i] = &RealDOMElement{element: elem}
	}
	return result
}

func (d *RealDOMDocument) GetElementByID(id string) DOMElement {
	elem := d.doc.GetElementByID(id)
	if elem == nil {
		return nil
	}
	return &RealDOMElement{element: elem}
}

func (d *RealDOMDocument) GetReadyState() string {
	return d.doc.Underlying().Get("readyState").String()
}

func (d *RealDOMDocument) AddEventListener(eventType string, handler func(DOMEvent)) {
	d.doc.AddEventListener(eventType, false, func(event dom.Event) {
		handler(&RealDOMEvent{event: event})
	})
}

// RealDOMEvent implements DOMEvent
type RealDOMEvent struct {
	event dom.Event
}

func (e *RealDOMEvent) PreventDefault() {
	e.event.PreventDefault()
}

func (e *RealDOMEvent) StopPropagation() {
	e.event.StopPropagation()
}

func (e *RealDOMEvent) GetTarget() DOMElement {
	target := e.event.Target()
	if target == nil {
		return nil
	}
	return &RealDOMElement{element: target}
}

func (e *RealDOMEvent) GetType() string {
	return e.event.Type()
}

// RealComponentBridge implements ComponentBridge
type RealComponentBridge struct {
	jsBridge JSBridge
	domBridge DOMBridge
}

func NewRealComponentBridge(jsBridge JSBridge, domBridge DOMBridge) *RealComponentBridge {
	return &RealComponentBridge{
		jsBridge: jsBridge,
		domBridge: domBridge,
	}
}

func (b *RealComponentBridge) InitializeComponent(componentName, selector string, options map[string]interface{}) error {
	logutil.Logf("Initializing %s component with selector: %s", componentName, selector)
	
	// Check if the component class is available
	componentClass := b.jsBridge.GetGlobal(componentName)
	if componentClass.IsUndefined() {
		logutil.Logf("Warning: %s not available, component may not be interactive", componentName)
		return fmt.Errorf("%s not available", componentName)
	}
	
	// Get DOM elements
	elements := b.domBridge.QuerySelectorAll(selector)
	if len(elements) == 0 {
		logutil.Logf("Warning: No elements found for selector: %s", selector)
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	// Convert options to JS object
	jsOptions := b.jsBridge.CreateObject()
	for key, value := range options {
		jsOptions.Set(key, value)
	}
	
	// Initialize component for each element
	for _, element := range elements {
		componentClass.New(element.Underlying(), jsOptions.Underlying())
	}
	
	logutil.Logf("%s component initialized successfully for %d elements", componentName, len(elements))
	return nil
}

func (b *RealComponentBridge) DestroyComponent(selector, componentType string) error {
	logutil.Logf("Destroying %s component with selector: %s", componentType, selector)
	
	elements := b.domBridge.QuerySelectorAll(selector)
	if len(elements) == 0 {
		return fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	for _, element := range elements {
		componentInstance := b.jsBridge.ValueOf(element.Underlying()).Get("hs" + componentType)
		if !componentInstance.IsUndefined() {
			destroy := componentInstance.Get("destroy")
			if !destroy.IsUndefined() {
				destroy.Invoke()
			}
		}
	}
	
	logutil.Logf("%s component destroyed successfully", componentType)
	return nil
}

func (b *RealComponentBridge) GetComponentInstance(selector, componentType string) (JSValue, error) {
	elements := b.domBridge.QuerySelectorAll(selector)
	if len(elements) == 0 {
		return nil, fmt.Errorf("no elements found for selector: %s", selector)
	}
	
	element := elements[0]
	componentInstance := b.jsBridge.ValueOf(element.Underlying()).Get("hs" + componentType)
	
	if componentInstance.IsUndefined() {
		return nil, fmt.Errorf("no %s instance found on element", componentType)
	}
	
	return componentInstance, nil
}

func (b *RealComponentBridge) InitializeAll(components []string) error {
	logutil.Logf("Initializing FlyonUI components: %v", components)
	
	// Check if HSStaticMethods is available
	hsStaticMethods := b.jsBridge.GetGlobal("HSStaticMethods")
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
	jsComponents := b.jsBridge.CreateArray(len(components))
	for i, component := range components {
		jsComponents.SetIndex(i, component)
	}
	
	// Call HSStaticMethods.autoInit with components
	autoInit.Invoke(jsComponents.Underlying())
	
	logutil.Log("FlyonUI components initialized successfully")
	return nil
}

// RealBridgeManager implements BridgeManager
type RealBridgeManager struct {
	jsBridge JSBridge
	domBridge DOMBridge
	componentBridge ComponentBridge
}

func NewRealBridgeManager() *RealBridgeManager {
	jsBridge := NewRealJSBridge()
	domBridge := NewRealDOMBridge()
	componentBridge := NewRealComponentBridge(jsBridge, domBridge)
	
	return &RealBridgeManager{
		jsBridge:        jsBridge,
		domBridge:       domBridge,
		componentBridge: componentBridge,
	}
}

// NewBridgeManager creates a new bridge manager with the provided bridges
func NewBridgeManager(jsBridge JSBridge, domBridge DOMBridge, componentBridge ComponentBridge) BridgeManager {
	return &RealBridgeManager{
		jsBridge:        jsBridge,
		domBridge:       domBridge,
		componentBridge: componentBridge,
	}
}

func (m *RealBridgeManager) JS() JSBridge {
	return m.jsBridge
}

func (m *RealBridgeManager) DOM() DOMBridge {
	return m.domBridge
}

func (m *RealBridgeManager) Component() ComponentBridge {
	return m.componentBridge
}

// GoStringsToJSArray converts a Go string slice to a JavaScript array
func (m *RealBridgeManager) GoStringsToJSArray(goSlice []string) JSValue {
	arr := m.jsBridge.CreateArray(len(goSlice))
	for i, str := range goSlice {
		arr.SetIndex(i, str)
	}
	return arr
}

// GoMapToJSObject converts a Go map to a JavaScript object
func (m *RealBridgeManager) GoMapToJSObject(goMap map[string]interface{}) JSValue {
	obj := m.jsBridge.CreateObject()
	for key, value := range goMap {
		obj.Set(key, value)
	}
	return obj
}

// InitializeFlyonComponents initializes FlyonUI components using HSStaticMethods.autoInit()
func (m *RealBridgeManager) InitializeFlyonComponents(components []string) error {
	return m.componentBridge.InitializeAll(components)
}

// InitializeSpecificComponent initializes a specific component
func (m *RealBridgeManager) InitializeSpecificComponent(componentName, selector string, options map[string]interface{}) error {
	return m.componentBridge.InitializeComponent(componentName, selector, options)
}

// InitializeAllComponents initializes all components
func (m *RealBridgeManager) InitializeAllComponents() error {
	return m.componentBridge.InitializeAll([]string{})
}

// DestroyComponent destroys a component
func (m *RealBridgeManager) DestroyComponent(selector, componentType string) error {
	return m.componentBridge.DestroyComponent(selector, componentType)
}

// GetComponentInstance gets a component instance
func (m *RealBridgeManager) GetComponentInstance(selector, componentType string) (JSValue, error) {
	return m.componentBridge.GetComponentInstance(selector, componentType)
}

// AddEventListener adds an event listener
func (m *RealBridgeManager) AddEventListener(selector, eventType string, handler func(DOMEvent)) error {
	return m.domBridge.AddEventListener(selector, eventType, handler)
}

// RemoveEventListener removes an event listener
func (m *RealBridgeManager) RemoveEventListener(selector, eventType string, handler func(DOMEvent)) error {
	// Note: The underlying DOMBridge.RemoveEventListener doesn't use the handler parameter
	// This is a limitation of the current implementation
	return m.domBridge.RemoveEventListener(selector, eventType)
}

// WaitForDOMReady waits for DOM ready
func (m *RealBridgeManager) WaitForDOMReady(callback func()) {
	m.domBridge.WaitForReady(callback)
}