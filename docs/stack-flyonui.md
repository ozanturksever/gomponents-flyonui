# Stack - FlyonUI

## Complete List of Classes
- **Base:**
  - `stack`: The main class for the stack component. It allows items to be stacked on top of each other.

The stack component itself does not have many variants. Its power comes from the components placed inside it.

## Variations and Sizes
The `stack` component does not have size or color variants. The appearance is determined by the children within the stack. Sizing is controlled by applying Tailwind CSS width and height utilities to the `.stack` container.

## HTML Examples

### Basic Stack
This example shows how to stack multiple `div` elements. The last element will be on top.
```html
<div class="stack w-48 h-24">
  <div class="grid w-32 h-20 rounded bg-primary text-primary-content place-content-center">1</div>
  <div class="grid w-32 h-20 rounded bg-accent text-accent-content place-content-center">2</div>
  <div class="grid w-32 h-20 rounded bg-secondary text-secondary-content place-content-center">3</div>
</div>
```

### Stacked Cards
```html
<div class="stack">
  <div class="card shadow-md bg-primary text-primary-content">
    <div class="card-body">
      <h2 class="card-title">Card 1</h2>
      <p>This is the first card.</p>
    </div>
  </div>
  <div class="card shadow-md bg-secondary text-secondary-content">
    <div class="card-body">
      <h2 class="card-title">Card 2</h2>
      <p>This is the second card, on top.</p>
    </div>
  </div>
</div>
```

### Stacked Images
```html
<div class="stack">
  <img src="https://daisyui.com/images/stock/photo-1559703248-dca8b94400de.jpg" alt="Image 1" class="rounded w-32"/>
  <img src="https://daisyui.com/images/stock/photo-1565098772267-60af42b81ef2.jpg" alt="Image 2" class="rounded w-32"/>
  <img src="https://daisyui.com/images/stock/photo-1572635148818-ef6fd45eb394.jpg" alt="Image 3" class="rounded w-32"/>
</div>
```

## JavaScript Interaction API
The stack component is a CSS-only component and does not have a JavaScript API.

## Configuration Options
There are no specific configuration options for the stack component. Customization is achieved through Tailwind CSS utility classes and the order of the child elements.