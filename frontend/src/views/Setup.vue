<script lang="ts" setup>
import { ref } from 'vue'
import { GenerateSecretPhrase, SetSecretPhrase } from '../../wailsjs/go/backend/App'
import { useRouter } from 'vue-router'
const router = useRouter()

const secretPhrase = ref<string>("")
const data = ref<string[]>(Array<string>(10).fill(""))


function generateSecret() {
  GenerateSecretPhrase().then((result: string) => {
    secretPhrase.value = result
    data.value = result.split(" ").slice(0, 10)
  })
}

function nextPage() {
  secretPhrase.value = data.value.join(" ").trim()
  SetSecretPhrase(secretPhrase.value).then(() => {
    router.push("/home")
  })
}

</script>

<template>
  <main>

    <h1>Enter a phrase or generate a new one</h1>
    <div class="grid">
      <input v-for="(_, i) in data" type="text" :key="i" v-model="data[i]" />
    </div>

    <button @click="generateSecret">Generate</button>
    <button @click="nextPage" :disabled="!data.every(word => word.trim() !== '')">Next</button>
  </main>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 8px;
  margin-bottom: 16px;
  margin-top: 16px;
}

input {
  padding: 6px 10px;
  border: 1px solid #ccc;
  text-transform: lowercase;
  /* border-radius: 4px; */
}

button {
  margin: 5px;
  padding: 8px 14px;
  border: none;
  /* border-radius: 4px; */
  cursor: pointer;
}

button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}
</style>
