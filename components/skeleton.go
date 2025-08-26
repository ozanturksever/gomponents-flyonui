package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// SkeletonShape represents the shape variants for skeleton components.
type SkeletonShape int

const (
	SkeletonRectangle SkeletonShape = iota
	SkeletonCircle
	SkeletonText
)

// String returns the CSS class suffix for the skeleton shape.
func (s SkeletonShape) String() string {
	switch s {
	case SkeletonCircle:
		return "circle"
	case SkeletonText:
		return "text"
	case SkeletonRectangle:
		fallthrough
	default:
		return "rectangle"
	}
}

// SkeletonComponent represents a skeleton loading placeholder component.
type SkeletonComponent struct {
	classes   []string
	attributes []g.Node
	children  []g.Node
	shape     *SkeletonShape
	pulse     bool
	wave      bool
}

// NewSkeleton creates a new skeleton component.
func NewSkeleton(children ...g.Node) *SkeletonComponent {
	return &SkeletonComponent{
		classes:   []string{"skeleton"},
		attributes: []g.Node{},
		children:  children,
	}
}

// WithShape sets the shape of the skeleton.
func (s *SkeletonComponent) WithShape(shape SkeletonShape) *SkeletonComponent {
	s.shape = &shape
	return s
}

// WithPulse adds pulse animation to the skeleton.
func (s *SkeletonComponent) WithPulse() *SkeletonComponent {
	s.pulse = true
	return s
}

// WithWave adds wave animation to the skeleton.
func (s *SkeletonComponent) WithWave() *SkeletonComponent {
	s.wave = true
	return s
}

// With applies modifiers to the skeleton component.
func (s *SkeletonComponent) With(children ...any) flyon.Component {
	for _, child := range children {
		switch c := child.(type) {
		case flyon.Color:
			s.classes = append(s.classes, "skeleton-"+c.String())
		case flyon.Size:
			s.classes = append(s.classes, "skeleton-"+c.String())
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

// isAttribute checks if a node is an HTML attribute
func isAttribute(node g.Node) bool {
	// Check if this is an attribute by trying to render it and seeing if it produces attribute-like output
	var buf strings.Builder
	if err := node.Render(&buf); err == nil {
		output := buf.String()
		// If the output contains '=' it's likely an attribute
		return strings.Contains(output, "=") && !strings.Contains(output, "<") && !strings.Contains(output, ">")
	}
	// If we can't render it, assume it's content
	return false
}

// Render implements the gomponents.Node interface
func (s *SkeletonComponent) Render(w io.Writer) error {
	// Collect all nodes for the element
	nodes := []g.Node{}

	// Add classes
	classes := make([]string, len(s.classes))
	copy(classes, s.classes)

	// Add shape class if specified
	if s.shape != nil {
		classes = append(classes, "skeleton-"+s.shape.String())
	}

	// Add animation classes
	if s.pulse {
		classes = append(classes, "skeleton-pulse")
	}
	if s.wave {
		classes = append(classes, "skeleton-wave")
	}

	// Add class attribute
	if len(classes) > 0 {
		nodes = append(nodes, h.Class(strings.Join(classes, " ")))
	}

	// Add other attributes
	nodes = append(nodes, s.attributes...)

	// Add children
	nodes = append(nodes, s.children...)

	return h.Div(nodes...).Render(w)
}