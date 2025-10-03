import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Home from './views/Home.vue'
import Setup from './views/Setup.vue'

// Route to component mapping
const routes: Array<RouteRecordRaw> = [
    { path: '/home', component: Home },
    { path: '/setup', component: Setup },
]

// New router instance
const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router