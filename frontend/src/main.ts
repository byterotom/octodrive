import { createApp } from 'vue'
import Toast, { PluginOptions, POSITION } from "vue-toastification"
import { EventsEmit, EventsOn } from '../wailsjs/runtime/runtime'
import App from './App.vue'
import router from './router'
import './style.css'
import "vue-toastification/dist/index.css";

EventsOn("backend:showHome", () => router.push("/home"))
EventsOn("backend:showSetup", () => router.push("/setup"))
EventsEmit("frontend:checkSecret")

const app = createApp(App)

app.use(router)

const options: PluginOptions = {
    position: POSITION.BOTTOM_RIGHT
}
app.use(Toast, options)

app.mount('#app')