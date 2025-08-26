package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// BreadcrumbComponent represents a breadcrumb navigation component
type BreadcrumbComponent struct {
	attributes []g.Node
	items      []g.Node
	classes    []string
	modifiers  []any
	separator  string
}

// NewBreadcrumb creates a new breadcrumb component with the given items
func NewBreadcrumb(items ...g.Node) *BreadcrumbComponent {
	var attributes []g.Node
	var breadcrumbItems []g.Node
	
	// Separate attributes from breadcrumb items
	for _, item := range items {
		switch v := item.(type) {
		case *BreadcrumbItemComponent:
			breadcrumbItems = append(breadcrumbItems, v)
		default:
			// Check if it's an HTML attribute (ID, Class, DataAttr, etc.)
			if isHTMLAttribute(v) {
				attributes = append(attributes, v)
			} else {
				// Treat as breadcrumb item content
				breadcrumbItems = append(breadcrumbItems, BreadcrumbItem(v))
			}
		}
	}
	
	return &BreadcrumbComponent{
		attributes: attributes,
		items:      breadcrumbItems,
		classes:    []string{"breadcrumbs"},
		modifiers:  []any{},
	}
}

// BreadcrumbItemComponent represents a single breadcrumb item
type BreadcrumbItemComponent struct {
	content g.Node
}

// BreadcrumbItem creates a new breadcrumb item
func BreadcrumbItem(content g.Node) *BreadcrumbItemComponent {
	return &BreadcrumbItemComponent{
		content: content,
	}
}

// Render implements gomponents.Node
func (b *BreadcrumbItemComponent) Render(w io.Writer) error {
	return h.Li(b.content).Render(w)
}

// WithSeparator sets a custom separator for the breadcrumb
func (b *BreadcrumbComponent) WithSeparator(separator string) *BreadcrumbComponent {
	newComponent := *b
	newComponent.separator = separator
	return &newComponent
}

// With applies modifiers to the breadcrumb component
func (b *BreadcrumbComponent) With(modifiers ...any) flyon.Component {
	newComponent := *b
	newComponent.modifiers = append([]any{}, b.modifiers...)
	newComponent.modifiers = append(newComponent.modifiers, modifiers...)
	
	// Apply modifiers to create new classes
	newComponent.classes = append([]string{}, b.classes...)
	
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Size:
			newComponent.classes = append(newComponent.classes, "breadcrumbs-"+m.String())
		case string:
			// Handle custom CSS classes
			newComponent.classes = append(newComponent.classes, m)
		case g.Node:
			// Handle additional attributes
			newComponent.attributes = append(newComponent.attributes, m)
		}
	}
	
	return &newComponent
}

// Render implements gomponents.Node
func (b *BreadcrumbComponent) Render(w io.Writer) error {
	// Build the class attribute
	classAttr := h.Class(strings.Join(b.classes, " "))
	
	// Combine all attributes
	allAttributes := []g.Node{classAttr}
	allAttributes = append(allAttributes, b.attributes...)
	
	// Add separator data attribute if custom separator is set
	if b.separator != "" {
		allAttributes = append(allAttributes, h.DataAttr("separator", b.separator))
	}
	
	// Create the breadcrumb list
	breadcrumbList := h.Ol(b.items...)
	
	// Combine attributes with the list
	allAttributes = append(allAttributes, breadcrumbList)
	
	return h.Nav(allAttributes...).Render(w)
}

// isHTMLAttribute checks if a node is an HTML attribute
func isHTMLAttribute(node g.Node) bool {
	// This is a simple heuristic - in practice, you might want to use type assertions
	// or a more sophisticated method to detect HTML attributes
	var buf strings.Builder
	node.Render(&buf)
	html := buf.String()
	
	// Check if it looks like an attribute (contains = but no < >)
	return strings.Contains(html, "=") && !strings.Contains(html, "<") && !strings.Contains(html, ">")
}