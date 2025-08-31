# Form Validation - FlyonUI

FlyonUI provides a custom form validation a-flyonui.mdpproach that enhances the browser's built-in HTML5 validation.

## Complete List of Classes
- **Form State:**
  - `needs-validation`: Add this class to your `<form>` to enable the custom validation script.
  - `was-validated`: This class is added by the script after the first submission attempt to show validation messages.
- **Message Styling:**
  - `error-message`: Class for the container of the error message for an input.
  - `success-message`: Class for the container of the success message for an input.

## How it Works
1.  Add the `needs-validation` class to your form and the `novalidate` attribute to disable default browser validation popups.
2.  Add standard HTML5 validation attributes to your inputs, like `required`, `type="email"`, etc.
3.  Provide `span` or `div` elements with the `error-message` and `success-message` classes after your inputs.
4.  Include the validation script which will handle the submission event and toggle the visibility of your custom messages.

## HTML Examples

### Basic Form Structure
```html
<form class="needs-validation" novalidate>
  <div>
    <label class="label-text" for="firstName">First Name</label>
    <input id="firstName" type="text" placeholder="John" class="input" required />
    <span class="error-message">Please enter your name.</span>
    <span class="success-message">Looks good!</span>
  </div>

  <button type="submit" class="btn btn-primary mt-4">Submit</button>
</form>
```

### Email and Password
```html
<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
  <div>
    <label class="label-text" for="userEmail">Email</label>
    <input id="userEmail" type="email" class="input" placeholder="john@gmail.com" required />
    <span class="error-message">Please enter a valid email</span>
    <span class="success-message">Looks good!</span>
  </div>
  <div>
    <label class="label-text" for="userPassword">Password</label>
    <input id="userPassword" type="password" class="input" required />
    <span class="error-message">Please enter a valid password</span>
    <span class="success-message">Looks good!</span>
  </div>
</div>
```

## JavaScript Interaction API

The validation is handled by a script that listens for the `submit` event on forms with the `.needs-validation` class.

### Validation Script
This script should be included on your page to enable the custom validation styling.
```javascript
document.addEventListener('DOMContentLoaded', function () {
  const forms = document.querySelectorAll('.needs-validation');

  Array.prototype.slice.call(forms).forEach(function (form) {
    form.addEventListener('submit', function (event) {
      if (!form.checkValidity()) {
        event.preventDefault();
        event.stopPropagation();
        const firstInvalidElement = form.querySelector(':invalid');
        if (firstInvalidElement) {
          firstInvalidElement.focus();
        }
      }
      form.classList.add('was-validated');
    }, false);
  });
});
```

## Configuration Options
The configuration is done through standard HTML5 attributes on the input fields.

-   `required`: Specifies that the input field must be filled out.
-   `type`: Defines the type of input, which comes with built-in validation (e.g., `email`, `number`).
-   `pattern`: Specifies a regular expression that the input's value must match.
-   `minlength` and `maxlength`: Specify the minimum and maximum length of the text.
-   `min` and `max`: Specify the minimum and maximum values for numerical inputs.