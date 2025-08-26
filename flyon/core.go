// Package flyon provides the core interfaces and types for the gomponents-flyonui library.
// It defines the Component interface and type-safe modifier system for building FlyonUI components.
package flyon

import "maragu.dev/gomponents"

// Modifier represents a type that can modify component properties.
// All modifier types (Color, Size, Variant) implement this interface.
type Modifier interface {
	// String returns the string representation for CSS class generation
	String() string
}

// Component represents a FlyonUI component that can be rendered and modified.
// It extends gomponents.Node to provide HTML rendering capabilities while adding
// a fluent builder pattern through the With method.
type Component interface {
	gomponents.Node
	// With applies modifiers to create a new component instance with updated configuration.
	// This follows the builder pattern and ensures immutability.
	With(modifiers ...Modifier) Component
}

// Color represents the color variants available in FlyonUI.
// These map directly to FlyonUI's color system and CSS classes.
type Color int

const (
	Primary Color = iota
	Secondary
	Success
	Warning
	Error
	Info
	Neutral
)

// String returns the string representation of the color for use in CSS classes.
func (c Color) String() string {
	switch c {
	case Primary:
		return "primary"
	case Secondary:
		return "secondary"
	case Success:
		return "success"
	case Warning:
		return "warning"
	case Error:
		return "error"
	case Info:
		return "info"
	case Neutral:
		return "neutral"
	default:
		return "primary"
	}
}

// ApplyToMock implements MockModifier for testing purposes.
func (c Color) ApplyToMock(config *mockConfig) {
	config.Color = c
}

// Size represents the size variants available in FlyonUI components.
type Size int

const (
	SizeXS Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
	SizeXL
)

// String returns the string representation of the size for use in CSS classes.
func (s Size) String() string {
	switch s {
	case SizeXS:
		return "xs"
	case SizeSmall:
		return "sm"
	case SizeMedium:
		return "md"
	case SizeLarge:
		return "lg"
	case SizeXL:
		return "xl"
	default:
		return "md"
	}
}

// ApplyToMock implements MockModifier for testing purposes.
func (s Size) ApplyToMock(config *mockConfig) {
	config.Size = s
}

// Variant represents the visual style variants available in FlyonUI components.
type Variant int

const (
	VariantSolid Variant = iota
	VariantOutline
	VariantGhost
	VariantSoft
)

// String returns the string representation of the variant for use in CSS classes.
func (v Variant) String() string {
	switch v {
	case VariantSolid:
		return "solid"
	case VariantOutline:
		return "outline"
	case VariantGhost:
		return "ghost"
	case VariantSoft:
		return "soft"
	default:
		return "solid"
	}
}

// ApplyToMock implements MockModifier for testing purposes.
func (v Variant) ApplyToMock(config *mockConfig) {
	config.Variant = v
}

// mockConfig is used for testing the modifier system.
// It's defined here to avoid circular imports in tests.
type mockConfig struct {
	Color   Color
	Size    Size
	Variant Variant
}

// MockModifier is a test interface for validating modifier application.
type MockModifier interface {
	ApplyToMock(*mockConfig)
}