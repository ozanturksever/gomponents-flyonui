# Drawer - FlyonUI

## Complete List of Classes
- **Base:**
  - `drawer`: The main class for the drawer component.
  - `drawer-header`: The header section of the drawer.
  - `drawer-body`: The main content area of the drawer.
  - `drawer-footer`: The footer section of the drawer.
  - `drawer-title`: For the title text within a drawer.
- **Positioning:**
  - `drawer-start`: Positions the drawer on the left (start).
  - `drawer-end`: Positions the drawer on the right (end).
  - `drawer-top`: Positions the drawer on the top.
  - `drawer-bottom`: Positions the drawer on the bottom.
- **Overlay & State:**
  - `overlay`: The main class for the overlay/backdrop.
  - `overlay-open:translate-x-0`: A utility to show the drawer when open (for horizontal drawers).
  - `overlay-open:translate-y-0`: A utility to show the drawer when open (for vertical drawers).
  - `hidden`: Initially hides the drawer.
  - `overlay-backdrop-open:bg-warning/40`: Customizes the backdrop color and opacity.

## Variations and Sizes
The size of the drawer is controlled by Tailwind CSS max-width utilities (e.g., `max-w-sm`, `max-w-72`). The position is controlled by `drawer-start`, `drawer-end`, `drawer-top`, and `drawer-bottom`.

## HTML Examples

### Basic Drawer (from left)
```html
<button type="button" class="btn btn-primary" aria-haspopup="dialog" aria-expanded="false" aria-controls="overlay-example" data-overlay="#overlay-example">Open drawer</button>

<div id="overlay-example" class="overlay overlay-open:translate-x-0 drawer drawer-start hidden" role="dialog" tabindex="-1">
  <div class="drawer-header">
    <h3 class="drawer-title">Drawer Title</h3>
    <button type="button" class="btn btn-text btn-circle btn-sm absolute end-3 top-3" aria-label="Close" data-overlay="#overlay-example">
      <span class="icon-[tabler--x] size-5"></span>
    </button>
  </div>
  <div class="drawer-body">
    <p>Some text as placeholder...</p>
  </div>
  <div class="drawer-footer">
    <button type="button" class="btn btn-soft btn-secondary" data-overlay="#overlay-example">Close</button>
    <button type="button" class="btn btn-primary">Save changes</button>
  </div>
</div>
```

### Drawer from the right
```html
<div id="overlay-end-example" class="overlay overlay-open:translate-x-0 drawer drawer-end hidden" role="dialog" tabindex="-1">
  ...
</div>
```

### Drawer from the top
```html
<div id="overlay-top-example" class="overlay drawer overlay-open:translate-y-0 drawer-top hidden" role="dialog" tabindex="-1">
  ...
</div>
```

### Drawer with a Form
```html
<div id="overlay-form-example" class="overlay overlay-open:translate-x-0 drawer drawer-end hidden justify-start" role="dialog" tabindex="-1" >
  <div class="drawer-header">...</div>
  <form>
    <div class="drawer-body justify-start">
      <div class="mb-4">
        <label class="label-text" for="fullName"> Full Name </label>
        <input type="text" placeholder="John Doe" class="input" id="fullName" />
      </div>
      ...
    </div>
    <div class="drawer-footer">
      <button type="button" class="btn btn-soft btn-secondary" data-overlay="#overlay-form-example">Close</button>
      <button type="submit" class="btn btn-primary">Save changes</button>
    </div>
  </form>
</div>
```

## JavaScript Interaction API
The Drawer is an "Overlay" component in FlyonUI and is controlled by the `HSOverlay` object.

### Initialization
The drawer is automatically initialized if you are using the FlyonUI script and the correct `data-overlay` attributes.

### Programmatic Control
You can open and close the drawer programmatically.
```javascript
// Get the overlay instance
const myDrawer = HSOverlay.getInstance('#my-drawer-id', true);

// Open the drawer
myDrawer.element.open();

// Close the drawer
myDrawer.element.close();
```

## Configuration Options
- `data-overlay`: The primary attribute to enable the drawer functionality. Its value should be a CSS selector pointing to the drawer element.
- `--body-scroll:true`: A CSS variable set on the drawer element to allow the body to scroll when the drawer is open.
- `--overlay-backdrop:false`: A CSS variable to disable the backdrop.
