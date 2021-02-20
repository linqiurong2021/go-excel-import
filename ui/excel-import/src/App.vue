<template>
  <div id="app">
    <!-- <ExcelImport /> -->
    <!-- <Map /> -->
    <component :is="componentName" ></component>
  </div>
</template>

<script>
import ExcelImport from "./components/excel-import.vue"
import EginMap from "./components/map.vue"
export default {
  name: 'App',
  components: {
    ExcelImport,
    EginMap
  },
  data() {
    return {
      componentName: 'EginMap'
    }
  }
}
</script>

<style>
#app {

}
</style>
