package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"golang.org/x/net/html"
	"maragu.dev/gomponents"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Test helper to render gomponents.Node to HTML string
func renderToHTML(node gomponents.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

// Test helper to parse HTML and check for specific attributes/classes
func parseHTML(htmlStr string) (*html.Node, error) {
	return html.Parse(strings.NewReader(htmlStr))
}

// Test helper to find element with specific tag
func findElement(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findElement(c, tag); result != nil {
			return result
		}
	}
	return nil
}

// Test helper to get attribute value
func getAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

// Test helper to check if class list contains specific class
func hasClass(classAttr, className string) bool {
	classes := strings.Fields(classAttr)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

// Test helper to check if boolean attribute exists
func hasAttribute(n *html.Node, key string) bool {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return true
		}
	}
	return false
}

func TestButton_BasicRendering(t *testing.T) {
	t.Run("renders basic button with default classes", func(t *testing.T) {
		btn := NewButton(g.Text("Click me"))
		html := renderToHTML(btn)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		buttonEl := findElement(doc, "button")
		if buttonEl == nil {
			t.Fatal("Expected button element not found")
		}
		
		classAttr := getAttribute(buttonEl, "class")
		if !hasClass(classAttr, "btn") {
			t.Errorf("Expected 'btn' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders button with text content", func(t *testing.T) {
		btn := NewButton(g.Text("Submit Form"))
		html := renderToHTML(btn)
		
		if !strings.Contains(html, "Submit Form") {
			t.Errorf("Expected button to contain 'Submit Form', got: %s", html)
		}
	})

	t.Run("renders button with multiple children", func(t *testing.T) {
		btn := NewButton(
			h.Span(g.Text("Icon")),
			g.Text(" Submit"),
		)
		html := renderToHTML(btn)
		
		if !strings.Contains(html, "<span>Icon</span>") {
			t.Errorf("Expected span with Icon, got: %s", html)
		}
		if !strings.Contains(html, " Submit") {
			t.Errorf("Expected ' Submit' text, got: %s", html)
		}
	})
}

func TestButton_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary button", flyon.Primary, "btn-primary"},
		{"Secondary button", flyon.Secondary, "btn-secondary"},
		{"Success button", flyon.Success, "btn-success"},
		{"Warning button", flyon.Warning, "btn-warning"},
		{"Error button", flyon.Error, "btn-error"},
		{"Info button", flyon.Info, "btn-info"},
		{"Neutral button", flyon.Neutral, "btn-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := NewButton(g.Text("Click me")).With(tt.color)
			html := renderToHTML(btn)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			buttonEl := findElement(doc, "button")
			if buttonEl == nil {
				t.Fatal("Expected button element not found")
			}
			
			classAttr := getAttribute(buttonEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestButton_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small button", flyon.SizeXS, "btn-xs"},
		{"Small button", flyon.SizeSmall, "btn-sm"},
		{"Medium button", flyon.SizeMedium, "btn-md"},
		{"Large button", flyon.SizeLarge, "btn-lg"},
		{"Extra large button", flyon.SizeXL, "btn-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := NewButton(g.Text("Click me")).With(tt.size)
			html := renderToHTML(btn)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			buttonEl := findElement(doc, "button")
			if buttonEl == nil {
				t.Fatal("Expected button element not found")
			}
			
			classAttr := getAttribute(buttonEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestButton_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid button", flyon.VariantSolid, "btn-solid"},
		{"Outline button", flyon.VariantOutline, "btn-outline"},
		{"Ghost button", flyon.VariantGhost, "btn-ghost"},
		{"Soft button", flyon.VariantSoft, "btn-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := NewButton(g.Text("Click me")).With(tt.variant)
			html := renderToHTML(btn)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			buttonEl := findElement(doc, "button")
			if buttonEl == nil {
				t.Fatal("Expected button element not found")
			}
			
			classAttr := getAttribute(buttonEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestButton_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		btn := NewButton(g.Text("Click me")).With(
			flyon.Primary,
			flyon.SizeLarge,
			flyon.VariantOutline,
		)
		html := renderToHTML(btn)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		buttonEl := findElement(doc, "button")
		if buttonEl == nil {
			t.Fatal("Expected button element not found")
		}
		
		classAttr := getAttribute(buttonEl, "class")
		expectedClasses := []string{"btn", "btn-primary", "btn-lg", "btn-outline"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestButton_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		btn := NewButton(
			h.ID("submit-btn"),
			h.Type("submit"),
			h.Disabled(),
			g.Text("Submit"),
		)
		html := renderToHTML(btn)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		buttonEl := findElement(doc, "button")
		if buttonEl == nil {
			t.Fatal("Expected button element not found")
		}
		
		if id := getAttribute(buttonEl, "id"); id != "submit-btn" {
			t.Errorf("Expected id='submit-btn', got id='%s'", id)
		}
		
		if btnType := getAttribute(buttonEl, "type"); btnType != "submit" {
			t.Errorf("Expected type='submit', got type='%s'", btnType)
		}
		
		if !hasAttribute(buttonEl, "disabled") {
			t.Error("Expected disabled attribute to be present")
		}
	})
}

func TestButton_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		btn := NewButton(g.Text("Click me"))
		
		// Test that it implements flyon.Component
		var _ flyon.Component = btn
		
		// Test that it implements gomponents.Node
		var _ gomponents.Node = btn
		
		// Test that With returns a new instance
		newBtn := btn.With(flyon.Primary)
		if btn == newBtn {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}