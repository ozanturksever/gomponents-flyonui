package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// RadioComponent represents a radio input with FlyonUI styling
type RadioComponent struct {
	id       string
	name     string
	value    string
	checked  bool
	disabled bool
	color    flyon.Color
	size     flyon.Size
	classes  []string
}

// NewRadio creates a new radio component
func NewRadio() *RadioComponent {
	return &RadioComponent{
		color: flyon.Primary,
		size:  flyon.SizeMedium,
	}
}

// WithID sets the radio ID
func (r *RadioComponent) WithID(id string) *RadioComponent {
	newRadio := r.copy()
	newRadio.id = id
	return newRadio
}

// WithName sets the radio name attribute
func (r *RadioComponent) WithName(name string) *RadioComponent {
	newRadio := r.copy()
	newRadio.name = name
	return newRadio
}

// WithValue sets the radio value attribute
func (r *RadioComponent) WithValue(value string) *RadioComponent {
	newRadio := r.copy()
	newRadio.value = value
	return newRadio
}

// WithChecked sets the radio checked state
func (r *RadioComponent) WithChecked(checked bool) *RadioComponent {
	newRadio := r.copy()
	newRadio.checked = checked
	return newRadio
}

// WithDisabled sets the radio disabled state
func (r *RadioComponent) WithDisabled(disabled bool) *RadioComponent {
	newRadio := r.copy()
	newRadio.disabled = disabled
	return newRadio
}

// WithColor sets the radio color
func (r *RadioComponent) WithColor(color flyon.Color) *RadioComponent {
	newRadio := r.copy()
	newRadio.color = color
	return newRadio
}

// WithSize sets the radio size
func (r *RadioComponent) WithSize(size flyon.Size) *RadioComponent {
	newRadio := r.copy()
	newRadio.size = size
	return newRadio
}

// WithClasses adds additional CSS classes
func (r *RadioComponent) WithClasses(classes ...string) *RadioComponent {
	newRadio := r.copy()
	newRadio.classes = append(newRadio.classes, classes...)
	return newRadio
}

// With applies modifiers to the radio
func (r *RadioComponent) With(modifiers ...any) flyon.Component {
	newRadio := r.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newRadio.color = m
		case flyon.Size:
			newRadio.size = m
		case string:
			newRadio.classes = append(newRadio.classes, m)
		}
	}
	return newRadio
}

// copy creates a deep copy of the radio component
func (r *RadioComponent) copy() *RadioComponent {
	newRadio := *r
	newRadio.classes = make([]string, len(r.classes))
	copy(newRadio.classes, r.classes)
	return &newRadio
}

// Render generates the HTML for the radio
func (r *RadioComponent) Render(w io.Writer) error {
	classes := []string{"radio"}
	
	// Add color class
	if r.color != flyon.Primary {
		classes = append(classes, "radio-"+r.color.String())
	}
	
	// Add size class
	if r.size != flyon.SizeMedium {
		classes = append(classes, "radio-"+r.size.String())
	}
	
	// Add custom classes
	classes = append(classes, r.classes...)
	
	// Build attributes
	attrs := []g.Node{
		h.Type("radio"),
		h.Class(strings.Join(classes, " ")),
	}
	
	if r.id != "" {
		attrs = append(attrs, h.ID(r.id))
	}
	
	if r.name != "" {
		attrs = append(attrs, h.Name(r.name))
	}
	
	if r.value != "" {
		attrs = append(attrs, h.Value(r.value))
	}
	
	if r.checked {
		attrs = append(attrs, h.Checked())
	}
	
	if r.disabled {
		attrs = append(attrs, h.Disabled())
	}
	
	return h.Input(attrs...).Render(w)
}