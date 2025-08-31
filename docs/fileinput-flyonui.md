# Fileinput - FlyonUI

## Complete List of Classes
- **Base:**
  - `input`: The base class for the file input component.
  - `file-input`: Can be used for more specific file input styling (though `input` is common).
- **Styling the "Browse" button:**
  - `file:text-bg-primary`: Sets the text color of the file selector button.
  - `file:px-4`: Sets the horizontal padding of the button.
  - `file:h-9.5`: Sets the height of the button.
  - `file:rounded-full`: Makes the button rounded.
  - `file:font-medium`: Sets the font weight.
  - `file:uppercase`: Makes the button text uppercase.
- **States:**
  - `is-valid`: For a success state.
  - `is-invalid`: For an error state.
- **Layout:**
  - `input-floating`: For a floating label layout.

## Variations and Sizes
File input sizes are controlled by Tailwind CSS utility classes on the `input` element itself, not by specific `file-input-*` classes.

## HTML Examples

### Default File Input
```html
<div class="max-w-sm">
  <label class="label-text" for="inpuFileTypeDefault"> Pick a file </label>
  <input type="file" class="input" id="inpuFileTypeDefault" />
</div>
```

### Styled File Input Button
```html
<input type="file" class="block cursor-pointer text-sm file:uppercase file:text-bg-primary file:px-4 file:h-9.5 file:rounded-field cursor-pointer file:font-medium file:text-base file:me-3" aria-label="file-input" />
```

### Floating Label File Input
```html
<div class="max-w-sm input-floating">
  <input type="file" placeholder="John Doe" class="input" id="inpuFileTypeFloating" />
  <label class="input-floating-label" for="inpuFileTypeFloating">Upload</label>
</div>
```

### Multiple File Input
```html
<input type="file" class="input max-w-sm" aria-label="file-input" multiple />
```

### Disabled File Input
```html
<input type="file" class="input max-w-sm" aria-label="file-input" disabled />
```

### File Input with Drag-and-Drop and Previews
FlyonUI uses a third-party library for advanced file upload functionality.
```html
<div data-file-upload='{ "url": "/upload", "acceptedFiles": "image/*"}'>
  <template data-file-upload-preview="">
    <!-- ... template for file preview with progress bar ... -->
  </template>
  <div class="bg-base-200/60 rounded-box flex flex-col justify-center border-2 border-base-content/20 border-dashed">
    <div class="text-center cursor-pointer p-12" data-file-upload-trigger="">
      <p class="text-base-content/50 mb-3 text-sm">Choose a file no larger than 2MB.</p>
      <button class="btn btn-soft btn-sm btn-primary text-nowrap">
        <span class="icon-[tabler--file-upload] size-4.5 shrink-0"></span> Drag & Drop to Upload
      </button>
      <p class="text-base-content/50 my-2 text-xs">or</p>
      <p class="link link-animated link-primary font-medium text-sm">Browse</p>
    </div>
    <div class="mx-12 mb-8 empty:m-0 grid grid-cols-4 gap-2 empty:gap-0 max-sm:grid-cols-2" data-file-upload-previews=""></div>
  </div>
</div>
```

## JavaScript Interaction API
For basic file inputs, you interact with them using standard DOM methods. For the advanced drag-and-drop component, FlyonUI uses a wrapper around a file upload library.

### Initialization (for advanced uploader)
```javascript
// The component is auto-initialized via the data-file-upload attribute.
// You might need to include a specific JS file for this component.
```

### Destroy and Reinitialize
```javascript
window.addEventListener('load', () => {
  const fileUpload = document.querySelector('#file-upload-to-destroy');
  const destroyBtn = document.querySelector('#destroy-btn');
  const reinitBtn = document.querySelector('#reinit-btn');

  destroyBtn.addEventListener('click', () => {
    const { element } = HSFileUpload.getInstance(fileUpload, true);
    element.destroy();
    // ... disable/enable buttons
  });

  reinitBtn.addEventListener('click', () => {
    HSFileUpload.autoInit();
    // ... disable/enable buttons
  });
});
```

## Configuration Options
- `multiple`: A standard HTML attribute to allow multiple file selection.
- `disabled`: A standard HTML attribute to disable the input.
- `data-file-upload`: A data attribute that holds a JSON object for configuring the advanced file uploader.
  - `url`: The endpoint to upload files to.
  - `maxFilesize`: The maximum file size in MB.
  - `acceptedFiles`: A string of accepted file types (e.g., `"image/*"`).