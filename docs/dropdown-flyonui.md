# Dropdown - FlyonUI

## Complete List of Classes
- **Base:**
  - `dropdown`: The main container for the dropdown component.
  - `dropdown-toggle`: The button that opens and closes the dropdown.
  - `dropdown-menu`: The container for the dropdown items.
  - `dropdown-item`: A single item within the dropdown menu.
  - `dropdown-header`: A header element within the dropdown menu.
  - `dropdown-footer`: A footer element within the dropdown menu.
- **State Modifiers:**
  - `dropdown-open:opacity-100`: Makes the dropdown visible when open.
  - `dropdown-open:rotate-180`: Rotates an icon when the dropdown is open.
- **Positioning:**
  - The dropdown's position is relative to the `dropdown` container. You can use positioning utilities if needed.
- **Other:**
  - `join-item`: For grouping a dropdown with a button.

## Variations and Sizes
The size of the dropdown menu is typically controlled by its content and can be constrained with `min-w-*` and `max-w-*` classes. The toggle button can be styled using standard button classes.

## HTML Examples

### Basic Dropdown
```html
<div class="dropdown relative inline-flex">
  <button id="dropdown-default" type="button" class="dropdown-toggle btn btn-primary" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
    Dropdown
    <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
  </button>
  <ul class="dropdown-menu dropdown-open:opacity-100 hidden min-w-60" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-default">
    <li><a class="dropdown-item" href="#">My Profile</a></li>
    <li><a class="dropdown-item" href="#">Settings</a></li>
    <li><a class="dropdown-item" href="#">Billing</a></li>
  </ul>
</div>
```

### Dropdown with Avatar
```html
<div class="dropdown relative inline-flex">
  <button id="dropdown-avatar" type="button" class="dropdown-toggle btn btn-outline btn-primary flex items-center gap-2 rounded-full" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
    <div class="avatar">
      <div class="size-6 rounded-full">
        <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-3.png" alt="User Avatar" />
      </div>
    </div>
    John Doe
    <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
  </button>
  <ul class="dropdown-menu dropdown-open:opacity-100 hidden min-w-60" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-avatar">
    <li class="dropdown-header gap-3">...</li>
    <li><a class="dropdown-item" href="#">My Profile</a></li>
  </ul>
</div>
```

### Dropdown with Form
```html
<div class="dropdown relative inline-flex [--auto-close:inside]">
  <button id="dropdown-form" type="button" class="dropdown-toggle btn btn-primary" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
    Dropdown form
  </button>
  <div class="dropdown-menu dropdown-open:opacity-100 min-w-70 hidden" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-form">
    <form class="p-4">
      ...form fields and button...
    </form>
  </div>
</div>
```

## JavaScript Interaction API
FlyonUI's dropdown is an interactive component that requires JavaScript.

### Initialization
The dropdowns are automatically initialized if you are using the main FlyonUI script and the correct data attributes.
```javascript
import { initDropdowns } from 'flowbite'; // Or your specific FlyonUI initialization script
initDropdowns();
```

### Programmatic Control
You can control the dropdown programmatically if you have an instance.
```javascript
// Example, actual API may vary
const myDropdown = HSDropdown.getInstance('#my-dropdown-id', true);
myDropdown.element.open();
myDropdown.element.close();
```

### Static Methods
```javascript
// Open a dropdown without an instance
HSDropdown.open('#dropdown-method');
```

## Configuration Options
- `[--auto-close:inside]`: A custom property to control if the dropdown closes when clicking inside it.
- `[--placement:bottom]`: A custom property to control the placement of the dropdown menu.
- `[--offset:5]`: A custom property to control the offset of the dropdown menu.
- `[--scope:window]`: A custom property to position the dropdown relative to the window.
- `[--has-autofocus:true]`: A custom property to autofocus on the dropdown.
- `[--gpu-acceleration:true]`: A custom property to enable GPU acceleration.