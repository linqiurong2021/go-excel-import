# PointMarker 点标绘组件
> 参数说明
  >> props 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | markerOptions  | 图例样式 |
      | data  | 表单详情数据 |
      | layer  | 当前图层 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElForm  | ElementUI 表单组件 |
      | ElFormItem  | ElementUI 表单组件 |
      | Button  | ElementUI 按钮组件 |
      | ElRadio  | ElementUI Radio组件 |
      | ElInput  | ElementUI Input组件 |
      | MyDialog  | 自定义弹窗组件 |
  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  表单数据 |
      | dialogVisible  |  是否显示弹窗 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  |  校验规则 |
      | iconList  |  可选icon列表 |
      | selectedIcon  | 当前选中的icon |
      | dialogVisible  | 是否显示弹窗 |
      | iconOptions  | 默认定位器 |
      | map  | 地图实例 |
      | temp  | 临时对象 |
      | temp.NAME  | 地图实例 |
      | temp.NOTE  | 默认定位器 |
      | temp.IS_PUBLIC  | 地图实例 |
      | temp.ICON_NAME  | 地图实例 |


  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | resetTemp  | 重置表单数据 |
      | getStyle  | 获取样式 |
      | changeIcon  | 修改ICON |
      | remove  | 删除 |
      | confirm  | 确认 |
      | close  | 关闭弹窗 |

> 组件代码
```html
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
```
# PolygonMarker 面标绘组件
> 参数说明
  >> props 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | markerOptions  | 图例样式 |
      | data  | 表单详情数据 |
      | layer  | 当前图层 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElForm  | ElementUI 表单组件 |
      | ElFormItem  | ElementUI 表单组件 |
      | Button  | ElementUI 按钮组件 |
      | ElSelect | ElementUI Select组件|
      | ElOption | ElementUI Option组件|
      | ElSlider|  ElementUI Slider组件|
      | ElRadio  | ElementUI Radio组件 |
      | ElInput  | ElementUI Input组件 |
      | MyDialog  | 自定义弹窗组件 |
  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  表单数据 |
      | temp  |  表单数据 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  |  校验规则 |
      | lineOptions  |  可选线类型 |
      | dialogVisible  | 是否显示弹窗 |
      | predefineColors  | 默认预定义颜色选项 |
      | map  | 地图实例 |
      | temp  | 临时对象 |
      | temp.NAME  | 地图实例 |
      | temp.NOTE  | 默认定位器 |
      | temp.IS_PUBLIC  | 地图实例 |
      | temp.BORDER_WIDTH  | 边框宽度 |
      | temp.BORDER_STYLE  | 边框类型 |
      | temp.FILL_COLOR  | 填充色 |
      | temp.FILL_OPACITY  | 填充透明度 |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | resetTemp  | 重置表单数据 |
      | setStyle  | 设置样式 |
      | closeDialog  | 关闭弹窗 |
      | remove  | 删除 |
      | confirm  | 确认 |
      | close  | 关闭弹窗 |

> 组件代码
```html
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
```


# PolylineMarker 线标绘组件

> 参数说明
  >> props 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | markerOptions  | 图例样式 |
      | data  | 表单详情数据 |
      | layer  | 当前图层 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElForm  | ElementUI 表单组件 |
      | ElFormItem  | ElementUI 表单组件 |
      | Button  | ElementUI 按钮组件 |
      | ElSelect | ElementUI Select组件|
      | ElOption | ElementUI Option组件|
      | ElSlider|  ElementUI Slider组件|
      | ElRadio  | ElementUI Radio组件 |
      | ElInput  | ElementUI Input组件 |
      | ElColorPicker  | ElementUI ColorPicker组件 |
      | MyDialog  | 自定义弹窗组件 |
  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  表单数据 |
      | temp  |  表单数据 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  |  校验规则 |
      | lineOptions  |  可选线类型 |
      | dialogVisible  | 是否显示弹窗 |
      | predefineColors  | 默认预定义颜色选项 |
      | map  | 地图实例 |
      | temp  | 临时对象 |
      | temp.NAME  | 地图实例 |
      | temp.NOTE  | 默认定位器 |
      | temp.IS_PUBLIC  | 地图实例 |
      | temp.LINE_WIDTH  | 线宽度 |
      | temp.LINE_STYLE  | 线类型 |
      | temp.LINE_COLOR  | 线颜色 |
     

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | resetTemp  | 重置表单数据 |
      | setStyle  | 设置样式 |
      | closeDialog  | 关闭弹窗 |
      | remove  | 删除 |
      | confirm  | 确认 |
      | close  | 关闭弹窗 |

> 组件代码
```html
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

```


# TextMarker 文字标绘组件
> 参数说明
  >> props 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | data  | 表单详情数据 |
      | layer  | 当前图层 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElForm  | ElementUI 表单组件 |
      | ElFormItem  | ElementUI 表单组件 |
      | Button  | ElementUI 按钮组件 |
      | ElSelect | ElementUI Select组件|
      | ElOption | ElementUI Option组件|
      | ElSlider|  ElementUI Slider组件|
      | ElRadio  | ElementUI Radio组件 |
      | ElInput  | ElementUI Input组件 |
      | ElColorPicker  | ElementUI ColorPicker组件 |
      | MyDialog  | 自定义弹窗组件 |
  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  表单数据 |



  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  |  校验规则 |
      | fontWeightOptions  |  字体加粗 |
      | dialogVisible  | 是否显示弹窗 |
      | predefineColors  | 默认预定义颜色选项 |
      | map  | 地图实例 |
      | temp  | 临时对象 |
      | temp.NAME  | 地图实例 |
      | temp.NOTE  | 默认定位器 |
      | temp.IS_PUBLIC  | 地图实例 |
      | temp.FONT_SIZE  | 字体大小 |
      | temp.FONT_WEIGHT  | 字体加粗 |
      | temp.FONT_COLOR  | 字体颜色 |
     

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | resetTemp  | 重置表单数据 |
      | getStyle  | 获取样式 |
      | setIcon  | 设置样式 |
      | closeDialog  | 关闭弹窗 |
      | remove  | 删除 |
      | confirm  | 确认 |
      | close  | 关闭弹窗 |

> 组件代码
```html

<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 11:34:29
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 11:24:07
 * @ Description: 文本标绘工具组件
 -->

<template>
  <my-dialog
    title="文本标绘"
    :visible="dialogVisible"
    :modal="false"
    :close-on-click-modal="false"
    width="25%"
    draggable
    @close="close"
  >

    <el-form ref="markerText" :rules="rules" :model="temp" label-width="80px">
      <el-form-item class="item" label="名称" prop="NAME">
        <el-input  v-model="temp.NAME" placeholder="请输入名称" size="small"  clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="备注" prop="NOTE">
        <el-input type="note" v-model="temp.NOTE" placeholder="请输入备注" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="字体大小" prop="FONT_SIZE">
        <el-input type="note" v-model="temp.FONT_SIZE" placeholder="请输入字体大小" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="字体颜色" prop="FONT_COLOR">
         <el-color-picker v-model="temp.FONT_COLOR" :predefine="predefineColors"/>
      </el-form-item>
       <el-form-item class="item" label="字体粗细" prop="FONT_WEIGHT">
         <el-select v-model="temp.FONT_WEIGHT" placeholder="请选择">
          <el-option
            v-for="item in fontWeightOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
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
import {
  Button,
  Input,
  Radio,
  Select,
  Option,
  ColorPicker,
  Form,
  FormItem
} from "element-ui";
import MyDialog from '../../dialog/index'
export default {
  name: "PointMarker",
  props: {
    eginMap: {
      type: Object,
      require: true,
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
      require: true,
    },
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
    data: {
      handler: function(newVal) {
        this.temp = newVal
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
      // 弹窗
      dialogVisible: false,
      // 默认颜色
      predefineColors: [
        "#ff4500",
        "#ff8c00",
        "#ffd700",
        "#90ee90",
        "#00ced1",
        "#1e90ff",
        "#c71585",
      ],
      // 字体样式
      fontWeightOptions: [
        {
          value: "normal",
          label: "正常",
        },
        {
          value: "bold",
          label: "加粗",
        },
        {
          value: "bolder",
          label: "更粗",
        },
      ],
      map: null,
      temp: {
        NAME: "",
        NOTE: "",
        FONT_SIZE: 12,
        FONT_WEIGHT: "normal",
        FONT_COLOR: "#000000",
        IS_PUBLIC: "0",
        USER_ID: ''
      },
    };
  },
  //
  mounted() {
    this.map = this.eginMap.map;
  },
  methods: {
    resetTemp() {
      this.temp = {
        NAME: "",
        NOTE: "",
        FONT_SIZE: 12,
        FONT_WEIGHT: "normal",
        FONT_COLOR: "#000000",
        IS_PUBLIC: "0",
        USER_ID: ''
      };
    },
    getStyle(iconSize) {
      return `"width:${iconSize[0]}";height:${iconSize[1]}`;
    },
    // DiVICON
    setIcon() {
      let html = `<div style="width:200px;font-size:${this.temp.FONT_SIZE}px;font-weight:${this.temp.FONT_WEIGHT};color:${this.temp.FONT_COLOR}">${this.temp.NAME}</div>`;
      //
      let divIcon = this.eginMap.MapTools.newDivIcon("iconic_location_on_px", [0, 0], html);
      //
      console.log(divIcon)
      const latlng = this.layer.getLatLng();
      let textMarker = this.eginMap.MapTools.newMarker(latlng, {
        icon: divIcon,
      });
      //
      textMarker.addTo(this.map);
      
      // 删除点位数据
      return textMarker;
    },
    closeDialog() {
      this.dialogVisible = false;
    },
    remove() {
      this.closeDialog();
      this.$emit("remove", this.layer);
    },
    confirm() {
      this.$refs['markerText'].validate(valid => {
        if (valid) {
          this.closeDialog();
          this.temp.MARKER_NAME = 'TextMarker'
          this.setIcon()
          this.$emit("confirm", this.layer, this.temp, null);
        }
      })
    },
    //
    close() {
      this.closeDialog();
      // this.$emit("remove", this.layer);
    },
  },
};
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

```

# 调用事例代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:21
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 16:08:51
 * @ Description: 标绘工具栏
 -->

<template>
  <div class="mark-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="mark-tools">
      <div
        class="item mark-item"
        v-for="(item, key) in tools"
        :key="key"
        @click="mark(item.type, key)"
      >
        {{ item.name }}
      </div>
    </div>
    <!-- 20201016 与潘核实取消清除功能--->
    <!-- <div class="item clear" @click="clear">清除</div> -->
    <div class="mark-tools item close" @click="showTools = false">关闭</div>
    <!--弹窗-->
    <!-- <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="confirm"
      @remove="remove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component> -->
    <!-- <mark-list
      :list="dataLists"
      @remove="markRemove"
      @location="markLocation"
    ></mark-list> -->
  </div>
</template>

<script>
import GeometryType from '../../const/GeometryType'

// import { mapApi } from '@api/mapApi/map'

// 工具提示
import './tooltip.js'
import { mapGetters } from 'vuex'
export default {
  name: 'Mark',
  props: {
    eginMap: Object
  },
  computed: {
    ...mapGetters({
      markTool: 'map/markTool'
    })
  },
  data() {
    return {
      title: '标绘',
      // 存放标绘的图层
      featureGroup: null,
      // 图层显示
      pane: null,
      //
      showTools: true,
      // 当前选中是哪个组件
      currentMarker: 'PointMarker',
      // 点工具 线工具 面工具 文本工具
      drawTools: [],
      // 当前工具
      currentTool: [],
      // 当前操作的图层
      currentLayer: null,
      // 临时图层(未确认的)
      tmpLayers: [],
      // 已添加的图层(图形)(已确认的)
      layers: [],
      // 是否显示弹窗
      show: [false, false, false, false],
      // 存储后的数组
      dataLists: [],
      map: null,
      tools: [
        {
          name: '点标绘',
          type: 'PointMarker'
        },
        {
          name: '线标绘',
          type: 'PolylineMarker'
        },
        {
          name: '面标绘',
          type: 'PolygonMarker'
        },
        {
          name: '文字标绘',
          type: 'TextMarker'
        }
      ],
      // 标记的数据
      markerData: {},
      businessLayerId: 'a191c15515da4016b71681394e5a8479',
      //
      pointLayerId: 'e6e25e52f6384d52a6e941cfe7ad551d', // 标绘点
      //
      poylineLayerId: 'b197aba58a934a8ba576632f44938f54', // 标绘线
      //
      polygonLayerId: '7fc9aff62ea640bc82de56ebd4d458b0', // 标绘面
      //
      textLayerId: 'b7c72b2ed74c448ba1664bcd22dc8878 ', // 文字标绘
      // 7fc9aff62ea640bc82de56ebd4d458b0 标绘面
      // b197aba58a934a8ba576632f44938f54 标绘线
      // e6e25e52f6384d52a6e941cfe7ad551d 标绘点
      // b7c72b2ed74c448ba1664bcd22dc8878 文字标绘
    }
  },
  //
  mounted() {
    this.map = this.eginMap.map

    const [featureGroup, pane] = this.eginMap.MapTools.initPane('markPane', 501)
    this.pane = pane
    this.featureGroup = featureGroup
    featureGroup.addTo(this.map)

    this.initTools()
    // 监听新增结果
    this.getResult()
  },
  methods: {
    //
    initTools() {
      //
      const markerOptions = {
        //
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }

      const textMarkerOptions = {
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }
      const drawPoint = this.eginMap.MapTools.initMarker(markerOptions)
      const drawPolyline = this.eginMap.MapTools.initDrawPolyline()
      const drawPolygon = this.eginMap.MapTools.initDrawPolygon()
      const drawText = this.eginMap.MapTools.initTextMarker(textMarkerOptions)
      const drawTools = [drawPoint, drawPolyline, drawPolygon, drawText]

      this.drawTools = drawTools
    },
    mark(type, key) {
      // 当前组件
      this.currentMarker = type
      //
      this.drawTools.map((tool, index) => {
        index == key ? tool.enable() : tool.disable()
        if (index == key) {
          // 开启哪个工具
          this.currentTool = tool
        }
      })
    },
    getResult() {
      this.map.off('draw:created').on('draw:created', e => {
        const layer = e.layer
        // 当前图层
        this.currentLayer = layer
        // 加入到数组中  删除时需要用到
        // this.tmpLayers.push(layer)
        // 先添加后地图上  如果选择删除后 再删除
        // layer.addTo(this.map)

        this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker , currentLayer: layer , markerData: {} })
        console.error('getResult#showDialog')
        layer.on('click', evt => {
          //
          let data = evt.target.options.data
          // 弹窗相应弹窗
          this.currentMarker = data.MARKER_NAME
          this.markerData = data
          // 
          this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker, currentLayer: layer , markerData: data })
          //
          // setTimeout(() => {
          //   this.$refs.markerDialog.dialogVisible = true
          // }, 300)
          console.log(evt, data)
        })
        // 不在这里做数据添加 需要在确认时做数据添加 点击时需要获取数据框中的数据
        // this.featureGroup.addLayer(layer)

        // setTimeout(() => {
        //   this.$refs.markerDialog.dialogVisible = true
        // }, 300)
        // console.log(this.show, 'show')
      })
    },
    // 删除
    remove(layer) {
      // 当前layer在datalist的位置再删除
      this.removeLayer(layer)
    },
    // 定位
    markLocation(layer, geoJSON) {
      const type = geoJSON.geometry.type
      if (type == GeometryType.Polygon || type == GeometryType.LineString) {
        // get
        const bounds = layer.getBounds()
        this.eginMap.map.fitBounds(bounds)
      } else {
        const latlng = layer.getLatLng()
        this.eginMap.flyTo(latlng, 16)
      }
    },
    // 标记移除
    markRemove(guid) {
      //
      this.dataLists.map(item => {
        if (item.data.guid == guid) {
          // 删除图层中的数据
          this.removeLayer(item.layer)
        }
      })
      //
      this.removeDataList(guid)
      console.log(this.dataLists, 'removeList')
    },
    // 从dataList中删除
    removeDataList(guid) {
      //
      this.dataLists.map((item, key) => {
        if (item.data.guid == guid) {
          this.dataLists.splice(key, 1)
        }
      })
    },
    // 删除某个图层
    removeLayer(layer) {
      //
      const hasLayer = this.featureGroup.hasLayer(layer)
      //
      if (hasLayer) {
        console.log(layer.options, 'options')
        //
        // 需要判断是否已添加到库中 如果是则需要调用接口
        if(!layer.options.data.OBJECTID){
          //
          let delParams = {
            layerId: '',
            causeField: 'OBJECTID',
            causeValue: layer.options.data.OBJECTID
          }
          console.log(delParams)
          //
          // mapApi.deleteFeatures(delParams).then(res => {
          //   console.log('union#delete', res)
          // })
        }
       
        
        this.removeDataList(layer.options.data.guid)
        this.featureGroup.removeLayer(layer)
      }
      // 重置
      this.markerData = {}
    },

    // 确认
    confirm(layer, data, selectedIcon) {
      // 删除
      // this.tmpLayers.pop()
      // 如果data.id不存在则新增
      data.MARKER_NAME = this.currentMarker
      if (!data.GUID) {
        data.GUID = this.eginMap.Tools.guid()
        // 把数据添加到列表中
        let tmpData = {
          data: data,
          layer: layer,
          geoJSON: layer.toGeoJSON()
        }
        this.dataLists.push(tmpData)
        //
        layer.options.data = data
        // 添加到保存
        this.layers.push(layer)
        //
        this.featureGroup.addLayer(layer)
      } else {
        //
        layer.options.data = data
      }
      let geoJSON = layer.toGeoJSON()
      let feature = this.eginMap.Tools.geoJsonToArcgis(geoJSON)
      // console.log(geoJSON)
      // return
      // 获取新增或更新的图层ID
      let layerId = this.getBusinessId(geoJSON.geometry.type)
      let isUpdate = false
      // 如果有OBJECTID则说明是更新
      isUpdate = !geoJSON.properties.OBJECTID ? true: false
      //
      feature.attributes = data 
      feature.attributes.USER_ID = ''


      // 添加 需要有图层
      let params = {
        layerId: layerId,
        features: JSON.stringify([feature])
      }
      console.log('return' , console.log(params))
      // return
      //
      if(isUpdate){
        // mapApi.addFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }else{
        // mapApi.updateFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }
      
      this.currentLayer.options.icon = selectedIcon
      this.show = false
    },
    // 获取业务图层ID
    getBusinessId(type) {
      let layerId = ''
      switch(type){
       case 'Point':  layerId = this.pointLayerId;break;
       case 'LineString': layerId = this.polygonLayerId;break;
       case 'Polygon': layerId = this.polygonLayerId; break;
      }
      return layerId
    },
    clear() {
      // 清除未点击确定的数据
      this.featureGroup.clearLayers()
      //  console.log('clear')
      // this.tmpLayers.map((layer)=>{
      //   this.map.remove(layer)
      // })
    }
  }
}
</script>

<style lang="scss" scoped>
.mark-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255,255,255,0.8);
  border-radius: 4px;
  padding: 0px;
}
.mark-wrap .title {
  height: 50px;
  line-height: 50px;
  opacity: 0.8;
  background: #3f51b5;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
  // @include top-header-gradients;
}
.mark-tools {
  display: flex;
  background: #ffffff;
  border-radius: 10px;
  margin: 0px 5px;
}

.mark-wrap .item {
  cursor: pointer;
  background: #ffffff;
  margin: 0px 5px;
  border-radius: 10px;
  padding: 10px 10px;
  font-size: 14px;
}
</style>
```