# FlyonUI Modal Component Documentation

## 1. Complete API Reference

The FlyonUI Modal component is controlled via HTML data attributes and a JavaScript API.

### HTML Data Attributes

- `data-overlay="#modal-id"`: Triggers the modal with the specified ID.
- `aria-haspopup="dialog"`: Indicates that the element triggers a dialog.
- `aria-expanded="false"`: Indicates that the modal is initially closed.
- `aria-controls="modal-id"`: Associates the trigger with the modal.
- `data-overlay-keyboard="false"`: Disables closing the modal with the Escape key.
- `[--overlay-backdrop:false]`: Disables the modal backdrop.
- `[--has-autofocus:false]`: Disables autofocus on elements within the modal.
- `[--body-scroll:true]`: Allows the body to scroll when the modal is open.

### JavaScript API

The modal component can be controlled programmatically using the `HSOverlay` class.

**Initialization:**

```javascript
const modalElement = document.querySelector('#modal-target');
const modal = new HSOverlay(modalElement);
```

**Methods:**

- `modal.open()`: Opens the modal.
- `modal.close()`: Closes the modal.
- `element.destroy()`: Destroys the modal instance.
- `HSOverlay.autoInit()`: Reinitializes all overlay components.

**Lifecycle Events:**

The documentation does not explicitly list lifecycle events, but you can manage the lifecycle with `destroy` and `reinit` methods.

## 2. Full CSS Class Hierarchy and Utility Classes

### Core Modal Structure

- `.modal`: Main container for the modal.
- `.modal-dialog`: Wrapper for the modal content, controls alignment and width.
- `.modal-content`: The actual content box of the modal.
- `.modal-header`: Header section of the modal.
- `.modal-title`: Title within the header.
- `.modal-body`: Main content area of the modal.
- `.modal-footer`: Footer section for actions.

### Sizing and Positioning

- `.modal-dialog-xl`: Extra-large modal.
- `.modal-middle`: Vertically centers the modal.
- `.modal-middle-end`: Aligns the modal to the middle-end of the viewport.
- `.modal-top-end`: Aligns the modal to the top-end of the viewport.
- `sm:modal-bottom`: Positions the modal at the bottom on small screens.

### Animations and Transitions

- `.overlay`: Base class for the overlay.
- `.overlay-open:opacity-100`: Sets opacity to 100% when open.
- `.overlay-open:duration-300`: Sets the transition duration.
- `.overlay-open:mt-12`: Adds a top margin when open for slide-down effect.
- `.transition-all`, `.ease-out`: Utility classes for smooth transitions.

## 3. Animation and Transition Details

FlyonUI modals use Tailwind CSS utility classes for animations.

- **Fade-in:** Achieved with `overlay-open:opacity-100` and `overlay-open:duration-300`.
- **Slide-down:** Add `overlay-open:mt-12` to the `.modal-dialog`.
- **Slide-up:** Add `overlay-open:mt-4` and `mt-12` to the `.modal-dialog`.

## 4. Accessibility Compliance Guidelines and ARIA Attributes

- `role="dialog"`: Defines the element as a dialog.
- `tabindex="-1"`: Makes the modal focusable but not tabbable.
- `aria-haspopup="dialog"`: On the trigger button.
- `aria-expanded="false"`: On the trigger button, toggles to `true` when open.
- `aria-controls="modal-id"`: Links the trigger to the modal.
- `aria-label="Close"`: On the close button for screen readers.

## 5. Responsive Behavior Specifications

- Use responsive prefixes like `sm:` to change modal behavior on different screen sizes.
- Example: `sm:modal-bottom` will position the modal at the bottom on screens larger than the `sm` breakpoint.
- For fullscreen modals, use `.max-w-none` on `.modal-dialog` and `h-full max-h-none` on `.modal-content`.

## 6. Browser Compatibility Information

The documentation does not provide a specific browser compatibility table. However, since it is built with Tailwind CSS, it is expected to be compatible with all modern browsers.

## 7. Integration Patterns

The provided documentation focuses on HTML and vanilla JavaScript. Integration with frameworks like React, Vue, or Angular would involve wrapping the modal HTML structure in components and using the framework's lifecycle hooks to initialize and manage the `HSOverlay` instance.

## 8. Configuration Options with Default Values

- **Keyboard Control:** Enabled by default. Disable with `data-overlay-keyboard="false"`.
- **Backdrop:** Enabled by default. Disable with `[--overlay-backdrop:false]`.
- **Autofocus:** Enabled by default. Disable with `[--has-autofocus:false]`.
- **Body Scroll:** Disabled by default. Enable with `[--body-scroll:true]`.

## 9. Event Lifecycle Documentation

The documentation does not provide a detailed event lifecycle. However, you can manage the modal's state using the `destroy` and `reinit` methods, as shown in the JavaScript examples.

## 10. Performance Optimization Recommendations

- Use the `--prevent-on-load-init` class to defer initialization of modals that are not immediately visible.
- Manually initialize modals when they are needed to improve initial page load performance.

## 11. Migration Guides or Version Changes

No migration guides or version changes were found in the provided documentation.

## 12. TypeScript Definitions

No TypeScript definitions were found in the provided documentation.

## 13. Common Troubleshooting Issues and Solutions

The documentation does not have a dedicated troubleshooting section. However, common issues might involve incorrect initialization or conflicting CSS classes. Ensure that the `HSOverlay` script is loaded and that you are targeting the correct elements.

## 14. Styling Customization Options

Styling is done primarily through Tailwind CSS utility classes. You can customize the modal's appearance by modifying these classes in your HTML or by extending your Tailwind CSS configuration.