package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// StatsOrientation represents the orientation of the stats
type StatsOrientation int

const (
	StatsVertical StatsOrientation = iota
	StatsHorizontal
)

// StatsComponent represents a stats component
type StatsComponent struct {
	children    []g.Node
	classes     []string
	attributes  []g.Node
	orientation *StatsOrientation
}

// NewStats creates a new stats component
func NewStats(children ...g.Node) *StatsComponent {
	return &StatsComponent{
		children: children,
		classes:  []string{"stats"},
	}
}

// WithOrientation sets the orientation of the stats
func (s *StatsComponent) WithOrientation(orientation StatsOrientation) *StatsComponent {
	s.orientation = &orientation
	return s
}

// With applies modifiers to the stats component
func (s *StatsComponent) With(items ...any) flyon.Component {
	for _, item := range items {
		switch v := item.(type) {
		case flyon.Color:
			s.classes = append(s.classes, "stats-"+v.String())
		case flyon.Size:
			s.classes = append(s.classes, "stats-"+v.String())
		case string:
			s.classes = append(s.classes, v)
		case g.Node:
			if isAttribute(v) {
				s.attributes = append(s.attributes, v)
			} else {
				s.children = append(s.children, v)
			}
		}
	}
	return s
}

// Render renders the stats component
func (s *StatsComponent) Render(w io.Writer) error {
	classes := make([]string, len(s.classes))
	copy(classes, s.classes)

	// Add orientation class
	if s.orientation != nil {
		switch *s.orientation {
		case StatsVertical:
			classes = append(classes, "stats-vertical")
		case StatsHorizontal:
			classes = append(classes, "stats-horizontal")
		}
	}

	// Build the element
	elements := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	elements = append(elements, s.attributes...)
	elements = append(elements, s.children...)

	return h.Div(elements...).Render(w)
}