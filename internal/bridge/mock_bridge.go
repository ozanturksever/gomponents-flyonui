//go:build !js || !wasm

package bridge

import (
	"fmt"
	"sync"
)

// MockJSBridge implements JSBridge for testing
type MockJSBridge struct {
	mu      sync.RWMutex
	globals map[string]JSValue
}

func NewMockJSBridge() *MockJSBridge {
	return &MockJSBridge{
		globals: make(map[string]JSValue),
	}
}

func (b *MockJSBridge) SetGlobal(name string, value JSValue) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.globals[name] = value
}

func (b *MockJSBridge) GetGlobal(name string) JSValue {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if value, exists := b.globals[name]; exists {
		return value
	}
	return &MockJSValue{isUndefined: true}
}

func (b *MockJSBridge) CreateArray(length int) JSValue {
	return &MockJSValue{
		isArray: true,
		arrayData: make([]interface{}, length),
	}
}

func (b *MockJSBridge) CreateObject() JSValue {
	return &MockJSValue{
		isObject: true,
		objectData: make(map[string]interface{}),
	}
}

func (b *MockJSBridge) ValueOf(v interface{}) JSValue {
	return &MockJSValue{
		value: v,
	}
}

// MockJSValue implements JSValue for testing
type MockJSValue struct {
	mu          sync.RWMutex
	value       interface{}
	isUndefined bool
	isArray     bool
	isObject    bool
	arrayData   []interface{}
	objectData  map[string]interface{}
	invokeCalls [][]interface{}
	newCalls    [][]interface{}
	callCalls   []CallRecord
}

type CallRecord struct {
	Method string
	Args   []interface{}
}

func (v *MockJSValue) IsUndefined() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.isUndefined
}

func (v *MockJSValue) String() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return fmt.Sprintf("%v", v.value)
}

func (v *MockJSValue) Int() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if i, ok := v.value.(int); ok {
		return i
	}
	return 0
}

func (v *MockJSValue) Bool() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if b, ok := v.value.(bool); ok {
		return b
	}
	return false
}

func (v *MockJSValue) Get(key string) JSValue {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.isObject && v.objectData != nil {
		if value, exists := v.objectData[key]; exists {
			return &MockJSValue{value: value}
		}
	}
	return &MockJSValue{isUndefined: true}
}

func (v *MockJSValue) Set(key string, value interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.isObject {
		if v.objectData == nil {
			v.objectData = make(map[string]interface{})
		}
		v.objectData[key] = value
	}
}

func (v *MockJSValue) Index(i int) JSValue {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.isArray && v.arrayData != nil && i >= 0 && i < len(v.arrayData) {
		return &MockJSValue{value: v.arrayData[i]}
	}
	return &MockJSValue{isUndefined: true}
}

func (v *MockJSValue) SetIndex(i int, value interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.isArray && v.arrayData != nil && i >= 0 && i < len(v.arrayData) {
		v.arrayData[i] = value
	}
}

func (v *MockJSValue) Invoke(args ...interface{}) JSValue {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.invokeCalls = append(v.invokeCalls, args)
	return &MockJSValue{value: "invoked"}
}

func (v *MockJSValue) New(args ...interface{}) JSValue {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.newCalls = append(v.newCalls, args)
	return &MockJSValue{value: "new_instance"}
}

func (v *MockJSValue) Call(method string, args ...interface{}) JSValue {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.callCalls = append(v.callCalls, CallRecord{Method: method, Args: args})
	return &MockJSValue{value: "called"}
}

func (v *MockJSValue) Underlying() interface{} {
	panic("Underlying() not supported in mock")
}

// Test helper methods
func (v *MockJSValue) GetInvokeCalls() [][]interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.invokeCalls
}

func (v *MockJSValue) GetNewCalls() [][]interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.newCalls
}

func (v *MockJSValue) GetCallCalls() []CallRecord {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.callCalls
}

// MockDOMBridge implements DOMBridge for testing
type MockDOMBridge struct {
	mu                sync.RWMutex
	elements          map[string][]DOMElement
	eventListeners    map[string][]EventListener
	document          DOMDocument
	readyCallbacks    []func()
	isReady           bool
}

type EventListener struct {
	Selector  string
	EventType string
	Handler   func(DOMEvent)
}

func NewMockDOMBridge() *MockDOMBridge {
	return &MockDOMBridge{
		elements:       make(map[string][]DOMElement),
		eventListeners: make(map[string][]EventListener),
		document:       &MockDOMDocument{},
		isReady:        true, // Default to ready for tests
	}
}

func (b *MockDOMBridge) AddMockElement(selector string, element DOMElement) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.elements[selector] = append(b.elements[selector], element)
}

func (b *MockDOMBridge) SetReady(ready bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.isReady = ready
	if ready {
		for _, callback := range b.readyCallbacks {
			callback()
		}
		b.readyCallbacks = nil
	}
}

func (b *MockDOMBridge) QuerySelectorAll(selector string) []DOMElement {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if elements, exists := b.elements[selector]; exists {
		return elements
	}
	return []DOMElement{}
}

func (b *MockDOMBridge) GetElementByID(id string) DOMElement {
	elements := b.QuerySelectorAll("#" + id)
	if len(elements) > 0 {
		return elements[0]
	}
	return nil
}

func (b *MockDOMBridge) GetDocument() DOMDocument {
	return b.document
}

func (b *MockDOMBridge) AddEventListener(selector, eventType string, handler func(DOMEvent)) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := selector + ":" + eventType
	b.eventListeners[key] = append(b.eventListeners[key], EventListener{
		Selector:  selector,
		EventType: eventType,
		Handler:   handler,
	})
	return nil
}

func (b *MockDOMBridge) RemoveEventListener(selector, eventType string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := selector + ":" + eventType
	delete(b.eventListeners, key)
	return nil
}

func (b *MockDOMBridge) WaitForReady(callback func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.isReady {
		callback()
	} else {
		b.readyCallbacks = append(b.readyCallbacks, callback)
	}
}

// Test helper methods
func (b *MockDOMBridge) GetEventListeners(selector, eventType string) []EventListener {
	b.mu.RLock()
	defer b.mu.RUnlock()
	key := selector + ":" + eventType
	return b.eventListeners[key]
}

func (b *MockDOMBridge) TriggerEvent(selector, eventType string) {
	b.mu.RLock()
	listeners := b.GetEventListeners(selector, eventType)
	b.mu.RUnlock()
	
	for _, listener := range listeners {
		listener.Handler(&MockDOMEvent{
			eventType: eventType,
			target:    &MockDOMElement{id: "mock-target"},
		})
	}
}

// MockDOMElement implements DOMElement for testing
type MockDOMElement struct {
	mu         sync.RWMutex
	id         string
	attributes map[string]string
	classes    map[string]bool
	parent     DOMElement
	children   map[string]DOMElement
}

func NewMockDOMElement(id string) *MockDOMElement {
	return &MockDOMElement{
		id:         id,
		attributes: make(map[string]string),
		classes:    make(map[string]bool),
		children:   make(map[string]DOMElement),
	}
}

func (e *MockDOMElement) GetAttribute(name string) string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.attributes[name]
}

func (e *MockDOMElement) SetAttribute(name, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.attributes[name] = value
}

func (e *MockDOMElement) AddClass(class string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classes[class] = true
}

func (e *MockDOMElement) RemoveClass(class string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.classes, class)
}

func (e *MockDOMElement) HasClass(class string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classes[class]
}

func (e *MockDOMElement) GetParent() DOMElement {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.parent
}

func (e *MockDOMElement) SetParent(parent DOMElement) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.parent = parent
}

func (e *MockDOMElement) QuerySelector(selector string) DOMElement {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if child, exists := e.children[selector]; exists {
		return child
	}
	return nil
}

func (e *MockDOMElement) AddChild(selector string, child DOMElement) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.children[selector] = child
}

func (e *MockDOMElement) Closest(selector string) DOMElement {
	// Simple implementation for testing
	if e.id == selector || "#"+e.id == selector {
		return e
	}
	if e.parent != nil {
		return e.parent.Closest(selector)
	}
	return nil
}

func (e *MockDOMElement) Underlying() interface{} {
	panic("Underlying() not supported in mock")
}

// Test helper methods
func (e *MockDOMElement) GetClasses() map[string]bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := make(map[string]bool)
	for k, v := range e.classes {
		result[k] = v
	}
	return result
}

func (e *MockDOMElement) GetAttributes() map[string]string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := make(map[string]string)
	for k, v := range e.attributes {
		result[k] = v
	}
	return result
}

// MockDOMDocument implements DOMDocument for testing
type MockDOMDocument struct {
	mu           sync.RWMutex
	elements     map[string][]DOMElement
	readyState   string
	eventHandlers map[string][]func(DOMEvent)
}

func NewMockDOMDocument() *MockDOMDocument {
	return &MockDOMDocument{
		elements:     make(map[string][]DOMElement),
		readyState:   "complete",
		eventHandlers: make(map[string][]func(DOMEvent)),
	}
}

func (d *MockDOMDocument) QuerySelectorAll(selector string) []DOMElement {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if elements, exists := d.elements[selector]; exists {
		return elements
	}
	return []DOMElement{}
}

func (d *MockDOMDocument) GetElementByID(id string) DOMElement {
	elements := d.QuerySelectorAll("#" + id)
	if len(elements) > 0 {
		return elements[0]
	}
	return nil
}

func (d *MockDOMDocument) GetReadyState() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.readyState
}

func (d *MockDOMDocument) SetReadyState(state string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.readyState = state
}

func (d *MockDOMDocument) AddEventListener(eventType string, handler func(DOMEvent)) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.eventHandlers[eventType] = append(d.eventHandlers[eventType], handler)
}

func (d *MockDOMDocument) AddMockElement(selector string, element DOMElement) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.elements[selector] = append(d.elements[selector], element)
}

// MockDOMEvent implements DOMEvent for testing
type MockDOMEvent struct {
	mu                sync.RWMutex
	eventType         string
	target            DOMElement
	preventedDefault  bool
	stoppedPropagation bool
}

func NewMockDOMEvent(eventType string, target DOMElement) *MockDOMEvent {
	return &MockDOMEvent{
		eventType: eventType,
		target:    target,
	}
}

func (e *MockDOMEvent) PreventDefault() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.preventedDefault = true
}

func (e *MockDOMEvent) StopPropagation() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.stoppedPropagation = true
}

func (e *MockDOMEvent) GetTarget() DOMElement {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.target
}

func (e *MockDOMEvent) GetType() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.eventType
}

// Test helper methods
func (e *MockDOMEvent) IsDefaultPrevented() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.preventedDefault
}

func (e *MockDOMEvent) IsPropagationStopped() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.stoppedPropagation
}

// MockComponentBridge implements ComponentBridge for testing
type MockComponentBridge struct {
	mu                   sync.RWMutex
	initializedComponents map[string][]string // componentName -> selectors
	destroyedComponents   map[string][]string // componentName -> selectors
	componentInstances    map[string]JSValue  // selector -> instance
	initializeAllCalls    [][]string
	errorOnInit           map[string]error // componentName -> error
}

func NewMockComponentBridge() *MockComponentBridge {
	return &MockComponentBridge{
		initializedComponents: make(map[string][]string),
		destroyedComponents:   make(map[string][]string),
		componentInstances:    make(map[string]JSValue),
		errorOnInit:           make(map[string]error),
	}
}

func (b *MockComponentBridge) SetInitError(componentName string, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.errorOnInit[componentName] = err
}

func (b *MockComponentBridge) InitializeComponent(componentName, selector string, options map[string]interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if err, exists := b.errorOnInit[componentName]; exists {
		return err
	}
	
	b.initializedComponents[componentName] = append(b.initializedComponents[componentName], selector)
	b.componentInstances[selector] = &MockJSValue{value: "mock_instance_" + componentName}
	return nil
}

func (b *MockComponentBridge) DestroyComponent(selector, componentType string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.destroyedComponents[componentType] = append(b.destroyedComponents[componentType], selector)
	delete(b.componentInstances, selector)
	return nil
}

func (b *MockComponentBridge) GetComponentInstance(selector, componentType string) (JSValue, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if instance, exists := b.componentInstances[selector]; exists {
		return instance, nil
	}
	return nil, fmt.Errorf("no %s instance found for selector %s", componentType, selector)
}

func (b *MockComponentBridge) InitializeAll(components []string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.initializeAllCalls = append(b.initializeAllCalls, components)
	return nil
}

// Test helper methods
func (b *MockComponentBridge) GetInitializedComponents() map[string][]string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	result := make(map[string][]string)
	for k, v := range b.initializedComponents {
		result[k] = append([]string{}, v...)
	}
	return result
}

func (b *MockComponentBridge) GetDestroyedComponents() map[string][]string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	result := make(map[string][]string)
	for k, v := range b.destroyedComponents {
		result[k] = append([]string{}, v...)
	}
	return result
}

func (b *MockComponentBridge) GetInitializeAllCalls() [][]string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	result := make([][]string, len(b.initializeAllCalls))
	for i, call := range b.initializeAllCalls {
		result[i] = append([]string{}, call...)
	}
	return result
}

// MockBridgeManager implements BridgeManager for testing
type MockBridgeManager struct {
	jsBridge        JSBridge
	domBridge       DOMBridge
	componentBridge ComponentBridge
}

func NewMockBridgeManager() *MockBridgeManager {
	return &MockBridgeManager{
		jsBridge:        NewMockJSBridge(),
		domBridge:       NewMockDOMBridge(),
		componentBridge: NewMockComponentBridge(),
	}
}

func (m *MockBridgeManager) JS() JSBridge {
	return m.jsBridge
}

func (m *MockBridgeManager) DOM() DOMBridge {
	return m.domBridge
}

func (m *MockBridgeManager) Component() ComponentBridge {
	return m.componentBridge
}