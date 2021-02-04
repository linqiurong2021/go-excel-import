<template>
 <div class="dialog">
  <el-dialog :title="title" :visible.sync="dialogVisible">
    <el-form :model="formData">
      <!--获取名称-->
      <div v-for="(item, key) in fields" :key="key">
        <template v-if="key != 'SYS_ID'">
          <el-form-item :label="names[key]" >
            <el-input v-model="formData[key]" :name="key" autocomplete="off" :placeholder="getPlaceholder(key,1)" v-if="types[key] == FieldType.TEXT" ></el-input>
            
            <div  v-else-if="types[key] == FieldType.SELECT">
              <el-select v-model="formData[key]" :placeholder="getPlaceholder(key,2)">
                <el-option
                  v-for="subItem in selectOptions[key]"
                  :key="subItem"
                  :label="subItem"
                  :value="subItem">
                </el-option>
              </el-select>
              <el-button type="primary" size="small" icon="el-icon-plus" circle  @click="showItem(key)"></el-button>
            </div>
          </el-form-item>
        </template>
      </div>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="submit">确 定</el-button>
    </div>
  </el-dialog>
  <newItem :name="newItemName" :item-key="newItemKey" @getItem="getItem" ref="newItem"/>
 </div>
</template>

<script>
import {Dialog, Button, Form, FormItem, Input, Select, Option, Message} from "element-ui"
import {FieldType,SearchField} from "../../utils/const.js"
import {getSelectOptions, updateBySysID, createData} from "../../api/excel-import.js"
import { mapGetters } from "vuex"
import newItem from "./newItem"
export default {
  name: "Dialog",
  props: {
    types: {
      type: Object,
      default() {
        return {}
      }
    },
    names: {
      type: Object,
      default() {
        return {}
      }
    },
    fields: {
      type: Object,
      default() {
        return {}
      }
    },
    title: {
      type: String,
      default: ''
    },
    detail: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  components: {
    ElForm: Form,
    ElDialog: Dialog,
    ElButton: Button,
    ElFormItem: FormItem,
    ElInput:Input,
    ElSelect: Select,
    ElOption: Option,
    newItem
  },
  data() {
    return {
      dialogVisible: false,
      formData: {},
      FieldType:FieldType,
      SearchField:SearchField,
      selectOptions: [],
      newItemName: "",
      newItemKey: ""
    }
  },
  computed: {
    ...mapGetters({
      getTableName: "form/getTableName"
    }),
       // 下拉字段
    selectKeys() {
      //
      let values = []
      for(let key in this.types) {
         if(this.types[key] == this.FieldType.SELECT) {
          values.push(key)
        }
      }
     
      return values
    },
    
  },
  watch: {
    types(newVal, oldVal){
      // 调用接口获取相应key的数组
      if(this.selectKeys.length > 0) {
        //
        let data = {
          table: this.getTableName,
          keys: this.selectKeys.join(',')
        }
        let tmpSearchOptions = []
        // data
        getSelectOptions(data).then((res)=>{
          let {data}= res
          // console.log(data,'data#res')
          this.selectKeys.forEach((item)=>{
            tmpSearchOptions[item] =data[item]
          })
          this.selectOptions = tmpSearchOptions
          // this.$set(this,"selectOptions",tmpSearchOptions)
        })
      }
    },
   
    names(newVal, oldVal){
      console.log(newVal, 'names')
    },
    detail(newData, oldData){
      //
      let tmpFormData = {}
      for(let item in newData) {
        tmpFormData[item] =newData[item]
      }
      this.formData = tmpFormData
      console.log(newData, 'newData', tmpFormData)
    },

  },
  methods: {
    getPlaceholder(key,type) {
      let label = type == 1 ? "请输入": "请选择"
      return label+this.names[key]
    },
    submit() {
      // 更新数据
      let data = {
        table: this.getTableName,
        params: JSON.stringify(this.formData)
      }
      if(this.title == "新增"){
        createData(data).then((res)=>{
          let {data} = res
          if (data >=0) {
            console.log("新增成功")
            Message({
              type: 'success',
              message: "新增成功"
            })
            this.dialogVisible = false
            this.formData = {} // 清除数据
            this.$emit('refresh')
          }else{
            console.error(data)
          }
          console.log(data.data)
        })
      }else {
        updateBySysID(data).then((res)=>{
          let {data} = res
          if (data >=0) {
            console.log("修改成功")
            Message({
              type: 'success',
              message: "修改成功"
            })
            this.formData = {} // 清除数据
            this.dialogVisible = false
            this.$emit('refresh')
          }else{
            console.error(data)
          }
          console.log(data.data)
        })
      }
     
      console.log(this.formData,'formData')
    },
    showItem(key) {
      this.newItemName = this.getItem(key,1)
      this.newItemKey = key
      this.$refs.newItem.show = true
      this.$refs.newItem.inputData = ""
    },
    getItem(data) {
      // 如果不存在则直接返回
      if (!this.selectOptions[data.key]) {
        return
      }
      let index = this.selectOptions[data.key].length
      this.$set(this.selectOptions[data.key], index , data.value)
      // 默认为新增的值
      this.formData[data.key] = data.value
    },
  }
}
</script>

<style scoped>

</style>