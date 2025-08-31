# Combobox - FlyonUI

## Complete List of Classes
- `combo-box-selected:dropdown-active`: Applied to a dropdown item when it's selected.
- `combo-box-selected:block`: Used to show an element (like a checkmark) when an item is selected.
- `dropdown-item`: Class for items within the dropdown.
- `input`: Standard input class.

The Combobox functionality is primarily driven by data attributes rather than a large set of CSS classes.

## Variations and Sizes
The Combobox component's size and appearance are determined by the standard `input` component classes and Tailwind CSS utilities. There are no specific size variants like `combobox-lg` or `combobox-sm`.

## HTML Examples

### Basic Combobox
```html
<div class="relative max-w-sm" data-combo-box>
  <div class="relative">
    <input class="input" type="text" value="India" role="combobox" aria-expanded="false" data-combo-box-input aria-label="Default combobox" />
    <span class="icon-[tabler--caret-up-down] text-base-content absolute end-3 top-1/2 size-4 shrink-0 -translate-y-1/2" data-combo-box-toggle></span>
  </div>
  <div class="bg-base-100 rounded-box shadow-base-300/20 absolute z-50 max-h-44 w-full space-y-0.5 overflow-y-auto p-2 shadow-lg" style="display: none" data-combo-box-output>
    <div class="dropdown-item combo-box-selected:dropdown-active" tabindex="0" data-combo-box-output-item>
      <div class="flex items-center justify-between">
        <span data-combo-box-search-text="Venezuela" data-combo-box-value="">Venezuela</span>
        <span class="icon-[tabler--check] text-primary combo-box-selected:block hidden size-4 shrink-0"></span>
      </div>
    </div>
    <div class="dropdown-item combo-box-selected:dropdown-active" tabindex="1" data-combo-box-output-item>
      <div class="flex items-center justify-between">
        <span data-combo-box-search-text="Papua New Guinea" data-combo-box-value="">Papua New Guinea</span>
        <span class="icon-[tabler--check] text-primary combo-box-selected:block hidden size-4 shrink-0"></span>
      </div>
    </div>
  </div>
</div>
```

### Combobox with API
```html
<div class="relative max-w-sm" data-combo-box='{
    "apiUrl": "https://www.freetestapi.com/api/v1/countries",
    "apiQuery": "limit=7",
    "apiSearchQuery": "search",
    "outputEmptyTemplate": "<div class=\"dropdown-item\">No countries found...</div>",
    "outputItemTemplate": "<div class=\"dropdown-item combo-box-selected:dropdown-active\" data-combo-box-output-item><div class=\"flex justify-between items-center w-full\"><div data-combo-box-output-item-field=\"name\" data-combo-box-search-text data-combo-box-value></div><span class=\"icon-[tabler--check] text-primary combo-box-selected:block hidden size-4 shrink-0\"></span></div></div>"
  }'>
  <div class="relative">
    <input class="input" type="text" value="India" role="combobox" aria-expanded="false" data-combo-box-input aria-label="Parameters in combobox" />
    <span class="icon-[tabler--caret-up-down] text-base-content absolute end-3 top-1/2 size-4 shrink-0 -translate-y-1/2" data-combo-box-toggle></span>
  </div>
  <div class="bg-base-100 rounded-box shadow-base-300/20 absolute z-50 max-h-44 w-full space-y-0.5 overflow-y-auto p-2 shadow-lg" style="display: none;" data-combo-box-output></div>
</div>
```

## JavaScript Interaction API

The FlyonUI Combobox component is controlled via the `HSComboBox` object.

### Initialization
```javascript
import HSComboBox from 'flyonui/combobox';

// To initialize all combobox components
HSComboBox.autoInit();

// To initialize a single component
const myComboBox = new HSComboBox(document.querySelector('#my-combobox'));
```

### Methods
- `HSComboBox.getInstance(element, isInstanceOwner)`: Retrieves the combobox instance.
- `instance.open()`: Programmatically opens the combobox dropdown.
- `instance.close()`: Programmatically closes the combobox dropdown.
- `instance.destroy()`: Removes the combobox functionality.

### Usage Example
```javascript
window.addEventListener('load', function () {
  const comboBoxEl = document.querySelector('#combo-box-method');
  if (comboBoxEl) {
      const comboBox = new HSComboBox(comboBoxEl);
      const openBtn = document.querySelector('#open-btn');

      openBtn.addEventListener('click', () => {
        comboBox.open();
      });
  }
});
```

## Configuration Options

Configuration is primarily done via `data-combo-box` attribute on the main container.

- `data-combo-box`: The main attribute to enable the component. It can hold a JSON object for configuration.
- `data-combo-box-input`: Identifies the input field.
- `data-combo-box-toggle`: Identifies the toggle button/element.
- `data-combo-box-output`: Identifies the container for the dropdown options.
- `data-combo-box-output-item`: Identifies a single item in the dropdown.
- `data-combo-box-search-text`: The text to use for searching within an item.
- `data-combo-box-value`: The value to be set in the input when an item is selected.

### API Configuration
- `apiUrl`: The base URL for the API.
- `apiSearchPath`: The path for search queries.
- `apiSearchDefaultPath`: The path to load initial data.
- `apiQuery`: A query string to append to the API URL.
- `apiSearchQuery`: The query parameter name for the search term.
- `outputItemTemplate`: An HTML string template for rendering each item from the API response.
- `outputEmptyTemplate`: An HTML string template for when no results are found.
- `outputLoaderTemplate`: An HTML string template for the loading state.