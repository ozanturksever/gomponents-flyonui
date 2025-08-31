# Input - FlyonUI

## Complete List of Classes
- **Base:**
  - `input`: The main class for a text input field.
- **Color Variants:**
  - `input-primary`
  - `input-secondary`
  - `input-accent`
  - `input-success`
  - `input-warning`
  - `input-info`
  - `input-error`
- **Size Variants:**
  - `input-lg`
  - `input-md` (default)
  - `input-sm`
  - `input-xs`
- **Other:**
  - `input-bordered`: Adds a border.
  - `input-ghost`: For a transparent input.
  - `join-item`: For grouping inputs.
  - `is-valid`: For success validation state.
  - `is-invalid`: For error validation state.
  - `input-floating` and `input-floating-label`: For floating label effect.
  - `pin-input`: For single-character inputs in a group (e.g., OTP).
  - `pin-input-underline`: A style variation for pin inputs.

## Variations and Sizes

### Size Variants
```html
<input type="text" placeholder="Extra Small" class="input input-xs" />
<input type="text" placeholder="Small" class="input input-sm" />
<input type="text" placeholder="Normal" class="input input-md" />
<input type="text" placeholder="Large" class="input input-lg" />
```
### Color Variants
```html
<input type="text" placeholder="Primary" class="input input-bordered input-primary" />
<input type="text" placeholder="Success" class="input input-bordered input-success" />
```

## HTML Examples

### Basic Input with Label
```html
<div>
  <label class="label-text" for="fullName">Full Name</label>
  <input type="text" placeholder="John Doe" class="input" id="fullName" />
</div>
```

### Floating Label Input
```html
<div class="input-floating w-96">
  <input type="text" placeholder="John Doe" class="input" id="floatingInput" />
  <label class="input-floating-label" for="floatingInput">Full Name</label>
</div>
```

### Input with Button (Join)
```html
<div class="join max-w-sm">
  <input class="input join-item" placeholder="Search"/>
  <button class="btn btn-outline btn-secondary join-item">Search</button>
</div>
```

### Pin Input
```html
<div class="flex space-x-3" data-pin-input>
  <input type="text" class="pin-input" placeholder="○" data-pin-input-item />
  <input type="text" class="pin-input" placeholder="○" data-pin-input-item />
  <input type="text" class="pin-input" placeholder="○" data-pin-input-item />
  <input type="text" class="pin-input" placeholder="○" data-pin-input-item />
</div>
```

### Input Number
```html
<div class="input max-w-sm" data-input-number>
  <input type="text" value="1" aria-label="Input number" data-input-number-input />
  <span class="my-auto flex gap-3">
    <button type="button" class="btn btn-primary btn-soft size-5.5 min-h-0 rounded-sm p-0" aria-label="Decrement button" data-input-number-decrement>
      <span class="icon-[tabler--minus] size-3.5 shrink-0"></span>
    </button>
    <button type="button" class="btn btn-primary btn-soft size-5.5 min-h-0 rounded-sm p-0" aria-label="Increment button" data-input-number-increment>
      <span class="icon-[tabler--plus] size-3.5 shrink-0"></span>
    </button>
  </span>
</div>
```

## JavaScript Interaction API
Standard text inputs are interacted with via the DOM. The advanced input types have their own JavaScript APIs.

- **Input Number:** `HSInputNumber`
- **Pin Input:** `HSPinInput`
- **Toggle Password:** Handled via `data-toggle-password` attribute.

### Example: Input Number Events
```javascript
const el = HSInputNumber.getInstance('#input-number');
el.on('change', ({inputValue}) => {console.log('Changed to:', inputValue)});
```

## Configuration Options
- **HTML Attributes:** Standard attributes like `placeholder`, `disabled`, `required`, `type`, `value`.
- **Data Attributes:**
  - `data-input-number`: Initializes the input number component. Can hold a JSON object for options like `min`, `max`, `step`.
  - `data-pin-input`: Initializes the pin input component.
  - `data-toggle-password`: Used on a button to toggle the visibility of a password field.
- **CSS Variables:**
  - `--input-color`: For color customization.
  - `--size`: For size customization.