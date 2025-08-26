package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// AlertComponent represents an alert UI component
type AlertComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
}

// NewAlert creates a new alert component with the given children
func NewAlert(children ...g.Node) *AlertComponent {
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
	
	return &AlertComponent{
		classes:    []string{"alert"},
		attributes: attributes,
		children:   content,
	}
}

// With applies modifiers to the alert and returns a new instance
func (a *AlertComponent) With(modifiers ...any) flyon.Component {
	newAlert := &AlertComponent{
		classes:    make([]string, len(a.classes)),
		attributes: make([]g.Node, len(a.attributes)),
		children:   make([]g.Node, len(a.children)),
	}
	copy(newAlert.classes, a.classes)
	copy(newAlert.attributes, a.attributes)
	copy(newAlert.children, a.children)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newAlert.classes = append(newAlert.classes, "alert-"+m.String())
		case flyon.Size:
			newAlert.classes = append(newAlert.classes, "alert-"+m.String())
		case flyon.Variant:
			newAlert.classes = append(newAlert.classes, "alert-"+m.String())
		}
	}

	return newAlert
}

// Render renders the alert component to HTML
func (a *AlertComponent) Render(w io.Writer) error {
	classAttr := h.Class(strings.Join(a.classes, " "))
	allAttributes := append([]g.Node{classAttr}, a.attributes...)
	allNodes := append(allAttributes, a.children...)
	return h.Div(allNodes...).Render(w)
}

// Interface compliance checks
var _ flyon.Component = (*AlertComponent)(nil)
var _ g.Node = (*AlertComponent)(nil)