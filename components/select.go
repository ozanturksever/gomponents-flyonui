package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// SelectComponent represents a select dropdown component
type SelectComponent struct {
	children   []gomponents.Node
	attributes []gomponents.Node
	modifiers  []flyon.Modifier
}

// NewSelect creates a new select component with the given children (options and attributes)
func NewSelect(children ...gomponents.Node) *SelectComponent {
	var attributes []gomponents.Node
	var content []gomponents.Node

	// Separate attributes from content children
	for _, child := range children {
		var buf strings.Builder
		child.Render(&buf)
		renderedChild := buf.String()

		// Check if this looks like an attribute (contains = or is "disabled")
		if strings.Contains(renderedChild, "=") || renderedChild == "disabled" {
			attributes = append(attributes, child)
		} else {
			content = append(content, child)
		}
	}

	return &SelectComponent{
		children:   content,
		attributes: attributes,
		modifiers:  []flyon.Modifier{},
	}
}

// With applies modifiers to the select component
func (s *SelectComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	s.modifiers = append(s.modifiers, modifiers...)
	return s
}

// Render outputs the select component as HTML
func (s *SelectComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"select"}

	// Apply modifiers
	for _, modifier := range s.modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			classes = append(classes, "select-"+m.String())
		case flyon.Size:
			classes = append(classes, "select-"+m.String())
		case flyon.Variant:
			classes = append(classes, "select-"+m.String())
		}
	}

	// Create the select element with class and attributes
	selectAttrs := []gomponents.Node{
		h.Class(strings.Join(classes, " ")),
	}
	selectAttrs = append(selectAttrs, s.attributes...)
	selectAttrs = append(selectAttrs, s.children...)

	return h.Select(selectAttrs...).Render(w)
}