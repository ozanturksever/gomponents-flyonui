# Card - FlyonUI

## Complete List of Classes
- **Base:**
  - `card`: The main container for the card component.
  - `card-header`: The header section of the card.
  - `card-body`: The main content area of the card.
  - `card-footer`: The footer section of the card.
  - `card-title`: For the title text within a card.
  - `card-text`: For the main text content within a card.
  - `card-actions`: A container for action buttons.
- **Style Variants:**
  - `card-border`: Adds a border to the card.
  - `image-full`: Used for cards where the image fills the background.
  - `glass`: for a frosted glass effect.
- **Size Variants:**
  - `card-xs`
  - `card-sm`
  - `card-lg`
  - `card-xl`
- **Layout Variants:**
  - `card-side`: For a horizontal layout.
- **Inside Card:**
  - `card-alert`: To be used when an alert is inside a card component.
- **Dismissible:**
  - `removing:translate-x-5`
  - `removing:opacity-0`

## Variations and Sizes

### Size Variants
```html
<div class="card card-xs">...</div>
<div class="card card-sm">...</div>
<div class="card">...</div> <!-- default size -->
<div class="card card-lg">...</div>
<div class="card card-xl">...</div>
```

### Style Variants
- **Default:** `card`
- **Bordered:** `card card-border`
- **Image Overlay:** `card image-full`
- **Glass:** `card glass`
- **Horizontal:** `card card-side`

## HTML Examples

### Basic Card
```html
<div class="card sm:max-w-sm">
  <div class="card-header">
    <h5 class="card-title">Innovative Solutions</h5>
  </div>
  <div class="card-body">
    <p>Explore our cutting-edge features designed to elevate your experience. Learn how our solutions can help you achieve your goals.</p>
  </div>
  <div class="card-footer text-center">
    <p class="text-base-content/50">Learn more about our features.</p>
  </div>
</div>
```

### Card with Image
```html
<div class="card sm:max-w-sm">
  <figure><img src="https://cdn.flyonui.com/fy-assets/components/card/image-9.png" alt="Watch" /></figure>
  <div class="card-body">
    <h5 class="card-title mb-2.5">Apple Smart Watch</h5>
    <p class="mb-4">Stay connected, motivated, and healthy with the latest Apple Watch.</p>
    <div class="card-actions">
      <button class="btn btn-primary">Buy Now</button>
      <button class="btn btn-secondary btn-soft">Add to cart</button>
    </div>
  </div>
</div>
```

### Image Overlay Card
```html
<div class="card image-full sm:max-w-sm">
  <figure><img src="https://cdn.flyonui.com/fy-assets/components/card/image-5.png" alt="overlay image" /></figure>
  <div class="card-body">
    <h2 class="card-title mb-2.5 text-white">Marketing</h2>
    <p class="text-white">Boost your brand's visibility and engagement through targeted marketing strategies.</p>
  </div>
</div>
```

### Dismissible Card
```html
<div class="card removing:translate-x-5 removing:opacity-0 w-full transition duration-300 ease-in-out" id="card-dismiss">
  <div class="card-header flex justify-between items-center">
    <span class="card-title">Card Actions</span>
    <div class="card-actions">
        <button class="tooltip-toggle btn btn-text btn-sm btn-circle" aria-label="Close Button" data-remove-element="#card-dismiss">
          <span class="icon-[tabler--x] size-5"></span>
        </button>
    </div>
  </div>
  <div class="card-body">
    <p>With a single click on the close button, this card will be effortlessly removed.</p>
  </div>
</div>
```

## JavaScript Interaction API

The main JavaScript interaction for the card component is the dismiss functionality, which is handled via a `data-remove-element` attribute.

### Dismissing a Card
To make a card dismissible, add an ID to the card and a button with the `data-remove-element` attribute pointing to that ID.

**HTML:**
```html
<div class="card" id="dismissible-card">
  <div class="card-body">
    <p>This card can be dismissed.</p>
    <div class="card-actions">
      <button data-remove-element="#dismissible-card" class="btn btn-error">Dismiss</button>
    </div>
  </div>
</div>
```

If you are using the FlyonUI script that handles `data-remove-element`, no extra JavaScript is needed.

## Configuration Options
- **CSS Variables:**
  - `--card-p`: Customizes the padding of the card.
  - `--card-border`: Customizes the border of the card.
  - `--card-shadow`: Customizes the shadow of the card.
- **Data Attributes:**
  - `data-remove-element`: Specifies the target element to remove when a button is clicked.

### Example of using CSS variables:
```css
.my-custom-card {
  --card-p: 2rem;
  --card-border: 2px solid blue;
}