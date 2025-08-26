package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringRange renders a range component to string for testing
func renderToStringRange(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestRangeComponent_WithID(t *testing.T) {
	range1 := NewRange().WithID("test-range")
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `id="test-range"`) {
		t.Errorf("Expected range to have id='test-range', got: %s", html)
	}
}

func TestRangeComponent_WithName(t *testing.T) {
	range1 := NewRange().WithName("volume")
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `name="volume"`) {
		t.Errorf("Expected range to have name='volume', got: %s", html)
	}
}

func TestRangeComponent_WithValue(t *testing.T) {
	range1 := NewRange().WithValue(75.5)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `value="75.5"`) {
		t.Errorf("Expected range to have value='75.5', got: %s", html)
	}
}

func TestRangeComponent_WithMin(t *testing.T) {
	range1 := NewRange().WithMin(10)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `min="10"`) {
		t.Errorf("Expected range to have min='10', got: %s", html)
	}
}

func TestRangeComponent_WithMax(t *testing.T) {
	range1 := NewRange().WithMax(200)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `max="200"`) {
		t.Errorf("Expected range to have max='200', got: %s", html)
	}
}

func TestRangeComponent_WithStep(t *testing.T) {
	range1 := NewRange().WithStep(0.5)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `step="0.5"`) {
		t.Errorf("Expected range to have step='0.5', got: %s", html)
	}
}

func TestRangeComponent_WithDisabled(t *testing.T) {
	range1 := NewRange().WithDisabled(true)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `disabled`) {
		t.Errorf("Expected range to be disabled, got: %s", html)
	}
}

func TestRangeComponent_WithColor(t *testing.T) {
	range1 := NewRange().WithColor(flyon.Secondary)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, "range-secondary") {
		t.Errorf("Expected range to have range-secondary class, got: %s", html)
	}
}

func TestRangeComponent_WithSize(t *testing.T) {
	range1 := NewRange().WithSize(flyon.SizeLarge)
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, "range-lg") {
		t.Errorf("Expected range to have range-lg class, got: %s", html)
	}
}

func TestRangeComponent_WithClasses(t *testing.T) {
	range1 := NewRange().WithClasses("custom-class", "another-class")
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected range to have custom-class, got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected range to have another-class, got: %s", html)
	}
}

func TestRangeComponent_With(t *testing.T) {
	range1 := NewRange().With(flyon.Success, flyon.SizeSmall, "custom-class")
	html := renderToStringRange(range1)
	if !strings.Contains(html, "range-success") {
		t.Error("Expected HTML to contain 'range-success' class")
	}
	if !strings.Contains(html, "range-sm") {
		t.Error("Expected HTML to contain 'range-sm' class")
	}
	if !strings.Contains(html, "custom-class") {
		t.Error("Expected HTML to contain 'custom-class'")
	}
}

func TestRangeComponent_Render(t *testing.T) {
	range1 := NewRange()
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `<input`) {
		t.Error("Expected HTML to contain input tag")
	}
	if !strings.Contains(html, `type="range"`) {
		t.Error("Expected HTML to contain type='range'")
	}
	if !strings.Contains(html, `class="range"`) {
		t.Error("Expected HTML to contain class='range'")
	}
	if !strings.Contains(html, `min="0"`) {
		t.Error("Expected HTML to contain min='0'")
	}
	if !strings.Contains(html, `max="100"`) {
		t.Error("Expected HTML to contain max='100'")
	}
	if !strings.Contains(html, `step="1"`) {
		t.Error("Expected HTML to contain step='1'")
	}
	if !strings.Contains(html, `value="50"`) {
		t.Error("Expected HTML to contain value='50'")
	}
}

func TestRangeComponent_RenderWithAllAttributes(t *testing.T) {
	range1 := NewRange().
		WithID("volume-slider").
		WithName("volume").
		WithValue(75).
		WithMin(0).
		WithMax(100).
		WithStep(5).
		WithDisabled(true).
		WithColor(flyon.Success).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-range")
	
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `id="volume-slider"`) {
		t.Error("Expected HTML to contain id='volume-slider'")
	}
	if !strings.Contains(html, `name="volume"`) {
		t.Error("Expected HTML to contain name='volume'")
	}
	if !strings.Contains(html, `value="75"`) {
		t.Error("Expected HTML to contain value='75'")
	}
	if !strings.Contains(html, `min="0"`) {
		t.Error("Expected HTML to contain min='0'")
	}
	if !strings.Contains(html, `max="100"`) {
		t.Error("Expected HTML to contain max='100'")
	}
	if !strings.Contains(html, `step="5"`) {
		t.Error("Expected HTML to contain step='5'")
	}
	if !strings.Contains(html, `disabled`) {
		t.Error("Expected HTML to contain disabled attribute")
	}
	if !strings.Contains(html, `range-success`) {
		t.Error("Expected HTML to contain 'range-success' class")
	}
	if !strings.Contains(html, `range-lg`) {
		t.Error("Expected HTML to contain 'range-lg' class")
	}
	if !strings.Contains(html, `custom-range`) {
		t.Error("Expected HTML to contain 'custom-range' class")
	}
}

func TestRangeComponent_RenderEnabled(t *testing.T) {
	range1 := NewRange().WithDisabled(false)
	html := renderToStringRange(range1)
	if strings.Contains(html, "disabled") {
		t.Error("Expected HTML to not contain disabled attribute")
	}
}

func TestRangeComponent_RenderDefaultColor(t *testing.T) {
	range1 := NewRange()
	html := renderToStringRange(range1)
	if strings.Contains(html, "range-primary") {
		t.Error("Expected HTML to not contain 'range-primary' class for default color")
	}
}

func TestRangeComponent_RenderDefaultSize(t *testing.T) {
	range1 := NewRange()
	html := renderToStringRange(range1)
	if strings.Contains(html, "range-md") {
		t.Error("Expected HTML to not contain 'range-md' class for default size")
	}
}

func TestRangeComponent_RenderWithFloatValues(t *testing.T) {
	range1 := NewRange().
		WithMin(0.5).
		WithMax(10.5).
		WithStep(0.25).
		WithValue(5.75)
	
	html := renderToStringRange(range1)
	
	if !strings.Contains(html, `min="0.5"`) {
		t.Error("Expected HTML to contain min='0.5'")
	}
	if !strings.Contains(html, `max="10.5"`) {
		t.Error("Expected HTML to contain max='10.5'")
	}
	if !strings.Contains(html, `step="0.25"`) {
		t.Error("Expected HTML to contain step='0.25'")
	}
	if !strings.Contains(html, `value="5.75"`) {
		t.Error("Expected HTML to contain value='5.75'")
	}
}

func TestRangeComponent_Immutability(t *testing.T) {
	original := NewRange()
	modified := original.WithValue(75).WithColor(flyon.Success)
	
	// Original should remain unchanged
	originalHTML := renderToStringRange(original)
	if !strings.Contains(originalHTML, `value="50"`) {
		t.Error("Original range value should remain unchanged")
	}
	if strings.Contains(originalHTML, "range-success") {
		t.Error("Original range color should remain unchanged")
	}
	
	// Modified should have new values
	modifiedHTML := renderToStringRange(modified)
	if !strings.Contains(modifiedHTML, `value="75"`) {
		t.Error("Modified range should have new value")
	}
	if !strings.Contains(modifiedHTML, "range-success") {
		t.Error("Modified range should have new color")
	}
}