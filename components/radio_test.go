package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringRadio is a helper function to render radio components to string
func renderToStringRadio(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestRadioComponent_WithID(t *testing.T) {
	radio := NewRadio().WithID("test-radio")
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, `id="test-radio"`) {
		t.Errorf("Expected radio to have id='test-radio', got: %s", html)
	}
}

func TestRadioComponent_WithName(t *testing.T) {
	radio := NewRadio().WithName("test-name")
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected radio to have name='test-name', got: %s", html)
	}
}

func TestRadioComponent_WithValue(t *testing.T) {
	radio := NewRadio().WithValue("test-value")
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("Expected radio to have value='test-value', got: %s", html)
	}
}

func TestRadioComponent_WithChecked(t *testing.T) {
	radio := NewRadio().WithChecked(true)
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, `checked`) {
		t.Errorf("Expected radio to be checked, got: %s", html)
	}
}

func TestRadioComponent_WithDisabled(t *testing.T) {
	radio := NewRadio().WithDisabled(true)
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, `disabled`) {
		t.Errorf("Expected radio to be disabled, got: %s", html)
	}
}

func TestRadioComponent_WithColor(t *testing.T) {
	radio := NewRadio().WithColor(flyon.Secondary)
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, "radio-secondary") {
		t.Errorf("Expected radio to have radio-secondary class, got: %s", html)
	}
}

func TestRadioComponent_WithSize(t *testing.T) {
	radio := NewRadio().WithSize(flyon.SizeLarge)
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, "radio-lg") {
		t.Errorf("Expected radio to have radio-lg class, got: %s", html)
	}
}

func TestRadioComponent_WithClasses(t *testing.T) {
	radio := NewRadio().WithClasses("custom-class", "another-class")
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected radio to have custom-class, got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected radio to have another-class, got: %s", html)
	}
}

func TestRadioComponent_With(t *testing.T) {
	radio := NewRadio().With(flyon.Success, flyon.SizeSmall, "custom-class")
	html := renderToStringRadio(radio)
	
	if !strings.Contains(html, "radio-success") {
		t.Errorf("Expected radio to have radio-success class, got: %s", html)
	}
	if !strings.Contains(html, "radio-sm") {
		t.Errorf("Expected radio to have radio-sm class, got: %s", html)
	}
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected radio to have custom-class, got: %s", html)
	}
}

func TestRadioComponent_Render(t *testing.T) {
	radio := NewRadio()
	html := renderToStringRadio(radio)
	
	// Check basic structure
	if !strings.Contains(html, `<input`) {
		t.Errorf("Expected radio to render as input element, got: %s", html)
	}
	if !strings.Contains(html, `type="radio"`) {
		t.Errorf("Expected radio to have type='radio', got: %s", html)
	}
	if !strings.Contains(html, `class="radio"`) {
		t.Errorf("Expected radio to have class='radio', got: %s", html)
	}
}

func TestRadioComponent_RenderWithAllAttributes(t *testing.T) {
	radio := NewRadio().
		WithID("test-id").
		WithName("test-name").
		WithValue("test-value").
		WithChecked(true).
		WithDisabled(true).
		WithColor(flyon.Secondary).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class")
	
	html := renderToStringRadio(radio)
	
	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-id"`,
		`name="test-name"`,
		`value="test-value"`,
		`checked`,
		`disabled`,
		`radio-secondary`,
		`radio-lg`,
		`custom-class`,
	}
	
	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected radio to contain %s, got: %s", attr, html)
		}
	}
}

func TestRadioComponent_RenderUnchecked(t *testing.T) {
	radio := NewRadio().WithChecked(false)
	html := renderToStringRadio(radio)
	
	if strings.Contains(html, `checked`) {
		t.Errorf("Expected radio to not be checked, got: %s", html)
	}
}

func TestRadioComponent_RenderEnabled(t *testing.T) {
	radio := NewRadio().WithDisabled(false)
	html := renderToStringRadio(radio)
	
	if strings.Contains(html, `disabled`) {
		t.Errorf("Expected radio to not be disabled, got: %s", html)
	}
}

func TestRadioComponent_RenderDefaultColor(t *testing.T) {
	radio := NewRadio()
	html := renderToStringRadio(radio)
	
	// Should not contain any color-specific classes
	colorClasses := []string{"radio-primary", "radio-secondary", "radio-success", "radio-error"}
	for _, colorClass := range colorClasses {
		if strings.Contains(html, colorClass) {
			t.Errorf("Expected radio to not have color class %s, got: %s", colorClass, html)
		}
	}
}

func TestRadioComponent_RenderDefaultSize(t *testing.T) {
	radio := NewRadio()
	html := renderToStringRadio(radio)
	
	// Should not contain any size-specific classes
	sizeClasses := []string{"radio-xs", "radio-sm", "radio-lg"}
	for _, sizeClass := range sizeClasses {
		if strings.Contains(html, sizeClass) {
			t.Errorf("Expected radio to not have size class %s, got: %s", sizeClass, html)
		}
	}
}

func TestRadioComponent_Immutability(t *testing.T) {
	original := NewRadio()
	modified := original.WithID("test-id").WithChecked(true)
	
	// Original should remain unchanged
	originalHTML := renderToStringRadio(original)
	modifiedHTML := renderToStringRadio(modified)
	
	if strings.Contains(originalHTML, `id="test-id"`) {
		t.Error("Original radio should not have been modified")
	}
	if strings.Contains(originalHTML, `checked`) {
		t.Error("Original radio should not have been modified")
	}
	
	if !strings.Contains(modifiedHTML, `id="test-id"`) {
		t.Error("Modified radio should have the new ID")
	}
	if !strings.Contains(modifiedHTML, `checked`) {
		t.Error("Modified radio should be checked")
	}
}