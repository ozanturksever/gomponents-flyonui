package bridge

// JSBridge defines the interface for JavaScript interop operations
type JSBridge interface {
	// Global JS operations
	GetGlobal(name string) JSValue
	CreateArray(length int) JSValue
	CreateObject() JSValue
	
	// Type conversion
	ValueOf(v interface{}) JSValue
}

// JSValue wraps js.Value with a testable interface
type JSValue interface {
	IsUndefined() bool
	String() string
	Int() int
	Bool() bool
	Get(key string) JSValue
	Set(key string, value interface{})
	Index(i int) JSValue
	SetIndex(i int, value interface{})
	Invoke(args ...interface{}) JSValue
	New(args ...interface{}) JSValue
	Call(method string, args ...interface{}) JSValue
	Underlying() interface{}
}

// DOMBridge defines the interface for DOM operations
type DOMBridge interface {
	QuerySelectorAll(selector string) []DOMElement
	GetElementByID(id string) DOMElement
	GetDocument() DOMDocument
	AddEventListener(selector, eventType string, handler func(DOMEvent)) error
	RemoveEventListener(selector, eventType string) error
	WaitForReady(callback func())
}

// DOMElement represents a DOM element
type DOMElement interface {
	GetAttribute(name string) string
	SetAttribute(name, value string)
	AddClass(class string)
	RemoveClass(class string)
	HasClass(class string) bool
	GetParent() DOMElement
	QuerySelector(selector string) DOMElement
	Closest(selector string) DOMElement
	Underlying() interface{}
}

// DOMDocument represents the document
type DOMDocument interface {
	QuerySelectorAll(selector string) []DOMElement
	GetElementByID(id string) DOMElement
	GetReadyState() string
	AddEventListener(eventType string, handler func(DOMEvent))
}

// DOMEvent represents a DOM event
type DOMEvent interface {
	PreventDefault()
	StopPropagation()
	GetTarget() DOMElement
	GetType() string
}

// ComponentBridge defines the interface for component operations
type ComponentBridge interface {
	InitializeComponent(componentName, selector string, options map[string]interface{}) error
	DestroyComponent(selector, componentType string) error
	GetComponentInstance(selector, componentType string) (JSValue, error)
	InitializeAll(components []string) error
}

// BridgeManager combines all bridge interfaces
type BridgeManager interface {
	JS() JSBridge
	DOM() DOMBridge
	Component() ComponentBridge
	
	// Utility methods
	GoStringsToJSArray(goSlice []string) JSValue
	GoMapToJSObject(goMap map[string]interface{}) JSValue
	
	// Component initialization methods
	InitializeFlyonComponents(components []string) error
	InitializeSpecificComponent(componentName, selector string, options map[string]interface{}) error
	InitializeAllComponents() error
	
	// Component management methods
	DestroyComponent(selector, componentType string) error
	GetComponentInstance(selector, componentType string) (JSValue, error)
	
	// Event handling methods
	AddEventListener(selector, eventType string, handler func(DOMEvent)) error
	RemoveEventListener(selector, eventType string, handler func(DOMEvent)) error
	
	// DOM ready method
	WaitForDOMReady(callback func())
}