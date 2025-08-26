package components

import (
	"io"
	"strings"

	h "maragu.dev/gomponents/html"
	g "maragu.dev/gomponents"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// BadgeComponent represents a badge UI component
type BadgeComponent struct {
	children   []g.Node
	attributes []g.Node
	classes    []string
}

// NewBadge creates a new badge component
func NewBadge(children ...g.Node) *BadgeComponent {
	// Separate attributes from content children
	var attributes []g.Node
	var content []g.Node
	
	for _, child := range children {
		// Check if this is an attribute by trying to render it and seeing if it produces attribute-like output
		var buf strings.Builder
		if err := child.Render(&buf); err == nil {
			output := buf.String()
			// If the output contains '=' it's likely an attribute
			if strings.Contains(output, "=") || strings.Contains(output, "title") || strings.Contains(output, "id") {
				attributes = append(attributes, child)
			} else {
				content = append(content, child)
			}
		} else {
			// If we can't render it, assume it's content
			content = append(content, child)
		}
	}
	
	return &BadgeComponent{
		children:   content,
		attributes: attributes,
		classes:    []string{"badge"}, // Default FlyonUI badge class
	}
}

// With applies modifiers to the badge and returns a new instance
func (b *BadgeComponent) With(modifiers ...any) flyon.Component {
	// Create a new instance to maintain immutability
	newBadge := &BadgeComponent{
		children:   make([]g.Node, len(b.children)),
		attributes: make([]g.Node, len(b.attributes)),
		classes:    make([]string, len(b.classes)),
	}
	
	// Copy children, attributes, and classes
	copy(newBadge.children, b.children)
	copy(newBadge.attributes, b.attributes)
	copy(newBadge.classes, b.classes)
	
	// Apply each modifier
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newBadge.classes = append(newBadge.classes, "badge-"+m.String())
		case flyon.Size:
			newBadge.classes = append(newBadge.classes, "badge-"+m.String())
		case flyon.Variant:
			newBadge.classes = append(newBadge.classes, "badge-"+m.String())
		}
	}
	
	return newBadge
}

// Render implements the gomponents.Node interface
func (b *BadgeComponent) Render(w io.Writer) error {
	// Build the class attribute
	classAttr := strings.Join(b.classes, " ")
	
	// Create the span element with class, attributes, and children
	allNodes := make([]g.Node, 0, len(b.attributes)+len(b.children)+1)
	allNodes = append(allNodes, h.Class(classAttr))
	allNodes = append(allNodes, b.attributes...)
	allNodes = append(allNodes, b.children...)
	
	spanEl := h.Span(allNodes...)
	
	return spanEl.Render(w)
}

// Ensure BadgeComponent implements the required interfaces
var (
	_ flyon.Component = (*BadgeComponent)(nil)
	_ g.Node          = (*BadgeComponent)(nil)
)