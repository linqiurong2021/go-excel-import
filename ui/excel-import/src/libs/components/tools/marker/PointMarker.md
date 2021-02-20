# 点标绘组件

## 引入

::: tip
  注意引用的路径
:::

```
import {  Button, Input, Radio ,Form,FormItem } from 'element-ui'
import MyDialog from '../../dialog/index'

```

## 使用说明

### props 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | markerOptions  | 图例样式 |
      | data  | 表单详情数据 |
      | layer  | 当前图层 |


  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElForm  | ElementUI 表单组件 |
      | ElFormItem  | ElementUI 表单组件 |
      | Button  | ElementUI 按钮组件 |
      | ElRadio  | ElementUI Radio组件 |
      | ElInput  | ElementUI Input组件 |
      | MyDialog  | 自定义弹窗组件 |
  
  ### watch 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  表单数据 |
      | dialogVisible  |  是否显示弹窗 |


  ### data 参数说明
    
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


  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | resetTemp  | 重置表单数据 |
      | getStyle  | 获取样式 |
      | changeIcon  | 修改ICON |
      | remove  | 删除 |
      | confirm  | 确认 |
      | close  | 关闭弹窗 |



# 事例

```html
<template>
  <div>
      <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="markerConfirm"
      @remove="markerRemove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component>
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {
    // 标绘弹窗
    PointMarker: () => import('../libs/components/tools/marker/PointMarker'),
    PolylineMarker: () => import('../libs/components/tools/marker/PolylineMarker'),
    PolygonMarker: () => import('../libs/components/tools/marker/PolygonMarker'),
    TextMarker: () => import('../libs/components/tools/marker/TextMarker'),  
  },
  computed: {
    ...mapGetters({
      markTool: 'map/markTool'
    })
  },
  methods: {
    // 删除某个图层
    markerRemove(layer) {
    },
     // 确认
    markerConfirm(layer, data, selectedIcon) {
    }
  }
}
</script>

```