# Collapse - FlyonUI

## Complete List of Classes
- **Base:**
  - `collapse`: The main container for the collapsible content.
  - `collapse-toggle`: The button that controls the collapse state.
- **State Modifiers:**
  - `collapse-open:hidden`: Hides an element when the collapse is open.
  - `collapse-open:block`: Shows an element when the collapse is open.
  - `collapse-open:rotate-180`: Rotates an element 180 degrees when the collapse is open.

## Variations and Sizes
The Collapse component itself doesn't have specific size or color variations. Its appearance is determined by the content within it and the classes applied to the `collapse-toggle` button. You can use standard button classes (`btn`, `btn-primary`, etc.) to style the toggle.

## HTML Examples

### Basic Collapse
```html
<button type="button" class="collapse-toggle btn btn-primary" id="basic-collapse" aria-expanded="false" aria-controls="basic-collapse-heading" data-collapse="#basic-collapse-heading">
  Collapse
  <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 size-4"></span>
</button>
<div id="basic-collapse-heading" class="collapse hidden w-full overflow-hidden transition-[height] duration-300" aria-labelledby="basic-collapse">
  <div class="border-base-content/25 mt-3 rounded-md border p-3">
    <p class="text-base-content/80">
      The collapsible body remains concealed by default until the collapse plugin dynamically adds specific classes. These classes are instrumental in styling each element, dictating the overall appearance, and managing visibility through CSS transitions.
    </p>
  </div>
</div>
```

### Show/Hide Content
```html
<div>
  <h6 class="text-base-content text-base">How can I track my order?</h6>
  <div id="show-hide-collapse-heading" class="collapse hidden w-full overflow-hidden transition-[height] duration-300" aria-labelledby="show-hide-collapse" >
    <p class="text-base-content/80"> To track your order, simply log in to your account and navigate to the order history section. You'll find detailed information about your order status and tracking number there.
    </p>
  </div>
</div>
<button type="button" class="collapse-toggle link link-primary inline-flex items-center" id="show-hide-collapse" aria-expanded="false" aria-controls="show-hide-collapse-heading" data-collapse="#show-hide-collapse-heading" >
  <span class="collapse-open:hidden">Read more</span>
  <span class="collapse-open:block hidden">Read less</span>
  <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 ms-2 size-4"></span>
</button>
```

### Nested Collapse in Dropdown
```html
<div class="dropdown relative inline-flex [--auto-close:inside]">
  <button id="dropdown-collapse" type="button" class="dropdown-toggle btn btn-primary" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
    Actions
    <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
  </button>
  <div class="dropdown-menu dropdown-open:opacity-100 hidden min-w-60" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-collapse">
    <div class="dropdown-header">Quick Actions</div>
    <div>
      <button id="nested-collapse-2" class="collapse-toggle dropdown-item justify-between" aria-expanded="false" aria-controls="nested-collapse-content" data-collapse="#nested-collapse-content" >
        More Options
        <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 size-4"></span>
      </button>
      <div class="collapse hidden w-full overflow-hidden transition-[height] duration-300" aria-labelledby="nested-collapse-2" id="nested-collapse-content" >
        <ul class="py-3 ps-3">
          <li><a class="dropdown-item" href="#">Download Documents</a></li>
        </ul>
      </div>
    </div>
  </div>
</div>
```

## JavaScript Interaction API
The FlyonUI Collapse component is controlled via the `HSCollapse` object.

### Initialization
```javascript
import HSCollapse from 'flyonui/collapse';

// To initialize all collapse components
HSCollapse.autoInit();
```

### Methods
- `HSCollapse.getInstance(element, isInstanceOwner)`: Retrieves the collapse instance.
- `instance.show()`: Programmatically shows the collapsible element.
- `instance.hide()`: Programmatically hides the collapsible element.
- `instance.destroy()`: Removes the collapse functionality.
- `HSCollapse.show(target)`: Static method to show a collapsible element.
- `HSCollapse.hide(target)`: Static method to hide a collapsible element.

### Events
- `open.hs.collapse`: Fires when the collage is opened.
- `hide.hs.collapse`: Fires when the collage is hidden.

### Usage Example
```javascript
window.addEventListener('load', function () {
  const collapseElement = document.querySelector('#myCollapse');
  if(collapseElement) {
      const collapse = new HSCollapse(collapseElement);
      const showBtn = document.querySelector('#show-btn');
      const hideBtn = document.querySelector('#hide-btn');

      showBtn.addEventListener('click', () => {
        collapse.show();
      });
      hideBtn.addEventListener('click', () => {
        collapse.hide();
      });

      collapseElement.addEventListener('open.hs.collapse', () => {
        console.log('Collapse opened!');
      });
  }
});
```

## Configuration Options
- `data-collapse`: The primary attribute to enable collapse functionality. The value should be a CSS selector pointing to the collapsible element.