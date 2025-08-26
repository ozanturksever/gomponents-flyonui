package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// ToggleComponent represents a toggle/switch input component
type ToggleComponent struct {
	id         string
	name       string
	value      string
	checked    bool
	disabled   bool
	color      flyon.Color
	size       flyon.Size
	colorSet   bool
	sizeSet    bool
	classes    []string
	attributes map[string]string
}

// NewToggle creates a new toggle component with default values
func NewToggle() *ToggleComponent {
	return &ToggleComponent{
		color:      flyon.Primary,
		size:       flyon.SizeMedium,
		colorSet:   false,
		sizeSet:    false,
		attributes: make(map[string]string),
	}
}

// WithID sets the ID attribute
func (t *ToggleComponent) WithID(id string) *ToggleComponent {
	new := *t
	new.id = id
	return &new
}

// WithName sets the name attribute
func (t *ToggleComponent) WithName(name string) *ToggleComponent {
	new := *t
	new.name = name
	return &new
}

// WithValue sets the value attribute
func (t *ToggleComponent) WithValue(value string) *ToggleComponent {
	new := *t
	new.value = value
	return &new
}

// WithChecked sets the checked state
func (t *ToggleComponent) WithChecked(checked bool) *ToggleComponent {
	new := *t
	new.checked = checked
	return &new
}

// WithDisabled sets the disabled state
func (t *ToggleComponent) WithDisabled(disabled bool) *ToggleComponent {
	new := *t
	new.disabled = disabled
	return &new
}

// WithColor sets the toggle color
func (t *ToggleComponent) WithColor(color flyon.Color) *ToggleComponent {
	new := *t
	new.attributes = make(map[string]string)
	for k, v := range t.attributes {
		new.attributes[k] = v
	}
	new.color = color
	new.colorSet = true
	return &new
}

// WithSize sets the toggle size
func (t *ToggleComponent) WithSize(size flyon.Size) *ToggleComponent {
	new := *t
	new.attributes = make(map[string]string)
	for k, v := range t.attributes {
		new.attributes[k] = v
	}
	new.size = size
	new.sizeSet = true
	return &new
}

// WithClasses adds CSS classes
func (t *ToggleComponent) WithClasses(classes ...string) *ToggleComponent {
	new := *t
	new.attributes = make(map[string]string)
	for k, v := range t.attributes {
		new.attributes[k] = v
	}
	new.classes = append([]string{}, t.classes...)
	new.classes = append(new.classes, classes...)
	return &new
}

// With applies modifiers to the toggle
func (t *ToggleComponent) With(modifiers ...any) flyon.Component {
	newToggle := t.copy()
	for i := 0; i < len(modifiers); i++ {
		modifier := modifiers[i]
		switch m := modifier.(type) {
		case flyon.Color:
			newToggle.color = m
			newToggle.colorSet = true
		case flyon.Size:
			newToggle.size = m
			newToggle.sizeSet = true
		case string:
			// If this is a string followed by another string, treat as key-value pair
			if i+1 < len(modifiers) {
				if value, ok := modifiers[i+1].(string); ok {
					newToggle.attributes[m] = value
					i++ // Skip the next item as it's the value
					continue
				}
			}
			// Otherwise treat as CSS class
			newToggle.classes = append(newToggle.classes, m)
		}
	}
	return newToggle
}

// copy creates a deep copy of the toggle component
func (t *ToggleComponent) copy() *ToggleComponent {
	newToggle := *t
	newToggle.classes = make([]string, len(t.classes))
	copy(newToggle.classes, t.classes)
	newToggle.attributes = make(map[string]string)
	for k, v := range t.attributes {
		newToggle.attributes[k] = v
	}
	return &newToggle
}

// Render renders the toggle component
func (t *ToggleComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"toggle"}

	// Add color class if explicitly set
	if t.colorSet {
		classes = append(classes, "toggle-"+t.color.String())
	}

	// Add size class if explicitly set
	if t.sizeSet {
		classes = append(classes, "toggle-"+t.size.String())
	}

	// Add custom classes
	classes = append(classes, t.classes...)

	// Build attributes
	attrs := []g.Node{
		h.Type("checkbox"),
		h.Class(strings.Join(classes, " ")),
	}

	if t.id != "" {
		attrs = append(attrs, h.ID(t.id))
	}

	if t.name != "" {
		attrs = append(attrs, h.Name(t.name))
	}

	if t.value != "" {
		attrs = append(attrs, h.Value(t.value))
	}

	if t.checked {
		attrs = append(attrs, h.Checked())
	}

	if t.disabled {
		attrs = append(attrs, h.Disabled())
	}

	// Add custom attributes
	for key, value := range t.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	return h.Input(attrs...).Render(w)
}