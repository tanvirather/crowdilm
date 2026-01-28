import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import { VitePWA } from 'vite-plugin-pwa'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    VitePWA({
      registerType: 'autoUpdate',
      manifest: {
        name: 'Zuhid',
        short_name: 'Zuhid',
        start_url: '/',
        display: 'standalone',
        background_color: '#000000',
        theme_color: '#3c317aff',
        icons: [
          { src: '/favicon.ico', sizes: '192x192', type: 'image/png' },
          { src: '/favicon.ico', sizes: '512x512', type: 'image/png' }
        ]
      },
      includeAssets: [
        "fonts/Scheherazade-Regular.ttf"
      ]
    })
  ],
  server: {
    port: 5173
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
