package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestCard_BasicRendering(t *testing.T) {
	t.Run("renders basic card with default classes", func(t *testing.T) {
		card := NewCard(g.Text("Card content"))
		html := renderToHTML(card)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		cardEl := findElement(doc, "div")
		if cardEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(cardEl, "class")
		if !hasClass(classAttr, "card") {
			t.Errorf("Expected 'card' class, got classes: %s", classAttr)
		}
		
		if !strings.Contains(html, "Card content") {
			t.Errorf("Expected 'Card content' text, got: %s", html)
		}
	})
	
	t.Run("renders card with multiple children", func(t *testing.T) {
		card := NewCard(
			h.H2(g.Text("Card Title")),
			h.P(g.Text("Card description")),
			h.Button(g.Text("Action")),
		)
		html := renderToHTML(card)
		
		if !strings.Contains(html, "<h2>Card Title</h2>") {
			t.Errorf("Expected h2 with Card Title, got: %s", html)
		}
		if !strings.Contains(html, "<p>Card description</p>") {
			t.Errorf("Expected p with Card description, got: %s", html)
		}
		if !strings.Contains(html, "<button>Action</button>") {
			t.Errorf("Expected button with Action, got: %s", html)
		}
	})
}

func TestCard_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary card", flyon.Primary, "card-primary"},
		{"Secondary card", flyon.Secondary, "card-secondary"},
		{"Success card", flyon.Success, "card-success"},
		{"Warning card", flyon.Warning, "card-warning"},
		{"Error card", flyon.Error, "card-error"},
		{"Info card", flyon.Info, "card-info"},
		{"Neutral card", flyon.Neutral, "card-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard(g.Text("Content")).With(tt.color)
			html := renderToHTML(card)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			cardEl := findElement(doc, "div")
			if cardEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(cardEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestCard_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small card", flyon.SizeXS, "card-xs"},
		{"Small card", flyon.SizeSmall, "card-sm"},
		{"Medium card", flyon.SizeMedium, "card-md"},
		{"Large card", flyon.SizeLarge, "card-lg"},
		{"Extra large card", flyon.SizeXL, "card-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard(g.Text("Content")).With(tt.size)
			html := renderToHTML(card)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			cardEl := findElement(doc, "div")
			if cardEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(cardEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestCard_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid card", flyon.VariantSolid, "card-solid"},
		{"Outline card", flyon.VariantOutline, "card-outline"},
		{"Ghost card", flyon.VariantGhost, "card-ghost"},
		{"Soft card", flyon.VariantSoft, "card-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard(g.Text("Content")).With(tt.variant)
			html := renderToHTML(card)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			cardEl := findElement(doc, "div")
			if cardEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(cardEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestCard_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		card := NewCard(g.Text("Content")).With(
			flyon.Primary,
			flyon.SizeLarge,
			flyon.VariantOutline,
		)
		html := renderToHTML(card)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		cardEl := findElement(doc, "div")
		if cardEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(cardEl, "class")
		expectedClasses := []string{"card", "card-primary", "card-lg", "card-outline"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestCard_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		card := NewCard(
			h.ID("main-card"),
			h.DataAttr("testid", "card-component"),
			g.Text("Card content"),
		)
		html := renderToHTML(card)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		cardEl := findElement(doc, "div")
		if cardEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		if id := getAttribute(cardEl, "id"); id != "main-card" {
			t.Errorf("Expected id='main-card', got id='%s'", id)
		}
		
		if testid := getAttribute(cardEl, "data-testid"); testid != "card-component" {
			t.Errorf("Expected data-testid='card-component', got data-testid='%s'", testid)
		}
	})
}

func TestCard_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		card := NewCard(g.Text("Content"))
		
		// Test that it implements flyon.Component
		var _ flyon.Component = card
		
		// Test that it implements gomponents.Node
		var _ g.Node = card
		
		// Test that With returns a new instance
		newCard := card.With(flyon.Primary)
		if card == newCard {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}