package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
)

func TestNewDropdown(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	item1 := DropdownItem(gomponents.Text("Item 1"))
	item2 := DropdownItem(gomponents.Text("Item 2"))
	
	dropdown := NewDropdown(trigger, item1, item2)
	
	if dropdown == nil {
		t.Fatal("NewDropdown returned nil")
	}
	
	foundDropdown := false
	for _, c := range dropdown.classes {
		if c == "dropdown" {
			foundDropdown = true
			break
		}
	}
	if !foundDropdown {
		t.Errorf("Expected classes to include 'dropdown', got %v", dropdown.classes)
	}
	
	if dropdown.position != DropdownBottom {
		t.Errorf("Expected default position to be DropdownBottom, got %v", dropdown.position)
	}
	
	if !dropdown.autoClose {
		t.Error("Expected autoClose to be true by default")
	}
	
	if dropdown.disabled {
		t.Error("Expected disabled to be false by default")
	}
}

func TestDropdownPosition_String(t *testing.T) {
	tests := []struct {
		position DropdownPosition
		expected string
	}{
		{DropdownBottom, "dropdown-bottom"},
		{DropdownTop, "dropdown-top"},
		{DropdownLeft, "dropdown-left"},
		{DropdownRight, "dropdown-right"},
		{DropdownBottomStart, "dropdown-bottom dropdown-start"},
		{DropdownBottomEnd, "dropdown-bottom dropdown-end"},
		{DropdownTopStart, "dropdown-top dropdown-start"},
		{DropdownTopEnd, "dropdown-top dropdown-end"},
	}
	
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if got := test.position.String(); got != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, got)
			}
		})
	}
}

func TestDropdownComponent_WithID(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger)
	
	newDropdown := dropdown.WithID("custom-dropdown")
	
	if newDropdown.id != "custom-dropdown" {
		t.Errorf("Expected ID to be 'custom-dropdown', got %s", newDropdown.id)
	}
	
	// Ensure original is unchanged
	if dropdown.id != "" {
		t.Error("Original dropdown ID should remain empty")
	}
}

func TestDropdownComponent_WithPosition(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger)
	
	newDropdown := dropdown.WithPosition(DropdownTop)
	
	if newDropdown.position != DropdownTop {
		t.Errorf("Expected position to be DropdownTop, got %v", newDropdown.position)
	}
	
	// Ensure original is unchanged
	if dropdown.position != DropdownBottom {
		t.Error("Original dropdown position should remain DropdownBottom")
	}
}

func TestDropdownComponent_WithAutoClose(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger)
	
	newDropdown := dropdown.WithAutoClose(false)
	
	if newDropdown.autoClose {
		t.Error("Expected autoClose to be false")
	}
	
	// Ensure original is unchanged
	if !dropdown.autoClose {
		t.Error("Original dropdown autoClose should remain true")
	}
}

func TestDropdownComponent_WithDisabled(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger)
	
	newDropdown := dropdown.WithDisabled(true)
	
	if !newDropdown.disabled {
		t.Error("Expected disabled to be true")
	}
	
	// Ensure original is unchanged
	if dropdown.disabled {
		t.Error("Original dropdown disabled should remain false")
	}
}

func TestDropdownComponent_With(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger)
	
	// Test with Size modifier
	newDropdown := dropdown.With(flyon.SizeLarge)
	component, ok := newDropdown.(*DropdownComponent)
	if !ok {
		t.Fatal("With() should return *DropdownComponent")
	}
	
	found := false
	for _, class := range component.classes {
		if class == "dropdown-lg" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'dropdown-lg' class to be added")
	}
	
	// Test with DropdownPosition modifier
	newDropdown2 := dropdown.With(DropdownTopStart)
	component2, ok := newDropdown2.(*DropdownComponent)
	if !ok {
		t.Fatal("With() should return *DropdownComponent")
	}
	
	if component2.position != DropdownTopStart {
		t.Errorf("Expected position to be DropdownTopStart, got %v", component2.position)
	}
	
	// Test with custom string class
	newDropdown3 := dropdown.With("custom-class")
	component3, ok := newDropdown3.(*DropdownComponent)
	if !ok {
		t.Fatal("With() should return *DropdownComponent")
	}
	
	found = false
	for _, class := range component3.classes {
		if class == "custom-class" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'custom-class' to be added")
	}
}

func TestDropdownComponent_Render(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	item1 := DropdownItem(gomponents.Text("Item 1"))
	item2 := DropdownItem(gomponents.Text("Item 2"))
	
	dropdown := NewDropdown(trigger, item1, item2).WithID("test-dropdown")
	
	var buf strings.Builder
	err := dropdown.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	// Check for essential dropdown structure
	if !strings.Contains(html, `id="test-dropdown"`) {
		t.Error("Expected dropdown ID in HTML")
	}
	
	if !strings.Contains(html, `class="dropdown`) {
		t.Error("Expected dropdown class in HTML")
	}
	
	if !strings.Contains(html, `class="dropdown-toggle"`) {
		t.Error("Expected dropdown-toggle class on trigger")
	}
	
	if !strings.Contains(html, `class="dropdown-menu`) {
		t.Error("Expected dropdown-menu class on menu")
	}
	
	if !strings.Contains(html, `aria-expanded="false"`) {
		t.Error("Expected aria-expanded attribute on trigger")
	}
	
	if !strings.Contains(html, `aria-haspopup="menu"`) {
		t.Error("Expected aria-haspopup=menu attribute on trigger")
	}
	
	if !strings.Contains(html, `role="menu"`) {
		t.Error("Expected role=menu on dropdown menu")
	}
	
	if !strings.Contains(html, "Open Menu") {
		t.Error("Expected trigger text in HTML")
	}
	
	if !strings.Contains(html, "Item 1") {
		t.Error("Expected first item text in HTML")
	}
	
	if !strings.Contains(html, "Item 2") {
		t.Error("Expected second item text in HTML")
	}
}

func TestDropdownComponent_RenderWithPosition(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger).WithPosition(DropdownTopStart)
	
	var buf strings.Builder
	err := dropdown.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "dropdown-top") {
		t.Error("Expected dropdown-top class in HTML")
	}
	
	if !strings.Contains(html, "dropdown-start") {
		t.Error("Expected dropdown-start class in HTML")
	}
}

func TestDropdownComponent_RenderDisabled(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	dropdown := NewDropdown(trigger).WithDisabled(true)
	
	var buf strings.Builder
	err := dropdown.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "disabled") {
		t.Error("Expected disabled attribute in HTML")
	}
}

func TestDropdownItem(t *testing.T) {
	item := DropdownItem(gomponents.Text("Test Item"))
	
	var buf strings.Builder
	err := item.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "<li>") {
		t.Error("Expected li element")
	}
	
	if !strings.Contains(html, "<a") {
		t.Error("Expected anchor element")
	}
	
	if !strings.Contains(html, `role="menuitem"`) {
		t.Error("Expected role=menuitem")
	}
	
	if !strings.Contains(html, "Test Item") {
		t.Error("Expected item text")
	}
	
	if !strings.Contains(html, "dropdown-item") {
		t.Error("Expected dropdown-item class")
	}
}

func TestDropdownDivider(t *testing.T) {
	divider := DropdownDivider()
	
	var buf strings.Builder
	err := divider.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "<li>") {
		t.Error("Expected li element")
	}
	
	if !strings.Contains(html, "<hr") {
		t.Error("Expected hr element")
	}
	
	if !strings.Contains(html, "my-1") {
		t.Error("Expected margin styling class")
	}
}

func TestDropdownHeader(t *testing.T) {
	header := DropdownHeader("Section Title")
	
	var buf strings.Builder
	err := header.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "<li>") {
		t.Error("Expected li element")
	}
	
	if !strings.Contains(html, "<div") {
		t.Error("Expected div element")
	}
	
	if !strings.Contains(html, "Section Title") {
		t.Error("Expected header text")
	}
	
	if !strings.Contains(html, "font-semibold") {
		t.Error("Expected font styling class")
	}
	
	if !strings.Contains(html, "text-base-content/70") {
		t.Error("Expected text color class")
	}
}

func TestDropdownComponent_Immutability(t *testing.T) {
	trigger := gomponents.Text("Open Menu")
	original := NewDropdown(trigger)
	
	// Test that modifications create new instances
	modified1 := original.WithID("test-id")
	modified2 := original.WithPosition(DropdownTop)
	modified3 := original.WithAutoClose(false)
	modified4 := original.WithDisabled(true)
	modified5 := original.With(flyon.SizeLarge)
	
	// Ensure original is unchanged
	if original.id != "" {
		t.Error("Original ID should remain empty")
	}
	if original.position != DropdownBottom {
		t.Error("Original position should remain DropdownBottom")
	}
	if !original.autoClose {
		t.Error("Original autoClose should remain true")
	}
	if original.disabled {
		t.Error("Original disabled should remain false")
	}
	// Default classes should still be the defaults from NewDropdown
	expected := map[string]bool{"dropdown": false, "relative inline-flex": false}
	for _, c := range original.classes {
		if _, ok := expected[c]; ok {
			expected[c] = true
		}
	}
	for k, v := range expected {
		if !v {
			t.Errorf("Original classes should include %q", k)
		}
	}
	
	// Ensure modifications are applied to new instances
	if modified1.id != "test-id" {
		t.Error("Modified1 should have new ID")
	}
	if modified2.position != DropdownTop {
		t.Error("Modified2 should have new position")
	}
	if modified3.autoClose {
		t.Error("Modified3 should have autoClose false")
	}
	if !modified4.disabled {
		t.Error("Modified4 should be disabled")
	}
	
	// Test With method returns flyon.Component
	modified5Dropdown, ok := modified5.(*DropdownComponent)
	if !ok {
		t.Error("With method should return *DropdownComponent")
	}
	
	// Check that modified5 has the size class
	found := false
	for _, class := range modified5Dropdown.classes {
		if class == "dropdown-lg" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Modified5 should have size class")
	}
}