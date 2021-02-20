# 图形搜索工具组件

## 引入
  无

## 使用说明

  ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |
        | searchUrl  | 搜索链接 |
        | where  | 搜索条件 |


  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | currentToolName  | 当前工具名称 |
      | drawTool  | 画图工具 |
      | map  | 地图 |
      | pane  | 空间数据存放图层 |
      | searchLayerId  | 搜索图层ID(没有传searchUrl时使用) |
      | searchOptions.url  | 搜索图层链接(没有传searchUrl时使用) |
      | map  | 地图 |
      | pane  | 空间数据存放图层 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |


  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | initDrawTools  | 初始化工具 |
      | getDrawResult  | 获取画完结果 |
      | search  | 搜索 |
      | searchResult  | 搜索结果 |
      | graphicSearch  | 图形搜索按钮 |
      | drawToolsDisable  | 禁用工具 |
      | close  | 关闭 |
      | clear  | 清除 |


# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'GraphicSearch',
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
    'GraphicSearch': () => import('./tools/GraphicSearch')
  },
  data() {
    return {
      currentTool: 'GraphicSearch',
    }
  }
}
</script>

```