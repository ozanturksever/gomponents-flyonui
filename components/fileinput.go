package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// FileInputComponent represents a file input component
type FileInputComponent struct {
	id         string
	name       string
	accept     string
	multiple   bool
	disabled   bool
	color      flyon.Color
	size       flyon.Size
	colorSet   bool
	sizeSet    bool
	classes    []string
	attributes map[string]string
}

// NewFileInput creates a new file input component
func NewFileInput() *FileInputComponent {
	return &FileInputComponent{
		color:      flyon.Primary,
		size:       flyon.SizeMedium,
		classes:    make([]string, 0),
		attributes: make(map[string]string),
	}
}

// WithID sets the ID of the file input
func (f *FileInputComponent) WithID(id string) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.id = id
	return newFileInput
}

// WithName sets the name of the file input
func (f *FileInputComponent) WithName(name string) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.name = name
	return newFileInput
}

// WithAccept sets the accept attribute of the file input
func (f *FileInputComponent) WithAccept(accept string) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.accept = accept
	return newFileInput
}

// WithMultiple sets whether the file input accepts multiple files
func (f *FileInputComponent) WithMultiple(multiple bool) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.multiple = multiple
	return newFileInput
}

// WithDisabled sets the disabled state of the file input
func (f *FileInputComponent) WithDisabled(disabled bool) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.disabled = disabled
	return newFileInput
}

// WithColor sets the color of the file input
func (f *FileInputComponent) WithColor(color flyon.Color) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.color = color
	newFileInput.colorSet = true
	return newFileInput
}

// WithSize sets the size of the file input
func (f *FileInputComponent) WithSize(size flyon.Size) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.size = size
	newFileInput.sizeSet = true
	return newFileInput
}

// WithClasses adds CSS classes to the file input
func (f *FileInputComponent) WithClasses(classes ...string) *FileInputComponent {
	newFileInput := f.copy()
	newFileInput.classes = append(newFileInput.classes, classes...)
	return newFileInput
}

// With applies modifiers to the file input
func (f *FileInputComponent) With(modifiers ...any) flyon.Component {
	newFileInput := f.copy()
	for i := 0; i < len(modifiers); i++ {
		modifier := modifiers[i]
		switch m := modifier.(type) {
		case flyon.Color:
			newFileInput.color = m
			newFileInput.colorSet = true
		case flyon.Size:
			newFileInput.size = m
			newFileInput.sizeSet = true
		case string:
			// If this is a string followed by another string, treat as key-value pair
			if i+1 < len(modifiers) {
				if value, ok := modifiers[i+1].(string); ok {
					newFileInput.attributes[m] = value
					i++ // Skip the next item as it's the value
					continue
				}
			}
			// Otherwise treat as CSS class
			newFileInput.classes = append(newFileInput.classes, m)
		}
	}
	return newFileInput
}

// copy creates a deep copy of the file input component
func (f *FileInputComponent) copy() *FileInputComponent {
	newFileInput := *f
	newFileInput.classes = make([]string, len(f.classes))
	copy(newFileInput.classes, f.classes)
	newFileInput.attributes = make(map[string]string)
	for k, v := range f.attributes {
		newFileInput.attributes[k] = v
	}
	return &newFileInput
}

// Render renders the file input component
func (f *FileInputComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"file-input"}

	// Add color class if explicitly set
	if f.colorSet {
		classes = append(classes, "file-input-"+f.color.String())
	}

	// Add size class if explicitly set
	if f.sizeSet {
		classes = append(classes, "file-input-"+f.size.String())
	}

	// Add custom classes
	classes = append(classes, f.classes...)

	// Build attributes
	attrs := []g.Node{
		h.Type("file"),
		h.Class(strings.Join(classes, " ")),
	}

	if f.id != "" {
		attrs = append(attrs, h.ID(f.id))
	}

	if f.name != "" {
		attrs = append(attrs, h.Name(f.name))
	}

	if f.accept != "" {
		attrs = append(attrs, h.Accept(f.accept))
	}

	if f.multiple {
		attrs = append(attrs, h.Multiple())
	}

	if f.disabled {
		attrs = append(attrs, h.Disabled())
	}

	// Add custom attributes
	for key, value := range f.attributes {
		attrs = append(attrs, g.Attr(key, value))
	}

	return h.Input(attrs...).Render(w)
}