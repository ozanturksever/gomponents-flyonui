# Accordion - FlyonUI

## Complete List of Classes
- **Base:**
  - `accordion`: Main container for the accordion component.
  - `accordion-item`: A single item within the accordion.
  - `accordion-toggle`: The button that controls the accordion item's state.
  - `accordion-content`: The collapsible content of an accordion item.
- **State Modifiers:**
  - `active`: Indicates the currently active/open accordion item.
  - `accordion-item-active`: A class applied to the active `accordion-item` for styling purposes.
- **Variations:**
  - `accordion-bordered`: Adds a border to the accordion.
  - `accordion-shadow`: Adds a shadow to the accordion items.
  - `divide-y` & `divide-neutral/20`: Used for creating dividers between accordion items.

## Variations and Sizes
FlyonUI accordions are highly customizable with various classes. Sizes are typically controlled by utility classes from Tailwind CSS (e.g., `text-sm`, `p-4`, `size-4.5`). There are no specific `xs, sm, md, lg, xl` size classes for the accordion component itself, but you can customize it using standard Tailwind classes.

**Color and Theme Variants:**
Colors are managed through Tailwind's color palette (e.g., `text-primary`, `bg-primary/10`).

## HTML Examples
### Basic Accordion
```html
<div class="accordion divide-neutral/20 divide-y">
  <div class="accordion-item active" id="payment-basic">
    <button class="accordion-toggle inline-flex items-center gap-x-4 text-start" aria-controls="payment-basic-collapse" aria-expanded="true" >
      <span class="icon-[tabler--plus] accordion-item-active:hidden text-base-content size-4.5 block shrink-0"></span>
      <span class="icon-[tabler--minus] accordion-item-active:block text-base-content size-4.5 hidden shrink-0"></span>
      When is payment taken for my order?
    </button>
    <div id="payment-basic-collapse" class="accordion-content w-full overflow-hidden transition-[height] duration-300" aria-labelledby="payment-basic" role="region" >
      <div class="px-5 pb-4">
        <p class="text-base-content/80 font-normal">
          Payment is taken during the checkout process when you pay for your order. The order number that appears on the confirmation screen indicates payment has been successfully processed.
        </p>
      </div>
    </div>
  </div>
  <div class="accordion-item" id="delivery-basic">
    <button class="accordion-toggle inline-flex items-center gap-x-4 text-start" aria-controls="delivery-basic-collapse" aria-expanded="false" >
        <span class="icon-[tabler--plus] accordion-item-active:hidden text-base-content size-4.5 block shrink-0"></span>
        <span class="icon-[tabler--minus] accordion-item-active:block text-base-content size-4.5 hidden shrink-0"></span>
        How would you ship my order?
    </button>
    <div id="delivery-basic-collapse" class="accordion-content hidden w-full overflow-hidden transition-[height] duration-300" aria-labelledby="delivery-basic" role="region" >
        <div class="px-5 pb-4">
            <p class="text-base-content/80 font-normal">For large products, we deliver your product via a third party logistics company offering you the “room of choice” scheduled delivery service. For small products, we offer free parcel delivery.</p>
        </div>
    </div>
  </div>
</div>
```

### Nested Accordion
```html
<div class="accordion divide-neutral/20 divide-y">
    <div class="accordion-item active" id="payment-nested">
        <button class="accordion-toggle inline-flex items-center gap-x-4 text-start" aria-controls="payment-nested-collapse" aria-expanded="true">
            <span class="icon-[tabler--plus] accordion-item-active:hidden text-base-content size-4.5 block shrink-0"></span>
            <span class="icon-[tabler--minus] accordion-item-active:block text-base-content size-4.5 hidden shrink-0"></span>
            Payment
        </button>
        <div id="payment-nested-collapse" class="accordion-content w-full overflow-hidden transition-[height] duration-300" aria-labelledby="payment-nested" role="region">
            <div class="accordion divide-neutral/20 divide-y ps-6">
                <div class="accordion-item active" id="payment-sub-one">
                    <button class="accordion-toggle inline-flex items-center gap-x-4 text-start" aria-controls="payment-sub-collapse-one" aria-expanded="true">
                        <span class="icon-[tabler--plus] accordion-item-active:hidden text-base-content size-4.5 block shrink-0"></span>
                        <span class="icon-[tabler--minus] accordion-item-active:block text-base-content size-4.5 hidden shrink-0"></span>
                        How do I pay for my order?
                    </button>
                    <div id="payment-sub-collapse-one" class="accordion-content w-full overflow-hidden transition-[height] duration-300" aria-labelledby="payment-sub-one" role="region">
                        <div class="px-5 pb-4">
                            <p class="text-base-content/80 font-normal">We accept Visa®, MasterCard®, American Express®, and PayPal®. Our servers encrypt all information submitted to them, so you can be confident that your credit card information will be kept safe and secure.</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
```

## JavaScript Interaction API

FlyonUI's accordion component can be controlled with JavaScript.

### Initialization
Include the accordion script in your HTML:
```html
<script src="../path/to/flyonui/dist/accordion.js"></script>
```

Or import it in your JavaScript module:
```javascript
import 'flyonui/dist/accordion.js';
```

### Methods
- `HSAccordion.autoInit()`: Initializes all accordion components on the page.
- `HSAccordion.getInstance(element, isInstanceOwner)`: Retrieves the accordion instance associated with a given element. `isInstanceOwner` should be set to `true`.
- `element.destroy()`: Removes the accordion functionality from the element.

### Usage Example
```javascript
window.addEventListener('load', () => {
    // Auto initialize all accordions
    HSAccordion.autoInit();

    const accordions = document.querySelectorAll('.accordion-to-destroy');
    const destroyBtn = document.querySelector('#destroy-btn');
    const reinitBtn = document.querySelector('#reinit-btn');

    if (destroyBtn) {
        destroyBtn.addEventListener('click', () => {
            accordions.forEach(el => {
                const { element } = HSAccordion.getInstance(el, true);
                element.destroy();
            });
            reinitBtn.removeAttribute('disabled');
            destroyBtn.setAttribute('disabled', true);
        });
    }

    if(reinitBtn) {
        reinitBtn.addEventListener('click', () => {
          HSAccordion.autoInit();
          reinitBtn.setAttribute('disabled', true);
          destroyBtn.removeAttribute('disabled');
        });
    }
});
```

## Configuration Options
The accordion component can be configured using data attributes on the HTML elements.
- `data-hs-accordion-always-open`: Allows multiple accordion items to be open at the same time.
