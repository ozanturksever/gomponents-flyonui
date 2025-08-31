# Button - FlyonUI

## Complete List of Classes
- **Base:**
  - `btn`: The main class for the button component.
- **Color Variants:**
  - `btn-primary`
  - `btn-secondary`
  - `btn-accent`
  - `btn-success`
  - `btn-error`
  - `btn-warning`
  - `btn-info`
- **Style Variants:**
  - `btn-soft`: For a softer color scheme.
  - `btn-outline`: For an outlined style.
  - `btn-gradient`: For a gradient background.
  - `btn-text`: For a text-only button.
  - `glass`: For a frosted glass effect.
- **Shape Variants:**
  - `btn-square`: For a square button.
  - `btn-circle`: For a circular button.
- **Size Variants:**
  - `btn-lg`
  - `btn-sm`
  - `btn-xs`
- **Other:**
  - `join-item`: For grouping buttons with other elements.

## Variations and Sizes

### Color and Style Variants
- **Solid:** `btn btn-primary`
- **Outline:** `btn btn-outline btn-primary`
- **Soft:** `btn btn-soft btn-primary`
- **Gradient:** `btn btn-gradient btn-primary`
- **Text:** `btn btn-text btn-primary`
- **Glass:** `btn glass`

### Shape Variants
- **Default (Pill):** `btn`
- **Square:** `btn btn-square`
- **Circle:** `btn btn-circle`

### Size Variants
```html
<button class="btn btn-lg">Large</button>
<button class="btn">Normal</button>
<button class="btn btn-sm">Small</button>
<button class="btn btn-xs">Extra Small</button>
```

## HTML Examples
### Basic Buttons
```html
<button class="btn">Default</button>
<button class="btn btn-primary">Primary</button>
<button class="btn btn-secondary">Secondary</button>
<button class="btn btn-accent">Accent</button>
```

### Buttons with Icons
```html
<button class="btn btn-primary">
  <span class="icon-[tabler--thumb-up] size-5"></span>
  Like
</button>
<button class="btn btn-circle btn-primary" aria-label="Search">
  <span class="icon-[tabler--search] size-5"></span>
</button>
```

### Social Media Buttons
```html
<button class="btn [--btn-color:#1877F2] text-white">
  <span class="icon-[tabler--brand-facebook] size-4.5 shrink-0"></span> Facebook
</button>
<button class="btn btn-circle btn-soft [--btn-color:#1da1f2]" aria-label="Twitter">
  <span class="icon-[tabler--brand-x] size-4.5 shrink-0"></span>
</button>
```

### Input with Button
```html
<div class="join max-w-sm">
 <input class="input join-item" placeholder="Search" />
 <button class="btn btn-outline btn-secondary join-item">Search</button>
</div>
```

## JavaScript Interaction API
Buttons are standard HTML elements and do not have a specific JavaScript API in FlyonUI beyond standard DOM manipulation. However, they can be used to trigger JavaScript actions. For example, some components like modals or accordions can be destroyed and reinitialized using buttons.

### Example: Destroying and Reinitializing a Component
```javascript
// This is a generic example inspired by FlyonUI docs.
// Replace `HSComponent` with the actual component you are using.
window.addEventListener('load', () => {
    const destroyBtn = document.querySelector('#destroy-btn');
    const reinitBtn = document.querySelector('#reinit-btn');

    if (destroyBtn) {
        destroyBtn.addEventListener('click', () => {
          // Logic to destroy a component
          // e.g., HSComponent.getInstance(element, true).element.destroy();
          destroyBtn.setAttribute('disabled', 'disabled');
          reinitBtn.removeAttribute('disabled');
        });
    }

    if(reinitBtn) {
        reinitBtn.addEventListener('click', () => {
          // Logic to re-initialize a component
          // e.g., HSComponent.autoInit();
          reinitBtn.setAttribute('disabled', 'disabled');
          destroyBtn.removeAttribute('disabled');
        });
    }
});
```

## Configuration Options
- **CSS Variables:** FlyonUI buttons can be customized using CSS variables.
  - `--btn-color`: Sets the primary color of the button.
  - `--btn-fg`: Sets the foreground/text color.
  - `--btn-shadow`: Sets the box-shadow.
  - `--btn-noise`: Adds a noise effect.
  - `--btn-p`: Sets the padding.
  - `--size`: Sets the size.
  - `--dark-shade`: Sets the dark shade for gradient buttons.
- **Data Attributes:** Buttons can use `data-*` attributes to interact with other components, like `data-remove-element` or `data-collapse`.

### Example of using CSS variables:
```css
.btn-custom {
  --btn-color: purple;
  --btn-fg: white;
}