package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// FlexComponent represents a flex layout component
type FlexComponent struct {
	classes []string
}

// Ensure FlexComponent implements both interfaces
var (
	_ flyon.Component = (*FlexComponent)(nil)
	_ gomponents.Node = (*FlexComponent)(nil)
)

// NewFlex creates a new flex component
func NewFlex() *FlexComponent {
	return &FlexComponent{
		classes: []string{"flex"},
	}
}

// With applies modifiers to the flex component and returns a new instance
func (f *FlexComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newClasses := make([]string, len(f.classes))
	copy(newClasses, f.classes)

	for _, modifier := range modifiers {
		switch mod := modifier.(type) {
		case flyon.Color:
			newClasses = append(newClasses, "flex-"+mod.String())
		case flyon.Size:
			newClasses = append(newClasses, "flex-"+mod.String())
		case flyon.Variant:
			newClasses = append(newClasses, "flex-"+mod.String())
		}
	}

	return &FlexComponent{
		classes: newClasses,
	}
}

// Render implements the gomponents.Node interface
func (f *FlexComponent) Render(w io.Writer) error {
	classAttr := strings.Join(f.classes, " ")

	flex := html.Div(
		html.Class(classAttr),
	)

	return flex.Render(w)
}