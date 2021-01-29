<template>
 <div>
  <search :config="searchConfig" :type="fieldType" :label="labelName"/>
  <list :config="listConfig" :type="fieldType" :label="labelName"/>
</div>
</template>

<script lang="ts">
import List from "./list/index.vue"
import Search from "./search/index.vue"

import {getConfig} from '../api/excel-import.js'
import { mapGetters } from "vuex"
export default {
  name: 'ExcelImport',
  components: {
    Search,List
  },
  computed: {
    ...mapGetters({
      tableName: "form/getTableName",
      tableConfigs: "form/getTableConfigs"
    })
  },
  created() {
    let tableName = "TemplateTest"
    this.$store.dispatch("form/setTableName", tableName)
    console.error(tableName,'tableName')
  },
  watch: {
    tableName(newName,oldName){
      if(newName == "") return
      // 
      getConfig({table: newName}).then((res)=>{
        let data = res.data
        this.$store.dispatch('form/setTableConfigs', res.data)
      })
    },
    tableConfigs(newConfig, oldConfig){
      if (newConfig.length >= 3) {
        this.fieldType = newConfig[0]
        this.searchConfig = newConfig[1]
        this.listConfig = newConfig[2]
        this.labelName = newConfig[3]
      }else{
        console.log("table config error: length less than 4 items")
      }
    }
  },
  
  data() {
    return {
      searchConfig: {},
      listConfig: {},
      fieldType: {},
      labelName: {},
    }
  }
}
</script>
