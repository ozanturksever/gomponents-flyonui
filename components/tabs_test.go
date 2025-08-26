package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
)

func TestTabsVariant_String(t *testing.T) {
	tests := []struct {
		name     string
		variant  TabsVariant
		expected string
	}{
		{"default", TabsDefault, ""},
		{"bordered", TabsBordered, "tabs-bordered"},
		{"lifted", TabsLifted, "tabs-lifted"},
		{"boxed", TabsBoxed, "tabs-boxed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.variant.String(); got != tt.expected {
				t.Errorf("TabsVariant.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTabsSize_String(t *testing.T) {
	tests := []struct {
		name     string
		size     TabsSize
		expected string
	}{
		{"xs", TabsSizeXS, "tabs-xs"},
		{"small", TabsSizeSmall, "tabs-sm"},
		{"medium", TabsSizeMedium, ""},
		{"large", TabsSizeLarge, "tabs-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.size.String(); got != tt.expected {
				t.Errorf("TabsSize.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewTabs(t *testing.T) {
	tab1 := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tab2 := NewTabItem("tab2", "Tab 2", gomponents.Text("Content 2"))
	tabs := NewTabs(tab1, tab2)

	if tabs == nil {
		t.Fatal("NewTabs() returned nil")
	}

	if len(tabs.tabs) != 2 {
		t.Errorf("Expected 2 tabs, got %d", len(tabs.tabs))
	}

	if tabs.variant != TabsDefault {
		t.Errorf("Expected default variant, got %v", tabs.variant)
	}

	if tabs.size != TabsSizeMedium {
		t.Errorf("Expected medium size, got %v", tabs.size)
	}

	if tabs.color != flyon.Primary {
		t.Errorf("Expected primary color, got %v", tabs.color)
	}
}

func TestTabsComponent_WithID(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)
	modified := tabs.WithID("custom-tabs")

	if modified.id != "custom-tabs" {
		t.Errorf("Expected ID 'custom-tabs', got %s", modified.id)
	}

	// Test immutability
	if tabs.id == "custom-tabs" {
		t.Error("Original tabs component was modified")
	}
}

func TestTabsComponent_WithVariant(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)
	modified := tabs.WithVariant(TabsBordered)

	if modified.variant != TabsBordered {
		t.Errorf("Expected TabsBordered variant, got %v", modified.variant)
	}

	// Test immutability
	if tabs.variant != TabsDefault {
		t.Error("Original tabs component was modified")
	}
}

func TestTabsComponent_WithSize(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)
	modified := tabs.WithSize(TabsSizeLarge)

	if modified.size != TabsSizeLarge {
		t.Errorf("Expected TabsSizeLarge, got %v", modified.size)
	}

	// Test immutability
	if tabs.size != TabsSizeMedium {
		t.Error("Original tabs component was modified")
	}
}

func TestTabsComponent_WithColor(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)
	modified := tabs.WithColor(flyon.Secondary)

	if modified.color != flyon.Secondary {
		t.Errorf("Expected Secondary color, got %v", modified.color)
	}

	// Test immutability
	if tabs.color != flyon.Primary {
		t.Error("Original tabs component was modified")
	}
}

func TestTabsComponent_WithClasses(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)
	modified := tabs.WithClasses("custom-class", "another-class")

	if len(modified.classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(modified.classes))
	}

	if modified.classes[0] != "custom-class" || modified.classes[1] != "another-class" {
		t.Errorf("Expected custom classes, got %v", modified.classes)
	}

	// Test immutability
	if len(tabs.classes) != 0 {
		t.Error("Original tabs component was modified")
	}
}

func TestTabsComponent_With(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab)

	// Test with flyon.Color
	modified1 := tabs.With(flyon.Secondary)
	modified1Tabs := modified1.(*TabsComponent)
	if modified1Tabs.color != flyon.Secondary {
		t.Errorf("Expected Secondary color, got %v", modified1Tabs.color)
	}

	// Test with flyon.Size
	modified2 := tabs.With(flyon.SizeLarge)
	modified2Tabs := modified2.(*TabsComponent)
	if modified2Tabs.size != TabsSizeLarge {
		t.Errorf("Expected TabsSizeLarge, got %v", modified2Tabs.size)
	}

	// Test with TabsVariant
	modified3 := tabs.With(TabsBoxed)
	modified3Tabs := modified3.(*TabsComponent)
	if modified3Tabs.variant != TabsBoxed {
		t.Errorf("Expected TabsBoxed variant, got %v", modified3Tabs.variant)
	}

	// Test with TabsSize
	modified4 := tabs.With(TabsSizeXS)
	modified4Tabs := modified4.(*TabsComponent)
	if modified4Tabs.size != TabsSizeXS {
		t.Errorf("Expected TabsSizeXS, got %v", modified4Tabs.size)
	}

	// Test with multiple modifiers
	modified5 := tabs.With(flyon.Warning, TabsLifted, TabsSizeSmall)
	modified5Tabs := modified5.(*TabsComponent)
	if modified5Tabs.color != flyon.Warning {
		t.Errorf("Expected Warning color, got %v", modified5Tabs.color)
	}
	if modified5Tabs.variant != TabsLifted {
		t.Errorf("Expected TabsLifted variant, got %v", modified5Tabs.variant)
	}
	if modified5Tabs.size != TabsSizeSmall {
		t.Errorf("Expected TabsSizeSmall, got %v", modified5Tabs.size)
	}
}

func TestTabsComponent_Render(t *testing.T) {
	tab1 := NewActiveTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tab2 := NewTabItem("tab2", "Tab 2", gomponents.Text("Content 2"))
	tabs := NewTabs(tab1, tab2).WithID("test-tabs")

	var buf strings.Builder
	err := tabs.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()

	// Check container structure
	if !strings.Contains(html, `id="test-tabs"`) {
		t.Error("Missing tabs ID")
	}
	if !strings.Contains(html, `class="tabs-container"`) {
		t.Error("Missing tabs-container class")
	}
	if !strings.Contains(html, `data-component="tabs"`) {
		t.Error("Missing data-component attribute")
	}

	// Check tab navigation
	if !strings.Contains(html, `class="tabs"`) {
		t.Error("Missing tabs class")
	}
	if !strings.Contains(html, `class="tab tab-active"`) {
		t.Error("Missing active tab class")
	}
	if !strings.Contains(html, `class="tab"`) {
		t.Error("Missing tab class")
	}
	if !strings.Contains(html, `href="#tab1"`) {
		t.Error("Missing tab1 href")
	}
	if !strings.Contains(html, `href="#tab2"`) {
		t.Error("Missing tab2 href")
	}
	if !strings.Contains(html, `data-tab-id="tab1"`) {
		t.Error("Missing tab1 data-tab-id")
	}
	if !strings.Contains(html, `data-tab-id="tab2"`) {
		t.Error("Missing tab2 data-tab-id")
	}

	// Check tab content
	if !strings.Contains(html, `class="tab-content-container"`) {
		t.Error("Missing tab-content-container class")
	}
	if !strings.Contains(html, `id="tab1"`) {
		t.Error("Missing tab1 content ID")
	}
	if !strings.Contains(html, `id="tab2"`) {
		t.Error("Missing tab2 content ID")
	}
	if !strings.Contains(html, `class="tab-content"`) {
		t.Error("Missing tab-content class for active tab")
	}
	if !strings.Contains(html, `class="tab-content hidden"`) {
		t.Error("Missing hidden tab-content class for inactive tab")
	}
	if !strings.Contains(html, `data-tab-panel="tab1"`) {
		t.Error("Missing tab1 data-tab-panel")
	}
	if !strings.Contains(html, `data-tab-panel="tab2"`) {
		t.Error("Missing tab2 data-tab-panel")
	}

	// Check content text
	if !strings.Contains(html, "Tab 1") {
		t.Error("Missing Tab 1 label")
	}
	if !strings.Contains(html, "Tab 2") {
		t.Error("Missing Tab 2 label")
	}
	if !strings.Contains(html, "Content 1") {
		t.Error("Missing Content 1 text")
	}
	if !strings.Contains(html, "Content 2") {
		t.Error("Missing Content 2 text")
	}
}

func TestTabsComponent_RenderWithVariant(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab).WithVariant(TabsBordered)

	var buf strings.Builder
	err := tabs.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, "tabs-bordered") {
		t.Error("Missing tabs-bordered class")
	}
}

func TestTabsComponent_RenderWithSize(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	tabs := NewTabs(tab).WithSize(TabsSizeLarge)

	var buf strings.Builder
	err := tabs.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, "tabs-lg") {
		t.Error("Missing tabs-lg class")
	}
}

func TestNewTabItem(t *testing.T) {
	content := gomponents.Text("Test content")
	tab := NewTabItem("test-id", "Test Label", content)

	if tab.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", tab.ID)
	}
	if tab.Label != "Test Label" {
		t.Errorf("Expected Label 'Test Label', got %s", tab.Label)
	}
	if tab.Content == nil {
		t.Error("Content should not be nil")
	}
	if tab.Active {
		t.Error("Expected tab to be inactive by default")
	}
}

func TestNewActiveTabItem(t *testing.T) {
	content := gomponents.Text("Test content")
	tab := NewActiveTabItem("test-id", "Test Label", content)

	if tab.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", tab.ID)
	}
	if tab.Label != "Test Label" {
		t.Errorf("Expected Label 'Test Label', got %s", tab.Label)
	}
	if tab.Content == nil {
		t.Error("Content should not be nil")
	}
	if !tab.Active {
		t.Error("Expected tab to be active")
	}
}

func TestTabsComponent_Immutability(t *testing.T) {
	tab := NewTabItem("tab1", "Tab 1", gomponents.Text("Content 1"))
	original := NewTabs(tab)
	originalID := original.id
	originalVariant := original.variant
	originalSize := original.size
	originalColor := original.color
	originalClassesLen := len(original.classes)

	// Apply various modifications
	_ = original.WithID("new-id")
	_ = original.WithVariant(TabsBordered)
	_ = original.WithSize(TabsSizeLarge)
	_ = original.WithColor(flyon.Secondary)
	_ = original.WithClasses("new-class")
	_ = original.With(flyon.Warning, TabsLifted)

	// Verify original is unchanged
	if original.id != originalID {
		t.Error("Original ID was modified")
	}
	if original.variant != originalVariant {
		t.Error("Original variant was modified")
	}
	if original.size != originalSize {
		t.Error("Original size was modified")
	}
	if original.color != originalColor {
		t.Error("Original color was modified")
	}
	if len(original.classes) != originalClassesLen {
		t.Error("Original classes were modified")
	}
}