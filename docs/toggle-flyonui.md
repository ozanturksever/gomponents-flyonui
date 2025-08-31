# Toggle - FlyonUI

The Toggle component is a checkbox styled to look like a switch.

## Complete List of Classes
- **Base:**
  - `toggle`: The main class for the toggle switch.
- **Color Variants:**
  - `toggle-primary`
  - `toggle-secondary`
  - `toggle-accent`
  - `toggle-success`
  - `toggle-warning`
  - `toggle-info`
  - `toggle-error`
- **Size Variants:**
  - `toggle-lg`
  - `toggle-md` (default)
  - `toggle-sm`
  - `toggle-xs`
- **Other:**
  - `label-text`: For styling the text of a `<label>`.

## Variations and Sizes

### Color Variants
```html
<input type="checkbox" class="toggle toggle-primary" />
<input type="checkbox" class="toggle toggle-success" />
```

### Size Variants
```html
<input type="checkbox" class="toggle toggle-xs" />
<input type="checkbox" class="toggle toggle-sm" />
<input type="checkbox" class="toggle toggle-md" />
<input type="checkbox" class="toggle toggle-lg" />
```

## HTML Examples

### Basic Toggle
```html
<input type="checkbox" class="toggle" checked />
```

### Toggle with Label
```html
<div class="form-control">
  <label class="label cursor-pointer">
    <span class="label-text">Remember me</span>
    <input type="checkbox" class="toggle" checked />
  </label>
</div>
```

### Disabled Toggle
```html
<input type="checkbox" class="toggle" disabled />
<input type="checkbox" class="toggle" checked disabled />
```

## JavaScript Interaction API
The toggle is a styled HTML `<input type="checkbox">` element. You interact with it using standard DOM methods to get or set its `checked` state.

```javascript
const myToggle = document.getElementById('my-toggle');

// Check the state
const isChecked = myToggle.checked;

// Set the state
myToggle.checked = true;
```
For more complex interactions, like toggling password visibility, FlyonUI provides data attributes (e.g., `data-toggle-password`) that are handled by its JavaScript library.

## Configuration Options
Configuration is done via standard HTML attributes:
- `checked`: Pre-selects the toggle to the "on" state.
- `disabled`: Disables the toggle.
- `name`: Sets the name for the input, useful in forms.
- `value`: Sets the value to be submitted with the form.