package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
)

func TestNewAccordion(t *testing.T) {
	item1 := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	item2 := NewAccordionItem("item2", "Item 2", gomponents.Text("Content 2"))
	accordion := NewAccordion(item1, item2)

	if accordion == nil {
		t.Fatal("NewAccordion() returned nil")
	}

	if len(accordion.items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(accordion.items))
	}

	if accordion.multiple {
		t.Error("Expected multiple to be false by default")
	}

	if accordion.color != flyon.Primary {
		t.Errorf("Expected primary color, got %v", accordion.color)
	}
}

func TestAccordionComponent_WithID(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	accordion := NewAccordion(item)
	modified := accordion.WithID("custom-accordion")

	if modified.id != "custom-accordion" {
		t.Errorf("Expected ID 'custom-accordion', got %s", modified.id)
	}

	// Test immutability
	if accordion.id == "custom-accordion" {
		t.Error("Original accordion component was modified")
	}
}

func TestAccordionComponent_WithMultiple(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	accordion := NewAccordion(item)
	modified := accordion.WithMultiple(true)

	if !modified.multiple {
		t.Error("Expected multiple to be true")
	}

	// Test immutability
	if accordion.multiple {
		t.Error("Original accordion component was modified")
	}
}

func TestAccordionComponent_WithColor(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	accordion := NewAccordion(item)
	modified := accordion.WithColor(flyon.Secondary)

	if modified.color != flyon.Secondary {
		t.Errorf("Expected Secondary color, got %v", modified.color)
	}

	// Test immutability
	if accordion.color != flyon.Primary {
		t.Error("Original accordion component was modified")
	}
}

func TestAccordionComponent_WithClasses(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	accordion := NewAccordion(item)
	modified := accordion.WithClasses("custom-class", "another-class")

	if len(modified.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(modified.classes))
	}

	if modified.classes[0] != "custom-class" || modified.classes[1] != "another-class" {
		t.Errorf("Expected custom classes, got %v", modified.classes)
	}

	// Test immutability
	if len(accordion.classes) != 0 {
		t.Error("Original accordion component was modified")
	}
}

func TestAccordionComponent_With(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	accordion := NewAccordion(item)

	// Test with flyon.Color
	modified1 := accordion.With(flyon.Warning)
	modified1Accordion := modified1.(*AccordionComponent)
	if modified1Accordion.color != flyon.Warning {
		t.Errorf("Expected Warning color, got %v", modified1Accordion.color)
	}

	// Test with multiple modifiers
	modified2 := accordion.With(flyon.Success)
	modified2Accordion := modified2.(*AccordionComponent)
	if modified2Accordion.color != flyon.Success {
		t.Errorf("Expected Success color, got %v", modified2Accordion.color)
	}
}

func TestAccordionComponent_Render(t *testing.T) {
	item1 := NewOpenAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	item2 := NewAccordionItem("item2", "Item 2", gomponents.Text("Content 2"))
	accordion := NewAccordion(item1, item2).WithID("test-accordion")

	var buf strings.Builder
	err := accordion.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()

	// Check container structure
	if !strings.Contains(html, `id="test-accordion"`) {
		t.Error("Missing accordion ID")
	}
	if !strings.Contains(html, `class="collapse-group"`) {
		t.Error("Missing collapse-group class")
	}
	if !strings.Contains(html, `data-component="accordion"`) {
		t.Error("Missing data-component attribute")
	}
	if !strings.Contains(html, `data-single="true"`) {
		t.Error("Missing data-single attribute for single mode")
	}

	// Check accordion items
	if !strings.Contains(html, `class="collapse collapse-arrow"`) {
		t.Error("Missing collapse class")
	}
	if !strings.Contains(html, `data-accordion-item="item1"`) {
		t.Error("Missing item1 data-accordion-item")
	}
	if !strings.Contains(html, `data-accordion-item="item2"`) {
		t.Error("Missing item2 data-accordion-item")
	}

	// Check radio inputs for single mode
	if !strings.Contains(html, `type="radio"`) {
		t.Error("Missing radio input type")
	}
	if !strings.Contains(html, `name="test-accordion-accordion"`) {
		t.Error("Missing radio input name")
	}
	if !strings.Contains(html, `class="collapse-toggle"`) {
		t.Error("Missing collapse-toggle class")
	}
	if !strings.Contains(html, `checked`) {
		t.Error("Missing checked attribute for open item")
	}

	// Check labels and content
	if !strings.Contains(html, `class="collapse-title text-xl font-medium cursor-pointer"`) {
		t.Error("Missing collapse-title class")
	}
	if !strings.Contains(html, `for="item1"`) {
		t.Error("Missing for attribute for item1")
	}
	if !strings.Contains(html, `for="item2"`) {
		t.Error("Missing for attribute for item2")
	}
	if !strings.Contains(html, "Item 1") {
		t.Error("Missing Item 1 title")
	}
	if !strings.Contains(html, "Item 2") {
		t.Error("Missing Item 2 title")
	}

	// Check content structure
	if !strings.Contains(html, `class="collapse-content"`) {
		t.Error("Missing collapse-content class")
	}
	if !strings.Contains(html, `class="pb-2"`) {
		t.Error("Missing pb-2 class")
	}
	if !strings.Contains(html, "Content 1") {
		t.Error("Missing Content 1 text")
	}
	if !strings.Contains(html, "Content 2") {
		t.Error("Missing Content 2 text")
	}
}

func TestAccordionComponent_RenderMultiple(t *testing.T) {
	item1 := NewOpenAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	item2 := NewAccordionItem("item2", "Item 2", gomponents.Text("Content 2"))
	accordion := NewAccordion(item1, item2).WithID("test-accordion").WithMultiple(true)

	var buf strings.Builder
	err := accordion.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()

	// Check multiple mode attributes
	if !strings.Contains(html, `data-multiple="true"`) {
		t.Error("Missing data-multiple attribute for multiple mode")
	}
	if strings.Contains(html, `data-single="true"`) {
		t.Error("Should not have data-single attribute in multiple mode")
	}

	// Check checkbox inputs for multiple mode
	if !strings.Contains(html, `type="checkbox"`) {
		t.Error("Missing checkbox input type")
	}
	if strings.Contains(html, `type="radio"`) {
		t.Error("Should not have radio input type in multiple mode")
	}

	// Check unique names for each checkbox
	if !strings.Contains(html, `name="test-accordion-accordion-item1"`) {
		t.Error("Missing unique name for item1 checkbox")
	}
	if !strings.Contains(html, `name="test-accordion-accordion-item2"`) {
		t.Error("Missing unique name for item2 checkbox")
	}
}

func TestNewAccordionItem(t *testing.T) {
	content := gomponents.Text("Test content")
	item := NewAccordionItem("test-id", "Test Title", content)

	if item.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", item.ID)
	}
	if item.Title != "Test Title" {
		t.Errorf("Expected Title 'Test Title', got %s", item.Title)
	}
	if item.Content != content {
		t.Error("Content mismatch")
	}
	if item.Open {
		t.Error("Expected item to be closed by default")
	}
}

func TestNewOpenAccordionItem(t *testing.T) {
	content := gomponents.Text("Test content")
	item := NewOpenAccordionItem("test-id", "Test Title", content)

	if item.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", item.ID)
	}
	if item.Title != "Test Title" {
		t.Errorf("Expected Title 'Test Title', got %s", item.Title)
	}
	if item.Content != content {
		t.Error("Content mismatch")
	}
	if !item.Open {
		t.Error("Expected item to be open")
	}
}

func TestAccordionComponent_Immutability(t *testing.T) {
	item := NewAccordionItem("item1", "Item 1", gomponents.Text("Content 1"))
	original := NewAccordion(item)
	originalID := original.id
	originalMultiple := original.multiple
	originalColor := original.color
	originalClassesLen := len(original.classes)

	// Apply various modifications
	_ = original.WithID("new-id")
	_ = original.WithMultiple(true)
	_ = original.WithColor(flyon.Secondary)
	_ = original.WithClasses("new-class")
	_ = original.With(flyon.Warning)

	// Verify original is unchanged
	if original.id != originalID {
		t.Error("Original ID was modified")
	}
	if original.multiple != originalMultiple {
		t.Error("Original multiple was modified")
	}
	if original.color != originalColor {
		t.Error("Original color was modified")
	}
	if len(original.classes) != originalClassesLen {
		t.Error("Original classes were modified")
	}
}