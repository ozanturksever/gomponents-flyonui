//go:build js && wasm

package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestDivider_BasicRendering(t *testing.T) {
	t.Run("renders basic divider with default classes", func(t *testing.T) {
		divider := NewDivider()
		html := renderToHTML(divider)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		dividerEl := findElement(doc, "div")
		if dividerEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(dividerEl, "class")
		if !hasClass(classAttr, "divider") {
			t.Errorf("Expected 'divider' class, got classes: %s", classAttr)
		}
	})
	
	t.Run("renders divider with text content", func(t *testing.T) {
		divider := NewDivider(g.Text("OR"))
		html := renderToHTML(divider)
		
		if !strings.Contains(html, "OR") {
			t.Errorf("Expected divider to contain 'OR', got: %s", html)
		}
	})
	
	t.Run("renders divider with multiple children", func(t *testing.T) {
		divider := NewDivider(
			h.Span(g.Text("Section")),
			g.Text(" Break"),
		)
		html := renderToHTML(divider)
		
		if !strings.Contains(html, "<span>Section</span>") {
			t.Errorf("Expected span with Section, got: %s", html)
		}
		if !strings.Contains(html, " Break") {
			t.Errorf("Expected ' Break' text, got: %s", html)
		}
	})
}

func TestDivider_ColorModifiers(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary color", flyon.Primary, "divider-primary"},
		{"Secondary color", flyon.Secondary, "divider-secondary"},
		{"Success color", flyon.Success, "divider-success"},
		{"Warning color", flyon.Warning, "divider-warning"},
		{"Error color", flyon.Error, "divider-error"},
		{"Info color", flyon.Info, "divider-info"},
		{"Neutral color", flyon.Neutral, "divider-neutral"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			divider := NewDivider().With(tt.color)
			html := renderToHTML(divider)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			dividerEl := findElement(doc, "div")
			if dividerEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(dividerEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
			if !hasClass(classAttr, "divider") {
				t.Errorf("Expected 'divider' base class, got classes: %s", classAttr)
			}
		})
	}
}

func TestDivider_OrientationModifiers(t *testing.T) {
	tests := []struct {
		name        string
		orientation string
		expected    string
	}{
		{"Horizontal orientation", "horizontal", "divider-horizontal"},
		{"Vertical orientation", "vertical", "divider-vertical"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			divider := NewDivider().WithOrientation(tt.orientation)
			html := renderToHTML(divider)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			dividerEl := findElement(doc, "div")
			if dividerEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(dividerEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestDivider_CombinedModifiers(t *testing.T) {
	t.Run("applies multiple modifiers correctly", func(t *testing.T) {
		divider := NewDivider(g.Text("Section")).WithOrientation("vertical").With(
			flyon.Primary,
			"custom-class",
		)
		
		html := renderToHTML(divider)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		dividerEl := findElement(doc, "div")
		if dividerEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(dividerEl, "class")
		expectedClasses := []string{"divider", "divider-primary", "custom-class", "divider-vertical"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestDivider_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		divider := NewDivider(g.Text("Content")).With(
			h.ID("test-divider"),
			g.Attr("data-test", "divider-component"),
			h.Title("Section Divider"),
		)
		
		html := renderToHTML(divider)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		dividerEl := findElement(doc, "div")
		if dividerEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		if id := getAttribute(dividerEl, "id"); id != "test-divider" {
			t.Errorf("Expected id='test-divider', got id='%s'", id)
		}
		
		if testAttr := getAttribute(dividerEl, "data-test"); testAttr != "divider-component" {
			t.Errorf("Expected data-test='divider-component', got data-test='%s'", testAttr)
		}
		
		if title := getAttribute(dividerEl, "title"); title != "Section Divider" {
			t.Errorf("Expected title='Section Divider', got title='%s'", title)
		}
	})
}

func TestDivider_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = NewDivider(g.Text("Test"))
		
		divider := NewDivider(g.Text("Interface Test"))
		html := renderToHTML(divider)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		dividerEl := findElement(doc, "div")
		if dividerEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		if dividerEl.FirstChild == nil || dividerEl.FirstChild.Data != "Interface Test" {
			t.Errorf("Expected text content 'Interface Test', got: %v", dividerEl.FirstChild)
		}
	})
}