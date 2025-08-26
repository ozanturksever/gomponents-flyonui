package components

import (
	"io"
	"strings"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// SwapComponent represents a swap component that can toggle between two states
type SwapComponent struct {
	id       string
	active   bool
	rotate   bool
	flip     bool
	color    flyon.Color
	classes  []string
	onState  g.Node
	offState g.Node
}

// NewSwap creates a new swap component with on and off states
func NewSwap(onState, offState g.Node) *SwapComponent {
	return &SwapComponent{
		onState:  onState,
		offState: offState,
		color:    flyon.Primary,
	}
}

// WithID sets the ID of the swap component
func (s *SwapComponent) WithID(id string) *SwapComponent {
	newSwap := s.copy()
	newSwap.id = id
	return newSwap
}

// WithActive sets the active state of the swap component
func (s *SwapComponent) WithActive(active bool) *SwapComponent {
	newSwap := s.copy()
	newSwap.active = active
	return newSwap
}

// WithRotate enables rotation animation for the swap
func (s *SwapComponent) WithRotate(rotate bool) *SwapComponent {
	newSwap := s.copy()
	newSwap.rotate = rotate
	return newSwap
}

// WithFlip enables flip animation for the swap
func (s *SwapComponent) WithFlip(flip bool) *SwapComponent {
	newSwap := s.copy()
	newSwap.flip = flip
	return newSwap
}

// WithColor sets the color of the swap component
func (s *SwapComponent) WithColor(color flyon.Color) *SwapComponent {
	newSwap := s.copy()
	newSwap.color = color
	return newSwap
}

// WithClasses adds additional CSS classes to the swap component
func (s *SwapComponent) WithClasses(classes ...string) *SwapComponent {
	newSwap := s.copy()
	newSwap.classes = append(newSwap.classes, classes...)
	return newSwap
}

// With applies modifiers to the swap component
func (s *SwapComponent) With(modifiers ...any) flyon.Component {
	newSwap := s.copy()
	for _, modifier := range modifiers {
		switch m := modifier.(type) {
		case flyon.Color:
			newSwap.color = m
		}
	}
	return newSwap
}

// copy creates a deep copy of the swap component
func (s *SwapComponent) copy() *SwapComponent {
	newClasses := make([]string, len(s.classes))
	copy(newClasses, s.classes)
	return &SwapComponent{
		id:       s.id,
		active:   s.active,
		rotate:   s.rotate,
		flip:     s.flip,
		color:    s.color,
		classes:  newClasses,
		onState:  s.onState,
		offState: s.offState,
	}
}

// Render renders the swap component to the provided writer
func (s *SwapComponent) Render(w io.Writer) error {
	// Build CSS classes
	classes := []string{"swap"}
	
	if s.rotate {
		classes = append(classes, "swap-rotate")
	}
	
	if s.flip {
		classes = append(classes, "swap-flip")
	}
	
	if s.color != flyon.Primary {
		classes = append(classes, "swap-"+s.color.String())
	}
	
	if len(s.classes) > 0 {
		classes = append(classes, s.classes...)
	}
	
	// Build attributes
	attrs := []g.Node{
		h.Class(strings.Join(classes, " ")),
	}
	
	if s.id != "" {
		attrs = append(attrs, h.ID(s.id))
	}
	
	// Create checkbox input for state management
	checkboxAttrs := []g.Node{
		h.Type("checkbox"),
		h.Class("swap-input"),
	}
	
	if s.active {
		checkboxAttrs = append(checkboxAttrs, h.Checked())
	}
	
	// Build the swap component
	return h.Label(
		g.Group(attrs),
		h.Input(g.Group(checkboxAttrs)),
		h.Div(
			h.Class("swap-on"),
			s.onState,
		),
		h.Div(
			h.Class("swap-off"),
			s.offState,
		),
	).Render(w)
}