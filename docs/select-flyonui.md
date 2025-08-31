# Select - FlyonUI

## Complete List of Classes
- **Base:**
  - `select`: The main class for the select component.
- **Color Variants:**
  - `select-primary`
  - `select-secondary`
  - `select-accent`
  - `select-success`
  - `select-warning`
  - `select-info`
  - `select-error`
- **Size Variants:**
  - `select-lg`
  - `select-md` (default)
  - `select-sm`
  - `select-xs`
- **Other:**
  - `select-bordered`: Adds a border.
  - `select-ghost`: For a transparent select.
  - `select-floating` and `select-floating-label`: For a floating label effect.
  - `label-text`: For styling the text of a `<label>`.
  - `helper-text`: For styling helper text below a select.
  - `is-valid`: For success validation state.
  - `is-invalid`: For error validation state.

## Variations and Sizes

### Color Variants
```html
<select class="select select-primary w-full max-w-xs">
  ...
</select>
```

### Size Variants
```html
<select class="select select-bordered select-lg w-full max-w-xs">...</select>
<select class="select select-bordered select-md w-full max-w-xs">...</select>
<select class="select select-bordered select-sm w-full max-w-xs">...</select>
<select class="select select-bordered select-xs w-full max-w-xs">...</select>
```

## HTML Examples

### Basic Select
```html
<select class="select w-full max-w-xs">
  <option disabled selected>Pick your favorite Simpson</option>
  <option>Homer</option>
  <option>Marge</option>
  <option>Bart</option>
  <option>Lisa</option>
  <option>Maggie</option>
</select>
```

### Select with Label and Helper Text
```html
<div class="w-96">
  <label class="label-text" for="selectHelperText"> Pick your favorite Movie </label>
  <select class="select" id="selectHelperText">
    <option>The Godfather</option>
    <option>The Shawshank Redemption</option>
  </select>
  <span class="helper-text">Please select one.</span>
</div>
```

### Floating Label Select
```html
<div class="select-floating w-96">
  <select class="select" aria-label="Select floating label" id="selectFloating">
    <option>The Godfather</option>
    <option>The Shawshank Redemption</option>
  </select>
  <label class="select-floating-label" for="selectFloating">Pick your favorite Movie</label>
</div>
```

### Disabled Select
```html
<select class="select w-full max-w-xs" disabled>
  <option>You can't touch this</option>
</select>
```

## JavaScript Interaction API
The select component is a standard HTML `<select>` element. You can interact with it using standard DOM methods. For more advanced "select" or "dropdown" controls with search, see the `Combobox` or `Dropdown` components. FlyonUI also seems to have an `HSSelect` class for advanced usage.

## Configuration Options
- **HTML Attributes:**
  - `disabled`: Disables the select input.
  - `multiple`: Allows multiple options to be selected.
  - `required`: Specifies that an option must be selected.
  - `size`: Controls the number of visible options in a scrolling list.