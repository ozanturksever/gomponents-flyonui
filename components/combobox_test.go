//go:build !js && !wasm

package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func renderToStringCombobox(component *ComboboxComponent) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestComboboxComponent_Defaults(t *testing.T) {
	combobox := NewCombobox()

	if combobox.id != "" {
		t.Errorf("Expected empty ID, got %s", combobox.id)
	}
	if combobox.name != "" {
		t.Errorf("Expected empty name, got %s", combobox.name)
	}
	if combobox.placeholder != "" {
		t.Errorf("Expected empty placeholder, got %s", combobox.placeholder)
	}
	if combobox.value != "" {
		t.Errorf("Expected empty value, got %s", combobox.value)
	}
	if combobox.disabled {
		t.Errorf("Expected disabled to be false, got %t", combobox.disabled)
	}
	if combobox.color != flyon.Primary {
		t.Errorf("Expected default color, got %s", combobox.color)
	}
	if combobox.size != flyon.SizeMedium {
		t.Errorf("Expected default size, got %s", combobox.size)
	}
	if len(combobox.options) != 0 {
		t.Errorf("Expected empty options, got %d", len(combobox.options))
	}
	if len(combobox.classes) != 0 {
		t.Errorf("Expected empty classes, got %d", len(combobox.classes))
	}
	if len(combobox.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %d", len(combobox.attributes))
	}
}

func TestComboboxComponent_WithID(t *testing.T) {
	combobox := NewCombobox().WithID("test-id")

	if combobox.id != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got %s", combobox.id)
	}
}

func TestComboboxComponent_WithName(t *testing.T) {
	combobox := NewCombobox().WithName("test-name")

	if combobox.name != "test-name" {
		t.Errorf("Expected name to be 'test-name', got %s", combobox.name)
	}
}

func TestComboboxComponent_WithPlaceholder(t *testing.T) {
	combobox := NewCombobox().WithPlaceholder("Select option...")

	if combobox.placeholder != "Select option..." {
		t.Errorf("Expected placeholder to be 'Select option...', got %s", combobox.placeholder)
	}
}

func TestComboboxComponent_WithValue(t *testing.T) {
	combobox := NewCombobox().WithValue("test-value")

	if combobox.value != "test-value" {
		t.Errorf("Expected value to be 'test-value', got %s", combobox.value)
	}
}

func TestComboboxComponent_WithDisabled(t *testing.T) {
	combobox := NewCombobox().WithDisabled(true)

	if !combobox.disabled {
		t.Errorf("Expected disabled to be true, got %t", combobox.disabled)
	}
}

func TestComboboxComponent_WithColor(t *testing.T) {
	tests := []struct {
		name  string
		color flyon.Color
	}{
		{"Primary color", flyon.Primary},
		{"Secondary color", flyon.Secondary},
		{"Success color", flyon.Success},
		{"Warning color", flyon.Warning},
		{"Error color", flyon.Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combobox := NewCombobox().WithColor(tt.color)
			if combobox.color != tt.color {
				t.Errorf("Expected color to be %s, got %s", tt.color, combobox.color)
			}
		})
	}
}

func TestComboboxComponent_WithSize(t *testing.T) {
	tests := []struct {
		name string
		size flyon.Size
	}{
		{"Extra small size", flyon.SizeXS},
		{"Small size", flyon.SizeSmall},
		{"Medium size", flyon.SizeMedium},
		{"Large size", flyon.SizeLarge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combobox := NewCombobox().WithSize(tt.size)
			if combobox.size != tt.size {
				t.Errorf("Expected size to be %s, got %s", tt.size, combobox.size)
			}
		})
	}
}

func TestComboboxComponent_WithOptions(t *testing.T) {
	options := []ComboboxOption{
		{Value: "option1", Label: "Option 1"},
		{Value: "option2", Label: "Option 2"},
	}
	combobox := NewCombobox().WithOptions(options)

	if len(combobox.options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(combobox.options))
	}
	if combobox.options[0].Value != "option1" {
		t.Errorf("Expected first option value to be 'option1', got %s", combobox.options[0].Value)
	}
	if combobox.options[0].Label != "Option 1" {
		t.Errorf("Expected first option label to be 'Option 1', got %s", combobox.options[0].Label)
	}
}

func TestComboboxComponent_WithClasses(t *testing.T) {
	combobox := NewCombobox().WithClasses("custom-class", "another-class")

	if len(combobox.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(combobox.classes))
	}
	if combobox.classes[0] != "custom-class" {
		t.Errorf("Expected first class to be 'custom-class', got %s", combobox.classes[0])
	}
	if combobox.classes[1] != "another-class" {
		t.Errorf("Expected second class to be 'another-class', got %s", combobox.classes[1])
	}
}

func TestComboboxComponent_With(t *testing.T) {
	combobox := NewCombobox().WithAttribute("data-test", "value")
	html := renderToStringCombobox(combobox)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestComboboxComponent_Render(t *testing.T) {
	combobox := NewCombobox()
	html := renderToStringCombobox(combobox)

	if !strings.Contains(html, "<div") {
		t.Errorf("Expected HTML to contain div element, got %s", html)
	}
	if !strings.Contains(html, "dropdown") {
		t.Errorf("Expected HTML to contain dropdown class, got %s", html)
	}
	if !strings.Contains(html, "<input") {
		t.Errorf("Expected HTML to contain input element, got %s", html)
	}
}

func TestComboboxComponent_RenderWithOptions(t *testing.T) {
	options := []ComboboxOption{
		{Value: "option1", Label: "Option 1"},
		{Value: "option2", Label: "Option 2"},
	}
	combobox := NewCombobox().WithOptions(options)
	html := renderToStringCombobox(combobox)

	if !strings.Contains(html, "Option 1") {
		t.Errorf("Expected HTML to contain 'Option 1', got %s", html)
	}
	if !strings.Contains(html, "Option 2") {
		t.Errorf("Expected HTML to contain 'Option 2', got %s", html)
	}
	if !strings.Contains(html, `value="option1"`) {
		t.Errorf("Expected HTML to contain option1 value, got %s", html)
	}
	if !strings.Contains(html, `value="option2"`) {
		t.Errorf("Expected HTML to contain option2 value, got %s", html)
	}
}

func TestComboboxComponent_RenderWithAllAttributes(t *testing.T) {
	options := []ComboboxOption{
		{Value: "option1", Label: "Option 1"},
	}
	combobox := NewCombobox().
		WithID("test-id").
		WithName("test-name").
		WithPlaceholder("Select option...").
		WithValue("option1").
		WithDisabled(true).
		WithColor(flyon.Primary).
		WithSize(flyon.SizeLarge).
		WithOptions(options).
		WithClasses("custom-class").
		WithAttribute("data-test", "value")

	html := renderToStringCombobox(combobox)

	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected HTML to contain ID, got %s", html)
	}
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected HTML to contain name, got %s", html)
	}
	if !strings.Contains(html, `placeholder="Select option..."`) {
		t.Errorf("Expected HTML to contain placeholder, got %s", html)
	}
	if !strings.Contains(html, `value="option1"`) {
		t.Errorf("Expected HTML to contain value, got %s", html)
	}
	if !strings.Contains(html, "disabled") {
		t.Errorf("Expected HTML to contain disabled attribute, got %s", html)
	}
	if !strings.Contains(html, "input-primary") {
		t.Errorf("Expected HTML to contain primary color class, got %s", html)
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
		t.Errorf("Expected HTML to contain option label, got %s", html)
	}
}

func TestComboboxComponent_Immutability(t *testing.T) {
	original := NewCombobox()
	modified := original.WithID("test-id")

	if original.id != "" {
		t.Errorf("Expected original ID to remain empty, got %s", original.id)
	}
	if modified.id != "test-id" {
		t.Errorf("Expected modified ID to be 'test-id', got %s", modified.id)
	}

	// Test that they are different instances
	if original == modified {
		t.Error("Expected original and modified to be different instances")
	}
}