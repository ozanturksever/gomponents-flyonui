//go:build !js && !wasm

package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringFileInput renders a FileInputComponent to a string for testing
func renderToStringFileInput(component flyon.Component) string {
	var builder strings.Builder
	component.Render(&builder)
	return builder.String()
}

func TestNewFileInput(t *testing.T) {
	fileInput := NewFileInput()

	if fileInput.id != "" {
		t.Errorf("Expected empty id, got %s", fileInput.id)
	}
	if fileInput.name != "" {
		t.Errorf("Expected empty name, got %s", fileInput.name)
	}
	if fileInput.accept != "" {
		t.Errorf("Expected empty accept, got %s", fileInput.accept)
	}
	if fileInput.multiple {
		t.Errorf("Expected multiple to be false, got %t", fileInput.multiple)
	}
	if fileInput.disabled {
		t.Errorf("Expected disabled to be false, got %t", fileInput.disabled)
	}
	if fileInput.color != flyon.Primary {
		t.Errorf("Expected color to be Primary, got %v", fileInput.color)
	}
	if fileInput.size != flyon.SizeMedium {
		t.Errorf("Expected size to be SizeMedium, got %v", fileInput.size)
	}
	if fileInput.colorSet {
		t.Errorf("Expected colorSet to be false, got %t", fileInput.colorSet)
	}
	if fileInput.sizeSet {
		t.Errorf("Expected sizeSet to be false, got %t", fileInput.sizeSet)
	}
	if len(fileInput.classes) != 0 {
		t.Errorf("Expected empty classes, got %v", fileInput.classes)
	}
	if len(fileInput.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %v", fileInput.attributes)
	}
}

func TestFileInputComponent_WithID(t *testing.T) {
	fileInput := NewFileInput().WithID("test-id")

	if fileInput.id != "test-id" {
		t.Errorf("Expected id to be 'test-id', got %s", fileInput.id)
	}
}

func TestFileInputComponent_WithName(t *testing.T) {
	fileInput := NewFileInput().WithName("test-name")

	if fileInput.name != "test-name" {
		t.Errorf("Expected name to be 'test-name', got %s", fileInput.name)
	}
}

func TestFileInputComponent_WithAccept(t *testing.T) {
	fileInput := NewFileInput().WithAccept("image/*")

	if fileInput.accept != "image/*" {
		t.Errorf("Expected accept to be 'image/*', got %s", fileInput.accept)
	}
}

func TestFileInputComponent_WithMultiple(t *testing.T) {
	tests := []struct {
		name     string
		multiple bool
	}{
		{"Multiple true", true},
		{"Multiple false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput := NewFileInput().WithMultiple(tt.multiple)

			if fileInput.multiple != tt.multiple {
				t.Errorf("Expected multiple to be %t, got %t", tt.multiple, fileInput.multiple)
			}
		})
	}
}

func TestFileInputComponent_WithDisabled(t *testing.T) {
	tests := []struct {
		name     string
		disabled bool
	}{
		{"Disabled true", true},
		{"Disabled false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput := NewFileInput().WithDisabled(tt.disabled)

			if fileInput.disabled != tt.disabled {
				t.Errorf("Expected disabled to be %t, got %t", tt.disabled, fileInput.disabled)
			}
		})
	}
}

func TestFileInputComponent_WithColor(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary color", flyon.Primary, "file-input-primary"},
		{"Secondary color", flyon.Secondary, "file-input-secondary"},
		{"Success color", flyon.Success, "file-input-success"},
		{"Warning color", flyon.Warning, "file-input-warning"},
		{"Error color", flyon.Error, "file-input-error"},
		{"Info color", flyon.Info, "file-input-info"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput := NewFileInput().WithColor(tt.color)
			html := renderToStringFileInput(fileInput)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected HTML to contain %s, got %s", tt.expected, html)
			}
		})
	}
}

func TestFileInputComponent_WithSize(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small size", flyon.SizeXS, "file-input-xs"},
		{"Small size", flyon.SizeSmall, "file-input-sm"},
		{"Medium size", flyon.SizeMedium, "file-input-md"},
		{"Large size", flyon.SizeLarge, "file-input-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput := NewFileInput().WithSize(tt.size)
			html := renderToStringFileInput(fileInput)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected HTML to contain %s, got %s", tt.expected, html)
			}
		})
	}
}

func TestFileInputComponent_WithClasses(t *testing.T) {
	fileInput := NewFileInput().WithClasses("custom-class", "another-class")
	html := renderToStringFileInput(fileInput)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got %s", html)
	}
}

func TestFileInputComponent_With(t *testing.T) {
	fileInput := NewFileInput().With("data-test", "value")
	html := renderToStringFileInput(fileInput)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestFileInputComponent_Render(t *testing.T) {
	fileInput := NewFileInput()
	html := renderToStringFileInput(fileInput)

	// Should contain basic file input structure
	if !strings.Contains(html, `type="file"`) {
		t.Errorf("Expected HTML to contain type=file, got %s", html)
	}
	if !strings.Contains(html, "file-input") {
		t.Errorf("Expected HTML to contain file-input class, got %s", html)
	}
}

func TestFileInputComponent_RenderWithAllAttributes(t *testing.T) {
	fileInput := NewFileInput().
		WithID("test-id").
		WithName("test-name").
		WithAccept("image/*").
		WithMultiple(true).
		WithDisabled(true).
		WithColor(flyon.Primary).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class").
		With("data-test", "value")

	html := renderToStringFileInput(fileInput)

	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-id"`,
		`name="test-name"`,
		`accept="image/*"`,
		`multiple`,
		`disabled`,
		"file-input-primary",
		"file-input-lg",
		"custom-class",
		`data-test="value"`,
	}

	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected HTML to contain %s, got %s", attr, html)
		}
	}
}

func TestFileInputComponent_Immutability(t *testing.T) {
	original := NewFileInput()
	modified := original.WithID("new-id").WithColor(flyon.Primary)

	// Original should remain unchanged
	if original.id != "" {
		t.Errorf("Expected original id to remain empty, got %s", original.id)
	}
	if original.colorSet {
		t.Errorf("Expected original colorSet to remain false, got %t", original.colorSet)
	}

	// Modified should have new values
	if modified.id != "new-id" {
		t.Errorf("Expected modified id to be 'new-id', got %s", modified.id)
	}
	if !modified.colorSet {
		t.Errorf("Expected modified colorSet to be true, got %t", modified.colorSet)
	}
}