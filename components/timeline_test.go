package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestTimeline_BasicRendering(t *testing.T) {
	t.Run("renders basic timeline", func(t *testing.T) {
		timeline := NewTimeline()
		html := renderToHTML(timeline)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		timelineEl := findElement(doc, "ul")
		if timelineEl == nil {
			t.Fatal("Expected ul element not found")
		}

		classAttr := getAttribute(timelineEl, "class")
		if !hasClass(classAttr, "timeline") {
			t.Errorf("Expected 'timeline' class, got classes: %s", classAttr)
		}
	})

	t.Run("renders timeline with content", func(t *testing.T) {
		timeline := NewTimeline(h.Li(g.Text("Timeline item")))
		html := renderToHTML(timeline)

		if !strings.Contains(html, "Timeline item") {
			t.Errorf("Expected timeline to contain 'Timeline item', got: %s", html)
		}
	})
}

func TestTimeline_OrientationModifiers(t *testing.T) {
	t.Run("applies orientation modifiers", func(t *testing.T) {
		tests := []struct {
			name        string
			orientation TimelineOrientation
			expected    string
		}{
			{"Vertical orientation", TimelineVertical, "timeline-vertical"},
			{"Horizontal orientation", TimelineHorizontal, "timeline-horizontal"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				timeline := NewTimeline().WithOrientation(tt.orientation)
				html := renderToHTML(timeline)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				timelineEl := findElement(doc, "ul")
				if timelineEl == nil {
					t.Fatal("Expected ul element not found")
				}

				classAttr := getAttribute(timelineEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestTimeline_ColorModifiers(t *testing.T) {
	t.Run("applies color modifiers", func(t *testing.T) {
		tests := []struct {
			name     string
			color    flyon.Color
			expected string
		}{
			{"Primary color", flyon.Primary, "timeline-primary"},
			{"Secondary color", flyon.Secondary, "timeline-secondary"},
			{"Success color", flyon.Success, "timeline-success"},
			{"Warning color", flyon.Warning, "timeline-warning"},
			{"Error color", flyon.Error, "timeline-error"},
			{"Info color", flyon.Info, "timeline-info"},
			{"Neutral color", flyon.Neutral, "timeline-neutral"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				timeline := NewTimeline().With(tt.color)
				html := renderToHTML(timeline)

				doc, err := parseHTML(html)
				if err != nil {
					t.Fatalf("Failed to parse HTML: %v", err)
				}

				timelineEl := findElement(doc, "ul")
				if timelineEl == nil {
					t.Fatal("Expected ul element not found")
				}

				classAttr := getAttribute(timelineEl, "class")
				if !hasClass(classAttr, tt.expected) {
					t.Errorf("Expected '%s' class, got classes: %s", tt.expected, classAttr)
				}
			})
		}
	})
}

func TestTimeline_CompactModifier(t *testing.T) {
	t.Run("applies compact modifier", func(t *testing.T) {
		timeline := NewTimeline().WithCompact()
		html := renderToHTML(timeline)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		timelineEl := findElement(doc, "ul")
		if timelineEl == nil {
			t.Fatal("Expected ul element not found")
		}

		classAttr := getAttribute(timelineEl, "class")
		if !hasClass(classAttr, "timeline-compact") {
			t.Errorf("Expected 'timeline-compact' class, got classes: %s", classAttr)
		}
	})
}

func TestTimeline_CombinedModifiers(t *testing.T) {
	t.Run("applies combined modifiers", func(t *testing.T) {
		timeline := NewTimeline(h.Li(g.Text("Item"))).WithOrientation(TimelineHorizontal).WithCompact().With(flyon.Primary)
		html := renderToHTML(timeline)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		timelineEl := findElement(doc, "ul")
		if timelineEl == nil {
			t.Fatal("Expected ul element not found")
		}

		classAttr := getAttribute(timelineEl, "class")
		if !hasClass(classAttr, "timeline") {
			t.Errorf("Expected 'timeline' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "timeline-horizontal") {
			t.Errorf("Expected 'timeline-horizontal' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "timeline-compact") {
			t.Errorf("Expected 'timeline-compact' class, got classes: %s", classAttr)
		}
		if !hasClass(classAttr, "timeline-primary") {
			t.Errorf("Expected 'timeline-primary' class, got classes: %s", classAttr)
		}
	})
}

func TestTimeline_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		timeline := NewTimeline().With(
			h.ID("main-timeline"),
			g.Attr("data-testid", "timeline-component"),
			h.Title("Timeline container"),
		)
		html := renderToHTML(timeline)

		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}

		timelineEl := findElement(doc, "ul")
		if timelineEl == nil {
			t.Fatal("Expected ul element not found")
		}

		if id := getAttribute(timelineEl, "id"); id != "main-timeline" {
			t.Errorf("Expected id='main-timeline', got id='%s'", id)
		}

		if testid := getAttribute(timelineEl, "data-testid"); testid != "timeline-component" {
			t.Errorf("Expected data-testid='timeline-component', got data-testid='%s'", testid)
		}

		if title := getAttribute(timelineEl, "title"); title != "Timeline container" {
			t.Errorf("Expected title='Timeline container', got title='%s'", title)
		}
	})
}

func TestTimeline_ComponentInterface(t *testing.T) {
	t.Run("implements flyon.Component interface", func(t *testing.T) {
		var _ flyon.Component = (*TimelineComponent)(nil)
	})

	t.Run("implements gomponents.Node interface", func(t *testing.T) {
		var _ g.Node = (*TimelineComponent)(nil)
	})

	t.Run("With method returns flyon.Component", func(t *testing.T) {
		timeline := NewTimeline()
		result := timeline.With("custom-class")
		if _, ok := result.(flyon.Component); !ok {
			t.Error("With method should return flyon.Component")
		}
	})
}