package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// Helper function to render AutocompleteComponent to string
func renderToStringAutocomplete(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestNewAutocomplete(t *testing.T) {
	autocomplete := NewAutocomplete()

	if autocomplete.id != "" {
		t.Errorf("Expected empty ID, got %s", autocomplete.id)
	}
	if autocomplete.name != "" {
		t.Errorf("Expected empty name, got %s", autocomplete.name)
	}
	if autocomplete.placeholder != "" {
		t.Errorf("Expected empty placeholder, got %s", autocomplete.placeholder)
	}
	if autocomplete.value != "" {
		t.Errorf("Expected empty value, got %s", autocomplete.value)
	}
	if autocomplete.disabled != false {
		t.Errorf("Expected disabled to be false, got %t", autocomplete.disabled)
	}
	if autocomplete.color != flyon.Primary {
		t.Errorf("Expected color to be Primary, got %v", autocomplete.color)
	}
	if autocomplete.size != flyon.SizeMedium {
		t.Errorf("Expected size to be SizeMedium, got %v", autocomplete.size)
	}
	if len(autocomplete.options) != 0 {
		t.Errorf("Expected empty options, got %v", autocomplete.options)
	}
	if len(autocomplete.classes) != 0 {
		t.Errorf("Expected empty classes, got %v", autocomplete.classes)
	}
	if len(autocomplete.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %v", autocomplete.attributes)
	}
}

func TestAutocompleteComponent_WithID(t *testing.T) {
	autocomplete := NewAutocomplete().WithID("test-id")

	if autocomplete.id != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got %s", autocomplete.id)
	}
}

func TestAutocompleteComponent_WithName(t *testing.T) {
	autocomplete := NewAutocomplete().WithName("test-name")

	if autocomplete.name != "test-name" {
		t.Errorf("Expected name to be 'test-name', got %s", autocomplete.name)
	}
}

func TestAutocompleteComponent_WithPlaceholder(t *testing.T) {
	autocomplete := NewAutocomplete().WithPlaceholder("Search...")

	if autocomplete.placeholder != "Search..." {
		t.Errorf("Expected placeholder to be 'Search...', got %s", autocomplete.placeholder)
	}
}

func TestAutocompleteComponent_WithValue(t *testing.T) {
	autocomplete := NewAutocomplete().WithValue("test-value")

	if autocomplete.value != "test-value" {
		t.Errorf("Expected value to be 'test-value', got %s", autocomplete.value)
	}
}

func TestAutocompleteComponent_WithDisabled(t *testing.T) {
	autocomplete := NewAutocomplete().WithDisabled(true)

	if autocomplete.disabled != true {
		t.Errorf("Expected disabled to be true, got %t", autocomplete.disabled)
	}
}

func TestAutocompleteComponent_WithColor(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected flyon.Color
	}{
		{"Primary color", flyon.Primary, flyon.Primary},
		{"Secondary color", flyon.Secondary, flyon.Secondary},
		{"Success color", flyon.Success, flyon.Success},
		{"Warning color", flyon.Warning, flyon.Warning},
		{"Error color", flyon.Error, flyon.Error},
		{"Info color", flyon.Info, flyon.Info},
		{"Neutral color", flyon.Neutral, flyon.Neutral},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			autocomplete := NewAutocomplete().WithColor(tt.color)
			if autocomplete.color != tt.expected {
				t.Errorf("Expected color to be %v, got %v", tt.expected, autocomplete.color)
			}
		})
	}
}

func TestAutocompleteComponent_WithSize(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected flyon.Size
	}{
		{"Extra small size", flyon.SizeXS, flyon.SizeXS},
		{"Small size", flyon.SizeSmall, flyon.SizeSmall},
		{"Medium size", flyon.SizeMedium, flyon.SizeMedium},
		{"Large size", flyon.SizeLarge, flyon.SizeLarge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			autocomplete := NewAutocomplete().WithSize(tt.size)
			if autocomplete.size != tt.expected {
				t.Errorf("Expected size to be %v, got %v", tt.expected, autocomplete.size)
			}
		})
	}
}

func TestAutocompleteComponent_WithOptions(t *testing.T) {
	options := []string{"Option 1", "Option 2", "Option 3"}
	autocomplete := NewAutocomplete().WithOptions(options...)

	if len(autocomplete.options) != 3 {
		t.Errorf("Expected 3 options, got %d", len(autocomplete.options))
	}
	for i, option := range options {
		if autocomplete.options[i] != option {
			t.Errorf("Expected option %d to be '%s', got '%s'", i, option, autocomplete.options[i])
		}
	}
}

func TestAutocompleteComponent_WithClasses(t *testing.T) {
	autocomplete := NewAutocomplete().WithClasses("custom-class", "another-class")
	html := renderToStringAutocomplete(autocomplete)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got %s", html)
	}
}

func TestAutocompleteComponent_With(t *testing.T) {
	autocomplete := NewAutocomplete().With("data-test", "value")
	html := renderToStringAutocomplete(autocomplete)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestAutocompleteComponent_Render(t *testing.T) {
	autocomplete := NewAutocomplete()
	html := renderToStringAutocomplete(autocomplete)

	if !strings.Contains(html, "input") {
		t.Errorf("Expected HTML to contain input element, got %s", html)
	}
	if !strings.Contains(html, "input-bordered") {
		t.Errorf("Expected HTML to contain input-bordered class, got %s", html)
	}
	if !strings.Contains(html, "dropdown") {
		t.Errorf("Expected HTML to contain dropdown element, got %s", html)
	}
}

func TestAutocompleteComponent_RenderWithOptions(t *testing.T) {
	options := []string{"Apple", "Banana", "Cherry"}
	autocomplete := NewAutocomplete().WithOptions(options...)
	html := renderToStringAutocomplete(autocomplete)

	for _, option := range options {
		if !strings.Contains(html, option) {
			t.Errorf("Expected HTML to contain option '%s', got %s", option, html)
		}
	}
}

func TestAutocompleteComponent_RenderWithAllAttributes(t *testing.T) {
	autocomplete := NewAutocomplete().
		WithID("test-id").
		WithName("test-name").
		WithPlaceholder("Search...").
		WithValue("test-value").
		WithDisabled(true).
		WithColor(flyon.Success).
		WithSize(flyon.SizeLarge).
		WithOptions("Option 1", "Option 2").
		WithClasses("custom-class").
		With("data-test", "value")

	html := renderToStringAutocomplete(autocomplete)

	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected HTML to contain id attribute, got %s", html)
	}
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected HTML to contain name attribute, got %s", html)
	}
	if !strings.Contains(html, `placeholder="Search..."`) {
		t.Errorf("Expected HTML to contain placeholder attribute, got %s", html)
	}
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("Expected HTML to contain value attribute, got %s", html)
	}
	if !strings.Contains(html, "disabled") {
		t.Errorf("Expected HTML to contain disabled attribute, got %s", html)
	}
	if !strings.Contains(html, "input-success") {
		t.Errorf("Expected HTML to contain success color class, got %s", html)
	}
	if !strings.Contains(html, "input-lg") {
		t.Errorf("Expected HTML to contain large size class, got %s", html)
	}
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain custom class, got %s", html)
	}
	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
	if !strings.Contains(html, "Option 1") {
		t.Errorf("Expected HTML to contain Option 1, got %s", html)
	}
	if !strings.Contains(html, "Option 2") {
		t.Errorf("Expected HTML to contain Option 2, got %s", html)
	}
}

func TestAutocompleteComponent_Immutability(t *testing.T) {
	original := NewAutocomplete()
	modified := original.WithID("new-id")

	if original.id == modified.id {
		t.Error("Expected original component to remain unchanged")
	}
	if original.id != "" {
		t.Errorf("Expected original ID to be empty, got %s", original.id)
	}
	if modified.id != "new-id" {
		t.Errorf("Expected modified ID to be 'new-id', got %s", modified.id)
	}
}