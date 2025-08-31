//go:build !js && !wasm

package main

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/ozanturksever/gomponents-flyonui/internal/devserver"
	"github.com/ozanturksever/gomponents-flyonui/internal/testhelpers"
)

func TestGomponentsShowcase_PageLoad(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	var title string
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		chromedp.Title(&title),
	)
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}

	expected := "Gomponents FlyonUI Showcase"
	if title != expected {
		t.Errorf("Expected title %q, got %q", expected, title)
	}
}

func TestGomponentsShowcase_WASMInitialization(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	var wasmStatus string
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		chromedp.Text("#wasm-status-text", &wasmStatus),
	)
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}

	if wasmStatus == "" {
		t.Error("WASM status should not be empty")
	}
	t.Logf("WASM Status: %s", wasmStatus)
}

func TestGomponentsShowcase_ComponentRendering(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test that all major component sections are rendered
	componentSections := []string{
		"#button-showcase",
		"#alert-showcase", 
		"#card-showcase",
		"#modal-showcase",
		"#dropdown-showcase",
		"#accordion-showcase",
		"#progress-showcase",
		"#layout-showcase",
		"#form-showcase",
	}

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
	)
	if err != nil {
		t.Fatalf("Failed to navigate and initialize: %v", err)
	}

	for _, selector := range componentSections {
		var exists bool
		err := chromedp.Run(chromedpCtx.Ctx,
			chromedp.WaitVisible(selector, chromedp.ByID),
			chromedp.Evaluate(`document.querySelector('`+selector+`') !== null`, &exists),
		)
		if err != nil {
			t.Errorf("Component section %s not found or not visible: %v", selector, err)
			continue
		}
		if !exists {
			t.Errorf("Component section %s does not exist", selector)
		}
	}
}

func TestGomponentsShowcase_ButtonInteraction(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		// Click on a primary button
		testhelpers.Actions.ClickAndWait("#button-showcase .btn-primary", 500*time.Millisecond),
	)
	if err != nil {
		t.Fatalf("Button interaction test failed: %v", err)
	}
}

func TestGomponentsShowcase_AlertClosing(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		// Wait for alert to be visible
		chromedp.WaitVisible("#alert-showcase .alert", chromedp.ByQuery),
		// Click close button on first alert
		testhelpers.Actions.ClickAndWait("#alert-showcase .alert .alert-close", 500*time.Millisecond),
	)
	if err != nil {
		t.Fatalf("Alert closing test failed: %v", err)
	}
}

func TestGomponentsShowcase_DropdownInteraction(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		// Wait for dropdown to be visible
		chromedp.WaitVisible("#dropdown-showcase .dropdown", chromedp.ByQuery),
		// Click dropdown trigger
		testhelpers.Actions.ClickAndWait("#dropdown-showcase .dropdown .dropdown-trigger", 500*time.Millisecond),
	)
	if err != nil {
		t.Fatalf("Dropdown interaction test failed: %v", err)
	}
}

func TestGomponentsShowcase_ModalInteraction(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.ExtendedTimeoutConfig())
	defer chromedpCtx.Cancel()

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		// Wait for modal trigger to be visible
		chromedp.WaitVisible("#modal-showcase .modal-trigger", chromedp.ByQuery),
		// Click modal trigger
		testhelpers.Actions.ClickAndWait("#modal-showcase .modal-trigger", 1*time.Second),
	)
	if err != nil {
		t.Fatalf("Modal interaction test failed: %v", err)
	}
}

func TestGomponentsShowcase_AccordionInteraction(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		// Wait for accordion to be visible
		chromedp.WaitVisible("#accordion-showcase .accordion", chromedp.ByQuery),
		// Click first accordion item
		testhelpers.Actions.ClickAndWait("#accordion-showcase .accordion .accordion-item:first-child .accordion-header", 500*time.Millisecond),
	)
	if err != nil {
		t.Fatalf("Accordion interaction test failed: %v", err)
	}
}

func TestGomponentsShowcase_ResponsiveLayout(t *testing.T) {
	server := devserver.NewServer("gomponents_showcase", "localhost:0")
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start dev server: %v", err)
	}
	defer server.Stop()

	chromedpCtx := testhelpers.MustNewChromedpContext(testhelpers.DefaultConfig())
	defer chromedpCtx.Cancel()

	// Test desktop layout
	err := chromedp.Run(chromedpCtx.Ctx,
		testhelpers.Actions.NavigateAndWaitForLoad(server.URL(), "body"),
		testhelpers.Actions.WaitForWASMInit(),
		chromedp.EmulateViewport(1200, 800),
		chromedp.WaitVisible("#layout-showcase", chromedp.ByID),
	)
	if err != nil {
		t.Fatalf("Desktop layout test failed: %v", err)
	}

	// Test mobile layout
	err = chromedp.Run(chromedpCtx.Ctx,
		chromedp.EmulateViewport(375, 667),
		chromedp.Sleep(500*time.Millisecond), // Allow layout to adjust
		chromedp.WaitVisible("#layout-showcase", chromedp.ByID),
	)
	if err != nil {
		t.Fatalf("Mobile layout test failed: %v", err)
	}
}