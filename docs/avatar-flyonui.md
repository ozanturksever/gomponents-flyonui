# Avatar - FlyonUI

## Complete List of Classes
- **Base:**
  - `avatar`: The main class for the avatar component.
- **Placeholder:**
  - `avatar-placeholder`: Used for avatars that are placeholders (e.g., solid color with initials).
- **Status Indicators:**
  - `avatar-online-top`, `avatar-away-top`, `avatar-busy-top`, `avatar-offline-top`
  - `avatar-online-bottom`, `avatar-away-bottom`, `avatar-busy-bottom`, `avatar-offline-bottom`
- **Group:**
  - `avatar-group`: Used to group multiple avatars together.
  - `pull-up`: A modifier for the group that creates a pull-up effect.
  - `-space-x-5`: A utility class to create a negative margin between avatars in a group.
- **Shape:**
  - `rounded-full`: Creates a circular avatar.
  - `rounded-md`: Creates a squared avatar with rounded corners.
- **Outline:**
  - `border`, `border-primary`, `border-secondary`, etc. for outlined avatars.
- **Sizes:**
  - Avatars are sized using Tailwind CSS width and height utilities (e.g., `w-10`, `h-10`, `size-12`).

## Variations and Sizes

### Sizes
Sizes are controlled by Tailwind CSS utilities.
```html
<div class="avatar">
  <div class="size-6 rounded-full">
    <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-1.png" alt="avatar" />
  </div>
</div>
<div class="avatar">
  <div class="size-16 rounded-full">
    <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-1.png" alt="avatar" />
  </div>
</div>
```

### Shapes
- **Circular:** `rounded-full`
- **Squared:** `rounded-md`

### Placeholders (Solid Color)
```html
<div class="avatar avatar-placeholder">
  <div class="bg-primary text-primary-content w-10 rounded-full">
    <span class="text-md uppercase">cl</span>
  </div>
</div>
```

### With Status Indicators
```html
<div class="avatar avatar-online-top">
  <div class="w-10 rounded-full">
    <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-1.png" alt="avatar" />
  </div>
</div>
```

### Avatar Group
```html
<div class="avatar-group pull-up -space-x-5">
  <div class="avatar">
    <div class="w-13">
      <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-16.png" alt="avatar" />
    </div>
  </div>
  <div class="avatar">
    <div class="w-13">
      <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-8.png" alt="avatar" />
    </div>
  </div>
</div>
```

## HTML Examples
### Basic Avatar
```html
<div class="avatar">
  <div class="w-24 rounded-full">
    <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-13.png" alt="avatar" />
  </div>
</div>
```

### Avatar with Tooltip
```html
<div class="tooltip">
  <div class="tooltip-toggle avatar">
    <div class="w-13">
      <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-6.png" alt="avatar" />
    </div>
  </div>
  <span class="tooltip-content tooltip-shown:opacity-100 tooltip-shown:visible" role="tooltip">
    <span class="tooltip-body">Jasmine Rivera</span>
  </span>
</div>
```

### Avatar in a Stat Component
```html
<div class="stats">
  <div class="stat">
    <div class="stat-figure">
      <div class="avatar">
        <div class="size-12 rounded-full">
          <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-1.png" alt="User Avatar" />
        </div>
      </div>
    </div>
    <div class="stat-title">Total page views</div>
    <div class="stat-value">89,400</div>
    <div class="stat-desc">21% ↗︎ than last month</div>
  </div>
</div>
```

## JavaScript Interaction API
Avatars are primarily styled with CSS and do not have a dedicated JavaScript API. Interactions like tooltips are handled by the tooltip component's JavaScript. Animated avatars can be created using a supporting animation library or custom CSS.

## Configuration Options
There are no specific JavaScript configuration options for the avatar component itself. Customization is done through HTML classes.