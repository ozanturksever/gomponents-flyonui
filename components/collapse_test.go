package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
)

// renderToString is a helper function to render a component to string for testing
func renderToString(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestCollapseComponent_WithID(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithID("custom-id")
	
	if modified.id != "custom-id" {
		t.Errorf("Expected ID to be 'custom-id', got '%s'", modified.id)
	}
	
	// Ensure original is unchanged
	if collapse.id == "custom-id" {
		t.Error("Original collapse component should not be modified")
	}
}

func TestCollapseComponent_WithOpen(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithOpen(true)
	
	if !modified.open {
		t.Error("Expected collapse to be open")
	}
	
	// Ensure original is unchanged
	if collapse.open {
		t.Error("Original collapse component should not be modified")
	}
}

func TestCollapseComponent_WithArrow(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithArrow(false)
	
	if modified.arrow {
		t.Error("Expected arrow to be disabled")
	}
	
	// Test enabling arrow disables plus
	collapseWithPlus := collapse.WithPlus(true)
	modifiedWithArrow := collapseWithPlus.WithArrow(true)
	
	if !modifiedWithArrow.arrow {
		t.Error("Expected arrow to be enabled")
	}
	if modifiedWithArrow.plus {
		t.Error("Expected plus to be disabled when arrow is enabled")
	}
}

func TestCollapseComponent_WithPlus(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithPlus(true)
	
	if !modified.plus {
		t.Error("Expected plus to be enabled")
	}
	if modified.arrow {
		t.Error("Expected arrow to be disabled when plus is enabled")
	}
	
	// Ensure original is unchanged
	if collapse.plus {
		t.Error("Original collapse component should not be modified")
	}
}

func TestCollapseComponent_WithColor(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithColor(flyon.Secondary)
	
	if modified.color != flyon.Secondary {
		t.Errorf("Expected color to be Secondary, got %v", modified.color)
	}
	
	// Ensure original is unchanged
	if collapse.color == flyon.Secondary {
		t.Error("Original collapse component should not be modified")
	}
}

func TestCollapseComponent_WithClasses(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.WithClasses("custom-class", "another-class")
	
	if len(modified.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(modified.classes))
	}
	if modified.classes[0] != "custom-class" || modified.classes[1] != "another-class" {
		t.Errorf("Expected classes to be ['custom-class', 'another-class'], got %v", modified.classes)
	}
	
	// Ensure original is unchanged
	if len(collapse.classes) != 0 {
		t.Error("Original collapse component should not be modified")
	}
}

func TestCollapseComponent_With(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	modified := collapse.With(flyon.Success)
	
	modifiedCollapse, ok := modified.(*CollapseComponent)
	if !ok {
		t.Fatal("Expected modified component to be *CollapseComponent")
	}
	
	if modifiedCollapse.color != flyon.Success {
		t.Errorf("Expected color to be Success, got %v", modifiedCollapse.color)
	}
}

func TestCollapseComponent_Render(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content"))
	html := renderToString(collapse)
	
	// Check for basic structure
	if !strings.Contains(html, `class="collapse collapse-arrow"`) {
		t.Error("Expected collapse container with arrow class")
	}
	if !strings.Contains(html, `data-component="collapse"`) {
		t.Error("Expected data-component attribute")
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Error("Expected checkbox input")
	}
	if !strings.Contains(html, `class="collapse-toggle"`) {
		t.Error("Expected collapse-toggle class on checkbox")
	}
	if !strings.Contains(html, `class="collapse-title text-xl font-medium cursor-pointer"`) {
		t.Error("Expected collapse-title class")
	}
	if !strings.Contains(html, "Test Title") {
		t.Error("Expected title text")
	}
	if !strings.Contains(html, `class="collapse-content"`) {
		t.Error("Expected collapse-content class")
	}
	if !strings.Contains(html, "Test content") {
		t.Error("Expected content text")
	}
}

func TestCollapseComponent_RenderOpen(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content")).WithOpen(true)
	html := renderToString(collapse)
	
	if !strings.Contains(html, "checked") {
		t.Error("Expected checkbox to be checked when collapse is open")
	}
}

func TestCollapseComponent_RenderPlus(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content")).WithPlus(true)
	html := renderToString(collapse)
	
	if !strings.Contains(html, `class="collapse collapse-plus"`) {
		t.Error("Expected collapse container with plus class")
	}
}

func TestCollapseComponent_RenderNoIndicator(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content")).WithArrow(false)
	html := renderToString(collapse)
	
	if !strings.Contains(html, `class="collapse"`) {
		t.Error("Expected basic collapse class")
	}
	if strings.Contains(html, "collapse-arrow") {
		t.Error("Should not contain arrow class when disabled")
	}
	if strings.Contains(html, "collapse-plus") {
		t.Error("Should not contain plus class when disabled")
	}
}

func TestCollapseComponent_RenderWithClasses(t *testing.T) {
	collapse := NewCollapse("Test Title", g.Text("Test content")).WithClasses("custom-class")
	html := renderToString(collapse)
	
	if !strings.Contains(html, "custom-class") {
		t.Error("Expected custom class in rendered HTML")
	}
}

func TestCollapseComponent_Immutability(t *testing.T) {
	original := NewCollapse("Test Title", g.Text("Test content"))
	originalID := original.id
	originalOpen := original.open
	originalArrow := original.arrow
	originalPlus := original.plus
	originalColor := original.color
	originalClassCount := len(original.classes)
	
	// Apply various modifications
	_ = original.WithID("new-id")
	_ = original.WithOpen(true)
	_ = original.WithArrow(false)
	_ = original.WithPlus(true)
	_ = original.WithColor(flyon.Secondary)
	_ = original.WithClasses("new-class")
	_ = original.With(flyon.Success)
	
	// Verify original is unchanged
	if original.id != originalID {
		t.Error("Original ID should not change")
	}
	if original.open != originalOpen {
		t.Error("Original open state should not change")
	}
	if original.arrow != originalArrow {
		t.Error("Original arrow state should not change")
	}
	if original.plus != originalPlus {
		t.Error("Original plus state should not change")
	}
	if original.color != originalColor {
		t.Error("Original color should not change")
	}
	if len(original.classes) != originalClassCount {
		t.Error("Original classes should not change")
	}
}