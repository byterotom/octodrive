<script lang="ts" setup>
import { ref } from 'vue'
import { DiscoverDevices } from '../../wailsjs/go/backend/App'
import { EventsOff, EventsOn } from '../../wailsjs/runtime/runtime'

const ips = ref<string[]>([])
const disableButton = ref<boolean>(false)

function discover() {
  disableButton.value = true
  ips.value = []
  EventsOff("backend:discover")
  EventsOn("backend:discover", (src: string) => {
    ips.value.push(src)
  })
  DiscoverDevices().then(() => { }).finally(() => {
    disableButton.value = false
  })
}
</script>

<template>
  <div class="sidebar">
    <button class="discover-btn" @click="discover" :disabled="disableButton">
      <div v-if="disableButton">Discovering...</div>
      <div v-else>Discover</div>
    </button>

    <ul class="ip-list">
      <li v-for="ip in ips" :key="ip" class="ip-item">
        {{ ip }}
      </li>
    </ul>
  </div>
</template>

<style scoped>
.sidebar {
  position: fixed;
  top: 0;
  right: 0;
  width: 160px;
  height: 100vh;
  background: #1e1e2f;
  color: #f0f0f0;
  display: flex;
  flex-direction: column;
  padding: 16px;
  box-shadow: -2px 0 8px rgba(0, 0, 0, 0.3);
}

.discover-btn {
  background: #4cafef;
  border: none;
  padding: 10px 14px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  color: #fff;
  margin-bottom: 16px;
  transition: background 0.2s;
}

.discover-btn:hover {
  background: #3399dd;
}

.discover-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.ip-list {
  list-style: none;
  padding: 0;
  margin: 0;
  flex-grow: 1;
  overflow-y: auto;
}

.ip-item {
  background: #2c2c40;
  padding: 8px 12px;
  border-radius: 4px;
  margin-bottom: 8px;
  font-family: monospace;
  transition: background 0.2s;
}

.ip-item:hover {
  background: #3a3a55;
}
</style>
