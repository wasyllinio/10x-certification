# 10xCertification

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 20.3.5.

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Karma](https://karma-runner.github.io) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Styling System

This project uses **PrimeNG** and **PrimeFlex** for UI components and styling:

### PrimeNG (v20.2.0)
- Comprehensive component library with 100+ components
- Fully configured and themed in `app.config.ts`
- CSS-in-JS theming system with Aura preset

### PrimeFlex (v4.0.0)
- Utility-first CSS framework
- Provides spacing, layout, and utility classes
- Imported in `styles.css` and ready to use

### Usage Examples

**PrimeNG Component:**
```html
<p-button label="Click me" icon="pi pi-check"></p-button>
```

**PrimeFlex Utilities:**
```html
<!-- Spacing -->
<div class="p-m-3 p-p-4">Margins and padding</div>

<!-- Flexbox Layout -->
<div class="flex align-items-center justify-content-between gap-3">
  <span>Left content</span>
  <span>Right content</span>
</div>

<!-- Responsive -->
<div class="grid">
  <div class="col-12 md:col-6 lg:col-4">Content</div>
</div>
```

### Custom Design System

Additional custom utilities are available in:
- `src/app/shared/styles/design-system.css` - Typography, spacing, elevation
- `src/app/shared/styles/cards.css` - Card components
- `src/app/shared/styles/components.css` - Enhanced PrimeNG components
- `src/app/shared/styles/utilities.css` - Additional utilities

## Additional Resources

- [Angular CLI Documentation](https://angular.dev/tools/cli)
- [PrimeNG Documentation](https://primeng.org/)
- [PrimeFlex Documentation](https://primeflex.org/)
