package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestNewFlex(t *testing.T) {
	flex := NewFlex()

	if flex == nil {
		t.Fatal("NewFlex() returned nil")
	}

	html := renderToHTML(flex)
	if !strings.Contains(html, "<div") {
		t.Errorf("Expected flex to render as div element, got: %s", html)
	}

	if !strings.Contains(html, "flex") {
		t.Errorf("Expected flex to have 'flex' class, got: %s", html)
	}
}

func TestFlex_With_Color(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary flex", flyon.Primary, "flex-primary"},
		{"Secondary flex", flyon.Secondary, "flex-secondary"},
		{"Success flex", flyon.Success, "flex-success"},
		{"Warning flex", flyon.Warning, "flex-warning"},
		{"Error flex", flyon.Error, "flex-error"},
		{"Info flex", flyon.Info, "flex-info"},
		{"Neutral flex", flyon.Neutral, "flex-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flex := NewFlex().With(tt.color)
			html := renderToHTML(flex)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			divEl := findElement(doc, "div")
			if divEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(divEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected flex to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestFlex_With_Size(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small flex", flyon.SizeXS, "flex-xs"},
		{"Small flex", flyon.SizeSmall, "flex-sm"},
		{"Medium flex", flyon.SizeMedium, "flex-md"},
		{"Large flex", flyon.SizeLarge, "flex-lg"},
		{"Extra large flex", flyon.SizeXL, "flex-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flex := NewFlex().With(tt.size)
			html := renderToHTML(flex)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			divEl := findElement(doc, "div")
			if divEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(divEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected flex to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestFlex_With_Variant(t *testing.T) {
	tests := []struct {
		name     string
		variant  flyon.Variant
		expected string
	}{
		{"Solid flex", flyon.VariantSolid, "flex-solid"},
		{"Outline flex", flyon.VariantOutline, "flex-outline"},
		{"Ghost flex", flyon.VariantGhost, "flex-ghost"},
		{"Soft flex", flyon.VariantSoft, "flex-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flex := NewFlex().With(tt.variant)
			html := renderToHTML(flex)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			divEl := findElement(doc, "div")
			if divEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(divEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected flex to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestFlex_With_CombinedModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []flyon.Modifier
		expected  []string
	}{
		{
			"Flex with size and color",
			[]flyon.Modifier{flyon.SizeLarge, flyon.Primary},
			[]string{"flex", "flex-lg", "flex-primary"},
		},
		{
			"Flex with variant and color",
			[]flyon.Modifier{flyon.VariantOutline, flyon.Secondary},
			[]string{"flex", "flex-outline", "flex-secondary"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flex := NewFlex()
			for _, mod := range tt.modifiers {
				flex = flex.With(mod).(*FlexComponent)
			}
			html := renderToHTML(flex)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			divEl := findElement(doc, "div")
			if divEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(divEl, "class")
			for _, expectedClass := range tt.expected {
				if !hasClass(classAttr, expectedClass) {
					t.Errorf("Expected flex to have class '%s', got: %s", expectedClass, classAttr)
				}
			}
		})
	}
}

func TestFlex_HTMLAttributes(t *testing.T) {
	flex := NewFlex()
	html := renderToHTML(flex)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	divEl := findElement(doc, "div")
	if divEl == nil {
		t.Fatal("Expected div element not found")
	}

	if !hasAttribute(divEl, "class") {
		t.Errorf("Expected flex to have class attribute, got: %s", html)
	}

	classValue := getAttribute(divEl, "class")
	if !strings.Contains(classValue, "flex") {
		t.Errorf("Expected flex class to contain 'flex', got: %s", classValue)
	}
}

func TestFlex_ComponentInterface(t *testing.T) {
	flex := NewFlex()

	// Test that Flex implements flyon.Component
	var _ flyon.Component = flex

	// Test that Flex can be rendered
	html := renderToHTML(flex)
	if html == "" {
		t.Error("Expected flex to render non-empty HTML")
	}
}

func TestFlex_Immutability(t *testing.T) {
	original := NewFlex()
	modified := original.With(flyon.Primary)

	originalHTML := renderToHTML(original)
	modifiedHTML := renderToHTML(modified)

	if originalHTML == modifiedHTML {
		t.Error("Expected original and modified flexes to be different")
	}

	// Check original flex
	originalDoc, err := parseHTML(originalHTML)
	if err != nil {
		t.Fatalf("Failed to parse original HTML: %v", err)
	}
	originalDiv := findElement(originalDoc, "div")
	if originalDiv == nil {
		t.Fatal("Expected div element not found in original")
	}
	originalClassAttr := getAttribute(originalDiv, "class")

	// Check modified flex
	modifiedDoc, err := parseHTML(modifiedHTML)
	if err != nil {
		t.Fatalf("Failed to parse modified HTML: %v", err)
	}
	modifiedDiv := findElement(modifiedDoc, "div")
	if modifiedDiv == nil {
		t.Fatal("Expected div element not found in modified")
	}
	modifiedClassAttr := getAttribute(modifiedDiv, "class")

	if hasClass(originalClassAttr, "flex-primary") {
		t.Error("Expected original flex to not have primary class after modification")
	}

	if !hasClass(modifiedClassAttr, "flex-primary") {
		t.Error("Expected modified flex to have primary class")
	}
}