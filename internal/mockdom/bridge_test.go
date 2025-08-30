//go:build !js || !wasm

package mockdom

import (
	"testing"
	"time"
)

func TestNewMockBridge(t *testing.T) {
	bridge := NewMockBridge()
	if bridge == nil {
		t.Fatal("NewMockBridge returned nil")
	}
	
	if bridge.GetDOM() == nil {
		t.Fatal("MockBridge should have a DOM instance")
	}
	
	// Check that components map is initialized
	bridge.mu.RLock()
	componentsLen := len(bridge.components)
	bridge.mu.RUnlock()
	
	if componentsLen != 0 {
		t.Error("MockBridge should have no components initially")
	}
}

func TestMockBridge_ConvertToJSValue(t *testing.T) {
	bridge := NewMockBridge()
	
	value := bridge.ConvertToJSValue("test string")
	if value == nil {
		t.Fatal("ConvertToJSValue returned nil")
	}
	
	mockValue, ok := value.(*MockJSValue)
	if !ok {
		t.Fatal("ConvertToJSValue should return *MockJSValue")
	}
	
	if mockValue.String() != "test string" {
		t.Errorf("Expected value 'test string', got '%v'", mockValue.String())
	}
}

func TestMockBridge_ConvertStringSliceToJSArray(t *testing.T) {
	bridge := NewMockBridge()
	
	slice := []string{"item1", "item2", "item3"}
	value := bridge.ConvertStringSliceToJSArray(slice)
	
	if value == nil {
		t.Fatal("ConvertStringSliceToJSArray returned nil")
	}
	
	mockValue, ok := value.(*MockJSValue)
	if !ok {
		t.Fatal("ConvertStringSliceToJSArray should return *MockJSValue")
	}
	
	array, ok := mockValue.Underlying().([]interface{})
	if !ok {
		t.Fatal("MockJSValue should contain array")
	}
	
	if len(array) != 3 {
		t.Errorf("Expected array length 3, got %d", len(array))
	}
	
	for i, item := range array {
		if item != slice[i] {
			t.Errorf("Array item %d: expected '%s', got '%v'", i, slice[i], item)
		}
	}
}

func TestMockBridge_ConvertMapToJSObject(t *testing.T) {
	bridge := NewMockBridge()
	
	m := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}
	
	value := bridge.ConvertMapToJSObject(m)
	
	if value == nil {
		t.Fatal("ConvertMapToJSObject returned nil")
	}
	
	mockValue, ok := value.(*MockJSValue)
	if !ok {
		t.Fatal("ConvertMapToJSObject should return *MockJSValue")
	}
	
	obj, ok := mockValue.Underlying().(map[string]interface{})
	if !ok {
		t.Fatal("MockJSValue should contain map")
	}
	
	if obj["key1"] != "value1" {
		t.Errorf("Expected key1 'value1', got '%v'", obj["key1"])
	}
	
	if obj["key2"] != 42 {
		t.Errorf("Expected key2 42, got '%v'", obj["key2"])
	}
	
	if obj["key3"] != true {
		t.Errorf("Expected key3 true, got '%v'", obj["key3"])
	}
}

func TestMockBridge_InitializeFlyonComponents(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create test elements
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	modal := NewMockElement("div", "modal1")
	modal.SetAttribute("data-hs-modal", "")
	bridge.GetDOM().AddElement(modal)
	
	// Test with specific components
	components := []string{"HSDropdown", "HSModal"}
	err := bridge.InitializeFlyonComponents(components)
	
	if err != nil {
		t.Errorf("InitializeFlyonComponents should not return error: %v", err)
	}
	
	// Check if components were initialized
	bridge.mu.RLock()
	componentsLen := len(bridge.components)
	bridge.mu.RUnlock()
	
	if componentsLen != 2 {
		t.Errorf("Expected 2 components to be initialized, got %d", componentsLen)
	}
	
	bridge.mu.RLock()
	_, hsDropdownExists := bridge.components["HSDropdown"]
	_, hsModalExists := bridge.components["HSModal"]
	bridge.mu.RUnlock()
	
	if !hsDropdownExists {
		t.Error("HSDropdown should be in components map")
	}
	
	if !hsModalExists {
		t.Error("HSModal should be in components map")
	}
}

func TestMockBridge_InitializeFlyonComponents_EmptyList(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Test with empty component list (should initialize all)
	err := bridge.InitializeFlyonComponents([]string{})
	
	if err != nil {
		t.Errorf("InitializeFlyonComponents should not return error: %v", err)
	}
	
	// Should have initialized all known components
	expectedComponents := []string{"HSDropdown", "HSModal", "HSCollapse", "HSTabs", "HSAccordion", "HSCarousel", "HSTooltip", "HSPopover"}
	for _, component := range expectedComponents {
		bridge.mu.RLock()
		_, exists := bridge.components[component]
		bridge.mu.RUnlock()
		if !exists {
			t.Errorf("%s should be in components map", component)
		}
	}
}

func TestMockBridge_InitializeDropdown(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create dropdown element
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	// Test with options
	options := map[string]interface{}{
		"autoClose": true,
		"delay": 100,
	}
	
	err := bridge.InitializeDropdown("#dropdown1", options)
	
	if err != nil {
		t.Errorf("InitializeDropdown should not return error: %v", err)
	}
	
	// Check if component instance was set
	instance := dropdown.GetComponentInstance("HSDropdown")
	if instance == nil {
		t.Fatal("Dropdown component instance should be set")
	}
	
	instanceMap, ok := instance.(map[string]interface{})
	if !ok {
		t.Fatal("Component instance should be a map")
	}
	
	if instanceMap["autoClose"] != true {
		t.Error("Component instance should preserve options")
	}
	
	if instanceMap["delay"] != 100 {
		t.Error("Component instance should preserve options")
	}
}

func TestMockBridge_InitializeModal(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create modal element
	modal := NewMockElement("div", "modal1")
	modal.SetAttribute("data-hs-modal", "")
	bridge.GetDOM().AddElement(modal)
	
	err := bridge.InitializeModal("#modal1", nil)
	
	if err != nil {
		t.Errorf("InitializeModal should not return error: %v", err)
	}
	
	// Check if component instance was set
	instance := modal.GetComponentInstance("HSModal")
	if instance == nil {
		t.Fatal("Modal component instance should be set")
	}
}

func TestMockBridge_DestroyComponent(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create and initialize dropdown
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	bridge.InitializeDropdown("#dropdown1", nil)
	
	// Verify component is initialized
	instance := dropdown.GetComponentInstance("HSDropdown")
	if instance == nil {
		t.Fatal("Component should be initialized")
	}
	
	// Destroy component - pass the element, not a selector
	mockDOMElement := &MockDOMElement{element: dropdown}
	err := bridge.DestroyComponent(mockDOMElement, "HSDropdown")
	
	if err != nil {
		t.Errorf("DestroyComponent should not return error: %v", err)
	}
	
	// Verify component is destroyed
	instance = dropdown.GetComponentInstance("HSDropdown")
	if instance != nil {
		t.Error("Component instance should be nil after destruction")
	}
}

func TestMockBridge_GetComponentInstance(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create and initialize dropdown
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	bridge.InitializeDropdown("#dropdown1", nil)
	
	// Get component instance - pass the element, not a selector
	mockDOMElement := &MockDOMElement{element: dropdown}
	instance, err := bridge.GetComponentInstance(mockDOMElement, "HSDropdown")
	
	if err != nil {
		t.Errorf("GetComponentInstance should not return error: %v", err)
	}
	
	if instance == nil {
		t.Fatal("GetComponentInstance should return instance")
	}
	
	_, ok := instance.(*MockJSValue)
	if !ok {
		t.Fatal("GetComponentInstance should return *MockJSValue")
	}
	
	// Test non-existent component
	instance, err = bridge.GetComponentInstance(mockDOMElement, "HSModal")
	if err == nil {
		t.Error("GetComponentInstance should return error for non-existent component")
	}
	if instance != nil {
		t.Error("GetComponentInstance should return nil for non-existent component")
	}
	
	// Test non-existent element
	nonExistent := NewMockElement("div", "non-existent")
	mockNonExistent := &MockDOMElement{element: nonExistent}
	instance, err = bridge.GetComponentInstance(mockNonExistent, "HSDropdown")
	if err == nil {
		t.Error("GetComponentInstance should return error for non-existent element")
	}
	if instance != nil {
		t.Error("GetComponentInstance should return nil for non-existent element")
	}
}

func TestMockBridge_AddEventListener(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create element
	button := NewMockElement("button", "test-button")
	bridge.GetDOM().AddElement(button)
	
	eventFired := false
	callback := func() {
		eventFired = true
	}
	
	// Add event listener - pass the element, not a selector
	mockDOMElement := &MockDOMElement{element: button}
	err := bridge.AddEventListener(mockDOMElement, "click", callback)
	
	if err != nil {
		t.Errorf("AddEventListener should not return error: %v", err)
	}
	
	// Trigger event
	button.Click()
	
	if !eventFired {
		t.Error("Event callback should have been called")
	}
	
	// Check if listener was stored
	bridge.mu.RLock()
	eventListenersLen := len(bridge.eventListeners)
	bridge.mu.RUnlock()
	
	if eventListenersLen != 1 {
		t.Errorf("Expected 1 event listener, got %d", eventListenersLen)
	}
}

func TestMockBridge_RemoveEventListener(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create element
	button := NewMockElement("button", "test-button")
	bridge.GetDOM().AddElement(button)
	
	eventFired := false
	callback := func() {
		eventFired = true
	}
	
	// Add event listener - pass the element, not a selector
	mockDOMElement := &MockDOMElement{element: button}
	bridge.AddEventListener(mockDOMElement, "click", callback)
	
	// Remove event listener
	err := bridge.RemoveEventListener(mockDOMElement, "click", callback)
	
	if err != nil {
		t.Errorf("RemoveEventListener should not return error: %v", err)
	}
	
	// Trigger event - should not fire
	button.Click()
	
	if eventFired {
		t.Error("Event callback should not have been called after removal")
	}
	
	// Check if listener was removed
	bridge.mu.RLock()
	eventListenersLen := len(bridge.eventListeners)
	bridge.mu.RUnlock()
	
	if eventListenersLen != 0 {
		t.Errorf("Expected 0 event listeners, got %d", eventListenersLen)
	}
}

func TestMockBridge_WaitForDOMReady(t *testing.T) {
	bridge := NewMockBridge()
	
	callbackCalled := false
	callback := func() {
		callbackCalled = true
	}
	
	// DOM not ready yet
	bridge.WaitForDOMReady(callback)
	
	if callbackCalled {
		t.Error("Callback should not be called when DOM is not ready")
	}
	
	// Set DOM ready
	bridge.GetDOM().SetReady()
	
	// Give some time for callback to be called
	time.Sleep(10 * time.Millisecond)
	
	if !callbackCalled {
		t.Error("Callback should be called when DOM becomes ready")
	}
}

func TestMockBridge_WaitForDOMReady_AlreadyReady(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	callbackCalled := false
	callback := func() {
		callbackCalled = true
	}
	
	// DOM already ready
	bridge.WaitForDOMReady(callback)
	
	if !callbackCalled {
		t.Error("Callback should be called immediately when DOM is already ready")
	}
}

func TestMockBridge_InitializeAllComponents(t *testing.T) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create various component elements
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	modal := NewMockElement("div", "modal1")
	modal.SetAttribute("data-hs-overlay", "")
	bridge.GetDOM().AddElement(modal)

	tabs := NewMockElement("div", "tabs1")
	tabs.SetAttribute("data-hs-tab", "")
	bridge.GetDOM().AddElement(tabs)
	
	// Initialize all components
	err := bridge.InitializeAllComponents()
	
	if err != nil {
		t.Errorf("InitializeAllComponents should not return error: %v", err)
	}
	
	// Check if all components were initialized by checking the bridge's component storage
	bridge.mu.RLock()
	defer bridge.mu.RUnlock()
	
	if _, exists := bridge.components["HSDropdown"]["dropdown1"]; !exists {
		t.Error("Dropdown should be initialized")
	}
	
	if _, exists := bridge.components["HSModal"]["modal1"]; !exists {
		t.Error("Modal should be initialized")
	}
	
	if _, exists := bridge.components["HSTabs"]["tabs1"]; !exists {
		t.Error("Tabs should be initialized")
	}
}

func TestMockJSValue_Methods(t *testing.T) {
	value := &MockJSValue{
		value: "test",
	}
	
	// Test String method
	if value.String() != "test" {
		t.Errorf("Expected String() to return 'test', got '%s'", value.String())
	}
	
	// Test Bool method
	boolValue := &MockJSValue{
		value: true,
	}
	
	if !boolValue.Bool() {
		t.Error("Expected Bool() to return true")
	}
	
	// Test Int method
	intValue := &MockJSValue{
		value: 42,
	}
	
	if intValue.Int() != 42 {
		t.Errorf("Expected Int() to return 42, got %d", intValue.Int())
	}
	
	// Test Float method
	floatValue := &MockJSValue{
		value: 3.14,
	}
	
	if floatValue.Float() != 3.14 {
		t.Errorf("Expected Float() to return 3.14, got %f", floatValue.Float())
	}
	
	// Test IsNull and IsUndefined
	nullValue := &MockJSValue{
		value: nil,
	}
	
	if !nullValue.IsNull() {
		t.Error("Expected IsNull() to return true for null value")
	}
	
	undefinedValue := &MockJSValue{
		value: nil,
	}
	
	if !undefinedValue.IsUndefined() {
		t.Error("Expected IsUndefined() to return true for undefined value")
	}
	
	// Test Underlying method
	underlying := value.Underlying()
	if underlying != "test" {
		t.Errorf("Expected Underlying() to return 'test', got '%v'", underlying)
	}
}

func TestMockDOMElement_Methods(t *testing.T) {
	bridge := NewMockBridge()
	element := &MockDOMElement{
		element: NewMockElement("div", "test"),
		bridge:  bridge,
	}
	
	// Test TagName
	if element.element.TagName != "DIV" {
		t.Errorf("Expected TagName 'DIV', got '%s'", element.element.TagName)
	}
	
	// Test ID
	if element.GetID() != "test" {
		t.Errorf("Expected ID 'test', got '%s'", element.GetID())
	}
	
	// Test SetID
	element.SetID("new-id")
	if element.GetID() != "new-id" {
		t.Errorf("Expected ID 'new-id' after SetID, got '%s'", element.GetID())
	}
	
	// Test ClassName
	if element.GetClassName() != "" {
		t.Errorf("Expected empty ClassName, got '%s'", element.GetClassName())
	}
	
	// Test AddClass
	element.AddClass("test-class")
	if !element.HasClass("test-class") {
		t.Error("Element should have class 'test-class' after AddClass")
	}
	
	// Test RemoveClass
	element.RemoveClass("test-class")
	if element.HasClass("test-class") {
		t.Error("Element should not have class 'test-class' after RemoveClass")
	}
	
	// Test attribute methods
	element.SetAttribute("data-test", "value")
	if element.GetAttribute("data-test") != "value" {
		t.Errorf("Expected GetAttribute() to return 'value', got '%s'", element.GetAttribute("data-test"))
	}
	
	if !element.HasAttribute("data-test") {
		t.Error("Expected HasAttribute() to return true")
	}
	
	element.RemoveAttribute("data-test")
	if element.HasAttribute("data-test") {
		t.Error("Expected HasAttribute() to return false after removal")
	}
	
	// Test Underlying method
	underlying := element.Underlying()
	if underlying == nil {
		t.Error("Expected Underlying() to return non-nil value")
	}
}

// Integration test
func TestMockBridge_Integration(t *testing.T) {
	bridge := NewMockBridge()
	
	// Create a complex component setup
	container := NewMockElement("div", "container")
	container.AddClass("container")
	bridge.GetDOM().AddElement(container)
	
	// Create dropdown
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	dropdown.AddClass("dropdown")
	container.AppendChild(dropdown)
	bridge.GetDOM().AddElement(dropdown)
	
	// Create modal
	modal := NewMockElement("div", "modal1")
	modal.SetAttribute("data-hs-overlay", "")
	modal.AddClass("modal")
	container.AppendChild(modal)
	bridge.GetDOM().AddElement(modal)
	
	// Set DOM ready
	bridge.GetDOM().SetReady()
	
	// Initialize all components
	err := bridge.InitializeAllComponents()
	if err != nil {
		t.Fatalf("Failed to initialize components: %v", err)
	}
	
	// Verify components are initialized
	if dropdown.GetComponentInstance("HSDropdown") == nil {
		t.Error("Dropdown should be initialized")
	}
	
	if modal.GetComponentInstance("HSModal") == nil {
		t.Error("Modal should be initialized")
	}
	
	// Add event listeners
	dropdownClicked := false
	modalClicked := false
	
	dropdownDOMElement := &MockDOMElement{element: dropdown}
	modalDOMElement := &MockDOMElement{element: modal}
	
	bridge.AddEventListener(dropdownDOMElement, "click", func() {
		dropdownClicked = true
	})
	
	bridge.AddEventListener(modalDOMElement, "click", func() {
		modalClicked = true
	})
	
	// Trigger events
	dropdown.Click()
	modal.Click()
	
	if !dropdownClicked {
		t.Error("Dropdown click event should have fired")
	}
	
	if !modalClicked {
		t.Error("Modal click event should have fired")
	}
	
	// Test component destruction
	bridge.DestroyComponent(dropdownDOMElement, "HSDropdown")
	if dropdown.GetComponentInstance("HSDropdown") != nil {
		t.Error("Dropdown component should be destroyed")
	}
	
	// Modal should still be intact
	if modal.GetComponentInstance("HSModal") == nil {
		t.Error("Modal component should still be initialized")
	}
	
	// Test DOM queries
	modals := bridge.GetDOM().QuerySelectorAll(".modal")
	if len(modals) != 1 {
		t.Errorf("Expected 1 modal element, got %d", len(modals))
	}
	
	dropdowns := bridge.GetDOM().QuerySelectorAll(".dropdown")
	if len(dropdowns) != 1 {
		t.Errorf("Expected 1 dropdown element, got %d", len(dropdowns))
	}
	
	// Test component instance retrieval
	modalInstance, _ := bridge.GetComponentInstance(modalDOMElement, "HSModal")
	if modalInstance == nil {
		t.Error("Should be able to retrieve modal instance")
	}
	
	droppedInstance, _ := bridge.GetComponentInstance(dropdownDOMElement, "HSDropdown")
	if droppedInstance != nil {
		t.Error("Should not be able to retrieve destroyed dropdown instance")
	}
}

// Benchmark tests
func BenchmarkMockBridge_InitializeComponent(b *testing.B) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create element
	dropdown := NewMockElement("div", "dropdown1")
	dropdown.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(dropdown)
	
	mockDOMElement := &MockDOMElement{element: dropdown}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bridge.InitializeDropdown("#dropdown1", nil)
		bridge.DestroyComponent(mockDOMElement, "HSDropdown")
	}
}

func BenchmarkMockBridge_EventHandling(b *testing.B) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create element
	button := NewMockElement("button", "test-button")
	bridge.GetDOM().AddElement(button)
	
	// Add event listener
	mockDOMElement := &MockDOMElement{element: button}
	bridge.AddEventListener(mockDOMElement, "click", func() {})
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		button.Click()
	}
}

func BenchmarkMockBridge_ComponentQuery(b *testing.B) {
	bridge := NewMockBridge()
	bridge.GetDOM().SetReady()
	
	// Create many elements
	for i := 0; i < 100; i++ {
		element := NewMockElement("div", "")
		element.SetAttribute("data-hs-dropdown", "")
		bridge.GetDOM().AddElement(element)
	}
	
	// Create target element
	target := NewMockElement("div", "target")
	target.SetAttribute("data-hs-dropdown", "")
	bridge.GetDOM().AddElement(target)
	
	b.ResetTimer()
	mockDOMElement := &MockDOMElement{element: target}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bridge.GetComponentInstance(mockDOMElement, "HSDropdown")
	}
}