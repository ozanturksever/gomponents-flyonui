package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// CheckboxComponent represents a checkbox input with FlyonUI styling
type CheckboxComponent struct {
	id       string
	name     string
	value    string
	checked  bool
	disabled bool
	color    flyon.Color
	size     flyon.Size
	classes  []string
}

// NewCheckbox creates a new checkbox component
func NewCheckbox() *CheckboxComponent {
	return &CheckboxComponent{
		color: flyon.Primary,
		size:  flyon.SizeMedium,
	}
}

// WithID sets the checkbox ID
func (c *CheckboxComponent) WithID(id string) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.id = id
	return newCheckbox
}

// WithName sets the checkbox name attribute
func (c *CheckboxComponent) WithName(name string) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.name = name
	return newCheckbox
}

// WithValue sets the checkbox value attribute
func (c *CheckboxComponent) WithValue(value string) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.value = value
	return newCheckbox
}

// WithChecked sets the checkbox checked state
func (c *CheckboxComponent) WithChecked(checked bool) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.checked = checked
	return newCheckbox
}

// WithDisabled sets the checkbox disabled state
func (c *CheckboxComponent) WithDisabled(disabled bool) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.disabled = disabled
	return newCheckbox
}

// WithColor sets the checkbox color
func (c *CheckboxComponent) WithColor(color flyon.Color) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.color = color
	return newCheckbox
}

// WithSize sets the checkbox size
func (c *CheckboxComponent) WithSize(size flyon.Size) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.size = size
	return newCheckbox
}

// WithClasses adds additional CSS classes
func (c *CheckboxComponent) WithClasses(classes ...string) *CheckboxComponent {
	newCheckbox := c.copy()
	newCheckbox.classes = append(newCheckbox.classes, classes...)
	return newCheckbox
}

// With applies modifiers to the checkbox
func (c *CheckboxComponent) With(modifiers ...any) flyon.Component {
	newCheckbox := c.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newCheckbox.color = m
		case flyon.Size:
			newCheckbox.size = m
		case string:
			newCheckbox.classes = append(newCheckbox.classes, m)
		}
	}
	return newCheckbox
}

// copy creates a deep copy of the checkbox component
func (c *CheckboxComponent) copy() *CheckboxComponent {
	newCheckbox := *c
	newCheckbox.classes = make([]string, len(c.classes))
	copy(newCheckbox.classes, c.classes)
	return &newCheckbox
}

// Render generates the HTML for the checkbox
func (c *CheckboxComponent) Render(w io.Writer) error {
	classes := []string{"checkbox"}
	
	// Add color class
	if c.color != flyon.Primary {
		classes = append(classes, "checkbox-"+c.color.String())
	}
	
	// Add size class
	if c.size != flyon.SizeMedium {
		classes = append(classes, "checkbox-"+c.size.String())
	}
	
	// Add custom classes
	classes = append(classes, c.classes...)
	
	// Build attributes
	attrs := []g.Node{
		h.Type("checkbox"),
		h.Class(strings.Join(classes, " ")),
	}
	
	if c.id != "" {
		attrs = append(attrs, h.ID(c.id))
	}
	
	if c.name != "" {
		attrs = append(attrs, h.Name(c.name))
	}
	
	if c.value != "" {
		attrs = append(attrs, h.Value(c.value))
	}
	
	if c.checked {
		attrs = append(attrs, h.Checked())
	}
	
	if c.disabled {
		attrs = append(attrs, h.Disabled())
	}
	
	return h.Input(attrs...).Render(w)
}