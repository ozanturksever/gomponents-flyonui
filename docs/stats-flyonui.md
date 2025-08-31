# Stats - FlyonUI

## Complete List of Classes
- **Base:**
  - `stats`: The main container for a group of statistics.
- **Children:**
  - `stat`: Represents a single statistic item within the `stats` container.
  - `stat-title`: The title or label for the statistic.
  - `stat-value`: The actual value of the statistic.
  - `stat-desc`: A description or additional context for the statistic.
  - `stat-figure`: A container for an icon or avatar.
- **Orientation:**
  - `stats-vertical`: Stacks the stat items vertically.
  - `stats-horizontal`: (Default) Lays out the stat items horizontally.

## Variations and Sizes
The Stats component is primarily a layout container. Sizing and colors are applied to the elements within each `.stat` item.

### Orientation
- **Horizontal (Default):**
  ```html
  <div class="stats">
    ...
  </div>
  ```
- **Vertical:**
  ```html
  <div class="stats stats-vertical">
    ...
  </div>
  ```

## HTML Examples

### Basic Stats
```html
<div class="stats shadow">
  <div class="stat">
    <div class="stat-title">Total Page Views</div>
    <div class="stat-value">89,400</div>
    <div class="stat-desc">21% more than last month</div>
  </div>
  <div class="stat">
    <div class="stat-title">New Users</div>
    <div class="stat-value">4,200</div>
    <div class="stat-desc">↗︎ 400 (22%)</div>
  </div>
</div>
```

### Stats with Icons or Avatars
```html
<div class="stats">
  <div class="stat">
    <div class="stat-figure text-primary">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-8 h-8 stroke-current"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
    </div>
    <div class="stat-title">Downloads</div>
    <div class="stat-value">31K</div>
  </div>

  <div class="stat">
    <div class="stat-figure text-secondary">
      <div class="avatar">
        <div class="w-12 rounded-full">
          <img src="https://cdn.flyonui.com/fy-assets/avatar/avatar-1.png" alt="User Avatar" />
        </div>
      </div>
    </div>
    <div class="stat-value">86%</div>
    <div class="stat-title">Tasks done</div>
  </div>
</div>
```

## JavaScript Interaction API
The stats component is a CSS-only component for layout and does not have a JavaScript API.

## Configuration Options
There are no specific configuration options for the stats component. Customization is achieved through utility classes and the structure of the content within it.