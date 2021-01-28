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

export default {
  name: 'ExcelImport',
  components: {
    Search,List
  },
  created() {
    let tableName = "TemplateTest"
    console.error(tableName,'tableName')
    getConfig({table: tableName}).then((res)=>{
      let data = res.data
      if (data.length >= 3) {
        this.fieldType = data[0]
        this.searchConfig = data[1]
        this.listConfig = data[2]
        this.labelName = data[3]
      }else{
        console.error(res,'res')
      }
    })
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
