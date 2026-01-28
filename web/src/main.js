import './assets/main.css'
import { apiClient, dbClient, storageClient } from '@/lib'
import { createApp } from 'vue'
import { registerSW } from 'virtual:pwa-register'
import App from '@/App.vue'
import router from '@/router'

registerSW()
const app = createApp(App)
app.use(router)
app.provide('apiClient', apiClient)
app.provide('dbClient', dbClient)
app.provide('storageClient', storageClient)
app.mount('#app')
