//go:build js && wasm

package main

import (
	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
	"honnef.co/go/js/dom/v2"
)

// main is the entry point for the WASM binary.
// It initializes the FlyonUI components and sets up hydration.
func main() {
	logutil.Log("WASM runtime initializing...")

	// Wait for DOM to be ready before hydrating
	bridge.WaitForDOMReady(func() {
		// Hydrate all components on the page
		hydrate()
	})

	logutil.Log("WASM runtime initialized successfully")

	// Keep the WASM binary alive - this is crucial for maintaining event listeners
	// and allowing the Go runtime to continue processing events
	select {}
}

// hydrate initializes all FlyonUI components and sets up client-side interactivity
func hydrate() {
	logutil.Log("Starting component hydration...")

	// Initialize all FlyonUI components found on the page
	// This will scan for data-* attributes and initialize the corresponding components
	if err := bridge.InitializeAllComponents(); err != nil {
		logutil.Logf("Error initializing FlyonUI components: %v", err)
		return
	}

	// Hydrate specific component types with custom Go logic
	hydrateDropdowns()
	hydrateModals()
	hydrateTooltips()
	hydrateAlerts()

	logutil.Log("Component hydration completed")
}

// hydrateDropdowns adds custom Go-based logic to dropdown components
func hydrateDropdowns() {
	doc := dom.GetWindow().Document()
	dropdownItems := doc.QuerySelectorAll(".dropdown-menu a")

	for _, item := range dropdownItems {
		item.AddEventListener("click", false, func(event dom.Event) {
			target := event.Target().(dom.Element)
			logutil.Logf("Dropdown item clicked: %s", target.TextContent())
		})
	}

	logutil.Logf("Hydrated %d dropdown items", len(dropdownItems))
}

// hydrateModals adds custom Go-based logic to modal components
func hydrateModals() {
	doc := dom.GetWindow().Document()
	modalTriggers := doc.QuerySelectorAll("[data-modal-toggle]")

	for _, trigger := range modalTriggers {
		trigger.AddEventListener("click", false, func(event dom.Event) {
			logutil.Log("Modal trigger clicked")
			// Additional custom logic can be added here
		})
	}

	logutil.Logf("Hydrated %d modal triggers", len(modalTriggers))
}

// hydrateTooltips adds custom Go-based logic to tooltip components
func hydrateTooltips() {
	doc := dom.GetWindow().Document()
	tooltipElements := doc.QuerySelectorAll("[data-tooltip]")

	for _, element := range tooltipElements {
		element.AddEventListener("mouseenter", false, func(event dom.Event) {
			logutil.Log("Tooltip hover started")
		})
		element.AddEventListener("mouseleave", false, func(event dom.Event) {
			logutil.Log("Tooltip hover ended")
		})
	}

	logutil.Logf("Hydrated %d tooltip elements", len(tooltipElements))
}

// hydrateAlerts adds custom Go-based logic to alert components
func hydrateAlerts() {
	doc := dom.GetWindow().Document()
	alertCloseButtons := doc.QuerySelectorAll(".alert [data-dismiss='alert']")

	for _, button := range alertCloseButtons {
		button.AddEventListener("click", false, func(event dom.Event) {
			logutil.Log("Alert close button clicked")
			// The actual close functionality is handled by FlyonUI JS
			// This is just for logging/analytics
		})
	}

	logutil.Logf("Hydrated %d alert close buttons", len(alertCloseButtons))
}