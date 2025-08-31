//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/components"
	"github.com/ozanturksever/gomponents-flyonui/flyon"
	"github.com/ozanturksever/gomponents-flyonui/internal/vite"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
	"honnef.co/go/js/dom/v2"
	"maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func main() {
	logutil.Log("Starting Gomponents FlyonUI Showcase...")

	// Create asset resolver for Vite-generated assets
	assetResolver := vite.NewAssetResolver("", true)

	// Prevent the program from exiting
	c := make(chan struct{})

	// Wait for DOM to be ready
	js.Global().Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//logutil.Log("DOM Content Loaded")
		hydrate(assetResolver)
		return nil
	}))

	// If DOM is already loaded
	if js.Global().Get("document").Get("readyState").String() != "loading" {
		//logutil.Log("DOM already loaded")
		hydrate(assetResolver)
	}

	<-c
}

func hydrate(assetResolver *vite.AssetResolver) {
	logutil.Log("Starting hydration process...")

	// Initialize Vite asset resolver
	if err := assetResolver.LoadManifest(); err != nil {
		logutil.Logf("Warning: Could not load Vite manifest: %v", err)
	}

	//// Log asset information
	//cssURL := assetResolver.GetAssetURL("css/main.css")
	//jsURL := assetResolver.GetAssetURL("js/main.js")
	//logutil.Logf("Using CSS asset: %s", cssURL)
	//logutil.Logf("Using JS asset: %s", jsURL)

	// Wait for JavaScript to be ready before initializing FlyonUI

	// Initialize FlyonUI JavaScript components
	//initializeFlyonUIComponents()

	// Initialize all FlyonUI components via bridge
	//bridge.InitializeAllComponents()

	// Render the showcase using Gomponents
	renderShowcase()

	// Setup interactive event listeners
	//setupInteractiveListeners()

	// Complete hydration
	completeHydration()
}

// renderShowcase renders the entire showcase using Gomponents
func renderShowcase() {
	logutil.Log("Rendering Gomponents showcase...")

	doc := dom.GetWindow().Document()
	container := doc.GetElementByID("showcase-container")
	if container == nil {
		logutil.Log("Warning: showcase-container not found")
		return
	}

	// Create the showcase content using Gomponents
	showcaseContent := createShowcaseContent()

	// Render to string and set innerHTML
	var htmlBuilder strings.Builder
	if err := showcaseContent.Render(&htmlBuilder); err != nil {
		logutil.Logf("Error rendering showcase: %v", err)
		return
	}

	container.SetInnerHTML(htmlBuilder.String())
	logutil.Log("Showcase rendered successfully using Gomponents")
}

// createShowcaseContent creates the main showcase content using Gomponents
func createShowcaseContent() gomponents.Node {
	return h.Div(
		//h.Class("space-y-12"),

		// Header Section
		//createHeaderSection(),

		// Button Showcase
		//createButtonShowcase(),

		// Alert Showcase
		//createAlertShowcase(),

		// Card Showcase
		//createCardShowcase(),

		// Modal Showcase
		createModalShowcase(),

		// Dropdown Showcase
		//createDropdownShowcase(),

		// Accordion Showcase
		//createAccordionShowcase(),

		// Progress Showcase
		//createProgressShowcase(),

		// Layout Showcase
		//createLayoutShowcase(),

		// Form Components Showcase
		//createFormShowcase(),

		// Status Section
		//createStatusSection(),
	)
}

// createHeaderSection creates the header using Gomponents
func createHeaderSection() gomponents.Node {
	return h.Div(
		h.Class("text-center mb-12"),
		h.H1(
			h.Class("text-4xl font-bold text-primary mb-4"),
			gomponents.Text("Gomponents FlyonUI Showcase"),
		),
		h.P(
			h.Class("text-lg text-base-content/70"),
			gomponents.Text("Comprehensive demonstration of FlyonUI components built with Gomponents in WebAssembly"),
		),
	)
}

// createButtonShowcase creates button examples using Gomponents
func createButtonShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Buttons"),
		),
		h.Div(
			h.Class("grid grid-cols-2 md:grid-cols-4 gap-4"),

			// Primary Button
			components.NewButton(gomponents.Text("Primary")).With(flyon.Primary),

			// Secondary Button
			components.NewButton(gomponents.Text("Secondary")).With(flyon.Secondary),

			// Success Button
			components.NewButton(gomponents.Text("Success")).With(flyon.Success),

			// Error Button
			components.NewButton(gomponents.Text("Error")).With(flyon.Error),

			// Large Button
			components.NewButton(gomponents.Text("Large")).With(flyon.Primary, flyon.SizeLarge),

			// Small Button
			components.NewButton(gomponents.Text("Small")).With(flyon.Secondary, flyon.SizeSmall),

			// Disabled Button
			components.NewButton(gomponents.Text("Disabled"), h.Disabled()).With(flyon.Primary),

			// Interactive Button
			components.NewButton(gomponents.Text("Click Me!"), h.ID("interactive-btn")).With(flyon.Success),
		),
	)
}

// createAlertShowcase creates alert examples using Gomponents
func createAlertShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Alerts"),
		),
		h.Div(
			h.Class("space-y-4"),

			// Info Alert
			components.NewAlert(
				h.SVG(
					gomponents.Attr("xmlns", "http://www.w3.org/2000/svg"),
					gomponents.Attr("fill", "none"),
					gomponents.Attr("viewBox", "0 0 24 24"),
					h.Class("stroke-current shrink-0 w-6 h-6"),
					gomponents.El("path",
						gomponents.Attr("stroke-linecap", "round"),
						gomponents.Attr("stroke-linejoin", "round"),
						gomponents.Attr("stroke-width", "2"),
						gomponents.Attr("d", "M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"),
					),
				),
				h.Span(
					gomponents.Text("This is an informational alert created with Gomponents."),
				),
				h.Button(
					h.Class("btn btn-sm btn-ghost alert-close"),
					h.ID("info-alert-close"),
					gomponents.Text("×"),
				),
			).With(flyon.Info),

			// Success Alert
			components.NewAlert(
				h.SVG(
					gomponents.Attr("xmlns", "http://www.w3.org/2000/svg"),
					h.Class("stroke-current shrink-0 h-6 w-6"),
					gomponents.Attr("fill", "none"),
					gomponents.Attr("viewBox", "0 0 24 24"),
					gomponents.El("path",
						gomponents.Attr("stroke-linecap", "round"),
						gomponents.Attr("stroke-linejoin", "round"),
						gomponents.Attr("stroke-width", "2"),
						gomponents.Attr("d", "M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"),
					),
				),
				h.Span(
					gomponents.Text("Success! Operation completed successfully using Gomponents."),
				),
				h.Button(
					h.Class("btn btn-sm btn-ghost alert-close"),
					h.ID("success-alert-close"),
					gomponents.Text("×"),
				),
			).With(flyon.Success),

			// Warning Alert
			components.NewAlert(
				h.SVG(
					gomponents.Attr("xmlns", "http://www.w3.org/2000/svg"),
					h.Class("stroke-current shrink-0 h-6 w-6"),
					gomponents.Attr("fill", "none"),
					gomponents.Attr("viewBox", "0 0 24 24"),
					gomponents.El("path",
						gomponents.Attr("stroke-linecap", "round"),
						gomponents.Attr("stroke-linejoin", "round"),
						gomponents.Attr("stroke-width", "2"),
						gomponents.Attr("d", "M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.34 16.5c-.77.833.192 2.5 1.732 2.5z"),
					),
				),
				h.Span(
					gomponents.Text("Warning: Please review your input before proceeding."),
				),
				h.Button(
					h.Class("btn btn-sm btn-ghost alert-close"),
					h.ID("warning-alert-close"),
					gomponents.Text("×"),
				),
			).With(flyon.Warning),
		),
	)
}

// createCardShowcase creates card examples using Gomponents
func createCardShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Cards"),
		),
		h.Div(
			h.Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"),

			// Basic Card
			components.NewCard(
				h.Div(
					h.Class("card-body"),
					h.H3(
						h.Class("card-title"),
						gomponents.Text("Basic Card"),
					),
					h.P(
						gomponents.Text("This is a basic card component created with Gomponents."),
					),
					h.Div(
						h.Class("card-actions justify-end"),
						components.NewButton(gomponents.Text("Action")).With(flyon.Primary, flyon.SizeSmall),
					),
				),
			).With(flyon.Primary),

			// Feature Card
			components.NewCard(
				h.Div(
					h.Class("card-body"),
					h.H3(
						h.Class("card-title"),
						gomponents.Text("Feature Card"),
					),
					h.P(
						gomponents.Text("Showcasing advanced features with interactive elements."),
					),
					h.Div(
						h.Class("card-actions justify-between"),
						components.NewButton(gomponents.Text("Learn More")).With(flyon.VariantGhost, flyon.SizeSmall),
						components.NewButton(gomponents.Text("Try It")).With(flyon.Secondary, flyon.SizeSmall),
					),
				),
			).With(flyon.Secondary),

			// Stats Card
			components.NewCard(
				h.Div(
					h.Class("card-body text-center"),
					h.H3(
						h.Class("card-title"),
						gomponents.Text("Statistics"),
					),
					h.Div(
						h.Class("stats stats-vertical shadow"),
						h.Div(
							h.Class("stat"),
							h.Div(
								h.Class("stat-value text-primary"),
								gomponents.Text("25.6K"),
							),
							h.Div(
								h.Class("stat-desc"),
								gomponents.Text("Components"),
							),
						),
					),
				),
			).With(flyon.Info),
		),
	)
}

// createModalShowcase creates modal examples using Gomponents
func createModalShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Modals"),
		),
		h.Div(
			h.Class("flex gap-4 flex-wrap"),

			// Modal Trigger Buttons
			components.NewButton(
				gomponents.Text("Open Modal"),
				h.DataAttr("overlay", "#gomponents-modal"),
				h.Aria("haspopup", "dialog"),
				h.Aria("expanded", "false"),
				h.Aria("controls", "gomponents-modal"),
			).With(flyon.Primary),

			components.NewButton(
				gomponents.Text("Confirm Dialog"),
				h.DataAttr("overlay", "#confirm-modal"),
				h.Aria("haspopup", "dialog"),
				h.Aria("expanded", "false"),
				h.Aria("controls", "confirm-modal"),
			).With(flyon.Secondary),
		),

		// Modal Structures
		components.NewModal("Gomponents Modal",
			h.P(
				gomponents.Text("This modal was created entirely using Gomponents! It demonstrates how to build complex UI components with type-safe Go code that compiles to WebAssembly."),
			),
			h.Ul(
				h.Class("list-disc list-inside mt-4 space-y-2"),
				h.Li(gomponents.Text("Type-safe component construction")),
				h.Li(gomponents.Text("Compile-time validation")),
				h.Li(gomponents.Text("WebAssembly performance")),
				h.Li(gomponents.Text("FlyonUI styling integration")),
			),
		).WithID("gomponents-modal").WithSize(components.ModalSizeLarge).WithActions(
			components.NewButton(gomponents.Text("Save Changes")).With(flyon.Primary),
			components.NewButton(
				gomponents.Text("Cancel"),
				h.DataAttr("overlay", "#gomponents-modal"),
			).With(flyon.VariantGhost),
		),

		components.NewModal("Confirm Action",
			h.P(
				gomponents.Text("Are you sure you want to proceed with this action? This operation cannot be undone."),
			),
		).WithID("confirm-modal").WithSize(components.ModalSizeSmall).WithActions(
			components.NewButton(gomponents.Text("Delete")).With(flyon.Error),
			components.NewButton(
				gomponents.Text("Cancel"),
				h.DataAttr("overlay", "#confirm-modal"),
			).With(flyon.VariantGhost),
		),
	)
}

// createDropdownShowcase creates dropdown examples using Gomponents
func createDropdownShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Dropdowns"),
		),
		h.Div(
			h.Class("flex gap-4 flex-wrap"),

			// Primary Dropdown
			components.NewDropdown(
				components.NewButton(gomponents.Text("Primary Menu")).With(flyon.Primary),
				components.DropdownItem(
					gomponents.Text("Profile Settings"),
				),
				components.DropdownItem(
					gomponents.Text("Account Management"),
				),
				components.DropdownDivider(),
				components.DropdownItem(
					gomponents.Text("Sign Out"),
				),
			),

			// Secondary Dropdown
			components.NewDropdown(
				components.NewButton(gomponents.Text("Actions")).With(flyon.Secondary),
				components.DropdownHeader("Quick Actions"),
				components.DropdownItem(
					gomponents.Text("Create New"),
				),
				components.DropdownItem(
					gomponents.Text("Import Data"),
				),
				components.DropdownItem(
					gomponents.Text("Export Report"),
				),
			),
		),
	)
}

// createAccordionShowcase creates accordion examples using Gomponents
func createAccordionShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Accordion"),
		),
		components.NewAccordion(
			components.AccordionItem{
				ID:      "item1",
				Title:   "What is Gomponents?",
				Content: h.P(gomponents.Text("Gomponents is a library for building HTML components in Go. It provides a type-safe way to construct HTML using Go's type system, making it easier to build and maintain web applications.")),
				Open:    true,
			},
			components.AccordionItem{
				ID:      "item2",
				Title:   "How does WebAssembly integration work?",
				Content: h.P(gomponents.Text("WebAssembly allows Go code to run in the browser. This showcase demonstrates how Gomponents can be used to build interactive web applications that run entirely in the browser using WASM.")),
				Open:    false,
			},
			components.AccordionItem{
				ID:    "item3",
				Title: "What are the benefits of this approach?",
				Content: h.Div(
					h.P(gomponents.Text("Using Gomponents with WebAssembly provides several benefits:")),
					h.Ul(
						h.Class("list-disc list-inside mt-2 space-y-1"),
						h.Li(gomponents.Text("Type safety at compile time")),
						h.Li(gomponents.Text("No runtime template errors")),
						h.Li(gomponents.Text("Better IDE support and refactoring")),
						h.Li(gomponents.Text("Shared code between server and client")),
					),
				),
				Open: false,
			},
		).WithID("showcase-accordion"),
	)
}

// createProgressShowcase creates progress examples using Gomponents
func createProgressShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Progress Indicators"),
		),
		h.Div(
			h.Class("space-y-4"),

			// Progress bars with different colors and values
			h.Div(
				h.Class("space-y-2"),
				h.Label(
					h.Class("text-sm font-medium"),
					gomponents.Text("Primary Progress (32%)"),
				),
				components.NewProgress(32).With(flyon.Primary),
			),

			h.Div(
				h.Class("space-y-2"),
				h.Label(
					h.Class("text-sm font-medium"),
					gomponents.Text("Secondary Progress (70%)"),
				),
				components.NewProgress(70).With(flyon.Secondary),
			),

			h.Div(
				h.Class("space-y-2"),
				h.Label(
					h.Class("text-sm font-medium"),
					gomponents.Text("Success Progress (90%)"),
				),
				components.NewProgress(90).With(flyon.Success),
			),

			h.Div(
				h.Class("space-y-2"),
				h.Label(
					h.Class("text-sm font-medium"),
					gomponents.Text("Indeterminate Progress"),
				),
				components.NewIndeterminateProgress().With(flyon.Info), // No value = indeterminate
			),
		),
	)
}

// createLayoutShowcase creates layout examples using Gomponents
func createLayoutShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Layout Components"),
		),
		h.Div(
			h.Class("space-y-8"),

			// Container Example
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Container"),
				),
				components.NewContainer().With(flyon.Primary),
			),

			// Grid Example
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Grid Layout"),
				),
				h.Div(
					h.Class("grid grid-cols-1 md:grid-cols-3 gap-4"),
					h.Div(
						h.Class("p-4 bg-primary text-primary-content rounded"),
						gomponents.Text("Grid Item 1"),
					),
					h.Div(
						h.Class("p-4 bg-secondary text-secondary-content rounded"),
						gomponents.Text("Grid Item 2"),
					),
					h.Div(
						h.Class("p-4 bg-accent text-accent-content rounded"),
						gomponents.Text("Grid Item 3"),
					),
				),
			),

			// Flex Example
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Flex Layout"),
				),
				h.Div(
					h.Class("flex justify-between items-center p-4 bg-base-200 rounded"),
					h.Span(
						gomponents.Text("Flex Start"),
					),
					h.Span(
						gomponents.Text("Flex Center"),
					),
					h.Span(
						gomponents.Text("Flex End"),
					),
				),
			),

			// Stack Example
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Stack Layout"),
				),
				h.Div(
					h.Class("stack space-y-2"),
					h.Div(
						h.Class("p-2 bg-info text-info-content rounded"),
						gomponents.Text("Stack Item 1"),
					),
					h.Div(
						h.Class("p-2 bg-warning text-warning-content rounded"),
						gomponents.Text("Stack Item 2"),
					),
					h.Div(
						h.Class("p-2 bg-error text-error-content rounded"),
						gomponents.Text("Stack Item 3"),
					),
				),
			),
		),
	)
}

// createFormShowcase creates form component examples using Gomponents
func createFormShowcase() gomponents.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-6"),
			gomponents.Text("Form Components"),
		),
		h.Div(
			h.Class("grid grid-cols-1 md:grid-cols-2 gap-8"),

			// Toggle Examples
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Toggle Switches"),
				),
				h.Div(
					h.Class("space-y-4"),
					components.NewToggle().WithID("toggle1").WithName("toggle1").WithChecked(true).With(flyon.Primary),
					h.Label(
						h.For("toggle1"),
						h.Class("cursor-pointer"),
						gomponents.Text("Primary Toggle (Checked)"),
					),

					components.NewToggle().WithID("toggle2").WithName("toggle2").With(flyon.Secondary),
					h.Label(
						h.For("toggle2"),
						h.Class("cursor-pointer"),
						gomponents.Text("Secondary Toggle"),
					),

					components.NewToggle().WithID("toggle3").WithName("toggle3").With(flyon.Success, flyon.SizeLarge),
					h.Label(
						h.For("toggle3"),
						h.Class("cursor-pointer"),
						gomponents.Text("Large Success Toggle"),
					),
				),
			),

			// Avatar Examples
			h.Div(
				h.H3(
					h.Class("text-lg font-medium mb-4"),
					gomponents.Text("Avatars"),
				),
				h.Div(
					h.Class("flex gap-4 items-center"),
					components.NewAvatar(
						h.Div(
							h.Class("w-8 h-8 rounded-full bg-primary flex items-center justify-center text-primary-content font-bold text-sm"),
							gomponents.Text("S"),
						),
					).With(flyon.Primary, flyon.SizeSmall),
					components.NewAvatar(
						h.Div(
							h.Class("w-12 h-12 rounded-full bg-secondary flex items-center justify-center text-secondary-content font-bold"),
							gomponents.Text("M"),
						),
					).With(flyon.Secondary),
					components.NewAvatar(
						h.Div(
							h.Class("w-16 h-16 rounded-full bg-accent flex items-center justify-center text-accent-content font-bold text-lg"),
							gomponents.Text("L"),
						),
					).With(flyon.Info, flyon.SizeLarge),
				),
			),
		),
	)
}

// createStatusSection creates the status section using Gomponents
func createStatusSection() gomponents.Node {
	return h.Section(
		h.Class("text-center"),
		h.Div(
			h.Class("stats shadow"),
			h.Div(
				h.Class("stat"),
				h.Div(
					h.Class("stat-title"),
					gomponents.Text("WASM Status"),
				),
				h.Div(
					h.Class("stat-value text-primary"),
					h.ID("wasm-status"),
					gomponents.Text("Loading..."),
				),
				h.Div(
					h.Class("stat-desc"),
					gomponents.Text("WebAssembly Runtime"),
				),
			),
			h.Div(
				h.Class("stat"),
				h.Div(
					h.Class("stat-title"),
					gomponents.Text("Components"),
				),
				h.Div(
					h.Class("stat-value text-secondary"),
					gomponents.Text("Gomponents"),
				),
				h.Div(
					h.Class("stat-desc"),
					gomponents.Text("Type-Safe Components"),
				),
			),
			h.Div(
				h.Class("stat"),
				h.Div(
					h.Class("stat-title"),
					gomponents.Text("Framework"),
				),
				h.Div(
					h.Class("stat-value text-accent"),
					gomponents.Text("FlyonUI"),
				),
				h.Div(
					h.Class("stat-desc"),
					gomponents.Text("Modern CSS Framework"),
				),
			),
		),
	)
}

// setupInteractiveListeners sets up event listeners for interactive components
func setupInteractiveListeners() {
	logutil.Log("Setting up interactive listeners for Gomponents showcase...")

	doc := dom.GetWindow().Document()

	// Interactive button click handler
	interactiveBtn := doc.GetElementByID("interactive-btn")
	if interactiveBtn != nil {
		interactiveBtn.AddEventListener("click", false, func(event dom.Event) {
			event.PreventDefault()
			logutil.Log("Interactive button clicked!")

			// Change button text temporarily
			originalText := interactiveBtn.TextContent()
			interactiveBtn.SetTextContent("Clicked!")

			// Reset after 1 second
			time.AfterFunc(1*time.Second, func() {
				interactiveBtn.SetTextContent(originalText)
			})
		})
	}

	// Alert close handlers
	//setupAlertCloseHandlers(doc)

	// Modal overlay handlers (FlyonUI handles most modal functionality)
	//setupModalOverlayHandlers(doc)

	logutil.Log("Interactive listeners setup complete")
}

// setupAlertCloseHandlers sets up close handlers for alerts
func setupAlertCloseHandlers(doc dom.Document) {
	alertCloseButtons := []string{"info-alert-close", "success-alert-close", "warning-alert-close"}

	for _, buttonID := range alertCloseButtons {
		closeBtn := doc.GetElementByID(buttonID)
		if closeBtn != nil {
			closeBtn.AddEventListener("click", false, func(event dom.Event) {
				event.PreventDefault()
				logutil.Logf("Alert close button clicked: %s", buttonID)

				// Find the parent alert and hide it
				alert := closeBtn.Closest(".alert")
				if alert != nil {
					alert.Underlying().Get("style").Set("display", "none")
					logutil.Log("Alert hidden")
				}
			})
		}
	}
}

// setupModalOverlayHandlers sets up additional modal functionality
func setupModalOverlayHandlers(doc dom.Document) {
	// FlyonUI handles most modal functionality via data-overlay attributes
	// We just need to ensure the overlays are properly initialized
	logutil.Log("Modal overlay handlers ready (handled by FlyonUI)")
}

// completeHydration finalizes the hydration process
func completeHydration() {
	// Update WASM status to indicate readiness
	doc := dom.GetWindow().Document()
	wasmStatus := doc.GetElementByID("wasm-status")
	if wasmStatus != nil {
		wasmStatus.SetTextContent("Ready")
	}

	logutil.Log("WASM status updated to Ready")
	// Dispatch a custom 'wasmReady' event
	time.AfterFunc(100*time.Millisecond, func() {
		evt := js.Global().Get("CustomEvent").New("wasmReady")
		js.Global().Call("dispatchEvent", evt)
		logutil.Log("WASM ready event dispatched", evt)
	})

	logutil.Log("Gomponents showcase hydration complete!")
}
