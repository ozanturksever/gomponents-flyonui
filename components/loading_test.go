//go:build js && wasm

package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestLoading_BasicRendering(t *testing.T) {
	tests := []struct {
		name     string
		loading  *LoadingComponent
		expected string
	}{
		{
			name:     "basic loading",
			loading:  NewLoading(),
			expected: "loading",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := renderToHTML(tt.loading)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected loading to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestLoading_WithContent(t *testing.T) {
	loading := NewLoading(g.Text("Loading..."))
	html := renderToHTML(loading)

	if !strings.Contains(html, "Loading...") {
		t.Errorf("Expected loading to contain 'Loading...', got: %s", html)
	}
}

func TestLoading_WithType(t *testing.T) {
	tests := []struct {
		name     string
		type_    LoadingType
		expected string
	}{
		{"spinner", LoadingSpinner, "loading-spinner"},
		{"dots", LoadingDots, "loading-dots"},
		{"ring", LoadingRing, "loading-ring"},
		{"ball", LoadingBall, "loading-ball"},
		{"bars", LoadingBars, "loading-bars"},
		{"infinity", LoadingInfinity, "loading-infinity"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loading := NewLoading().WithType(tt.type_)
			html := renderToHTML(loading)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected loading to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestLoading_WithColorModifiers(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"primary", flyon.Primary, "loading-primary"},
		{"secondary", flyon.Secondary, "loading-secondary"},
		{"success", flyon.Success, "loading-success"},
		{"warning", flyon.Warning, "loading-warning"},
		{"error", flyon.Error, "loading-error"},
		{"info", flyon.Info, "loading-info"},
		{"neutral", flyon.Neutral, "loading-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loading := NewLoading().With(tt.color)
			html := renderToHTML(loading)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected loading to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestLoading_WithSizeModifiers(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"xs", flyon.SizeXS, "loading-xs"},
		{"sm", flyon.SizeSmall, "loading-sm"},
		{"md", flyon.SizeMedium, "loading-md"},
		{"lg", flyon.SizeLarge, "loading-lg"},
		{"xl", flyon.SizeXL, "loading-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loading := NewLoading().With(tt.size)
			html := renderToHTML(loading)

			if !strings.Contains(html, tt.expected) {
				t.Errorf("Expected loading to contain '%s', got: %s", tt.expected, html)
			}
		})
	}
}

func TestLoading_CombinedModifiers(t *testing.T) {
	loading := NewLoading().WithType(LoadingSpinner).With(flyon.Primary, flyon.SizeLarge, "custom-class")
	html := renderToHTML(loading)

	expectedClasses := []string{
		"loading",
		"loading-spinner",
		"loading-primary",
		"loading-lg",
		"custom-class",
	}

	for _, class := range expectedClasses {
		if !strings.Contains(html, class) {
			t.Errorf("Expected loading to contain class '%s', got: %s", class, html)
		}
	}
}

func TestLoading_WithHTMLAttributes(t *testing.T) {
	loading := NewLoading().With(
		h.ID("test-loading"),
		g.Attr("data-testid", "loading"),
	)
	html := renderToHTML(loading)

	if !strings.Contains(html, `id="test-loading"`) {
		t.Errorf("Expected loading to contain id attribute, got: %s", html)
	}

	if !strings.Contains(html, `data-testid="loading"`) {
		t.Errorf("Expected loading to contain data-testid attribute, got: %s", html)
	}
}

func TestLoading_ImplementsInterfaces(t *testing.T) {
	// Test that LoadingComponent implements flyon.Component
	var _ flyon.Component = (*LoadingComponent)(nil)

	// Test that LoadingComponent implements gomponents.Node
	var _ g.Node = (*LoadingComponent)(nil)
}

func TestLoading_WithChildren(t *testing.T) {
	loading := NewLoading(
		h.Span(g.Text("Icon")),
		g.Text(" Loading"),
	)
	html := renderToHTML(loading)

	if !strings.Contains(html, "<span>Icon</span>") {
		t.Errorf("Expected loading to contain span element, got: %s", html)
	}

	if !strings.Contains(html, " Loading") {
		t.Errorf("Expected loading to contain ' Loading' text, got: %s", html)
	}
}