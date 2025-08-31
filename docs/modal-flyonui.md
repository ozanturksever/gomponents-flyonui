# Modal - FlyonUI

Last synced: 2025-08-31 • Source: https://flyonui.com/docs/docs/overlays/modal

## Official Classes and Structure
- Container:
  - `overlay modal`: Root modal container with backdrop behavior.
  - State utilities: `overlay-open:opacity-100`, `overlay-open:duration-300` (example transition utilities from docs).
  - Visibility: `hidden` is commonly used so the modal starts hidden when controlled via `data-overlay`.
- Structure inside the container:
  - `modal-dialog`: Wrapper for the dialog panel.
  - `modal-content`: Panel that holds header, body, and footer.
  - `modal-header`: Header region.
  - `modal-title`: Title element in header.
  - `modal-body`: Scrollable content area.
  - `modal-footer`: Action area (buttons, etc.).
- Position variants:
  - `modal-top-center` (default top-center positioning).
  - `modal-middle` (centered in the viewport).
- Triggers and close controls:
  - Use `data-overlay="#<modal-id>"` on any element to open or close the modal with that ID.

Accessibility best practices from docs examples:
- Triggers include `aria-haspopup="dialog"`, `aria-expanded="false"`, and `aria-controls="<modal-id>"`.
- Modal container includes `role="dialog"` and `tabindex="-1"`.

## HTML Examples

### Basic Modal
```html
<button type="button" class="btn btn-primary" aria-haspopup="dialog" aria-expanded="false" aria-controls="basic-modal" data-overlay="#basic-modal">
  Open modal
</button>

<div id="basic-modal" class="overlay modal overlay-open:opacity-100 overlay-open:duration-300 hidden" role="dialog" tabindex="-1">
  <div class="modal-dialog overlay-open:opacity-100 overlay-open:duration-300">
    <div class="modal-content">
      <div class="modal-header">
        <h3 class="modal-title">Dialog Title</h3>
        <button type="button" class="btn btn-text btn-circle btn-sm absolute end-3 top-3" aria-label="Close" data-overlay="#basic-modal">
          <span class="icon-[tabler--x] size-4"></span>
        </button>
      </div>
      <div class="modal-body">
        This is some placeholder content to show the scrolling behavior for modals.
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-soft btn-secondary" data-overlay="#basic-modal">Close</button>
        <button type="button" class="btn btn-primary">Save changes</button>
      </div>
    </div>
  </div>
</div>
```

### Middle Center Position
```html
<button type="button" class="btn btn-primary" aria-haspopup="dialog" aria-expanded="false" aria-controls="middle-center-modal" data-overlay="#middle-center-modal">
  Middle center
</button>

<div id="middle-center-modal" class="overlay modal overlay-open:opacity-100 overlay-open:duration-300 modal-middle hidden" role="dialog" tabindex="-1">
  <div class="modal-dialog overlay-open:opacity-100 overlay-open:duration-300">
    <div class="modal-content">
      <div class="modal-header">
        <h3 class="modal-title">Dialog Title</h3>
        <button type="button" class="btn btn-text btn-circle btn-sm absolute end-3 top-3" aria-label="Close" data-overlay="#middle-center-modal">
          <span class="icon-[tabler--x] size-4"></span>
        </button>
      </div>
      <div class="modal-body">
        This is some placeholder content to show the scrolling behavior for modals.
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-soft btn-secondary" data-overlay="#middle-center-modal">Close</button>
        <button type="button" class="btn btn-primary">Save changes</button>
      </div>
    </div>
  </div>
</div>
```

## Configuration via Data Attributes
- `data-overlay`: Primary attribute to control opening/closing. Its value should be the CSS selector for the modal element (e.g., `#basic-modal`). Apply it to triggers and close buttons.
- `data-overlay-options`: Optional JSON configuration inline. Example:
```html
<button type="button" class="btn btn-primary" aria-haspopup="dialog" aria-expanded="false" aria-controls="demo-hidden-modal"
  data-overlay="#demo-hidden-modal"
  data-overlay-options='{ "hiddenClass": "hidden" }'>
  Open modal
</button>
```
- Transition utilities like `overlay-open:*` classes can be added to the modal container and/or dialog to tune enter/leave animations as shown above.

## Notes
- The following are NOT part of current FlyonUI Modal and should not be used: `.modal-box`, `.modal-trigger`, `.modal-close`, `data-modal-target`, `modal-open`, `modal-sm`, `modal-md`, `modal-lg`, `modal-xl`, `modal-full`, `modal-bottom`.
- Use standard FlyonUI button classes (e.g., `btn`, `btn-primary`, etc.) for actions inside the footer.