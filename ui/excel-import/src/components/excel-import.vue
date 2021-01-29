<template>
 <div>
  <el-select v-model="currentTemplate" @change="changeTemplate">
    <el-option
      v-for="item in templateList"
      :key="item.NAME"
      :label="item.CNAME"
      :value="item.NAME">
    </el-option>
  </el-select>
  <search :config="searchConfig" :type="fieldType" :label="labelName"/>
  <list :config="listConfig" :type="fieldType" :label="labelName"/>
</div>
</template>

<script lang="ts">
import List from "./list/index.vue"
import Search from "./search/index.vue"

import {getConfig} from '../api/excel-import.js'
import { mapGetters } from "vuex"
import {Select , Option} from "element-ui"
export default {
  name: 'ExcelImport',
  components: {
    Search,List,
    ElSelect: Select,
    ElOption: Option,
  },

  computed: {
    ...mapGetters({
      tableName: "form/getTableName",
      tableConfigs: "form/getTableConfigs"
    })
  },
  created() {
    let tableName = "TemplateTest2"
    this.$store.dispatch("form/setTableName", tableName)
    console.error(tableName,'tableName')
  },
  watch: {
    tableName(newName,oldName){
      if(newName == "") return
      // 
      this.currentTemplate = newName
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
        console.log(this.fieldType, this.searchConfig,this.listConfig,this.labelName)
      }else{
        console.log("table config error: length less than 4 items")
      }
    }
  },
  methods: {
    changeTemplate(newVal, oldVal){
      // console.log(newVal,'newVal')
      if(newVal) {
        this.$store.dispatch('form/setTableName', newVal)
      }
    }
  },
  data() {
    return {
      searchConfig: {},
      listConfig: {},
      fieldType: {},
      labelName: {},
      //
      currentTemplate: '',
      // 
      templateList: [
        {"NAME":"TemplateTest","CNAME":"模板测试1","ORDER":1},
        {"NAME":"TemplateTest2","CNAME":"模板测试2","ORDER":2},
        {"NAME":"TemplateTest3","CNAME":"模板测试3","ORDER":3}
      ]
    }
  }
}
</script>
