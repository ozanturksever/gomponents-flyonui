//go:build !js && !wasm

package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringFormValidation renders a FormValidationComponent to a string for testing
func renderToStringFormValidation(component flyon.Component) string {
	var builder strings.Builder
	component.Render(&builder)
	return builder.String()
}

func TestNewFormValidation(t *testing.T) {
	validation := NewFormValidation()

	if validation.id != "" {
		t.Errorf("Expected empty id, got %s", validation.id)
	}
	if validation.message != "" {
		t.Errorf("Expected empty message, got %s", validation.message)
	}
	if validation.validationType != ValidationTypeError {
		t.Errorf("Expected validation type to be error, got %s", validation.validationType)
	}
	if validation.visible {
		t.Errorf("Expected visible to be false, got %t", validation.visible)
	}
	if len(validation.classes) != 0 {
		t.Errorf("Expected empty classes, got %v", validation.classes)
	}
	if len(validation.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %v", validation.attributes)
	}
}

func TestFormValidationComponent_WithID(t *testing.T) {
	validation := NewFormValidation().WithID("test-id")

	if validation.id != "test-id" {
		t.Errorf("Expected id to be 'test-id', got %s", validation.id)
	}
}

func TestFormValidationComponent_WithMessage(t *testing.T) {
	validation := NewFormValidation().WithMessage("Test message")

	if validation.message != "Test message" {
		t.Errorf("Expected message to be 'Test message', got %s", validation.message)
	}
}

func TestFormValidationComponent_WithType(t *testing.T) {
	tests := []struct {
		name           string
		validationType ValidationType
	}{
		{"Error type", ValidationTypeError},
		{"Warning type", ValidationTypeWarning},
		{"Success type", ValidationTypeSuccess},
		{"Info type", ValidationTypeInfo},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validation := NewFormValidation().WithType(tt.validationType)

			if validation.validationType != tt.validationType {
				t.Errorf("Expected validation type to be %s, got %s", tt.validationType, validation.validationType)
			}
		})
	}
}

func TestFormValidationComponent_WithVisible(t *testing.T) {
	tests := []struct {
		name    string
		visible bool
	}{
		{"Visible true", true},
		{"Visible false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validation := NewFormValidation().WithVisible(tt.visible)

			if validation.visible != tt.visible {
				t.Errorf("Expected visible to be %t, got %t", tt.visible, validation.visible)
			}
		})
	}
}

func TestFormValidationComponent_WithClasses(t *testing.T) {
	validation := NewFormValidation().WithClasses("custom-class", "another-class").WithVisible(true).WithMessage("Test")
	html := renderToStringFormValidation(validation)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got %s", html)
	}
}

func TestFormValidationComponent_With(t *testing.T) {
	validation := NewFormValidation().WithVisible(true).WithMessage("Test").With("data-test", "value")
	html := renderToStringFormValidation(validation)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestFormValidationComponent_Render(t *testing.T) {
	validation := NewFormValidation().WithMessage("Test message").WithVisible(true)
	html := renderToStringFormValidation(validation)

	// Should contain basic validation structure
	if !strings.Contains(html, "Test message") {
		t.Errorf("Expected HTML to contain message text, got %s", html)
	}
	if !strings.Contains(html, "label-text-alt") {
		t.Errorf("Expected HTML to contain label-text-alt class, got %s", html)
	}
}

func TestFormValidationComponent_RenderHidden(t *testing.T) {
	validation := NewFormValidation().WithMessage("Test message").WithVisible(false)
	html := renderToStringFormValidation(validation)

	// Should be empty when not visible
	if html != "" {
		t.Errorf("Expected empty HTML when not visible, got %s", html)
	}
}

func TestFormValidationComponent_RenderWithErrorType(t *testing.T) {
	validation := NewFormValidation().
		WithMessage("Error message").
		WithType(ValidationTypeError).
		WithVisible(true)
	html := renderToStringFormValidation(validation)

	// Should contain error styling
	if !strings.Contains(html, "text-error") {
		t.Errorf("Expected HTML to contain text-error class, got %s", html)
	}
	if !strings.Contains(html, "Error message") {
		t.Errorf("Expected HTML to contain error message, got %s", html)
	}
}

func TestFormValidationComponent_RenderWithWarningType(t *testing.T) {
	validation := NewFormValidation().
		WithMessage("Warning message").
		WithType(ValidationTypeWarning).
		WithVisible(true)
	html := renderToStringFormValidation(validation)

	// Should contain warning styling
	if !strings.Contains(html, "text-warning") {
		t.Errorf("Expected HTML to contain text-warning class, got %s", html)
	}
	if !strings.Contains(html, "Warning message") {
		t.Errorf("Expected HTML to contain warning message, got %s", html)
	}
}

func TestFormValidationComponent_RenderWithSuccessType(t *testing.T) {
	validation := NewFormValidation().
		WithMessage("Success message").
		WithType(ValidationTypeSuccess).
		WithVisible(true)
	html := renderToStringFormValidation(validation)

	// Should contain success styling
	if !strings.Contains(html, "text-success") {
		t.Errorf("Expected HTML to contain text-success class, got %s", html)
	}
	if !strings.Contains(html, "Success message") {
		t.Errorf("Expected HTML to contain success message, got %s", html)
	}
}

func TestFormValidationComponent_RenderWithInfoType(t *testing.T) {
	validation := NewFormValidation().
		WithMessage("Info message").
		WithType(ValidationTypeInfo).
		WithVisible(true)
	html := renderToStringFormValidation(validation)

	// Should contain info styling
	if !strings.Contains(html, "text-info") {
		t.Errorf("Expected HTML to contain text-info class, got %s", html)
	}
	if !strings.Contains(html, "Info message") {
		t.Errorf("Expected HTML to contain info message, got %s", html)
	}
}

func TestFormValidationComponent_RenderWithAllAttributes(t *testing.T) {
	validation := NewFormValidation().
		WithID("test-id").
		WithMessage("Test message").
		WithType(ValidationTypeError).
		WithVisible(true).
		WithClasses("custom-class").
		With("data-test", "value")

	html := renderToStringFormValidation(validation)

	// Check all attributes are present
	expectedContent := []string{
		`id="test-id"`,
		"Test message",
		"text-error",
		"custom-class",
		`data-test="value"`,
	}

	for _, content := range expectedContent {
		if !strings.Contains(html, content) {
			t.Errorf("Expected HTML to contain %s, got %s", content, html)
		}
	}
}

func TestFormValidationComponent_Immutability(t *testing.T) {
	original := NewFormValidation()
	modified := original.WithID("new-id").WithMessage("New message")

	// Original should remain unchanged
	if original.id != "" {
		t.Errorf("Expected original id to remain empty, got %s", original.id)
	}
	if original.message != "" {
		t.Errorf("Expected original message to remain empty, got %s", original.message)
	}

	// Modified should have new values
	if modified.id != "new-id" {
		t.Errorf("Expected modified id to be 'new-id', got %s", modified.id)
	}
	if modified.message != "New message" {
		t.Errorf("Expected modified message to be 'New message', got %s", modified.message)
	}
}