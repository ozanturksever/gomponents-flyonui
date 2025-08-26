package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ValidationType represents the type of validation message
type ValidationType string

const (
	ValidationTypeError   ValidationType = "error"
	ValidationTypeWarning ValidationType = "warning"
	ValidationTypeSuccess ValidationType = "success"
	ValidationTypeInfo    ValidationType = "info"
)

// String returns the string representation of ValidationType
func (vt ValidationType) String() string {
	return string(vt)
}

// FormValidationComponent represents a validation message display component
type FormValidationComponent struct {
	id             string
	message        string
	validationType ValidationType
	visible        bool
	classes        []string
	attributes     map[string]string
}

// NewFormValidation creates a new FormValidation component with default values
func NewFormValidation() *FormValidationComponent {
	return &FormValidationComponent{
		id:             "",
		message:        "",
		validationType: ValidationTypeError,
		visible:        false,
		classes:        []string{},
		attributes:     make(map[string]string),
	}
}

// WithID sets the ID of the validation message
func (fv *FormValidationComponent) WithID(id string) *FormValidationComponent {
	new := fv.copy()
	new.id = id
	return new
}

// WithMessage sets the validation message text
func (fv *FormValidationComponent) WithMessage(message string) *FormValidationComponent {
	new := fv.copy()
	new.message = message
	return new
}

// WithType sets the validation type
func (fv *FormValidationComponent) WithType(validationType ValidationType) *FormValidationComponent {
	new := fv.copy()
	new.validationType = validationType
	return new
}

// WithVisible sets whether the validation message is visible
func (fv *FormValidationComponent) WithVisible(visible bool) *FormValidationComponent {
	new := fv.copy()
	new.visible = visible
	return new
}

// WithClasses adds CSS classes to the validation message
func (fv *FormValidationComponent) WithClasses(classes ...string) *FormValidationComponent {
	new := fv.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// With adds custom attributes or handles other modifiers
func (fv *FormValidationComponent) With(modifiers ...any) flyon.Component {
	new := fv.copy()
	for i := 0; i < len(modifiers); i++ {
		switch v := modifiers[i].(type) {
		case ValidationType:
			new.validationType = v
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
		case bool:
			new.visible = v
		}
	}
	return new
}

// copy creates a deep copy of the FormValidationComponent
func (fv *FormValidationComponent) copy() *FormValidationComponent {
	newClasses := make([]string, len(fv.classes))
	copy(newClasses, fv.classes)

	newAttributes := make(map[string]string)
	for k, v := range fv.attributes {
		newAttributes[k] = v
	}

	return &FormValidationComponent{
		id:             fv.id,
		message:        fv.message,
		validationType: fv.validationType,
		visible:        fv.visible,
		classes:        newClasses,
		attributes:     newAttributes,
	}
}

// Render generates the HTML for the validation message
func (fv *FormValidationComponent) Render(w io.Writer) error {
	// Don't render anything if not visible
	if !fv.visible {
		return nil
	}

	// Build class list
	classes := []string{"label-text-alt"}
	
	// Add validation type class
	switch fv.validationType {
	case ValidationTypeError:
		classes = append(classes, "text-error")
	case ValidationTypeWarning:
		classes = append(classes, "text-warning")
	case ValidationTypeSuccess:
		classes = append(classes, "text-success")
	case ValidationTypeInfo:
		classes = append(classes, "text-info")
	}
	
	// Add custom classes
	classes = append(classes, fv.classes...)

	// Build attributes
	attrs := []g.Node{h.Class(strings.Join(classes, " "))}
	if fv.id != "" {
		attrs = append(attrs, h.ID(fv.id))
	}
	for key, value := range fv.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	// Add message text
	attrs = append(attrs, g.Text(fv.message))

	// Render the validation message
	node := h.Label(attrs...)
	return node.Render(w)
}