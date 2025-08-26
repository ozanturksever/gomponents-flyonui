package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DrawerSide represents the side where the drawer appears.
type DrawerSide int

const (
	DrawerLeft DrawerSide = iota
	DrawerRight
	DrawerTop
	DrawerBottom
)

// String returns the CSS class for the drawer side.
func (ds DrawerSide) String() string {
	switch ds {
	case DrawerLeft:
		return "drawer-start"
	case DrawerRight:
		return "drawer-end"
	case DrawerTop:
		return "drawer-top"
	case DrawerBottom:
		return "drawer-bottom"
	default:
		return "drawer-start"
	}
}

// DrawerComponent represents a drawer/sidebar component.
type DrawerComponent struct {
	id       string
	side     DrawerSide
	open     bool
	overlay  bool // Show overlay when open
	content  gomponents.Node // Main content
	sidebar  gomponents.Node // Drawer sidebar content
	classes  []string
}

// NewDrawer creates a new drawer component with the given content and sidebar.
func NewDrawer(content, sidebar gomponents.Node) *DrawerComponent {
	return &DrawerComponent{
		id:      generateID(),
		side:    DrawerLeft,
		open:    false,
		overlay: true, // Default to overlay mode
		content: content,
		sidebar: sidebar,
		classes: []string{},
	}
}

// WithID sets the ID for the drawer component.
func (dc *DrawerComponent) WithID(id string) *DrawerComponent {
	newDC := dc.copy()
	newDC.id = id
	return newDC
}

// WithSide sets the side where the drawer appears.
func (dc *DrawerComponent) WithSide(side DrawerSide) *DrawerComponent {
	newDC := dc.copy()
	newDC.side = side
	return newDC
}

// WithOpen sets the initial open state of the drawer.
func (dc *DrawerComponent) WithOpen(open bool) *DrawerComponent {
	newDC := dc.copy()
	newDC.open = open
	return newDC
}

// WithOverlay enables or disables the overlay when drawer is open.
func (dc *DrawerComponent) WithOverlay(overlay bool) *DrawerComponent {
	newDC := dc.copy()
	newDC.overlay = overlay
	return newDC
}

// WithClasses adds additional CSS classes to the drawer component.
func (dc *DrawerComponent) WithClasses(classes ...string) *DrawerComponent {
	newDC := dc.copy()
	newDC.classes = append(newDC.classes, classes...)
	return newDC
}

// With applies modifiers to the drawer component.
func (dc *DrawerComponent) With(modifiers ...any) flyon.Component {
	newDC := dc.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case DrawerSide:
			newDC.side = m
		}
	}
	return newDC
}

// copy creates a deep copy of the drawer component.
func (dc *DrawerComponent) copy() *DrawerComponent {
	newClasses := make([]string, len(dc.classes))
	copy(newClasses, dc.classes)
	
	return &DrawerComponent{
		id:      dc.id,
		side:    dc.side,
		open:    dc.open,
		overlay: dc.overlay,
		content: dc.content,
		sidebar: dc.sidebar,
		classes: newClasses,
	}
}

// Render renders the drawer component to HTML.
func (dc *DrawerComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"drawer"}
	classes = append(classes, dc.side.String())
	classes = append(classes, dc.classes...)
	
	// Create the drawer component
	drawerContainer := h.Div(
		h.ID(dc.id),
		h.Class(strings.Join(classes, " ")),
		gomponents.Attr("data-component", "drawer"),
		
		// Hidden checkbox for state management
		h.Input(
			h.Type("checkbox"),
			h.ID(dc.id+"-toggle"),
			h.Class("drawer-toggle"),
			gomponents.If(dc.open, h.Checked()),
		),
		
		// Drawer content (main content area)
		h.Div(
			h.Class("drawer-content flex flex-col"),
			dc.content,
		),
		
		// Drawer side (sidebar)
		h.Div(
			h.Class("drawer-side"),
			
			// Overlay (if enabled)
			gomponents.If(dc.overlay,
				h.Label(
					h.For(dc.id+"-toggle"),
					h.Class("drawer-overlay"),
					gomponents.Attr("aria-label", "close sidebar"),
				),
			),
			
			// Sidebar content
		h.Aside(
				h.Class("bg-base-200 min-h-full w-80 p-4"),
				dc.sidebar,
			),
		),
	)
	
	return drawerContainer.Render(w)
}

// DrawerToggleButton creates a button that toggles the drawer.
func DrawerToggleButton(drawerID, text string) gomponents.Node {
	return h.Label(
		h.For(drawerID+"-toggle"),
		h.Class("btn btn-square btn-ghost drawer-button"),
		gomponents.Attr("aria-label", "toggle drawer"),
		gomponents.Text(text),
	)
}

// DrawerCloseButton creates a button that closes the drawer.
func DrawerCloseButton(drawerID, text string) gomponents.Node {
	return h.Label(
		h.For(drawerID+"-toggle"),
		h.Class("btn btn-sm btn-circle btn-ghost absolute right-2 top-2"),
		gomponents.Attr("aria-label", "close drawer"),
		gomponents.Text(text),
	)
}