# Rating - FlyonUI

## Complete List of Classes
- **Base:**
  - `rating`: The main class for the rating component.
- **Color Variants:**
  - `rating-primary`
  - `rating-secondary`
  - `rating-accent`
  - `rating-success`
  - `rating-warning`
  - `rating-info`
  - `rating-error`
- **Size Variants:**
  - `rating-lg`
  - `rating-md` (default)
  - `rating-sm`
  - `rating-xs`

## Variations and Sizes
The Rating component can be styled with different colors and sizes.

### Color Variants
```html
<div class="rating rating-primary">...</div>
<div class="rating rating-secondary">...</div>
```

### Size Variants
```html
<div class="rating rating-xs">...</div>
<div class="rating rating-sm">...</div>
<div class="rating rating-md">...</div>
<div class="rating rating-lg">...</div>
```

## HTML Examples

### Basic Rating
The rating component uses radio buttons internally to select a value. The `checked` attribute determines the selected rating.
```html
<div class="rating">
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" checked />
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" />
  <input type="radio" name="rating-1" class="mask mask-star" />
</div>
```

### Star-2 mask
```html
<div class="rating">
  <input type="radio" name="rating-2" class="mask mask-star-2 bg-orange-400" />
  <input type="radio" name="rating-2" class="mask mask-star-2 bg-orange-400" checked />
  <input type="radio" name="rating-2" class="mask mask-star-2 bg-orange-400" />
  <input type="radio" name="rating-2" class="mask mask-star-2 bg-orange-400" />
  <input type="radio" name="rating-2" class="mask mask-star-2 bg-orange-400" />
</div>
```

### Heart mask with multiple colors
```html
<div class="rating gap-1">
  <input type="radio" name="rating-3" class="mask mask-heart bg-red-400" />
  <input type="radio" name="rating-3" class="mask mask-heart bg-orange-400" checked />
  <input type="radio" name="rating-3" class="mask mask-heart bg-yellow-400" />
  <input type="radio" name="rating-3" class="mask mask-heart bg-lime-400" />
  <input type="radio" name="rating-3" class="mask mask-heart bg-green-400" />
</div>
```

## JavaScript Interaction API
The rating component is primarily CSS-driven. You can interact with the underlying radio buttons using standard DOM JavaScript to get or set the rating value. There is a `HSRating` class for more advanced uses.

## Configuration Options
- **CSS Variables:** The rating component does not expose specific CSS variables for customization. Styling is done through utility classes.
- The `name` attribute on the `input` elements is crucial for grouping the radio buttons.