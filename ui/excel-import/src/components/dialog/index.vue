<template>
 <div class="dialog">
  <el-dialog title="详情" :visible.sync="dialogVisible">
    <el-form :model="formData">
      <!--获取名称-->
      <div v-for="(item, key) in formData" :key="key">
        <template v-if="key != 'SYS_ID'">
          <el-form-item :label="names[key]" >
            <el-input v-model="formData[key]" :name="key" autocomplete="off" :placeholder="getPlaceholder(key,1)" v-if="types[key] == FieldType.TEXT" ></el-input>
            <el-select v-model="formData[key]" :placeholder="getPlaceholder(key,2)" v-else-if="types[key] == FieldType.SELECT">
              <el-option
                v-for="subItem in selectOptions[key]"
                :key="subItem"
                :label="subItem"
                :value="subItem">
              </el-option>
            </el-select>
            
          </el-form-item>
        </template>
      </div>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
    </div>
  </el-dialog>
 </div>
</template>

<script>
import {Dialog, Button, Form, FormItem, Input, Select, Option} from "element-ui"
import {FieldType,SearchField} from "../../utils/const.js"
import {getSelectOptions} from "../../api/excel-import.js"
import { mapGetters } from "vuex"
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
  },
  data() {
    return {
      dialogVisible: false,
      formData: {},
      FieldType:FieldType,
      SearchField:SearchField,
      selectOptions: []
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
          console.log(data,'data#res')
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
    }
  }
}
</script>

<style scoped>

</style>