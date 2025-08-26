package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// AutocompleteComponent represents an autocomplete input with dropdown options
type AutocompleteComponent struct {
	id         string
	name       string
	placeholder string
	value      string
	disabled   bool
	color      flyon.Color
	size       flyon.Size
	colorSet   bool
	sizeSet    bool
	options    []string
	classes    []string
	attributes map[string]string
}

// NewAutocomplete creates a new Autocomplete component with default values
func NewAutocomplete() *AutocompleteComponent {
	return &AutocompleteComponent{
		id:         "",
		name:       "",
		placeholder: "",
		value:      "",
		disabled:   false,
		color:      flyon.Primary,
		size:       flyon.SizeMedium,
		colorSet:   false,
		sizeSet:    false,
		options:    []string{},
		classes:    []string{},
		attributes: make(map[string]string),
	}
}

// WithID sets the ID of the autocomplete input
func (ac *AutocompleteComponent) WithID(id string) *AutocompleteComponent {
	new := ac.copy()
	new.id = id
	return new
}

// WithName sets the name of the autocomplete input
func (ac *AutocompleteComponent) WithName(name string) *AutocompleteComponent {
	new := ac.copy()
	new.name = name
	return new
}

// WithPlaceholder sets the placeholder text of the autocomplete input
func (ac *AutocompleteComponent) WithPlaceholder(placeholder string) *AutocompleteComponent {
	new := ac.copy()
	new.placeholder = placeholder
	return new
}

// WithValue sets the value of the autocomplete input
func (ac *AutocompleteComponent) WithValue(value string) *AutocompleteComponent {
	new := ac.copy()
	new.value = value
	return new
}

// WithDisabled sets whether the autocomplete input is disabled
func (ac *AutocompleteComponent) WithDisabled(disabled bool) *AutocompleteComponent {
	new := ac.copy()
	new.disabled = disabled
	return new
}

// WithColor sets the color of the autocomplete input
func (ac *AutocompleteComponent) WithColor(color flyon.Color) *AutocompleteComponent {
	new := ac.copy()
	new.color = color
	new.colorSet = true
	return new
}

// WithSize sets the size of the autocomplete input
func (ac *AutocompleteComponent) WithSize(size flyon.Size) *AutocompleteComponent {
	new := ac.copy()
	new.size = size
	new.sizeSet = true
	return new
}

// WithOptions sets the dropdown options for the autocomplete
func (ac *AutocompleteComponent) WithOptions(options ...string) *AutocompleteComponent {
	new := ac.copy()
	new.options = append(new.options, options...)
	return new
}

// WithClasses adds CSS classes to the autocomplete input
func (ac *AutocompleteComponent) WithClasses(classes ...string) *AutocompleteComponent {
	new := ac.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// With adds custom attributes or handles other modifiers
func (ac *AutocompleteComponent) With(modifiers ...any) flyon.Component {
	new := ac.copy()
	for i := 0; i < len(modifiers); i++ {
		switch v := modifiers[i].(type) {
		case flyon.Color:
			new.color = v
			new.colorSet = true
		case flyon.Size:
			new.size = v
			new.sizeSet = true
		case string:
			// Handle CSS classes or key-value pairs
			if i+1 < len(modifiers) {
				if value, ok := modifiers[i+1].(string); ok {
					// Key-value pair for attributes
					new.attributes[v] = value
					i++ // Skip the value in next iteration
					continue
				}
			}
			// Single string is treated as CSS class
			new.classes = append(new.classes, v)
		case bool:
			new.disabled = v
		}
	}
	return new
}

// copy creates a deep copy of the AutocompleteComponent
func (ac *AutocompleteComponent) copy() *AutocompleteComponent {
	newClasses := make([]string, len(ac.classes))
	copy(newClasses, ac.classes)

	newOptions := make([]string, len(ac.options))
	copy(newOptions, ac.options)

	newAttributes := make(map[string]string)
	for k, v := range ac.attributes {
		newAttributes[k] = v
	}

	return &AutocompleteComponent{
		id:         ac.id,
		name:       ac.name,
		placeholder: ac.placeholder,
		value:      ac.value,
		disabled:   ac.disabled,
		color:      ac.color,
		size:       ac.size,
		colorSet:   ac.colorSet,
		sizeSet:    ac.sizeSet,
		options:    newOptions,
		classes:    newClasses,
		attributes: newAttributes,
	}
}

// Render generates the HTML for the autocomplete component
func (ac *AutocompleteComponent) Render(w io.Writer) error {
	// Build class list for input
	classes := []string{"input", "input-bordered"}

	// Add color class if not default
	if ac.colorSet && ac.color != flyon.Primary {
		switch ac.color {
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

	// Add size class if not default
	if ac.sizeSet && ac.size != flyon.SizeMedium {
		switch ac.size {
		case flyon.SizeXS:
			classes = append(classes, "input-xs")
		case flyon.SizeSmall:
			classes = append(classes, "input-sm")
		case flyon.SizeLarge:
			classes = append(classes, "input-lg")
		}
	}

	// Add custom classes
	classes = append(classes, ac.classes...)

	// Build input attributes
	inputAttrs := []g.Node{
		h.Type("text"),
		h.Class(strings.Join(classes, " ")),
	}

	if ac.id != "" {
		inputAttrs = append(inputAttrs, h.ID(ac.id))
	}
	if ac.name != "" {
		inputAttrs = append(inputAttrs, h.Name(ac.name))
	}
	if ac.placeholder != "" {
		inputAttrs = append(inputAttrs, h.Placeholder(ac.placeholder))
	}
	if ac.value != "" {
		inputAttrs = append(inputAttrs, h.Value(ac.value))
	}
	if ac.disabled {
		inputAttrs = append(inputAttrs, h.Disabled())
	}

	// Add custom attributes
	for key, value := range ac.attributes {
		inputAttrs = append(inputAttrs, g.Attr(key, value))
	}

	// Build dropdown options
	var dropdownItems []g.Node
	for _, option := range ac.options {
		dropdownItems = append(dropdownItems, h.Li(
			h.A(
				h.Href("#"),
				g.Text(option),
			),
		))
	}

	// Create the autocomplete structure with dropdown
	node := h.Div(
		h.Class("dropdown"),
		h.Input(inputAttrs...),
		h.Ul(
			h.Class("dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow"),
			g.Group(dropdownItems),
		),
	)

	return node.Render(w)
}