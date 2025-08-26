package components

import (
	"io"
	"strconv"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TextareaComponent represents a textarea form element with FlyonUI styling
type TextareaComponent struct {
	id          string
	name        string
	value       string
	placeholder string
	rows        int
	cols        int
	disabled    bool
	readonly    bool
	required    bool
	color       flyon.Color
	size        flyon.Size
	colorSet    bool
	sizeSet     bool
	classes     []string
}

// NewTextarea creates a new textarea component
func NewTextarea() *TextareaComponent {
	return &TextareaComponent{
		rows:  3,
		cols:  50,
		color: flyon.Primary,
		size:  flyon.SizeMedium,
	}
}

// WithID sets the ID attribute
func (t *TextareaComponent) WithID(id string) *TextareaComponent {
	new := t.copy()
	new.id = id
	return new
}

// WithName sets the name attribute
func (t *TextareaComponent) WithName(name string) *TextareaComponent {
	new := t.copy()
	new.name = name
	return new
}

// WithValue sets the value (content) of the textarea
func (t *TextareaComponent) WithValue(value string) *TextareaComponent {
	new := t.copy()
	new.value = value
	return new
}

// WithPlaceholder sets the placeholder text
func (t *TextareaComponent) WithPlaceholder(placeholder string) *TextareaComponent {
	new := t.copy()
	new.placeholder = placeholder
	return new
}

// WithRows sets the number of visible text lines
func (t *TextareaComponent) WithRows(rows int) *TextareaComponent {
	new := t.copy()
	new.rows = rows
	return new
}

// WithCols sets the visible width of the text control
func (t *TextareaComponent) WithCols(cols int) *TextareaComponent {
	new := t.copy()
	new.cols = cols
	return new
}

// WithDisabled sets the disabled state
func (t *TextareaComponent) WithDisabled(disabled bool) *TextareaComponent {
	new := t.copy()
	new.disabled = disabled
	return new
}

// WithReadonly sets the readonly state
func (t *TextareaComponent) WithReadonly(readonly bool) *TextareaComponent {
	new := t.copy()
	new.readonly = readonly
	return new
}

// WithRequired sets the required state
func (t *TextareaComponent) WithRequired(required bool) *TextareaComponent {
	new := t.copy()
	new.required = required
	return new
}

// WithColor sets the color variant
func (t *TextareaComponent) WithColor(color flyon.Color) *TextareaComponent {
	new := t.copy()
	new.color = color
	new.colorSet = true
	return new
}

// WithSize sets the size variant
func (t *TextareaComponent) WithSize(size flyon.Size) *TextareaComponent {
	new := t.copy()
	new.size = size
	new.sizeSet = true
	return new
}

// WithClasses adds additional CSS classes
func (t *TextareaComponent) WithClasses(classes ...string) *TextareaComponent {
	new := t.copy()
	new.classes = append(new.classes, classes...)
	return new
}

// With applies modifiers to the textarea component
func (t *TextareaComponent) With(modifiers ...any) flyon.Component {
	new := t.copy()
	
	// Apply each modifier
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
func (t *TextareaComponent) copy() *TextareaComponent {
	new := &TextareaComponent{
		id:          t.id,
		name:        t.name,
		value:       t.value,
		placeholder: t.placeholder,
		rows:        t.rows,
		cols:        t.cols,
		disabled:    t.disabled,
		readonly:    t.readonly,
		required:    t.required,
		color:       t.color,
		size:        t.size,
		colorSet:    t.colorSet,
		sizeSet:     t.sizeSet,
		classes:     make([]string, len(t.classes)),
	}
	copy(new.classes, t.classes)
	return new
}

// Render implements the gomponents.Node interface
func (t *TextareaComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"textarea"}
	
	// Add color class if explicitly set
	if t.colorSet {
		classes = append(classes, "textarea-"+t.color.String())
	}
	
	// Add size class if explicitly set
	if t.sizeSet {
		classes = append(classes, "textarea-"+t.size.String())
	}
	
	// Add additional classes
	classes = append(classes, t.classes...)
	
	// Build attributes
	attrs := []gomponents.Node{
		h.Class(strings.Join(classes, " ")),
	}
	
	// Add ID if set
	if t.id != "" {
		attrs = append(attrs, g.Attr("id", t.id))
	}
	
	// Add name if set
	if t.name != "" {
		attrs = append(attrs, g.Attr("name", t.name))
	}
	
	// Add placeholder if set
	if t.placeholder != "" {
		attrs = append(attrs, g.Attr("placeholder", t.placeholder))
	}
	
	// Add rows if set
	if t.rows > 0 {
		attrs = append(attrs, g.Attr("rows", strconv.Itoa(t.rows)))
	}
	
	// Add cols if set
	if t.cols > 0 {
		attrs = append(attrs, g.Attr("cols", strconv.Itoa(t.cols)))
	}
	
	// Add boolean attributes
	if t.disabled {
		attrs = append(attrs, g.Attr("disabled", "disabled"))
	}
	
	if t.readonly {
		attrs = append(attrs, g.Attr("readonly", "readonly"))
	}
	
	if t.required {
		attrs = append(attrs, g.Attr("required", "required"))
	}
	
	// Create textarea element with value as content
	var content gomponents.Node
	if t.value != "" {
		content = g.Text(t.value)
	}
	
	textareaEl := h.Textarea(append(attrs, content)...)
	
	return textareaEl.Render(w)
}

// Ensure TextareaComponent implements the required interfaces
var (
	_ flyon.Component = (*TextareaComponent)(nil)
	_ gomponents.Node = (*TextareaComponent)(nil)
)