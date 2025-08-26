package components

import (
	"io"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// RangeComponent represents a range slider input with FlyonUI styling
type RangeComponent struct {
	id       string
	name     string
	value    float64
	min      float64
	max      float64
	step     float64
	disabled bool
	color    flyon.Color
	size     flyon.Size
	classes  []string
}

// NewRange creates a new range component
func NewRange() *RangeComponent {
	return &RangeComponent{
		min:   0,
		max:   100,
		step:  1,
		value: 50,
		color: flyon.Primary,
		size:  flyon.SizeMedium,
	}
}

// WithID sets the range ID
func (r *RangeComponent) WithID(id string) *RangeComponent {
	newRange := r.copy()
	newRange.id = id
	return newRange
}

// WithName sets the range name attribute
func (r *RangeComponent) WithName(name string) *RangeComponent {
	newRange := r.copy()
	newRange.name = name
	return newRange
}

// WithValue sets the range value
func (r *RangeComponent) WithValue(value float64) *RangeComponent {
	newRange := r.copy()
	newRange.value = value
	return newRange
}

// WithMin sets the minimum value
func (r *RangeComponent) WithMin(min float64) *RangeComponent {
	newRange := r.copy()
	newRange.min = min
	return newRange
}

// WithMax sets the maximum value
func (r *RangeComponent) WithMax(max float64) *RangeComponent {
	newRange := r.copy()
	newRange.max = max
	return newRange
}

// WithStep sets the step value
func (r *RangeComponent) WithStep(step float64) *RangeComponent {
	newRange := r.copy()
	newRange.step = step
	return newRange
}

// WithDisabled sets the range disabled state
func (r *RangeComponent) WithDisabled(disabled bool) *RangeComponent {
	newRange := r.copy()
	newRange.disabled = disabled
	return newRange
}

// WithColor sets the range color
func (r *RangeComponent) WithColor(color flyon.Color) *RangeComponent {
	newRange := r.copy()
	newRange.color = color
	return newRange
}

// WithSize sets the range size
func (r *RangeComponent) WithSize(size flyon.Size) *RangeComponent {
	newRange := r.copy()
	newRange.size = size
	return newRange
}

// WithClasses adds additional CSS classes
func (r *RangeComponent) WithClasses(classes ...string) *RangeComponent {
	newRange := r.copy()
	newRange.classes = append(newRange.classes, classes...)
	return newRange
}

// With applies modifiers to the range
func (r *RangeComponent) With(modifiers ...any) flyon.Component {
	newRange := r.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newRange.color = m
		case flyon.Size:
			newRange.size = m
		case string:
			newRange.classes = append(newRange.classes, m)
		}
	}
	return newRange
}

// copy creates a deep copy of the range component
func (r *RangeComponent) copy() *RangeComponent {
	newRange := *r
	newRange.classes = make([]string, len(r.classes))
	copy(newRange.classes, r.classes)
	return &newRange
}

// Render generates the HTML for the range
func (r *RangeComponent) Render(w io.Writer) error {
	classes := []string{"range"}
	
	// Add color class
	if r.color != flyon.Primary {
		classes = append(classes, "range-"+r.color.String())
	}
	
	// Add size class
	if r.size != flyon.SizeMedium {
		classes = append(classes, "range-"+r.size.String())
	}
	
	// Add custom classes
	classes = append(classes, r.classes...)
	
	// Build attributes
	attrs := []g.Node{
		h.Type("range"),
		h.Class(strings.Join(classes, " ")),
		g.Attr("min", strconv.FormatFloat(r.min, 'f', -1, 64)),
		g.Attr("max", strconv.FormatFloat(r.max, 'f', -1, 64)),
		g.Attr("step", strconv.FormatFloat(r.step, 'f', -1, 64)),
		h.Value(strconv.FormatFloat(r.value, 'f', -1, 64)),
	}
	
	if r.id != "" {
		attrs = append(attrs, h.ID(r.id))
	}
	
	if r.name != "" {
		attrs = append(attrs, h.Name(r.name))
	}
	
	if r.disabled {
		attrs = append(attrs, h.Disabled())
	}
	
	return h.Input(attrs...).Render(w)
}