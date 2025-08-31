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

// setupModalListeners sets up event listeners for modal functionality using FlyonUI's HSOverlay API
func setupModalListeners() {
	doc := dom.GetWindow().Document()

	// Initialize HSOverlay for all modals
	if hsOverlay := js.Global().Get("HSOverlay"); !hsOverlay.IsUndefined() {
		// Auto-initialize all overlays
		hsOverlay.Call("autoInit")
		logutil.Log("HSOverlay auto-initialized")
	} else {
		logutil.Log("HSOverlay not found, falling back to manual implementation")
		// Fallback to manual implementation
		setupManualModalListeners(doc)
		return
	}

	//// Handle modal trigger buttons with HSOverlay API
	//modalTriggers := doc.QuerySelectorAll(".modal-trigger")
	//for _, trigger := range modalTriggers {
	//	trigger.AddEventListener("click", false, func(event dom.Event) {
	//		event.PreventDefault()
	//
	//		// Get the target modal ID from data attribute
	//		targetID := trigger.GetAttribute("data-modal-target")
	//		if targetID == "" {
	//			return
	//		}
	//
	//		// Use HSOverlay API to open the modal
	//		modal := doc.GetElementByID(targetID)
	//		if modal != nil {
	//			if hsOverlay := js.Global().Get("HSOverlay"); !hsOverlay.IsUndefined() {
	//				// Create HSOverlay instance and open
	//				instance := hsOverlay.Call("getInstance", modal.Underlying())
	//				if instance.IsUndefined() {
	//					// Create new instance if it doesn't exist
	//					instance = js.Global().Get("HSOverlay").New(modal.Underlying())
	//				}
	//				instance.Call("open")
	//				logutil.Logf("Opened modal using HSOverlay: %s", targetID)
	//			} else {
	//				// Fallback
	//				modal.Class().Remove("hidden")
	//				logutil.Logf("Opened modal (fallback): %s", targetID)
	//			}
	//		}
	//	})
	//}
	//
	//// Handle modal close buttons with HSOverlay API
	//modalCloseButtons := doc.QuerySelectorAll(".modal-close")
	//for _, closeBtn := range modalCloseButtons {
	//	closeBtn.AddEventListener("click", false, func(event dom.Event) {
	//		event.PreventDefault()
	//
	//		// Find the parent modal and close it using HSOverlay
	//		modal := closeBtn.Closest(".modal")
	//		if modal != nil {
	//			if hsOverlay := js.Global().Get("HSOverlay"); !hsOverlay.IsUndefined() {
	//				instance := hsOverlay.Call("getInstance", modal.Underlying())
	//				if !instance.IsUndefined() {
	//					instance.Call("close")
	//					logutil.Log("Closed modal using HSOverlay")
	//				} else {
	//					// Fallback
	//					modal.Class().Add("hidden")
	//					logutil.Log("Closed modal (fallback)")
	//				}
	//			} else {
	//				// Fallback
	//				modal.Class().Add("hidden")
	//				logutil.Log("Closed modal (fallback)")
	//			}
	//		}
	//	})
	//}
	//
	//logutil.Logf("Set up %d modal triggers and %d close buttons with HSOverlay", len(modalTriggers), len(modalCloseButtons))
}

// setupManualModalListeners provides simple modal toggle via [data-overlay] attributes per updated FlyonUI docs
func setupManualModalListeners(doc dom.Document) {
	// Find all elements with data-overlay attribute
	triggers := doc.QuerySelectorAll("[data-overlay]")
	for _, el := range triggers {
		el.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			// Get selector from data-overlay (e.g., "#demo-modal")
			sel := el.GetAttribute("data-overlay")
			if sel == "" {
				return
			}
			// Query target modal by selector
			target := doc.QuerySelector(sel)
			if target == nil {
				logutil.Logf("data-overlay target not found: %s", sel)
				return
			}
			// Toggle hidden class
			if target.Class().Contains("hidden") {
				target.Class().Remove("hidden")
				logutil.Logf("Opened modal: %s", sel)
			} else {
				target.Class().Add("hidden")
				logutil.Logf("Closed modal: %s", sel)
			}
		})
	}
	logutil.Logf("Set up %d [data-overlay] triggers", len(triggers))
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
