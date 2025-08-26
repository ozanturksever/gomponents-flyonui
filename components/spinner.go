package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// SpinnerType represents the type of spinner animation
type SpinnerType int

const (
	SpinnerDefault SpinnerType = iota
	SpinnerDots
	SpinnerRing
	SpinnerBall
	SpinnerBars
	SpinnerInfinity
)

// String returns the CSS class suffix for the spinner type
func (st SpinnerType) String() string {
	switch st {
	case SpinnerDots:
		return "dots"
	case SpinnerRing:
		return "ring"
	case SpinnerBall:
		return "ball"
	case SpinnerBars:
		return "bars"
	case SpinnerInfinity:
		return "infinity"
	default:
		return ""
	}
}

// SpinnerComponent represents a loading spinner component
type SpinnerComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
	spinnerType *SpinnerType
}

// NewSpinner creates a new spinner component with the given children
func NewSpinner(children ...g.Node) *SpinnerComponent {
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

	return &SpinnerComponent{
		classes:    []string{"loading"},
		attributes: attributes,
		children:   content,
	}
}

// WithType sets the spinner type
func (s *SpinnerComponent) WithType(spinnerType SpinnerType) *SpinnerComponent {
	s.spinnerType = &spinnerType
	return s
}

// With applies modifiers to the spinner component.
func (s *SpinnerComponent) With(children ...any) flyon.Component {
	for _, child := range children {
		switch c := child.(type) {
		case flyon.Color:
			s.classes = append(s.classes, "loading-"+c.String())
		case flyon.Size:
			s.classes = append(s.classes, "loading-"+c.String())
		case string:
			s.classes = append(s.classes, c)
		case g.Node:
			// Check if it's an attribute or content
			if isAttribute(c) {
				s.attributes = append(s.attributes, c)
			} else {
				s.children = append(s.children, c)
			}
		}
	}
	return s
}



// Render implements the gomponents.Node interface
func (s *SpinnerComponent) Render(w io.Writer) error {
	// Collect all nodes for the element
	nodes := []g.Node{}

	// Add classes
	classes := make([]string, len(s.classes))
	copy(classes, s.classes)

	// Add spinner type class if specified
	if s.spinnerType != nil {
		typeClass := s.spinnerType.String()
		if typeClass != "" {
			classes = append(classes, "loading-"+typeClass)
		}
	}

	// Add class attribute
	if len(classes) > 0 {
		nodes = append(nodes, h.Class(strings.Join(classes, " ")))
	}

	// Add other attributes
	nodes = append(nodes, s.attributes...)

	// Add children
	nodes = append(nodes, s.children...)

	return h.Span(nodes...).Render(w)
}