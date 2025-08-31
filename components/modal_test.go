package components

import (
	"strings"
	"testing"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"maragu.dev/gomponents"
)

func TestNewModal(t *testing.T) {
	content := gomponents.Text("Modal content")
	modal := NewModal("Test Modal", content)
	
	if modal == nil {
		t.Fatal("NewModal should not return nil")
	}
	
	if modal.title != "Test Modal" {
		t.Errorf("Expected title 'Test Modal', got '%s'", modal.title)
	}
	
	if len(modal.content) != 1 {
		t.Errorf("Expected 1 content item, got %d", len(modal.content))
	}
	
	if len(modal.classes) != 0 {
		t.Errorf("Expected no default custom classes, got %v", modal.classes)
	}
	
	if modal.size != ModalSizeDefault {
		t.Errorf("Expected default size ModalSizeDefault, got %v", modal.size)
	}
	
	if !modal.closable {
		t.Error("Expected modal to be closable by default")
	}
	
	if !modal.backdrop {
		t.Error("Expected modal to have backdrop by default")
	}
	
	if modal.open {
		t.Error("Expected modal to be closed by default")
	}
}

func TestModalSize_String(t *testing.T) {
	tests := []struct {
		name     string
		size     ModalSize
		expected string
	}{
		{"modal-dialog-sm", ModalSizeSmall, "modal-dialog-sm"},
		{"modal-dialog-md", ModalSizeMedium, "modal-dialog-md"},
		{"modal-dialog-lg", ModalSizeLarge, "modal-dialog-lg"},
		{"modal-dialog-xl", ModalSizeExtraLarge, "modal-dialog-xl"},
		{"modal-dialog-full", ModalSizeFullWidth, "modal-dialog-full"},
		{"default", ModalSize(999), ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.size.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestModalComponent_WithID(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	modified := original.WithID("custom-modal")
	
	// Original should be unchanged
	if original.id != "" {
		t.Error("Original modal ID should remain empty")
	}
	
	// Modified should have new ID
	if modified.id != "custom-modal" {
		t.Errorf("Expected ID 'custom-modal', got '%s'", modified.id)
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithID should return a new instance")
	}
}

func TestModalComponent_WithSize(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	modified := original.WithSize(ModalSizeLarge)
	
	// Original should be unchanged
	if original.size != ModalSizeDefault {
		t.Error("Original modal size should remain ModalSizeDefault")
	}
	
	// Modified should have new size
	if modified.size != ModalSizeLarge {
		t.Errorf("Expected size ModalSizeLarge, got %v", modified.size)
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithSize should return a new instance")
	}
}

func TestModalComponent_WithClosable(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	modified := original.WithClosable(false)
	
	// Original should be unchanged
	if !original.closable {
		t.Error("Original modal should remain closable")
	}
	
	// Modified should have new closable state
	if modified.closable {
		t.Error("Modified modal should not be closable")
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithClosable should return a new instance")
	}
}

func TestModalComponent_WithBackdrop(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	modified := original.WithBackdrop(false)
	
	// Original should be unchanged
	if !original.backdrop {
		t.Error("Original modal should have backdrop")
	}
	
	// Modified should have new backdrop state
	if modified.backdrop {
		t.Error("Modified modal should not have backdrop")
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithBackdrop should return a new instance")
	}
}

func TestModalComponent_WithOpen(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	modified := original.WithOpen(true)
	
	// Original should be unchanged
	if original.open {
		t.Error("Original modal should remain closed")
	}
	
	// Modified should have new open state
	if !modified.open {
		t.Error("Modified modal should be open")
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithOpen should return a new instance")
	}
}

func TestModalComponent_WithActions(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	action1 := gomponents.Text("OK")
	action2 := gomponents.Text("Cancel")
	modified := original.WithActions(action1, action2)
	
	// Original should be unchanged
	if len(original.actions) != 0 {
		t.Error("Original modal should have no actions")
	}
	
	// Modified should have new actions
	if len(modified.actions) != 2 {
		t.Errorf("Expected 2 actions, got %d", len(modified.actions))
	}
	
	// Should be different instances
	if original == modified {
		t.Error("WithActions should return a new instance")
	}
}

func TestModalComponent_With(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	
	// Test with flyon.Size
	modified1 := original.With(flyon.SizeLarge)
	modified1Modal, ok := modified1.(*ModalComponent)
	if !ok {
		t.Fatal("With should return *ModalComponent")
	}
	if modified1Modal.size != ModalSizeLarge {
		t.Errorf("Expected size ModalSizeLarge, got %v", modified1Modal.size)
	}
	
	// Test with ModalSize
	modified2 := original.With(ModalSizeExtraLarge)
	modified2Modal, ok := modified2.(*ModalComponent)
	if !ok {
		t.Fatal("With should return *ModalComponent")
	}
	if modified2Modal.size != ModalSizeExtraLarge {
		t.Errorf("Expected size ModalSizeExtraLarge, got %v", modified2Modal.size)
	}
	
	// Test with custom class
	modified3 := original.With("custom-class")
	modified3Modal, ok := modified3.(*ModalComponent)
	if !ok {
		t.Fatal("With should return *ModalComponent")
	}
	found := false
	for _, class := range modified3Modal.classes {
		if class == "custom-class" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Custom class should be added")
	}
	
	// Original should be unchanged
	if original.size != ModalSizeDefault {
		t.Error("Original modal size should remain unchanged")
	}
	// The classes slice holds only custom classes; base classes are added at render-time
	if len(original.classes) != 0 {
		t.Error("Original modal classes should remain unchanged (no custom classes by default)")
	}
}

func TestModalComponent_Render(t *testing.T) {
	content := gomponents.Text("Modal content")
	modal := NewModal("Test Modal", content)
	
	var buf strings.Builder
	err := modal.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	// Check for modal structure per updated docs
	if !strings.Contains(html, `class="overlay modal overlay-open:opacity-100 overlay-open:duration-300`) {
		t.Error("Modal should have 'overlay modal overlay-open:opacity-100 overlay-open:duration-300' classes")
	}
	
	if !strings.Contains(html, `role="dialog"`) {
		t.Error("Modal should have dialog role")
	}
	
	if !strings.Contains(html, `tabindex="-1"`) {
		t.Error("Modal should have tabindex -1 for accessibility")
	}
	
	// Check for wrappers
	if !strings.Contains(html, "modal-dialog") {
		t.Error("Modal should include modal-dialog wrapper")
	}
	if !strings.Contains(html, "modal-content") {
		t.Error("Modal should include modal-content wrapper")
	}
	
	// Check for title and content
	if !strings.Contains(html, "Test Modal") {
		t.Error("Modal should contain the title")
	}
	if !strings.Contains(html, "Modal content") {
		t.Error("Modal should contain the content")
	}
	
	// Check for close button (default closable)
	if !strings.Contains(html, `data-overlay="#`) {
		t.Error("Close button should use data-overlay to target modal id")
	}
}

func TestModalComponent_RenderWithSize(t *testing.T) {
	content := gomponents.Text("Modal content")
	modal := NewModal("Test Modal", content).WithSize(ModalSizeLarge)
	
	var buf strings.Builder
	err := modal.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	// Size classes should be applied to modal-dialog element
	if !strings.Contains(html, `class="modal-dialog modal-dialog-lg"`) {
		t.Error("Modal dialog should include modal-dialog-lg size class")
	}

	// Size classes should NOT be on the main modal container
	if strings.Contains(html, `class="overlay modal overlay-open:opacity-100 overlay-open:duration-300 hidden modal-dialog-lg"`) {
		t.Error("Modal container should not include size classes")
	}
}

func TestModalComponent_RenderNotClosable(t *testing.T) {
	content := gomponents.Text("Modal content")
	modal := NewModal("Test Modal", content).WithClosable(false)
	
	var buf strings.Builder
	err := modal.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	// Check that close button is not present (data-overlay close should be absent in header)
	if strings.Contains(html, `data-overlay="#`) {
		t.Error("Modal should not have close button when not closable")
	}

	// Check that keyboard control is disabled when not closable
	if !strings.Contains(html, `data-overlay-keyboard="false"`) {
		t.Error("Modal should have data-overlay-keyboard=false when not closable")
	}
}

func TestModalComponent_RenderWithKeyboard(t *testing.T) {
	content := gomponents.Text("Modal content")
	modal := NewModal("Test Modal", content).WithKeyboard(false)

	var buf strings.Builder
	err := modal.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	html := buf.String()

	// Check that keyboard control attribute is present when disabled
	if !strings.Contains(html, `data-overlay-keyboard="false"`) {
		t.Error("Modal should have data-overlay-keyboard=false when keyboard is disabled")
	}
}

func TestModalComponent_RenderWithActions(t *testing.T) {
	content := gomponents.Text("Modal content")
	action1 := ModalAction("OK", flyon.Primary)
	action2 := ModalCloseAction("Cancel", flyon.Secondary)
	modal := NewModal("Test Modal", content).WithActions(action1, action2)
	
	var buf strings.Builder
	err := modal.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	// Check for modal footer
	if !strings.Contains(html, "modal-footer") {
		t.Error("Modal should have footer when actions are provided")
	}
	
	// Check for action buttons
	if !strings.Contains(html, "OK") {
		t.Error("Modal should contain OK button")
	}
	
	if !strings.Contains(html, "Cancel") {
		t.Error("Modal should contain Cancel button")
	}
	
	if !strings.Contains(html, "btn-primary") {
		t.Error("Modal should have primary button")
	}
	
	if !strings.Contains(html, "btn-secondary") {
		t.Error("Modal should have secondary button")
	}
}

func TestModalAction(t *testing.T) {
	action := ModalAction("Test Action", flyon.Primary)
	
	var buf strings.Builder
	err := action.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "Test Action") {
		t.Error("Action should contain the text")
	}
	
	if !strings.Contains(html, "btn-primary") {
		t.Error("Action should have primary button class")
	}
	
	if !strings.Contains(html, `type="button"`) {
		t.Error("Action should be a button")
	}
}

func TestModalCloseAction(t *testing.T) {
	action := ModalCloseAction("Close", flyon.Secondary)
	
	var buf strings.Builder
	err := action.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	html := buf.String()
	
	if !strings.Contains(html, "Close") {
		t.Error("Close action should contain the text")
	}
	
	if !strings.Contains(html, "btn-secondary") {
		t.Error("Close action should have secondary button class")
	}
	
	if !strings.Contains(html, `class="btn btn-secondary`) || !strings.Contains(html, "modal-close") {
		t.Error("Close action should have secondary btn class and modal-close helper class")
	}
}

func TestModalComponent_Immutability(t *testing.T) {
	original := NewModal("Test", gomponents.Text("content"))
	
	// Test that modifications create new instances
	modified1 := original.WithID("test-id")
	modified2 := original.WithSize(ModalSizeLarge)
	modified3 := original.WithClosable(false)
	modified4 := original.WithBackdrop(false)
	modified5 := original.WithOpen(true)
	modified6 := original.With(flyon.SizeLarge)
	
	// Ensure original is unchanged
	if original.id != "" {
		t.Error("Original ID should remain empty")
	}
	if original.size != ModalSizeDefault {
		t.Error("Original size should remain ModalSizeDefault")
	}
	if !original.closable {
		t.Error("Original closable should remain true")
	}
	if !original.backdrop {
		t.Error("Original backdrop should remain true")
	}
	if original.open {
		t.Error("Original open should remain false")
	}
	if len(original.classes) != 0 {
		t.Error("Original classes should remain empty by default")
	}
	
	// Ensure modifications are applied to new instances
	if modified1.id != "test-id" {
		t.Error("Modified1 should have new ID")
	}
	if modified2.size != ModalSizeLarge {
		t.Error("Modified2 should have new size")
	}
	if modified3.closable {
		t.Error("Modified3 should not be closable")
	}
	if modified4.backdrop {
		t.Error("Modified4 should not have backdrop")
	}
	if !modified5.open {
		t.Error("Modified5 should be open")
	}
	
	// Test With method returns flyon.Component
	modified6Modal, ok := modified6.(*ModalComponent)
	if !ok {
		t.Error("With method should return *ModalComponent")
	}
	if modified6Modal.size != ModalSizeLarge {
		t.Error("Modified6 should have large size")
	}
}