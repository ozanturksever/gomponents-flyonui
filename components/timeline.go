package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// TimelineOrientation represents the orientation of the timeline
type TimelineOrientation int

const (
	TimelineVertical TimelineOrientation = iota
	TimelineHorizontal
)

// TimelineComponent represents a timeline component
type TimelineComponent struct {
	children   []g.Node
	classes    []string
	attributes []g.Node
	compact    bool
	orientation *TimelineOrientation
}

// NewTimeline creates a new timeline component
func NewTimeline(children ...g.Node) *TimelineComponent {
	return &TimelineComponent{
		children: children,
		classes:  []string{"timeline"},
	}
}

// WithOrientation sets the orientation of the timeline
func (t *TimelineComponent) WithOrientation(orientation TimelineOrientation) *TimelineComponent {
	t.orientation = &orientation
	return t
}

// WithCompact sets the timeline to compact mode
func (t *TimelineComponent) WithCompact() *TimelineComponent {
	t.compact = true
	return t
}

// With applies modifiers to the timeline component
func (t *TimelineComponent) With(items ...any) flyon.Component {
	for _, item := range items {
		switch v := item.(type) {
		case flyon.Color:
			t.classes = append(t.classes, "timeline-"+v.String())
		case string:
			t.classes = append(t.classes, v)
		case g.Node:
			if isAttribute(v) {
				t.attributes = append(t.attributes, v)
			} else {
				t.children = append(t.children, v)
			}
		}
	}
	return t
}

// Render renders the timeline component
func (t *TimelineComponent) Render(w io.Writer) error {
	classes := make([]string, len(t.classes))
	copy(classes, t.classes)

	// Add orientation class
	if t.orientation != nil {
		switch *t.orientation {
		case TimelineVertical:
			classes = append(classes, "timeline-vertical")
		case TimelineHorizontal:
			classes = append(classes, "timeline-horizontal")
		}
	}

	// Add compact class
	if t.compact {
		classes = append(classes, "timeline-compact")
	}

	// Build the element
	elements := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	elements = append(elements, t.attributes...)
	elements = append(elements, t.children...)

	return h.Ul(elements...).Render(w)
}