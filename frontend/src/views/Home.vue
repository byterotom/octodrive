<script lang="ts" setup>
import { SelectFilesToUpload, UploadFile } from '../../wailsjs/go/backend/App';
import Discover from '../components/Discover.vue';

async function upload() {
    try {
        const files = await SelectFilesToUpload()
        for (const file of files) {
            uploadSingle(file)
        }
    } catch (err: unknown) {

    } finally { }
}

async function uploadSingle(file: string) {
    try {
        await UploadFile(file)
    } catch (err: unknown) {

    } finally { }
}

</script>
<template>
    <button @click="upload">Select</button>
    <Discover />
</template>