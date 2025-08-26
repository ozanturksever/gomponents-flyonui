package components

import (
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	h "maragu.dev/gomponents/html"
)

func TestTextarea_BasicRendering(t *testing.T) {
	tests := []struct {
		name     string
		textarea *TextareaComponent
		want     string
	}{
		{
			name:     "basic textarea",
			textarea: NewTextarea(),
			want:     "textarea",
		},
		{
			name:     "textarea with placeholder",
			textarea: NewTextarea(h.Placeholder("Enter text...")),
			want:     "textarea",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			htmlStr := renderToHTML(tt.textarea)
			doc, err := parseHTML(htmlStr)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			textarea := findElement(doc, "textarea")
			if textarea == nil {
				t.Fatalf("Expected textarea element, got none")
			}

			// Check for base classes
			classAttr := getAttribute(textarea, "class")
			if !hasClass(classAttr, "textarea") {
				t.Errorf("Expected textarea to have 'textarea' class, got: %s", classAttr)
			}
		})
	}
}

func TestTextarea_ColorModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifier  flyon.Modifier
		expected  string
	}{
		{"Primary textarea", flyon.Primary, "textarea-primary"},
		{"Secondary textarea", flyon.Secondary, "textarea-secondary"},
		{"Success textarea", flyon.Success, "textarea-success"},
		{"Warning textarea", flyon.Warning, "textarea-warning"},
		{"Error textarea", flyon.Error, "textarea-error"},
		{"Info textarea", flyon.Info, "textarea-info"},
		{"Neutral textarea", flyon.Neutral, "textarea-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textarea := NewTextarea().With(tt.modifier)
			htmlStr := renderToHTML(textarea)
			doc, err := parseHTML(htmlStr)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			textareaEl := findElement(doc, "textarea")
			if textareaEl == nil {
				t.Fatalf("Expected textarea element, got none")
			}

			classAttr := getAttribute(textareaEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected textarea to have '%s' class, got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestTextarea_SizeModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifier  flyon.Modifier
		expected  string
	}{
		{"Extra small textarea", flyon.SizeXS, "textarea-xs"},
		{"Small textarea", flyon.SizeSmall, "textarea-sm"},
		{"Medium textarea", flyon.SizeMedium, "textarea-md"},
		{"Large textarea", flyon.SizeLarge, "textarea-lg"},
		{"Extra large textarea", flyon.SizeXL, "textarea-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textarea := NewTextarea().With(tt.modifier)
			htmlStr := renderToHTML(textarea)
			doc, err := parseHTML(htmlStr)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			textareaEl := findElement(doc, "textarea")
			if textareaEl == nil {
				t.Fatalf("Expected textarea element, got none")
			}

			classAttr := getAttribute(textareaEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected textarea to have '%s' class, got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestTextarea_VariantModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifier  flyon.Modifier
		expected  string
	}{
		{"Solid textarea", flyon.VariantSolid, "textarea-solid"},
		{"Outline textarea", flyon.VariantOutline, "textarea-outline"},
		{"Ghost textarea", flyon.VariantGhost, "textarea-ghost"},
		{"Soft textarea", flyon.VariantSoft, "textarea-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textarea := NewTextarea().With(tt.modifier)
			htmlStr := renderToHTML(textarea)
			doc, err := parseHTML(htmlStr)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			textareaEl := findElement(doc, "textarea")
			if textareaEl == nil {
				t.Fatalf("Expected textarea element, got none")
			}

			classAttr := getAttribute(textareaEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected textarea to have '%s' class, got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestTextarea_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		textarea := NewTextarea().With(flyon.Primary, flyon.SizeLarge, flyon.VariantOutline)
		htmlStr := renderToHTML(textarea)
		doc, err := parseHTML(htmlStr)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		textareaEl := findElement(doc, "textarea")
		if textareaEl == nil {
			t.Fatalf("Expected textarea element, got none")
		}

		classAttr := getAttribute(textareaEl, "class")
		expectedClasses := []string{"textarea", "textarea-primary", "textarea-lg", "textarea-outline"}
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected textarea to have '%s' class, got: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestTextarea_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		textarea := NewTextarea(
			h.ID("my-textarea"),
			h.Name("description"),
			h.Placeholder("Enter description..."),
			h.Rows("5"),
			h.Cols("40"),
		)
		htmlStr := renderToHTML(textarea)
		doc, err := parseHTML(htmlStr)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		textareaEl := findElement(doc, "textarea")
		if textareaEl == nil {
			t.Fatalf("Expected textarea element, got none")
		}

		if id := getAttribute(textareaEl, "id"); id != "my-textarea" {
			t.Errorf("Expected id='my-textarea', got: %s", id)
		}
		if name := getAttribute(textareaEl, "name"); name != "description" {
			t.Errorf("Expected name='description', got: %s", name)
		}
		if placeholder := getAttribute(textareaEl, "placeholder"); placeholder != "Enter description..." {
			t.Errorf("Expected placeholder='Enter description...', got: %s", placeholder)
		}
		if rows := getAttribute(textareaEl, "rows"); rows != "5" {
			t.Errorf("Expected rows='5', got: %s", rows)
		}
		if cols := getAttribute(textareaEl, "cols"); cols != "40" {
			t.Errorf("Expected cols='40', got: %s", cols)
		}
	})
}

func TestTextarea_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = NewTextarea()

		textarea := NewTextarea().With(flyon.Primary)
		htmlStr := renderToHTML(textarea)
		doc, err := parseHTML(htmlStr)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		textareaEl := findElement(doc, "textarea")
		if textareaEl == nil {
			t.Fatalf("Expected textarea element, got none")
		}

		classAttr := getAttribute(textareaEl, "class")
		if !hasClass(classAttr, "textarea-primary") {
			t.Errorf("Expected textarea to have 'textarea-primary' class, got: %s", classAttr)
		}
	})
}