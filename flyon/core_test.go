package flyon

import (
	"io"
	"strings"
	"testing"

	"maragu.dev/gomponents"
	"golang.org/x/net/html"
)

// Test helper to render gomponents.Node to HTML string
func renderToHTML(node gomponents.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

// Test helper to parse HTML and check for specific attributes/classes
func parseHTML(htmlStr string) (*html.Node, error) {
	return html.Parse(strings.NewReader(htmlStr))
}

// Test helper to find element with specific tag
func findElement(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findElement(c, tag); result != nil {
			return result
		}
	}
	return nil
}

// Test helper to get attribute value
func getAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func TestComponent_Interface(t *testing.T) {
	t.Run("Component interface compliance", func(t *testing.T) {
		// Test that our Component interface extends gomponents.Node
		var _ gomponents.Node = (*mockComponent)(nil)
		var _ Component = (*mockComponent)(nil)
	})
}

func TestColor_Enum(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"Primary color", Primary, "primary"},
		{"Secondary color", Secondary, "secondary"},
		{"Success color", Success, "success"},
		{"Warning color", Warning, "warning"},
		{"Error color", Error, "error"},
		{"Info color", Info, "info"},
		{"Neutral color", Neutral, "neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.color.String(); got != tt.expected {
				t.Errorf("Color.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSize_Enum(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected string
	}{
		{"Extra small size", SizeXS, "xs"},
		{"Small size", SizeSmall, "sm"},
		{"Medium size", SizeMedium, "md"},
		{"Large size", SizeLarge, "lg"},
		{"Extra large size", SizeXL, "xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.size.String(); got != tt.expected {
				t.Errorf("Size.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVariant_Enum(t *testing.T) {
	tests := []struct {
		name     string
		variant  Variant
		expected string
	}{
		{"Solid variant", VariantSolid, "solid"},
		{"Outline variant", VariantOutline, "outline"},
		{"Ghost variant", VariantGhost, "ghost"},
		{"Soft variant", VariantSoft, "soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.variant.String(); got != tt.expected {
				t.Errorf("Variant.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestModifierApplication(t *testing.T) {
	t.Run("Color modifier application", func(t *testing.T) {
		config := &mockConfig{}
		modifier := Primary
		
		modifier.ApplyToMock(config)
		
		if config.Color != Primary {
			t.Errorf("Expected color to be Primary, got %v", config.Color)
		}
	})

	t.Run("Size modifier application", func(t *testing.T) {
		config := &mockConfig{}
		modifier := SizeLarge
		
		modifier.ApplyToMock(config)
		
		if config.Size != SizeLarge {
			t.Errorf("Expected size to be SizeLarge, got %v", config.Size)
		}
	})

	t.Run("Variant modifier application", func(t *testing.T) {
		config := &mockConfig{}
		modifier := VariantOutline
		
		modifier.ApplyToMock(config)
		
		if config.Variant != VariantOutline {
			t.Errorf("Expected variant to be VariantOutline, got %v", config.Variant)
		}
	})
}

// Mock implementations for testing
type mockComponent struct {
	config *mockConfig
}

func (m *mockComponent) Render(w io.Writer) error {
	// Simple mock render
	_, err := w.Write([]byte("<div></div>"))
	return err
}

func (m *mockComponent) With(modifiers ...any) Component {
	newConfig := *m.config
	for _, mod := range modifiers {
		switch modifier := mod.(type) {
		case MockModifier:
			modifier.ApplyToMock(&newConfig)
		}
	}
	return &mockComponent{config: &newConfig}
}

// Ensure our enums implement the MockModifier interface
var _ MockModifier = Primary
var _ MockModifier = SizeLarge
var _ MockModifier = VariantOutline