package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestSpinner_BasicRendering(t *testing.T) {
	t.Run("renders basic spinner", func(t *testing.T) {
		spinner := NewSpinner()
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		if spinnerEl == nil {
			t.Fatal("Expected span element not found")
		}

		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading") {
			t.Errorf("Expected 'loading' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders spinner with content", func(t *testing.T) {
		spinner := NewSpinner(g.Text("Loading..."))
		html := renderToHTML(spinner)

		if !strings.Contains(html, "Loading...") {
			t.Errorf("Expected spinner to contain 'Loading...', got: %s", html)
		}
	})
}

func TestSpinner_TypeModifiers(t *testing.T) {
	t.Run("applies spinner type", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerDots)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		if spinnerEl == nil {
			t.Fatal("Expected span element not found")
		}

		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading-dots") {
			t.Errorf("Expected 'loading-dots' class, got classes: %s", classAttr)
		}
	})

	t.Run("applies ring type", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerRing)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading-ring") {
			t.Errorf("Expected 'loading-ring' class, got classes: %s", classAttr)
		}
	})

	t.Run("applies ball type", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerBall)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading-ball") {
			t.Errorf("Expected 'loading-ball' class, got classes: %s", classAttr)
		}
	})

	t.Run("applies bars type", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerBars)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading-bars") {
			t.Errorf("Expected 'loading-bars' class, got classes: %s", classAttr)
		}
	})

	t.Run("applies infinity type", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerInfinity)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading-infinity") {
			t.Errorf("Expected 'loading-infinity' class, got classes: %s", classAttr)
		}
	})
}

func TestSpinner_SizeModifiers(t *testing.T) {
	t.Run("applies size modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			size     flyon.Size
			expected string
		}{
			{"XS size", flyon.SizeXS, "loading-xs"},
			{"Small size", flyon.SizeSmall, "loading-sm"},
			{"Medium size", flyon.SizeMedium, "loading-md"},
			{"Large size", flyon.SizeLarge, "loading-lg"},
			{"XL size", flyon.SizeXL, "loading-xl"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				spinner := NewSpinner().With(tt.size)
				html := renderToHTML(spinner)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				spinnerEl := findElement(doc, "span")
				if spinnerEl == nil {
					t.Fatal("Expected span element not found")
				}

				classAttr := getAttribute(spinnerEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestSpinner_CombinedModifiers(t *testing.T) {
	t.Run("applies combined modifiers", func(t *testing.T) {
		spinner := NewSpinner().WithType(SpinnerDots).With(flyon.SizeLarge)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		if spinnerEl == nil {
			t.Fatal("Expected span element not found")
		}

		classAttr := getAttribute(spinnerEl, "class")
		if !hasClass(classAttr, "loading") {
			t.Errorf("Expected 'loading' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "loading-dots") {
			t.Errorf("Expected 'loading-dots' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "loading-lg") {
			t.Errorf("Expected 'loading-lg' class, got classes: %s", classAttr)
		}
	})
}

func TestSpinner_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		spinner := NewSpinner().With(
			h.ID("loading-spinner"),
			g.Attr("data-testid", "spinner-component"),
			h.Title("Loading content"),
		)
		html := renderToHTML(spinner)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		spinnerEl := findElement(doc, "span")
		if spinnerEl == nil {
			t.Fatal("Expected span element not found")
		}

		if id := getAttribute(spinnerEl, "id"); id != "loading-spinner" {
			t.Errorf("Expected id='loading-spinner', got id='%s'", id)
		}

		if testid := getAttribute(spinnerEl, "data-testid"); testid != "spinner-component" {
			t.Errorf("Expected data-testid='spinner-component', got data-testid='%s'", testid)
		}

		if title := getAttribute(spinnerEl, "title"); title != "Loading content" {
			t.Errorf("Expected title='Loading content', got title='%s'", title)
		}
	})
}

func TestSpinner_ComponentInterface(t *testing.T) {
	t.Run("implements flyon.Component interface", func(t *testing.T) {
		var _ flyon.Component = (*SpinnerComponent)(nil)
	})

	t.Run("implements gomponents.Node interface", func(t *testing.T) {
		var _ g.Node = (*SpinnerComponent)(nil)
	})

	t.Run("With method returns flyon.Component", func(t *testing.T) {
		spinner := NewSpinner()
		result := spinner.With("custom-class")
		if _, ok := result.(flyon.Component); !ok {
			t.Error("With method should return flyon.Component")
		}
	})
}