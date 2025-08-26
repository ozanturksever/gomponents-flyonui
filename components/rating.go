package components

import (
	"io"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// RatingComponent represents a FlyonUI rating component.
type RatingComponent struct {
	value    int
	classes  []string
	attrs    []g.Node
	children []g.Node
}

// NewRating creates a new rating component with the specified value.
func NewRating(value int, children ...g.Node) *RatingComponent {
	return &RatingComponent{
		value:    value,
		classes:  []string{"rating"},
		children: children,
	}
}

// With applies modifiers to the rating component.
// It accepts flyon.Color, flyon.Size, string classes, and gomponents.Node attributes.
func (r *RatingComponent) With(items ...any) flyon.Component {
	newRating := &RatingComponent{
		value:    r.value,
		classes:  make([]string, len(r.classes)),
		attrs:    make([]g.Node, len(r.attrs)),
		children: make([]g.Node, len(r.children)),
	}
	copy(newRating.classes, r.classes)
	copy(newRating.attrs, r.attrs)
	copy(newRating.children, r.children)

	for _, item := range items {
		switch v := item.(type) {
		case flyon.Color:
			newRating.classes = append(newRating.classes, "rating-"+v.String())
		case flyon.Size:
			newRating.classes = append(newRating.classes, "rating-"+v.String())
		case string:
			newRating.classes = append(newRating.classes, v)
		case g.Node:
			newRating.attrs = append(newRating.attrs, v)
		}
	}

	return newRating
}

// Render generates the HTML for the rating component.
func (r *RatingComponent) Render(w io.Writer) error {
	attrs := []g.Node{
		h.Class(strings.Join(r.classes, " ")),
		g.Attr("data-rating", strconv.Itoa(r.value)),
	}
	attrs = append(attrs, r.attrs...)
	attrs = append(attrs, r.children...)

	return h.Div(attrs...).Render(w)
}