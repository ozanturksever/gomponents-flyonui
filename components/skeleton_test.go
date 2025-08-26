package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"golang.org/x/net/html"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Test helper to get text content from HTML node
func getTextContent(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text.WriteString(getTextContent(c))
	}
	return text.String()
}

func TestSkeleton_BasicRendering(t *testing.T) {
	t.Run("renders basic skeleton", func(t *testing.T) {
		skeleton := NewSkeleton()
		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(skeletonEl, "class")
		if !hasClass(classAttr, "skeleton") {
			t.Errorf("Expected 'skeleton' base class, got classes: %s", classAttr)
		}
	})

	t.Run("renders skeleton with content", func(t *testing.T) {
		skeleton := NewSkeleton(g.Text("Loading..."))
		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		textContent := getTextContent(skeletonEl)
		if textContent != "Loading..." {
			t.Errorf("Expected text content 'Loading...', got: %s", textContent)
		}
	})
}

func TestSkeleton_ShapeModifiers(t *testing.T) {
	tests := []struct {
		name     string
		shape    SkeletonShape
		expected string
	}{
		{"Rectangle shape", SkeletonRectangle, "skeleton-rectangle"},
		{"Circle shape", SkeletonCircle, "skeleton-circle"},
		{"Text shape", SkeletonText, "skeleton-text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skeleton := NewSkeleton().WithShape(tt.shape)
			html := renderToHTML(skeleton)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			skeletonEl := findElement(doc, "div")
			if skeletonEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(skeletonEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
			if !hasClass(classAttr, "skeleton") {
				t.Errorf("Expected 'skeleton' base class, got classes: %s", classAttr)
			}
		})
	}
}

func TestSkeleton_SizeModifiers(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small size", flyon.SizeXS, "skeleton-xs"},
		{"Small size", flyon.SizeSmall, "skeleton-sm"},
		{"Medium size", flyon.SizeMedium, "skeleton-md"},
		{"Large size", flyon.SizeLarge, "skeleton-lg"},
		{"Extra large size", flyon.SizeXL, "skeleton-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skeleton := NewSkeleton().With(tt.size)
			html := renderToHTML(skeleton)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			skeletonEl := findElement(doc, "div")
			if skeletonEl == nil {
				t.Fatal("Expected div element not found")
			}

			classAttr := getAttribute(skeletonEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestSkeleton_AnimationModifiers(t *testing.T) {
	t.Run("applies pulse animation", func(t *testing.T) {
		skeleton := NewSkeleton().WithPulse()
		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(skeletonEl, "class")
		if !hasClass(classAttr, "skeleton-pulse") {
			t.Errorf("Expected 'skeleton-pulse' class, got classes: %s", classAttr)
		}
	})

	t.Run("applies wave animation", func(t *testing.T) {
		skeleton := NewSkeleton().WithWave()
		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(skeletonEl, "class")
		if !hasClass(classAttr, "skeleton-wave") {
			t.Errorf("Expected 'skeleton-wave' class, got classes: %s", classAttr)
		}
	})
}

func TestSkeleton_CombinedModifiers(t *testing.T) {
	t.Run("applies multiple modifiers correctly", func(t *testing.T) {
		skeleton := NewSkeleton().WithShape(SkeletonCircle).WithPulse().With(
			flyon.SizeLarge,
			"custom-class",
		)

		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(skeletonEl, "class")
		expectedClasses := []string{"skeleton", "skeleton-circle", "skeleton-pulse", "skeleton-lg", "custom-class"}

		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestSkeleton_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		skeleton := NewSkeleton().With(
			h.ID("test-skeleton"),
			g.Attr("data-test", "skeleton-component"),
			h.Title("Loading Skeleton"),
		)

		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		if id := getAttribute(skeletonEl, "id"); id != "test-skeleton" {
			t.Errorf("Expected id='test-skeleton', got id='%s'", id)
		}

		if testAttr := getAttribute(skeletonEl, "data-test"); testAttr != "skeleton-component" {
			t.Errorf("Expected data-test='skeleton-component', got data-test='%s'", testAttr)
		}

		if title := getAttribute(skeletonEl, "title"); title != "Loading Skeleton" {
			t.Errorf("Expected title='Loading Skeleton', got title='%s'", title)
		}
	})
}

func TestSkeleton_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = NewSkeleton()

		skeleton := NewSkeleton(g.Text("Content"))
		html := renderToHTML(skeleton)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		skeletonEl := findElement(doc, "div")
		if skeletonEl == nil {
			t.Fatal("Expected div element not found")
		}

		textContent := getTextContent(skeletonEl)
		if textContent != "Content" {
			t.Errorf("Expected text content 'Content', got: %s", textContent)
		}
	})
}