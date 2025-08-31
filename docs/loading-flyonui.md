# Loading - FlyonUI

## Complete List of Classes
- **Base:**
  - `loading`: The main class for the loading component.
- **Type Variants:**
  - `loading-spinner`: A spinner-style loading indicator.
  - `loading-dots`: A dot-style loading indicator.
  - `loading-ring`: A ring-style loading indicator.
  - `loading-ball`: A ball-style loading indicator.
  - `loading-bars`: A bar-style loading indicator.
  - `loading-infinity`: An infinity-style loading indicator.
- **Size Variants:**
  - `loading-lg`
  - `loading-md` (default)
  - `loading-sm`
  - `loading-xs`
- **Animation:**
  - `animate-pulse`: Can be used on a badge to indicate loading.

## Variations and Sizes

### Type Variants
```html
<span class="loading loading-spinner"></span>
<span class="loading loading-dots"></span>
<span class="loading loading-ring"></span>
<span class="loading loading-ball"></span>
<span class="loading loading-bars"></span>
<span class="loading loading-infinity"></span>
```

### Size Variants
```html
<span class="loading loading-spinner loading-xs"></span>
<span class="loading loading-spinner loading-sm"></span>
<span class="loading loading-spinner loading-md"></span>
<span class="loading loading-spinner loading-lg"></span>
```

## HTML Examples

### Basic Spinners
```html
<span class="loading loading-spinner text-primary"></span>
<span class="loading loading-spinner text-secondary"></span>
<span class="loading loading-spinner text-accent"></span>
<span class="loading loading-spinner text-success"></span>
<span class="loading loading-spinner text-warning"></span>
<span class="loading loading-spinner text-info"></span>
<span class="loading loading-spinner text-error"></span>
```

### Loading overlay on a Card
This example shows a loading spinner centered on top of a card.
```html
<div class="card group max-w-sm hover:shadow">
  <figure>
    <img src="https://cdn.flyonui.com/fy-assets/components/carousel/image-7.png" alt="Album" />
  </figure>
  <div class="card-body">
    <h5 class="card-title">Card title</h5>
    <p>This is a wider card with supporting text below...</p>
  </div>
  <div class="bg-base-100/50 absolute start-0 top-0 size-full"></div>
  <div class="absolute start-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 transform">
    <span class="loading loading-spinner loading-lg text-primary"></span>
  </div>
</div>
```

### Pulsing Badge Indicator
```html
<div class="indicator">
  <span class="badge badge-primary animate-pulse rounded-full">loading...</span>
</div>
```

## JavaScript Interaction API
The loading component is a CSS-only component and does not have a JavaScript API. You would typically use JavaScript to dynamically add or remove the loading element from the DOM based on the application's state.

## Configuration Options
- Customization is done through Tailwind CSS utility classes (e.g., `text-primary` for color). There are no specific data attributes or CSS variables for the loading component itself.