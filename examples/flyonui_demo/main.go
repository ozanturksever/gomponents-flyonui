//go:build js && wasm

package main

import (
	"syscall/js"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/internal/bridge"
	"github.com/ozanturksever/gomponents-flyonui/internal/vite"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
	"honnef.co/go/js/dom/v2"
)

func main() {
	// Wait for DOM to be ready
	bridge.WaitForDOMReady(func() {
		logutil.Log("DOM ready, starting hydration...")

		// Initialize Vite asset resolver
		assetResolver := vite.NewAssetResolver("", true) // Empty baseURL for development, development=true
		if err := assetResolver.LoadManifest(); err != nil {
			logutil.Logf("Warning: Could not load Vite manifest: %v", err)
		}

		hydrate(assetResolver)
	})

	// Keep the program running
	select {}
}

func hydrate(assetResolver *vite.AssetResolver) {
	// Log asset information
	cssURL := assetResolver.GetAssetURL("css/main.css")
	jsURL := assetResolver.GetAssetURL("js/main.js")
	logutil.Logf("Using CSS asset: %s", cssURL)
	logutil.Logf("Using JS asset: %s", jsURL)

	// Wait for JavaScript to be ready before initializing FlyonUI
	waitForJSReady(func() {
		// Initialize FlyonUI JavaScript components
		initializeFlyonUIComponents()

		// Initialize all FlyonUI components via bridge
		bridge.InitializeAllComponents()

		// Setup event listeners for interactive components
		setupDropdownListeners()
		setupModalListeners()
		setupAlertListeners()

		// Complete hydration
		completeHydration()
	})
}

// waitForJSReady waits for the JavaScript initialization to complete
func waitForJSReady(callback func()) {
	// Check if JS is ready
	if js.Global().Get("flyonUIManager").Truthy() {
		callback()
		return
	}

	// Listen for jsReady event
	js.Global().Call("addEventListener", "jsReady", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		logutil.Log("JavaScript ready event received")
		callback()
		return nil
	}))
}

// initializeFlyonUIComponents initializes FlyonUI components via JavaScript
func initializeFlyonUIComponents() {
	logutil.Log("Initializing FlyonUI components...")

	// Check if HSStaticMethods is available (FlyonUI's initialization method)
	hsStaticMethods := js.Global().Get("HSStaticMethods")
	if !hsStaticMethods.Truthy() {
		logutil.Log("Warning: HSStaticMethods not found, FlyonUI may not be loaded")
		return
	}

	// Check if FlyonUI manager is available
	flyonUIManager := js.Global().Get("flyonUIManager")
	if !flyonUIManager.Truthy() {
		logutil.Log("Warning: FlyonUI manager not found")
		return
	}

	// Check if FlyonUI is already initialized
	if flyonUIManager.Get("initialized").Bool() {
		logutil.Log("FlyonUI already initialized")
		return
	}

	// Initialize FlyonUI via the manager
	flyonUIManager.Call("init")
	logutil.Log("FlyonUI components initialized via JavaScript")
}

// reinitializeFlyonUIComponents reinitializes FlyonUI components for dynamic content
func reinitializeFlyonUIComponents() {
	logutil.Log("Reinitializing FlyonUI components...")

	// Use the registered Go callback
	goWASMUtils := js.Global().Get("GoWASMUtils")
	if goWASMUtils.Truthy() {
		goWASMUtils.Call("callGoFunction", "reinitializeFlyonUI")
		logutil.Log("FlyonUI components reinitialized")
	} else {
		logutil.Log("Warning: GoWASMUtils not available for reinitializing FlyonUI")
	}
}

// completeHydration finalizes the hydration process
func completeHydration() {

	// Update WASM status to indicate readiness
	doc := dom.GetWindow().Document()
	wasmStatus := doc.GetElementByID("wasm-status")
	if wasmStatus != nil {
		wasmStatus.SetTextContent("Ready")
		logutil.Log("WASM status updated to Ready")
		// Emit a simple console log to satisfy test expectations

		// Dispatch a custom 'wasmReady' event in a way compatible with the browser and JS listeners
		// Note: honnef.co/go/js/dom/v2 may not provide a direct CustomEvent constructor with init across versions,
		// so we construct it via syscall/js and dispatch on window (assets/js/main.js listens on window).
		time.AfterFunc(10*time.Millisecond, func() {
			evt := js.Global().Get("CustomEvent").New("wasmReady")
			js.Global().Call("dispatchEvent", evt)
			logutil.Log("WASM ready event dispatched", evt)
		})
	}

	logutil.Log("Hydration complete with Vite assets")
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
