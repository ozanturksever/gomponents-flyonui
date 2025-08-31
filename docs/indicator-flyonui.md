# Indicator - FlyonUI

## Complete List of Classes
- **Base:**
  - `indicator`: The main wrapper class for an element and its indicator.
  - `indicator-item`: The class for the indicator element itself, which is positioned relative to the main element.
- **Positioning:**
  - `indicator-top`, `indicator-middle`, `indicator-bottom`: Vertical alignment.
  - `indicator-start`, `indicator-center`, `indicator-end`: Horizontal alignment.
- **Indicator Content:**
  - `badge`: Often used for indicators with text (e.g., a number).
  - `status`: Used for dot-like status indicators.

## Variations and Sizes
The indicator component is highly flexible. The size and style of the indicator depend on what element is used as the `indicator-item` (e.g., a `badge` or `status` dot).

## HTML Examples

### Basic Indicator
A simple dot indicator.
```html
<div class="indicator">
  <span class="indicator-item bg-primary size-3 rounded-full"></span>
  <div class="bg-primary/10 border-primary grid place-items-center rounded-md border p-3">
    <span class="icon-[tabler--bell] text-primary size-5"></span>
  </div>
</div>
```

### Indicator with a Badge
```html
<div class="indicator">
  <span class="indicator-item badge badge-primary">+999</span>
  <div class="bg-primary/10 border-primary grid place-items-center rounded-md border p-3">
    <span class="icon-[tabler--bell] text-primary size-5"></span>
  </div>
</div>
```

### Indicator on an Avatar
```html
<div class="avatar indicator">
  <span class="indicator-item badge badge-primary">typing…</span>
  <div class="size-16 rounded-md">
    <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-8.png" alt="avatar" />
  </div>
</div>
```

### Indicator on a Button (as a close button)
```html
<div class="indicator mt-4">
  <div class="indicator-item indicator-top">
    <button class="btn btn-error btn-circle btn-sm" aria-label="Close Button Indicator">
      <span class="icon-[tabler--x] size-5"></span>
    </button>
  </div>
  <div class="card">
    <div class="card-body">
      <h2 class="card-title">Card Title</h2>
      <p>Rerum reiciendis beatae tenetur excepturi</p>
    </div>
  </div>
</div>
```

### Positioning
Combine vertical and horizontal classes to position the indicator.
```html
<div class="indicator">
  <span class="indicator-item indicator-top indicator-start bg-primary size-3 rounded-full"></span>
  <div class="grid h-32 w-60 place-items-center bg-base-300">content</div>
</div>
```

## JavaScript Interaction API
The indicator component is a CSS-only component and does not have a JavaScript API.

## Configuration Options
- **CSS Variables:** You can control the position of the indicator with these CSS variables.
  - `--indicator-t`
  - `--indicator-b`
  - `--indicator-s`
  - `--indicator-e`
  - `--indicator-y`
  - `--indicator-x`