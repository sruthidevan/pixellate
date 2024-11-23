# pixellate-frontend

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
```
pixellate-frontend/
├── public/                # Static files (e.g., favicon, images)
│   ├── favicon.ico        # Favicon
│   └── index.html         # Root HTML template
├── src/                   # Main application source
│   ├── assets/            # Images, fonts, CSS files
│   │   └── styles/        # Global styles (optional)
│   ├── components/        # Reusable Vue components
│   │   └── Button.vue     # Example button component
│   ├── views/             # Page components for routing
│   │   ├── Home.vue       # Home page
│   │   ├── About.vue      # About page
│   │   ├── Contact.vue    # Contact page
│   │   └── Services.vue   # Services page
│   ├── router/            # Vue Router configuration
│   │   └── index.ts       # Router setup
│   ├── store/             # State management (e.g., Pinia or Vuex)
│   │   └── index.ts       # Store configuration (optional)
│   ├── App.vue            # Root Vue component
│   ├── main.ts            # Entry point for the app
│   └── index.css          # Tailwind or global styles
├── .gitignore             # Files to ignore in Git
├── package.json           # Project metadata and dependencies
├── tailwind.config.js     # Tailwind CSS configuration
├── postcss.config.js      # PostCSS configuration (for Tailwind)
└── vite.config.ts         # Vite configuration
```