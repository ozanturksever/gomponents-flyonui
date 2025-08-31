# Tabs - FlyonUI

## Complete List of Classes
- **Base:**
  - `tabs`: The main container for the tab navigation.
  - `tab`: A single tab item.
- **State:**
  - `tab-active`: Class for the currently active tab.
- **Variants:**
  - `tabs-bordered`: Adds a border below the tabs.
  - `tabs-lifted`: Gives the active tab a "lifted" appearance.
  - `tabs-boxed`: Styles tabs to look like buttons in a row.
- **Size Variants:**
  - `tabs-lg`
  - `tabs-md` (default)
  - `tabs-sm`
  - `tabs-xs`
- **Content:**
  - `tab-content`: (Not a direct class, but a pattern) A container for the tab panels. Panels are shown/hidden based on the active tab.

## Variations and Sizes

### Variants
```html
<!-- Default -->
<div role="tablist" class="tabs">...</div>

<!-- Bordered -->
<div role="tablist" class="tabs tabs-bordered">...</div>

<!-- Lifted -->
<div role="tablist" class="tabs tabs-lifted">...</div>

<!-- Boxed -->
<div role="tablist" class="tabs tabs-boxed">...</div>
```

### Size Variants
```html
<div role="tablist" class="tabs tabs-lg">...</div>
<div role="tablist" class="tabs tabs-md">...</div>
<div role="tablist" class="tabs tabs-sm">...</div>
<div role="tablist" class="tabs tabs-xs">...</div>
```

## HTML Examples

### Basic Tabs
```html
<div role="tablist" class="tabs">
  <a role="tab" class="tab tab-active">Tab 1</a>
  <a role="tab" class="tab">Tab 2</a>
  <a role="tab" class="tab">Tab 3</a>
</div>
```

### Tabs with Content
This pattern requires JavaScript to toggle the visibility of the content panels.
```html
<div data-component="tabs">
  <div role="tablist" class="tabs tabs-lifted">
    <a role="tab" class="tab tab-active" data-tab-id="panel-1">Tab 1</a>
    <a role="tab" class="tab" data-tab-id="panel-2">Tab 2</a>
  </div>
  <div class="tab-content-container">
    <div id="panel-1" class="tab-content">
      Content for Tab 1
    </div>
    <div id="panel-2" class="tab-content hidden">
      Content for Tab 2
    </div>
  </div>
</div>
```

## JavaScript Interaction API
The tabs component is interactive and relies on the `HSTabs` JavaScript class.

### Initialization
Tabs are typically initialized automatically via data attributes.
```javascript
// Manually initialize a tabs component
const tabsElement = document.querySelector('[data-component="tabs"]');
const tabs = new HSTabs(tabsElement);
```

### Methods
The documentation for `HSTabs` shows it's part of an ecosystem of components, but specific methods for the Tabs component beyond initialization are not detailed in the provided Go source. Interaction is mainly handled through `data` attributes.

## Configuration Options
- `data-component="tabs"`: The main attribute to initialize the tabs component.
- `data-tab-id="panel-id"`: On a tab link, points to the ID of the content panel it controls.
- `data-tab-panel="panel-id"`: On a content panel, identifies it as a target for a tab link.