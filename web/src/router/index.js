import { createRouter, createWebHistory } from 'vue-router'
import QuranView from '../views/QuranView.vue'
import SettingView from '../views/SettingView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: QuranView },
    { path: '/Quran', name: 'Quran', component: QuranView },
    { path: '/Setting', name: 'Setting', component: SettingView },
  ],
})

export default router
