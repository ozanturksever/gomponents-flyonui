# Form Group - FlyonUI

## Form Grouping Pattern

In FlyonUI, form elements are typically grouped together for structure and spacing. There isn't a specific `.form-group` class; instead, this is achieved by wrapping labels and inputs within a `div` and using utility classes for layout.

## Complete List of Classes
The "form group" pattern relies on standard spacing and layout utilities from Tailwind CSS.

- `mb-4`: A common utility to add a margin to the bottom of the group.
- `flex`, `flex-col`, `gap-2`: For more complex layouts within a form group.
- `label-text`: For styling the text of a `<label>`.
- `helper-text`: For styling helper text below an input.

## HTML Examples

### Basic Form Group
A simple group consists of a `label` and an `input` wrapped in a `div`.
```html
<div class="mb-4">
  <label class="label-text" for="fullName"> Full Name </label>
  <input type="text" placeholder="John Doe" class="input" id="fullName" />
</div>
```

### Form Group with Helper Text
```html
<div class="max-w-sm">
  <label class="label-text" for="fileInputHelperText"> Pick a file </label>
  <input type="file" class="input" id="fileInputHelperText" />
  <span class="helper-text">Helper text</span>
</div>
```

### Floating Label Group
The `input-floating` class on the parent `div` creates a floating label effect.
```html
<div class="input-floating max-w-sm">
  <input type="text" placeholder="name@example.com" class="input" id="emailFloating" />
  <label class="input-floating-label" for="emailFloating">Email</label>
</div>
```

### Input Group (Join)
To group an input with a button or another element, the `join` class is used on the parent, and `join-item` on the children.
```html
<div class="join w-96">
  <div class="input join-item flex">
    <span class="icon-[tabler--user] text-base-content/80 my-auto me-3 size-5 shrink-0"></span>
    <label class="sr-only" for="groupInput">Full Name</label>
    <input type="text" class="grow" placeholder="John Doe" id="groupInput" />
  </div>
  <button class="btn btn-outline btn-secondary join-item h-auto">Search</button>
</div>
```

## JavaScript Interaction API
The form group pattern is a structural and styling concept, so it does not have a dedicated JavaScript API.

## Configuration Options
All configuration for form groups is handled through HTML classes and the structure of the elements.