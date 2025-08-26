package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringCheckbox is a helper function to render checkbox components to string
func renderToStringCheckbox(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestCheckboxComponent_WithID(t *testing.T) {
	checkbox := NewCheckbox().WithID("test-checkbox")
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, `id="test-checkbox"`) {
		t.Errorf("Expected checkbox to have id='test-checkbox', got: %s", html)
	}
}

func TestCheckboxComponent_WithName(t *testing.T) {
	checkbox := NewCheckbox().WithName("test-name")
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected checkbox to have name='test-name', got: %s", html)
	}
}

func TestCheckboxComponent_WithValue(t *testing.T) {
	checkbox := NewCheckbox().WithValue("test-value")
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("Expected checkbox to have value='test-value', got: %s", html)
	}
}

func TestCheckboxComponent_WithChecked(t *testing.T) {
	checkbox := NewCheckbox().WithChecked(true)
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, `checked`) {
		t.Errorf("Expected checkbox to be checked, got: %s", html)
	}
}

func TestCheckboxComponent_WithDisabled(t *testing.T) {
	checkbox := NewCheckbox().WithDisabled(true)
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, `disabled`) {
		t.Errorf("Expected checkbox to be disabled, got: %s", html)
	}
}

func TestCheckboxComponent_WithColor(t *testing.T) {
	checkbox := NewCheckbox().WithColor(flyon.Secondary)
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, "checkbox-secondary") {
		t.Errorf("Expected checkbox to have checkbox-secondary class, got: %s", html)
	}
}

func TestCheckboxComponent_WithSize(t *testing.T) {
	checkbox := NewCheckbox().WithSize(flyon.SizeLarge)
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, "checkbox-lg") {
		t.Errorf("Expected checkbox to have checkbox-lg class, got: %s", html)
	}
}

func TestCheckboxComponent_WithClasses(t *testing.T) {
	checkbox := NewCheckbox().WithClasses("custom-class", "another-class")
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected checkbox to have custom-class, got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected checkbox to have another-class, got: %s", html)
	}
}

func TestCheckboxComponent_With(t *testing.T) {
	checkbox := NewCheckbox().With(flyon.Success, flyon.SizeSmall, "custom-class")
	html := renderToStringCheckbox(checkbox)
	
	if !strings.Contains(html, "checkbox-success") {
		t.Errorf("Expected checkbox to have checkbox-success class, got: %s", html)
	}
	if !strings.Contains(html, "checkbox-sm") {
		t.Errorf("Expected checkbox to have checkbox-sm class, got: %s", html)
	}
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected checkbox to have custom-class, got: %s", html)
	}
}

func TestCheckboxComponent_Render(t *testing.T) {
	checkbox := NewCheckbox()
	html := renderToStringCheckbox(checkbox)
	
	// Check basic structure
	if !strings.Contains(html, `<input`) {
		t.Errorf("Expected checkbox to render as input element, got: %s", html)
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Errorf("Expected checkbox to have type='checkbox', got: %s", html)
	}
	if !strings.Contains(html, `class="checkbox"`) {
		t.Errorf("Expected checkbox to have class='checkbox', got: %s", html)
	}
}

func TestCheckboxComponent_RenderWithAllAttributes(t *testing.T) {
	checkbox := NewCheckbox().
		WithID("test-id").
		WithName("test-name").
		WithValue("test-value").
		WithChecked(true).
		WithDisabled(true).
		WithColor(flyon.Secondary).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class")
	
	html := renderToStringCheckbox(checkbox)
	
	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-id"`,
		`name="test-name"`,
		`value="test-value"`,
		`checked`,
		`disabled`,
		`checkbox-secondary`,
		`checkbox-lg`,
		`custom-class`,
	}
	
	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected checkbox to contain %s, got: %s", attr, html)
		}
	}
}

func TestCheckboxComponent_RenderUnchecked(t *testing.T) {
	checkbox := NewCheckbox().WithChecked(false)
	html := renderToStringCheckbox(checkbox)
	
	if strings.Contains(html, `checked`) {
		t.Errorf("Expected checkbox to not be checked, got: %s", html)
	}
}

func TestCheckboxComponent_RenderEnabled(t *testing.T) {
	checkbox := NewCheckbox().WithDisabled(false)
	html := renderToStringCheckbox(checkbox)
	
	if strings.Contains(html, `disabled`) {
		t.Errorf("Expected checkbox to not be disabled, got: %s", html)
	}
}

func TestCheckboxComponent_RenderDefaultColor(t *testing.T) {
	checkbox := NewCheckbox()
	html := renderToStringCheckbox(checkbox)
	
	// Should not contain any color-specific classes
	colorClasses := []string{"checkbox-primary", "checkbox-secondary", "checkbox-success", "checkbox-error"}
	for _, colorClass := range colorClasses {
		if strings.Contains(html, colorClass) {
			t.Errorf("Expected checkbox to not have color class %s, got: %s", colorClass, html)
		}
	}
}

func TestCheckboxComponent_RenderDefaultSize(t *testing.T) {
	checkbox := NewCheckbox()
	html := renderToStringCheckbox(checkbox)
	
	// Should not contain any size-specific classes
	sizeClasses := []string{"checkbox-xs", "checkbox-sm", "checkbox-lg"}
	for _, sizeClass := range sizeClasses {
		if strings.Contains(html, sizeClass) {
			t.Errorf("Expected checkbox to not have size class %s, got: %s", sizeClass, html)
		}
	}
}

func TestCheckboxComponent_Immutability(t *testing.T) {
	original := NewCheckbox()
	modified := original.WithID("test-id").WithChecked(true)
	
	// Original should remain unchanged
	originalHTML := renderToStringCheckbox(original)
	modifiedHTML := renderToStringCheckbox(modified)
	
	if strings.Contains(originalHTML, `id="test-id"`) {
		t.Error("Original checkbox should not have been modified")
	}
	if strings.Contains(originalHTML, `checked`) {
		t.Error("Original checkbox should not have been modified")
	}
	
	if !strings.Contains(modifiedHTML, `id="test-id"`) {
		t.Error("Modified checkbox should have the new ID")
	}
	if !strings.Contains(modifiedHTML, `checked`) {
		t.Error("Modified checkbox should be checked")
	}
}