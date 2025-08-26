package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ButtonComponent represents a FlyonUI button component
type ButtonComponent struct {
	children   []gomponents.Node
	attributes []gomponents.Node
	classes    []string
}

// NewButton creates a new button component with FlyonUI styling
func NewButton(children ...gomponents.Node) *ButtonComponent {
	// Separate attributes from content children
	var attributes []gomponents.Node
	var content []gomponents.Node
	
	for _, child := range children {
		// Check if this is an attribute by trying to render it and seeing if it produces attribute-like output
		var buf strings.Builder
		if err := child.Render(&buf); err == nil {
			output := buf.String()
			// If the output contains '=' it's likely an attribute
			if strings.Contains(output, "=") || strings.Contains(output, "disabled") {
				attributes = append(attributes, child)
			} else {
				content = append(content, child)
			}
		} else {
			// If we can't render it, assume it's content
			content = append(content, child)
		}
	}
	
	return &ButtonComponent{
		children:   content,
		attributes: attributes,
		classes:    []string{"btn"}, // Default FlyonUI button class
	}
}

// With applies modifiers to the button and returns a new instance
func (b *ButtonComponent) With(modifiers ...any) flyon.Component {
	// Create a new instance to maintain immutability
	newBtn := &ButtonComponent{
		children:   make([]gomponents.Node, len(b.children)),
		attributes: make([]gomponents.Node, len(b.attributes)),
		classes:    make([]string, len(b.classes)),
	}
	
	// Copy children, attributes, and classes
	copy(newBtn.children, b.children)
	copy(newBtn.attributes, b.attributes)
	copy(newBtn.classes, b.classes)
	
	// Apply each modifier
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newBtn.classes = append(newBtn.classes, "btn-"+m.String())
		case flyon.Size:
			newBtn.classes = append(newBtn.classes, "btn-"+m.String())
		case flyon.Variant:
			newBtn.classes = append(newBtn.classes, "btn-"+m.String())
		}
	}
	
	return newBtn
}

// Render implements the gomponents.Node interface
func (b *ButtonComponent) Render(w io.Writer) error {
	// Build the class attribute
	classAttr := strings.Join(b.classes, " ")
	
	// Create the button element with class, attributes, and children
	allNodes := make([]gomponents.Node, 0, len(b.attributes)+len(b.children)+1)
	allNodes = append(allNodes, h.Class(classAttr))
	allNodes = append(allNodes, b.attributes...)
	allNodes = append(allNodes, b.children...)
	
	buttonEl := h.Button(allNodes...)
	
	return buttonEl.Render(w)
}

// Ensure ButtonComponent implements the required interfaces
var (
	_ flyon.Component   = (*ButtonComponent)(nil)
	_ gomponents.Node   = (*ButtonComponent)(nil)
)