<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-30 11:17:06
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 17:04:38
 * @ Description: 自定义弹窗 合并图形列表组件
 -->

<template>
 <div class="layer-dialog">
   <my-dialog title="合并列表" :visible="visible" draggable @close="close">
      <div>请选择一个图斑继承未切割前的图斑业务属性：</div>
      <el-table :data="tableData" highlight-current-row @current-change="handleCurrentChange">
        <el-table-column :property="field" label="名称"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancleUnion">
          取消
        </el-button>
        <el-button
          type="primary"
          @click="confirmUnion"
        >
          确定
        </el-button>
      </div>
   </my-dialog>
 </div>
</template>

<script>

import MyDialog from '../../dialog/index'
import { Table, Button, TableColumn} from 'element-ui'
export default {
  name: 'UnionList',
  props: {
    data: {
      type: Array,
      required: true
    },
    field: {
      type: String,
      require: true
    },
    colorList: {
      type: Array,
      default() {
        return []
      }
    }
  },
  components: {
    ElTable: Table,
    ElButton: Button, 
    ElTableColumn: TableColumn,
    MyDialog 
  },
  data() {
    return {
      visible: true,
      tableData: [],
      confirmData: {}
    }
  },
  watch:{
    data: {
      handler: function(newVal){
     
        let tmpDataList = []
        newVal.forEach((item)=>{
          tmpDataList.push(item.feature.properties)
        })
        this.tableData = tmpDataList
        this.visible = true
        // console.log(this.tableData,'data###watch')
      },
      immediate: true
    }
  },
  methods: {
    close() {
      this.$store.dispatch('map/setCancleUnion', true)
      // 重置颜色 取消后则删除
      this.visible = false
    },
    // 修改颜色
    handleCurrentChange(currentRow, oldRow) {
      console.log(currentRow, oldRow)
      if(!currentRow) return
      // console.log(this.tableData,'handleCurrentChange###tableData',currentRow)
      this.tableData.map((item, key)=>{
        // console.log(item,'item')
        if(item[this.field] == currentRow[this.field]) {
          const selected = {
            selected: this.data[key],
            key: key
          }
          this.confirmData = selected
          this.$store.dispatch('map/setSelected',this.confirmData)
        }
      })

    },
    cancleUnion() {
      this.close()
    },
    confirmUnion() {
       //
      const confirmData = {
        selected: this.confirmData.selected,
        selectList: this.data,
        key: this.confirmData.key
      }
      this.$store.dispatch('map/setConfirmUnion', confirmData)
      // console.log('confirmCut',)
      this.close()
    }
  }
}
</script>

<style scoped>
.my-dialog{
  opacity: 0.9;
}
</style>