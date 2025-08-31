# Spinner - FlyonUI

The Spinner is a type of loading indicator in FlyonUI. It's used to signify that a process is ongoing. It is essentially the same as the `loading` component with the `loading-spinner` class applied.

## Complete List of Classes
- **Base:**
  - `loading`: The base class for all loading indicators.
  - `loading-spinner`: The class that creates the spinner style.
- **Color Variants:**
  - `text-primary`, `text-secondary`, etc. (uses text color utilities).
- **Size Variants:**
  - `loading-lg`
  - `loading-md` (default)
  - `loading-sm`
  - `loading-xs`

## Variations and Sizes

### Color Variants
The spinner's color is controlled by Tailwind's text color utilities.
```html
<span class="loading loading-spinner text-primary"></span>
<span class="loading loading-spinner text-success"></span>
```

### Size Variants
```html
<span class="loading loading-spinner loading-xs"></span>
<span class="loading loading-spinner loading-sm"></span>
<span class="loading loading-spinner loading-md"></span>
<span class="loading loading-spinner loading-lg"></span>
```

## HTML Examples

### Simple Spinner
```html
<span class="loading loading-spinner"></span>
```

### Spinner inside a Card Overlay
This is a common use case, showing a spinner while card content is loading.
```html
<div class="card group max-w-sm hover:shadow">
  <!-- Card content -->
  <div class="bg-base-100/50 absolute start-0 top-0 size-full"></div>
  <div class="absolute start-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 transform">
    <span class="loading loading-spinner loading-lg text-primary"></span>
  </div>
</div>
```

## JavaScript Interaction API
The spinner is a CSS-only component. It does not have its own JavaScript API. You would typically use JavaScript to dynamically add or remove the spinner element from the DOM to show or hide the loading state.

## Configuration Options
There are no specific configuration options for the spinner. All customization is done via utility classes for color, size, etc.