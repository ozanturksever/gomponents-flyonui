# Grid - FlyonUI

## Grid Layout Utility

FlyonUI uses Tailwind CSS's powerful grid system to create complex and responsive layouts. "Grid" is not a single component, but a set of utility classes for building grid-based structures.

## Complete List of Classes
- **Display:**
  - `grid`: Creates a grid container.
- **Columns:**
  - `grid-cols-1`, `grid-cols-2`, ..., `grid-cols-12`: Defines the number of columns in the grid.
  - Responsive variants like `md:grid-cols-2`, `lg:grid-cols-3` are used to change the column count at different breakpoints.
- **Rows:**
  - `grid-rows-1`, `grid-rows-2`, etc.: Defines the number of rows.
  - `grid-flow-col`, `grid-flow-row`: Controls the direction of the grid flow.
- **Column Span:**
  - `col-span-1`, `col-span-2`, etc.: Makes an element span across multiple columns.
  - `col-start-2`, `col-end-4`: Specifies the start and end lines of an element.
- **Gap:**
  - `gap-4`, `gap-6`: Defines the space between grid cells.

## HTML Examples

### Basic Grid
A simple 3-column grid that becomes a single column on smaller screens.
```html
<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
  <div class="bg-primary/20 p-4 rounded-box">Item 1</div>
  <div class="bg-primary/20 p-4 rounded-box">Item 2</div>
  <div class="bg-primary/20 p-4 rounded-box">Item 3</div>
</div>
```

### Grid with Spanning Columns
This example shows a more complex layout where items can span multiple columns.
```html
<div class="grid grid-cols-3 gap-4">
  <div class="col-span-2 bg-secondary/20 p-4 rounded-box">Spans 2 columns</div>
  <div class="bg-secondary/20 p-4 rounded-box">Item 2</div>
  <div class="bg-secondary/20 p-4 rounded-box">Item 3</div>
  <div class="col-span-3 bg-secondary/20 p-4 rounded-box">Spans 3 columns</div>
</div>
```

### Bento Grid Example
FlyonUI provides "Bento Grid" blocks, which are complex layouts built using the grid utilities.
```html
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
  <div class="bg-blue-500 h-48 rounded-lg">Block 1</div>
  <div class="bg-green-500 h-32 rounded-lg">Block 2</div>
  <div class="bg-red-500 h-64 rounded-lg">Block 3</div>
  <div class="bg-yellow-500 h-40 rounded-lg">Block 4</div>
  <div class="bg-purple-500 h-56 rounded-lg">Block 5</div>
  <div class="bg-pink-500 h-48 rounded-lg">Block 6</div>
</div>
```

## JavaScript Interaction API
The grid system is purely CSS-based and has no JavaScript API.

## Configuration Options
All configuration is done through Tailwind CSS utility classes in your HTML.