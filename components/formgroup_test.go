//go:build !js && !wasm

package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringFormGroup renders a FormGroupComponent to a string for testing
func renderToStringFormGroup(component flyon.Component) string {
	var builder strings.Builder
	component.Render(&builder)
	return builder.String()
}

func TestNewFormGroup(t *testing.T) {
	formGroup := NewFormGroup()

	if formGroup.id != "" {
		t.Errorf("Expected empty id, got %s", formGroup.id)
	}
	if formGroup.label != "" {
		t.Errorf("Expected empty label, got %s", formGroup.label)
	}
	if formGroup.description != "" {
		t.Errorf("Expected empty description, got %s", formGroup.description)
	}
	if formGroup.required {
		t.Errorf("Expected required to be false, got %t", formGroup.required)
	}
	if formGroup.error != "" {
		t.Errorf("Expected empty error, got %s", formGroup.error)
	}
	if formGroup.input != nil {
		t.Errorf("Expected input to be nil, got %v", formGroup.input)
	}
	if len(formGroup.classes) != 0 {
		t.Errorf("Expected empty classes, got %v", formGroup.classes)
	}
	if len(formGroup.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %v", formGroup.attributes)
	}
}

func TestFormGroupComponent_WithID(t *testing.T) {
	formGroup := NewFormGroup().WithID("test-id")

	if formGroup.id != "test-id" {
		t.Errorf("Expected id to be 'test-id', got %s", formGroup.id)
	}
}

func TestFormGroupComponent_WithLabel(t *testing.T) {
	formGroup := NewFormGroup().WithLabel("Test Label")

	if formGroup.label != "Test Label" {
		t.Errorf("Expected label to be 'Test Label', got %s", formGroup.label)
	}
}

func TestFormGroupComponent_WithDescription(t *testing.T) {
	formGroup := NewFormGroup().WithDescription("Test description")

	if formGroup.description != "Test description" {
		t.Errorf("Expected description to be 'Test description', got %s", formGroup.description)
	}
}

func TestFormGroupComponent_WithRequired(t *testing.T) {
	tests := []struct {
		name     string
		required bool
	}{
		{"Required true", true},
		{"Required false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formGroup := NewFormGroup().WithRequired(tt.required)

			if formGroup.required != tt.required {
				t.Errorf("Expected required to be %t, got %t", tt.required, formGroup.required)
			}
		})
	}
}

func TestFormGroupComponent_WithError(t *testing.T) {
	formGroup := NewFormGroup().WithError("Test error")

	if formGroup.error != "Test error" {
		t.Errorf("Expected error to be 'Test error', got %s", formGroup.error)
	}
}

func TestFormGroupComponent_WithInput(t *testing.T) {
	input := NewInput().WithID("test-input")
	formGroup := NewFormGroup().WithInput(input)

	if formGroup.input != input {
		t.Errorf("Expected input to be set, got %v", formGroup.input)
	}
}

func TestFormGroupComponent_WithClasses(t *testing.T) {
	formGroup := NewFormGroup().WithClasses("custom-class", "another-class")
	html := renderToStringFormGroup(formGroup)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got %s", html)
	}
}

func TestFormGroupComponent_With(t *testing.T) {
	formGroup := NewFormGroup().With("data-test", "value")
	html := renderToStringFormGroup(formGroup)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestFormGroupComponent_Render(t *testing.T) {
	formGroup := NewFormGroup()
	html := renderToStringFormGroup(formGroup)

	// Should contain basic form group structure
	if !strings.Contains(html, "form-control") {
		t.Errorf("Expected HTML to contain form-control class, got %s", html)
	}
}

func TestFormGroupComponent_RenderWithLabel(t *testing.T) {
	formGroup := NewFormGroup().WithLabel("Test Label")
	html := renderToStringFormGroup(formGroup)

	// Should contain label
	if !strings.Contains(html, "<label") {
		t.Errorf("Expected HTML to contain label element, got %s", html)
	}
	if !strings.Contains(html, "Test Label") {
		t.Errorf("Expected HTML to contain label text, got %s", html)
	}
}

func TestFormGroupComponent_RenderWithRequired(t *testing.T) {
	formGroup := NewFormGroup().WithLabel("Test Label").WithRequired(true)
	html := renderToStringFormGroup(formGroup)

	// Should contain required indicator
	if !strings.Contains(html, "*") {
		t.Errorf("Expected HTML to contain required indicator (*), got %s", html)
	}
}

func TestFormGroupComponent_RenderWithDescription(t *testing.T) {
	formGroup := NewFormGroup().WithDescription("Test description")
	html := renderToStringFormGroup(formGroup)

	// Should contain description
	if !strings.Contains(html, "Test description") {
		t.Errorf("Expected HTML to contain description text, got %s", html)
	}
	if !strings.Contains(html, "label-text-alt") {
		t.Errorf("Expected HTML to contain label-text-alt class, got %s", html)
	}
}

func TestFormGroupComponent_RenderWithError(t *testing.T) {
	formGroup := NewFormGroup().WithError("Test error")
	html := renderToStringFormGroup(formGroup)

	// Should contain error message
	if !strings.Contains(html, "Test error") {
		t.Errorf("Expected HTML to contain error text, got %s", html)
	}
	if !strings.Contains(html, "label-text-alt") {
		t.Errorf("Expected HTML to contain label-text-alt class, got %s", html)
	}
	if !strings.Contains(html, "text-error") {
		t.Errorf("Expected HTML to contain text-error class, got %s", html)
	}
}

func TestFormGroupComponent_RenderWithInput(t *testing.T) {
	input := NewInput().WithID("test-input").WithName("test-name")
	formGroup := NewFormGroup().WithInput(input)
	html := renderToStringFormGroup(formGroup)

	// Should contain input
	if !strings.Contains(html, `id="test-input"`) {
		t.Errorf("Expected HTML to contain input id, got %s", html)
	}
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected HTML to contain input name, got %s", html)
	}
}

func TestFormGroupComponent_RenderWithAllAttributes(t *testing.T) {
	input := NewInput().WithID("test-input")
	formGroup := NewFormGroup().
		WithID("test-id").
		WithLabel("Test Label").
		WithDescription("Test description").
		WithRequired(true).
		WithError("Test error").
		WithInput(input).
		WithClasses("custom-class").
		With("data-test", "value")

	html := renderToStringFormGroup(formGroup)

	// Check all attributes are present
	expectedContent := []string{
		`id="test-id"`,
		"Test Label",
		"*", // required indicator
		"Test description",
		"Test error",
		`id="test-input"`,
		"custom-class",
		`data-test="value"`,
	}

	for _, content := range expectedContent {
		if !strings.Contains(html, content) {
			t.Errorf("Expected HTML to contain %s, got %s", content, html)
		}
	}
}

func TestFormGroupComponent_Immutability(t *testing.T) {
	original := NewFormGroup()
	modified := original.WithID("new-id").WithLabel("New Label")

	// Original should remain unchanged
	if original.id != "" {
		t.Errorf("Expected original id to remain empty, got %s", original.id)
	}
	if original.label != "" {
		t.Errorf("Expected original label to remain empty, got %s", original.label)
	}

	// Modified should have new values
	if modified.id != "new-id" {
		t.Errorf("Expected modified id to be 'new-id', got %s", modified.id)
	}
	if modified.label != "New Label" {
		t.Errorf("Expected modified label to be 'New Label', got %s", modified.label)
	}
}