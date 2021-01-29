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
    <el-table-column
      fixed="right"
      label="操作"
      width="100">
      <template slot-scope="scope">
        <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
        <el-button type="text" size="small">编辑</el-button>
      </template>
    </el-table-column>
  </el-table>
    <el-pagination
      small
      layout="total, prev, pager, next"
      :page-size.sync="pageSize"
      :current-page.sync="currentPage"
      :total="total">
    </el-pagination>
 </div>
</template>

<script lang="ts">
import {Table, TableColumn, Pagination, Button} from "element-ui"
import {ListField} from '../../utils/const.js'
import {getListByPage} from '../../api/excel-import.js'
import { mapGetters } from 'vuex'
 export default {
  name: "List",
  components: {
    ElTable: Table,
    ElTableColumn : TableColumn,
    ElPagination: Pagination,
    ElButton: Button
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
        // console.log(newParams,'newParams')
        // console.log(newParams,'newParams')
        this.newParams = newParams
        let defaultParams = {
          page: this.currentPage,
          page_size: this.pageSize,
          table: this.tableName,
          params: this.newParams,
        }
        this.getListByPage(defaultParams)
        // console.log("end")
      }
    },
    currentPage(newVal, oldVal) {
      let defaultParams = {
        page: this.currentPage,
        page_size: this.pageSize,
        table: this.tableName,
        params: this.newParams,
      }
      this.getListByPage(defaultParams)
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
      tableName: "form/getTableName"
    })
  },
  methods: {
    handleClick(row) {
      console.log(row)
    },
    tableRowClassName({row, rowIndex}) {
      if (rowIndex === 1) {
        return 'warning-row';
      } else if (rowIndex === 3) {
        return 'success-row';
      }
      return '';
    },
    // 获取分页数据内容
    getListByPage(params) {
      // 获取列表数据
      getListByPage(params).then((res)=>{
        let data = res.data
        if (data.total) {
          this.total = data.total
        }
        // console.log(data,'data')
        if (data.list || data.list == null) {
          let list = data.list == null ? [] : data.list
          this.tableData = list
        }
        
      })
    }
  },
  data() {
    return {
      // 字段
      ListField: ListField,
      // 名称
      labelNames: [],
      tableData: [],
      total: 0,
      currentPage: 1,
      pageSize: 5,
      // 搜索参数
      newParams: ''
    }
  },
  created() {
    // 默认获取全部数据
    let defaultParams = {
      "page": this.currentPage,
      "page_size": this.pageSize,
      "table": this.tableName,
      "params": ""
    }
    this.getListByPage(defaultParams)
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