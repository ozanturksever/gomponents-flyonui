package components

import (
	"io"
	"strings"
	"time"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// DatePickerComponent represents a date picker input component
type DatePickerComponent struct {
	id         string
	name       string
	placeholder string
	value      time.Time
	disabled   bool
	color      flyon.Color
	size       flyon.Size
	colorSet   bool
	sizeSet    bool
	format     string
	minDate    time.Time
	maxDate    time.Time
	classes    []string
	attributes map[string]string
}

// NewDatePicker creates a new DatePicker component with default values
func NewDatePicker() *DatePickerComponent {
	return &DatePickerComponent{
		color:      flyon.Primary,
		size:       flyon.SizeMedium,
		format:     "2006-01-02", // ISO format for HTML date input
		classes:    []string{},
		attributes: make(map[string]string),
	}
}

// WithID sets the ID attribute
func (d *DatePickerComponent) WithID(id string) *DatePickerComponent {
	new := d.copy()
	new.id = id
	return new
}

// WithName sets the name attribute
func (d *DatePickerComponent) WithName(name string) *DatePickerComponent {
	new := d.copy()
	new.name = name
	return new
}

// WithPlaceholder sets the placeholder text
func (d *DatePickerComponent) WithPlaceholder(placeholder string) *DatePickerComponent {
	new := d.copy()
	new.placeholder = placeholder
	return new
}

// WithValue sets the date value
func (d *DatePickerComponent) WithValue(value time.Time) *DatePickerComponent {
	new := d.copy()
	new.value = value
	return new
}

// WithDisabled sets the disabled state
func (d *DatePickerComponent) WithDisabled(disabled bool) *DatePickerComponent {
	new := d.copy()
	new.disabled = disabled
	return new
}

// WithColor sets the color theme
func (d *DatePickerComponent) WithColor(color flyon.Color) *DatePickerComponent {
	new := d.copy()
	new.color = color
	new.colorSet = true
	return new
}

// WithSize sets the size
func (d *DatePickerComponent) WithSize(size flyon.Size) *DatePickerComponent {
	new := d.copy()
	new.size = size
	new.sizeSet = true
	return new
}

// WithFormat sets the date format
func (d *DatePickerComponent) WithFormat(format string) *DatePickerComponent {
	new := d.copy()
	new.format = format
	return new
}

// WithMinDate sets the minimum selectable date
func (d *DatePickerComponent) WithMinDate(minDate time.Time) *DatePickerComponent {
	new := d.copy()
	new.minDate = minDate
	return new
}

// WithMaxDate sets the maximum selectable date
func (d *DatePickerComponent) WithMaxDate(maxDate time.Time) *DatePickerComponent {
	new := d.copy()
	new.maxDate = maxDate
	return new
}

// WithClasses adds custom CSS classes
func (d *DatePickerComponent) WithClasses(classes ...string) *DatePickerComponent {
	new := d.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// WithAttribute sets a custom attribute
func (d *DatePickerComponent) WithAttribute(key, value string) *DatePickerComponent {
	new := d.copy()
	new.attributes[key] = value
	return new
}

// With applies modifiers to the date picker
func (d *DatePickerComponent) With(modifiers ...any) flyon.Component {
	new := d.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			new.color = m
			new.colorSet = true
		case flyon.Size:
			new.size = m
			new.sizeSet = true
		case string:
			new.classes = append(new.classes, m)
		}
	}
	return new
}

// copy creates a deep copy of the component for immutability
func (d *DatePickerComponent) copy() *DatePickerComponent {
	newClasses := make([]string, len(d.classes))
	copy(newClasses, d.classes)

	newAttributes := make(map[string]string)
	for k, v := range d.attributes {
		newAttributes[k] = v
	}

	return &DatePickerComponent{
		id:         d.id,
		name:       d.name,
		placeholder: d.placeholder,
		value:      d.value,
		disabled:   d.disabled,
		color:      d.color,
		size:       d.size,
		colorSet:   d.colorSet,
		sizeSet:    d.sizeSet,
		format:     d.format,
		minDate:    d.minDate,
		maxDate:    d.maxDate,
		classes:    newClasses,
		attributes: newAttributes,
	}
}

// Render generates the HTML for the date picker component
func (d *DatePickerComponent) Render(w io.Writer) error {
	classes := []string{"input", "input-bordered"}

	// Add color class if set
	if d.colorSet {
		switch d.color {
		case flyon.Primary:
			classes = append(classes, "input-primary")
		case flyon.Secondary:
			classes = append(classes, "input-secondary")
		case flyon.Success:
			classes = append(classes, "input-success")
		case flyon.Warning:
			classes = append(classes, "input-warning")
		case flyon.Error:
			classes = append(classes, "input-error")
		case flyon.Info:
			classes = append(classes, "input-info")
		case flyon.Neutral:
			classes = append(classes, "input-neutral")
		}
	}

	// Add size class if set
	if d.sizeSet {
		switch d.size {
		case flyon.SizeXS:
			classes = append(classes, "input-xs")
		case flyon.SizeSmall:
			classes = append(classes, "input-sm")
		case flyon.SizeMedium:
			// Medium is default, no additional class needed
		case flyon.SizeLarge:
			classes = append(classes, "input-lg")
		}
	}

	// Add custom classes
	classes = append(classes, d.classes...)

	// Build attributes
	attrs := []g.Node{
		h.Type("date"),
		h.Class(strings.Join(classes, " ")),
	}

	if d.id != "" {
		attrs = append(attrs, h.ID(d.id))
	}
	if d.name != "" {
		attrs = append(attrs, h.Name(d.name))
	}
	if d.placeholder != "" {
		attrs = append(attrs, h.Placeholder(d.placeholder))
	}
	if !d.value.IsZero() {
		attrs = append(attrs, h.Value(d.value.Format("2006-01-02")))
	}
	if d.disabled {
		attrs = append(attrs, h.Disabled())
	}
	if !d.minDate.IsZero() {
		attrs = append(attrs, h.Min(d.minDate.Format("2006-01-02")))
	}
	if !d.maxDate.IsZero() {
		attrs = append(attrs, h.Max(d.maxDate.Format("2006-01-02")))
	}

	// Add custom attributes
	for key, value := range d.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	return h.Input(attrs...).Render(w)
}