package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// InputComponent represents an input UI component
type InputComponent struct {
	classes    []string
	attributes []g.Node
}

// NewInput creates a new input component with the given attributes
func NewInput(attributes ...g.Node) *InputComponent {
	return &InputComponent{
		classes:    []string{"input"},
		attributes: attributes,
	}
}

// With applies modifiers to the input and returns a new instance
func (i *InputComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newInput := &InputComponent{
		classes:    make([]string, len(i.classes)),
		attributes: make([]g.Node, len(i.attributes)),
	}
	copy(newInput.classes, i.classes)
	copy(newInput.attributes, i.attributes)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newInput.classes = append(newInput.classes, "input-"+m.String())
		case flyon.Size:
			newInput.classes = append(newInput.classes, "input-"+m.String())
		case flyon.Variant:
			newInput.classes = append(newInput.classes, "input-"+m.String())
		}
	}

	return newInput
}

// Render renders the input component to HTML
func (i *InputComponent) Render(w io.Writer) error {
	classAttr := h.Class(strings.Join(i.classes, " "))
	allAttributes := append([]g.Node{classAttr}, i.attributes...)
	return h.Input(allAttributes...).Render(w)
}

// Interface compliance checks
var _ flyon.Component = (*InputComponent)(nil)
var _ g.Node = (*InputComponent)(nil)