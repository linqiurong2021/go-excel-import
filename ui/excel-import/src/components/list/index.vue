<template>
 <div>
  <!--v-loading 需要引用并使用 位置main.js-->
   <el-table
    v-loading="loading"
    element-loading-text="拼命加载中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
    :data="tableData"
    style="width: 100%"
    @selection-change="rowSelectionChange"
    :row-class-name="tableRowClassName">
    <el-table-column
      type="selection"
      width="55">
    </el-table-column>
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
        <el-button @click="handleDelete(scope.row)" type="text" size="small" >删除</el-button>
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
    <!--详情弹窗-->
    <DetailDialog :title="dialogTitle" :types="fieldsType" :detail="detailData" :fields="fields" :names="fieldsName" ref="dialog" @refresh="refresh"/>
 </div>
</template>

<script>
import {Table, TableColumn, Pagination, Button, Popconfirm,MessageBox, Message} from "element-ui"
import {ListField} from '../../utils/const.js'
import {getListByPage, getDataByID, getFieldsType,getFieldsName, deleteBySysIDs, getFields} from '../../api/excel-import.js'
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
    ElPopconfirm: Popconfirm
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
    },
    // 是否新增
    isAdd: {
      type: Boolean,
      default: false
    },
  },
  watch: {
     // 新增操作
    isAdd(newVal, oldVal) {
      this.addOp = newVal
    },
    addOp(newVal, oldVal){
      if(newVal) {
        let tableParams= {table: this.tableName}
        // console.log("newData")
        this.dialogTitle = '新增'
        this.getFieldsType(tableParams)
        this.getFieldsName(tableParams)
        this.getFields(tableParams)
        this.$emit('addChange',false)
      }
    },
    config(newConf, oldConf){
      console.log(newConf,'newConf')
    },
    // 模板切换
    tableName() {
      this.newParams = {}
      
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
    },

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
    
    // 更新成功后需要刷新
    refresh() {
      this.getListByPage()
    },
    getFieldsType(tableParams) {
     getFieldsType(tableParams).then((res)=>{
        let {data} = res
        this.fieldsType = data
        this.$refs.dialog.dialogVisible = true
        console.log(data,"fieldsType")
      })
    },
    getFieldsName(tableParams) {
      getFieldsName(tableParams).then((res)=>{
        let {data} = res
        this.fieldsName = data
        this.$refs.dialog.dialogVisible = true
      })
    },
    getFields(tableParams) {
      getFields(tableParams).then((res)=>{
        let {data} = res
        this.fields = data
        this.$refs.dialog.dialogVisible = true
      })
    },
    //
    getDataByID(sysID) {
       let params = {
        table: this.tableName,
        sysID: sysID
      }
      getDataByID(params).then((res)=>{
        let {data} = res
        this.detailData = data
        this.$refs.dialog.dialogVisible = true
          console.log(data,'res')
      })
    },
    handleClick(row) {
     
      let tableParams= {table: this.tableName}
      //
      this.getFieldsType(tableParams)
      //
      this.getDataByID(row.SYS_ID)
      // 
      this.getFieldsName(tableParams)
      //
      this.getFields(tableParams)
    },
    handleDelete(row) {
      // 提示是否需要删除
      MessageBox.confirm('删除后不可恢复,是否删除？', '确认信息', {
        distinguishCancelAndClose: true,
        confirmButtonText: '删除',
        cancelButtonText: '取消'
      }).then(() => {
        //
        let params = {
          table: this.tableName,
          sys_ids: row.SYS_ID
        }
        deleteBySysIDs(params).then((res)=> {
          let {data} = res
          if (data > 0) {
            // 删除成功
            Message({
              type: 'success',
              message: "删除成功"
            })
            // 刷新页面
            this.refresh()
            console.log("删除成功")
          }else{

            console.error(res)
          }
        })
        console.log("删除了")
      })
      .catch(action => {
        console.log("取消了",action)
      });

      
    },
    tableRowClassName({row, rowIndex}) {
      if (rowIndex === 1) {
        return 'warning-row';
      } else if (rowIndex === 3) {
        return 'success-row';
      }
      return '';
    },
    // 行选中状态
    rowSelectionChange(val) {
      this.$emit('selectedChange', val)
    },
    // 获取分页数据内容
    getListByPage() {
      this.loading = true
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
        this.loading = false
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
      fields: {},
      //
      detailData: {},
      // 显示加载中
      loading: false,
      //
      addOp: false,
      // 选中的项
      selectedRows: [],
      // 弹窗标题
      dialogTitle: '详情'
    }
  },
  created() {
    this.newParams = {}
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