package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// InputType represents the type of input
type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypePassword InputType = "password"
	InputTypeEmail    InputType = "email"
	InputTypeNumber   InputType = "number"
	InputTypeURL      InputType = "url"
	InputTypeTel      InputType = "tel"
	InputTypeSearch   InputType = "search"
	InputTypeDate     InputType = "date"
	InputTypeTime     InputType = "time"
	InputTypeDatetime InputType = "datetime-local"
)

// InputComponent represents an input field with FlyonUI styling
type InputComponent struct {
	id          string
	name        string
	value       string
	placeholder string
	inputType   InputType
	disabled    bool
	readonly    bool
	required    bool
	color       flyon.Color
	size        flyon.Size
	classes     []string
}

// NewInput creates a new input component
func NewInput() *InputComponent {
	return &InputComponent{
		inputType: InputTypeText,
		color:     flyon.Primary,
		size:      flyon.SizeMedium,
	}
}

// WithID sets the input ID
func (i *InputComponent) WithID(id string) *InputComponent {
	newInput := i.copy()
	newInput.id = id
	return newInput
}

// WithName sets the input name attribute
func (i *InputComponent) WithName(name string) *InputComponent {
	newInput := i.copy()
	newInput.name = name
	return newInput
}

// WithValue sets the input value
func (i *InputComponent) WithValue(value string) *InputComponent {
	newInput := i.copy()
	newInput.value = value
	return newInput
}

// WithPlaceholder sets the input placeholder
func (i *InputComponent) WithPlaceholder(placeholder string) *InputComponent {
	newInput := i.copy()
	newInput.placeholder = placeholder
	return newInput
}

// WithType sets the input type
func (i *InputComponent) WithType(inputType InputType) *InputComponent {
	newInput := i.copy()
	newInput.inputType = inputType
	return newInput
}

// WithDisabled sets the input disabled state
func (i *InputComponent) WithDisabled(disabled bool) *InputComponent {
	newInput := i.copy()
	newInput.disabled = disabled
	return newInput
}

// WithReadonly sets the input readonly state
func (i *InputComponent) WithReadonly(readonly bool) *InputComponent {
	newInput := i.copy()
	newInput.readonly = readonly
	return newInput
}

// WithRequired sets the input required state
func (i *InputComponent) WithRequired(required bool) *InputComponent {
	newInput := i.copy()
	newInput.required = required
	return newInput
}

// WithColor sets the input color
func (i *InputComponent) WithColor(color flyon.Color) *InputComponent {
	newInput := i.copy()
	newInput.color = color
	return newInput
}

// WithSize sets the input size
func (i *InputComponent) WithSize(size flyon.Size) *InputComponent {
	newInput := i.copy()
	newInput.size = size
	return newInput
}

// WithClasses adds additional CSS classes
func (i *InputComponent) WithClasses(classes ...string) *InputComponent {
	newInput := i.copy()
	newInput.classes = append(newInput.classes, classes...)
	return newInput
}

// With applies modifiers to the input
func (i *InputComponent) With(modifiers ...any) flyon.Component {
	newInput := i.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newInput.color = m
		case flyon.Size:
			newInput.size = m
		case string:
			newInput.classes = append(newInput.classes, m)
		}
	}
	return newInput
}

// copy creates a deep copy of the input component
func (i *InputComponent) copy() *InputComponent {
	newInput := *i
	newInput.classes = make([]string, len(i.classes))
	copy(newInput.classes, i.classes)
	return &newInput
}

// Render generates the HTML for the input
func (i *InputComponent) Render(w io.Writer) error {
	classes := []string{"input", "input-bordered"}
	
	// Add color class
	if i.color != flyon.Primary {
		classes = append(classes, "input-"+i.color.String())
	}
	
	// Add size class
	if i.size != flyon.SizeMedium {
		classes = append(classes, "input-"+i.size.String())
	}
	
	// Add custom classes
	classes = append(classes, i.classes...)
	
	// Build attributes
	attrs := []g.Node{
		h.Type(string(i.inputType)),
		h.Class(strings.Join(classes, " ")),
	}
	
	if i.id != "" {
		attrs = append(attrs, h.ID(i.id))
	}
	
	if i.name != "" {
		attrs = append(attrs, h.Name(i.name))
	}
	
	if i.value != "" {
		attrs = append(attrs, h.Value(i.value))
	}
	
	if i.placeholder != "" {
		attrs = append(attrs, h.Placeholder(i.placeholder))
	}
	
	if i.disabled {
		attrs = append(attrs, h.Disabled())
	}
	
	if i.readonly {
		attrs = append(attrs, h.ReadOnly())
	}
	
	if i.required {
		attrs = append(attrs, h.Required())
	}
	
	return h.Input(attrs...).Render(w)
}