package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// FormGroupComponent represents a form group with label, input, description, and error message
type FormGroupComponent struct {
	id          string
	label       string
	description string
	required    bool
	error       string
	input       flyon.Component
	classes     []string
	attributes  map[string]string
}

// NewFormGroup creates a new FormGroup component with default values
func NewFormGroup() *FormGroupComponent {
	return &FormGroupComponent{
		id:         "",
		label:      "",
		description: "",
		required:   false,
		error:      "",
		input:      nil,
		classes:    []string{},
		attributes: make(map[string]string),
	}
}

// WithID sets the ID of the form group
func (fg *FormGroupComponent) WithID(id string) *FormGroupComponent {
	new := fg.copy()
	new.id = id
	return new
}

// WithLabel sets the label text
func (fg *FormGroupComponent) WithLabel(label string) *FormGroupComponent {
	new := fg.copy()
	new.label = label
	return new
}

// WithDescription sets the description text
func (fg *FormGroupComponent) WithDescription(description string) *FormGroupComponent {
	new := fg.copy()
	new.description = description
	return new
}

// WithRequired sets whether the field is required
func (fg *FormGroupComponent) WithRequired(required bool) *FormGroupComponent {
	new := fg.copy()
	new.required = required
	return new
}

// WithError sets the error message
func (fg *FormGroupComponent) WithError(error string) *FormGroupComponent {
	new := fg.copy()
	new.error = error
	return new
}

// WithInput sets the input component
func (fg *FormGroupComponent) WithInput(input flyon.Component) *FormGroupComponent {
	new := fg.copy()
	new.input = input
	return new
}

// WithClasses adds CSS classes to the form group
func (fg *FormGroupComponent) WithClasses(classes ...string) *FormGroupComponent {
	new := fg.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// With adds custom attributes or handles other modifiers
func (fg *FormGroupComponent) With(modifiers ...any) flyon.Component {
	new := fg.copy()
	for i := 0; i < len(modifiers); i++ {
		switch v := modifiers[i].(type) {
		case string:
			// Handle CSS classes
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
		}
	}
	return new
}

// copy creates a deep copy of the FormGroupComponent
func (fg *FormGroupComponent) copy() *FormGroupComponent {
	newClasses := make([]string, len(fg.classes))
	copy(newClasses, fg.classes)

	newAttributes := make(map[string]string)
	for k, v := range fg.attributes {
		newAttributes[k] = v
	}

	return &FormGroupComponent{
		id:         fg.id,
		label:      fg.label,
		description: fg.description,
		required:   fg.required,
		error:      fg.error,
		input:      fg.input,
		classes:    newClasses,
		attributes: newAttributes,
	}
}

// Render generates the HTML for the form group
func (fg *FormGroupComponent) Render(w io.Writer) error {
	// Build class list
	classes := []string{"form-control"}
	classes = append(classes, fg.classes...)

	// Build attributes
	attrs := []g.Node{h.Class(strings.Join(classes, " "))}
	if fg.id != "" {
		attrs = append(attrs, h.ID(fg.id))
	}
	for key, value := range fg.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	// Build children
	var children []g.Node

	// Add label if present
	if fg.label != "" {
		labelChildren := []g.Node{
			h.Class("label-text"),
			g.Text(fg.label),
		}
		if fg.required {
			labelChildren = append(labelChildren, g.Text(" *"))
		}
		children = append(children, h.Label(labelChildren...))
	}

	// Add description if present
	if fg.description != "" {
		children = append(children, h.Label(
			h.Class("label-text-alt"),
			g.Text(fg.description),
		))
	}

	// Add input if present
	if fg.input != nil {
		// Render input to a buffer and then add as raw HTML
		var inputBuffer strings.Builder
		if err := fg.input.Render(&inputBuffer); err != nil {
			return err
		}
		children = append(children, g.Raw(inputBuffer.String()))
	}

	// Add error message if present
	if fg.error != "" {
		children = append(children, h.Label(
			h.Class("label-text-alt text-error"),
			g.Text(fg.error),
		))
	}

	// Render the complete form group
	node := h.Div(append(attrs, children...)...)
	return node.Render(w)
}