package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringInput renders an input component to a string for testing
func renderToStringInput(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestInputComponent_WithID(t *testing.T) {
	input := NewInput().WithID("test-id")
	html := renderToStringInput(input)
	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected input to have id='test-id', got: %s", html)
	}
}

func TestInputComponent_WithName(t *testing.T) {
	input := NewInput().WithName("test-name")
	html := renderToStringInput(input)
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected input to have name='test-name', got: %s", html)
	}
}

func TestInputComponent_WithValue(t *testing.T) {
	input := NewInput().WithValue("test-value")
	html := renderToStringInput(input)
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("Expected input to have value='test-value', got: %s", html)
	}
}

func TestInputComponent_WithPlaceholder(t *testing.T) {
	input := NewInput().WithPlaceholder("Enter text")
	html := renderToStringInput(input)
	if !strings.Contains(html, `placeholder="Enter text"`) {
		t.Errorf("Expected input to have placeholder='Enter text', got: %s", html)
	}
}

func TestInputComponent_WithType(t *testing.T) {
	input := NewInput().WithType(InputTypePassword)
	html := renderToStringInput(input)
	if !strings.Contains(html, `type="password"`) {
		t.Errorf("Expected input to have type='password', got: %s", html)
	}
}

func TestInputComponent_WithDisabled(t *testing.T) {
	input := NewInput().WithDisabled(true)
	html := renderToStringInput(input)
	if !strings.Contains(html, `disabled`) {
		t.Errorf("Expected input to be disabled, got: %s", html)
	}
}

func TestInputComponent_WithReadonly(t *testing.T) {
	input := NewInput().WithReadonly(true)
	html := renderToStringInput(input)
	if !strings.Contains(html, `readonly`) {
		t.Errorf("Expected input to be readonly, got: %s", html)
	}
}

func TestInputComponent_WithRequired(t *testing.T) {
	input := NewInput().WithRequired(true)
	html := renderToStringInput(input)
	if !strings.Contains(html, `required`) {
		t.Errorf("Expected input to be required, got: %s", html)
	}
}

func TestInputComponent_WithColor(t *testing.T) {
	input := NewInput().WithColor(flyon.Secondary)
	html := renderToStringInput(input)
	if !strings.Contains(html, "input-secondary") {
		t.Errorf("Expected input to have secondary color class, got: %s", html)
	}
}

func TestInputComponent_WithSize(t *testing.T) {
	input := NewInput().WithSize(flyon.SizeLarge)
	html := renderToStringInput(input)
	if !strings.Contains(html, "input-lg") {
		t.Errorf("Expected input to have large size class, got: %s", html)
	}
}

func TestInputComponent_WithClasses(t *testing.T) {
	input := NewInput().WithClasses("custom-class", "another-class")
	html := renderToStringInput(input)
	if !strings.Contains(html, "custom-class") || !strings.Contains(html, "another-class") {
		t.Errorf("Expected input to have custom classes, got: %s", html)
	}
}

func TestInputComponent_With(t *testing.T) {
	input := NewInput().With(flyon.Success, flyon.SizeSmall, "custom-class")
	html := renderToStringInput(input)
	if !strings.Contains(html, "input-success") {
		t.Errorf("Expected input to have success color class, got: %s", html)
	}
	if !strings.Contains(html, "input-sm") {
		t.Errorf("Expected input to have small size class, got: %s", html)
	}
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected input to have custom class, got: %s", html)
	}
}

func TestInputComponent_Render(t *testing.T) {
	input := NewInput()
	html := renderToStringInput(input)
	
	// Check for basic input structure
	if !strings.Contains(html, "<input") {
		t.Errorf("Expected input element, got: %s", html)
	}
	
	// Check for default classes
	if !strings.Contains(html, "input") {
		t.Errorf("Expected input class, got: %s", html)
	}
	if !strings.Contains(html, "input-bordered") {
		t.Errorf("Expected input-bordered class, got: %s", html)
	}
	
	// Check for default type
	if !strings.Contains(html, `type="text"`) {
		t.Errorf("Expected default type='text', got: %s", html)
	}
}

func TestInputComponent_RenderWithAllAttributes(t *testing.T) {
	input := NewInput().
		WithID("test-input").
		WithName("test-name").
		WithValue("test-value").
		WithPlaceholder("Enter text").
		WithType(InputTypeEmail).
		WithDisabled(true).
		WithReadonly(true).
		WithRequired(true).
		WithColor(flyon.Primary).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class")
	
	html := renderToStringInput(input)
	
	// Check all attributes are present
	expectedAttrs := []string{
		`id="test-input"`,
		`name="test-name"`,
		`value="test-value"`,
		`placeholder="Enter text"`,
		`type="email"`,
		`disabled`,
		`readonly`,
		`required`,
		"input-lg",
		"custom-class",
	}
	
	for _, attr := range expectedAttrs {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected attribute %s in HTML, got: %s", attr, html)
		}
	}
}

func TestInputComponent_RenderNotDisabled(t *testing.T) {
	input := NewInput().WithDisabled(false)
	html := renderToStringInput(input)
	if strings.Contains(html, "disabled") {
		t.Errorf("Expected input to not be disabled, got: %s", html)
	}
}

func TestInputComponent_RenderNotReadonly(t *testing.T) {
	input := NewInput().WithReadonly(false)
	html := renderToStringInput(input)
	if strings.Contains(html, "readonly") {
		t.Errorf("Expected input to not be readonly, got: %s", html)
	}
}

func TestInputComponent_RenderNotRequired(t *testing.T) {
	input := NewInput().WithRequired(false)
	html := renderToStringInput(input)
	if strings.Contains(html, "required") {
		t.Errorf("Expected input to not be required, got: %s", html)
	}
}

func TestInputComponent_RenderDefaultColor(t *testing.T) {
	input := NewInput()
	html := renderToStringInput(input)
	// Primary color should not add a class (it's the default)
	if strings.Contains(html, "input-primary") {
		t.Errorf("Expected no primary color class for default, got: %s", html)
	}
}

func TestInputComponent_RenderDefaultSize(t *testing.T) {
	input := NewInput()
	html := renderToStringInput(input)
	// Medium size should not add a class (it's the default)
	if strings.Contains(html, "input-md") {
		t.Errorf("Expected no medium size class for default, got: %s", html)
	}
}

func TestInputComponent_Immutability(t *testing.T) {
	original := NewInput().WithValue("original")
	modified := original.WithValue("modified")
	
	originalHTML := renderToStringInput(original)
	modifiedHTML := renderToStringInput(modified)
	
	if !strings.Contains(originalHTML, `value="original"`) {
		t.Errorf("Original input should still have original value, got: %s", originalHTML)
	}
	
	if !strings.Contains(modifiedHTML, `value="modified"`) {
		t.Errorf("Modified input should have modified value, got: %s", modifiedHTML)
	}
	
	if strings.Contains(originalHTML, `value="modified"`) {
		t.Errorf("Original input should not be affected by modification, got: %s", originalHTML)
	}
}