import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import router from './router'
import { EventsEmit, EventsOn } from '../wailsjs/runtime/runtime'

EventsOn("server:showHome", () => router.push("/home"))
EventsOn("server:showSetup", () => router.push("/setup"))
EventsEmit("frontend:checkSecret")

const app = createApp(App)

app.use(router).mount('#app')
