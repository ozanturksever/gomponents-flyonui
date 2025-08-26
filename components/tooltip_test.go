package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestTooltip_BasicRendering(t *testing.T) {
	t.Run("renders basic tooltip", func(t *testing.T) {
		tooltip := NewTooltip("Tooltip text", h.Button(g.Text("Hover me")))
		html := renderToHTML(tooltip)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		tooltipEl := findElement(doc, "div")
		if tooltipEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(tooltipEl, "class")
		if !hasClass(classAttr, "tooltip") {
			t.Errorf("Expected 'tooltip' class, got classes: %s", classAttr)
		}

		tipAttr := getAttribute(tooltipEl, "data-tip")
		if tipAttr != "Tooltip text" {
			t.Errorf("Expected data-tip='Tooltip text', got data-tip='%s'", tipAttr)
		}
	})

	t.Run("renders tooltip with content", func(t *testing.T) {
		tooltip := NewTooltip("Help text", h.Span(g.Text("Help icon")))
		html := renderToHTML(tooltip)

		if !strings.Contains(html, "Help icon") {
			t.Errorf("Expected tooltip to contain 'Help icon', got: %s", html)
		}
		if !strings.Contains(html, `data-tip="Help text"`) {
			t.Errorf("Expected tooltip to contain data-tip attribute, got: %s", html)
		}
	})
}

func TestTooltip_PositionModifiers(t *testing.T) {
	t.Run("applies position modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			position TooltipPosition
			expected string
		}{
			{"Top position", TooltipTop, "tooltip-top"},
			{"Bottom position", TooltipBottom, "tooltip-bottom"},
			{"Left position", TooltipLeft, "tooltip-left"},
			{"Right position", TooltipRight, "tooltip-right"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tooltip := NewTooltip("Test", h.Button(g.Text("Button"))).WithPosition(tt.position)
				html := renderToHTML(tooltip)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				tooltipEl := findElement(doc, "div")
				if tooltipEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(tooltipEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestTooltip_ColorModifiers(t *testing.T) {
	t.Run("applies color modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			color    flyon.Color
			expected string
		}{
			{"Primary color", flyon.Primary, "tooltip-primary"},
			{"Secondary color", flyon.Secondary, "tooltip-secondary"},
			{"Success color", flyon.Success, "tooltip-success"},
			{"Warning color", flyon.Warning, "tooltip-warning"},
			{"Error color", flyon.Error, "tooltip-error"},
			{"Info color", flyon.Info, "tooltip-info"},
			{"Neutral color", flyon.Neutral, "tooltip-neutral"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tooltip := NewTooltip("Test", h.Button(g.Text("Button"))).With(tt.color)
				html := renderToHTML(tooltip)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				tooltipEl := findElement(doc, "div")
				if tooltipEl == nil {
					t.Fatal("Expected div element not found")
				}

				classAttr := getAttribute(tooltipEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestTooltip_OpenModifier(t *testing.T) {
	t.Run("applies open modifier", func(t *testing.T) {
		tooltip := NewTooltip("Test", h.Button(g.Text("Button"))).WithOpen()
		html := renderToHTML(tooltip)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		tooltipEl := findElement(doc, "div")
		if tooltipEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(tooltipEl, "class")
		if !hasClass(classAttr, "tooltip-open") {
			t.Errorf("Expected 'tooltip-open' class, got classes: %s", classAttr)
		}
	})
}

func TestTooltip_CombinedModifiers(t *testing.T) {
	t.Run("applies combined modifiers", func(t *testing.T) {
		tooltip := NewTooltip("Test tooltip", h.Button(g.Text("Button"))).WithPosition(TooltipTop).WithOpen().With(flyon.Primary)
		html := renderToHTML(tooltip)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		tooltipEl := findElement(doc, "div")
		if tooltipEl == nil {
			t.Fatal("Expected div element not found")
		}

		classAttr := getAttribute(tooltipEl, "class")
		if !hasClass(classAttr, "tooltip") {
			t.Errorf("Expected 'tooltip' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "tooltip-top") {
			t.Errorf("Expected 'tooltip-top' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "tooltip-primary") {
			t.Errorf("Expected 'tooltip-primary' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "tooltip-open") {
			t.Errorf("Expected 'tooltip-open' class, got classes: %s", classAttr)
		}
	})
}

func TestTooltip_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		tooltip := NewTooltip("Help", h.Button(g.Text("Help"))).With(
			h.ID("help-tooltip"),
			g.Attr("data-testid", "tooltip-component"),
			h.Title("Tooltip container"),
		)
		html := renderToHTML(tooltip)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		tooltipEl := findElement(doc, "div")
		if tooltipEl == nil {
			t.Fatal("Expected div element not found")
		}

		if id := getAttribute(tooltipEl, "id"); id != "help-tooltip" {
			t.Errorf("Expected id='help-tooltip', got id='%s'", id)
		}

		if testid := getAttribute(tooltipEl, "data-testid"); testid != "tooltip-component" {
			t.Errorf("Expected data-testid='tooltip-component', got data-testid='%s'", testid)
		}

		if title := getAttribute(tooltipEl, "title"); title != "Tooltip container" {
			t.Errorf("Expected title='Tooltip container', got title='%s'", title)
		}
	})
}

func TestTooltip_ComponentInterface(t *testing.T) {
	t.Run("implements flyon.Component interface", func(t *testing.T) {
		var _ flyon.Component = (*TooltipComponent)(nil)
	})

	t.Run("implements gomponents.Node interface", func(t *testing.T) {
		var _ g.Node = (*TooltipComponent)(nil)
	})

	t.Run("With method returns flyon.Component", func(t *testing.T) {
		tooltip := NewTooltip("Test", h.Button(g.Text("Button")))
		result := tooltip.With("custom-class")
		if _, ok := result.(flyon.Component); !ok {
			t.Error("With method should return flyon.Component")
		}
	})
}