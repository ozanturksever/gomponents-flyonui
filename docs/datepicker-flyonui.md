# Datepicker - FlyonUI

FlyonUI uses the popular [Flatpickr](https://flatpickr.js.org/) library for its datepicker component. This provides a powerful and lightweight date and time picker.

## Complete List of Classes
FlyonUI's Datepicker doesn't have a specific set of component classes. It is styled through a combination of the standard `input` classes and Flatpickr's own CSS, which can be customized.

- **Input Styling:** `input`, `input-floating`, `is-valid`, `is-invalid`
- **Flatpickr States:**
  - `flatpickr-success`: for success state styling.
  - `flatpickr-error`: for error state styling.

## Variations and Sizes
The datepicker's appearance and size are controlled by the classes on the `input` element and the options passed to the Flatpickr instance.

## HTML Examples

### Default Datepicker
```html
<input type="text" class="input max-w-sm" placeholder="YYYY-MM-DD" id="flatpickr-default" />
```

### Date Range Picker
```html
<input type="text" class="input max-w-sm" placeholder="YYYY-MM-DD to YYYY-MM-DD" id="flatpickr-range" />
```

### DateTime Picker
```html
<input type="text" class="input max-w-sm" placeholder="YYYY-MM-DD HH:MM" id="flatpickr-date-time" />
```

### Floating Label Datepicker
```html
<div class="input-floating max-w-sm">
  <input type="text" placeholder="YYYY-MM-DD" class="input" id="flatpickr-floating" />
  <label class="input-floating-label" for="flatpickr-floating">Date</label>
</div>
```

## JavaScript Interaction API

The Datepicker is initialized and controlled via the `flatpickr` JavaScript object.

### Initialization
First, include the Flatpickr library.
```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css">
<script src="https://cdn.jsdelivr.net/npm/flatpickr"></script>
```
And import the custom FlyonUI styles:
```css
@import "flyonui/src/vendor/flatpickr.css";
```

Then, initialize the datepicker on your input element.

```javascript
window.addEventListener('load', function () {
  // Default Datepicker
  flatpickr('#flatpickr-default', {
    monthSelectorType: 'static'
  });

  // Range Picker
  flatpickr('#flatpickr-range', {
    mode: 'range'
  });

  // DateTime Picker
  flatpickr('#flatpickr-date-time', {
    enableTime: true,
    dateFormat: 'Y-m-d H:i'
  });
});
```

## Configuration Options
Flatpickr has a rich set of configuration options. Here are some of the most common ones used with FlyonUI:

- `mode`: `'single'`, `'multiple'`, or `'range'`.
- `dateFormat`: A string to format the date. e.g., `'Y-m-d'`.
- `enableTime`: Set to `true` to enable time selection.
- `noCalendar`: Set to `true` to create a time-only picker.
- `altInput`: Set to `true` to show a human-friendly date format to the user.
- `altFormat`: The format for the `altInput`. e.g., `'F j, Y'`.
- `inline`: Set to `true` to display the calendar inline.
- `weekNumbers`: Set to `true` to show week numbers.
- `locale`: To translate the calendar (e.g., `'ru'` for Russian).

For a complete list of options, please refer to the [Flatpickr documentation](https://flatpickr.js.org/options/).