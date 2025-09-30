<script lang="ts" setup>
import { SendFile } from '../../wailsjs/go/backend/App';
import Discover from '../components/Discover.vue';

function send() {
    SendFile("")
}

</script>
<template>
    <button @click="send">Send</button>
    <Discover />
</template>