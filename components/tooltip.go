package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// TooltipPosition represents the position of the tooltip
type TooltipPosition int

const (
	TooltipTop TooltipPosition = iota
	TooltipBottom
	TooltipLeft
	TooltipRight
)

// String returns the CSS class suffix for the tooltip position
func (tp TooltipPosition) String() string {
	switch tp {
	case TooltipTop:
		return "top"
	case TooltipBottom:
		return "bottom"
	case TooltipLeft:
		return "left"
	case TooltipRight:
		return "right"
	default:
		return ""
	}
}

// TooltipComponent represents a tooltip component
type TooltipComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
	text       string
	position   *TooltipPosition
	isOpen     bool
}

// NewTooltip creates a new tooltip component with the given text and children
func NewTooltip(text string, children ...g.Node) *TooltipComponent {
	// Separate attributes from content children
	var attributes []g.Node
	var content []g.Node

	for _, child := range children {
		// Check if this is an attribute by trying to render it and seeing if it produces attribute-like output
		if isAttribute(child) {
			attributes = append(attributes, child)
		} else {
			content = append(content, child)
		}
	}

	return &TooltipComponent{
		classes:    []string{"tooltip"},
		attributes: attributes,
		children:   content,
		text:       text,
	}
}

// WithPosition sets the tooltip position
func (t *TooltipComponent) WithPosition(position TooltipPosition) *TooltipComponent {
	t.position = &position
	return t
}

// WithOpen sets the tooltip to be always open
func (t *TooltipComponent) WithOpen() *TooltipComponent {
	t.isOpen = true
	return t
}

// With applies modifiers to the tooltip component.
func (t *TooltipComponent) With(children ...any) flyon.Component {
	for _, child := range children {
		switch c := child.(type) {
		case flyon.Color:
			t.classes = append(t.classes, "tooltip-"+c.String())
		case string:
			t.classes = append(t.classes, c)
		case g.Node:
			// Check if it's an attribute or content
			if isAttribute(c) {
				t.attributes = append(t.attributes, c)
			} else {
				t.children = append(t.children, c)
			}
		}
	}
	return t
}

// Render implements the gomponents.Node interface
func (t *TooltipComponent) Render(w io.Writer) error {
	// Collect all nodes for the element
	nodes := []g.Node{}

	// Add classes
	classes := make([]string, len(t.classes))
	copy(classes, t.classes)

	// Add position class if specified
	if t.position != nil {
		positionClass := t.position.String()
		if positionClass != "" {
			classes = append(classes, "tooltip-"+positionClass)
		}
	}

	// Add open class if specified
	if t.isOpen {
		classes = append(classes, "tooltip-open")
	}

	// Add class attribute
	if len(classes) > 0 {
		nodes = append(nodes, h.Class(strings.Join(classes, " ")))
	}

	// Add data-tip attribute
	if t.text != "" {
		nodes = append(nodes, g.Attr("data-tip", t.text))
	}

	// Add other attributes
	nodes = append(nodes, t.attributes...)

	// Add children
	nodes = append(nodes, t.children...)

	return h.Div(nodes...).Render(w)
}