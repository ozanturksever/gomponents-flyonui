package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// StackComponent represents a stack layout component
type StackComponent struct {
	classes []string
}

// Ensure StackComponent implements both interfaces
var (
	_ flyon.Component = (*StackComponent)(nil)
	_ gomponents.Node = (*StackComponent)(nil)
)

// NewStack creates a new stack component
func NewStack() *StackComponent {
	return &StackComponent{
		classes: []string{"stack"},
	}
}

// With applies modifiers to the stack component and returns a new instance
func (s *StackComponent) With(modifiers ...any) flyon.Component {
	newClasses := make([]string, len(s.classes))
	copy(newClasses, s.classes)

	for _, modifier := range modifiers {
		switch mod := modifier.(type) {
		case flyon.Color:
			newClasses = append(newClasses, "stack-"+mod.String())
		case flyon.Size:
			newClasses = append(newClasses, "stack-"+mod.String())
		case flyon.Variant:
			newClasses = append(newClasses, "stack-"+mod.String())
		}
	}

	return &StackComponent{
		classes: newClasses,
	}
}

// Render implements the gomponents.Node interface
func (s *StackComponent) Render(w io.Writer) error {
	classAttr := strings.Join(s.classes, " ")

	stack := html.Div(
		html.Class(classAttr),
	)

	return stack.Render(w)
}