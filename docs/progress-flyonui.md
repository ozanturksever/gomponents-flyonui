# Progress - FlyonUI

## Complete List of Classes
- **Base:**
  - `progress`: The main class for the progress bar component.
- **Color Variants:**
  - `progress-primary`
  - `progress-secondary`
  - `progress-accent`
  - `progress-success`
  - `progress-warning`
  - `progress-info`
  - `progress-error`
- **Size Variants:**
  - The height of the progress bar is controlled by Tailwind CSS height utilities (e.g., `h-2`), there are no specific `progress-*` size classes.

## Variations and Sizes

### Color Variants
```html
<progress class="progress progress-primary" value="30" max="100"></progress>
<progress class="progress progress-success" value="70" max="100"></progress>
```

### Sizes
Size is controlled with utility classes.
```html
<progress class="progress h-1" value="50" max="100"></progress>
<progress class="progress h-4" value="50" max="100"></progress>
```

## HTML Examples

### Basic Progress Bar
```html
<progress class="progress w-56" value="70" max="100"></progress>
```

### Indeterminate Progress Bar
An indeterminate progress bar is created by not setting the `value` attribute.
```html
<progress class="progress w-56"></progress>
```

### Colored Progress Bars
```html
<progress class="progress progress-primary w-56" value="10" max="100"></progress>
<progress class="progress progress-secondary w-56" value="40" max="100"></progress>
<progress class="progress progress-accent w-56" value="70" max="100"></progress>
```

## JavaScript Interaction API
The progress bar is a standard HTML `<progress>` element. You can interact with it using standard DOM methods to change the `value` and `max` attributes dynamically.

```javascript
const progressBar = document.getElementById('my-progress');

// Set progress
progressBar.value = 50;
```

## Configuration Options
- `value`: (Required for determinate) The current value of the progress bar.
- `max`: The maximum value of the progress bar (defaults to 100).
- If the `value` attribute is omitted, the progress bar is indeterminate.