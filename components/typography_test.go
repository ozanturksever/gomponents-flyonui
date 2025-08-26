//go:build js && wasm

package components

import (
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestTypography_BasicRendering(t *testing.T) {
	tests := []struct {
		name     string
		element  func(...g.Node) *TypographyComponent
		expected string
	}{
		{"Heading 1", H1, "h1"},
		{"Heading 2", H2, "h2"},
		{"Heading 3", H3, "h3"},
		{"Heading 4", H4, "h4"},
		{"Heading 5", H5, "h5"},
		{"Heading 6", H6, "h6"},
		{"Paragraph", P, "p"},
		{"Span", Span, "span"},
		{"Strong", Strong, "strong"},
		{"Em", Em, "em"},
		{"Small", Small, "small"},
		{"Code", Code, "code"},
		{"Pre", Pre, "pre"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typo := tt.element(g.Text("Test content"))
			html := renderToHTML(typo)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			el := findElement(doc, tt.expected)
			if el == nil {
				t.Fatalf("Expected %s element not found", tt.expected)
			}

			if el.FirstChild == nil || el.FirstChild.Data != "Test content" {
				t.Errorf("Expected text content 'Test content', got: %v", el.FirstChild)
			}
		})
	}
}

func TestTypography_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small text", flyon.SizeXS, "text-xs"},
		{"Small text", flyon.SizeSmall, "text-sm"},
		{"Medium text", flyon.SizeMedium, "text-md"},
		{"Large text", flyon.SizeLarge, "text-lg"},
		{"Extra large text", flyon.SizeXL, "text-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typo := NewTypography("p", g.Text("Content")).With(tt.size)
			html := renderToHTML(typo)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			el := findElement(doc, "p")
			if el == nil {
				t.Fatal("Expected p element not found")
			}

			classAttr := getAttribute(el, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestTypography_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary text", flyon.Primary, "text-primary"},
		{"Secondary text", flyon.Secondary, "text-secondary"},
		{"Neutral text", flyon.Neutral, "text-neutral"},
		{"Info text", flyon.Info, "text-info"},
		{"Success text", flyon.Success, "text-success"},
		{"Warning text", flyon.Warning, "text-warning"},
		{"Error text", flyon.Error, "text-error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typo := NewTypography("p", g.Text("Content")).With(tt.color)
			html := renderToHTML(typo)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			el := findElement(doc, "p")
			if el == nil {
				t.Fatal("Expected p element not found")
			}

			classAttr := getAttribute(el, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestTypography_WeightModifiers(t *testing.T) {
	tests := []struct {
		name          string
		weight        string
		expectedClass string
	}{
		{"Thin weight", "thin", "font-thin"},
		{"Light weight", "light", "font-light"},
		{"Normal weight", "normal", "font-normal"},
		{"Medium weight", "medium", "font-medium"},
		{"Semibold weight", "semibold", "font-semibold"},
		{"Bold weight", "bold", "font-bold"},
		{"Extrabold weight", "extrabold", "font-extrabold"},
		{"Black weight", "black", "font-black"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typo := NewTypography("p", g.Text("Content")).WithWeight(tt.weight)
			html := renderToHTML(typo)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			el := findElement(doc, "p")
			if el == nil {
				t.Fatal("Expected p element not found")
			}

			classAttr := getAttribute(el, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestTypography_AlignmentModifiers(t *testing.T) {
	tests := []struct {
		name          string
		alignment     string
		expectedClass string
	}{
		{"Left alignment", "left", "text-left"},
		{"Center alignment", "center", "text-center"},
		{"Right alignment", "right", "text-right"},
		{"Justify alignment", "justify", "text-justify"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typo := NewTypography("p", g.Text("Content")).WithAlign(tt.alignment)
			html := renderToHTML(typo)

			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			el := findElement(doc, "p")
			if el == nil {
				t.Fatal("Expected p element not found")
			}

			classAttr := getAttribute(el, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestTypography_CombinedModifiers(t *testing.T) {
	typo := NewTypography("h1", g.Text("Title")).With(
		flyon.Primary,
		flyon.SizeLarge,
		"custom-class",
	).(*TypographyComponent).WithWeight("bold").WithAlign("center")

	html := renderToHTML(typo)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	h1El := findElement(doc, "h1")
	if h1El == nil {
		t.Fatal("Expected h1 element not found")
	}

	classAttr := getAttribute(h1El, "class")
	expectedClasses := []string{"text-primary", "text-lg", "font-bold", "text-center", "custom-class"}
	for _, expectedClass := range expectedClasses {
		if !hasClass(classAttr, expectedClass) {
			t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
		}
	}
}

func TestTypography_HTMLAttributes(t *testing.T) {
	typo := NewTypography("p", g.Text("Content")).With(
		h.ID("test-id"),
		g.Attr("data-test", "value"),
	)

	html := renderToHTML(typo)

	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	pEl := findElement(doc, "p")
	if pEl == nil {
		t.Fatal("Expected p element not found")
	}

	if !hasAttribute(pEl, "id") {
		t.Error("Expected id attribute")
	}

	if getAttribute(pEl, "id") != "test-id" {
		t.Errorf("Expected id 'test-id', got: %s", getAttribute(pEl, "id"))
	}

	if !hasAttribute(pEl, "data-test") {
		t.Error("Expected data-test attribute")
	}

	if getAttribute(pEl, "data-test") != "value" {
		t.Errorf("Expected data-test 'value', got: %s", getAttribute(pEl, "data-test"))
	}
}

func TestTypography_ComponentInterface(t *testing.T) {
	var _ flyon.Component = NewTypography("p", g.Text("Test"))
}