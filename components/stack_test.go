package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if stack == nil {
		t.Fatal("NewStack() returned nil")
	}

	html := renderToHTML(stack)
	if !strings.Contains(html, "<div") {
		t.Errorf("Expected stack to render as div element, got: %s", html)
	}

	if !strings.Contains(html, "stack") {
		t.Errorf("Expected stack to have 'stack' class, got: %s", html)
	}
}

func TestStack_With_Color(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary stack", flyon.Primary, "stack-primary"},
		{"Secondary stack", flyon.Secondary, "stack-secondary"},
		{"Success stack", flyon.Success, "stack-success"},
		{"Warning stack", flyon.Warning, "stack-warning"},
		{"Error stack", flyon.Error, "stack-error"},
		{"Info stack", flyon.Info, "stack-info"},
		{"Neutral stack", flyon.Neutral, "stack-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack().With(tt.color)
			html := renderToHTML(stack)

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
				t.Errorf("Expected stack to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestStack_With_Size(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small stack", flyon.SizeXS, "stack-xs"},
		{"Small stack", flyon.SizeSmall, "stack-sm"},
		{"Medium stack", flyon.SizeMedium, "stack-md"},
		{"Large stack", flyon.SizeLarge, "stack-lg"},
		{"Extra large stack", flyon.SizeXL, "stack-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack().With(tt.size)
			html := renderToHTML(stack)

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
				t.Errorf("Expected stack to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestStack_With_Variant(t *testing.T) {
	tests := []struct {
		name     string
		variant  flyon.Variant
		expected string
	}{
		{"Solid stack", flyon.VariantSolid, "stack-solid"},
		{"Outline stack", flyon.VariantOutline, "stack-outline"},
		{"Ghost stack", flyon.VariantGhost, "stack-ghost"},
		{"Soft stack", flyon.VariantSoft, "stack-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack().With(tt.variant)
			html := renderToHTML(stack)

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
				t.Errorf("Expected stack to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestStack_With_CombinedModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []flyon.Modifier
		expected  []string
	}{
		{
			"Stack with size and color",
			[]flyon.Modifier{flyon.SizeLarge, flyon.Primary},
			[]string{"stack", "stack-lg", "stack-primary"},
		},
		{
			"Stack with variant and color",
			[]flyon.Modifier{flyon.VariantOutline, flyon.Secondary},
			[]string{"stack", "stack-outline", "stack-secondary"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack()
			for _, mod := range tt.modifiers {
				stack = stack.With(mod).(*StackComponent)
			}
			html := renderToHTML(stack)

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
					t.Errorf("Expected stack to have class '%s', got: %s", expectedClass, classAttr)
				}
			}
		})
	}
}

func TestStack_HTMLAttributes(t *testing.T) {
	stack := NewStack()
	html := renderToHTML(stack)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	divEl := findElement(doc, "div")
	if divEl == nil {
		t.Fatal("Expected div element not found")
	}

	if !hasAttribute(divEl, "class") {
		t.Errorf("Expected stack to have class attribute, got: %s", html)
	}

	classValue := getAttribute(divEl, "class")
	if !strings.Contains(classValue, "stack") {
		t.Errorf("Expected stack class to contain 'stack', got: %s", classValue)
	}
}

func TestStack_ComponentInterface(t *testing.T) {
	stack := NewStack()

	// Test that Stack implements flyon.Component
	var _ flyon.Component = stack

	// Test that Stack can be rendered
	html := renderToHTML(stack)
	if html == "" {
		t.Error("Expected stack to render non-empty HTML")
	}
}

func TestStack_Immutability(t *testing.T) {
	original := NewStack()
	modified := original.With(flyon.Primary)

	originalHTML := renderToHTML(original)
	modifiedHTML := renderToHTML(modified)

	if originalHTML == modifiedHTML {
		t.Error("Expected original and modified stacks to be different")
	}

	// Check original stack
	originalDoc, err := parseHTML(originalHTML)
	if err != nil {
		t.Fatalf("Failed to parse original HTML: %v", err)
	}
	originalDiv := findElement(originalDoc, "div")
	if originalDiv == nil {
		t.Fatal("Expected div element not found in original")
	}
	originalClassAttr := getAttribute(originalDiv, "class")

	// Check modified stack
	modifiedDoc, err := parseHTML(modifiedHTML)
	if err != nil {
		t.Fatalf("Failed to parse modified HTML: %v", err)
	}
	modifiedDiv := findElement(modifiedDoc, "div")
	if modifiedDiv == nil {
		t.Fatal("Expected div element not found in modified")
	}
	modifiedClassAttr := getAttribute(modifiedDiv, "class")

	if hasClass(originalClassAttr, "stack-primary") {
		t.Error("Expected original stack to not have primary class after modification")
	}

	if !hasClass(modifiedClassAttr, "stack-primary") {
		t.Error("Expected modified stack to have primary class")
	}
}