//go:build !js && !wasm

package main

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/ozanturksever/gomponents-flyonui/internal/devserver"
	"github.com/ozanturksever/gomponents-flyonui/internal/testhelpers"
)

func TestFlyonUIDemo_PageLoad(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test page load and WASM initialization
	var wasmStatus string
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		chromedp.Text("#wasm-status", &wasmStatus, chromedp.ByID),
	)
	if err != nil {
		t.Fatalf("Page load test failed: %v", err)
	}

	if wasmStatus != "Ready" {
		t.Errorf("Expected WASM status to be 'Ready', got '%s'", wasmStatus)
	}
}

func TestFlyonUIDemo_DropdownInteraction(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test dropdown functionality
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Test primary dropdown
		chromedp.WaitVisible(".dropdown-trigger", chromedp.ByQuery),
		testhelpers.Actions.ClickAndWait(".dropdown-trigger", 500*time.Millisecond),
		
		// Verify dropdown content is visible
		chromedp.WaitVisible(".dropdown-content", chromedp.ByQuery),
		
		// Click dropdown option
		testhelpers.Actions.ClickAndWait(".dropdown-content a:first-child", 500*time.Millisecond),
	)
	if err != nil {
		t.Fatalf("Dropdown interaction test failed: %v", err)
	}
}

func TestFlyonUIDemo_ModalInteraction(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test modal functionality
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Open demo modal
		chromedp.WaitVisible("[data-modal-target='demo-modal']", chromedp.ByQuery),
		testhelpers.Actions.ClickAndWait("[data-modal-target='demo-modal']", 500*time.Millisecond),
		
		// Verify modal is visible
		chromedp.WaitVisible("#demo-modal", chromedp.ByID),
		
		// Close modal
		testhelpers.Actions.ClickAndWait(".modal-close", 500*time.Millisecond),
		
		// Verify modal is hidden
		chromedp.WaitNotPresent("#demo-modal:not(.hidden)", chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("Modal interaction test failed: %v", err)
	}
}

func TestFlyonUIDemo_ConfirmModalInteraction(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test confirm modal functionality
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Open confirm modal
		chromedp.WaitVisible("[data-modal-target='confirm-modal']", chromedp.ByQuery),
		testhelpers.Actions.ClickAndWait("[data-modal-target='confirm-modal']", 500*time.Millisecond),
		
		// Verify modal is visible
		chromedp.WaitVisible("#confirm-modal", chromedp.ByID),
		
		// Test cancel button
		testhelpers.Actions.ClickAndWait(".btn-error", 500*time.Millisecond),
		
		// Verify modal is hidden
		chromedp.WaitNotPresent("#confirm-modal:not(.hidden)", chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("Confirm modal interaction test failed: %v", err)
	}
}

func TestFlyonUIDemo_AlertInteraction(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test alert close functionality
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Verify alerts are visible
		chromedp.WaitVisible("#info-alert", chromedp.ByID),
		chromedp.WaitVisible("#success-alert", chromedp.ByID),
		
		// Close info alert
		testhelpers.Actions.ClickAndWait("#info-alert .alert-close", 500*time.Millisecond),
		
		// Verify info alert is hidden
		chromedp.WaitNotVisible("#info-alert", chromedp.ByID),
		
		// Close success alert
		testhelpers.Actions.ClickAndWait("#success-alert .alert-close", 500*time.Millisecond),
		
		// Verify success alert is hidden
		chromedp.WaitNotVisible("#success-alert", chromedp.ByID),
	)
	if err != nil {
		t.Fatalf("Alert interaction test failed: %v", err)
	}
}

func TestFlyonUIDemo_ComponentShowcase(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test component showcase rendering
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Verify badges are rendered
		chromedp.WaitVisible(".badge-primary", chromedp.ByQuery),
		chromedp.WaitVisible(".badge-secondary", chromedp.ByQuery),
		chromedp.WaitVisible(".badge-accent", chromedp.ByQuery),
		chromedp.WaitVisible(".badge-success", chromedp.ByQuery),
		
		// Verify progress bars are rendered
		chromedp.WaitVisible(".progress-primary", chromedp.ByQuery),
		chromedp.WaitVisible(".progress-secondary", chromedp.ByQuery),
		chromedp.WaitVisible(".progress-accent", chromedp.ByQuery),
		
		// Verify avatars are rendered
		chromedp.WaitVisible(".avatar", chromedp.ByQuery),
		
		// Verify stats are rendered
		chromedp.WaitVisible(".stats", chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("Component showcase test failed: %v", err)
	}
}

func TestFlyonUIDemo_ResponsiveLayout(t *testing.T) {
	// Start dev server
	server := devserver.NewServer("flyonui_demo", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	// Setup chromedp context with visible browser for responsive testing
	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test responsive layout at different viewport sizes
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit("#wasm-status", 3*time.Second),
		
		// Test mobile viewport
		chromedp.EmulateViewport(375, 667),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.WaitVisible(".container", chromedp.ByQuery),
		
		// Test tablet viewport
		chromedp.EmulateViewport(768, 1024),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.WaitVisible(".grid", chromedp.ByQuery),
		
		// Test desktop viewport
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.WaitVisible(".lg\\:grid-cols-3", chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("Responsive layout test failed: %v", err)
	}
}