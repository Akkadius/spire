import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

window.__vue3_router__ = router
createApp(App).use(router).mount('#app')
