//go:build js && wasm

package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// TypographyComponent represents a typography element with FlyonUI styling.
type TypographyComponent struct {
	tag        string
	attributes []g.Node
	classes    []string
	children   []g.Node
}

// NewTypography creates a new typography component with the specified HTML tag.
func NewTypography(tag string, children ...g.Node) *TypographyComponent {
	return &TypographyComponent{
		tag:        tag,
		attributes: []g.Node{},
		classes:    []string{},
		children:   children,
	}
}

// WithWeight applies a font weight modifier to the typography.
func (t *TypographyComponent) WithWeight(weight string) *TypographyComponent {
	newComponent := *t
	newComponent.classes = append([]string{}, t.classes...)
	newComponent.classes = append(newComponent.classes, "font-"+weight)
	return &newComponent
}

// WithAlign applies a text alignment modifier to the typography.
func (t *TypographyComponent) WithAlign(alignment string) *TypographyComponent {
	newComponent := *t
	newComponent.classes = append([]string{}, t.classes...)
	newComponent.classes = append(newComponent.classes, "text-"+alignment)
	return &newComponent
}

// With applies modifiers to the typography and returns a new instance.
func (t *TypographyComponent) With(modifiers ...any) flyon.Component {
	newComponent := *t
	newComponent.attributes = append([]g.Node{}, t.attributes...)
	newComponent.classes = append([]string{}, t.classes...)
	newComponent.children = append([]g.Node{}, t.children...)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newComponent.classes = append(newComponent.classes, "text-"+m.String())
		case flyon.Size:
			newComponent.classes = append(newComponent.classes, "text-"+m.String())
		case string:
			// Handle custom CSS classes
			newComponent.classes = append(newComponent.classes, m)
		case g.Node:
			// Handle gomponents attributes and children
			newComponent.attributes = append(newComponent.attributes, m)
		}
	}

	return &newComponent
}

// Render renders the typography component to the provided writer.
func (t *TypographyComponent) Render(w io.Writer) error {
	// Build all nodes to pass to the element
	var nodes []g.Node

	// Add classes if any
	if len(t.classes) > 0 {
		nodes = append(nodes, h.Class(strings.Join(t.classes, " ")))
	}

	// Add attributes
	nodes = append(nodes, t.attributes...)

	// Add children
	nodes = append(nodes, t.children...)

	// Create the element based on the tag with all nodes
	var element g.Node
	switch t.tag {
	case "h1":
		element = h.H1(nodes...)
	case "h2":
		element = h.H2(nodes...)
	case "h3":
		element = h.H3(nodes...)
	case "h4":
		element = h.H4(nodes...)
	case "h5":
		element = h.H5(nodes...)
	case "h6":
		element = h.H6(nodes...)
	case "p":
		element = h.P(nodes...)
	case "span":
		element = h.Span(nodes...)
	case "strong":
		element = h.Strong(nodes...)
	case "em":
		element = h.Em(nodes...)
	case "small":
		element = h.Small(nodes...)
	case "code":
		element = h.Code(nodes...)
	case "pre":
		element = h.Pre(nodes...)
	default:
		// Fallback to a generic element
		element = h.Div(nodes...)
	}

	return element.Render(w)
}

// Convenience functions for common typography elements

// H1 creates a new h1 typography component.
func H1(children ...g.Node) *TypographyComponent {
	return NewTypography("h1", children...)
}

// H2 creates a new h2 typography component.
func H2(children ...g.Node) *TypographyComponent {
	return NewTypography("h2", children...)
}

// H3 creates a new h3 typography component.
func H3(children ...g.Node) *TypographyComponent {
	return NewTypography("h3", children...)
}

// H4 creates a new h4 typography component.
func H4(children ...g.Node) *TypographyComponent {
	return NewTypography("h4", children...)
}

// H5 creates a new h5 typography component.
func H5(children ...g.Node) *TypographyComponent {
	return NewTypography("h5", children...)
}

// H6 creates a new h6 typography component.
func H6(children ...g.Node) *TypographyComponent {
	return NewTypography("h6", children...)
}

// P creates a new p typography component.
func P(children ...g.Node) *TypographyComponent {
	return NewTypography("p", children...)
}

// Span creates a new span typography component.
func Span(children ...g.Node) *TypographyComponent {
	return NewTypography("span", children...)
}

// Strong creates a new strong typography component.
func Strong(children ...g.Node) *TypographyComponent {
	return NewTypography("strong", children...)
}

// Em creates a new em typography component.
func Em(children ...g.Node) *TypographyComponent {
	return NewTypography("em", children...)
}

// Small creates a new small typography component.
func Small(children ...g.Node) *TypographyComponent {
	return NewTypography("small", children...)
}

// Code creates a new code typography component.
func Code(children ...g.Node) *TypographyComponent {
	return NewTypography("code", children...)
}

// Pre creates a new pre typography component.
func Pre(children ...g.Node) *TypographyComponent {
	return NewTypography("pre", children...)
}