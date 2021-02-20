# 图形编辑工具组件

## 引入
::: tip
  注意引用的路径
:::

```
import { mapGetters } from 'vuex'
import clipPolygonUtil from '../../ClipPolygon'
import L from 'leaflet'
import { feature } from '@turf/turf';

```

## 使用说明

 ### props 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
  ### computed 参数说明
     
      |  字段   |  说明  |
      |  ----  | ----  |
      | editLayer  | 编辑的图形 |
      | graphicEditType  | 编辑的图形类型 |
      | getSelected  | 当前选择的图形(主要用于切割与合并) |
      | confirmCut  | 确认切割 |
      | confirmUnion  | 确认合并 |
      | cancleCut  | 取消切割 |
      | cancleUnion  | 取消合并 |
    
  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | layer  | 当前图形 |
      | layers  | 一个图形中有多个图形组合拆分后的图形存储 |
      | map  | 地图实例 |
      | featureGroup  | 切割或合并后的图形存放图层 |
      | drawTool  | 画图工具(主要是线用于做合并与切割还有整形) |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |
      | pane  | 图形存放的图层 |
      | errorMessage  | 错误信息 |
      | showError | 是否显示错误信息 |
      | colorList  | 颜色(切割与合并时会用到) |
      | searchLayerId  | 搜索图层ID(合并时用到) |
      | businessLayerId  | 空间数据保存时用到 |
      | searchOptions  | 搜索url |
      | layersGroups  | 编辑时有可能同时编辑多个图形 |
     
  ### watch 参数说明
     
      |  字段   |  说明  |
      |  ----  | ----  |
      | editLayer  | 编辑的图形 |
      | graphicEditType  | 编辑的图形类型 |
      | getSelected  | 当前选择的图形(主要用于切割与合并) |
      | confirmCut  | 确认切割 |
      | confirmUnion  | 确认合并 |
      | cancleCut  | 取消切割 |
      | cancleUnion  | 取消合并 |


  ### 方法 说明
    
      |  名称   |  说明  |
      |  ----  | ----  |
      | setTooltip  | 设置工具提示信息 |
      | authCheck  | 编辑权限判断 |
      | graphicEdit  | 选择哪个工具操作 |
      | cut  | 切割操作 |
      | remove  | 整形操作 |
      | union  | 合并操作 |
      | searchResult  | 合并时空间搜索列表 |
      | edit  | 编辑() |
      | confirm  | 编辑后确认 |
      | clear  | 清除 |
      | clearToolsStatus  | 清除工具状态 |


# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'GraphicEdit',
  props: {
    eginMap: Object,
    show: {
      type: Boolean,
      default() {
        return false
      }
    }
  },
  components: {
    'GraphicEdit': () => import('./tools/GraphicEdit')
  },
  data() {
    return {
      currentTool: 'GraphicEdit',
    }
  }
}
</script>

```