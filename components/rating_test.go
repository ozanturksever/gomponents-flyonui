//go:build js && wasm

package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestRating_BasicRendering(t *testing.T) {
	t.Run("renders basic rating", func(t *testing.T) {
		rating := NewRating(5)
		html := renderToHTML(rating)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		ratingEl := findElement(doc, "div")
		if ratingEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(ratingEl, "class")
		if !hasClass(classAttr, "rating") {
			t.Errorf("Expected 'rating' class, got classes: %s", classAttr)
		}
	})
}

func TestRating_ContentRendering(t *testing.T) {
	t.Run("renders rating with content", func(t *testing.T) {
		rating := NewRating(3, h.Span(g.Text("3 stars")))
		html := renderToHTML(rating)

		if !strings.Contains(html, "3 stars") {
			t.Errorf("Expected rating to contain '3 stars', got: %s", html)
		}
	})
}

func TestRating_ValueAttribute(t *testing.T) {
	t.Run("sets rating value attribute", func(t *testing.T) {
		rating := NewRating(4)
		html := renderToHTML(rating)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		ratingEl := findElement(doc, "div")
		if ratingEl == nil {
			t.Fatal("Expected div element not found")
		}

		valueAttr := getAttribute(ratingEl, "data-rating")
		if valueAttr != "4" {
			t.Errorf("Expected data-rating='4', got data-rating='%s'", valueAttr)
		}
	})
}

func TestRating_ColorModifiers(t *testing.T) {
	t.Run("applies color modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			color    flyon.Color
			expected string
		}{
			{"Primary color", flyon.Primary, "rating-primary"},
			{"Secondary color", flyon.Secondary, "rating-secondary"},
			{"Success color", flyon.Success, "rating-success"},
			{"Warning color", flyon.Warning, "rating-warning"},
			{"Error color", flyon.Error, "rating-error"},
			{"Info color", flyon.Info, "rating-info"},
			{"Neutral color", flyon.Neutral, "rating-neutral"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rating := NewRating(5).With(tt.color)
				html := renderToHTML(rating)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				ratingEl := findElement(doc, "div")
				if ratingEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(ratingEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestRating_SizeModifiers(t *testing.T) {
	t.Run("applies size modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			size     flyon.Size
			expected string
		}{
			{"XS size", flyon.SizeXS, "rating-xs"},
			{"SM size", flyon.SizeSmall, "rating-sm"},
			{"MD size", flyon.SizeMedium, "rating-md"},
			{"LG size", flyon.SizeLarge, "rating-lg"},
			{"XL size", flyon.SizeXL, "rating-xl"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rating := NewRating(5).With(tt.size)
				html := renderToHTML(rating)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				ratingEl := findElement(doc, "div")
				if ratingEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(ratingEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestRating_CombinedModifiers(t *testing.T) {
	t.Run("applies combined modifiers", func(t *testing.T) {
		rating := NewRating(4).With(flyon.Primary, flyon.SizeLarge)
		html := renderToHTML(rating)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		ratingEl := findElement(doc, "div")
		if ratingEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(ratingEl, "class")
		if !hasClass(classAttr, "rating") {
			t.Errorf("Expected 'rating' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "rating-primary") {
			t.Errorf("Expected 'rating-primary' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "rating-lg") {
			t.Errorf("Expected 'rating-lg' class, got classes: %s", classAttr)
		}
	})
}

func TestRating_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		rating := NewRating(3).With(
			h.ID("main-rating"),
			g.Attr("data-testid", "rating-component"),
			h.Title("Rating widget"),
		)
		html := renderToHTML(rating)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		ratingEl := findElement(doc, "div")
		if ratingEl == nil {
			t.Fatal("Expected div element not found")
		}

		if id := getAttribute(ratingEl, "id"); id != "main-rating" {
			t.Errorf("Expected id='main-rating', got id='%s'", id)
		}

		if testid := getAttribute(ratingEl, "data-testid"); testid != "rating-component" {
			t.Errorf("Expected data-testid='rating-component', got data-testid='%s'", testid)
		}

		if title := getAttribute(ratingEl, "title"); title != "Rating widget" {
			t.Errorf("Expected title='Rating widget', got title='%s'", title)
		}
	})
}

func TestRating_ComponentInterface(t *testing.T) {
	t.Run("implements flyon.Component interface", func(t *testing.T) {
		var _ flyon.Component = (*RatingComponent)(nil)
	})

	t.Run("implements gomponents.Node interface", func(t *testing.T) {
		var _ g.Node = (*RatingComponent)(nil)
	})

	t.Run("With method returns flyon.Component", func(t *testing.T) {
		rating := NewRating(5)
		result := rating.With("custom-class")
		if _, ok := result.(flyon.Component); !ok {
			t.Error("With method should return flyon.Component")
		}
	})
}