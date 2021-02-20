<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 11:33:53
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 17:14:16
 * @ Description: 点标绘工具组件
 -->

<template>
  <my-dialog title="点标绘" :visible="dialogVisible" :modal="false" :close-on-click-modal="false"
    width="25%"
    draggable
    titlePosition="center"
    @close="close">

    <el-form ref="markerPoint" :rules="rules" :model="temp" label-width="80px">
      <el-form-item class="item" label="名称" prop="NAME">
        <el-input  v-model="temp.NAME" placeholder="请输入名称" size="small"  clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="备注" prop="NOTE">
        <el-input type="note" v-model="temp.NOTE" placeholder="请输入备注" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="是否公开" prop="IS_PUBLIC">
        <el-radio  v-model="temp.IS_PUBLIC" label="1">是</el-radio>
        <el-radio  v-model="temp.IS_PUBLIC" label="0">否</el-radio>
      </el-form-item>
      <el-form-item class="item" label="图标" prop="ICON_NAME">
        <div class="icon iconfont" :class="temp.ICON_NAME"></div>
      </el-form-item>
      <div class="item icon-list">
          <i class="icon iconfont" :class="item.iconName" v-for="(item, key) in iconOptions" :key="key" @click="changeIcon(item,key)" :style="getStyle(item.iconSize)"></i>
      </div>
    </el-form>
   
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="confirm">确 定</el-button>
      <el-button @click="remove" type="danger">删除</el-button>
    </span>
  </my-dialog>
</template>

<script>

import {  Button, Input, Radio ,Form,FormItem } from 'element-ui'
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
    ElForm: Form,
    ElFormItem: FormItem
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
        
      // iconList 为修改地图上的icon做准备
      iconList: [],
      // 已选中的iocn
      selectedIcon: null,
      // 弹窗
      dialogVisible: false,
      // 默认定位器
      iconOptions: [{
          iconName: 'iconsoph_location',
          iconSize: [20,20]
        },
        {
          iconName: 'iconic_location_on_px',
          iconSize: [20,20]
        }
      ],
      map: null,
      temp: {
        NAME: '',
        NOTE: '',
        IS_PUBLIC: 0,
        ICON_NAME: 'iconic_location_on_px'
      }
    }
  },
  watch: {
    dialogVisible: function(newVal){
      if(!newVal) {
        this.temp = this.data
      }
    },
    data: {
      handler:  function(newVal) {
        this.temp = newVal
      },
      deep: true,
      immediate: true
    }
  },
  //
  mounted() {
    this.dialogVisible = true
    if(this.markerOptions && this.markerOptions.length > 0) {
      this.iconOptions = this.markerOptions
    }
    // 初始化icon
    this.iconOptions.map((item)=>{
      let marker = this.eginMap.MapTools.newDivIcon(item.iconName, item.iconSize)
      this.iconList.push(marker)
    })
    this.map = this.eginMap.map
    
    // this.temp = this.data
  },
  methods: {
    resetTemp() {
      this.temp = {
        longitude: '',
        latitude: '',
        IS_PUBLIC: 0,
        NOTE: '',
        // MARKER_NAME: 'PointMarker',
        ICON_NAME: 'iconic_location_on_px'
      }
    },
    getStyle(iconSize) {
      return `"width:${iconSize[0]}";height:${iconSize[1]}`
    },
    // 修改icon 并设置当前选中的icon
    changeIcon(item,key){
      // console.log(item,'item')
      this.selectedIcon = this.iconList[key] ? this.iconList[key] : this.iconList[0]
      // 切换时有问题
      // this.temp.ICON_NAME = item.iconName
      // this.temp.tmp_time = new Date().getMilliseconds()
      this.$set(this.temp,'ICON_NAME',item.iconName)
      // 图层设置
      console.log(this.temp,'temp')
      this.layer.setIcon(this.selectedIcon)
    },
    closeDialog() {
      this.dialogVisible = false
    },
    remove() {
      // console.log(this.layer, 'remove###layer')
      this.closeDialog()
      this.$emit('remove', this.layer)
    },
    confirm() {
      this.$refs['markerPoint'].validate(valid => {
        if (valid) {
          this.closeDialog()
          this.temp.MARKER_NAME = 'PointMarker'
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

div.icon-list {
  display: flex;
  flex-wrap: wrap;
  border-bottom: none;
}
.icon-list .icon{
  font-size: 30px;
  padding: 5px;
}
i.iconfont{
  cursor: pointer;
}
.item >>> .el-input  input {
  border: none
}
</style>