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
        <el-button @click="handleDelete(scope.row)" type="text" size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
    <el-pagination
      small
      layout="total, prev, pager, next"
      :page-size.sync="pageSize"
      :current-page.sync="currentPage"
      :total.sync="total">
    </el-pagination>

    <DetailDialog :types="fieldsType" :detail="detailData" :names="fieldsName" ref="dialog"/>
 </div>
</template>

<script>
import {Table, TableColumn, Pagination, Button} from "element-ui"
import {ListField} from '../../utils/const.js'
import {getListByPage, getDataByID, getFieldsType,getFieldsName} from '../../api/excel-import.js'
import { mapGetters } from 'vuex'

import DetailDialog from "../dialog/index"

 export default {
  name: "List",
  components: {
    DetailDialog,
    ElTable: Table,
    ElTableColumn : TableColumn,
    ElPagination: Pagination,
    ElButton: Button,
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
    // 模板切换
    tableName() {
      this.newParams = ""
      
      this.getListByPage()
    },
    getSearchParams: { 
      handler: function(newParams, oldParams) {
        this.newParams = newParams
        this.getListByPage()
        // console.log("end")
      }
    },
    currentPage(newVal, oldVal) {
      this.currentPage = newVal
      this.getListByPage()
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
      let params = {
        table: this.tableName,
        sysID: row.SYS_ID
      }
      let tableParams= {table: this.tableName}
      getFieldsType(tableParams).then((res)=>{
        let {data} = res
        this.fieldsType = data
        this.$refs.dialog.dialogVisible = true
        console.log(data,"fieldsType")
      })
      //
      getDataByID(params).then((res)=>{
        let {data} = res
        this.detailData = data
        this.$refs.dialog.dialogVisible = true
         console.log(data,'res')
      })
      // 
      getFieldsName(tableParams).then((res)=>{
        let {data} = res
        this.fieldsName = data
        console.log(data,'names')
        this.$refs.dialog.dialogVisible = true
      })
    },
    handleDelete(row) {
      console.log(row, 'delete')
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
    getListByPage() {
      let params = {
        page: this.currentPage,
        page_size: this.pageSize,
        table: this.tableName,
        params: this.newParams,
      }
      // 获取列表数据
      getListByPage(params).then((res)=>{
        let data = res.data
        console.log(data,'data')
        if (data.list || data.list == null) {
          let list = data.list == null ? [] : data.list
          this.tableData = list
          this.total = data.list == null ? 0 : data.total
          console.log(this.total, 'total')
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
      newParams: '',
      //
      fieldsType: {},
      //
      fieldsName: {},
      //
      detailData: {},

    }
  },
  created() {
    this.newParams = "" 
    // 默认获取全部数据
    this.getListByPage()
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