package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// ContainerComponent represents a responsive container component
type ContainerComponent struct {
	classes []string
}

// Ensure ContainerComponent implements both interfaces
var (
	_ flyon.Component = (*ContainerComponent)(nil)
	_ gomponents.Node = (*ContainerComponent)(nil)
)

// NewContainer creates a new container component with default styling
func NewContainer() *ContainerComponent {
	return &ContainerComponent{
		classes: []string{"container"},
	}
}

// With applies modifiers to the container and returns a new instance
func (c *ContainerComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newClasses := make([]string, len(c.classes))
	copy(newClasses, c.classes)

	for _, modifier := range modifiers {
		switch mod := modifier.(type) {
		case flyon.Color:
			newClasses = append(newClasses, "container-"+mod.String())
		case flyon.Size:
			newClasses = append(newClasses, "container-"+mod.String())
		case flyon.Variant:
			newClasses = append(newClasses, "container-"+mod.String())
		}
	}

	return &ContainerComponent{
		classes: newClasses,
	}
}

// Render implements the gomponents.Node interface
func (c *ContainerComponent) Render(w io.Writer) error {
	classAttr := strings.Join(c.classes, " ")

	container := html.Div(
		html.Class(classAttr),
	)

	return container.Render(w)
}