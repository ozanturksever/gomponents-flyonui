//go:build js && wasm

package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// DividerComponent represents a divider element with FlyonUI styling.
type DividerComponent struct {
	attributes []g.Node
	classes    []string
	children   []g.Node
}

// NewDivider creates a new divider component.
func NewDivider(children ...g.Node) *DividerComponent {
	return &DividerComponent{
		attributes: []g.Node{},
		classes:    []string{"divider"},
		children:   children,
	}
}

// WithOrientation applies an orientation modifier to the divider.
func (d *DividerComponent) WithOrientation(orientation string) *DividerComponent {
	newComponent := *d
	newComponent.classes = append([]string{}, d.classes...)
	newComponent.classes = append(newComponent.classes, "divider-"+orientation)
	return &newComponent
}

// With applies modifiers to the divider and returns a new instance.
func (d *DividerComponent) With(modifiers ...any) flyon.Component {
	newComponent := *d
	newComponent.attributes = append([]g.Node{}, d.attributes...)
	newComponent.classes = append([]string{}, d.classes...)
	newComponent.children = append([]g.Node{}, d.children...)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newComponent.classes = append(newComponent.classes, "divider-"+m.String())
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

// Render renders the divider component to the provided writer.
func (d *DividerComponent) Render(w io.Writer) error {
	// Build all nodes to pass to the element
	var nodes []g.Node

	// Add classes
	nodes = append(nodes, h.Class(strings.Join(d.classes, " ")))

	// Add attributes
	nodes = append(nodes, d.attributes...)

	// Add children
	nodes = append(nodes, d.children...)

	// Create the div element with all nodes
	element := h.Div(nodes...)

	return element.Render(w)
}