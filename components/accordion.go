package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// AccordionItem represents a single accordion item with title and content.
type AccordionItem struct {
	ID      string
	Title   string
	Content gomponents.Node
	Open    bool
}

// AccordionComponent represents an accordion component with multiple collapsible items.
type AccordionComponent struct {
	id       string
	items    []AccordionItem
	multiple bool // Allow multiple items to be open simultaneously
	color    flyon.Color
	classes  []string
}

// NewAccordion creates a new accordion component with the given items.
func NewAccordion(items ...AccordionItem) *AccordionComponent {
	return &AccordionComponent{
		id:       generateID(),
		items:    items,
		multiple: false,
		color:    flyon.Primary,
		classes:  []string{},
	}
}

// WithID sets the ID for the accordion component.
func (ac *AccordionComponent) WithID(id string) *AccordionComponent {
	newAC := ac.copy()
	newAC.id = id
	return newAC
}

// WithMultiple allows multiple accordion items to be open simultaneously.
func (ac *AccordionComponent) WithMultiple(multiple bool) *AccordionComponent {
	newAC := ac.copy()
	newAC.multiple = multiple
	return newAC
}

// WithColor sets the color for the accordion component.
func (ac *AccordionComponent) WithColor(color flyon.Color) *AccordionComponent {
	newAC := ac.copy()
	newAC.color = color
	return newAC
}

// WithClasses adds additional CSS classes to the accordion component.
func (ac *AccordionComponent) WithClasses(classes ...string) *AccordionComponent {
	newAC := ac.copy()
	newAC.classes = append(newAC.classes, classes...)
	return newAC
}

// With applies modifiers to the accordion component.
func (ac *AccordionComponent) With(modifiers ...any) flyon.Component {
	newAC := ac.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newAC.color = m
		}
	}
	return newAC
}

// copy creates a deep copy of the accordion component.
func (ac *AccordionComponent) copy() *AccordionComponent {
	newItems := make([]AccordionItem, len(ac.items))
	copy(newItems, ac.items)
	
	newClasses := make([]string, len(ac.classes))
	copy(newClasses, ac.classes)
	
	return &AccordionComponent{
		id:       ac.id,
		items:    newItems,
		multiple: ac.multiple,
		color:    ac.color,
		classes:  newClasses,
	}
}

// Render renders the accordion component to HTML.
func (ac *AccordionComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"collapse-group"}
	classes = append(classes, ac.classes...)
	
	// Create accordion items
	accordionItems := make([]gomponents.Node, 0, len(ac.items))
	for _, item := range ac.items {
		// Generate unique names for radio buttons if not multiple
		name := ac.id + "-accordion"
		if ac.multiple {
			name = ac.id + "-accordion-" + item.ID
		}
		
		// Create input element (radio for single, checkbox for multiple)
		var inputElement gomponents.Node
		if ac.multiple {
			inputElement = h.Input(
				h.Type("checkbox"),
				h.ID(item.ID),
				h.Name(name),
				h.Class("collapse-toggle"),
				gomponents.If(item.Open, h.Checked()),
			)
		} else {
			inputElement = h.Input(
				h.Type("radio"),
				h.ID(item.ID),
				h.Name(name),
				h.Class("collapse-toggle"),
				gomponents.If(item.Open, h.Checked()),
			)
		}
		
		// Create the accordion item
		accordionItem := h.Div(
			h.Class("collapse collapse-arrow"),
			gomponents.Attr("data-accordion-item", item.ID),
			
			// Hidden input for state management
			inputElement,
			
			// Accordion header/title
			h.Label(
				h.For(item.ID),
				h.Class("collapse-title text-xl font-medium cursor-pointer"),
				gomponents.Text(item.Title),
			),
			
			// Accordion content
			h.Div(
				h.Class("collapse-content"),
				h.Div(
					h.Class("pb-2"),
					item.Content,
				),
			),
		)
		
		accordionItems = append(accordionItems, accordionItem)
	}
	
	// Render the complete accordion component
	accordionContainer := h.Div(
		h.ID(ac.id),
		h.Class(strings.Join(classes, " ")),
		gomponents.Attr("data-component", "accordion"),
		gomponents.If(!ac.multiple, gomponents.Attr("data-single", "true")),
		gomponents.If(ac.multiple, gomponents.Attr("data-multiple", "true")),
		gomponents.Group(accordionItems),
	)
	
	return accordionContainer.Render(w)
}

// NewAccordionItem creates a new accordion item with the given ID, title, and content.
func NewAccordionItem(id, title string, content gomponents.Node) AccordionItem {
	return AccordionItem{
		ID:      id,
		Title:   title,
		Content: content,
		Open:    false,
	}
}

// NewOpenAccordionItem creates a new open accordion item with the given ID, title, and content.
func NewOpenAccordionItem(id, title string, content gomponents.Node) AccordionItem {
	return AccordionItem{
		ID:      id,
		Title:   title,
		Content: content,
		Open:    true,
	}
}