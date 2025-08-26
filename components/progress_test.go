//go:build js && wasm

package components

import (
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestProgress_BasicRendering(t *testing.T) {
	t.Run("renders basic progress with default classes", func(t *testing.T) {
		progress := NewProgress(50)
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		classAttr := getAttribute(progressEl, "class")
		if !hasClass(classAttr, "progress") {
			t.Errorf("Expected 'progress' class, got classes: %s", classAttr)
		}
		
		valueAttr := getAttribute(progressEl, "value")
		if valueAttr != "50" {
			t.Errorf("Expected value='50', got value='%s'", valueAttr)
		}
		
		maxAttr := getAttribute(progressEl, "max")
		if maxAttr != "100" {
			t.Errorf("Expected max='100', got max='%s'", maxAttr)
		}
	})
	
	t.Run("renders progress with custom max value", func(t *testing.T) {
		progress := NewProgressWithMax(75, 200)
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		valueAttr := getAttribute(progressEl, "value")
		if valueAttr != "75" {
			t.Errorf("Expected value='75', got value='%s'", valueAttr)
		}
		
		maxAttr := getAttribute(progressEl, "max")
		if maxAttr != "200" {
			t.Errorf("Expected max='200', got max='%s'", maxAttr)
		}
	})
	
	t.Run("renders indeterminate progress", func(t *testing.T) {
		progress := NewIndeterminateProgress()
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		// Indeterminate progress should not have value attribute
		valueAttr := getAttribute(progressEl, "value")
		if valueAttr != "" {
			t.Errorf("Expected no value attribute for indeterminate progress, got value='%s'", valueAttr)
		}
		
		maxAttr := getAttribute(progressEl, "max")
		if maxAttr != "100" {
			t.Errorf("Expected max='100', got max='%s'", maxAttr)
		}
	})
}

func TestProgress_ColorModifiers(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected string
	}{
		{"Primary color", flyon.Primary, "progress-primary"},
		{"Secondary color", flyon.Secondary, "progress-secondary"},
		{"Success color", flyon.Success, "progress-success"},
		{"Warning color", flyon.Warning, "progress-warning"},
		{"Error color", flyon.Error, "progress-error"},
		{"Info color", flyon.Info, "progress-info"},
		{"Neutral color", flyon.Neutral, "progress-neutral"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			progress := NewProgress(60).With(tt.color)
			html := renderToHTML(progress)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			progressEl := findElement(doc, "progress")
			if progressEl == nil {
				t.Fatal("Expected progress element not found")
			}
			
			classAttr := getAttribute(progressEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
			if !hasClass(classAttr, "progress") {
				t.Errorf("Expected 'progress' base class, got classes: %s", classAttr)
			}
		})
	}
}

func TestProgress_SizeModifiers(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected string
	}{
		{"Extra small size", flyon.SizeXS, "progress-xs"},
		{"Small size", flyon.SizeSmall, "progress-sm"},
		{"Medium size", flyon.SizeMedium, "progress-md"},
		{"Large size", flyon.SizeLarge, "progress-lg"},
		{"Extra large size", flyon.SizeXL, "progress-xl"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			progress := NewProgress(40).With(tt.size)
			html := renderToHTML(progress)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			progressEl := findElement(doc, "progress")
			if progressEl == nil {
				t.Fatal("Expected progress element not found")
			}
			
			classAttr := getAttribute(progressEl, "class")
			if !hasClass(classAttr, tt.expected) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
			}
		})
	}
}

func TestProgress_CombinedModifiers(t *testing.T) {
	t.Run("applies multiple modifiers correctly", func(t *testing.T) {
		progress := NewProgress(80).With(
			flyon.Success,
			flyon.SizeLarge,
			"custom-class",
		)
		
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		classAttr := getAttribute(progressEl, "class")
		expectedClasses := []string{"progress", "progress-success", "progress-lg", "custom-class"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
		
		valueAttr := getAttribute(progressEl, "value")
		if valueAttr != "80" {
			t.Errorf("Expected value='80', got value='%s'", valueAttr)
		}
	})
}

func TestProgress_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		progress := NewProgress(90).With(
			h.ID("test-progress"),
			g.Attr("data-test", "progress-component"),
			h.Title("Loading Progress"),
		)
		
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		if id := getAttribute(progressEl, "id"); id != "test-progress" {
			t.Errorf("Expected id='test-progress', got id='%s'", id)
		}
		
		if testAttr := getAttribute(progressEl, "data-test"); testAttr != "progress-component" {
			t.Errorf("Expected data-test='progress-component', got data-test='%s'", testAttr)
		}
		
		if title := getAttribute(progressEl, "title"); title != "Loading Progress" {
			t.Errorf("Expected title='Loading Progress', got title='%s'", title)
		}
	})
}

func TestProgress_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		var _ flyon.Component = NewProgress(25)
		
		progress := NewProgress(25)
		html := renderToHTML(progress)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		progressEl := findElement(doc, "progress")
		if progressEl == nil {
			t.Fatal("Expected progress element not found")
		}
		
		valueAttr := getAttribute(progressEl, "value")
		if valueAttr != "25" {
			t.Errorf("Expected value='25', got value='%s'", valueAttr)
		}
	})
}