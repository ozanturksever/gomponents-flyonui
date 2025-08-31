package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ModalSize represents different modal sizes
type ModalSize int

const (
	ModalSizeDefault ModalSize = iota // No size class (default)
	ModalSizeSmall
	ModalSizeMedium
	ModalSizeLarge
	ModalSizeExtraLarge
	ModalSizeFullWidth
)

// String returns the CSS class for the modal size
func (s ModalSize) String() string {
	switch s {
	case ModalSizeSmall:
		return "modal-dialog-sm"
	case ModalSizeMedium:
		return "modal-dialog-md"
	case ModalSizeLarge:
		return "modal-dialog-lg"
	case ModalSizeExtraLarge:
		return "modal-dialog-xl"
	case ModalSizeFullWidth:
		return "modal-dialog-full"
	default:
		return "" // No size class for default
	}
}

// ModalPosition represents different modal positions
type ModalPosition int

const (
	ModalPositionDefault ModalPosition = iota // No position class (default)
	ModalPositionMiddle
	ModalPositionBottom
)

// String returns the CSS class for the modal position
func (p ModalPosition) String() string {
	switch p {
	case ModalPositionMiddle:
		return "modal-middle"
	case ModalPositionBottom:
		return "modal-bottom"
	default:
		return "" // No position class for default
	}
}

// ModalComponent represents a modal dialog component
type ModalComponent struct {
	title      string
	content    []gomponents.Node
	actions    []gomponents.Node
	attributes []gomponents.Node
	classes    []string
	id         string
	size       ModalSize
	position   ModalPosition
	closable   bool
	backdrop   bool
	open       bool
	keyboard   bool
}

// NewModal creates a new modal component with FlyonUI styling
func NewModal(title string, content ...gomponents.Node) *ModalComponent {
	return &ModalComponent{
		title:    title,
		content:  content,
		classes:  []string{},
		size:     ModalSizeDefault,
		position: ModalPositionDefault,
		closable: true,
		backdrop: true,
		open:     false,
		keyboard: true,
	}
}

// WithID sets a custom ID for the modal
func (m *ModalComponent) WithID(id string) *ModalComponent {
	newModal := m.copy()
	newModal.id = id
	return newModal
}

// WithSize sets the size of the modal
func (m *ModalComponent) WithSize(size ModalSize) *ModalComponent {
	newModal := m.copy()
	newModal.size = size
	return newModal
}

// WithClosable controls whether the modal can be closed by clicking outside or pressing escape
func (m *ModalComponent) WithClosable(closable bool) *ModalComponent {
	newModal := m.copy()
	newModal.closable = closable
	// When not closable, also disable keyboard
	if !closable {
		newModal.keyboard = false
	}
	return newModal
}

// WithBackdrop controls whether the modal has a backdrop
func (m *ModalComponent) WithBackdrop(backdrop bool) *ModalComponent {
	newModal := m.copy()
	newModal.backdrop = backdrop
	return newModal
}

// WithOpen sets the initial open state of the modal
func (m *ModalComponent) WithOpen(open bool) *ModalComponent {
	newModal := m.copy()
	newModal.open = open
	return newModal
}

// WithPosition sets the position of the modal
func (m *ModalComponent) WithPosition(position ModalPosition) *ModalComponent {
	newModal := m.copy()
	newModal.position = position
	return newModal
}

// WithKeyboard controls whether the modal can be closed with the Escape key
func (m *ModalComponent) WithKeyboard(keyboard bool) *ModalComponent {
	newModal := m.copy()
	newModal.keyboard = keyboard
	return newModal
}

// WithActions sets the action buttons for the modal
func (m *ModalComponent) WithActions(actions ...gomponents.Node) *ModalComponent {
	newModal := m.copy()
	newModal.actions = make([]gomponents.Node, len(actions))
	copy(newModal.actions, actions)
	return newModal
}

// With applies modifiers to the modal and returns a new instance
func (m *ModalComponent) With(modifiers ...any) flyon.Component {
	newModal := m.copy()

	// Apply each modifier
	for _, modifier := range modifiers {
		switch mod := modifier.(type) {
		case flyon.Size:
			// Map flyon.Size to ModalSize
			switch mod {
			case flyon.SizeSmall:
				newModal.size = ModalSizeSmall
			case flyon.SizeLarge:
				newModal.size = ModalSizeLarge
			default:
				newModal.size = ModalSizeMedium
			}
		case ModalSize:
			newModal.size = mod
		case string:
			// Allow custom CSS classes
			newModal.classes = append(newModal.classes, mod)
		}
	}

	return newModal
}

// copy creates a deep copy of the modal component
func (m *ModalComponent) copy() *ModalComponent {
	newModal := &ModalComponent{
		title:      m.title,
		content:    make([]gomponents.Node, len(m.content)),
		actions:    make([]gomponents.Node, len(m.actions)),
		attributes: make([]gomponents.Node, len(m.attributes)),
		classes:    make([]string, len(m.classes)),
		id:         m.id,
		size:       m.size,
		position:   m.position,
		closable:   m.closable,
		backdrop:   m.backdrop,
		open:       m.open,
		keyboard:   m.keyboard,
	}

	copy(newModal.content, m.content)
	copy(newModal.actions, m.actions)
	copy(newModal.attributes, m.attributes)
	copy(newModal.classes, m.classes)

	return newModal
}

// Render implements the gomponents.Node interface
func (m *ModalComponent) Render(w io.Writer) error {
	// Build the class list for the modal container
	classes := make([]string, 0, len(m.classes)+5)
	// Ensure required container classes per latest FlyonUI docs
	classes = append(classes, "overlay", "modal", "overlay-open:opacity-100", "overlay-open:duration-300")
	classes = append(classes, m.classes...)

	// Add position class if specified
	if positionClass := m.position.String(); positionClass != "" {
		classes = append(classes, positionClass)
	}

	// Visibility: use hidden class when not open
	if !m.open {
		classes = append(classes, "hidden")
	}

	// Generate ID if not provided
	id := m.id
	if id == "" {
		id = "modal-" + generateID()
	}

	// Create modal attributes
	modalAttrs := []gomponents.Node{
		h.ID(id),
		h.Class(strings.Join(classes, " ")),
		h.Role("dialog"),
		gomponents.Attr("tabindex", "-1"),
	}

	// Keyboard behavior
	if !m.keyboard {
		modalAttrs = append(modalAttrs, gomponents.Attr("data-overlay-keyboard", "false"))
	}

	// Add custom attributes (e.g., overlay-open:* utilities)
	modalAttrs = append(modalAttrs, m.attributes...)

	// Build dialog/content structure classes
	dialogClasses := []string{"modal-dialog"}
	if sizeClass := m.size.String(); sizeClass != "" {
		dialogClasses = append(dialogClasses, sizeClass)
	}
	contentClasses := []string{"modal-content"}

	// Header
	var headerNodes []gomponents.Node
	if m.title != "" {
		headerNodes = append(headerNodes, h.H3(
			h.Class("modal-title"),
			gomponents.Text(m.title),
		))
	}
	if m.closable {
		headerNodes = append(headerNodes, h.Button(
			h.Type("button"),
			h.Class("btn btn-text btn-circle btn-sm absolute end-3 top-3"),
			// Close via data-overlay selector to this modal id
			gomponents.Attr("data-overlay", "#"+id),
			h.Aria("label", "Close"),
			// icon placeholder span per docs
			h.Span(h.Class("icon-[tabler--x] size-4")),
		))
	}

	// Body and footer
	bodyNode := h.Div(h.Class("modal-body"), gomponents.Group(m.content))
	var footerNode gomponents.Node
	if len(m.actions) > 0 {
		footerNode = h.Div(h.Class("modal-footer"), gomponents.Group(m.actions))
	}

	// Assemble modal element
	modalEl := h.Div(append(modalAttrs,
		h.Div(
			h.Class(strings.Join(dialogClasses, " ")),
			h.Div(
				h.Class(strings.Join(contentClasses, " ")),
				// header (only if has content)
				gomponents.If(len(headerNodes) > 0, h.Div(h.Class("modal-header"), gomponents.Group(headerNodes))),
				bodyNode,
				gomponents.If(footerNode != nil, footerNode),
			),
		),
	)...)

	return modalEl.Render(w)
}

// ModalAction creates a modal action button
func ModalAction(text string, variant flyon.Color, attrs ...gomponents.Node) gomponents.Node {
	buttonAttrs := []gomponents.Node{
		h.Type("button"),
		h.Class("btn btn-" + variant.String()),
	}
	buttonAttrs = append(buttonAttrs, attrs...)
	buttonAttrs = append(buttonAttrs, gomponents.Text(text))

	return h.Button(buttonAttrs...)
}

// ModalCloseAction creates a modal close action button
func ModalCloseAction(text string, variant flyon.Color) gomponents.Node {
	// 'modal-close' class is kept for internal/demo JS hooks; no HSOverlay data attribute per new docs
	return ModalAction(text, variant, h.Class("modal-close"))
}

// Ensure ModalComponent implements the required interfaces
var (
	_ flyon.Component = (*ModalComponent)(nil)
	_ gomponents.Node = (*ModalComponent)(nil)
)
