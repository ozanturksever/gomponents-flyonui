# Radio - FlyonUI

## Complete List of Classes
- **Base:**
  - `radio`: The main class for the radio input component.
- **Color Variants:**
  - `radio-primary`
  - `radio-secondary`
  - `radio-accent`
  - `radio-success`
  - `radio-warning`
  - `radio-info`
  - `radio-error`
- **Size Variants:**
  - `radio-lg`
  - `radio-md` (default)
  - `radio-sm`
  - `radio-xs`
- **Other:**
  - `label-text`: For styling the text of a `<label>`.

## Variations and Sizes

### Color Variants
```html
<input type="radio" name="radio-color" class="radio radio-primary" />
<input type="radio" name="radio-color" class="radio radio-secondary" />
```

### Size Variants
```html
<input type="radio" name="radio-size" class="radio radio-xs" />
<input type="radio" name="radio-size" class="radio radio-sm" />
<input type="radio" name="radio-size" class="radio radio-md" />
<input type="radio" name="radio-size" class="radio radio-lg" />
```

## HTML Examples

### Basic Radio Buttons
A group of radio buttons must share the same `name` attribute to ensure only one can be selected.
```html
<div class="flex items-center gap-2">
  <input type="radio" id="male" name="gender" class="radio radio-primary" checked />
  <label class="label-text text-base" for="male">Male</label>
</div>
<div class="flex items-center gap-2">
  <input type="radio" id="female" name="gender" class="radio radio-primary" />
  <label class="label-text text-base" for="female">Female</label>
</div>
```

### Vertical Radio Group
```html
<ul class="border-base-content/25 divide-base-content/25 divide-y max-w-sm rounded-md border *:cursor-pointer">
  <li>
    <label class="flex items-center gap-3 p-3">
      <input type="radio" name="expertise" class="radio radio-primary" />
      <span class="label-text text-base">Project Management</span>
    </label>
  </li>
  <li>
    <label class="flex items-center gap-3 p-3">
      <input type="radio" name="expertise" class="radio radio-primary" checked />
      <span class="label-text text-base">Marketing Strategy</span>
    </label>
  </li>
</ul>
```

### Disabled Radio Button
```html
<input type="radio" name="radio-disabled" class="radio" disabled />
<input type="radio" name="radio-disabled" class="radio" checked disabled />
```

## JavaScript Interaction API
Radio buttons are standard HTML elements. You interact with them using standard DOM methods to get or set their `checked` state, `value`, etc. There is no specific JavaScript API for the radio component in FlyonUI.

## Configuration Options
Configuration is done via standard HTML attributes:
- `name`: (Required) Groups radio buttons together.
- `value`: The value to be submitted for the selected radio button.
- `checked`: Pre-selects a radio button.
- `disabled`: Disables a radio button.