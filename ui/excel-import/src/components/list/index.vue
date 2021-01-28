<template>
 <div>
   <el-table
    :data="tableData"
    style="width: 100%"
    :row-class-name="tableRowClassName">
    <el-table-column
      v-for="(item,index) in listKeys"
      :key="index"
      :prop="item"
      :label="labelNames[index]"
    >
    </el-table-column>
  </el-table>
  
 </div>
</template>

<script lang="ts">
import {Table, TableColumn} from "element-ui"
import {ListField} from '../../utils/const.js'
import {getTableListByPage} from '../../api/excel-import.js'
import { mapGetters } from 'vuex'
 export default {
  name: "List",
  components: {
    ElTable: Table,
    ElTableColumn : TableColumn
  },

  props: {
    config: {
      type: Object,
      default() {
        return {}
      }
    },
    label: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  watch: {
    config(newConf, oldConf){
      console.log(newConf,'newConf')
    },
    getSearchParams: { 
      handler: function(newParams, oldParams) {
        console.log(newParams, oldParams,'getSearchParams')
      }
    }
  },
  computed: {
        // 所有KEY
    configKeys () {
      return Object.keys(this.config)
    },
    // 所有值
    configValues() {
      return Object.values(this.config)
    },
    // 所有值
    labelValues() {
      return Object.values(this.label)
    },

    // 列表key
    listKeys() {
      let keys = []
      let labels = []
      this.configValues.forEach((item,index)=>{
        if(item == this.ListField.VALUE) {
          labels.push(this.labelValues[index])
          keys.push(this.configKeys[index])
        }
      })
      this.labelNames = labels
      return keys
    },
    ...mapGetters({
      getSearchParams: 'form/getSearchParams',
    })
  },
  methods: {
    tableRowClassName({row, rowIndex}) {
      if (rowIndex === 1) {
        return 'warning-row';
      } else if (rowIndex === 3) {
        return 'success-row';
      }
      return '';
    },
    // 获取分页数据内容
    getTableListByPage() {
      let params = {
        page: 1,
        pageSize: 10,
        params: "name=''&id=''"
      }
      // 获取列表数据
      getTableListByPage(params).then((res)=>{
        console.log(res,'res')
      })
    }
  },
  data() {
    return {
      // 字段
      ListField: ListField,
      // 名称
      labelNames: [],
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄',
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄',
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }]
    }
  },
  mounted() {
    this.getTableListByPage()
  }
 }
</script>

<style scoped>
.el-table .warning-row {
  background: rgb(238, 194, 113);
}
.el-table .success-row {
  background: #bef3a2;
}
</style>