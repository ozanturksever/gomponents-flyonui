# Textarea - FlyonUI

## Complete List of Classes
- **Base:**
  - `textarea`: The main class for the textarea component.
- **Color Variants:**
  - `textarea-primary`
  - `textarea-secondary`
  - `textarea-accent`
  - `textarea-success`
  - `textarea-warning`
  - `textarea-info`
  - `textarea-error`
- **Size Variants:**
  - `textarea-lg`
  - `textarea-md` (default)
  - `textarea-sm`
  - `textarea-xs`
- **Other:**
  - `textarea-bordered`: Adds a border.
  - `textarea-ghost`: For a transparent textarea.
  - `textarea-floating` and `textarea-floating-label`: For a floating label effect.
  - `is-valid`: For success validation state.
  - `is-invalid`: For error validation state.
  - `label-text`: For styling the text of a `<label>`.
  - `helper-text`: For styling helper text below a textarea.

## Variations and Sizes

### Color Variants
```html
<textarea class="textarea textarea-bordered textarea-primary"></textarea>
<textarea class="textarea textarea-bordered textarea-success"></textarea>
```

### Size Variants
```html
<textarea class="textarea textarea-bordered textarea-lg"></textarea>
<textarea class="textarea textarea-bordered textarea-md"></textarea>
<textarea class="textarea textarea-bordered textarea-sm"></textarea>
<textarea class="textarea textarea-bordered textarea-xs"></textarea>
```

## HTML Examples

### Basic Textarea with Label
```html
<div>
  <label class="label-text" for="bio">Bio</label>
  <textarea class="textarea" id="bio" placeholder="Tell us about yourself..."></textarea>
</div>
```

### Floating Label Textarea
```html
<div class="textarea-floating sm:w-96">
  <textarea class="textarea" placeholder="Hello!!!" id="textareaStateSuccessFloating"></textarea>
  <label class="textarea-floating-label" for="textareaStateSuccessFloating">Your bio</label>
</div>
```

### Disabled Textarea
```html
<textarea class="textarea" placeholder="You can't write here" disabled></textarea>
```

## JavaScript Interaction API
The textarea is a standard HTML `<textarea>` element. You interact with it using standard DOM methods to get or set its `value`.

```javascript
const myTextarea = document.getElementById('my-textarea');

// Get the value
const content = myTextarea.value;

// Set the value
myTextarea.value = 'New content for the textarea.';
```

## Configuration Options
Configuration is done via standard HTML attributes:
- `placeholder`: The placeholder text.
- `rows`: The visible number of lines in a text area.
- `cols`: The visible width of a text area.
- `disabled`: Disables the textarea.
- `readonly`: Makes the textarea read-only.
- `required`: Specifies that the textarea must be filled out.