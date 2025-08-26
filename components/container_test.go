package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if container == nil {
		t.Fatal("NewContainer() returned nil")
	}

	html := renderToHTML(container)
	if !strings.Contains(html, "<div") {
		t.Errorf("Expected container to render as div element, got: %s", html)
	}

	if !strings.Contains(html, "container") {
		t.Errorf("Expected container to have 'container' class, got: %s", html)
	}
}

func TestContainer_With_Color(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary container", flyon.Primary, "container-primary"},
		{"Secondary container", flyon.Secondary, "container-secondary"},
		{"Success container", flyon.Success, "container-success"},
		{"Warning container", flyon.Warning, "container-warning"},
		{"Error container", flyon.Error, "container-error"},
		{"Info container", flyon.Info, "container-info"},
		{"Neutral container", flyon.Neutral, "container-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := NewContainer().With(tt.color)
			html := renderToHTML(container)

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
				t.Errorf("Expected container to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestContainer_With_Size(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small container", flyon.SizeXS, "container-xs"},
		{"Small container", flyon.SizeSmall, "container-sm"},
		{"Medium container", flyon.SizeMedium, "container-md"},
		{"Large container", flyon.SizeLarge, "container-lg"},
		{"Extra large container", flyon.SizeXL, "container-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := NewContainer().With(tt.size)
			html := renderToHTML(container)

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
				t.Errorf("Expected container to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestContainer_With_Variant(t *testing.T) {
	tests := []struct {
		name     string
		variant  flyon.Variant
		expected string
	}{
		{"Solid container", flyon.VariantSolid, "container-solid"},
		{"Outline container", flyon.VariantOutline, "container-outline"},
		{"Ghost container", flyon.VariantGhost, "container-ghost"},
		{"Soft container", flyon.VariantSoft, "container-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := NewContainer().With(tt.variant)
			html := renderToHTML(container)

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
				t.Errorf("Expected container to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestContainer_With_CombinedModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []flyon.Modifier
		expected  []string
	}{
		{
			"Container with size and color",
			[]flyon.Modifier{flyon.SizeLarge, flyon.Primary},
			[]string{"container", "container-lg", "container-primary"},
		},
		{
			"Container with variant and color",
			[]flyon.Modifier{flyon.VariantOutline, flyon.Secondary},
			[]string{"container", "container-outline", "container-secondary"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := NewContainer()
			for _, mod := range tt.modifiers {
				container = container.With(mod).(*ContainerComponent)
			}
			html := renderToHTML(container)

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
					t.Errorf("Expected container to have class '%s', got: %s", expectedClass, classAttr)
				}
			}
		})
	}
}

func TestContainer_HTMLAttributes(t *testing.T) {
	container := NewContainer()
	html := renderToHTML(container)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	divEl := findElement(doc, "div")
	if divEl == nil {
		t.Fatal("Expected div element not found")
	}

	if !hasAttribute(divEl, "class") {
		t.Errorf("Expected container to have class attribute, got: %s", html)
	}

	classValue := getAttribute(divEl, "class")
	if !strings.Contains(classValue, "container") {
		t.Errorf("Expected container class to contain 'container', got: %s", classValue)
	}
}

func TestContainer_ComponentInterface(t *testing.T) {
	container := NewContainer()

	// Test that Container implements flyon.Component
	var _ flyon.Component = container

	// Test that Container can be rendered
	html := renderToHTML(container)
	if html == "" {
		t.Error("Expected container to render non-empty HTML")
	}
}

func TestContainer_Immutability(t *testing.T) {
	original := NewContainer()
	modified := original.With(flyon.Primary)

	originalHTML := renderToHTML(original)
	modifiedHTML := renderToHTML(modified)

	if originalHTML == modifiedHTML {
		t.Error("Expected original and modified containers to be different")
	}

	// Check original container
	originalDoc, err := parseHTML(originalHTML)
	if err != nil {
		t.Fatalf("Failed to parse original HTML: %v", err)
	}
	originalDiv := findElement(originalDoc, "div")
	if originalDiv == nil {
		t.Fatal("Expected div element not found in original")
	}
	originalClassAttr := getAttribute(originalDiv, "class")

	// Check modified container
	modifiedDoc, err := parseHTML(modifiedHTML)
	if err != nil {
		t.Fatalf("Failed to parse modified HTML: %v", err)
	}
	modifiedDiv := findElement(modifiedDoc, "div")
	if modifiedDiv == nil {
		t.Fatal("Expected div element not found in modified")
	}
	modifiedClassAttr := getAttribute(modifiedDiv, "class")

	if hasClass(originalClassAttr, "container-primary") {
		t.Error("Expected original container to not have primary class after modification")
	}

	if !hasClass(modifiedClassAttr, "container-primary") {
		t.Error("Expected modified container to have primary class")
	}
}