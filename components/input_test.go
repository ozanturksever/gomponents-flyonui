package components

import (
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	h "maragu.dev/gomponents/html"
)

func TestInput_BasicRendering(t *testing.T) {
	t.Run("renders basic input with default classes", func(t *testing.T) {
		input := NewInput()
		html := renderToHTML(input)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		inputEl := findElement(doc, "input")
		if inputEl == nil {
			t.Fatal("Expected input element not found")
		}
		
		classAttr := getAttribute(inputEl, "class")
		if !hasClass(classAttr, "input") {
			t.Errorf("Expected 'input' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders input with placeholder and type", func(t *testing.T) {
		input := NewInput(h.Placeholder("Enter text"), h.Type("email"))
		html := renderToHTML(input)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		inputEl := findElement(doc, "input")
		if inputEl == nil {
			t.Fatal("Expected input element not found")
		}
		
		if placeholder := getAttribute(inputEl, "placeholder"); placeholder != "Enter text" {
			t.Errorf("Expected placeholder='Enter text', got placeholder='%s'", placeholder)
		}
		
		if inputType := getAttribute(inputEl, "type"); inputType != "email" {
			t.Errorf("Expected type='email', got type='%s'", inputType)
		}
	})
}

func TestInput_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary input", flyon.Primary, "input-primary"},
		{"Secondary input", flyon.Secondary, "input-secondary"},
		{"Success input", flyon.Success, "input-success"},
		{"Warning input", flyon.Warning, "input-warning"},
		{"Error input", flyon.Error, "input-error"},
		{"Info input", flyon.Info, "input-info"},
		{"Neutral input", flyon.Neutral, "input-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput().With(tt.color)
			html := renderToHTML(input)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			inputEl := findElement(doc, "input")
			if inputEl == nil {
				t.Fatal("Expected input element not found")
			}
			
			classAttr := getAttribute(inputEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestInput_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small input", flyon.SizeXS, "input-xs"},
		{"Small input", flyon.SizeSmall, "input-sm"},
		{"Medium input", flyon.SizeMedium, "input-md"},
		{"Large input", flyon.SizeLarge, "input-lg"},
		{"Extra large input", flyon.SizeXL, "input-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput().With(tt.size)
			html := renderToHTML(input)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			inputEl := findElement(doc, "input")
			if inputEl == nil {
				t.Fatal("Expected input element not found")
			}
			
			classAttr := getAttribute(inputEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestInput_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid input", flyon.VariantSolid, "input-solid"},
		{"Outline input", flyon.VariantOutline, "input-outline"},
		{"Ghost input", flyon.VariantGhost, "input-ghost"},
		{"Soft input", flyon.VariantSoft, "input-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput().With(tt.variant)
			html := renderToHTML(input)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			inputEl := findElement(doc, "input")
			if inputEl == nil {
				t.Fatal("Expected input element not found")
			}
			
			classAttr := getAttribute(inputEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestInput_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		input := NewInput().With(flyon.Primary, flyon.SizeLarge, flyon.VariantOutline)
		html := renderToHTML(input)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		inputEl := findElement(doc, "input")
		if inputEl == nil {
			t.Fatal("Expected input element not found")
		}
		
		classAttr := getAttribute(inputEl, "class")
		expectedClasses := []string{"input", "input-primary", "input-lg", "input-outline"}
		for _, class := range expectedClasses {
			if !hasClass(classAttr, class) {
				t.Errorf("Expected class '%s', got classes: %s", class, classAttr)
			}
		}
	})
}

func TestInput_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		input := NewInput(h.ID("test-input"), h.Disabled(), h.Placeholder("Test"))
		html := renderToHTML(input)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		inputEl := findElement(doc, "input")
		if inputEl == nil {
			t.Fatal("Expected input element not found")
		}
		
		if id := getAttribute(inputEl, "id"); id != "test-input" {
			t.Errorf("Expected id='test-input', got id='%s'", id)
		}
		
		if !hasAttribute(inputEl, "disabled") {
			t.Error("Expected disabled attribute to be present")
		}
		
		if placeholder := getAttribute(inputEl, "placeholder"); placeholder != "Test" {
			t.Errorf("Expected placeholder='Test', got placeholder='%s'", placeholder)
		}
	})
}

func TestInput_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = (*InputComponent)(nil)
		
		input := NewInput(h.Placeholder("Test"))
		html := renderToHTML(input)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		inputEl := findElement(doc, "input")
		if inputEl == nil {
			t.Fatal("Expected input element not found")
		}
		
		if placeholder := getAttribute(inputEl, "placeholder"); placeholder != "Test" {
			t.Errorf("Expected placeholder='Test', got placeholder='%s'", placeholder)
		}
	})
}