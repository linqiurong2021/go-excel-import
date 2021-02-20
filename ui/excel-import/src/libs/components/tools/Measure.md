# 测量工具组件

## 引入
::: tip
  注意引用的路径
:::

```
import './tooltip.js'

```

## 使用说明

  ### props 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |

  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | map  | 地图实例 |
      | title  | 标题 |
      | drawPolyline  | 画线工具 |
      | drawPolygon  | 画面工具 |
      | drawType  | 画的类型 |
      | layers  | 存储图形 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型 |

  ### watch 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | showTools  | 是否显示弹窗 |


  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | measure  | 工具点击 |
      | length  | 长度测量  |
      | area  | 面积测量 |
      | getResult  | 图形结果 |
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
  name: 'Measure',
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
    'Measure': () => import('./tools/Measure')
  },
  data() {
    return {
      currentTool: 'Measure',
    }
  }
}
</script>

```