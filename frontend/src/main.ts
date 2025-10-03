import { createApp } from 'vue'
import Toast, { PluginOptions, POSITION } from "vue-toastification"
import { EventsEmit, EventsOn } from '../wailsjs/runtime/runtime'
import App from './App.vue'
import router from './router'
import './style.css'
import "vue-toastification/dist/index.css";

// Setup event handlers for routes
EventsOn("backend:showHome", () => router.push("/home"))
EventsOn("backend:showSetup", () => router.push("/setup"))
EventsEmit("frontend:checkSecret")

// Create new app instance
const app = createApp(App)

// Use the router
app.use(router)

// Setting toast position at bottom right
const options: PluginOptions = {
    position: POSITION.BOTTOM_RIGHT
}
app.use(Toast, options)

// Mount the app
app.mount('#app')