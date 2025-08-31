# Typography - FlyonUI

The Typography component in FlyonUI provides a comprehensive set of text styling utilities for creating consistent and visually appealing text content. It leverages Tailwind CSS's typography system to offer flexible control over font sizes, weights, colors, and alignment.

## Classes

The Typography component uses Tailwind CSS classes for styling:

### Font Sizes
- `text-xs`: 0.75rem (12px)
- `text-sm`: 0.875rem (14px)
- `text-base`: 1rem (16px) - default
- `text-lg`: 1.125rem (18px)
- `text-xl`: 1.25rem (20px)
- `text-2xl`: 1.5rem (24px)
- `text-3xl`: 1.875rem (30px)
- `text-4xl`: 2.25rem (36px)
- `text-5xl`: 3rem (48px)
- `text-6xl`: 3.75rem (60px)
- `text-7xl`: 4.5rem (72px)
- `text-8xl`: 6rem (96px)
- `text-9xl`: 8rem (128px)

### Font Weights
- `font-thin`: 100
- `font-extralight`: 200
- `font-light`: 300
- `font-normal`: 400
- `font-medium`: 500
- `font-semibold`: 600
- `font-bold`: 700
- `font-extrabold`: 800
- `font-black`: 900

### Text Colors
- `text-primary`: Primary brand color
- `text-secondary`: Secondary brand color
- `text-accent`: Accent color
- `text-neutral`: Neutral color
- `text-base-100`: Base color 100
- `text-base-200`: Base color 200
- `text-base-300`: Base color 300
- `text-info`: Information color
- `text-success`: Success color
- `text-warning`: Warning color
- `text-error`: Error color

### Text Alignment
- `text-left`: Left-aligned text
- `text-center`: Center-aligned text
- `text-right`: Right-aligned text
- `text-justify`: Justified text

### HTML Elements
The component supports various HTML elements:
- Headings: `h1`, `h2`, `h3`, `h4`, `h5`, `h6`
- Paragraphs: `p`
- Inline elements: `span`, `strong`, `em`, `small`, `code`
- Preformatted: `pre`

## HTML Structure

Typography elements are rendered as standard HTML elements with appropriate classes:

```html
<!-- Headings -->
<h1 class="text-4xl font-bold text-primary">Main Heading</h1>
<h2 class="text-3xl font-semibold text-secondary">Section Heading</h2>
<h3 class="text-2xl font-medium text-accent">Subsection Heading</h3>

<!-- Paragraphs -->
<p class="text-base text-neutral">Regular paragraph text</p>
<p class="text-sm text-base-300">Small paragraph text</p>

<!-- Inline elements -->
<span class="text-lg font-semibold text-info">Inline span</span>
<strong class="text-base font-bold text-success">Strong text</strong>
<em class="text-base italic text-warning">Emphasized text</em>
<code class="text-sm font-mono text-error">Code text</code>

<!-- Preformatted -->
<pre class="text-sm text-left bg-base-200 p-4 rounded">
Preformatted text
with line breaks
</pre>
```

## JavaScript Interaction API

Typography components are static text elements and do not require JavaScript interaction. However, you can dynamically update text content in WebAssembly (WASM) applications:

```go
// wasm/main.go
func updateTypographyContent() {
	doc := dom.GetWindow().Document()
	
	// Update heading text
	heading := doc.QuerySelector("h1")
	if heading != nil {
		heading.SetTextContent("Updated Heading Text")
	}
	
	// Update paragraph text
	paragraph := doc.QuerySelector("p")
	if paragraph != nil {
		paragraph.SetTextContent("This paragraph text was updated from WASM")
	}
	
	logutil.Log("Typography content updated")
}
```

## Examples

### Basic Typography

```html
<!-- Main heading -->
<h1 class="text-4xl font-bold text-primary mb-4">Welcome to FlyonUI</h1>

<!-- Section heading -->
<h2 class="text-3xl font-semibold text-secondary mb-3">Getting Started</h2>

<!-- Paragraph -->
<p class="text-base text-neutral leading-relaxed mb-4">
  FlyonUI is a powerful toolkit that combines semantic classes and interactive 
  headless JavaScript plugins, enabling developers to build stunning, 
  interactive user interfaces with ease.
</p>

<!-- Small text -->
<small class="text-sm text-base-300">Last updated: 2024</small>
```

### Typography with Different Sizes

```html
<div class="space-y-2">
  <h1 class="text-5xl font-bold">Heading 1 (5xl)</h1>
  <h2 class="text-4xl font-bold">Heading 2 (4xl)</h2>
  <h3 class="text-3xl font-bold">Heading 3 (3xl)</h3>
  <h4 class="text-2xl font-bold">Heading 4 (2xl)</h4>
  <h5 class="text-xl font-bold">Heading 5 (xl)</h5>
  <h6 class="text-lg font-bold">Heading 6 (lg)</h6>
  <p class="text-base">Paragraph (base)</p>
  <p class="text-sm">Small text (sm)</p>
  <p class="text-xs">Extra small text (xs)</p>
</div>
```

### Typography with Different Weights

```html
<div class="space-y-1">
  <p class="font-thin">Thin weight text</p>
  <p class="font-light">Light weight text</p>
  <p class="font-normal">Normal weight text</p>
  <p class="font-medium">Medium weight text</p>
  <p class="font-semibold">Semibold weight text</p>
  <p class="font-bold">Bold weight text</p>
  <p class="font-extrabold">Extrabold weight text</p>
  <p class="font-black">Black weight text</p>
</div>
```

### Typography with Colors

```html
<div class="space-y-2">
  <h3 class="text-2xl font-bold text-primary">Primary Color Heading</h3>
  <p class="text-base text-secondary">Secondary color paragraph</p>
  <p class="text-base text-accent">Accent color paragraph</p>
  <p class="text-base text-neutral">Neutral color paragraph</p>
  <p class="text-base text-info">Info color paragraph</p>
  <p class="text-base text-success">Success color paragraph</p>
  <p class="text-base text-warning">Warning color paragraph</p>
  <p class="text-base text-error">Error color paragraph</p>
</div>
```

### Code and Preformatted Text

```html
<!-- Inline code -->
<p>Here's some <code class="text-sm font-mono bg-base-200 px-2 py-1 rounded">inline code</code> in a paragraph.</p>

<!-- Code block -->
<pre class="text-sm text-left bg-base-200 p-4 rounded overflow-x-auto">
<code>// JavaScript example
function greet(name) {
  return `Hello, ${name}!`;
}

console.log(greet('World'));</code>
</pre>
```

### Responsive Typography

```html
<!-- Responsive heading sizes -->
<h1 class="text-2xl md:text-3xl lg:text-4xl font-bold text-primary">
  Responsive Heading
</h1>

<!-- Responsive paragraph -->
<p class="text-sm md:text-base lg:text-lg text-neutral">
  This paragraph text will adjust based on screen size.
</p>
```

### Typography with Alignment

```html
<!-- Centered heading -->
<h2 class="text-3xl font-bold text-center text-primary mb-4">
  Centered Heading
</h2>

<!-- Right-aligned text -->
<p class="text-base text-right text-neutral mb-4">
  Right-aligned paragraph text.
</p>

<!-- Justified text -->
<p class="text-base text-justify text-neutral">
  This text is justified, meaning it will be aligned to both the left and right margins,
  creating a clean and professional appearance for longer paragraphs of text.
</p>