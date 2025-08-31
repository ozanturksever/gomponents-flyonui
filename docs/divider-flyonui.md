# Divider - FlyonUI

## Complete List of Classes
- **Base:**
  - `divider`: The main class for the divider component.
- **Direction:**
  - `divider-horizontal`: Creates a horizontal divider.
  - `divider-vertical`: (Implied) a standard divider is vertical in a horizontal flex container.
- **Color Variants:**
  - `divider-neutral`
  - `divider-primary`
  - `divider-secondary`
  - `divider-accent`
  - `divider-success`
  - `divider-warning`
  - `divider-info`
  - `divider-error`
- **Style Variants:**
  - `divider-dotted`: For a dotted line style.
  - `divider-dashed`: For a dashed line style.
- **Content Position:**
  - `divider-start`: Aligns content to the start.
  - `divider-end`: Aligns content to the end.

## Variations and Sizes
FlyonUI dividers can be styled with colors, styles, and can contain text or icons. Their thickness can be customized using pseudo-elements and border utilities.

### Color Variants
```html
<div class="divider divider-primary"></div>
<div class="divider divider-secondary"></div>
```

### Style Variants
```html
<div class="divider divider-dotted">Dotted</div>
<div class="divider divider-dashed">Dashed</div>
```

### Thickness
You can customize the thickness using `before` and `after` pseudo-elements with Tailwind's border utilities.
```html
<div class="divider after:border-t-4 before:border-t-4">Thick</div>
```

## HTML Examples

### Default Divider
```html
<div class="flex flex-col w-full">
  <div>Content 1</div>
  <div class="divider">OR</div>
  <div>Content 2</div>
</div>
```

### Horizontal Divider
```html
<div class="flex w-full">
  <div>Item 1</div>
  <div class="divider divider-horizontal"></div>
  <div>Item 2</div>
</div>
```

### Divider with Icon
```html
<div class="divider">
  <span class="flex items-center justify-center">
    <span class="icon-[tabler--crown] size-5"></span>
  </span>
</div>
```

### Responsive Divider
This example shows a divider that is vertical by default and becomes horizontal on large screens.
```html
<div class="flex w-full flex-col lg:flex-row">
  <div class="grid flex-grow h-32 card bg-base-300 rounded-box place-items-center">content</div>
  <div class="divider lg:divider-horizontal">OR</div>
  <div class="grid flex-grow h-32 card bg-base-300 rounded-box place-items-center">content</div>
</div>
```

## JavaScript Interaction API
The divider is a purely stylistic component and does not have a JavaScript API.

## Configuration Options
There are no specific configuration options for the divider component. Customization is achieved through utility classes.