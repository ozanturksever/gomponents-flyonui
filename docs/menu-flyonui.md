# Menu - FlyonUI

The Menu component provides a flexible navigation system that can be used for sidebars, navigation bars, and hierarchical navigation structures. It supports vertical and horizontal layouts, nested sub-menus, icons, badges, and various styling options.

## Classes

| Class | Type | Description |
|-------|------|-------------|
| `menu` | Component | Base class for menu container. |
| `menu-title` | Part | Base class for menu title. |
| `menu-disabled` | State | Makes menu item (`li`) disabled. |
| `menu-active` | State | Makes element inside menu item (`li`) active. |
| `menu-xs` | Size | Extra small size. |
| `menu-sm` | Size | Small size. |
| `menu-md` | Size | Medium (default) size. |
| `menu-lg` | Size | Large size. |
| `menu-xl` | Size | Extra-Large size. |
| `menu-vertical` | Direction | Vertical menu (default). |
| `menu-horizontal` | Direction | Horizontal menu. |

## Basic Usage

### Simple Vertical Menu

```html
<ul class="menu">
  <li><a href="#">Home</a></li>
  <li><a href="#">Account</a></li>
  <li><a href="#">Notifications</a></li>
</ul>
```

### Horizontal Menu

```html
<ul class="menu menu-horizontal">
  <li><a href="#">Home</a></li>
  <li><a href="#">Account</a></li>
  <li><a href="#">Notifications</a></li>
</ul>
```

### Responsive Horizontal Menu

```html
<ul class="menu sm:menu-horizontal">
  <li><a href="#">Home</a></li>
  <li><a href="#">Account</a></li>
  <li><a href="#">Notifications</a></li>
</ul>
```

## Menu Sizes

### Available Sizes

```html
<ul class="menu menu-xs w-fit">
  <li><a href="#">Extra Small</a></li>
</ul>

<ul class="menu menu-sm w-fit">
  <li><a href="#">Small</a></li>
</ul>

<ul class="menu menu-md w-fit">
  <li><a href="#">Medium</a></li>
</ul>

<ul class="menu menu-lg w-fit">
  <li><a href="#">Large</a></li>
</ul>

<ul class="menu menu-xl w-fit">
  <li><a href="#">Extra Large</a></li>
</ul>
```

## Menu with Icons

### Icons with Text

```html
<ul class="menu">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--user] size-5"></span>
      Account
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--message] size-5"></span>
      Notifications
    </a>
  </li>
</ul>
```

### Icon-Only Menu

```html
<ul class="menu">
  <li>
    <a href="#" aria-label="Home Link">
      <span class="icon-[tabler--home] size-5"></span>
    </a>
  </li>
  <li>
    <a href="#" aria-label="User Link">
      <span class="icon-[tabler--user] size-5"></span>
    </a>
  </li>
  <li>
    <a href="#" aria-label="Message Link">
      <span class="icon-[tabler--message] size-5"></span>
    </a>
  </li>
</ul>
```

### Horizontal Icon-Only Menu

```html
<ul class="menu menu-horizontal">
  <li>
    <a href="#" aria-label="Home Link">
      <span class="icon-[tabler--home] size-5"></span>
    </a>
  </li>
  <li>
    <a href="#" aria-label="User Link">
      <span class="icon-[tabler--user] size-5"></span>
    </a>
  </li>
  <li>
    <a href="#" aria-label="Message Link">
      <span class="icon-[tabler--message] size-5"></span>
    </a>
  </li>
</ul>
```

## Menu with Badges

```html
<ul class="menu lg:menu-horizontal">
  <li>
    <a href="#">
      <span class="icon-[tabler--mail] size-5"></span>
      Inbox
      <span class="badge badge-sm badge-primary">1K+</span>
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--info-circle] size-5"></span>
      Updates
      <span class="badge badge-sm badge-warning">NEW</span>
    </a>
  </li>
  <li>
    <a href="#">
      Status
      <span class="badge badge-success size-3 p-0"></span>
    </a>
  </li>
</ul>
```

## Menu Titles

### Simple Title

```html
<ul class="menu">
  <li class="menu-title">Apps</li>
  <li>
    <a href="#">
      <span class="icon-[tabler--message] size-5"></span>
      Chat
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--calendar] size-5"></span>
      Calendar
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--book] size-5"></span>
      Academy
    </a>
  </li>
</ul>
```

### Title as Parent with Sub-menu

```html
<ul class="menu">
  <li>
    <p class="menu-title">Apps</p>
    <ul>
      <li>
        <a href="#">
          <span class="icon-[tabler--message] size-5"></span>
          Chat
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--calendar] size-5"></span>
          Calendar
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--book] size-5"></span>
          Academy
        </a>
      </li>
    </ul>
  </li>
</ul>
```

## Nested Sub-menus

### Static Sub-menus

```html
<ul class="menu">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--apps] size-5"></span>
      Apps
    </a>
    <ul class="menu">
      <li>
        <a href="#">
          <span class="icon-[tabler--message] size-5"></span>
          Chat
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--calendar] size-5"></span>
          Calendar
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--book] size-5"></span>
          Academy
        </a>
        <ul class="menu">
          <li>
            <a href="#">
              <span class="icon-[tabler--books] size-5"></span>
              Courses
            </a>
          </li>
          <li>
            <a href="#">
              <span class="icon-[tabler--list-details] size-5"></span>
              Course details
            </a>
          </li>
        </ul>
      </li>
    </ul>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--settings] size-5"></span>
      Settings
    </a>
  </li>
</ul>
```

### Horizontal Menu with Sub-menu

```html
<ul class="menu sm:menu-horizontal">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--apps] size-5"></span>
      Apps
    </a>
    <ul class="menu">
      <li>
        <a href="#">
          <span class="icon-[tabler--message] size-5"></span>
          Chat
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--calendar] size-5"></span>
          Calendar
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--book] size-5"></span>
          Academy
        </a>
      </li>
    </ul>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--settings] size-5"></span>
      Settings
    </a>
  </li>
</ul>
```

## Collapsible Sub-menus

### Basic Collapsible Menu

```html
<ul class="menu w-64 space-y-0.5">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li class="space-y-0.5">
    <a class="collapse-toggle collapse-open:bg-base-content/10 open" id="menu-app" data-collapse="#menu-app-collapse">
      <span class="icon-[tabler--apps] size-5"></span>
      Apps
      <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 size-4 transition-all duration-300"></span>
    </a>
    <ul id="menu-app-collapse" class="open collapse w-auto space-y-0.5 overflow-hidden transition-[height] duration-300" aria-labelledby="menu-app">
      <li>
        <a href="#">
          <span class="icon-[tabler--message] size-5"></span>
          Chat
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--calendar] size-5"></span>
          Calendar
        </a>
      </li>
    </ul>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--settings] size-5"></span>
      Settings
    </a>
  </li>
</ul>
```

### Multi-level Collapsible Menu

```html
<ul class="menu w-64 space-y-0.5">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li class="space-y-0.5">
    <a class="collapse-toggle collapse-open:bg-base-content/10 open" id="menu-app" data-collapse="#menu-app-collapse">
      <span class="icon-[tabler--apps] size-5"></span>
      Apps
      <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 size-4 transition-all duration-300"></span>
    </a>
    <ul id="menu-app-collapse" class="open collapse w-auto space-y-0.5 overflow-hidden transition-[height] duration-300" aria-labelledby="menu-app">
      <li>
        <a href="#">
          <span class="icon-[tabler--message] size-5"></span>
          Chat
        </a>
      </li>
      <li>
        <a href="#">
          <span class="icon-[tabler--calendar] size-5"></span>
          Calendar
        </a>
      </li>
      <li class="space-y-0.5">
        <a class="collapse-toggle collapse-open:bg-base-content/10 open" id="sub-menu-academy" data-collapse="#sub-menu-academy-collapse">
          <span class="icon-[tabler--book] size-5"></span>
          Academy
          <span class="icon-[tabler--chevron-down] collapse-open:rotate-180 size-4 transition-all duration-300"></span>
        </a>
        <ul id="sub-menu-academy-collapse" class="open collapse w-auto space-y-0.5 overflow-hidden transition-[height] duration-300" aria-labelledby="sub-menu-academy">
          <li>
            <a href="#">
              <span class="icon-[tabler--books] size-5"></span>
              Courses
            </a>
          </li>
          <li>
            <a href="#">
              <span class="icon-[tabler--list-details] size-5"></span>
              Course details
            </a>
          </li>
        </ul>
      </li>
    </ul>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--settings] size-5"></span>
      Settings
    </a>
  </li>
</ul>
```

## Mega Menu

### Basic Mega Menu

```html
<ul class="menu sm:menu-horizontal">
  <li>
    <span class="menu-title">Services</span>
    <ul class="menu">
      <li><a href="#">Design Solutions</a></li>
      <li><a href="#">Software Development</a></li>
      <li><a href="#">Web Hosting</a></li>
      <li><a href="#">Domain Registration</a></li>
    </ul>
  </li>
  <li>
    <span class="menu-title">Corporate Solutions</span>
    <ul class="menu">
      <li><a href="#">CRM</a></li>
      <li><a href="#">Management Solutions</a></li>
      <li><a href="#">Security Services</a></li>
      <li><a href="#">Consulting Services</a></li>
    </ul>
  </li>
  <li>
    <span class="menu-title">Product Offerings</span>
    <ul class="menu">
      <li><a href="#">UI Kits</a></li>
      <li><a href="#">Component Library</a></li>
      <li><a href="#">WordPress Plugins</a></li>
      <li>
        <span class="menu-title">Open Source Projects</span>
        <ul class="menu">
          <li><a href="#">Authentication System</a></li>
          <li><a href="#">FlyonUI Theme</a></li>
        </ul>
      </li>
    </ul>
  </li>
</ul>
```

## Menu States

### Active Item

```html
<ul class="menu">
  <li>
    <a href="#" class="menu-active">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--user] size-5"></span>
      Account
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--message] size-5"></span>
      Notifications
    </a>
  </li>
</ul>
```

### Disabled Item

```html
<ul class="menu">
  <li>
    <a href="#">
      <span class="icon-[tabler--home] size-5"></span>
      Home
    </a>
  </li>
  <li>
    <a href="#">
      <span class="icon-[tabler--settings] size-5"></span>
      Settings
    </a>
  </li>
  <li class="menu-disabled">
    <a href="#">
      <span class="icon-[tabler--lock] size-5"></span>
      Security
    </a>
  </li>
</ul>
```

## Menu with Tooltips

```html
<ul class="menu">
  <li class="tooltip [--placement:right]">
    <a href="#" class="tooltip-toggle" aria-label="Home Link">
      <span class="icon-[tabler--home] size-5"></span>
    </a>
    <span class="tooltip-content tooltip-shown:opacity-100 tooltip-shown:visible" role="tooltip">
      <span class="tooltip-body">Home</span>
    </span>
  </li>
  <li class="tooltip [--placement:right]">
    <a href="#" class="tooltip-toggle" aria-label="User Link">
      <span class="icon-[tabler--user] size-5"></span>
    </a>
    <span class="tooltip-content tooltip-shown:opacity-100 tooltip-shown:visible" role="tooltip">
      <span class="tooltip-body">Account</span>
    </span>
  </li>
  <li class="tooltip [--placement:right]">
    <a href="#" class="tooltip-toggle" aria-label="Message Link">
      <span class="icon-[tabler--message] size-5"></span>
    </a>
    <span class="tooltip-content tooltip-shown:opacity-100 tooltip-shown:visible" role="tooltip">
      <span class="tooltip-body">Notifications</span>
    </span>
  </li>
</ul>
```

## Simple Menu Without Styling

```html
<ul class="menu rounded-none p-0 [&_li>*]:rounded-none">
  <li><a href="#">Home</a></li>
  <li><a href="#">Account</a></li>
  <li><a href="#">Notifications</a></li>
</ul>
```

## Dropdown Integration

### Menu with Dropdown

```html
<ul class="menu menu-horizontal space-x-0.5">
  <li><a href="#">Home</a></li>
  <li><a href="#">Services</a></li>
  <li class="dropdown relative inline-flex [--auto-close:inside] [--offset:9] [--placement:bottom-end]">
    <button id="dropdown-end" type="button" class="dropdown-toggle dropdown-open:bg-base-content/10 dropdown-open:text-base-content max-sm:px-2" aria-haspopup="menu" aria-expanded="false" aria-label="Dropdown">
      Products
      <span class="icon-[tabler--chevron-down] dropdown-open:rotate-180 size-4"></span>
    </button>
    <ul class="dropdown-menu dropdown-open:opacity-100 hidden" role="menu" aria-orientation="vertical" aria-labelledby="dropdown-end">
      <li><a class="dropdown-item" href="#">UI kits</a></li>
      <li><a class="dropdown-item" href="#">Templates</a></li>
      <li><a class="dropdown-item" href="#">Component library</a></li>
      <hr class="border-base-content/25 -mx-2 my-3" />
      <li><a class="dropdown-item" href="#">Figma designs</a></li>
    </ul>
  </li>
</ul>
```

## Accessibility

The menu component includes proper ARIA attributes for accessibility:

- Use `aria-label` for icon-only menu items
- Include `aria-expanded` for collapsible menus
- Use semantic HTML with `<ul>` and `<li>` elements
- Support for keyboard navigation
- Proper focus management

## Responsive Design

The menu component is fully responsive:

- Uses responsive prefixes like `sm:`, `md:`, `lg:` for different screen sizes
- Supports horizontal scrolling for long menus
- Adapts to different screen sizes automatically
- Mobile-friendly touch targets