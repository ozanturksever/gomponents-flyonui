# Tooltip - FlyonUI

The Tooltip component in FlyonUI provides a small, contextual popup that displays information when a user hovers over or focuses on an element. It's built using semantic HTML and styled with Tailwind CSS classes.

## Classes

The core of the Tooltip component relies on the following classes:

*   `tooltip`: The base class that defines the tooltip's general appearance and behavior.
*   `tooltip-[position]`: Modifiers for positioning the tooltip relative to its parent element.
    *   `tooltip-top`: Positions the tooltip above the element.
    *   `tooltip-bottom`: Positions the tooltip below the element.
    *   `tooltip-left`: Positions the tooltip to the left of the element.
    *   `tooltip-right`: Positions the tooltip to the right of the element.
*   `tooltip-[color]`: Modifiers for applying different color themes to the tooltip. These correspond to the `flyon.Color` enum.
    *   `tooltip-primary`
    *   `tooltip-secondary`
    *   `tooltip-accent`
    *   `tooltip-success`
    *   `tooltip-warning`
    *   `tooltip-error`
    *   `tooltip-info`
    *   `tooltip-neutral`
*   `tooltip-open`: A class that forces the tooltip to be visible by default, overriding the hover/focus behavior.

## HTML Structure

A FlyonUI tooltip is typically a `div` element that wraps the content it's associated with. The tooltip text itself is provided via the `data-tip` attribute on the `tooltip` element.

```html
<div class="tooltip" data-tip="This is a simple tooltip">
  <button class="btn">Hover me</button>
</div>

<div class="tooltip tooltip-bottom" data-tip="Tooltip at the bottom">
  <span class="link">Link with tooltip</span>
</div>

<div class="tooltip tooltip-left tooltip-primary" data-tip="Primary tooltip on the left">
  <img src="https://via.placeholder.com/50" alt="Image" class="rounded-full">
</div>

<div class="tooltip tooltip-open tooltip-right tooltip-success" data-tip="Always open success tooltip">
  <p>Always visible</p>
</div>
```

## JavaScript Interaction API

FlyonUI tooltips are designed to be interactive, appearing on `mouseenter` and `mouseleave` events. While the core behavior is often handled by an underlying JavaScript library (referred to as `HSTooltip` in the codebase), you can also add custom Go-based logic in WebAssembly (WASM) to enhance or observe tooltip interactions.

### Initialization

The FlyonUI JavaScript library typically handles the automatic initialization of tooltips based on their class names and `data-tip` attributes.

In a WASM context, you might find hydration logic similar to this:

```go
// wasm/main.go
func hydrateTooltips() {
	doc := dom.GetWindow().Document()
	tooltipElements := doc.QuerySelectorAll("[data-tooltip]") // Note: The Go component uses data-tip, but JS might use data-tooltip
	
	for _, element := range tooltipElements {
		element.AddEventListener("mouseenter", false, func(event dom.Event) {
			logutil.Log("Tooltip hover started")
			// Custom logic on hover start
		})
		element.AddEventListener("mouseleave", false, func(event dom.Event) {
			logutil.Log("Tooltip hover ended")
			// Custom logic on hover end
		})
	}
	logutil.Logf("Hydrated %d tooltip elements", len(tooltipElements))
}
```

**Note on `data-tip` vs `data-tooltip`**: The Go `TooltipComponent` uses `data-tip` to specify the tooltip text. However, the WASM hydration example in `wasm/main.go` queries for `[data-tooltip]`. This suggests that the underlying JavaScript library might use `data-tooltip` for its initialization or that `data-tooltip` could be used for more advanced configurations. For basic usage, `data-tip` is sufficient for the text content. If `HSTooltip` is used, it likely handles the display based on these attributes.

### Configuration Options

The primary configuration for tooltips is done through CSS classes and the `data-tip` attribute. There are no explicit JavaScript methods exposed for direct manipulation of individual tooltip instances from the Go `TooltipComponent` itself, as the interaction is primarily CSS-driven or handled by the `HSTooltip` JavaScript library.

For advanced scenarios, you might need to interact with the `HSTooltip` JavaScript object directly if it exposes a public API.

## Examples

### Basic Tooltip

```html
<div class="tooltip" data-tip="Hello, I'm a tooltip!">
  <button class="btn btn-info">Hover over me</button>
</div>
```

### Tooltip Positions

```html
<div class="flex flex-wrap gap-4 justify-center">
  <div class="tooltip tooltip-top" data-tip="Top tooltip">
    <button class="btn">Top</button>
  </div>
  <div class="tooltip tooltip-bottom" data-tip="Bottom tooltip">
    <button class="btn">Bottom</button>
  </div>
  <div class="tooltip tooltip-left" data-tip="Left tooltip">
    <button class="btn">Left</button>
  </div>
  <div class="tooltip tooltip-right" data-tip="Right tooltip">
    <button class="btn">Right</button>
  </div>
</div>
```

### Tooltip Colors

```html
<div class="flex flex-wrap gap-4 justify-center">
  <div class="tooltip tooltip-primary" data-tip="Primary color">
    <button class="btn btn-primary">Primary</button>
  </div>
  <div class="tooltip tooltip-secondary" data-tip="Secondary color">
    <button class="btn btn-secondary">Secondary</button>
  </div>
  <div class="tooltip tooltip-accent" data-tip="Accent color">
    <button class="btn btn-accent">Accent</button>
  </div>
  <div class="tooltip tooltip-success" data-tip="Success color">
    <button class="btn btn-success">Success</button>
  </div>
  <div class="tooltip tooltip-warning" data-tip="Warning color">
    <button class="btn btn-warning">Warning</button>
  </div>
  <div class="tooltip tooltip-error" data-tip="Error color">
    <button class="btn btn-error">Error</button>
  </div>
  <div class="tooltip tooltip-info" data-tip="Info color">
    <button class="btn btn-info">Info</button>
  </div>
  <div class="tooltip tooltip-neutral" data-tip="Neutral color">
    <button class="btn btn-neutral">Neutral</button>
  </div>
</div>
```

### Always Open Tooltip

```html
<div class="tooltip tooltip-open tooltip-bottom" data-tip="This tooltip is always open">
  <button class="btn">Always Open</button>
</div>