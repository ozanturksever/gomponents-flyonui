package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
	h "maragu.dev/gomponents/html"
)

// ComboboxOption represents an option in the combobox dropdown
type ComboboxOption struct {
	Value    string
	Label    string
	Disabled bool
}

// ComboboxComponent represents a combobox input with dropdown functionality
type ComboboxComponent struct {
	id         string
	name       string
	placeholder string
	value      string
	disabled   bool
	color      flyon.Color
	size       flyon.Size
	colorSet   bool
	sizeSet    bool
	options    []ComboboxOption
	classes    []string
	attributes map[string]string
}

// NewCombobox creates a new combobox component
func NewCombobox() *ComboboxComponent {
	return &ComboboxComponent{
		color:      flyon.Primary,
		size:       flyon.SizeMedium,
		options:    make([]ComboboxOption, 0),
		classes:    make([]string, 0),
		attributes: make(map[string]string),
	}
}

// WithID sets the ID attribute
func (c *ComboboxComponent) WithID(id string) *ComboboxComponent {
	new := c.copy()
	new.id = id
	return new
}

// WithName sets the name attribute
func (c *ComboboxComponent) WithName(name string) *ComboboxComponent {
	new := c.copy()
	new.name = name
	return new
}

// WithPlaceholder sets the placeholder text
func (c *ComboboxComponent) WithPlaceholder(placeholder string) *ComboboxComponent {
	new := c.copy()
	new.placeholder = placeholder
	return new
}

// WithValue sets the current value
func (c *ComboboxComponent) WithValue(value string) *ComboboxComponent {
	new := c.copy()
	new.value = value
	return new
}

// WithDisabled sets the disabled state
func (c *ComboboxComponent) WithDisabled(disabled bool) *ComboboxComponent {
	new := c.copy()
	new.disabled = disabled
	return new
}

// WithColor sets the color theme
func (c *ComboboxComponent) WithColor(color flyon.Color) *ComboboxComponent {
	new := c.copy()
	new.color = color
	new.colorSet = true
	return new
}

// WithSize sets the size
func (c *ComboboxComponent) WithSize(size flyon.Size) *ComboboxComponent {
	new := c.copy()
	new.size = size
	new.sizeSet = true
	return new
}

// WithOptions sets the dropdown options
func (c *ComboboxComponent) WithOptions(options []ComboboxOption) *ComboboxComponent {
	new := c.copy()
	new.options = make([]ComboboxOption, len(options))
	copy(new.options, options)
	return new
}

// WithClasses adds custom CSS classes
func (c *ComboboxComponent) WithClasses(classes ...string) *ComboboxComponent {
	new := c.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// WithAttribute sets a custom attribute
func (c *ComboboxComponent) WithAttribute(key, value string) *ComboboxComponent {
	new := c.copy()
	new.attributes[key] = value
	return new
}

// With applies modifiers to the combobox
func (c *ComboboxComponent) With(modifiers ...any) flyon.Component {
	new := c.copy()
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

// copy creates a deep copy of the component
func (c *ComboboxComponent) copy() *ComboboxComponent {
	new := &ComboboxComponent{
		id:         c.id,
		name:       c.name,
		placeholder: c.placeholder,
		value:      c.value,
		disabled:   c.disabled,
		color:      c.color,
		size:       c.size,
		colorSet:   c.colorSet,
		sizeSet:    c.sizeSet,
		options:    make([]ComboboxOption, len(c.options)),
		classes:    make([]string, len(c.classes)),
		attributes: make(map[string]string),
	}

	copy(new.options, c.options)
	copy(new.classes, c.classes)
	for k, v := range c.attributes {
		new.attributes[k] = v
	}

	return new
}

// Render generates the HTML for the combobox component
func (c *ComboboxComponent) Render(w io.Writer) error {
	// Build dropdown container classes
	dropdownClasses := []string{"dropdown"}

	// Build input classes
	inputClasses := []string{"input", "input-bordered"}

	// Add color class if set
	if c.colorSet {
		inputClasses = append(inputClasses, "input-"+c.color.String())
	}

	// Add size class if set
	if c.sizeSet {
		inputClasses = append(inputClasses, "input-"+c.size.String())
	}

	// Add custom classes
	inputClasses = append(inputClasses, c.classes...)

	// Build input attributes
	attrs := []g.Node{
		h.Type("text"),
		h.Class(strings.Join(inputClasses, " ")),
	}

	if c.id != "" {
		attrs = append(attrs, h.ID(c.id))
	}
	if c.name != "" {
		attrs = append(attrs, h.Name(c.name))
	}
	if c.placeholder != "" {
		attrs = append(attrs, h.Placeholder(c.placeholder))
	}
	if c.value != "" {
		attrs = append(attrs, h.Value(c.value))
	}
	if c.disabled {
		attrs = append(attrs, h.Disabled())
	}

	// Add custom attributes
	for key, value := range c.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	// Build dropdown content
	dropdownContent := []g.Node{
		h.Input(attrs...),
	}

	// Add dropdown menu if options exist
	if len(c.options) > 0 {
		menuItems := make([]g.Node, 0, len(c.options))
		for _, option := range c.options {
			itemAttrs := []g.Node{
				h.Class("dropdown-item"),
				g.Attr("data-value", option.Value),
			}
			if option.Disabled {
				itemAttrs = append(itemAttrs, h.Disabled())
			}
			menuItems = append(menuItems, h.Li(append(itemAttrs, g.Text(option.Label))...))
		}

		dropdownMenu := h.Ul(
			h.Class("dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow"),
			g.Group(menuItems),
		)
		dropdownContent = append(dropdownContent, dropdownMenu)
	}

	return h.Div(
		h.Class(strings.Join(dropdownClasses, " ")),
		g.Group(dropdownContent),
	).Render(w)
}