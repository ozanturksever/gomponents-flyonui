//go:build js && wasm

package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestIndicator_BasicRendering(t *testing.T) {
	tests := []struct {
		name     string
		indicator *IndicatorComponent
		expected string
	}{
		{
			name:     "basic indicator",
			indicator: NewIndicator(),
			expected: "indicator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := renderToHTML(tt.indicator)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected indicator to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestIndicator_WithContent(t *testing.T) {
	indicator := NewIndicator(g.Text("Badge content"))
	html := renderToHTML(indicator)

	if !strings.Contains(html, "Badge content") {
		t.Errorf("Expected indicator to contain 'Badge content', got: %s", html)
	}
}

func TestIndicator_WithPosition(t *testing.T) {
	tests := []struct {
		name     string
		position IndicatorPosition
		expected string
	}{
		{"top start", IndicatorTopStart, "indicator-top-start"},
		{"top center", IndicatorTopCenter, "indicator-top-center"},
		{"top end", IndicatorTopEnd, "indicator-top-end"},
		{"middle start", IndicatorMiddleStart, "indicator-middle-start"},
		{"middle center", IndicatorMiddleCenter, "indicator-middle-center"},
		{"middle end", IndicatorMiddleEnd, "indicator-middle-end"},
		{"bottom start", IndicatorBottomStart, "indicator-bottom-start"},
		{"bottom center", IndicatorBottomCenter, "indicator-bottom-center"},
		{"bottom end", IndicatorBottomEnd, "indicator-bottom-end"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indicator := NewIndicator().WithPosition(tt.position)
			html := renderToHTML(indicator)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected indicator to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestIndicator_WithColorModifiers(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"primary", flyon.Primary, "indicator-primary"},
		{"secondary", flyon.Secondary, "indicator-secondary"},
		{"success", flyon.Success, "indicator-success"},
		{"warning", flyon.Warning, "indicator-warning"},
		{"error", flyon.Error, "indicator-error"},
		{"info", flyon.Info, "indicator-info"},
		{"neutral", flyon.Neutral, "indicator-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indicator := NewIndicator().With(tt.color)
			html := renderToHTML(indicator)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected indicator to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestIndicator_WithSizeModifiers(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"xs", flyon.SizeXS, "indicator-xs"},
		{"sm", flyon.SizeSmall, "indicator-sm"},
		{"md", flyon.SizeMedium, "indicator-md"},
		{"lg", flyon.SizeLarge, "indicator-lg"},
		{"xl", flyon.SizeXL, "indicator-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indicator := NewIndicator().With(tt.size)
			html := renderToHTML(indicator)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected indicator to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestIndicator_CombinedModifiers(t *testing.T) {
	indicator := NewIndicator().WithPosition(IndicatorTopEnd).With(flyon.Primary, flyon.SizeLarge, "custom-class")
	html := renderToHTML(indicator)

	expectedClasses := []string{
		"indicator",
		"indicator-top-end",
		"indicator-primary",
		"indicator-lg",
		"custom-class",
	}

	for _, class := range expectedClasses {
		if !strings.Contains(html, class) {
			t.Errorf("Expected indicator to contain class '%s', got: %s", class, html)
		}
	}
}

func TestIndicator_WithHTMLAttributes(t *testing.T) {
	indicator := NewIndicator().With(
		h.ID("test-indicator"),
		g.Attr("data-testid", "indicator"),
	)
	html := renderToHTML(indicator)

	if !strings.Contains(html, `id="test-indicator"`) {
		t.Errorf("Expected indicator to contain id attribute, got: %s", html)
	}

	if !strings.Contains(html, `data-testid="indicator"`) {
		t.Errorf("Expected indicator to contain data-testid attribute, got: %s", html)
	}
}

func TestIndicator_ImplementsInterfaces(t *testing.T) {
	// Test that IndicatorComponent implements flyon.Component
	var _ flyon.Component = (*IndicatorComponent)(nil)

	// Test that IndicatorComponent implements gomponents.Node
	var _ g.Node = (*IndicatorComponent)(nil)
}

func TestIndicator_WithChildren(t *testing.T) {
	indicator := NewIndicator(
		h.Span(g.Text("Icon")),
		g.Text(" Badge"),
	)
	html := renderToHTML(indicator)

	if !strings.Contains(html, "<span>Icon</span>") {
		t.Errorf("Expected indicator to contain span element, got: %s", html)
	}

	if !strings.Contains(html, " Badge") {
		t.Errorf("Expected indicator to contain ' Badge' text, got: %s", html)
	}
}