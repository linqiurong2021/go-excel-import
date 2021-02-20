<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 11:34:12
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:57:38
 * @ Description: 面标绘工具组件
 -->

<template>
  <my-dialog title="面标绘" :visible="dialogVisible" :modal="false" :close-on-click-modal="false"
    width="25%"
    draggable
    titlePosition="center"
    @close="close">

     <el-form ref="markerPolygon" :rules="rules" :model="temp" label-width="80px">
      <el-form-item class="item" label="名称" prop="NAME">
        <el-input  v-model="temp.NAME" placeholder="请输入名称" size="small"  clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="备注" prop="NOTE">
        <el-input type="note" v-model="temp.NOTE" placeholder="请输入备注" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="边框宽度" prop="BORDER_WIDTH">
        <el-input v-model="temp.BORDER_WIDTH" placeholder="请输入线条宽度" size="small"></el-input>
      </el-form-item>
      <el-form-item class="item" label="边框样式" prop="BORDER_STYLE">
         <el-select v-model="temp.BORDER_STYLE" placeholder="请选择">
          <el-option
            v-for="item in lineOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
       <el-form-item class="item" label="边框颜色" prop="BORDER_COLOR">
         <el-color-picker v-model="temp.BORDER_COLOR" :predefine="predefineColors"/>
      </el-form-item>
       <el-form-item class="item" label="填充颜色" prop="FILL_COLOR">
         <el-color-picker v-model="temp.FILL_COLOR" :predefine="predefineColors"/>
      </el-form-item>
      <el-form-item class="item" label="透明度" prop="FILL_OPACITY">
         <el-slider v-model="temp.FILL_OPACITY" :step="0.01" :min="0" :max="1"  style="width:220px"></el-slider>
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

import {  Button, Input, Radio, Select, Option, ColorPicker,Slider,Form,FormItem } from 'element-ui'
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
    // ElDialog: Dialog,
    ElInput: Input,
    ElButton: Button,
    ElRadio: Radio,
    ElSelect: Select,
    ElColorPicker: ColorPicker,
    ElOption: Option,
    ElSlider: Slider,
    ElForm: Form,
    ElFormItem: FormItem
  },
  watch: {
    data: {
      handler: function(newVal){
        this.temp = newVal
      },
      deep: true,
      immediate: true
    },
    temp: {
      handler: function(){
        this.setStyle()
      },
      deep: true,
      immediate: true
    }
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
        BORDER_WIDTH: '3',
        BORDER_STYLE: 'solid',
        BORDER_COLOR: '#2E7DDD',
        FILL_COLOR: '#2E7DDD',
        FILL_OPACITY: 0.8
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
        BORDER_WIDTH: '3',
        BORDER_STYLE: 'solid',
        BORDER_COLOR: '#2E7DDD',
        FILL_COLOR: '#2E7DDD',
        FILL_OPACITY: 0.8
      }
    },
    setStyle() {
      console.log(this.layer, 'layer')
      if(!this.layer) return
      this.layer.setStyle({
        color: this.temp.BORDER_COLOR,
        weight: this.temp.BORDER_WIDTH,
        fillColor: this.temp.FILL_COLOR,
        fillOpacity: this.temp.FILL_OPACITY,
        dashArray: this.temp.line == 'dash' ? 10 : ''
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
      this.$refs['markerPolygon'].validate(valid => {
        if (valid) {
          this.closeDialog()
          this.temp.MARKER_NAME = 'PolygonMarker'
          console.log(this.temp,'this.temp')
          this.$emit('confirm', this.layer, this.temp, this.selectedIcon)
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