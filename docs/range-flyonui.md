# Range - FlyonUI

## Complete List of Classes
- **Base:**
  - `range`: The main class for the range slider component.
- **Color Variants:**
  - `range-primary`
  - `range-secondary`
  - `range-accent`
  - `range-success`
  - `range-warning`
  - `range-info`
  - `range-error`
- **Size Variants:**
  - `range-lg`
  - `range-md` (default)
  - `range-sm`
  - `range-xs`

## Variations and Sizes

### Color Variants
```html
<input type="range" min="0" max="100" class="range range-primary" />
<input type="range" min="0" max="100" class="range range-success" />
```

### Size Variants
```html
<input type="range" min="0" max="100" class="range range-xs" />
<input type="range" min="0" max="100" class="range range-sm" />
<input type="range" min="0" max="100" class="range range-md" />
<input type="range" min="0" max="100" class="range range-lg" />
```

## HTML Examples

### Basic Range Slider
```html
<input type="range" min="0" max="100" value="40" class="range" />
```

### Range with Steps
The `step` attribute allows you to specify the legal number intervals.
```html
<input type="range" min="0" max="100" value="25" class="range" step="25" />
<div class="w-full flex justify-between text-xs px-2">
  <span>0</span>
  <span>25</span>
  <span>50</span>
  <span>75</span>
  <span>100</span>
</div>
```

### Disabled Range Slider
```html
<input type="range" min="0" max="100" value="50" class="range" disabled />
```

## JavaScript Interaction API
The range slider is a standard HTML `<input type="range">` element. You can interact with it using standard DOM methods to get or set its `value`, `min`, `max`, and `step` attributes.

```javascript
const rangeSlider = document.getElementById('my-range');

// Get the current value
const currentValue = rangeSlider.value;

// Set the value
rangeSlider.value = 75;

// Add an event listener
rangeSlider.addEventListener('input', (event) => {
  console.log('Range value:', event.target.value);
});
```
FlyonUI also appears to have an advanced range slider with more features, likely implemented as a third-party plugin (`HSAdvancedRangeSlider`).

## Configuration Options
Configuration is done via standard HTML attributes on the `<input>` element:
- `min`: The minimum value of the range.
- `max`: The maximum value of the range.
- `step`: The interval to increment or decrement the value.
- `value`: The initial value.
- `disabled`: Disables the range slider.