package components

import (
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"golang.org/x/net/html"
)

func TestSelect_BasicRendering(t *testing.T) {
	t.Run("renders basic select with default classes", func(t *testing.T) {
		select_ := NewSelect(
			h.Option(g.Attr("value", "1"), g.Text("Option 1")),
			h.Option(g.Attr("value", "2"), g.Text("Option 2")),
		)
		html := renderToHTML(select_)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		selectEl := findElement(doc, "select")
		if selectEl == nil {
			t.Fatal("Expected select element not found")
		}
		
		classAttr := getAttribute(selectEl, "class")
		if !hasClass(classAttr, "select") {
			t.Errorf("Expected 'select' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders select with options", func(t *testing.T) {
		select_ := NewSelect(
			h.Option(g.Attr("value", "apple"), g.Text("Apple")),
			h.Option(g.Attr("value", "banana"), g.Text("Banana")),
			h.Option(g.Attr("value", "cherry"), g.Text("Cherry")),
		)
		html := renderToHTML(select_)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		selectEl := findElement(doc, "select")
		if selectEl == nil {
			t.Fatal("Expected select element not found")
		}
		
		// Check that options are present
		options := findAllElements(selectEl, "option")
		if len(options) != 3 {
			t.Errorf("Expected 3 options, got %d", len(options))
		}
		
		// Check first option
		if value := getAttribute(options[0], "value"); value != "apple" {
			t.Errorf("Expected first option value='apple', got value='%s'", value)
		}
	})
}

func TestSelect_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary select", flyon.Primary, "select-primary"},
		{"Secondary select", flyon.Secondary, "select-secondary"},
		{"Success select", flyon.Success, "select-success"},
		{"Warning select", flyon.Warning, "select-warning"},
		{"Error select", flyon.Error, "select-error"},
		{"Info select", flyon.Info, "select-info"},
		{"Neutral select", flyon.Neutral, "select-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			select_ := NewSelect(h.Option(g.Text("Test"))).With(tt.color)
			html := renderToHTML(select_)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			selectEl := findElement(doc, "select")
			if selectEl == nil {
				t.Fatal("Expected select element not found")
			}
			
			classAttr := getAttribute(selectEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestSelect_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small select", flyon.SizeXS, "select-xs"},
		{"Small select", flyon.SizeSmall, "select-sm"},
		{"Medium select", flyon.SizeMedium, "select-md"},
		{"Large select", flyon.SizeLarge, "select-lg"},
		{"Extra large select", flyon.SizeXL, "select-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			select_ := NewSelect(h.Option(g.Text("Test"))).With(tt.size)
			html := renderToHTML(select_)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			selectEl := findElement(doc, "select")
			if selectEl == nil {
				t.Fatal("Expected select element not found")
			}
			
			classAttr := getAttribute(selectEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestSelect_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid select", flyon.VariantSolid, "select-solid"},
		{"Outline select", flyon.VariantOutline, "select-outline"},
		{"Ghost select", flyon.VariantGhost, "select-ghost"},
		{"Soft select", flyon.VariantSoft, "select-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			select_ := NewSelect(h.Option(g.Text("Test"))).With(tt.variant)
			html := renderToHTML(select_)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			selectEl := findElement(doc, "select")
			if selectEl == nil {
				t.Fatal("Expected select element not found")
			}
			
			classAttr := getAttribute(selectEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestSelect_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		select_ := NewSelect(h.Option(g.Text("Test"))).With(flyon.Primary, flyon.SizeLarge, flyon.VariantOutline)
		html := renderToHTML(select_)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		selectEl := findElement(doc, "select")
		if selectEl == nil {
			t.Fatal("Expected select element not found")
		}
		
		classAttr := getAttribute(selectEl, "class")
		expectedClasses := []string{"select", "select-primary", "select-lg", "select-outline"}
		for _, class := range expectedClasses {
			if !hasClass(classAttr, class) {
				t.Errorf("Expected class '%s', got classes: %s", class, classAttr)
			}
		}
	})
}

func TestSelect_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		select_ := NewSelect(
			h.ID("test-select"),
			h.Disabled(),
			h.Name("fruits"),
			h.Option(g.Attr("value", "apple"), g.Text("Apple")),
		)
		html := renderToHTML(select_)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		selectEl := findElement(doc, "select")
		if selectEl == nil {
			t.Fatal("Expected select element not found")
		}
		
		if id := getAttribute(selectEl, "id"); id != "test-select" {
			t.Errorf("Expected id='test-select', got id='%s'", id)
		}
		
		if !hasAttribute(selectEl, "disabled") {
			t.Error("Expected disabled attribute to be present")
		}
		
		if name := getAttribute(selectEl, "name"); name != "fruits" {
			t.Errorf("Expected name='fruits', got name='%s'", name)
		}
	})
}

func TestSelect_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = (*SelectComponent)(nil)
		
		select_ := NewSelect(h.Option(g.Attr("value", "test"), g.Text("Test")))
		html := renderToHTML(select_)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		selectEl := findElement(doc, "select")
		if selectEl == nil {
			t.Fatal("Expected select element not found")
		}
		
		option := findElement(selectEl, "option")
		if option == nil {
			t.Fatal("Expected option element not found")
		}
		
		if value := getAttribute(option, "value"); value != "test" {
			t.Errorf("Expected option value='test', got value='%s'", value)
		}
	})
}

// Helper function to find all elements with a specific tag
func findAllElements(n *html.Node, tag string) []*html.Node {
	var elements []*html.Node
	findAllElementsRecursive(n, tag, &elements)
	return elements
}

func findAllElementsRecursive(n *html.Node, tag string, elements *[]*html.Node) {
	if n.Type == html.ElementNode && n.Data == tag {
		*elements = append(*elements, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAllElementsRecursive(c, tag, elements)
	}
}