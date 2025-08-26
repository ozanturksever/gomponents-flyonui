package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringToggle renders a ToggleComponent to a string for testing
func renderToStringToggle(component flyon.Component) string {
	var builder strings.Builder
	component.Render(&builder)
	return builder.String()
}

func TestNewToggle(t *testing.T) {
	toggle := NewToggle()

	if toggle == nil {
		t.Fatal("NewToggle() returned nil")
	}

	// Test default values
	if toggle.checked {
		t.Error("Expected default checked to be false")
	}
	if toggle.disabled {
		t.Error("Expected default disabled to be false")
	}
	if toggle.color != flyon.Primary {
		t.Errorf("Expected default color to be Primary, got %v", toggle.color)
	}
	if toggle.size != flyon.SizeMedium {
		t.Errorf("Expected default size to be SizeMedium, got %v", toggle.size)
	}
}

func TestToggleComponent_WithID(t *testing.T) {
	toggle := NewToggle().WithID("test-toggle")
	html := renderToStringToggle(toggle)

	if !strings.Contains(html, `id="test-toggle"`) {
		t.Errorf("Expected HTML to contain id='test-toggle', got: %s", html)
	}
}

func TestToggleComponent_WithName(t *testing.T) {
	toggle := NewToggle().WithName("toggle-name")
	html := renderToStringToggle(toggle)

	if !strings.Contains(html, `name="toggle-name"`) {
		t.Errorf("Expected HTML to contain name='toggle-name', got: %s", html)
	}
}

func TestToggleComponent_WithValue(t *testing.T) {
	toggle := NewToggle().WithValue("toggle-value")
	html := renderToStringToggle(toggle)

	if !strings.Contains(html, `value="toggle-value"`) {
		t.Errorf("Expected HTML to contain value='toggle-value', got: %s", html)
	}
}

func TestToggleComponent_WithChecked(t *testing.T) {
	tests := []struct {
		name     string
		checked  bool
		expected string
	}{
		{"Checked true", true, "checked"},
		{"Checked false", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggle := NewToggle().WithChecked(tt.checked)
			html := renderToStringToggle(toggle)

			if tt.expected != "" {
				if !strings.Contains(html, tt.expected) {
					t.Errorf("Expected HTML to contain '%s', got: %s", tt.expected, html)
				}
			} else {
				if strings.Contains(html, "checked") {
					t.Errorf("Expected HTML not to contain 'checked', got: %s", html)
				}
			}
		})
	}
}

func TestToggleComponent_WithDisabled(t *testing.T) {
	tests := []struct {
		name     string
		disabled bool
		expected string
	}{
		{"Disabled true", true, "disabled"},
		{"Disabled false", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggle := NewToggle().WithDisabled(tt.disabled)
			html := renderToStringToggle(toggle)

			if tt.expected != "" {
				if !strings.Contains(html, tt.expected) {
					t.Errorf("Expected HTML to contain '%s', got: %s", tt.expected, html)
				}
			} else {
				if strings.Contains(html, "disabled") {
					t.Errorf("Expected HTML not to contain 'disabled', got: %s", html)
				}
			}
		})
	}
}

func TestToggleComponent_WithColor(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary color", flyon.Primary, "toggle-primary"},
		{"Secondary color", flyon.Secondary, "toggle-secondary"},
		{"Success color", flyon.Success, "toggle-success"},
		{"Warning color", flyon.Warning, "toggle-warning"},
		{"Error color", flyon.Error, "toggle-error"},
		{"Info color", flyon.Info, "toggle-info"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggle := NewToggle().WithColor(tt.color)
			html := renderToStringToggle(toggle)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected HTML to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestToggleComponent_WithSize(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small size", flyon.SizeXS, "toggle-xs"},
		{"Small size", flyon.SizeSmall, "toggle-sm"},
		{"Medium size", flyon.SizeMedium, "toggle-md"},
		{"Large size", flyon.SizeLarge, "toggle-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggle := NewToggle().WithSize(tt.size)
			html := renderToStringToggle(toggle)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected HTML to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestToggleComponent_WithClasses(t *testing.T) {
	toggle := NewToggle().WithClasses("custom-class", "another-class")
	html := renderToStringToggle(toggle)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got: %s", html)
	}
}

func TestToggleComponent_With(t *testing.T) {
	toggle := NewToggle().With("data-testid", "toggle-test")
	html := renderToStringToggle(toggle)

	if !strings.Contains(html, `data-testid="toggle-test"`) {
		t.Errorf("Expected HTML to contain data-testid='toggle-test', got: %s", html)
	}
}

func TestToggleComponent_Render(t *testing.T) {
	toggle := NewToggle()
	html := renderToStringToggle(toggle)

	// Check for basic toggle structure
	if !strings.Contains(html, `<input`) {
		t.Errorf("Expected HTML to contain input element, got: %s", html)
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Errorf("Expected HTML to contain type='checkbox', got: %s", html)
	}
	if !strings.Contains(html, `class="toggle"`) {
		t.Errorf("Expected HTML to contain class='toggle', got: %s", html)
	}
}

func TestToggleComponent_RenderWithAllAttributes(t *testing.T) {
	toggle := NewToggle().
		WithID("test-toggle").
		WithName("toggle-name").
		WithValue("toggle-value").
		WithChecked(true).
		WithDisabled(true).
		WithColor(flyon.Success).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class").
		With("data-testid", "toggle-test")

	html := renderToStringToggle(toggle)

	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-toggle"`,
		`name="toggle-name"`,
		`value="toggle-value"`,
		`checked`,
		`disabled`,
		`toggle-success`,
		`toggle-lg`,
		`custom-class`,
		`data-testid="toggle-test"`,
	}

	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected HTML to contain '%s', got: %s", attr, html)
		}
	}
}

func TestToggleComponent_Immutability(t *testing.T) {
	original := NewToggle()
	modified := original.WithChecked(true).WithColor(flyon.Success)

	// Original should remain unchanged
	if original.checked {
		t.Error("Original toggle should not be checked")
	}
	if original.color != flyon.Primary {
		t.Error("Original toggle color should remain Primary")
	}

	// Modified should have new values
	if !modified.checked {
		t.Error("Modified toggle should be checked")
	}
	if modified.color != flyon.Success {
		t.Error("Modified toggle color should be Success")
	}
}