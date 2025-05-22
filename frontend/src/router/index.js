import { createRouter, createWebHistory } from 'vue-router'
import ModernHome from '@/components/HelloWorld.vue'

export default createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/items', component: ModernHome },
        // Add legacy fallback, for example:
        { path: '/:catchAll(.*)', component: () => import('@/components/LegacyMain.vue') }
    ]
})
