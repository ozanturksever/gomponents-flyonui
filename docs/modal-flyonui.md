# Modal - FlyonUI

## Complete List of Classes
- **Core Modal Structure:**
  - `.modal`: Main container for the modal and its backdrop.
  - `.modal-box`: The visible container for the modal's content.
  - `.modal-header`: Header section of the modal.
  - `.modal-title`: Title within the header.
  - `.modal-body`: Main content area of the modal.
  - `.modal-footer`: Footer section for actions.
  - `.modal-action`: Wrapper for action buttons in the footer.
- **Sizing and Positioning:**
  - `modal-sm`, `modal-md`, `modal-lg`, `modal-xl`, `modal-full`: For sizing the modal.
  - `modal-middle`: Vertically centers the modal.
  - `modal-bottom`: Positions the modal at the bottom of the screen.
- **State & Behavior:**
  - `modal-open`: Added to the `.modal` element to make it visible.
  - `modal-trigger`: Class for buttons that open a modal.
  - `modal-close`: Class for buttons that close a modal.
- **Overlay (Backdrop):**
  - `overlay`: Base class, often used with `.modal`.
  - `overlay-open:opacity-100`: Controls the backdrop's visibility.

## Variations and Sizes

### Size Variants
```html
<!-- Default (Medium) -->
<div class="modal-box">...</div>

<!-- Small -->
<div class="modal-box modal-sm">...</div>

<!-- Large -->
<div class="modal-box modal-lg">...</div>

<!-- Extra Large -->
<div class="modal-box modal-xl">...</div>
```

### Position Variants
```html
<!-- Vertically Centered -->
<div class="modal modal-middle">...</div>

<!-- At the bottom -->
<div class="modal modal-bottom">...</div>
```

## HTML Examples

### Basic Modal
```html
<!-- Trigger -->
<button class="modal-trigger btn btn-primary" data-modal-target="demo-modal">Open Modal</button>

<!-- Modal Structure -->
<div id="demo-modal" class="modal hidden">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Demo Modal</h3>
    <p class="py-4">This is a demonstration modal.</p>
    <div class="modal-action">
      <button class="modal-close btn">Close</button>
    </div>
  </div>
</div>
```

### Confirmation Dialog
```html
<!-- Trigger -->
<button class="modal-trigger btn btn-warning" data-modal-target="confirm-modal">Confirm Dialog</button>

<!-- Modal Structure -->
<div id="confirm-modal" class="modal hidden">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Confirm Action</h3>
    <p class="py-4">Are you sure you want to proceed?</p>
    <div class="modal-action flex gap-2">
      <button class="modal-close btn btn-error">Cancel</button>
      <button class="modal-close btn btn-success">Confirm</button>
    </div>
  </div>
</div>
```

## JavaScript Interaction API
The modal component in FlyonUI is controlled via the `HSOverlay` JavaScript class.

### Initialization
The modal is typically initialized automatically via data attributes. To initialize manually:
```javascript
const modalElement = document.querySelector('#demo-modal');
const modal = new HSOverlay(modalElement);
```

### Methods
- `modal.open()`: Opens the modal.
- `modal.close()`: Closes the modal.
- `element.destroy()`: Destroys the modal instance, removing its functionality.
- `HSOverlay.autoInit()`: Reinitializes all overlay components, including modals.

### Events
The documentation does not specify a public event API, but functionality is primarily controlled via data attributes and methods.

## Configuration Options
Configuration is handled through `data` attributes on the trigger and modal elements.

### Triggers (`<button>`)
- `data-modal-target="modal-id"`: Specifies the ID of the modal to open.

### Modal (`<div class="modal">`)
- `data-overlay-keyboard="false"`: Disables closing the modal with the Escape key.
- `[--overlay-backdrop:false]`: Disables the modal backdrop via CSS custom property.
- `[--body-scroll:true]`: Allows the body to scroll when the modal is open.
- `[--has-autofocus:false]`: Disables autofocus on elements within the modal.