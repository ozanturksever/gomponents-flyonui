package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TabsVariant represents the visual style variants for tabs.
type TabsVariant int

const (
	TabsDefault TabsVariant = iota
	TabsBordered
	TabsLifted
	TabsBoxed
)

// String returns the CSS class for the tabs variant.
func (tv TabsVariant) String() string {
	switch tv {
	case TabsBordered:
		return "tabs-bordered"
	case TabsLifted:
		return "tabs-lifted"
	case TabsBoxed:
		return "tabs-boxed"
	default:
		return ""
	}
}

// TabsSize represents the size variants for tabs.
type TabsSize int

const (
	TabsSizeXS TabsSize = iota
	TabsSizeSmall
	TabsSizeMedium
	TabsSizeLarge
)

// String returns the CSS class for the tabs size.
func (ts TabsSize) String() string {
	switch ts {
	case TabsSizeXS:
		return "tabs-xs"
	case TabsSizeSmall:
		return "tabs-sm"
	case TabsSizeLarge:
		return "tabs-lg"
	default:
		return ""
	}
}

// TabItem represents a single tab item with label and content.
type TabItem struct {
	ID      string
	Label   string
	Content gomponents.Node
	Active  bool
}

// TabsComponent represents a tabs component with multiple tab items.
type TabsComponent struct {
	id       string
	tabs     []TabItem
	variant  TabsVariant
	size     TabsSize
	color    flyon.Color
	classes  []string
}

// NewTabs creates a new tabs component with the given tab items.
func NewTabs(tabs ...TabItem) *TabsComponent {
	return &TabsComponent{
		id:      generateID(),
		tabs:    tabs,
		variant: TabsDefault,
		size:    TabsSizeMedium,
		color:   flyon.Primary,
		classes: []string{},
	}
}

// WithID sets the ID for the tabs component.
func (tc *TabsComponent) WithID(id string) *TabsComponent {
	newTC := tc.copy()
	newTC.id = id
	return newTC
}

// WithVariant sets the variant for the tabs component.
func (tc *TabsComponent) WithVariant(variant TabsVariant) *TabsComponent {
	newTC := tc.copy()
	newTC.variant = variant
	return newTC
}

// WithSize sets the size for the tabs component.
func (tc *TabsComponent) WithSize(size TabsSize) *TabsComponent {
	newTC := tc.copy()
	newTC.size = size
	return newTC
}

// WithColor sets the color for the tabs component.
func (tc *TabsComponent) WithColor(color flyon.Color) *TabsComponent {
	newTC := tc.copy()
	newTC.color = color
	return newTC
}

// WithClasses adds additional CSS classes to the tabs component.
func (tc *TabsComponent) WithClasses(classes ...string) *TabsComponent {
	newTC := tc.copy()
	newTC.classes = append(newTC.classes, classes...)
	return newTC
}

// With applies modifiers to the tabs component.
func (tc *TabsComponent) With(modifiers ...any) flyon.Component {
	newTC := tc.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newTC.color = m
		case flyon.Size:
			// Map flyon.Size to TabsSize
			switch m {
			case flyon.SizeXS:
				newTC.size = TabsSizeXS
			case flyon.SizeSmall:
				newTC.size = TabsSizeSmall
			case flyon.SizeLarge:
				newTC.size = TabsSizeLarge
			default:
				newTC.size = TabsSizeMedium
			}
		case TabsVariant:
			newTC.variant = m
		case TabsSize:
			newTC.size = m
		}
	}
	return newTC
}

// copy creates a deep copy of the tabs component.
func (tc *TabsComponent) copy() *TabsComponent {
	newTabs := make([]TabItem, len(tc.tabs))
	copy(newTabs, tc.tabs)
	
	newClasses := make([]string, len(tc.classes))
	copy(newClasses, tc.classes)
	
	return &TabsComponent{
		id:      tc.id,
		tabs:    newTabs,
		variant: tc.variant,
		size:    tc.size,
		color:   tc.color,
		classes: newClasses,
	}
}

// Render renders the tabs component to HTML.
func (tc *TabsComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"tabs"}
	
	if tc.variant != TabsDefault {
		classes = append(classes, tc.variant.String())
	}
	
	if tc.size != TabsSizeMedium {
		classes = append(classes, tc.size.String())
	}
	
	classes = append(classes, tc.classes...)
	
	// Create tab navigation
	tabNavItems := make([]gomponents.Node, 0, len(tc.tabs))
	for _, tab := range tc.tabs {
		tabClasses := []string{"tab"}
		if tab.Active {
			tabClasses = append(tabClasses, "tab-active")
		}
		
		tabNavItems = append(tabNavItems, h.A(
			h.Class(strings.Join(tabClasses, " ")),
			h.Href("#"+tab.ID),
			gomponents.Attr("data-tab-id", tab.ID),
			gomponents.Text(tab.Label),
		))
	}
	
	// Create tab content panels
	tabContentItems := make([]gomponents.Node, 0, len(tc.tabs))
	for _, tab := range tc.tabs {
		contentClasses := []string{"tab-content"}
		if !tab.Active {
			contentClasses = append(contentClasses, "hidden")
		}
		
		tabContentItems = append(tabContentItems, h.Div(
			h.ID(tab.ID),
			h.Class(strings.Join(contentClasses, " ")),
			gomponents.Attr("data-tab-panel", tab.ID),
			tab.Content,
		))
	}
	
	// Render the complete tabs component
	tabsContainer := h.Div(
		h.ID(tc.id),
		h.Class("tabs-container"),
		gomponents.Attr("data-component", "tabs"),
		
		// Tab navigation
		h.Div(
			h.Class(strings.Join(classes, " ")),
			gomponents.Group(tabNavItems),
		),
		
		// Tab content
		h.Div(
			h.Class("tab-content-container"),
			gomponents.Group(tabContentItems),
		),
	)
	
	return tabsContainer.Render(w)
}

// NewTabItem creates a new tab item with the given ID, label, and content.
func NewTabItem(id, label string, content gomponents.Node) TabItem {
	return TabItem{
		ID:      id,
		Label:   label,
		Content: content,
		Active:  false,
	}
}

// NewActiveTabItem creates a new active tab item with the given ID, label, and content.
func NewActiveTabItem(id, label string, content gomponents.Node) TabItem {
	return TabItem{
		ID:      id,
		Label:   label,
		Content: content,
		Active:  true,
	}
}