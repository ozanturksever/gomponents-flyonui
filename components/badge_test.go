package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestBadge_BasicRendering(t *testing.T) {
	t.Run("renders basic badge with default classes", func(t *testing.T) {
		badge := NewBadge(g.Text("New"))
		html := renderToHTML(badge)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		badgeEl := findElement(doc, "span")
		if badgeEl == nil {
			t.Fatal("Expected span element not found")
		}
		
		classAttr := getAttribute(badgeEl, "class")
		if !hasClass(classAttr, "badge") {
			t.Errorf("Expected 'badge' class, got classes: %s", classAttr)
		}
		
		if !strings.Contains(html, "New") {
			t.Errorf("Expected 'New' text, got: %s", html)
		}
	})
	
	t.Run("renders badge with multiple children", func(t *testing.T) {
		badge := NewBadge(
			h.Span(g.Text("Icon")),
			g.Text(" Count: 5"),
		)
		html := renderToHTML(badge)
		
		if !strings.Contains(html, "<span>Icon</span>") {
			t.Errorf("Expected span with Icon, got: %s", html)
		}
		if !strings.Contains(html, " Count: 5") {
			t.Errorf("Expected ' Count: 5' text, got: %s", html)
		}
	})
}

func TestBadge_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary badge", flyon.Primary, "badge-primary"},
		{"Secondary badge", flyon.Secondary, "badge-secondary"},
		{"Success badge", flyon.Success, "badge-success"},
		{"Warning badge", flyon.Warning, "badge-warning"},
		{"Error badge", flyon.Error, "badge-error"},
		{"Info badge", flyon.Info, "badge-info"},
		{"Neutral badge", flyon.Neutral, "badge-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := NewBadge(g.Text("New")).With(tt.color)
			html := renderToHTML(badge)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			badgeEl := findElement(doc, "span")
			if badgeEl == nil {
				t.Fatal("Expected span element not found")
			}
			
			classAttr := getAttribute(badgeEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestBadge_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small badge", flyon.SizeXS, "badge-xs"},
		{"Small badge", flyon.SizeSmall, "badge-sm"},
		{"Medium badge", flyon.SizeMedium, "badge-md"},
		{"Large badge", flyon.SizeLarge, "badge-lg"},
		{"Extra large badge", flyon.SizeXL, "badge-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := NewBadge(g.Text("New")).With(tt.size)
			html := renderToHTML(badge)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			badgeEl := findElement(doc, "span")
			if badgeEl == nil {
				t.Fatal("Expected span element not found")
			}
			
			classAttr := getAttribute(badgeEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestBadge_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid badge", flyon.VariantSolid, "badge-solid"},
		{"Outline badge", flyon.VariantOutline, "badge-outline"},
		{"Ghost badge", flyon.VariantGhost, "badge-ghost"},
		{"Soft badge", flyon.VariantSoft, "badge-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := NewBadge(g.Text("New")).With(tt.variant)
			html := renderToHTML(badge)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			badgeEl := findElement(doc, "span")
			if badgeEl == nil {
				t.Fatal("Expected span element not found")
			}
			
			classAttr := getAttribute(badgeEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestBadge_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		badge := NewBadge(g.Text("New")).With(
			flyon.Primary,
			flyon.SizeLarge,
			flyon.VariantOutline,
		)
		html := renderToHTML(badge)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		badgeEl := findElement(doc, "span")
		if badgeEl == nil {
			t.Fatal("Expected span element not found")
		}
		
		classAttr := getAttribute(badgeEl, "class")
		expectedClasses := []string{"badge", "badge-primary", "badge-lg", "badge-outline"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestBadge_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		badge := NewBadge(
			h.ID("status-badge"),
			h.Title("Status indicator"),
			g.Text("Active"),
		)
		html := renderToHTML(badge)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		badgeEl := findElement(doc, "span")
		if badgeEl == nil {
			t.Fatal("Expected span element not found")
		}
		
		if id := getAttribute(badgeEl, "id"); id != "status-badge" {
			t.Errorf("Expected id='status-badge', got id='%s'", id)
		}
		
		if title := getAttribute(badgeEl, "title"); title != "Status indicator" {
			t.Errorf("Expected title='Status indicator', got title='%s'", title)
		}
	})
}

func TestBadge_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		badge := NewBadge(g.Text("New"))
		
		// Test that it implements flyon.Component
		var _ flyon.Component = badge
		
		// Test that it implements gomponents.Node
		var _ g.Node = badge
		
		// Test that With returns a new instance
		newBadge := badge.With(flyon.Primary)
		if badge == newBadge {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}