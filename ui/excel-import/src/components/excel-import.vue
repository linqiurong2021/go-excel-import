<template>
 <div>
  <!--表单选择-->
  <el-select v-model="currentTemplate" @change="changeTemplate">
    <el-option
      v-for="item in templateList"
      :key="item.SYS_KEY"
      :label="item.SYS_NAME"
      :value="item.SYS_KEY">
    </el-option>
  </el-select>
  <!--工具栏-->
  <tool-bar @newData="newData" @importData="importData" @exportData="exportData" @deleteData="deleteData" :delDisabled="delDisabled" @refresh="refresh"/>
  <!--搜索工具-->
  <search :config="searchConfig" :type="fieldType" :label="labelName"/>
  <!--列表工具-->
  <list ref="list" :config="listConfig" :type="fieldType" :label="labelName" @selectedChange="rowSelectedChange" :isAdd="isAdd" @addChange="addChange"/>


</div>
</template>

<script lang="ts">
import List from "./list/index.vue"
import Search from "./search/index.vue"
import ToolBar from "./toolbar/index.vue"
import { mapGetters } from "vuex"
import {Select , Option, MessageBox, Message} from "element-ui"
import {getConfig, deleteBySysIDs, getTemplates,exportData, download} from "../api/excel-import.js"
export default {
  name: 'ExcelImport',
  components: {
    Search,List,
    ElSelect: Select,
    ElOption: Option,
    ToolBar
  },

  computed: {
    ...mapGetters({
      tableName: "form/getTableName",
      tableConfigs: "form/getTableConfigs"
    }),
    // 禁用按钮状态
    delDisabled() {
      // 未选中则禁用
      return this.selectedRows.length <=0 
    }
  },
  created() {
    this.getTemplates()
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
      console.log(newConfig,'newConfig')
      if (newConfig.length >= 6) {
        this.labelName = newConfig[0]
        this.fieldType = newConfig[3]
        this.searchConfig = newConfig[4]
        this.listConfig = newConfig[5]
    
        console.log(this.fieldType, this.searchConfig,this.listConfig,this.labelName)
      }else{
        console.log("table config error: length less than 4 items")
      }
    }
  },
  methods: {
    getTemplates() {
      getTemplates().then((res)=>{
        let {data} = res
        if (data) {
          this.templateList = data
          if(data[0] && data[0]['SYS_KEY']){
            this.$store.dispatch("form/setTableName", data[0]['SYS_KEY'])
          }
          
        }
        console.log(res,'getTemplates')
      })
    },
    changeTemplate(newVal, oldVal){
      // console.log(newVal,'newVal')
      if(newVal) {
        this.$store.dispatch('form/setTableName', newVal)
      }
    },
    addChange(val) {
      this.isAdd = val
    },
    newData() {
      this.isAdd =  true
    },
      importData() {
      console.log('importData')
     
    },
    download(data,fileName) {
      let url = window.URL.createObjectURL(new Blob([data]));
      let link = document.createElement('a');
      link.style.display = 'none';
      link.href = url;
      link.setAttribute('download', fileName);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    },
    exportData() {
      //
    
      let ids = []
      this.selectedRows.forEach(element => {
          ids.push(element.SYS_ID)
      });
      let data = {
        table: this.tableName,
        sys_ids: ids.join(",")
      }
      exportData(data).then((res)=>{
        var filename = 'exportData.xlsx'
        let contentDisposition = res.headers['content-disposition'];
        if(contentDisposition){
          let tmpArr = contentDisposition.split('=')
          // 删除 前后的 ""
          filename = tmpArr[1].substr(1,tmpArr[1].length-2)
        }
        // console.log(filename,'filename')
        download(res.data, filename)
      })
      console.log('exportData')
     
    },
    // 刷新列表
    refresh() {
      //
      this.refreshTable()
      // 刷新模板列表
      this.getTemplates()
    },
    // 
    refreshTable() {
      this.$refs.list.getListByPage()
    },
    deleteData() {
      //
      if (this.selectedRows.length <=0) {
        return
      }
      let ids = []
      this.selectedRows.forEach(element => {
          ids.push(element.SYS_ID)
      });
      // 提示是否需要删除
      MessageBox.confirm('删除后不可恢复,是否删除？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '删除',
        cancelButtonText: '取消'
      }).then(() => {
        //
        let params = {
          table: this.tableName,
          sys_ids: ids.join(",")
        }
        deleteBySysIDs(params).then((res)=> {
          let {data} = res
          if (data > 0) {
            // 删除成功
            Message({
              type: 'success',
              message: "删除成功"
            })
            // 刷新列表
            this.refreshTable()
            // console.log("删除成功")
          }else{

            console.error(res)
          }
        })
        console.log("删除了")
      })
      .catch(action => {
        console.log("取消了",action)
      });

      console.log("deleteData")
    },
    //
    rowSelectedChange(val) {
      this.selectedRows = val
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
      ],
      // checked 选中的项
      selectedRows: [],
      // 是否新增操作
      isAdd: false
    }
  }
}
</script>
