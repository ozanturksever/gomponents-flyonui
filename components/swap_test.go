package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// renderToStringSwap is a helper function to render a component to string for testing
func renderToStringSwap(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestSwapComponent_WithID(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithID("test-swap")
	
	if swap.id != "test-swap" {
		t.Errorf("Expected id to be 'test-swap', got '%s'", swap.id)
	}
}

func TestSwapComponent_WithActive(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithActive(true)
	
	if !swap.active {
		t.Error("Expected active to be true")
	}
}

func TestSwapComponent_WithRotate(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithRotate(true)
	
	if !swap.rotate {
		t.Error("Expected rotate to be true")
	}
}

func TestSwapComponent_WithFlip(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithFlip(true)
	
	if !swap.flip {
		t.Error("Expected flip to be true")
	}
}

func TestSwapComponent_WithColor(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithColor(flyon.Secondary)
	
	if swap.color != flyon.Secondary {
		t.Errorf("Expected color to be Secondary, got %v", swap.color)
	}
}

func TestSwapComponent_WithClasses(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithClasses("custom-class", "another-class")
	
	if len(swap.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(swap.classes))
	}
	if swap.classes[0] != "custom-class" {
		t.Errorf("Expected first class to be 'custom-class', got '%s'", swap.classes[0])
	}
	if swap.classes[1] != "another-class" {
		t.Errorf("Expected second class to be 'another-class', got '%s'", swap.classes[1])
	}
}

func TestSwapComponent_With(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).With(flyon.Success)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "swap-success") {
		t.Error("Expected HTML to contain swap-success class")
	}
}

func TestSwapComponent_Render(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, `class="swap"`) {
		t.Error("Expected HTML to contain swap class")
	}
	if !strings.Contains(html, `class="swap-input"`) {
		t.Error("Expected HTML to contain swap-input class")
	}
	if !strings.Contains(html, `class="swap-on"`) {
		t.Error("Expected HTML to contain swap-on class")
	}
	if !strings.Contains(html, `class="swap-off"`) {
		t.Error("Expected HTML to contain swap-off class")
	}
	if !strings.Contains(html, "ON") {
		t.Error("Expected HTML to contain ON text")
	}
	if !strings.Contains(html, "OFF") {
		t.Error("Expected HTML to contain OFF text")
	}
}

func TestSwapComponent_RenderWithID(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithID("test-swap")
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, `id="test-swap"`) {
		t.Error("Expected HTML to contain id attribute")
	}
}

func TestSwapComponent_RenderActive(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithActive(true)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "checked") {
		t.Error("Expected HTML to contain checked attribute")
	}
}

func TestSwapComponent_RenderRotate(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithRotate(true)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "swap-rotate") {
		t.Error("Expected HTML to contain swap-rotate class")
	}
}

func TestSwapComponent_RenderFlip(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithFlip(true)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "swap-flip") {
		t.Error("Expected HTML to contain swap-flip class")
	}
}

func TestSwapComponent_RenderWithColor(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithColor(flyon.Secondary)
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "swap-secondary") {
		t.Error("Expected HTML to contain swap-secondary class")
	}
}

func TestSwapComponent_RenderWithClasses(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	swap := NewSwap(onState, offState).WithClasses("custom-class")
	html := renderToStringSwap(swap)
	
	if !strings.Contains(html, "custom-class") {
		t.Error("Expected HTML to contain custom-class")
	}
}

func TestSwapComponent_Immutability(t *testing.T) {
	onState := h.Span(g.Text("ON"))
	offState := h.Span(g.Text("OFF"))
	original := NewSwap(onState, offState)
	
	// Test that methods return new instances
	withID := original.WithID("test")
	if withID == original {
		t.Error("WithID should return a new instance")
	}
	
	withActive := original.WithActive(true)
	if withActive == original {
		t.Error("WithActive should return a new instance")
	}
	
	withRotate := original.WithRotate(true)
	if withRotate == original {
		t.Error("WithRotate should return a new instance")
	}
	
	withFlip := original.WithFlip(true)
	if withFlip == original {
		t.Error("WithFlip should return a new instance")
	}
	
	withColor := original.WithColor(flyon.Secondary)
	if withColor == original {
		t.Error("WithColor should return a new instance")
	}
	
	withClasses := original.WithClasses("test")
	if withClasses == original {
		t.Error("WithClasses should return a new instance")
	}
	
	withModifier := original.With(flyon.Success)
	if withModifier == original {
		t.Error("With should return a new instance")
	}
	
	// Test that original is unchanged
	if original.id != "" {
		t.Error("Original id should be unchanged")
	}
	if original.active {
		t.Error("Original active should be unchanged")
	}
	if original.rotate {
		t.Error("Original rotate should be unchanged")
	}
	if original.flip {
		t.Error("Original flip should be unchanged")
	}
	if original.color != flyon.Primary {
		t.Error("Original color should be unchanged")
	}
	if len(original.classes) != 0 {
		t.Error("Original classes should be unchanged")
	}
}