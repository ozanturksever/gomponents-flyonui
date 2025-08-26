package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// GridComponent represents a grid layout component
type GridComponent struct {
	classes []string
}

// Ensure GridComponent implements both interfaces
var (
	_ flyon.Component = (*GridComponent)(nil)
	_ gomponents.Node = (*GridComponent)(nil)
)

// NewGrid creates a new grid component
func NewGrid() *GridComponent {
	return &GridComponent{
		classes: []string{"grid"},
	}
}

// With applies modifiers to the grid component and returns a new instance
func (g *GridComponent) With(modifiers ...flyon.Modifier) flyon.Component {
	newClasses := make([]string, len(g.classes))
	copy(newClasses, g.classes)

	for _, modifier := range modifiers {
		switch mod := modifier.(type) {
		case flyon.Color:
			newClasses = append(newClasses, "grid-"+mod.String())
		case flyon.Size:
			newClasses = append(newClasses, "grid-"+mod.String())
		case flyon.Variant:
			newClasses = append(newClasses, "grid-"+mod.String())
		}
	}

	return &GridComponent{
		classes: newClasses,
	}
}

// Render implements the gomponents.Node interface
func (g *GridComponent) Render(w io.Writer) error {
	classAttr := strings.Join(g.classes, " ")

	grid := html.Div(
		html.Class(classAttr),
	)

	return grid.Render(w)
}