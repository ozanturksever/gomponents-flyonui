//go:build !js || !wasm

package mockdom

import (
	"testing"
)

func TestNewMockDOM(t *testing.T) {
	dom := NewMockDOM()
	if dom == nil {
		t.Fatal("NewMockDOM returned nil")
	}
	
	if dom.IsReady() {
		t.Error("DOM should not be ready initially")
	}
	
	if len(dom.elements) != 0 {
		t.Error("DOM should have no elements initially")
	}
}

func TestNewMockElement(t *testing.T) {
	element := NewMockElement("div", "test-id")
	if element == nil {
		t.Fatal("NewMockElement returned nil")
	}
	
	if element.TagName != "DIV" {
		t.Errorf("Expected TagName 'DIV', got '%s'", element.TagName)
	}
	
	if element.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got '%s'", element.ID)
	}
	
	if element.ClassName != "" {
		t.Errorf("Expected empty ClassName, got '%s'", element.ClassName)
	}
	
	if len(element.Attributes) != 0 {
		t.Error("Element should have no attributes initially")
	}
	
	if len(element.Children) != 0 {
		t.Error("Element should have no children initially")
	}
}

func TestMockDOM_SetReady(t *testing.T) {
	dom := NewMockDOM()
	callbackCalled := false
	
	// Add callback before setting ready
	dom.OnReady(func() {
		callbackCalled = true
	})
	
	if dom.IsReady() {
		t.Error("DOM should not be ready initially")
	}
	
	dom.SetReady()
	
	if !dom.IsReady() {
		t.Error("DOM should be ready after SetReady()")
	}
	
	if !callbackCalled {
		t.Error("Ready callback should have been called")
	}
}

func TestMockDOM_OnReady_WhenAlreadyReady(t *testing.T) {
	dom := NewMockDOM()
	dom.SetReady()
	
	callbackCalled := false
	dom.OnReady(func() {
		callbackCalled = true
	})
	
	if !callbackCalled {
		t.Error("Callback should be called immediately when DOM is already ready")
	}
}

func TestMockDOM_CreateElement(t *testing.T) {
	dom := NewMockDOM()
	element := dom.CreateElement("span")
	
	if element == nil {
		t.Fatal("CreateElement returned nil")
	}
	
	if element.TagName != "SPAN" {
		t.Errorf("Expected TagName 'SPAN', got '%s'", element.TagName)
	}
	
	if element.ID != "" {
		t.Error("Created element should have empty ID")
	}
}

func TestMockDOM_AddElement_GetElementByID(t *testing.T) {
	dom := NewMockDOM()
	element := NewMockElement("div", "test-element")
	
	dom.AddElement(element)
	
	retrieved := dom.GetElementByID("test-element")
	if retrieved == nil {
		t.Fatal("GetElementByID returned nil")
	}
	
	if retrieved != element {
		t.Error("Retrieved element is not the same as added element")
	}
	
	// Test non-existent element
	nonExistent := dom.GetElementByID("non-existent")
	if nonExistent != nil {
		t.Error("GetElementByID should return nil for non-existent element")
	}
}

func TestMockDOM_QuerySelector(t *testing.T) {
	dom := NewMockDOM()
	
	// Create elements
	divElement := NewMockElement("div", "test-div")
	spanElement := NewMockElement("span", "test-span")
	classElement := NewMockElement("p", "test-p")
	classElement.AddClass("test-class")
	
	dom.AddElement(divElement)
	dom.AddElement(spanElement)
	dom.AddElement(classElement)
	
	// Test ID selector
	result := dom.QuerySelector("#test-div")
	if result == nil {
		t.Fatal("QuerySelector with ID selector returned nil")
	}
	if result.ID != "test-div" {
		t.Errorf("Expected ID 'test-div', got '%s'", result.ID)
	}
	
	// Test class selector
	result = dom.QuerySelector(".test-class")
	if result == nil {
		t.Fatal("QuerySelector with class selector returned nil")
	}
	if !result.HasClass("test-class") {
		t.Error("Element should have 'test-class' class")
	}
	
	// Test tag selector
	result = dom.QuerySelector("span")
	if result == nil {
		t.Fatal("QuerySelector with tag selector returned nil")
	}
	if result.TagName != "SPAN" {
		t.Errorf("Expected TagName 'SPAN', got '%s'", result.TagName)
	}
	
	// Test non-existent selector
	result = dom.QuerySelector("#non-existent")
	if result != nil {
		t.Error("QuerySelector should return nil for non-existent selector")
	}
}

func TestMockDOM_QuerySelectorAll(t *testing.T) {
	dom := NewMockDOM()
	
	// Create multiple elements with same class
	for i := 0; i < 3; i++ {
		element := NewMockElement("div", "")
		element.AddClass("test-class")
		dom.AddElement(element)
	}
	
	// Create element with different class
	differentElement := NewMockElement("div", "")
	differentElement.AddClass("different-class")
	dom.AddElement(differentElement)
	
	// Test class selector
	results := dom.QuerySelectorAll(".test-class")
	if len(results) != 3 {
		t.Errorf("Expected 3 elements with 'test-class', got %d", len(results))
	}
	
	for _, element := range results {
		if !element.HasClass("test-class") {
			t.Error("All returned elements should have 'test-class' class")
		}
	}
	
	// Test tag selector
	results = dom.QuerySelectorAll("div")
	if len(results) != 4 {
		t.Errorf("Expected 4 div elements, got %d", len(results))
	}
	
	// Test non-existent selector
	results = dom.QuerySelectorAll(".non-existent")
	if len(results) != 0 {
		t.Errorf("Expected 0 elements for non-existent class, got %d", len(results))
	}
}

func TestMockDOM_EventHandling(t *testing.T) {
	dom := NewMockDOM()
	eventFired := false
	
	// Add global event listener
	dom.AddEventListener("test-event", func(event *MockEvent) {
		eventFired = true
		if event.Type != "test-event" {
			t.Errorf("Expected event type 'test-event', got '%s'", event.Type)
		}
	})
	
	// Dispatch event
	event := NewMockEvent("test-event", nil)
	dom.DispatchEvent(event)
	
	if !eventFired {
		t.Error("Event should have been fired")
	}
}

func TestMockElement_Attributes(t *testing.T) {
	element := NewMockElement("div", "test-id")
	
	// Test SetAttribute and GetAttribute
	element.SetAttribute("data-test", "test-value")
	value := element.GetAttribute("data-test")
	if value != "test-value" {
		t.Errorf("Expected attribute value 'test-value', got '%s'", value)
	}
	
	// Test HasAttribute
	if !element.HasAttribute("data-test") {
		t.Error("Element should have 'data-test' attribute")
	}
	
	if element.HasAttribute("non-existent") {
		t.Error("Element should not have 'non-existent' attribute")
	}
	
	// Test RemoveAttribute
	element.RemoveAttribute("data-test")
	if element.HasAttribute("data-test") {
		t.Error("Element should not have 'data-test' attribute after removal")
	}
	
	value = element.GetAttribute("data-test")
	if value != "" {
		t.Errorf("Expected empty attribute value after removal, got '%s'", value)
	}
}

func TestMockElement_Classes(t *testing.T) {
	element := NewMockElement("div", "test-id")
	
	// Test AddClass
	element.AddClass("class1")
	if !element.HasClass("class1") {
		t.Error("Element should have 'class1' class")
	}
	
	if element.ClassName != "class1" {
		t.Errorf("Expected ClassName 'class1', got '%s'", element.ClassName)
	}
	
	// Test adding multiple classes
	element.AddClass("class2")
	if !element.HasClass("class1") || !element.HasClass("class2") {
		t.Error("Element should have both 'class1' and 'class2' classes")
	}
	
	// Test adding duplicate class
	element.AddClass("class1")
	if element.ClassName != "class1 class2" {
		t.Errorf("Expected ClassName 'class1 class2', got '%s'", element.ClassName)
	}
	
	// Test RemoveClass
	element.RemoveClass("class1")
	if element.HasClass("class1") {
		t.Error("Element should not have 'class1' class after removal")
	}
	if !element.HasClass("class2") {
		t.Error("Element should still have 'class2' class")
	}
	
	// Test ToggleClass
	element.ToggleClass("class3")
	if !element.HasClass("class3") {
		t.Error("Element should have 'class3' class after toggle")
	}
	
	element.ToggleClass("class3")
	if element.HasClass("class3") {
		t.Error("Element should not have 'class3' class after second toggle")
	}
}

func TestMockElement_Children(t *testing.T) {
	parent := NewMockElement("div", "parent")
	child1 := NewMockElement("span", "child1")
	child2 := NewMockElement("p", "child2")
	
	// Test AppendChild
	parent.AppendChild(child1)
	parent.AppendChild(child2)
	
	if len(parent.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(parent.Children))
	}
	
	if child1.Parent != parent {
		t.Error("child1's parent should be set to parent")
	}
	
	if child2.Parent != parent {
		t.Error("child2's parent should be set to parent")
	}
	
	// Test RemoveChild
	err := parent.RemoveChild(child1)
	if err != nil {
		t.Errorf("RemoveChild should not return error: %v", err)
	}
	
	if len(parent.Children) != 1 {
		t.Errorf("Expected 1 child after removal, got %d", len(parent.Children))
	}
	
	if child1.Parent != nil {
		t.Error("child1's parent should be nil after removal")
	}
	
	// Test removing non-existent child
	nonChild := NewMockElement("div", "non-child")
	err = parent.RemoveChild(nonChild)
	if err == nil {
		t.Error("RemoveChild should return error for non-existent child")
	}
}

func TestMockElement_QuerySelector(t *testing.T) {
	parent := NewMockElement("div", "parent")
	child1 := NewMockElement("span", "child1")
	child2 := NewMockElement("p", "child2")
	grandchild := NewMockElement("a", "grandchild")
	
	child1.AddClass("test-class")
	child2.AppendChild(grandchild)
	parent.AppendChild(child1)
	parent.AppendChild(child2)
	
	// Test ID selector
	result := parent.QuerySelector("#child1")
	if result == nil {
		t.Fatal("QuerySelector should find child1")
	}
	if result.ID != "child1" {
		t.Errorf("Expected ID 'child1', got '%s'", result.ID)
	}
	
	// Test class selector
	result = parent.QuerySelector(".test-class")
	if result == nil {
		t.Fatal("QuerySelector should find element with test-class")
	}
	if !result.HasClass("test-class") {
		t.Error("Found element should have test-class")
	}
	
	// Test tag selector (should find first match)
	result = parent.QuerySelector("span")
	if result == nil {
		t.Fatal("QuerySelector should find span element")
	}
	if result.TagName != "SPAN" {
		t.Errorf("Expected TagName 'SPAN', got '%s'", result.TagName)
	}
	
	// Test deep selector
	result = parent.QuerySelector("a")
	if result == nil {
		t.Fatal("QuerySelector should find grandchild element")
	}
	if result.ID != "grandchild" {
		t.Errorf("Expected ID 'grandchild', got '%s'", result.ID)
	}
}

func TestMockElement_EventHandling(t *testing.T) {
	element := NewMockElement("button", "test-button")
	eventFired := false
	
	// Add event listener
	element.AddEventListener("click", func(event *MockEvent) {
		eventFired = true
		if event.Target != element {
			t.Error("Event target should be the element")
		}
		if event.Type != "click" {
			t.Errorf("Expected event type 'click', got '%s'", event.Type)
		}
	})
	
	// Test Click method
	element.Click()
	if !eventFired {
		t.Error("Click event should have been fired")
	}
	
	// Test Focus and Blur
	focusFired := false
	blurFired := false
	
	element.AddEventListener("focus", func(event *MockEvent) {
		focusFired = true
	})
	
	element.AddEventListener("blur", func(event *MockEvent) {
		blurFired = true
	})
	
	element.Focus()
	if !focusFired {
		t.Error("Focus event should have been fired")
	}
	
	element.Blur()
	if !blurFired {
		t.Error("Blur event should have been fired")
	}
}

func TestMockElement_EventBubbling(t *testing.T) {
	parent := NewMockElement("div", "parent")
	child := NewMockElement("button", "child")
	parent.AppendChild(child)
	
	parentEventFired := false
	childEventFired := false
	
	// Add event listeners
	parent.AddEventListener("click", func(event *MockEvent) {
		parentEventFired = true
		if event.Target != child {
			t.Error("Event target should be the child element")
		}
		if event.CurrentTarget != parent {
			t.Error("Event currentTarget should be the parent element")
		}
	})
	
	child.AddEventListener("click", func(event *MockEvent) {
		childEventFired = true
		if event.Target != child {
			t.Error("Event target should be the child element")
		}
		if event.CurrentTarget != child {
			t.Error("Event currentTarget should be the child element")
		}
	})
	
	// Click on child should bubble to parent
	child.Click()
	
	if !childEventFired {
		t.Error("Child event should have been fired")
	}
	
	if !parentEventFired {
		t.Error("Parent event should have been fired due to bubbling")
	}
}

func TestMockElement_ComponentInstance(t *testing.T) {
	element := NewMockElement("div", "test-element")
	
	// Test setting and getting component instance
	instance := map[string]interface{}{"initialized": true}
	element.SetComponentInstance("HSDropdown", instance)
	
	retrieved := element.GetComponentInstance("HSDropdown")
	if retrieved == nil {
		t.Fatal("GetComponentInstance returned nil")
	}
	
	if retrievedMap, ok := retrieved.(map[string]interface{}); ok {
		if retrievedMap["initialized"] != true {
			t.Error("Component instance data should be preserved")
		}
	} else {
		t.Error("Retrieved instance should be a map")
	}
	
	// Test getting non-existent component
	nonExistent := element.GetComponentInstance("HSModal")
	if nonExistent != nil {
		t.Error("GetComponentInstance should return nil for non-existent component")
	}
}

func TestMockElement_String(t *testing.T) {
	element := NewMockElement("div", "test-id")
	element.SetAttribute("class", "test-class")
	element.SetAttribute("data-test", "value")
	
	str := element.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	
	// Should contain tag name
	if !contains(str, "div") {
		t.Error("String representation should contain tag name")
	}
}

func TestNewMockEvent(t *testing.T) {
	element := NewMockElement("div", "test-element")
	event := NewMockEvent("click", element)
	
	if event.Type != "click" {
		t.Errorf("Expected event type 'click', got '%s'", event.Type)
	}
	
	if event.Target != element {
		t.Error("Event target should be the provided element")
	}
	
	if !event.Bubbles {
		t.Error("Event should bubble by default")
	}
	
	if !event.Cancelable {
		t.Error("Event should be cancelable by default")
	}
	
	if event.DefaultPrevented {
		t.Error("Event should not be default prevented initially")
	}
}

func TestMockEvent_PreventDefault(t *testing.T) {
	event := NewMockEvent("click", nil)
	
	if event.DefaultPrevented {
		t.Error("Event should not be default prevented initially")
	}
	
	event.PreventDefault()
	
	if !event.DefaultPrevented {
		t.Error("Event should be default prevented after calling PreventDefault")
	}
}

func TestMockEvent_StopPropagation(t *testing.T) {
	event := NewMockEvent("click", nil)
	
	if !event.Bubbles {
		t.Error("Event should bubble initially")
	}
	
	event.StopPropagation()
	
	if event.Bubbles {
		t.Error("Event should not bubble after calling StopPropagation")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			 s[len(s)-len(substr):] == substr || 
			 indexOfSubstring(s, substr) >= 0)))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// Benchmark tests
func BenchmarkMockDOM_QuerySelector(b *testing.B) {
	dom := NewMockDOM()
	
	// Create many elements
	for i := 0; i < 1000; i++ {
		element := NewMockElement("div", "")
		element.AddClass("test-class")
		dom.AddElement(element)
	}
	
	// Add target element
	target := NewMockElement("div", "target")
	dom.AddElement(target)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dom.QuerySelector("#target")
	}
}

func BenchmarkMockElement_AddClass(b *testing.B) {
	element := NewMockElement("div", "test")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		element.AddClass("test-class")
		element.RemoveClass("test-class")
	}
}

func BenchmarkMockElement_EventDispatch(b *testing.B) {
	element := NewMockElement("button", "test")
	
	// Add event listener
	element.AddEventListener("click", func(event *MockEvent) {
		// Do nothing
	})
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		element.Click()
	}
}

// Integration test
func TestMockDOM_Integration(t *testing.T) {
	dom := NewMockDOM()
	
	// Create a complex DOM structure
	container := NewMockElement("div", "container")
	container.AddClass("container")
	dom.AddElement(container)
	
	header := NewMockElement("header", "header")
	header.AddClass("header")
	container.AppendChild(header)
	dom.AddElement(header) // Add to DOM for querying
	
	nav := NewMockElement("nav", "nav")
	nav.AddClass("navigation")
	header.AppendChild(nav)
	dom.AddElement(nav) // Add to DOM for querying
	
	main := NewMockElement("main", "main")
	main.AddClass("content")
	container.AppendChild(main)
	dom.AddElement(main) // Add to DOM for querying
	
	// Create multiple buttons
	for i := 0; i < 3; i++ {
		button := NewMockElement("button", "")
		button.AddClass("btn")
		button.SetAttribute("data-action", "click")
		main.AppendChild(button)
		dom.AddElement(button) // Add to DOM for querying
	}
	
	dom.SetReady()
	
	// Test complex queries
	buttons := dom.QuerySelectorAll(".btn")
	if len(buttons) != 3 {
		t.Errorf("Expected 3 buttons, got %d", len(buttons))
	}
	
	headerEl := dom.QuerySelector("#header")
	if headerEl == nil {
		t.Fatal("Should find header element")
	}
	
	navEl := headerEl.QuerySelector("nav")
	if navEl == nil {
		t.Fatal("Should find nav element within header")
	}
	
	// Test event handling with bubbling
	clickCount := 0
	container.AddEventListener("click", func(event *MockEvent) {
		clickCount++
	})
	
	// Click on each button - should bubble to container
	for _, button := range buttons {
		button.Click()
	}
	
	if clickCount != 3 {
		t.Errorf("Expected 3 click events to bubble to container, got %d", clickCount)
	}
	
	// Test DOM ready functionality
	readyCallbackCount := 0
	dom.OnReady(func() {
		readyCallbackCount++
	})
	
	// Should be called immediately since DOM is already ready
	if readyCallbackCount != 1 {
		t.Errorf("Expected ready callback to be called immediately, count: %d", readyCallbackCount)
	}
}