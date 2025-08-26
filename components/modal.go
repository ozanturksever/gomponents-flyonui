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
	ModalSizeSmall ModalSize = iota
	ModalSizeMedium
	ModalSizeLarge
	ModalSizeExtraLarge
	ModalSizeFullWidth
)

// String returns the CSS class for the modal size
func (s ModalSize) String() string {
	switch s {
	case ModalSizeSmall:
		return "modal-sm"
	case ModalSizeMedium:
		return "modal-md"
	case ModalSizeLarge:
		return "modal-lg"
	case ModalSizeExtraLarge:
		return "modal-xl"
	case ModalSizeFullWidth:
		return "modal-full"
	default:
		return "modal-md"
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
	closable   bool
	backdrop   bool
	open       bool
}

// NewModal creates a new modal component with FlyonUI styling
func NewModal(title string, content ...gomponents.Node) *ModalComponent {
	return &ModalComponent{
		title:    title,
		content:  content,
		classes:  []string{"modal"},
		size:     ModalSizeMedium,
		closable: true,
		backdrop: true,
		open:     false,
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
		closable:   m.closable,
		backdrop:   m.backdrop,
		open:       m.open,
	}
	
	copy(newModal.content, m.content)
	copy(newModal.actions, m.actions)
	copy(newModal.attributes, m.attributes)
	copy(newModal.classes, m.classes)
	
	return newModal
}

// Render implements the gomponents.Node interface
func (m *ModalComponent) Render(w io.Writer) error {
	// Build the class list
	classes := make([]string, len(m.classes))
	copy(classes, m.classes)
	
	// Add size class if not default
	if m.size != ModalSizeMedium {
		classes = append(classes, m.size.String())
	}
	
	// Add open class if modal should be open
	if m.open {
		classes = append(classes, "modal-open")
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
		h.DataAttr("hs-overlay", ""),
		h.Role("dialog"),
		h.Aria("labelledby", id+"-title"),
		h.Aria("modal", "true"),
	}
	
	// Add backdrop attributes
	if !m.backdrop {
		modalAttrs = append(modalAttrs, h.DataAttr("hs-overlay-backdrop", "false"))
	}
	
	// Add closable attributes
	if !m.closable {
		modalAttrs = append(modalAttrs, h.DataAttr("hs-overlay-keyboard", "false"))
	}
	
	// Add custom attributes
	modalAttrs = append(modalAttrs, m.attributes...)
	
	// Create modal header
	header := h.Div(
		h.Class("modal-header flex items-center justify-between p-4 border-b"),
		h.H3(
			h.ID(id+"-title"),
			h.Class("text-lg font-semibold"),
			gomponents.Text(m.title),
		),
	)
	
	// Add close button if closable
	if m.closable {
		header = h.Div(
			h.Class("modal-header flex items-center justify-between p-4 border-b"),
			h.H3(
				h.ID(id+"-title"),
				h.Class("text-lg font-semibold"),
				gomponents.Text(m.title),
			),
			h.Button(
				h.Type("button"),
				h.Class("btn btn-sm btn-circle btn-ghost"),
				h.DataAttr("hs-overlay-close", ""),
				h.Aria("label", "Close"),
				gomponents.Text("✕"),
			),
		)
	}
	
	// Create modal body
	body := h.Div(
		h.Class("modal-body p-4"),
		gomponents.Group(m.content),
	)
	
	// Create modal footer if actions are provided
	var footer gomponents.Node
	if len(m.actions) > 0 {
		footer = h.Div(
			h.Class("modal-footer flex justify-end gap-2 p-4 border-t"),
			gomponents.Group(m.actions),
		)
	}
	
	// Create modal content container
	modalContent := []gomponents.Node{header, body}
	if footer != nil {
		modalContent = append(modalContent, footer)
	}
	
	// Create the complete modal structure
	modalEl := h.Div(
		append(modalAttrs,
			h.Div(
				h.Class("modal-dialog relative w-auto pointer-events-none"),
				h.Div(
					h.Class("modal-content relative flex flex-col w-full pointer-events-auto bg-white border border-gray-200 rounded-lg shadow-lg"),
					gomponents.Group(modalContent),
				),
			),
		)...,
	)
	
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
	return ModalAction(text, variant, h.DataAttr("hs-overlay-close", ""))
}

// Ensure ModalComponent implements the required interfaces
var (
	_ flyon.Component = (*ModalComponent)(nil)
	_ gomponents.Node = (*ModalComponent)(nil)
)