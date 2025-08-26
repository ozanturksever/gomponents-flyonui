package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// renderToStringTextarea is a helper function to render TextareaComponent to string
func renderToStringTextarea(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestTextareaComponent_WithID(t *testing.T) {
	textarea := NewTextarea().WithID("test-id")
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected textarea to have id='test-id', got: %s", html)
	}
}

func TestTextareaComponent_WithName(t *testing.T) {
	textarea := NewTextarea().WithName("test-name")
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected textarea to have name='test-name', got: %s", html)
	}
}

func TestTextareaComponent_WithValue(t *testing.T) {
	textarea := NewTextarea().WithValue("test content")
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, "test content") {
		t.Errorf("Expected textarea to contain 'test content', got: %s", html)
	}
}

func TestTextareaComponent_WithPlaceholder(t *testing.T) {
	textarea := NewTextarea().WithPlaceholder("Enter text...")
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `placeholder="Enter text..."`) {
		t.Errorf("Expected textarea to have placeholder='Enter text...', got: %s", html)
	}
}

func TestTextareaComponent_WithRows(t *testing.T) {
	textarea := NewTextarea().WithRows(5)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `rows="5"`) {
		t.Errorf("Expected textarea to have rows='5', got: %s", html)
	}
}

func TestTextareaComponent_WithCols(t *testing.T) {
	textarea := NewTextarea().WithCols(40)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `cols="40"`) {
		t.Errorf("Expected textarea to have cols='40', got: %s", html)
	}
}

func TestTextareaComponent_WithDisabled(t *testing.T) {
	textarea := NewTextarea().WithDisabled(true)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `disabled="disabled"`) {
		t.Errorf("Expected textarea to have disabled='disabled', got: %s", html)
	}
}

func TestTextareaComponent_WithReadonly(t *testing.T) {
	textarea := NewTextarea().WithReadonly(true)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `readonly="readonly"`) {
		t.Errorf("Expected textarea to have readonly='readonly', got: %s", html)
	}
}

func TestTextareaComponent_WithRequired(t *testing.T) {
	textarea := NewTextarea().WithRequired(true)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, `required="required"`) {
		t.Errorf("Expected textarea to have required='required', got: %s", html)
	}
}

func TestTextareaComponent_WithColor(t *testing.T) {
	textarea := NewTextarea().WithColor(flyon.Primary)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, "textarea-primary") {
		t.Errorf("Expected textarea to have 'textarea-primary' class, got: %s", html)
	}
}

func TestTextareaComponent_WithSize(t *testing.T) {
	textarea := NewTextarea().WithSize(flyon.SizeLarge)
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, "textarea-lg") {
		t.Errorf("Expected textarea to have 'textarea-lg' class, got: %s", html)
	}
}

func TestTextareaComponent_WithClasses(t *testing.T) {
	textarea := NewTextarea().WithClasses("custom-class", "another-class")
	html := renderToStringTextarea(textarea)
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected textarea to have 'custom-class' class, got: %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected textarea to have 'another-class' class, got: %s", html)
	}
}

func TestTextareaComponent_With(t *testing.T) {
	textarea1 := NewTextarea().With(flyon.Success, flyon.SizeSmall, "custom-modifier").(*TextareaComponent)
	html := renderToStringTextarea(textarea1)
	if !strings.Contains(html, "textarea-success") {
		t.Errorf("Expected textarea to have 'textarea-success' class, got: %s", html)
	}
	if !strings.Contains(html, "textarea-sm") {
		t.Errorf("Expected textarea to have 'textarea-sm' class, got: %s", html)
	}
	if !strings.Contains(html, "custom-modifier") {
		t.Errorf("Expected textarea to have 'custom-modifier' class, got: %s", html)
	}
}

func TestTextareaComponent_Render(t *testing.T) {
	textarea := NewTextarea()
	html := renderToStringTextarea(textarea)
	
	// Check basic structure
	if !strings.Contains(html, "<textarea") {
		t.Errorf("Expected textarea element, got: %s", html)
	}
	if !strings.Contains(html, "</textarea>") {
		t.Errorf("Expected closing textarea tag, got: %s", html)
	}
	if !strings.Contains(html, "class=\"textarea\"") {
		t.Errorf("Expected textarea to have 'textarea' class, got: %s", html)
	}
}

func TestTextareaComponent_RenderWithAllAttributes(t *testing.T) {
	textarea := NewTextarea().
		WithID("test-textarea").
		WithName("description").
		WithValue("Initial content").
		WithPlaceholder("Enter description...").
		WithRows(5).
		WithCols(40).
		WithDisabled(true).
		WithReadonly(true).
		WithRequired(true).
		WithColor(flyon.Primary).
		WithSize(flyon.SizeLarge).
		WithClasses("custom-class")
	
	html := renderToStringTextarea(textarea)
	
	// Check all attributes are present
	expectedAttributes := []string{
		`id="test-textarea"`,
		`name="description"`,
		`placeholder="Enter description..."`,
		`rows="5"`,
		`cols="40"`,
		`disabled="disabled"`,
		`readonly="readonly"`,
		`required="required"`,
		"textarea-primary",
		"textarea-lg",
		"custom-class",
		"Initial content",
	}
	
	for _, attr := range expectedAttributes {
		if !strings.Contains(html, attr) {
			t.Errorf("Expected textarea to contain '%s', got: %s", attr, html)
		}
	}
}

func TestTextareaComponent_RenderNotDisabled(t *testing.T) {
	textarea := NewTextarea().WithDisabled(false)
	html := renderToStringTextarea(textarea)
	if strings.Contains(html, "disabled") {
		t.Errorf("Expected textarea to not have disabled attribute, got: %s", html)
	}
}

func TestTextareaComponent_RenderNotReadonly(t *testing.T) {
	textarea := NewTextarea().WithReadonly(false)
	html := renderToStringTextarea(textarea)
	if strings.Contains(html, "readonly") {
		t.Errorf("Expected textarea to not have readonly attribute, got: %s", html)
	}
}

func TestTextareaComponent_RenderNotRequired(t *testing.T) {
	textarea := NewTextarea().WithRequired(false)
	html := renderToStringTextarea(textarea)
	if strings.Contains(html, "required") {
		t.Errorf("Expected textarea to not have required attribute, got: %s", html)
	}
}

func TestTextareaComponent_RenderDefaultColor(t *testing.T) {
	textarea := NewTextarea() // Default color
	html := renderToStringTextarea(textarea)
	if strings.Contains(html, "textarea-default") {
		t.Errorf("Expected textarea to not have 'textarea-default' class, got: %s", html)
	}
}

func TestTextareaComponent_RenderDefaultSize(t *testing.T) {
	textarea := NewTextarea() // Default size
	html := renderToStringTextarea(textarea)
	if strings.Contains(html, "textarea-default") {
		t.Errorf("Expected textarea to not have 'textarea-default' class, got: %s", html)
	}
}

func TestTextareaComponent_Immutability(t *testing.T) {
	original := NewTextarea()
	modified := original.WithID("new-id")
	
	// Check that original is unchanged
	originalHTML := renderToStringTextarea(original)
	modifiedHTML := renderToStringTextarea(modified)
	
	if strings.Contains(originalHTML, "new-id") {
		t.Errorf("Original textarea should not be modified, got: %s", originalHTML)
	}
	if !strings.Contains(modifiedHTML, "new-id") {
		t.Errorf("Modified textarea should contain new-id, got: %s", modifiedHTML)
	}
	
	// Test color immutability
	originalColor := NewTextarea()
	modifiedColor := originalColor.WithColor(flyon.Primary)
	
	originalColorHTML := renderToStringTextarea(originalColor)
	modifiedColorHTML := renderToStringTextarea(modifiedColor)
	
	if strings.Contains(originalColorHTML, "textarea-primary") {
		t.Errorf("Original textarea should not have primary color, got: %s", originalColorHTML)
	}
	if !strings.Contains(modifiedColorHTML, "textarea-primary") {
		t.Errorf("Modified textarea should have primary color, got: %s", modifiedColorHTML)
	}
}