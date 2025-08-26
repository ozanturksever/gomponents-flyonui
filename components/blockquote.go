package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// BlockquoteComponent represents a blockquote component
type BlockquoteComponent struct {
	children   []g.Node
	attributes []g.Node
	classes    []string
	author     string
	source     string
}

// NewBlockquote creates a new blockquote component
func NewBlockquote(children ...g.Node) *BlockquoteComponent {
	// Separate attributes from content children
	var attributes []g.Node
	var content []g.Node
	
	for _, child := range children {
		// Check if this is an attribute by trying to render it and seeing if it produces attribute-like output
		var buf strings.Builder
		if err := child.Render(&buf); err == nil {
			output := buf.String()
			// If the output contains '=' it's likely an attribute
			if strings.Contains(output, "=") || strings.Contains(output, "disabled") {
				attributes = append(attributes, child)
			} else {
				content = append(content, child)
			}
		} else {
			// If we can't render it, assume it's content
			content = append(content, child)
		}
	}
	
	return &BlockquoteComponent{
		children:   content,
		attributes: attributes,
		classes:    []string{"blockquote"}, // Default FlyonUI blockquote class
	}
}

// Blockquote creates a new blockquote component (convenience function)
func Blockquote(children ...g.Node) *BlockquoteComponent {
	return NewBlockquote(children...)
}

// WithAuthor adds an author attribution to the blockquote
func (b *BlockquoteComponent) WithAuthor(author string) *BlockquoteComponent {
	newBlockquote := &BlockquoteComponent{
		children:   make([]g.Node, len(b.children)),
		attributes: make([]g.Node, len(b.attributes)),
		classes:    make([]string, len(b.classes)),
		author:     author,
		source:     b.source,
	}
	copy(newBlockquote.children, b.children)
	copy(newBlockquote.attributes, b.attributes)
	copy(newBlockquote.classes, b.classes)
	return newBlockquote
}

// WithSource adds a source attribution to the blockquote
func (b *BlockquoteComponent) WithSource(source string) *BlockquoteComponent {
	newBlockquote := &BlockquoteComponent{
		children:   make([]g.Node, len(b.children)),
		attributes: make([]g.Node, len(b.attributes)),
		classes:    make([]string, len(b.classes)),
		author:     b.author,
		source:     source,
	}
	copy(newBlockquote.children, b.children)
	copy(newBlockquote.attributes, b.attributes)
	copy(newBlockquote.classes, b.classes)
	return newBlockquote
}

// With adds modifiers to the blockquote
func (b *BlockquoteComponent) With(modifiers ...any) flyon.Component {
	newBlockquote := &BlockquoteComponent{
		children:   make([]g.Node, len(b.children)),
		attributes: make([]g.Node, len(b.attributes)),
		classes:    make([]string, len(b.classes)),
		author:     b.author,
		source:     b.source,
	}
	copy(newBlockquote.children, b.children)
	copy(newBlockquote.attributes, b.attributes)
	copy(newBlockquote.classes, b.classes)
	
	// Apply each modifier
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newBlockquote.classes = append(newBlockquote.classes, "blockquote-"+m.String())
		case flyon.Size:
			newBlockquote.classes = append(newBlockquote.classes, "blockquote-"+m.String())
		case flyon.Variant:
			newBlockquote.classes = append(newBlockquote.classes, "blockquote-"+m.String())
		case string:
			// Treat string modifiers as additional CSS classes
			newBlockquote.classes = append(newBlockquote.classes, m)
		case g.Node:
			// Treat Node modifiers as additional attributes
			newBlockquote.attributes = append(newBlockquote.attributes, m)
		}
	}
	
	return newBlockquote
}

// Render renders the blockquote component
func (b *BlockquoteComponent) Render(w io.Writer) error {
	// Build the class attribute
	classAttr := h.Class(strings.Join(b.classes, " "))
	
	// Build the blockquote content
	content := make([]g.Node, 0, len(b.children)+1)
	content = append(content, b.children...)
	
	// Add author and/or source if provided
	if b.author != "" || b.source != "" {
		citeContent := make([]g.Node, 0, 3)
		
		if b.author != "" {
			citeContent = append(citeContent, g.Text(b.author))
		}
		
		if b.author != "" && b.source != "" {
			citeContent = append(citeContent, g.Text(", "))
		}
		
		if b.source != "" {
			citeContent = append(citeContent, g.Text(b.source))
		}
		
		cite := h.Cite(citeContent...)
		content = append(content, cite)
	}
	
	// Create the blockquote element with class, attributes, and children
	allNodes := make([]g.Node, 0, len(b.attributes)+len(content)+1)
	allNodes = append(allNodes, classAttr)
	allNodes = append(allNodes, b.attributes...)
	allNodes = append(allNodes, content...)
	
	blockquoteEl := h.BlockQuote(allNodes...)
	
	return blockquoteEl.Render(w)
}

// Ensure BlockquoteComponent implements the required interfaces
var (
	_ flyon.Component = (*BlockquoteComponent)(nil)
	_ g.Node          = (*BlockquoteComponent)(nil)
)