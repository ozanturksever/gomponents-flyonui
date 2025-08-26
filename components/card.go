package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// CardComponent represents a card UI component
type CardComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
}

// NewCard creates a new card component with the given children
func NewCard(children ...g.Node) *CardComponent {
	// Separate attributes from content children
	var attributes []g.Node
	var content []g.Node
	
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
	
	return &CardComponent{
		classes:    []string{"card"},
		attributes: attributes,
		children:   content,
	}
}

// With applies modifiers to the card and returns a new instance
func (c *CardComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newCard := &CardComponent{
		classes:    make([]string, len(c.classes)),
		attributes: make([]g.Node, len(c.attributes)),
		children:   make([]g.Node, len(c.children)),
	}
	copy(newCard.classes, c.classes)
	copy(newCard.attributes, c.attributes)
	copy(newCard.children, c.children)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newCard.classes = append(newCard.classes, "card-"+m.String())
		case flyon.Size:
			newCard.classes = append(newCard.classes, "card-"+m.String())
		case flyon.Variant:
			newCard.classes = append(newCard.classes, "card-"+m.String())
		}
	}

	return newCard
}

// Render renders the card component to HTML
func (c *CardComponent) Render(w io.Writer) error {
	classAttr := h.Class(strings.Join(c.classes, " "))
	allAttributes := append([]g.Node{classAttr}, c.attributes...)
	allNodes := append(allAttributes, c.children...)
	return h.Div(allNodes...).Render(w)
}

// Interface compliance checks
var _ flyon.Component = (*CardComponent)(nil)
var _ g.Node = (*CardComponent)(nil)