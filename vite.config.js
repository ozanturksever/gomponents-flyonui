import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [
    tailwindcss()
  ],
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    manifest: true,
    rollupOptions: {
      input: {
        main: './assets/main.js'
      }
    }
  },
  server: {
    port: 3000,
    open: true
  }
})