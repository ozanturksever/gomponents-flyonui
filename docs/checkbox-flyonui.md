# Checkbox - FlyonUI

## Complete List of Classes
- **Base:**
  - `checkbox`: The main class for the checkbox input.
- **Color Variants:**
  - `checkbox-primary`
  - `checkbox-secondary`
  - `checkbox-accent`
  - `checkbox-success`
  - `checkbox-warning`
  - `checkbox-info`
  - `checkbox-error`
- **Size Variants:**
  - `checkbox-lg`
  - `checkbox-md` (default)
  - `checkbox-sm`
  - `checkbox-xs`
- **Other:**
  - `custom-option`: A label that acts as a container for a custom-styled checkbox.
  - `custom-soft-option`: A soft-styled version of `custom-option`.
  - `is-valid`: For displaying a valid state.
  - `is-invalid`: For displaying an invalid state.
  - `label-text`: For styling the label text.
  - `switch`: To style a checkbox as a switch.

## Variations and Sizes

### Colors
- `checkbox-primary`
- `checkbox-secondary`
- `checkbox-accent`
- etc.

### Sizes
```html
<input type="checkbox" class="checkbox checkbox-lg" />
<input type="checkbox" class="checkbox checkbox-md" />
<input type="checkbox" class="checkbox checkbox-sm" />
<input type="checkbox" class="checkbox checkbox-xs" />
```

### As a Switch
The `switch` class can be used to render a checkbox as a toggle switch.
```html
<input type="checkbox" class="switch" />
```

## HTML Examples

### Basic Checkbox
```html
<div class="flex items-center gap-1">
  <input type="checkbox" class="checkbox" id="defaultCheckbox1" />
  <label class="label-text text-base" for="defaultCheckbox1">Default</label>
</div>
<div class="flex items-center gap-1">
  <input type="checkbox" class="checkbox" id="defaultCheckbox2" checked />
  <label class="label-text text-base" for="defaultCheckbox2">Checked</label>
</div>
```

### Checkbox with Label and Helper Text
```html
<div class="flex gap-2">
  <input type="checkbox" class="checkbox checkbox-primary mt-2" id="checkboxLabel" />
  <label class="label-text cursor-pointer flex flex-col" for="checkboxLabel">
    <span class="text-base">Archive</span>
    <span>Notify me when this action happens.</span>
  </label>
</div>
```

### Checkbox List
```html
<ul class="border-base-content/25 divide-base-content/25 divide-y max-w-sm rounded-md border *:cursor-pointer">
  <li>
    <label class="flex items-center gap-3 p-3">
      <input type="checkbox" class="checkbox checkbox-primary" />
      <span class="label-text text-base"> Web Development </span>
    </label>
  </li>
  <li>
    <label class="flex items-center gap-3 p-3">
      <input type="checkbox" class="checkbox checkbox-primary" checked />
      <span class="label-text text-base"> Data Analysis </span>
    </label>
  </li>
</ul>
```

### Custom Option Checkbox
```html
<label class="custom-option flex flex-row items-start gap-3 sm:w-1/2">
  <input type="checkbox" class="checkbox checkbox-primary mt-2" checked required />
  <span class="label-text w-full text-start">
    <span class="flex justify-between mb-1">
      <span class="text-base font-medium">Basic</span>
      <span class="text-base-content/50 text-base">Free</span>
    </span>
    <span class="text-base-content/80">Get 1 project with 1 teams members.</span>
  </span>
</label>
```

## JavaScript Interaction API
Checkboxes are standard HTML elements and do not have a specific JavaScript API in FlyonUI. You can interact with them using standard DOM methods. For example, you can use JavaScript to toggle the visibility of a password field.

### Example: Toggle Password Visibility
```html
<div class="mb-3 max-w-sm">
  <label class="label-text" for="toggle-password-checkbox">Password</label>
  <input id="toggle-password-checkbox" type="password" class="input" placeholder="Enter password" value="Pwd_1242@mA1" />
</div>

<div class="flex items-center gap-2">
  <input id="toggleCheckboxPassword" type="checkbox" data-toggle-password='{ "target": "#toggle-password-checkbox" }' class="checkbox checkbox-primary" />
  <label class="label-text text-base" for="toggleCheckboxPassword">Show password</label>
</div>
```
FlyonUI provides a `data-toggle-password` attribute that handles the toggling functionality automatically when the corresponding JavaScript is included.

## Configuration Options
- **CSS Variables:**
  - `--size`: Customizes the size of the checkbox.
  - `--input-color`: Customizes the color of the checked state.
- **Data Attributes:**
  - `data-toggle-password`: Used to associate a checkbox with a password input for toggling visibility.

### Example of using CSS variables:
```css
.my-custom-checkbox {
  --size: 2rem;
  --input-color: #ff0000;
}