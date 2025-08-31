# Flex - FlyonUI

## Flexbox Utility

FlyonUI leverages Tailwind CSS's powerful flexbox utilities to create flexible and responsive layouts. "Flex" is not a component in itself, but a set of utility classes that you can apply to any container element.

## Common Flexbox Classes
Here are some of the most commonly used flexbox utility classes in FlyonUI components:

- **Display:**
  - `flex`: Creates a flex container.
  - `inline-flex`: Creates an inline-flex container.
- **Direction:**
  - `flex-row`: (Default) Arranges items horizontally.
  - `flex-row-reverse`: Reverses the horizontal order.
  - `flex-col`: Arranges items vertically.
  - `flex-col-reverse`: Reverses the vertical order.
- **Wrapping:**
  - `flex-wrap`: Allows items to wrap to the next line.
  - `flex-nowrap`: Prevents wrapping.
- **Alignment:**
  - `items-start`, `items-center`, `items-end`, `items-stretch`: Align items along the cross-axis.
  - `justify-start`, `justify-center`, `justify-end`, `justify-between`, `justify-around`: Justify content along the main-axis.
- **Gap:**
  - `gap-1`, `gap-2`, `gap-4`: Adds space between flex items.
  - `gap-x-2`, `gap-y-4`: Adds space on a specific axis.

## HTML Examples

### Basic Flex Row
```html
<div class="flex flex-row gap-4">
  <div class="bg-primary/20 p-4 rounded-box">Item 1</div>
  <div class="bg-primary/20 p-4 rounded-box">Item 2</div>
  <div class="bg-primary/20 p-4 rounded-box">Item 3</div>
</div>
```

### Responsive Flex Layout
This example shows a column layout on small screens that becomes a row layout on larger screens.
```html
<div class="flex flex-col sm:flex-row gap-4">
  <div class="bg-secondary/20 p-4 rounded-box">Responsive Item 1</div>
  <div class="bg-secondary/20 p-4 rounded-box">Responsive Item 2</div>
</div>
```

### Centering Items
```html
<div class="flex justify-center items-center h-32 bg-base-200 rounded-box">
  <p>Centered Content</p>
</div>
```

## JavaScript Interaction API
Flexbox utilities are purely CSS-based and do not have a JavaScript API.

## Configuration Options
All configuration is done through Tailwind CSS utility classes in your HTML.