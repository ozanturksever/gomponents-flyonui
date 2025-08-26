//go:build js && wasm

package main

import (
	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
	"honnef.co/go/js/dom/v2"
)

func main() {
	// Wait for DOM to be ready
	bridge.WaitForDOMReady(func() {
		logutil.Log("DOM ready, starting hydration...")
		hydrate()
	})

	// Keep the program running
	select {}
}

func hydrate() {
	// Initialize all FlyonUI components
	bridge.InitializeAllComponents()

	// Setup event listeners for interactive components
	setupDropdownListeners()
	setupModalListeners()
	setupAlertListeners()

	logutil.Log("Hydration complete")
}

func setupDropdownListeners() {
	doc := dom.GetWindow().Document()
	dropdownTriggers := doc.QuerySelectorAll(".dropdown-trigger")

	for _, trigger := range dropdownTriggers {
		trigger.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			logutil.Log("Dropdown triggered")
			
			// Find the associated dropdown menu
			dropdown := trigger.ParentElement().QuerySelector(".dropdown-content")
			if dropdown != nil {
				if dropdown.Class().Contains("hidden") {
					dropdown.Class().Remove("hidden")
					logutil.Log("Dropdown opened")
				} else {
					dropdown.Class().Add("hidden")
					logutil.Log("Dropdown closed")
				}
			}
		})
	}
}

func setupModalListeners() {
	doc := dom.GetWindow().Document()
	modalTriggers := doc.QuerySelectorAll(".modal-trigger")

	for _, trigger := range modalTriggers {
		trigger.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			logutil.Log("Modal triggered")
			
			// Find the target modal
			modalId := trigger.GetAttribute("data-modal-target")
			if modalId != "" {
				modal := doc.GetElementByID(modalId)
				if modal != nil {
					modal.Class().Remove("hidden")
					logutil.Log("Modal opened:", modalId)
				}
			}
		})
	}

	// Setup modal close listeners
	modalCloses := doc.QuerySelectorAll(".modal-close")
	for _, closeBtn := range modalCloses {
		closeBtn.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			logutil.Log("Modal close triggered")
			
			// Find the parent modal
			modal := closeBtn.Closest(".modal")
			if modal != nil {
				modal.Class().Add("hidden")
				logutil.Log("Modal closed")
			}
		})
	}
}

func setupAlertListeners() {
	doc := dom.GetWindow().Document()
	alertCloses := doc.QuerySelectorAll(".alert-close")

	for _, closeBtn := range alertCloses {
		closeBtn.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			logutil.Log("Alert close triggered")
			
			// Find the parent alert
			alert := closeBtn.Closest(".alert")
			if alert != nil {
				// Use the underlying JS value to set style
				alert.Underlying().Get("style").Set("display", "none")
				logutil.Log("Alert closed")
			}
		})
	}
}