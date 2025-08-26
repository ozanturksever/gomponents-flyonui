package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid()

	if grid == nil {
		t.Fatal("NewGrid() returned nil")
	}

	html := renderToHTML(grid)
	if !strings.Contains(html, "<div") {
		t.Errorf("Expected grid to render as div element, got: %s", html)
	}

	if !strings.Contains(html, "grid") {
		t.Errorf("Expected grid to have 'grid' class, got: %s", html)
	}
}

func TestGrid_With_Color(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary grid", flyon.Primary, "grid-primary"},
		{"Secondary grid", flyon.Secondary, "grid-secondary"},
		{"Success grid", flyon.Success, "grid-success"},
		{"Warning grid", flyon.Warning, "grid-warning"},
		{"Error grid", flyon.Error, "grid-error"},
		{"Info grid", flyon.Info, "grid-info"},
		{"Neutral grid", flyon.Neutral, "grid-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid().With(tt.color)
			html := renderToHTML(grid)

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
				t.Errorf("Expected grid to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestGrid_With_Size(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small grid", flyon.SizeXS, "grid-xs"},
		{"Small grid", flyon.SizeSmall, "grid-sm"},
		{"Medium grid", flyon.SizeMedium, "grid-md"},
		{"Large grid", flyon.SizeLarge, "grid-lg"},
		{"Extra large grid", flyon.SizeXL, "grid-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid().With(tt.size)
			html := renderToHTML(grid)

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
				t.Errorf("Expected grid to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestGrid_With_Variant(t *testing.T) {
	tests := []struct {
		name     string
		variant  flyon.Variant
		expected string
	}{
		{"Solid grid", flyon.VariantSolid, "grid-solid"},
		{"Outline grid", flyon.VariantOutline, "grid-outline"},
		{"Ghost grid", flyon.VariantGhost, "grid-ghost"},
		{"Soft grid", flyon.VariantSoft, "grid-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid().With(tt.variant)
			html := renderToHTML(grid)

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
				t.Errorf("Expected grid to have class '%s', got: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestGrid_With_CombinedModifiers(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []flyon.Modifier
		expected  []string
	}{
		{
			"Grid with size and color",
			[]flyon.Modifier{flyon.SizeLarge, flyon.Primary},
			[]string{"grid", "grid-lg", "grid-primary"},
		},
		{
			"Grid with variant and color",
			[]flyon.Modifier{flyon.VariantOutline, flyon.Secondary},
			[]string{"grid", "grid-outline", "grid-secondary"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid()
			for _, mod := range tt.modifiers {
				grid = grid.With(mod).(*GridComponent)
			}
			html := renderToHTML(grid)

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
					t.Errorf("Expected grid to have class '%s', got: %s", expectedClass, classAttr)
				}
			}
		})
	}
}

func TestGrid_HTMLAttributes(t *testing.T) {
	grid := NewGrid()
	html := renderToHTML(grid)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	divEl := findElement(doc, "div")
	if divEl == nil {
		t.Fatal("Expected div element not found")
	}

	if !hasAttribute(divEl, "class") {
		t.Errorf("Expected grid to have class attribute, got: %s", html)
	}

	classValue := getAttribute(divEl, "class")
	if !strings.Contains(classValue, "grid") {
		t.Errorf("Expected grid class to contain 'grid', got: %s", classValue)
	}
}

func TestGrid_ComponentInterface(t *testing.T) {
	grid := NewGrid()

	// Test that Grid implements flyon.Component
	var _ flyon.Component = grid

	// Test that Grid can be rendered
	html := renderToHTML(grid)
	if html == "" {
		t.Error("Expected grid to render non-empty HTML")
	}
}

func TestGrid_Immutability(t *testing.T) {
	original := NewGrid()
	modified := original.With(flyon.Primary)

	originalHTML := renderToHTML(original)
	modifiedHTML := renderToHTML(modified)

	if originalHTML == modifiedHTML {
		t.Error("Expected original and modified grids to be different")
	}

	// Check original grid
	originalDoc, err := parseHTML(originalHTML)
	if err != nil {
		t.Fatalf("Failed to parse original HTML: %v", err)
	}
	originalDiv := findElement(originalDoc, "div")
	if originalDiv == nil {
		t.Fatal("Expected div element not found in original")
	}
	originalClassAttr := getAttribute(originalDiv, "class")

	// Check modified grid
	modifiedDoc, err := parseHTML(modifiedHTML)
	if err != nil {
		t.Fatalf("Failed to parse modified HTML: %v", err)
	}
	modifiedDiv := findElement(modifiedDoc, "div")
	if modifiedDiv == nil {
		t.Fatal("Expected div element not found in modified")
	}
	modifiedClassAttr := getAttribute(modifiedDiv, "class")

	if hasClass(originalClassAttr, "grid-primary") {
		t.Error("Expected original grid to not have primary class after modification")
	}

	if !hasClass(modifiedClassAttr, "grid-primary") {
		t.Error("Expected modified grid to have primary class")
	}
}