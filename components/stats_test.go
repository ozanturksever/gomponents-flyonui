package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestStats_BasicRendering(t *testing.T) {
	t.Run("renders basic stats", func(t *testing.T) {
		stats := NewStats()
		html := renderToHTML(stats)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		statsEl := findElement(doc, "div")
		if statsEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(statsEl, "class")
		if !hasClass(classAttr, "stats") {
			t.Errorf("Expected 'stats' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders stats with content", func(t *testing.T) {
		stats := NewStats(h.Div(g.Text("Stat item")))
		html := renderToHTML(stats)

		if !strings.Contains(html, "Stat item") {
			t.Errorf("Expected stats to contain 'Stat item', got: %s", html)
		}
	})
}

func TestStats_OrientationModifiers(t *testing.T) {
	t.Run("applies orientation modifiers", func(t *testing.T) {
		tests := []struct {
			name        string
			orientation StatsOrientation
			expected    string
		}{
			{"Vertical orientation", StatsVertical, "stats-vertical"},
			{"Horizontal orientation", StatsHorizontal, "stats-horizontal"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stats := NewStats().WithOrientation(tt.orientation)
				html := renderToHTML(stats)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				statsEl := findElement(doc, "div")
				if statsEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(statsEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestStats_ColorModifiers(t *testing.T) {
	t.Run("applies color modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			color    flyon.Color
			expected string
		}{
			{"Primary color", flyon.Primary, "stats-primary"},
			{"Secondary color", flyon.Secondary, "stats-secondary"},
			{"Success color", flyon.Success, "stats-success"},
			{"Warning color", flyon.Warning, "stats-warning"},
			{"Error color", flyon.Error, "stats-error"},
			{"Info color", flyon.Info, "stats-info"},
			{"Neutral color", flyon.Neutral, "stats-neutral"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stats := NewStats().With(tt.color)
				html := renderToHTML(stats)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				statsEl := findElement(doc, "div")
				if statsEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(statsEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestStats_SizeModifiers(t *testing.T) {
	t.Run("applies size modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			size     flyon.Size
			expected string
		}{
			{"XS size", flyon.SizeXS, "stats-xs"},
			{"SM size", flyon.SizeSmall, "stats-sm"},
			{"MD size", flyon.SizeMedium, "stats-md"},
			{"LG size", flyon.SizeLarge, "stats-lg"},
			{"XL size", flyon.SizeXL, "stats-xl"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stats := NewStats().With(tt.size)
				html := renderToHTML(stats)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				statsEl := findElement(doc, "div")
				if statsEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(statsEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestStats_CombinedModifiers(t *testing.T) {
	t.Run("applies combined modifiers", func(t *testing.T) {
		stats := NewStats(h.Div(g.Text("Item"))).WithOrientation(StatsVertical).With(flyon.Primary, flyon.SizeLarge)
		html := renderToHTML(stats)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		statsEl := findElement(doc, "div")
		if statsEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(statsEl, "class")
		if !hasClass(classAttr, "stats") {
			t.Errorf("Expected 'stats' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "stats-vertical") {
			t.Errorf("Expected 'stats-vertical' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "stats-primary") {
			t.Errorf("Expected 'stats-primary' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "stats-lg") {
			t.Errorf("Expected 'stats-lg' class, got classes: %s", classAttr)
		}
	})
}

func TestStats_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		stats := NewStats().With(
			h.ID("main-stats"),
			g.Attr("data-testid", "stats-component"),
			h.Title("Stats container"),
		)
		html := renderToHTML(stats)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		statsEl := findElement(doc, "div")
		if statsEl == nil {
			t.Fatal("Expected div element not found")
		}

		if id := getAttribute(statsEl, "id"); id != "main-stats" {
			t.Errorf("Expected id='main-stats', got id='%s'", id)
		}

		if testid := getAttribute(statsEl, "data-testid"); testid != "stats-component" {
			t.Errorf("Expected data-testid='stats-component', got data-testid='%s'", testid)
		}

		if title := getAttribute(statsEl, "title"); title != "Stats container" {
			t.Errorf("Expected title='Stats container', got title='%s'", title)
		}
	})
}

func TestStats_ComponentInterface(t *testing.T) {
	t.Run("implements flyon.Component interface", func(t *testing.T) {
		var _ flyon.Component = (*StatsComponent)(nil)
	})

	t.Run("implements gomponents.Node interface", func(t *testing.T) {
		var _ g.Node = (*StatsComponent)(nil)
	})

	t.Run("With method returns flyon.Component", func(t *testing.T) {
		stats := NewStats()
		result := stats.With("custom-class")
		if _, ok := result.(flyon.Component); !ok {
			t.Error("With method should return flyon.Component")
		}
	})
}