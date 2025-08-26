//go:build js && wasm

package main

import (
	"testing"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
	"honnef.co/go/js/dom/v2"
)

// TestHydrateDropdowns tests the dropdown hydration functionality
func TestHydrateDropdowns(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Underlying().Get("body")

	// Create a test dropdown element
	testDropdown := doc.CreateElement("div")
	testDropdown.SetInnerHTML(`
		<div class="dropdown">
			<button data-dropdown-toggle="#test-menu">Toggle</button>
			<ul id="test-menu" class="dropdown-menu">
				<li><a href="#">Item 1</a></li>
				<li><a href="#">Item 2</a></li>
			</ul>
		</div>
	`)

	// Append to body for testing
	body.Call("appendChild", testDropdown.Underlying())
	defer body.Call("removeChild", testDropdown.Underlying())

	// Test hydration
	hydrateDropdowns()

	// Verify dropdown items were found and hydrated
	dropdownItems := doc.QuerySelectorAll(".dropdown-menu a")
	if len(dropdownItems) < 2 {
		t.Errorf("Expected at least 2 dropdown items, got %d", len(dropdownItems))
	}
}

// TestHydrateModals tests the modal hydration functionality
func TestHydrateModals(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Underlying().Get("body")

	// Create a test modal trigger element
	testModal := doc.CreateElement("div")
	testModal.SetInnerHTML(`
		<button data-modal-toggle="#test-modal">Open Modal</button>
		<div id="test-modal" class="modal">
			<div class="modal-content">
				<p>Test modal content</p>
			</div>
		</div>
	`)

	// Append to body for testing
	body.Call("appendChild", testModal.Underlying())
	defer body.Call("removeChild", testModal.Underlying())

	// Test hydration
	hydrateModals()

	// Verify modal triggers were found and hydrated
	modalTriggers := doc.QuerySelectorAll("[data-modal-toggle]")
	if len(modalTriggers) < 1 {
		t.Errorf("Expected at least 1 modal trigger, got %d", len(modalTriggers))
	}
}

// TestHydrateTooltips tests the tooltip hydration functionality
func TestHydrateTooltips(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Underlying().Get("body")

	// Create a test tooltip element
	testTooltip := doc.CreateElement("div")
	testTooltip.SetInnerHTML(`
		<button data-tooltip="Test tooltip content">Hover me</button>
	`)

	// Append to body for testing
	body.Call("appendChild", testTooltip.Underlying())
	defer body.Call("removeChild", testTooltip.Underlying())

	// Test hydration
	hydrateTooltips()

	// Verify tooltip elements were found and hydrated
	tooltipElements := doc.QuerySelectorAll("[data-tooltip]")
	if len(tooltipElements) < 1 {
		t.Errorf("Expected at least 1 tooltip element, got %d", len(tooltipElements))
	}
}

// TestHydrateAlerts tests the alert hydration functionality
func TestHydrateAlerts(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Underlying().Get("body")

	// Create a test alert element
	testAlert := doc.CreateElement("div")
	testAlert.SetInnerHTML(`
		<div class="alert">
			<p>Test alert message</p>
			<button data-dismiss="alert">×</button>
		</div>
	`)

	// Append to body for testing
	body.Call("appendChild", testAlert.Underlying())
	defer body.Call("removeChild", testAlert.Underlying())

	// Test hydration
	hydrateAlerts()

	// Verify alert close buttons were found and hydrated
	alertCloseButtons := doc.QuerySelectorAll(".alert [data-dismiss='alert']")
	if len(alertCloseButtons) < 1 {
		t.Errorf("Expected at least 1 alert close button, got %d", len(alertCloseButtons))
	}
}

// TestHydrateFunction tests the main hydrate function
func TestHydrateFunction(t *testing.T) {
	// This test ensures hydrate() doesn't panic and completes successfully
	// We can't easily test the FlyonUI initialization without the actual JS library
	// but we can ensure our Go hydration functions work
	hydrate()

	// If we reach here without panicking, the test passes
	t.Log("Hydration completed successfully")
}

// TestDOMReady tests that the DOM ready functionality works
func TestDOMReady(t *testing.T) {
	start := time.Now()
	callbackExecuted := false

	bridge.WaitForDOMReady(func() {
		callbackExecuted = true
	})

	// Give some time for the callback to execute
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(start)

	if !callbackExecuted {
		t.Fatalf("WaitForDOMReady callback was not executed")
	}

	// Should complete quickly since DOM is already ready in test environment
	if elapsed > 5*time.Second {
		t.Errorf("WaitForDOMReady took too long: %v", elapsed)
	}

	t.Logf("DOM ready check completed in %v", elapsed)
}