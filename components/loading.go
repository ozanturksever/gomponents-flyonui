package components

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// LoadingType represents the type of loading animation
type LoadingType string

const (
	LoadingSpinner   LoadingType = "loading-spinner"
	LoadingDots      LoadingType = "loading-dots"
	LoadingRing      LoadingType = "loading-ring"
	LoadingBall      LoadingType = "loading-ball"
	LoadingBars      LoadingType = "loading-bars"
	LoadingInfinity  LoadingType = "loading-infinity"
)

// LoadingComponent represents a loading indicator component
type LoadingComponent struct {
	classes    []string
	attributes []g.Node
	children   []g.Node
	loadingType LoadingType
}

// NewLoading creates a new loading component
func NewLoading(children ...g.Node) *LoadingComponent {
	return &LoadingComponent{
		classes:  []string{"loading"},
		children: children,
	}
}

// WithType sets the loading animation type
func (l *LoadingComponent) WithType(loadingType LoadingType) *LoadingComponent {
	newLoading := &LoadingComponent{
		classes:     make([]string, len(l.classes)),
		attributes:  make([]g.Node, len(l.attributes)),
		children:    make([]g.Node, len(l.children)),
		loadingType: loadingType,
	}
	copy(newLoading.classes, l.classes)
	copy(newLoading.attributes, l.attributes)
	copy(newLoading.children, l.children)
	return newLoading
}

// With applies modifiers to the loading component
func (l *LoadingComponent) With(items ...any) flyon.Component {
	newLoading := &LoadingComponent{
		classes:     make([]string, len(l.classes)),
		attributes:  make([]g.Node, len(l.attributes)),
		children:    make([]g.Node, len(l.children)),
		loadingType: l.loadingType,
	}
	copy(newLoading.classes, l.classes)
	copy(newLoading.attributes, l.attributes)
	copy(newLoading.children, l.children)

	for _, item := range items {
		switch v := item.(type) {
		case flyon.Color:
			newLoading.classes = append(newLoading.classes, "loading-"+v.String())
		case flyon.Size:
			newLoading.classes = append(newLoading.classes, "loading-"+v.String())
		case string:
			newLoading.classes = append(newLoading.classes, v)
		case g.Node:
			newLoading.attributes = append(newLoading.attributes, v)
		}
	}

	return newLoading
}

// Render renders the loading component to HTML
func (l *LoadingComponent) Render(w io.Writer) error {
	classes := make([]string, len(l.classes))
	copy(classes, l.classes)
	
	// Add loading type class if specified
	if l.loadingType != "" {
		classes = append(classes, string(l.loadingType))
	}
	
	nodes := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	nodes = append(nodes, l.attributes...)
	nodes = append(nodes, l.children...)
	
	return h.Span(nodes...).Render(w)
}