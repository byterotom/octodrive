<script lang="ts" setup>
import {reactive} from 'vue'
import {GenerateSecretPhrase, SetSecretPhrase} from '../../wailsjs/go/server/App'
import { useRouter } from 'vue-router'
const router = useRouter()

const data = reactive<{secretPhrase:string}>({
  secretPhrase : ""
})


function generateSecret() {
  GenerateSecretPhrase().then((result:string) => {
    data.secretPhrase = result
  })
}

function setSecret(){
  SetSecretPhrase(data.secretPhrase).then(()=>{})
}

function nextPage(){
  router.push("/home")
}

</script>

<template>
  <main>
    <div id="result" >{{ data.secretPhrase }}</div>
    <button @click="generateSecret">Generate</button>
    <br>
    --OR--
    <br>  
    <input type="text" v-model="data.secretPhrase" placeholder="Enter your secret phrase"/>
    <button @click="setSecret">Set</button>
    
    <button @click="nextPage">Next</button>
  </main>
</template>

<style scoped>
</style>
