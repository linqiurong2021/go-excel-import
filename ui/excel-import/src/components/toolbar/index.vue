<template>
 <div class="toolbar">
  <el-button icon="el-icon-plus" circle @click="newData">新增</el-button>
  <el-button icon="el-icon-delete" circle @click="deleteData" :disabled="delDisabled">删除</el-button>
  <el-button icon="el-icon-upload2" circle @click="importData">导入</el-button>
  <el-button icon="el-icon-download" circle @click="exportData">导出</el-button>
 </div>
</template>

<script>
import {Button} from "element-ui"
export default {
  name: "ToolBar",
  components: {
    ElButton: Button
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

    }
  },
  methods: {
    newData() {
      this.$emit('newData')
    },
      importData() {
      console.log('importData')
       this.$emit('importData')
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