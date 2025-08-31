package components

import (
	"fmt"
	"io"
	"strings"
	"sync/atomic"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DropdownComponent represents a FlyonUI dropdown component with interactive behavior
type DropdownComponent struct {
	trigger    gomponents.Node
	content    []gomponents.Node
	attributes []gomponents.Node
	classes    []string
	id         string
	position   DropdownPosition
	autoClose  bool
	disabled   bool
}

// DropdownPosition represents the position of the dropdown menu
type DropdownPosition int

const (
	DropdownBottom DropdownPosition = iota
	DropdownTop
	DropdownLeft
	DropdownRight
	DropdownBottomStart
	DropdownBottomEnd
	DropdownTopStart
	DropdownTopEnd
)

// String returns the CSS class for the dropdown position
func (p DropdownPosition) String() string {
	switch p {
	case DropdownTop:
		return "dropdown-top"
	case DropdownLeft:
		return "dropdown-left"
	case DropdownRight:
		return "dropdown-right"
	case DropdownBottomStart:
		return "dropdown-bottom dropdown-start"
	case DropdownBottomEnd:
		return "dropdown-bottom dropdown-end"
	case DropdownTopStart:
		return "dropdown-top dropdown-start"
	case DropdownTopEnd:
		return "dropdown-top dropdown-end"
	default:
		return "dropdown-bottom"
	}
}

// NewDropdown creates a new dropdown component with FlyonUI styling
func NewDropdown(trigger gomponents.Node, content ...gomponents.Node) *DropdownComponent {
	return &DropdownComponent{
		trigger:   trigger,
		content:   content,
		classes:   []string{"dropdown", "relative inline-flex"},
		position:  DropdownBottom,
		autoClose: true,
		disabled:  false,
	}
}

// WithID sets a custom ID for the dropdown
func (d *DropdownComponent) WithID(id string) *DropdownComponent {
	newDropdown := d.copy()
	newDropdown.id = id
	return newDropdown
}

// WithPosition sets the position of the dropdown menu
func (d *DropdownComponent) WithPosition(position DropdownPosition) *DropdownComponent {
	newDropdown := d.copy()
	newDropdown.position = position
	return newDropdown
}

// WithAutoClose controls whether the dropdown closes automatically when clicking outside
func (d *DropdownComponent) WithAutoClose(autoClose bool) *DropdownComponent {
	newDropdown := d.copy()
	newDropdown.autoClose = autoClose
	return newDropdown
}

// WithDisabled sets the disabled state of the dropdown
func (d *DropdownComponent) WithDisabled(disabled bool) *DropdownComponent {
	newDropdown := d.copy()
	newDropdown.disabled = disabled
	return newDropdown
}

// With applies modifiers to the dropdown and returns a new instance
func (d *DropdownComponent) With(modifiers ...any) flyon.Component {
	newDropdown := d.copy()
	
	// Apply each modifier
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Size:
			newDropdown.classes = append(newDropdown.classes, "dropdown-"+m.String())
		case DropdownPosition:
			newDropdown.position = m
		case string:
			// Allow custom CSS classes
			newDropdown.classes = append(newDropdown.classes, m)
		}
	}
	
	return newDropdown
}

// copy creates a deep copy of the dropdown component
func (d *DropdownComponent) copy() *DropdownComponent {
	newDropdown := &DropdownComponent{
		trigger:   d.trigger,
		content:   make([]gomponents.Node, len(d.content)),
		attributes: make([]gomponents.Node, len(d.attributes)),
		classes:   make([]string, len(d.classes)),
		id:        d.id,
		position:  d.position,
		autoClose: d.autoClose,
		disabled:  d.disabled,
	}
	
	copy(newDropdown.content, d.content)
	copy(newDropdown.attributes, d.attributes)
	copy(newDropdown.classes, d.classes)
	
	return newDropdown
}

// Render implements the gomponents.Node interface
func (d *DropdownComponent) Render(w io.Writer) error {
	// Build the class list
	classes := make([]string, len(d.classes))
	copy(classes, d.classes)
	
	// Add position class
	if d.position != DropdownBottom {
		positionClasses := strings.Split(d.position.String(), " ")
		classes = append(classes, positionClasses...)
	}
	// Auto-close behavior via CSS variable class
	if d.autoClose {
		classes = append(classes, "[--auto-close:inside]")
	} else {
		classes = append(classes, "[--auto-close:outside]")
	}

	// Generate ID if not provided
	id := d.id
	if id == "" {
		id = "dropdown-" + generateID()
	}

	// Create trigger with proper attributes
	triggerAttrs := []gomponents.Node{
		h.ID(id + "-toggle"),
		h.Type("button"),
		h.Class("dropdown-toggle"),
		h.Aria("expanded", "false"),
		h.Aria("haspopup", "menu"),
	}

	if d.disabled {
		triggerAttrs = append(triggerAttrs, h.Disabled())
	}

	// Wrap trigger if it's not already a button
	triggerElement := h.Button(
		append(triggerAttrs, d.trigger)...,
	)

	// Create dropdown menu
	menuClasses := "dropdown-menu dropdown-open:opacity-100 hidden min-w-60"
	menu := h.Ul(
		h.Class(menuClasses),
		h.Aria("labelledby", id+"-toggle"),
		h.Aria("orientation", "vertical"),
		h.Role("menu"),
		gomponents.Group(d.content),
	)

	// Create the complete dropdown structure
	dropdownEl := h.Div(
		h.ID(id),
		h.Class(strings.Join(classes, " ")),
		gomponents.Group(d.attributes),
		triggerElement,
		menu,
	)

	return dropdownEl.Render(w)
}

// DropdownItem creates a dropdown menu item
func DropdownItem(children ...gomponents.Node) gomponents.Node {
	return h.Li(
		h.A(
			h.Class("dropdown-item"),
			h.Role("menuitem"),
			gomponents.Group(children),
		),
	)
}

// DropdownDivider creates a divider in the dropdown menu
func DropdownDivider() gomponents.Node {
	return h.Li(
		h.Hr(h.Class("my-1")),
	)
}

// DropdownHeader creates a header in the dropdown menu
func DropdownHeader(text string) gomponents.Node {
	return h.Li(
		h.Div(
			h.Class("dropdown-header font-semibold text-base-content/70"),
			gomponents.Text(text),
		),
	)
}

// generateID generates a unique ID suffix for components
var dropdownIDCounter uint64

func generateID() string {
	n := atomic.AddUint64(&dropdownIDCounter, 1)
	return fmt.Sprintf("d%06d", n)
}

// Ensure DropdownComponent implements the required interfaces
var (
	_ flyon.Component = (*DropdownComponent)(nil)
	_ gomponents.Node = (*DropdownComponent)(nil)
)