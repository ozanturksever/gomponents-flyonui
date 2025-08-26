package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TextareaComponent represents a textarea form element with FlyonUI styling
type TextareaComponent struct {
	attributes []gomponents.Node
	classes    []string
}

// NewTextarea creates a new textarea component
func NewTextarea(attrs ...gomponents.Node) *TextareaComponent {
	return &TextareaComponent{
		attributes: attrs,
		classes:    []string{"textarea"}, // Default FlyonUI textarea class
	}
}

// With applies modifiers to the textarea component
func (t *TextareaComponent) With(modifiers ...any) flyon.Component {
	// Create a new instance to maintain immutability
	newTextarea := &TextareaComponent{
		attributes: make([]gomponents.Node, len(t.attributes)),
		classes:    make([]string, len(t.classes)),
	}
	
	// Copy attributes and classes
	copy(newTextarea.attributes, t.attributes)
	copy(newTextarea.classes, t.classes)
	
	// Apply each modifier
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newTextarea.classes = append(newTextarea.classes, "textarea-"+m.String())
		case flyon.Size:
			newTextarea.classes = append(newTextarea.classes, "textarea-"+m.String())
		case flyon.Variant:
			newTextarea.classes = append(newTextarea.classes, "textarea-"+m.String())
		}
	}
	
	return newTextarea
}

// Render implements the gomponents.Node interface
func (t *TextareaComponent) Render(w io.Writer) error {
	// Build the class attribute
	classAttr := strings.Join(t.classes, " ")
	
	// Create the textarea element with class and attributes
	allNodes := make([]gomponents.Node, 0, len(t.attributes)+1)
	allNodes = append(allNodes, h.Class(classAttr))
	allNodes = append(allNodes, t.attributes...)
	
	textareaEl := h.Textarea(allNodes...)
	
	return textareaEl.Render(w)
}

// Ensure TextareaComponent implements the required interfaces
var (
	_ flyon.Component   = (*TextareaComponent)(nil)
	_ gomponents.Node   = (*TextareaComponent)(nil)
)