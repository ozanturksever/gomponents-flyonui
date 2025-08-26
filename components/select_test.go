package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringSelect renders a SelectComponent to a string for testing
func renderToStringSelect(component *SelectComponent) string {
	var builder strings.Builder
	component.Render(&builder)
	return builder.String()
}

func TestSelectComponent_WithID(t *testing.T) {
	select1 := NewSelect().WithID("test-id")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected HTML to contain id='test-id', got: %s", html)
	}
}

func TestSelectComponent_WithName(t *testing.T) {
	select1 := NewSelect().WithName("test-name")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected HTML to contain name='test-name', got: %s", html)
	}
}

func TestSelectComponent_WithValue(t *testing.T) {
	select1 := NewSelect().WithValue("test-value")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("Expected HTML to contain value='test-value', got: %s", html)
	}
}

func TestSelectComponent_WithDisabled(t *testing.T) {
	select1 := NewSelect().WithDisabled(true)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "disabled") {
		t.Errorf("Expected HTML to contain disabled attribute, got: %s", html)
	}
}

func TestSelectComponent_WithRequired(t *testing.T) {
	select1 := NewSelect().WithRequired(true)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "required") {
		t.Errorf("Expected HTML to contain required attribute, got: %s", html)
	}
}

func TestSelectComponent_WithMultiple(t *testing.T) {
	select1 := NewSelect().WithMultiple(true)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "multiple") {
		t.Errorf("Expected HTML to contain multiple attribute, got: %s", html)
	}
}

func TestSelectComponent_WithSize(t *testing.T) {
	select1 := NewSelect().WithSize(5)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `size="5"`) {
		t.Errorf("Expected HTML to contain size='5', got: %s", html)
	}
}

func TestSelectComponent_WithColor(t *testing.T) {
	select1 := NewSelect().WithColor(flyon.Secondary)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "select-secondary") {
		t.Errorf("Expected HTML to contain 'select-secondary' class, got: %s", html)
	}
}

func TestSelectComponent_WithCompSize(t *testing.T) {
	select1 := NewSelect().WithCompSize(flyon.SizeLarge)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "select-lg") {
		t.Errorf("Expected HTML to contain 'select-lg' class, got: %s", html)
	}
}

func TestSelectComponent_WithOption(t *testing.T) {
	select1 := NewSelect().WithOption("value1", "Label 1")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `<option value="value1">Label 1</option>`) {
		t.Errorf("Expected HTML to contain option with value='value1' and text 'Label 1', got: %s", html)
	}
}

func TestSelectComponent_WithSelectedOption(t *testing.T) {
	select1 := NewSelect().WithSelectedOption("value1", "Label 1")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `<option value="value1" selected>Label 1</option>`) {
		t.Errorf("Expected HTML to contain selected option with value='value1' and text 'Label 1', got: %s", html)
	}
}

func TestSelectComponent_WithDisabledOption(t *testing.T) {
	select1 := NewSelect().WithDisabledOption("value1", "Label 1")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `<option value="value1" disabled>Label 1</option>`) {
		t.Errorf("Expected HTML to contain disabled option with value='value1' and text 'Label 1', got: %s", html)
	}
}

func TestSelectComponent_WithOptions(t *testing.T) {
	options := []SelectOption{
		{Value: "value1", Label: "Label 1"},
		{Value: "value2", Label: "Label 2", Selected: true},
		{Value: "value3", Label: "Label 3", Disabled: true},
	}
	
	select1 := NewSelect().WithOptions(options)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, `<option value="value1">Label 1</option>`) {
		t.Errorf("Expected HTML to contain first option, got: %s", html)
	}
	if !strings.Contains(html, `<option value="value2" selected>Label 2</option>`) {
		t.Errorf("Expected HTML to contain selected second option, got: %s", html)
	}
	if !strings.Contains(html, `<option value="value3" disabled>Label 3</option>`) {
		t.Errorf("Expected HTML to contain disabled third option, got: %s", html)
	}
}

func TestSelectComponent_WithClasses(t *testing.T) {
	select1 := NewSelect().WithClasses("custom-class", "another-class")
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got: %s", html)
	}
}

func TestSelectComponent_With(t *testing.T) {
	select1 := NewSelect().With(flyon.Success, flyon.SizeSmall, "custom-modifier").(*SelectComponent)
	html := renderToStringSelect(select1)
	
	if !strings.Contains(html, "select-success") {
		t.Errorf("Expected HTML to contain 'select-success' class, got: %s", html)
	}
	if !strings.Contains(html, "select-sm") {
		t.Errorf("Expected HTML to contain 'select-sm' class, got: %s", html)
	}
	if !strings.Contains(html, "custom-modifier") {
		t.Errorf("Expected HTML to contain 'custom-modifier' class, got: %s", html)
	}
}

func TestSelectComponent_Render(t *testing.T) {
	select1 := NewSelect()
	html := renderToStringSelect(select1)
	
	// Check for basic select structure
	if !strings.Contains(html, "<select") {
		t.Errorf("Expected HTML to contain <select tag, got: %s", html)
	}
	if !strings.Contains(html, "</select>") {
		t.Errorf("Expected HTML to contain </select> tag, got: %s", html)
	}
	if !strings.Contains(html, "select") {
		t.Errorf("Expected HTML to contain 'select' class, got: %s", html)
	}
	if !strings.Contains(html, "select-bordered") {
		t.Errorf("Expected HTML to contain 'select-bordered' class, got: %s", html)
	}
}

func TestSelectComponent_RenderWithAllAttributes(t *testing.T) {
	select1 := NewSelect().
		WithID("test-select").
		WithName("test-name").
		WithValue("test-value").
		WithDisabled(true).
		WithRequired(true).
		WithMultiple(true).
		WithSize(3).
		WithColor(flyon.Secondary).
		WithCompSize(flyon.SizeLarge).
		WithOption("opt1", "Option 1").
		WithSelectedOption("opt2", "Option 2").
		WithClasses("custom-class")
	
	html := renderToStringSelect(select1)
	
	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-select"`,
		`name="test-name"`,
		`value="test-value"`,
		"disabled",
		"required",
		"multiple",
		`size="3"`,
		"select-secondary",
		"select-lg",
		"custom-class",
		`<option value="opt1">Option 1</option>`,
		`<option value="opt2" selected>Option 2</option>`,
	}
	
	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected HTML to contain '%s', got: %s", attr, html)
		}
	}
}

func TestSelectComponent_RenderNotDisabled(t *testing.T) {
	select1 := NewSelect().WithDisabled(false)
	html := renderToStringSelect(select1)
	
	if strings.Contains(html, "disabled") {
		t.Errorf("Expected HTML to not contain disabled attribute, got: %s", html)
	}
}

func TestSelectComponent_RenderNotRequired(t *testing.T) {
	select1 := NewSelect().WithRequired(false)
	html := renderToStringSelect(select1)
	
	if strings.Contains(html, "required") {
		t.Errorf("Expected HTML to not contain required attribute, got: %s", html)
	}
}

func TestSelectComponent_RenderNotMultiple(t *testing.T) {
	select1 := NewSelect().WithMultiple(false)
	html := renderToStringSelect(select1)
	
	if strings.Contains(html, "multiple") {
		t.Errorf("Expected HTML to not contain multiple attribute, got: %s", html)
	}
}

func TestSelectComponent_RenderDefaultColor(t *testing.T) {
	select1 := NewSelect()
	html := renderToStringSelect(select1)
	
	// Should not contain color class for primary (default)
	if strings.Contains(html, "select-primary") {
		t.Errorf("Expected HTML to not contain 'select-primary' class for default color, got: %s", html)
	}
}

func TestSelectComponent_RenderDefaultSize(t *testing.T) {
	select1 := NewSelect()
	html := renderToStringSelect(select1)
	
	// Should not contain size class for medium (default)
	if strings.Contains(html, "select-md") {
		t.Errorf("Expected HTML to not contain 'select-md' class for default size, got: %s", html)
	}
}

func TestSelectComponent_Immutability(t *testing.T) {
	original := NewSelect().WithOption("value1", "Label 1")
	modified := original.WithColor(flyon.Success).WithOption("value2", "Label 2")
	
	originalHTML := renderToStringSelect(original)
	modifiedHTML := renderToStringSelect(modified)
	
	// Original should not contain success color
	if strings.Contains(originalHTML, "select-success") {
		t.Errorf("Original component should not be modified, but contains 'select-success'")
	}
	
	// Original should not contain second option
	if strings.Contains(originalHTML, "value2") {
		t.Errorf("Original component should not be modified, but contains 'value2'")
	}
	
	// Modified should contain success color
	if !strings.Contains(modifiedHTML, "select-success") {
		t.Errorf("Modified component should contain 'select-success'")
	}
	
	// Modified should contain both options
	if !strings.Contains(modifiedHTML, "value1") {
		t.Errorf("Modified component should contain 'value1'")
	}
	if !strings.Contains(modifiedHTML, "value2") {
		t.Errorf("Modified component should contain 'value2'")
	}
}