# Badge - FlyonUI

## Complete List of Classes
- **Base:**
  - `badge`: The main class for the badge component.
- **Color Variants:**
  - `badge-primary`
  - `badge-secondary`
  - `badge-success`
  - `badge-error`
  - `badge-warning`
  - `badge-info`
- **Style Variants:**
  - `badge-soft`: For a softer color scheme.
  - `badge-outline`: For an outlined style.
- **Size Variants:**
  - `badge-xl`
  - `badge-lg`
  - `badge-sm`
  - `badge-xs`
- **As an Indicator:**
  - `indicator-item`: Used when a badge is an indicator on another element.
- **Dismissible:**
  - `removing:translate-x-5`
  - `removing:opacity-0`

## Variations and Sizes
FlyonUI badges come in several color, style, and size variations.

### Color Variants
- `badge-primary`
- `badge-secondary`
- `badge-success`
- `badge-error`
- `badge-warning`
- `badge-info`

### Style Variants
- **Solid (default):**
  ```html
  <span class="badge badge-primary">Primary</span>
  ```
- **Soft:**
  ```html
  <span class="badge badge-soft badge-primary">Primary</span>
  ```
- **Outline:**
  ```html
  <span class="badge badge-outline badge-primary">Primary</span>
  ```

### Size Variants
```html
<span class="badge badge-secondary badge-xl">XL</span>
<span class="badge badge-secondary badge-lg">LG</span>
<span class="badge badge-secondary">BASE</span>
<span class="badge badge-secondary badge-sm">SM</span>
<span class="badge badge-secondary badge-xs">XS</span>
```

## HTML Examples
### Basic Badges
```html
<span class="badge badge-primary">Primary</span>
<a href="#"><span class="badge badge-secondary">Secondary</span></a>
<span class="badge badge-success">Success</span>
<span class="badge badge-error">Error</span>
<span class="badge badge-warning">Warning</span>
<span class="badge badge-info">Info</span>
```

### Badges with Icons
```html
<span class="badge badge-primary size-6 p-0">
  <span class="icon-[tabler--star]"></span>
</span>
<span class="badge badge-soft badge-primary badge-lg">
  <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-2.png" alt="Anna" class="size-4.5 rounded-full"/>
  Anna
</span>
```

### Dismissible Badge
```html
<span class="badge badge-primary badge-lg removing:translate-x-5 removing:opacity-0 transition duration-300 ease-in-out" id="badge-1">
  Badge
  <button class="icon-[tabler--circle-x-filled] size-5 min-h-0 cursor-pointer px-0" data-remove-element="#badge-1" aria-label="Dismiss Button"></button>
</span>
```

### Badge as an Indicator
```html
<div class="indicator">
  <span class="indicator-item badge badge-primary">+999</span>
  <div class="bg-primary/10 border-primary grid place-items-center rounded-md border p-3">
    <span class="icon-[tabler--bell] text-primary size-5"></span>
  </div>
</div>
```

## JavaScript Interaction API
The main JavaScript interaction for badges is the dismiss functionality, which is handled via a `data-remove-element` attribute. No other specific JavaScript API is needed for the badge component itself.

### Dismissing a Badge
To make a badge dismissible, add an ID to the badge and a button with the `data-remove-element` attribute pointing to that ID.

**HTML:**
```html
<span class="badge badge-soft badge-lg badge-primary removing:translate-x-5 removing:opacity-0 transition duration-300 ease-in-out" id="badge-chip">
  Badge
  <button class="icon-[tabler--circle-x-filled] size-5 min-h-0 cursor-pointer px-0 opacity-70" data-remove-element="#badge-chip" aria-label="Close Button"></button>
</span>
```
If you are using the FlyonUI script that handles `data-remove-element`, no additional JavaScript is required.

## Configuration Options
- `--badge-color`: CSS variable to customize the color of the badge.
- `--size`: CSS variable to customize the size of the badge.
- `data-remove-element`: Attribute to specify the target element to remove when the button is clicked.

### Example of using CSS variables:
```css
:root {
  --badge-bg: #your-custom-bg-color;
  --badge-text: #your-custom-text-color;
  --badge-color: #your-custom-color;
  --size: 1.5rem;
}