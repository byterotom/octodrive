import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import router from './router'
import { EventsEmit, EventsOn } from '../wailsjs/runtime/runtime'

EventsOn("backend:showHome", () => router.push("/home"))
EventsOn("backend:showSetup", () => router.push("/setup"))
EventsEmit("frontend:checkSecret")

const app = createApp(App)

app.use(router).mount('#app')
