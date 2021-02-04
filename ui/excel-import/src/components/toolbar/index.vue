<template>
 <div class="toolbar">
  <el-button icon="el-icon-plus" circle @click="newData">新增</el-button>
 
  <el-button icon="el-icon-delete" circle @click="deleteData" :disabled="delDisabled">删除</el-button>
  <el-upload
    class="upload-demo"
    action="noaction"
    :show-file-list="false"
    :before-upload="beforeUpload"
    >
    <el-button icon="el-icon-upload2" circle>导入</el-button>
  </el-upload>
  <el-button icon="el-icon-download" circle @click="exportData">导出</el-button>
 </div>
</template>

<script>
import {Button, Upload, Message} from "element-ui"
import { importData } from '../../api/excel-import';

export default {
  name: "ToolBar",
  components: {
    ElButton: Button,
    ElUpload: Upload
  },
  props: {
    // 删除按钮状态 默认为禁用
    delDisabled: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      uploadApiUrl: '/import',
      uploadAccept: '.csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' // .csv .xls .xlsx
    }
  },
  methods: {
    newData() {
      this.$emit('newData')
    },
    beforeUpload(file) {
      let fd = new FormData()
      fd.append('file', file)
      //
      importData(fd).then((res)=>{
        console.log(res)
        let {data} = res
        if (data == 'success') {
          Message.success("导入成功")
          this.$emit('refresh')
        }else{
          Message.error(data)
        }
      })
      return false
    },
    
    exportData() {
      console.log('exportData')
       this.$emit('exportData')
    },
    deleteData() {
      this.$emit('deleteData')
    }
  }
}
</script>

<style scoped>

</style>