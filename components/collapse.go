package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// CollapseComponent represents a collapsible content component.
type CollapseComponent struct {
	id       string
	title    string
	content  gomponents.Node
	open     bool
	arrow    bool // Show arrow indicator
	plus     bool // Show plus/minus indicator
	color    flyon.Color
	classes  []string
}

// NewCollapse creates a new collapse component with the given title and content.
func NewCollapse(title string, content gomponents.Node) *CollapseComponent {
	return &CollapseComponent{
		id:      generateID(),
		title:   title,
		content: content,
		open:    false,
		arrow:   true, // Default to arrow indicator
		plus:    false,
		color:   flyon.Primary,
		classes: []string{},
	}
}

// WithID sets the ID for the collapse component.
func (cc *CollapseComponent) WithID(id string) *CollapseComponent {
	newCC := cc.copy()
	newCC.id = id
	return newCC
}

// WithOpen sets the initial open state of the collapse component.
func (cc *CollapseComponent) WithOpen(open bool) *CollapseComponent {
	newCC := cc.copy()
	newCC.open = open
	return newCC
}

// WithArrow enables or disables the arrow indicator.
func (cc *CollapseComponent) WithArrow(arrow bool) *CollapseComponent {
	newCC := cc.copy()
	newCC.arrow = arrow
	if arrow {
		newCC.plus = false // Only one indicator type at a time
	}
	return newCC
}

// WithPlus enables or disables the plus/minus indicator.
func (cc *CollapseComponent) WithPlus(plus bool) *CollapseComponent {
	newCC := cc.copy()
	newCC.plus = plus
	if plus {
		newCC.arrow = false // Only one indicator type at a time
	}
	return newCC
}

// WithColor sets the color for the collapse component.
func (cc *CollapseComponent) WithColor(color flyon.Color) *CollapseComponent {
	newCC := cc.copy()
	newCC.color = color
	return newCC
}

// WithClasses adds additional CSS classes to the collapse component.
func (cc *CollapseComponent) WithClasses(classes ...string) *CollapseComponent {
	newCC := cc.copy()
	newCC.classes = append(newCC.classes, classes...)
	return newCC
}

// With applies modifiers to the collapse component.
func (cc *CollapseComponent) With(modifiers ...any) flyon.Component {
	newCC := cc.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newCC.color = m
		}
	}
	return newCC
}

// copy creates a deep copy of the collapse component.
func (cc *CollapseComponent) copy() *CollapseComponent {
	newClasses := make([]string, len(cc.classes))
	copy(newClasses, cc.classes)
	
	return &CollapseComponent{
		id:      cc.id,
		title:   cc.title,
		content: cc.content,
		open:    cc.open,
		arrow:   cc.arrow,
		plus:    cc.plus,
		color:   cc.color,
		classes: newClasses,
	}
}

// Render renders the collapse component to HTML.
func (cc *CollapseComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"collapse"}
	
	if cc.arrow {
		classes = append(classes, "collapse-arrow")
	} else if cc.plus {
		classes = append(classes, "collapse-plus")
	}
	
	classes = append(classes, cc.classes...)
	
	// Create the collapse component
	collapseContainer := h.Div(
		h.ID(cc.id),
		h.Class(strings.Join(classes, " ")),
		gomponents.Attr("data-component", "collapse"),
		
		// Hidden checkbox for state management
		h.Input(
			h.Type("checkbox"),
			h.ID(cc.id+"-toggle"),
			h.Class("collapse-toggle"),
			gomponents.If(cc.open, h.Checked()),
		),
		
		// Collapse title/header
		h.Label(
			h.For(cc.id+"-toggle"),
			h.Class("collapse-title text-xl font-medium cursor-pointer"),
			gomponents.Text(cc.title),
		),
		
		// Collapse content
		h.Div(
			h.Class("collapse-content"),
			h.Div(
				h.Class("pb-2"),
				cc.content,
			),
		),
	)
	
	return collapseContainer.Render(w)
}