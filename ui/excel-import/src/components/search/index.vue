<template>
  <div class="search-wrap">
    <div v-for="(item,key) in searchKeys" :key="key">
      <el-input v-model="searchParams[item]" :placeholder="labelName(key,1)" v-if="searchTypeValues[key] == FieldType.TEXT" clearable></el-input>
      <el-select v-model="searchParams[item]" :placeholder="labelName(key,2)" v-else-if="searchTypeValues[key] == FieldType.SELECT " clearable>
        <el-option
          v-for="subItem in selectOptions[item]"
          :key="subItem"
          :label="subItem"
          :value="subItem">
        </el-option>
      </el-select>
    </div>
    <div><el-button icon="el-icon-search" circle @click="search"></el-button> </div>
  </div>
</template>

<script>
import { Input, Select, Button, Option} from "element-ui"
import {IgnoreFields, FieldType,SearchField} from "../../utils/const.js"
import {getSelectOptions} from '../../api/excel-import.js'
import { mapGetters } from "vuex"
export default {
  name: "Search",
  components: {
    ElInput: Input,
    ElSelect: Select,
    ElButton: Button,
    ElOption: Option
  },
  props: {
    config: {
      type: Object,
      default() {
        return {}
      }
    },
    type: {
      type: Object,
      default() {
        return {}
      }
    },
    // 名称
    label: {
      type: Object,
      default() {
        return {}
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
      console.log("configValues")
      return Object.values(this.config)
    },
    typeKeys() {
      return Object.keys(this.type)
    },
    typeValues() {
      return Object.values(this.type)
    },
    // 需要搜索的字段
    searchKeys() {
      // 搜索key
      let searchKeys = []
      // 搜索字段类型
      let searchTypeValues = []
      // 搜索字段
      let searchFields = {}
      // 搜索名称
      let searchLabels = []
      //
      this.configValues.forEach((item,index)=>{
       
        // 通过判断值不为空则说明需要有搜索
        if(item.trim() == this.SearchField.VALUE) {
          let key = this.configKeys[index]
          if (this.IgnoreFields.indexOf(key) == -1) {
            searchKeys.push(key)
            searchTypeValues.push(this.typeValues[index])
            searchFields[key]= ""
            searchLabels.push(this.label[key])
          }
        }
      })
      // 搜索字段类型
      this.searchTypeValues = searchTypeValues
      this.searchParams = searchFields
      this.searchLabels = searchLabels
      // console.log(this.searchTypeValues,this.searchParams,this.searchLabels,'result')
      return searchKeys
    },
    // 下拉字段
    selectKeys() {
      //
      let values = []
      this.typeValues.forEach((item,index)=>{
        if(item == this.FieldType.SELECT) {
          values.push(this.typeKeys[index])
        }
      })
      console.log(values,'selectKeys')
      return values
    },
    ...mapGetters({
      getTableName: "form/getTableName"
    })
  },
  watch: {
    type(newType, oldType){
      // 调用接口获取相应key的数组
      if(this.selectKeys.length > 0) {
        //
        let data = {
          table: this.getTableName,
          keys: this.selectKeys.join(',')
        }
        let tmpSearchOptions = {}
        // data
        getSelectOptions(data).then((res)=>{
          let {data}= res
          this.selectKeys.forEach((item)=>{
            tmpSearchOptions[item] =data[item]
          })
          this.selectOptions = tmpSearchOptions
        })
      }
    }
  },
  data() {
    return {
      //
      FieldType: FieldType,
      //
      IgnoreFields: IgnoreFields,
      //
      SearchField: SearchField,
      // 搜索名称
      searchTypeValues: [],
      // 
      searchParams: {},
      // key
      selectOptions: {},
      // 搜索名称
      searchLabels: {}
    }
  },
  methods: {
    search() {
       // 过滤空的数据
      let newParams = {}
      for (let item in this.searchParams ){
        if(this.searchParams[item]!=""){
           newParams[item]= this.searchParams[item]
        }
      }
      this.$store.dispatch('form/setSearchParams', newParams)
     
    },
    labelName(key ,type ) {
      let label = type == 1 ? "请输入" : "请选择"
      return label + this.searchLabels[key]
    }
  }

}
</script>

<style scoped>
.search-wrap {
  display: flex
}
</style>