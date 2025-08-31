# Breadcrumb - FlyonUI

The Breadcrumb component provides a navigation aid that helps users understand their current location within a website's hierarchy and allows them to navigate back to higher-level pages.

## Classes

| Class | Type | Description |
|-------|------|-------------|
| `breadcrumbs` | Component | Base class for breadcrumb component. |
| `breadcrumbs-separator` | Part | Used to separate icons or text within breadcrumbs. |

## Basic Usage

### Simple Breadcrumb with Chevron Separators

```html
<div class="breadcrumbs">
  <ul>
    <li>
      <a href="#">Home</a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li>
      <a href="#">Components</a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li aria-current="page">Breadcrumb</li>
  </ul>
</div>
```

### Breadcrumb with Slash Separators

```html
<div class="breadcrumbs">
  <ul>
    <li>
      <a href="#">Home</a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <a href="#">Components</a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li aria-current="page">Breadcrumb</li>
  </ul>
</div>
```

## Variations

### With Icons

Add icons to breadcrumb items for enhanced visual navigation.

```html
<div class="breadcrumbs">
  <ol>
    <li>
      <a href="#">
        <span class="icon-[tabler--star-filled] size-5"></span>
        Home
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li>
      <a href="#">
        <span class="icon-[tabler--star-filled] size-5"></span>
        Components
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li aria-current="page">
      <span class="icon-[tabler--star-filled] me-1 size-5"></span>
      Breadcrumb
    </li>
  </ol>
</div>
```

### With Ellipsis for Long Paths

Use an ellipsis icon to represent intermediate pages in long navigation paths.

```html
<div class="breadcrumbs">
  <ol>
    <li>
      <a href="#">
        <span class="icon-[tabler--folder] size-5"></span>
        Home
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li>
      <a href="#" aria-label="More Pages">
        <span class="icon-[tabler--dots]"></span>
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:rotate-180">
      <span class="icon-[tabler--chevron-right]"></span>
    </li>
    <li aria-current="page">
      <span class="icon-[tabler--file] me-1 size-5"></span>
      Breadcrumb
    </li>
  </ol>
</div>
```

### With Dropdowns

Integrate dropdown menus within breadcrumb items for complex navigation.

```html
<div class="breadcrumbs">
  <ol>
    <li>
      <a href="#">Home</a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <div class="dropdown relative inline-flex">
        <button id="dropdown-default" type="button" class="dropdown-toggle btn btn-text font-normal" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
          Components
          <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
        </button>
        <ul class="dropdown-menu dropdown-open:opacity-100 hidden min-w-10" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-default">
          <li><a class="dropdown-item" href="#">Overlay</a></li>
          <li><a class="dropdown-item" href="#">Navigation</a></li>
          <li><a class="dropdown-item" href="#">Collapse</a></li>
          <li><a class="dropdown-item" href="#">Form</a></li>
        </ul>
      </div>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li aria-current="page">
      <div class="dropdown relative inline-flex">
        <button id="dropdown-default" type="button" class="dropdown-toggle btn btn-text btn-secondary" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
          Component
          <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
        </button>
        <ul class="dropdown-menu dropdown-open:opacity-100 hidden min-w-10" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-default">
          <li><a class="dropdown-item" href="#">Modal</a></li>
          <li><a class="dropdown-item" href="#">Breadcrumb</a></li>
          <li><a class="dropdown-item" href="#">Accordion</a></li>
          <li><a class="dropdown-item" href="#">Input</a></li>
        </ul>
      </div>
    </li>
  </ol>
</div>
```

## Styling Options

### With Background Styling

Apply background styling to highlight the current page or important sections.

```html
<div class="w-full rounded-lg border px-4 py-2">
  <div class="breadcrumbs">
    <ul>
      <li>
        <a href="#">Home</a>
      </li>
      <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
      <li>
        <a href="#">Components</a>
      </li>
      <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
      <li aria-current="page">
        <span class="bg-primary/20 !text-primary rounded-xs px-1.5 py-0.5">Breadcrumb</span>
      </li>
    </ul>
  </div>
</div>
```

### Bordered Container

Add a border to the breadcrumb container for visual separation.

```html
<div class="border-base-content/25 w-full rounded-lg border px-4 py-2">
  <div class="breadcrumbs">
    <ul>
      <li>
        <a href="#">Home</a>
      </li>
      <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
      <li>
        <a href="#">Components</a>
      </li>
      <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
      <li aria-current="page">Breadcrumb</li>
    </ul>
  </div>
</div>
```

### Responsive Width

Apply max-width to breadcrumbs for responsive layouts with long navigation paths.

```html
<div class="breadcrumbs h-fit max-w-xs">
  <ul>
    <li>
      <a href="#">
        <span class="icon-[tabler--folder] size-5"></span>
        Home
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <a href="#">
        <span class="icon-[tabler--folder] size-5"></span>
        App
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <a href="#">
        <span class="icon-[tabler--folder] size-5"></span>
        Components
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <a href="#">
        <span class="icon-[tabler--folder] size-5"></span>
        Navigation
      </a>
    </li>
    <li class="breadcrumbs-separator rtl:-rotate-[40deg]">/</li>
    <li>
      <span class="icon-[tabler--file] me-1 size-5"></span>
      Breadcrumb
    </li>
  </ul>
</div>
```

## Accessibility

The breadcrumb component includes proper ARIA attributes for accessibility:

- Use `aria-current="page"` on the current page item
- Include `aria-label` attributes for icon-only links
- Use semantic HTML with `<nav>`, `<ol>`, and `<li>` elements
- Support for RTL languages with `rtl:` modifiers

## Responsive Design

The breadcrumb component is fully responsive:

- Uses `max-w-*` classes for width constraints
- Supports horizontal scrolling for long paths
- Includes RTL support with `rtl:` modifiers
- Adapts to different screen sizes automatically