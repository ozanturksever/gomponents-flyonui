package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// AvatarComponent represents an avatar UI component
type AvatarComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
}

// NewAvatar creates a new avatar component with the given children
func NewAvatar(children ...g.Node) *AvatarComponent {
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
	
	return &AvatarComponent{
		classes:    []string{"avatar"},
		attributes: attributes,
		children:   content,
	}
}

// With applies modifiers to the avatar and returns a new instance
func (a *AvatarComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newAvatar := &AvatarComponent{
		classes:    make([]string, len(a.classes)),
		attributes: make([]g.Node, len(a.attributes)),
		children:   make([]g.Node, len(a.children)),
	}
	copy(newAvatar.classes, a.classes)
	copy(newAvatar.attributes, a.attributes)
	copy(newAvatar.children, a.children)

	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newAvatar.classes = append(newAvatar.classes, "avatar-"+m.String())
		case flyon.Size:
			newAvatar.classes = append(newAvatar.classes, "avatar-"+m.String())
		case flyon.Variant:
			newAvatar.classes = append(newAvatar.classes, "avatar-"+m.String())
		}
	}

	return newAvatar
}

// Render renders the avatar component to HTML
func (a *AvatarComponent) Render(w io.Writer) error {
	classAttr := h.Class(strings.Join(a.classes, " "))
	allAttributes := append([]g.Node{classAttr}, a.attributes...)
	allNodes := append(allAttributes, a.children...)
	return h.Div(allNodes...).Render(w)
}

// Interface compliance checks
var _ flyon.Component = (*AvatarComponent)(nil)
var _ g.Node = (*AvatarComponent)(nil)