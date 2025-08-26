package components

import (
	"io"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// SelectOption represents an option in a select dropdown
type SelectOption struct {
	Value    string
	Label    string
	Selected bool
	Disabled bool
}

// SelectComponent represents a select dropdown with FlyonUI styling
type SelectComponent struct {
	id       string
	name     string
	value    string
	disabled bool
	required bool
	multiple bool
	size     int
	color    flyon.Color
	compSize flyon.Size
	options  []SelectOption
	classes  []string
}

// NewSelect creates a new select component
func NewSelect() *SelectComponent {
	return &SelectComponent{
		color:    flyon.Primary,
		compSize: flyon.SizeMedium,
		options:  make([]SelectOption, 0),
	}
}

// WithID sets the select ID
func (s *SelectComponent) WithID(id string) *SelectComponent {
	newSelect := s.copy()
	newSelect.id = id
	return newSelect
}

// WithName sets the select name attribute
func (s *SelectComponent) WithName(name string) *SelectComponent {
	newSelect := s.copy()
	newSelect.name = name
	return newSelect
}

// WithValue sets the select value
func (s *SelectComponent) WithValue(value string) *SelectComponent {
	newSelect := s.copy()
	newSelect.value = value
	return newSelect
}

// WithDisabled sets the select disabled state
func (s *SelectComponent) WithDisabled(disabled bool) *SelectComponent {
	newSelect := s.copy()
	newSelect.disabled = disabled
	return newSelect
}

// WithRequired sets the select required state
func (s *SelectComponent) WithRequired(required bool) *SelectComponent {
	newSelect := s.copy()
	newSelect.required = required
	return newSelect
}

// WithMultiple sets the select multiple state
func (s *SelectComponent) WithMultiple(multiple bool) *SelectComponent {
	newSelect := s.copy()
	newSelect.multiple = multiple
	return newSelect
}

// WithSize sets the select size (number of visible options)
func (s *SelectComponent) WithSize(size int) *SelectComponent {
	newSelect := s.copy()
	newSelect.size = size
	return newSelect
}

// WithColor sets the select color
func (s *SelectComponent) WithColor(color flyon.Color) *SelectComponent {
	newSelect := s.copy()
	newSelect.color = color
	return newSelect
}

// WithCompSize sets the select component size
func (s *SelectComponent) WithCompSize(size flyon.Size) *SelectComponent {
	newSelect := s.copy()
	newSelect.compSize = size
	return newSelect
}

// WithOption adds an option to the select
func (s *SelectComponent) WithOption(value, label string) *SelectComponent {
	newSelect := s.copy()
	newSelect.options = append(newSelect.options, SelectOption{
		Value: value,
		Label: label,
	})
	return newSelect
}

// WithSelectedOption adds a selected option to the select
func (s *SelectComponent) WithSelectedOption(value, label string) *SelectComponent {
	newSelect := s.copy()
	newSelect.options = append(newSelect.options, SelectOption{
		Value:    value,
		Label:    label,
		Selected: true,
	})
	return newSelect
}

// WithDisabledOption adds a disabled option to the select
func (s *SelectComponent) WithDisabledOption(value, label string) *SelectComponent {
	newSelect := s.copy()
	newSelect.options = append(newSelect.options, SelectOption{
		Value:    value,
		Label:    label,
		Disabled: true,
	})
	return newSelect
}

// WithOptions sets multiple options at once
func (s *SelectComponent) WithOptions(options []SelectOption) *SelectComponent {
	newSelect := s.copy()
	newSelect.options = make([]SelectOption, len(options))
	copy(newSelect.options, options)
	return newSelect
}

// WithClasses adds additional CSS classes
func (s *SelectComponent) WithClasses(classes ...string) *SelectComponent {
	newSelect := s.copy()
	newSelect.classes = append(newSelect.classes, classes...)
	return newSelect
}

// With applies modifiers to the select
func (s *SelectComponent) With(modifiers ...any) flyon.Component {
	newSelect := s.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newSelect.color = m
		case flyon.Size:
			newSelect.compSize = m
		case string:
			newSelect.classes = append(newSelect.classes, m)
		}
	}
	return newSelect
}

// copy creates a deep copy of the select component
func (s *SelectComponent) copy() *SelectComponent {
	newSelect := *s
	newSelect.options = make([]SelectOption, len(s.options))
	copy(newSelect.options, s.options)
	newSelect.classes = make([]string, len(s.classes))
	copy(newSelect.classes, s.classes)
	return &newSelect
}

// Render generates the HTML for the select
func (s *SelectComponent) Render(w io.Writer) error {
	classes := []string{"select", "select-bordered"}
	
	// Add color class
	if s.color != flyon.Primary {
		classes = append(classes, "select-"+s.color.String())
	}
	
	// Add size class
	if s.compSize != flyon.SizeMedium {
		classes = append(classes, "select-"+s.compSize.String())
	}
	
	// Add custom classes
	classes = append(classes, s.classes...)
	
	// Build attributes
	attrs := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	
	if s.id != "" {
		attrs = append(attrs, h.ID(s.id))
	}
	
	if s.name != "" {
		attrs = append(attrs, h.Name(s.name))
	}
	
	if s.value != "" {
		attrs = append(attrs, h.Value(s.value))
	}
	
	if s.disabled {
		attrs = append(attrs, h.Disabled())
	}
	
	if s.required {
		attrs = append(attrs, h.Required())
	}
	
	if s.multiple {
		attrs = append(attrs, h.Multiple())
	}
	
	if s.size > 0 {
		attrs = append(attrs, g.Attr("size", strconv.Itoa(s.size)))
	}
	
	// Build options
	optionNodes := make([]g.Node, 0, len(s.options))
	for _, option := range s.options {
		optionAttrs := []g.Node{
			h.Value(option.Value),
		}
		
		if option.Selected {
			optionAttrs = append(optionAttrs, h.Selected())
		}
		
		if option.Disabled {
			optionAttrs = append(optionAttrs, h.Disabled())
		}
		
		optionNodes = append(optionNodes, h.Option(append(optionAttrs, g.Text(option.Label))...))
	}
	
	attrs = append(attrs, optionNodes...)
	
	return h.Select(attrs...).Render(w)
}