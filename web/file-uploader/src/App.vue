<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <FileUploadDragAndDrop/>
</template>

<script>
import FileUploadDragAndDrop from "@/FileUploadDragAndDrop.vue";


export default {
  name: 'App',
  components: {
    FileUploadDragAndDrop,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
