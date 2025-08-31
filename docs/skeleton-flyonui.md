# Skeleton - FlyonUI

## Complete List of Classes
- **Base:**
  - `skeleton`: The main class for a skeleton placeholder.
- **Shape Variants:**
  - `skeleton-rectangle`: (Default) A rectangular skeleton.
  - `skeleton-circle`: A circular skeleton.
  - `skeleton-text`: a skeleton designed to mimic a line of text.
- **Animation:**
  - `skeleton-pulse`: Adds a pulse animation.
  - `skeleton-wave`: Adds a wave animation.

## Variations and Sizes
Sizing of skeleton components is controlled by standard Tailwind CSS height and width utilities (e.g., `h-4`, `w-full`).

### Shape Variants
```html
<!-- Rectangle -->
<div class="skeleton h-32 w-full"></div>

<!-- Circle -->
<div class="skeleton w-16 h-16 rounded-full shrink-0"></div>

<!-- Text Lines -->
<div class="flex flex-col gap-4 w-52">
  <div class="skeleton h-4 w-28"></div>
  <div class="skeleton h-4 w-full"></div>
  <div class="skeleton h-4 w-full"></div>
</div>
```

### Animation Variants
```html
<!-- Pulse Animation -->
<div class="skeleton skeleton-pulse h-32 w-full"></div>

<!-- Wave Animation -->
<div class="skeleton skeleton-wave h-32 w-full"></div>
```

## HTML Examples

### Skeleton for a Card
```html
<div class="flex flex-col gap-4 w-52">
  <div class="skeleton h-32 w-full"></div>
  <div class="skeleton h-4 w-28"></div>
  <div class="skeleton h-4 w-full"></div>
  <div class="skeleton h-4 w-full"></div>
</div>
```

### Skeleton for a Profile
```html
<div class="flex gap-4 items-center">
  <div class="skeleton w-16 h-16 rounded-full shrink-0"></div>
  <div class="flex flex-col gap-4">
    <div class="skeleton h-4 w-20"></div>
    <div class="skeleton h-4 w-28"></div>
  </div>
</div>
```

## JavaScript Interaction API
The skeleton component is a CSS-only component used for loading states. It does not have a JavaScript API. You would use JavaScript to conditionally render the skeleton component while data is being fetched, and then replace it with the actual content once loaded.

## Configuration Options
All customization is done through utility classes. There are no specific data attributes or CSS variables for the skeleton component.