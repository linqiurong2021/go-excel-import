<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-30 11:13:12
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-15 08:27:11
 * @ Description: 自定义弹窗 切割图形列表组件
 -->

<template>
 <div class="layer-dialog">
   <my-dialog title="切割列表" :visible="visible" draggable @close="close">
      <div>请选择一个图斑继承未切割前的图斑业务属性：</div>
      <el-table :data="tableData" highlight-current-row @current-change="handleCurrentChange">
        <!--根据具体业务修改ID-->
        <el-table-column :property="field" label="图斑编号"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancleCut">
          取消
        </el-button>
        <el-button
          type="primary"
          @click="confirmCut"
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
  name: 'CutList',
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
        console.log(newVal,'data###watch')
        let tmpDataList = []
        newVal.forEach((item)=>{
          tmpDataList.push(item.feature.properties)
        })
        
        // this.$set(this,this.tableData, tmpDataList)
        this.tableData = tmpDataList
        this.visible = true
        // console.log(this.tableData,'tableData',tmpDataList)
      },
      immediate: true
    }
  },
  methods: {
    handleCurrentChange(currentRow) {
      if(!currentRow) return
      // console.log(this.tableData,'handleCurrentChange###tableData',currentRow)
      this.tableData.map((item, key)=>{
        // console.log(item,'item')
        if(item[this.field] == currentRow[this.field]) {
          //
          const selected = {
            selected: this.data[key],
            key: key
          }
          
          this.confirmData = selected
          this.$store.dispatch('map/setSelected',this.confirmData)
        }
        console.log('item@handleCurrentChange',item)
      })
      // console.log('dispatch')
    },
    close() {
      this.$store.dispatch('map/setCancleCut', true)
      // 重置颜色 取消后则删除
      this.visible = false
    },
    cancleCut() {
      
      this.close()
      this.$store.dispatch('map/setCancleCut', true)
    },
    confirmCut() {
      //
      const confirmData = {
        selected: this.confirmData.selected,
        selectList: this.data,
        key: this.confirmData.key
      }
      this.$store.dispatch('map/setConfirmCut', confirmData)
      // console.log('confirmCut',)
      this.close()
    }
  }
}
</script>

<style scoped>

</style>