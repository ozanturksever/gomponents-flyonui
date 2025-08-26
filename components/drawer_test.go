package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
)

// renderToString is a helper function to render a component to string for testing
func renderToStringDrawer(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

// renderNodeToString is a helper function to render a node to string for testing
func renderNodeToString(node g.Node) string {
	var sb strings.Builder
	node.Render(&sb)
	return sb.String()
}

func TestDrawerSide_String(t *testing.T) {
	tests := []struct {
		side     DrawerSide
		expected string
	}{
		{DrawerLeft, "drawer-start"},
		{DrawerRight, "drawer-end"},
		{DrawerTop, "drawer-top"},
		{DrawerBottom, "drawer-bottom"},
		{DrawerSide(999), "drawer-start"}, // Invalid value defaults to left
	}
	
	for _, test := range tests {
		if got := test.side.String(); got != test.expected {
			t.Errorf("DrawerSide(%d).String() = %q, want %q", test.side, got, test.expected)
		}
	}
}

func TestDrawerComponent_WithID(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.WithID("custom-id")
	
	if modified.id != "custom-id" {
		t.Errorf("Expected ID to be 'custom-id', got '%s'", modified.id)
	}
	
	// Ensure original is unchanged
	if drawer.id == "custom-id" {
		t.Error("Original drawer component should not be modified")
	}
}

func TestDrawerComponent_WithSide(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.WithSide(DrawerRight)
	
	if modified.side != DrawerRight {
		t.Errorf("Expected side to be DrawerRight, got %v", modified.side)
	}
	
	// Ensure original is unchanged
	if drawer.side == DrawerRight {
		t.Error("Original drawer component should not be modified")
	}
}

func TestDrawerComponent_WithOpen(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.WithOpen(true)
	
	if !modified.open {
		t.Error("Expected drawer to be open")
	}
	
	// Ensure original is unchanged
	if drawer.open {
		t.Error("Original drawer component should not be modified")
	}
}

func TestDrawerComponent_WithOverlay(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.WithOverlay(false)
	
	if modified.overlay {
		t.Error("Expected overlay to be disabled")
	}
	
	// Ensure original is unchanged (default is true)
	if !drawer.overlay {
		t.Error("Original drawer component should not be modified")
	}
}

func TestDrawerComponent_WithClasses(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.WithClasses("custom-class", "another-class")
	
	if len(modified.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(modified.classes))
	}
	if modified.classes[0] != "custom-class" || modified.classes[1] != "another-class" {
		t.Errorf("Expected classes to be ['custom-class', 'another-class'], got %v", modified.classes)
	}
	
	// Ensure original is unchanged
	if len(drawer.classes) != 0 {
		t.Error("Original drawer component should not be modified")
	}
}

func TestDrawerComponent_With(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	modified := drawer.With(DrawerRight)
	
	modifiedDrawer, ok := modified.(*DrawerComponent)
	if !ok {
		t.Fatal("Expected modified component to be *DrawerComponent")
	}
	
	if modifiedDrawer.side != DrawerRight {
		t.Errorf("Expected side to be DrawerRight, got %v", modifiedDrawer.side)
	}
}

func TestDrawerComponent_Render(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	html := renderToStringDrawer(drawer)
	
	// Check for basic structure
	if !strings.Contains(html, `class="drawer drawer-start"`) {
		t.Error("Expected drawer container with start class")
	}
	if !strings.Contains(html, `data-component="drawer"`) {
		t.Error("Expected data-component attribute")
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Error("Expected checkbox input")
	}
	if !strings.Contains(html, `class="drawer-toggle"`) {
		t.Error("Expected drawer-toggle class on checkbox")
	}
	if !strings.Contains(html, `class="drawer-content flex flex-col"`) {
		t.Error("Expected drawer-content class")
	}
	if !strings.Contains(html, "Main content") {
		t.Error("Expected main content text")
	}
	if !strings.Contains(html, `class="drawer-side"`) {
		t.Error("Expected drawer-side class")
	}
	if !strings.Contains(html, `class="drawer-overlay"`) {
		t.Error("Expected drawer-overlay class")
	}
	if !strings.Contains(html, `class="bg-base-200 min-h-full w-80 p-4"`) {
		t.Error("Expected sidebar styling classes")
	}
	if !strings.Contains(html, "Sidebar content") {
		t.Error("Expected sidebar content text")
	}
}

func TestDrawerComponent_RenderOpen(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content")).WithOpen(true)
	html := renderToStringDrawer(drawer)
	
	if !strings.Contains(html, "checked") {
		t.Error("Expected checkbox to be checked when drawer is open")
	}
}

func TestDrawerComponent_RenderSide(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content")).WithSide(DrawerRight)
	html := renderToStringDrawer(drawer)
	
	if !strings.Contains(html, `class="drawer drawer-end"`) {
		t.Error("Expected drawer container with end class")
	}
}

func TestDrawerComponent_RenderNoOverlay(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content")).WithOverlay(false)
	html := renderToStringDrawer(drawer)
	
	if strings.Contains(html, "drawer-overlay") {
		t.Error("Should not contain overlay when disabled")
	}
}

func TestDrawerComponent_RenderWithClasses(t *testing.T) {
	drawer := NewDrawer(g.Text("Main content"), g.Text("Sidebar content")).WithClasses("custom-class")
	html := renderToStringDrawer(drawer)
	
	if !strings.Contains(html, "custom-class") {
		t.Error("Expected custom class in rendered HTML")
	}
}

func TestDrawerToggleButton(t *testing.T) {
	button := DrawerToggleButton("test-drawer", "☰")
	html := renderNodeToString(button)
	
	if !strings.Contains(html, `for="test-drawer-toggle"`) {
		t.Error("Expected for attribute pointing to drawer toggle")
	}
	if !strings.Contains(html, `class="btn btn-square btn-ghost drawer-button"`) {
		t.Error("Expected button classes")
	}
	if !strings.Contains(html, "☰") {
		t.Error("Expected button text")
	}
}

func TestDrawerCloseButton(t *testing.T) {
	button := DrawerCloseButton("test-drawer", "✕")
	html := renderNodeToString(button)
	
	if !strings.Contains(html, `for="test-drawer-toggle"`) {
		t.Error("Expected for attribute pointing to drawer toggle")
	}
	if !strings.Contains(html, `class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"`) {
		t.Error("Expected close button classes")
	}
	if !strings.Contains(html, "✕") {
		t.Error("Expected close button text")
	}
}

func TestDrawerComponent_Immutability(t *testing.T) {
	original := NewDrawer(g.Text("Main content"), g.Text("Sidebar content"))
	originalID := original.id
	originalSide := original.side
	originalOpen := original.open
	originalOverlay := original.overlay
	originalClassCount := len(original.classes)
	
	// Apply various modifications
	_ = original.WithID("new-id")
	_ = original.WithSide(DrawerRight)
	_ = original.WithOpen(true)
	_ = original.WithOverlay(false)
	_ = original.WithClasses("new-class")
	_ = original.With(DrawerBottom)
	
	// Verify original is unchanged
	if original.id != originalID {
		t.Error("Original ID should not change")
	}
	if original.side != originalSide {
		t.Error("Original side should not change")
	}
	if original.open != originalOpen {
		t.Error("Original open state should not change")
	}
	if original.overlay != originalOverlay {
		t.Error("Original overlay state should not change")
	}
	if len(original.classes) != originalClassCount {
		t.Error("Original classes should not change")
	}
}