# Timeline - FlyonUI

## Complete List of Classes
- **Base:**
  - `timeline`: The main container for the timeline component, which is a `<ul>` element.
  - `timeline-item`: A `<li>` element representing a single event in the timeline.
  - `timeline-start`: The content aligned to the start of the timeline item.
  - `timeline-end`: The content aligned to the end of the timeline item.
  - `timeline-middle`: The element in the middle of the timeline, usually the icon or bubble.
  - `timeline-box`: A styled box for the content of a timeline item.
- **Orientation:**
  - `timeline-vertical`: (Default) A vertical timeline.
  - `timeline-horizontal`: A horizontal timeline.
- **Modifiers:**
  - `timeline-compact`: A more compact version of the timeline.

## Variations and Sizes
The timeline component is a structural component. Its size and appearance are determined by the content and utility classes applied to its items.

### Orientation
- **Vertical:**
  ```html
  <ul class="timeline">...</ul>
  ```
- **Horizontal:**
  ```html
  <ul class="timeline timeline-horizontal">...</ul>
  ```

### Compact
```html
<ul class="timeline timeline-compact">...</ul>
```

## HTML Examples

### Basic Vertical Timeline
```html
<ul class="timeline">
  <li>
    <div class="timeline-start">1984</div>
    <div class="timeline-middle">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd" /></svg>
    </div>
    <div class="timeline-end timeline-box">First Macintosh computer</div>
    <hr/>
  </li>
  <li>
    <hr/>
    <div class="timeline-start">1998</div>
    <div class="timeline-middle">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd" /></svg>
    </div>
    <div class="timeline-end timeline-box">iMac G3 released</div>
  </li>
</ul>
```

### Horizontal Timeline
```html
<ul class="timeline timeline-horizontal">
  <li>
    <div class="timeline-middle">...</div>
    <div class="timeline-end timeline-box">...</div>
    <hr/>
  </li>
  ...
</ul>
```

## JavaScript Interaction API
The timeline is a CSS-only component and does not have a JavaScript API.

## Configuration Options
All customization is done through utility classes and the structure of the `<li>` elements within the `<ul>`. There are no specific data attributes or CSS variables for this component.