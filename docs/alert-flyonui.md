# Alert - FlyonUI

## Complete List of Classes
- **Base:**
  - `alert`: The main class for the alert component.
- **Color Variants:**
  - `alert-primary`
  - `alert-secondary`
  - `alert-success`
  - `alert-danger`
  - `alert-warning`
  - `alert-info`
- **Style Variants:**
  - `alert-soft`: For a softer color scheme.
  - `alert-outline`: For an outlined style.
  - `alert-dashed`: For a dashed border style.
- **Responsive:**
  - `alert-responsive`: For responsive styling.
- **Dismissible:**
  - `removing:translate-x-5`
  - `removing:opacity-0`
- **Inside Card:**
  - `card-alert`: To be used when an alert is inside a card component.

## Variations and Sizes
FlyonUI alerts come in several color and style variations. Sizes are controlled by Tailwind CSS utility classes.

### Color Variants
- `alert-primary`
- `alert-secondary`
- `alert-success`
- `alert-danger`
- `alert-warning`
- `alert-info`

### Style Variants
- **Solid (default):**
  ```html
  <div class="alert alert-primary" role="alert">...</div>
  ```
- **Soft:**
  ```html
  <div class="alert alert-soft alert-primary" role="alert">...</div>
  ```
- **Outline:**
  ```html
  <div class="alert alert-outline alert-primary" role="alert">...</div>
  ```
- **Dashed:**
  ```html
  <div class="alert alert-dashed alert-primary" role="alert">...</div>
  ```

## HTML Examples
### Basic Alerts
```html
<div class="alert alert-primary" role="alert">
    This is a primary alert message.
</div>
<div class="alert alert-success" role="alert">
    This is a success alert message.
</div>
```

### Alerts with Icons
```html
<div class="alert alert-warning flex items-center gap-4" role="alert">
  <span class="icon-[tabler--alert-triangle] shrink-0 size-6"></span>
  <p><span class="text-lg font-semibold">Warning alert:</span> Stay informed about the latest updates and upcoming events.</p>
</div>
```

### Dismissible Alert
```html
<div class="alert alert-soft alert-primary removing:translate-x-5 removing:opacity-0 flex items-center gap-4 transition duration-300 ease-in-out" role="alert" id="dismiss-alert">
  Dive into our platform to discover exciting new features and updates.
  <button class="ms-auto cursor-pointer leading-none" data-remove-element="#dismiss-alert" aria-label="Close Button">
    <span class="icon-[tabler--x] size-5"></span>
  </button>
</div>
```

### Alert with List
```html
<div class="alert alert-soft alert-primary flex items-start gap-4">
    <span class="icon-[tabler--info-circle] shrink-0 size-6"></span>
    <div class="flex flex-col gap-1">
        <h5 class="text-lg font-semibold">Please ensure that your password meets the following requirements:</h5>
        <ul class="mt-1.5 list-inside list-disc">
            <li>Contains a minimum of 10 characters and a maximum of 100 characters.</li>
            <li>Includes at least one lowercase character.</li>
            <li>Incorporates at least one special character such as !, @, #, or ?.</li>
        </ul>
    </div>
</div>
```

## JavaScript Interaction API
The main JavaScript interaction for alerts is the dismiss functionality, which is handled via a `data-remove-element` attribute.

### Dismissing an Alert
To make an alert dismissible, add an ID to the alert and a button with the `data-remove-element` attribute pointing to that ID.

**HTML:**
```html
<div class="alert alert-soft alert-primary removing:translate-x-5 removing:opacity-0 flex items-center gap-4 transition duration-300 ease-in-out" role="alert" id="dismiss-alert-example">
  This is a dismissible alert.
  <button class="ms-auto cursor-pointer leading-none" data-remove-element="#dismiss-alert-example" aria-label="Close Button">
    <span class="icon-[tabler--x] size-5"></span>
  </button>
</div>
```
No extra JavaScript is needed for the dismiss functionality if you are using the FlyonUI script that handles `data-remove-element`.

## Configuration Options
- `--alert-color`: CSS variable to customize the color of the alert.
- `data-remove-element`: Attribute to specify the target element to remove when the button is clicked.

### Example of using CSS variable:
```css
:root {
  --alert-color: #your-custom-color;
}