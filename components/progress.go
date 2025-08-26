package components

import (
	"fmt"
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// ProgressComponent represents a progress element with FlyonUI styling.
type ProgressComponent struct {
	attributes []g.Node
	classes    []string
	value      *int // nil for indeterminate progress
	max        int
}

// NewProgress creates a new progress component with the specified value (0-100).
func NewProgress(value int) *ProgressComponent {
	return &ProgressComponent{
		attributes: []g.Node{},
		classes:    []string{"progress"},
		value:      &value,
		max:        100,
	}
}

// NewProgressWithMax creates a new progress component with custom max value.
func NewProgressWithMax(value, max int) *ProgressComponent {
	return &ProgressComponent{
		attributes: []g.Node{},
		classes:    []string{"progress"},
		value:      &value,
		max:        max,
	}
}

// NewIndeterminateProgress creates a new indeterminate progress component.
func NewIndeterminateProgress() *ProgressComponent {
	return &ProgressComponent{
		attributes: []g.Node{},
		classes:    []string{"progress"},
		value:      nil, // nil indicates indeterminate
		max:        100,
	}
}

// With applies modifiers to the progress and returns a new instance.
func (p *ProgressComponent) With(modifiers ...any) flyon.Component {
	newComponent := *p
	newComponent.attributes = append([]g.Node{}, p.attributes...)
	newComponent.classes = append([]string{}, p.classes...)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newComponent.classes = append(newComponent.classes, "progress-"+m.String())
		case flyon.Size:
			newComponent.classes = append(newComponent.classes, "progress-"+m.String())
		case string:
			// Handle custom CSS classes
			newComponent.classes = append(newComponent.classes, m)
		case g.Node:
			// Handle gomponents attributes
			newComponent.attributes = append(newComponent.attributes, m)
		}
	}

	return &newComponent
}

// Render renders the progress component to the provided writer.
func (p *ProgressComponent) Render(w io.Writer) error {
	// Build all nodes to pass to the element
	var nodes []g.Node

	// Add classes
	nodes = append(nodes, h.Class(strings.Join(p.classes, " ")))

	// Add max attribute
	nodes = append(nodes, g.Attr("max", fmt.Sprintf("%d", p.max)))

	// Add value attribute only if not indeterminate
	if p.value != nil {
		nodes = append(nodes, g.Attr("value", fmt.Sprintf("%d", *p.value)))
	}

	// Add custom attributes
	nodes = append(nodes, p.attributes...)

	// Create the progress element with all nodes
	element := h.Progress(nodes...)

	return element.Render(w)
}