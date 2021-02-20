<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 11:34:04
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 11:20:24
 * @ Description: 线标绘工具组件
 -->

<template>
  <my-dialog title="线标绘" :visible="dialogVisible" :modal="false" :close-on-click-modal="false"
    width="25%"
    draggable
    @close="close">


    <el-form ref="markerPolyline" :rules="rules" :model="temp" label-width="80px">
      <el-form-item class="item" label="名称" prop="NAME">
        <el-input  v-model="temp.NAME" placeholder="请输入名称" size="small"  clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="备注" prop="NOTE">
        <el-input type="note" v-model="temp.NOTE" placeholder="请输入备注" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="线条宽度" prop="LINE_WIDTH">
        <el-input v-model="temp.LINE_WIDTH" placeholder="请输入线条宽度" size="small"></el-input>
      </el-form-item>
      <el-form-item class="item" label="线条样式" prop="LINE_STYLE">
         <el-select v-model="temp.LINE_STYLE" placeholder="请选择">
          <el-option
            v-for="item in lineOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
       <el-form-item class="item" label="线条颜色" prop="LINE_COLOR">
         <el-color-picker v-model="temp.LINE_COLOR" :predefine="predefineColors"/>
      </el-form-item>
      <el-form-item class="item" label="是否公开" prop="IS_PUBLIC">
        <el-radio  v-model="temp.IS_PUBLIC" label="1">是</el-radio>
        <el-radio  v-model="temp.IS_PUBLIC" label="0">否</el-radio>
      </el-form-item>
    </el-form>
  
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="confirm">确 定</el-button>
      <el-button @click="remove" type="danger">删除</el-button>
     
    </span>
  </my-dialog>
</template>

<script>

import { Button, Input, Radio, Select, Option, ColorPicker,Form,FormItem } from 'element-ui'
import MyDialog from '../../dialog/index'

export default {
  name: 'PointMarker',
  props: {
    eginMap: {
      type: Object,
      require: true
    },
    // icon
    markerOptions: {
      type: Array,
      default() {
        return []
      }
    },
    // 显示详情数据
    data: {
      type: Object,
      default() {
        return {}
      }
    },
    // 当前图层
    layer: {
      type: Object,
      require: true
    }
  },
  components: {
    MyDialog: MyDialog,
    ElInput: Input,
    ElButton: Button,
    ElRadio: Radio,
    ElSelect: Select,
    ElColorPicker: ColorPicker,
    ElOption: Option,
    ElForm: Form,
    ElFormItem: FormItem
  },
  watch: {
    temp: {
      handler: function(newVal){
        this.setStyle(newVal)
      },
      deep: true,
      immediate: true
    },
    data: {
      handler: function(newVal) {
        this.temp = newVal
      },
      deep: true,
      immediate: true
    },
  },
  data() {
    return {
      // 校验规则
      rules: {
        NAME: [
          {
            required: true,
            message: '名称不能为空',
            trigger: 'blur'
          }
        ],
      },
      lineOptions: [
        {
          value: 'solid',
          label: '—————————'
        }, {
          value: 'dash',
          label: '---------------'
        }
      ],
      // 弹窗
      dialogVisible: false,
      // 默认颜色
      predefineColors: [
        '#ff4500',
        '#ff8c00',
        '#ffd700',
        '#90ee90',
        '#00ced1',
        '#1e90ff',
        '#c71585',
      ],
      map: null,
      temp: {
        NAME: '',
        NOTE: '',
        LINE_WIDTH: '3',
        LINE_STYLE: 'solid',
        LINE_COLOR: '#2E7DDD'
      }
    }
  },
  mounted() {
    this.map = this.eginMap.map
  },
  methods: {
    //
    resetTemp() {
      this.temp = {
        NAME: '',
        NOTE: '',
        LINE_WIDTH: '3',
        LINE_STYLE: 'solid',
        LINE_COLOR: '#2E7DDD'
      }
    },
    setStyle() {
      console.log(this.layer, 'layerId')
      if(!this.layer) return
      this.layer.setStyle({
        color: this.temp.LINE_COLOR,
        weight: this.temp.LINE_WIDTH,
        dashArray: this.temp.LINE_STYLE == 'dash' ? 10 : ''
      })
    },
    closeDialog() {
      this.dialogVisible = false
    },
    remove() {
      this.closeDialog()
      this.$emit('remove', this.layer)
    },
    confirm() {
      this.$refs['markerPolyline'].validate(valid => {
        if (valid) {
          this.closeDialog()
          this.temp.MARKER_NAME = 'PolylineMarker'
          this.$emit('confirm', this.layer, this.temp, null)
        }
      })
    },
    //
    close() {
      this.closeDialog()
      // this.$emit('remove', this.layer)
    }
  },
}
</script>

<style scoped>
.item{
  font-size: 14px;
  border-bottom: 1px solid #cccccc;
}
.item >>> .el-input  input {
  border: none
}
</style>