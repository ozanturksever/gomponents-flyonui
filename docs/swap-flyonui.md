# Swap - FlyonUI

## Complete List of Classes
- **Base:**
  - `swap`: The main container for the swap component.
  - `swap-on`: The element that is visible when the swap is in the "on" state.
  - `swap-off`: The element that is visible when the swap is in the "off" state.
  - `swap-input`: The hidden checkbox that controls the state.
- **Animation Variants:**
  - `swap-rotate`: Adds a rotation animation on state change.
  - `swap-flip`: Adds a flip animation on state change.
- **Color Variants:**
  - `swap-primary`, `swap-secondary`, `swap-accent`, etc., applied to the main `swap` element.

## Variations and Sizes
The Swap component is a single component with animation variations. Its size is determined by the content within the `swap-on` and `swap-off` elements.

### Animation Variations
- **Rotate:**
  ```html
  <label class="swap swap-rotate">...</label>
  ```
- **Flip:**
  ```html
  <label class="swap swap-flip">...</label>
  ```

## HTML Examples

### Basic Swap with Text
```html
<label class="swap">
  <input type="checkbox" class="swap-input" />
  <div class="swap-on">ON</div>
  <div class="swap-off">OFF</div>
</label>
```

### Swap with Icons (Hamburger Menu)
This is a common use case for the swap component.
```html
<label class="btn btn-circle swap swap-rotate">
  <input type="checkbox" class="swap-input" />
  <span class="icon-[tabler--menu-2] swap-off"></span>
  <span class="icon-[tabler--x] swap-on"></span>
</label>
```

### Swap with Flip Animation
```html
<label class="swap swap-flip text-9xl">
  <input type="checkbox" class="swap-input" />
  <div class="swap-on">😈</div>
  <div class="swap-off">😇</div>
</label>
```

## JavaScript Interaction API
The swap component is controlled by a standard HTML checkbox. You can programmatically control its state by changing the `checked` property of the hidden checkbox.

```javascript
// Get the checkbox element
const swapCheckbox = document.querySelector('#my-swap .swap-input');

// Toggle the swap state
swapCheckbox.checked = !swapCheckbox.checked;
```
There is no specific JavaScript API for the swap component in FlyonUI beyond standard DOM manipulation.

## Configuration Options
- **HTML Attributes:**
  - `checked`: on the `<input type="checkbox">` to set the initial state to "on".