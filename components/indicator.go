package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// IndicatorPosition represents the position of the indicator
type IndicatorPosition int

const (
	IndicatorTopStart IndicatorPosition = iota
	IndicatorTopCenter
	IndicatorTopEnd
	IndicatorMiddleStart
	IndicatorMiddleCenter
	IndicatorMiddleEnd
	IndicatorBottomStart
	IndicatorBottomCenter
	IndicatorBottomEnd
)

// IndicatorComponent represents a FlyonUI indicator component.
type IndicatorComponent struct {
	position   *IndicatorPosition
	classes    []string
	attrs      []g.Node
	children   []g.Node
}

// NewIndicator creates a new indicator component.
func NewIndicator(children ...g.Node) *IndicatorComponent {
	return &IndicatorComponent{
		classes:  []string{"indicator"},
		children: children,
	}
}

// WithPosition sets the position of the indicator.
func (i *IndicatorComponent) WithPosition(position IndicatorPosition) *IndicatorComponent {
	newIndicator := &IndicatorComponent{
		position: &position,
		classes:  make([]string, len(i.classes)),
		attrs:    make([]g.Node, len(i.attrs)),
		children: make([]g.Node, len(i.children)),
	}
	copy(newIndicator.classes, i.classes)
	copy(newIndicator.attrs, i.attrs)
	copy(newIndicator.children, i.children)

	return newIndicator
}

// With applies modifiers to the indicator component.
// It accepts flyon.Color, flyon.Size, string classes, and gomponents.Node attributes.
func (i *IndicatorComponent) With(items ...any) flyon.Component {
	newIndicator := &IndicatorComponent{
		position: i.position,
		classes:  make([]string, len(i.classes)),
		attrs:    make([]g.Node, len(i.attrs)),
		children: make([]g.Node, len(i.children)),
	}
	copy(newIndicator.classes, i.classes)
	copy(newIndicator.attrs, i.attrs)
	copy(newIndicator.children, i.children)

	for _, item := range items {
		switch v := item.(type) {
		case flyon.Color:
			newIndicator.classes = append(newIndicator.classes, "indicator-"+v.String())
		case flyon.Size:
			newIndicator.classes = append(newIndicator.classes, "indicator-"+v.String())
		case string:
			newIndicator.classes = append(newIndicator.classes, v)
		case g.Node:
			newIndicator.attrs = append(newIndicator.attrs, v)
		}
	}

	return newIndicator
}

// Render generates the HTML for the indicator component.
func (i *IndicatorComponent) Render(w io.Writer) error {
	classes := make([]string, len(i.classes))
	copy(classes, i.classes)

	// Add position class
	if i.position != nil {
		switch *i.position {
		case IndicatorTopStart:
			classes = append(classes, "indicator-top-start")
		case IndicatorTopCenter:
			classes = append(classes, "indicator-top-center")
		case IndicatorTopEnd:
			classes = append(classes, "indicator-top-end")
		case IndicatorMiddleStart:
			classes = append(classes, "indicator-middle-start")
		case IndicatorMiddleCenter:
			classes = append(classes, "indicator-middle-center")
		case IndicatorMiddleEnd:
			classes = append(classes, "indicator-middle-end")
		case IndicatorBottomStart:
			classes = append(classes, "indicator-bottom-start")
		case IndicatorBottomCenter:
			classes = append(classes, "indicator-bottom-center")
		case IndicatorBottomEnd:
			classes = append(classes, "indicator-bottom-end")
		}
	}

	// Build the element
	elements := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	elements = append(elements, i.attrs...)
	elements = append(elements, i.children...)

	return h.Span(elements...).Render(w)
}