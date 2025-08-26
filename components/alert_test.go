package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestAlert_BasicRendering(t *testing.T) {
	t.Run("renders basic alert with default classes", func(t *testing.T) {
		alert := NewAlert(g.Text("Alert message"))
		html := renderToHTML(alert)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		alertEl := findElement(doc, "div")
		if alertEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(alertEl, "class")
		if !hasClass(classAttr, "alert") {
			t.Errorf("Expected 'alert' class, got classes: %s", classAttr)
		}
		
		if !strings.Contains(html, "Alert message") {
			t.Errorf("Expected 'Alert message' text, got: %s", html)
		}
	})
	
	t.Run("renders alert with multiple children", func(t *testing.T) {
		alert := NewAlert(
			h.Strong(g.Text("Warning!")),
			g.Text(" This is an important message."),
			h.A(h.Href("#"), g.Text("Learn more")),
		)
		html := renderToHTML(alert)
		
		if !strings.Contains(html, "<strong>Warning!</strong>") {
			t.Errorf("Expected strong with Warning!, got: %s", html)
		}
		if !strings.Contains(html, " This is an important message.") {
			t.Errorf("Expected message text, got: %s", html)
		}
		if !strings.Contains(html, "<a href=\"#\">Learn more</a>") {
			t.Errorf("Expected link with Learn more, got: %s", html)
		}
	})
}

func TestAlert_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary alert", flyon.Primary, "alert-primary"},
		{"Secondary alert", flyon.Secondary, "alert-secondary"},
		{"Success alert", flyon.Success, "alert-success"},
		{"Warning alert", flyon.Warning, "alert-warning"},
		{"Error alert", flyon.Error, "alert-error"},
		{"Info alert", flyon.Info, "alert-info"},
		{"Neutral alert", flyon.Neutral, "alert-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := NewAlert(g.Text("Content")).With(tt.color)
			html := renderToHTML(alert)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			alertEl := findElement(doc, "div")
			if alertEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(alertEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAlert_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small alert", flyon.SizeXS, "alert-xs"},
		{"Small alert", flyon.SizeSmall, "alert-sm"},
		{"Medium alert", flyon.SizeMedium, "alert-md"},
		{"Large alert", flyon.SizeLarge, "alert-lg"},
		{"Extra large alert", flyon.SizeXL, "alert-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := NewAlert(g.Text("Content")).With(tt.size)
			html := renderToHTML(alert)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			alertEl := findElement(doc, "div")
			if alertEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(alertEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAlert_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid alert", flyon.VariantSolid, "alert-solid"},
		{"Outline alert", flyon.VariantOutline, "alert-outline"},
		{"Ghost alert", flyon.VariantGhost, "alert-ghost"},
		{"Soft alert", flyon.VariantSoft, "alert-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := NewAlert(g.Text("Content")).With(tt.variant)
			html := renderToHTML(alert)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			alertEl := findElement(doc, "div")
			if alertEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(alertEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAlert_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		alert := NewAlert(g.Text("Content")).With(
			flyon.Warning,
			flyon.SizeLarge,
			flyon.VariantOutline,
		)
		html := renderToHTML(alert)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		alertEl := findElement(doc, "div")
		if alertEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(alertEl, "class")
		expectedClasses := []string{"alert", "alert-warning", "alert-lg", "alert-outline"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestAlert_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		alert := NewAlert(
			h.ID("main-alert"),
			h.DataAttr("testid", "alert-component"),
			h.Role("alert"),
			g.Text("Alert content"),
		)
		html := renderToHTML(alert)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		alertEl := findElement(doc, "div")
		if alertEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		if id := getAttribute(alertEl, "id"); id != "main-alert" {
			t.Errorf("Expected id='main-alert', got id='%s'", id)
		}
		
		if testid := getAttribute(alertEl, "data-testid"); testid != "alert-component" {
			t.Errorf("Expected data-testid='alert-component', got data-testid='%s'", testid)
		}
		
		if role := getAttribute(alertEl, "role"); role != "alert" {
			t.Errorf("Expected role='alert', got role='%s'", role)
		}
	})
}

func TestAlert_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		alert := NewAlert(g.Text("Content"))
		
		// Test that it implements flyon.Component
		var _ flyon.Component = alert
		
		// Test that it implements gomponents.Node
		var _ g.Node = alert
		
		// Test that With returns a new instance
		newAlert := alert.With(flyon.Success)
		if alert == newAlert {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}