package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestBreadcrumb_BasicRendering(t *testing.T) {
	t.Run("renders basic breadcrumb with default classes", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			BreadcrumbItem(g.Text("Home")),
			BreadcrumbItem(g.Text("Products")),
			BreadcrumbItem(g.Text("Laptops")),
		)
		html := renderToHTML(breadcrumb)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		breadcrumbEl := findElement(doc, "nav")
		if breadcrumbEl == nil {
			t.Fatal("Expected nav element not found")
		}
		
		classAttr := getAttribute(breadcrumbEl, "class")
		if !hasClass(classAttr, "breadcrumbs") {
			t.Errorf("Expected 'breadcrumbs' class, got classes: %s", classAttr)
		}
		
		if !strings.Contains(html, "Home") {
			t.Errorf("Expected 'Home' text, got: %s", html)
		}
		if !strings.Contains(html, "Products") {
			t.Errorf("Expected 'Products' text, got: %s", html)
		}
		if !strings.Contains(html, "Laptops") {
			t.Errorf("Expected 'Laptops' text, got: %s", html)
		}
	})
	
	t.Run("renders breadcrumb with single item", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			BreadcrumbItem(g.Text("Home")),
		)
		html := renderToHTML(breadcrumb)
		
		if !strings.Contains(html, "Home") {
			t.Errorf("Expected 'Home' text, got: %s", html)
		}
	})
}

func TestBreadcrumb_WithLinks(t *testing.T) {
	t.Run("renders breadcrumb items with links", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			BreadcrumbItem(h.A(h.Href("/"), g.Text("Home"))),
			BreadcrumbItem(h.A(h.Href("/products"), g.Text("Products"))),
			BreadcrumbItem(g.Text("Current Page")),
		)
		html := renderToHTML(breadcrumb)
		
		if !strings.Contains(html, `<a href="/">Home</a>`) {
			t.Errorf("Expected Home link, got: %s", html)
		}
		if !strings.Contains(html, `<a href="/products">Products</a>`) {
			t.Errorf("Expected Products link, got: %s", html)
		}
		if !strings.Contains(html, "Current Page") {
			t.Errorf("Expected 'Current Page' text, got: %s", html)
		}
	})
}

func TestBreadcrumb_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small breadcrumb", flyon.SizeXS, "breadcrumbs-xs"},
		{"Small breadcrumb", flyon.SizeSmall, "breadcrumbs-sm"},
		{"Medium breadcrumb", flyon.SizeMedium, "breadcrumbs-md"},
		{"Large breadcrumb", flyon.SizeLarge, "breadcrumbs-lg"},
		{"Extra large breadcrumb", flyon.SizeXL, "breadcrumbs-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			breadcrumb := NewBreadcrumb(
				BreadcrumbItem(g.Text("Home")),
				BreadcrumbItem(g.Text("Page")),
			).With(tt.size)
			html := renderToHTML(breadcrumb)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			breadcrumbEl := findElement(doc, "nav")
			if breadcrumbEl == nil {
				t.Fatal("Expected nav element not found")
			}
			
			classAttr := getAttribute(breadcrumbEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestBreadcrumb_CustomSeparator(t *testing.T) {
	t.Run("renders breadcrumb with custom separator", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			BreadcrumbItem(g.Text("Home")),
			BreadcrumbItem(g.Text("Products")),
			BreadcrumbItem(g.Text("Laptops")),
		).WithSeparator(">")
		html := renderToHTML(breadcrumb)
		
		// Check that custom separator is applied via CSS variable or data attribute
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		breadcrumbEl := findElement(doc, "nav")
		if breadcrumbEl == nil {
			t.Fatal("Expected nav element not found")
		}
		
		// Check for data-separator attribute or style attribute
		separatorAttr := getAttribute(breadcrumbEl, "data-separator")
		if separatorAttr != ">" {
			t.Errorf("Expected data-separator='>', got: %s", separatorAttr)
		}
	})
}

func TestBreadcrumb_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			h.ID("main-breadcrumb"),
			h.DataAttr("testid", "breadcrumb-component"),
			g.Attr("aria-label", "Breadcrumb navigation"),
			BreadcrumbItem(g.Text("Home")),
			BreadcrumbItem(g.Text("Page")),
		)
		html := renderToHTML(breadcrumb)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		breadcrumbEl := findElement(doc, "nav")
		if breadcrumbEl == nil {
			t.Fatal("Expected nav element not found")
		}
		
		if id := getAttribute(breadcrumbEl, "id"); id != "main-breadcrumb" {
			t.Errorf("Expected id='main-breadcrumb', got id='%s'", id)
		}
		
		if testid := getAttribute(breadcrumbEl, "data-testid"); testid != "breadcrumb-component" {
			t.Errorf("Expected data-testid='breadcrumb-component', got data-testid='%s'", testid)
		}
		
		if ariaLabel := getAttribute(breadcrumbEl, "aria-label"); ariaLabel != "Breadcrumb navigation" {
			t.Errorf("Expected aria-label='Breadcrumb navigation', got aria-label='%s'", ariaLabel)
		}
	})
}

func TestBreadcrumb_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		breadcrumb := NewBreadcrumb(
			BreadcrumbItem(g.Text("Home")),
		)
		
		// Test that it implements flyon.Component
		var _ flyon.Component = breadcrumb
		
		// Test that it implements gomponents.Node
		var _ g.Node = breadcrumb
		
		// Test that With returns a new instance
		newBreadcrumb := breadcrumb.With(flyon.SizeLarge)
		if breadcrumb == newBreadcrumb {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}