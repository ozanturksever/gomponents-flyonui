# Container - FlyonUI

## Complete List of Classes
FlyonUI's container is not a specific component with a `.container` class in the same way as other components. It's a general concept of wrapping content. It relies on standard Tailwind CSS utility classes for layout, padding, and responsive design.

- **Layout & Sizing:** `mx-auto`, `max-w-7xl`, `p-6`, `py-8`, `px-4`, etc.
- **Styling:** `bg-base-100`, `rounded-box`, `border`, `shadow-md`, etc.

The containers in FlyonUI are flexible and built using these utilities rather than a single, rigid `container` class.

## Variations and Sizes
Containers are sized and styled using Tailwind CSS utility classes. There are no predefined `container-sm` or `container-lg` classes.

**Examples of container definitions:**
- `mx-auto max-w-7xl px-4 sm:px-6 lg:px-8`: A centered container with a max-width and responsive padding.
- `rounded-box border p-6`: A container with rounded corners, a border, and padding.

## HTML Examples

### Basic Content Container
```html
<div class="bg-base-100 py-8 sm:py-16 lg:py-24">
  <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
    <!-- Your content here -->
    <h2>Page Title</h2>
    <p>This is some content inside a responsive container.</p>
  </div>
</div>
```

### Card as a Container
A `card` is a common container for content.
```html
<div class="card sm:max-w-sm">
  <div class="card-body">
    <h5 class="card-title">Card Title</h5>
    <p>This card acts as a container for this text and button.</p>
    <a href="#" class="btn btn-primary">Action</a>
  </div>
</div>
```

### Stats Container
The `stats` class creates a container for stat items.
```html
<div class="stats">
  <div class="stat">
    <div class="stat-title">Total Page Views</div>
    <div class="stat-value">89,400</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Users</div>
    <div class="stat-value">4,200</div>
  </div>
</div>
```

## JavaScript Interaction API
Since the container is a structural element, it does not have a dedicated JavaScript API. Any interaction would be with the components placed inside the container.

## Configuration Options
There are no JavaScript-based configuration options for the container itself. All configuration is done through HTML classes.