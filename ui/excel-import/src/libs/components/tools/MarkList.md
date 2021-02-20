# 标绘列表工具组件(目前未使用)

## 引入
::: tip
  注意引用的路径
:::

```
import MyDialog from '../dialog/index'
import GeometryType from '../../const/GeometryType'

```

## 使用说明

  ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | list  | 标绘数据列表 |

  ### component 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | MyDialog  | 自定义弹窗 |
  

  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  | 是否显示标绘列表 |
      | markList  | 标绘数据列表 |
      | geometryType  | 标绘数据项类型(点、线、面、文本) |


  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | getIcon  | 获取标绘数据项的图标(icon) |
      | close  | 关闭标绘列表  |
      | location  | 定位 |
      | remove  | 移除 |

# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'MarkList',
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
    'MarkList': () => import('./tools/MarkList')
  },
  data() {
    return {
      currentTool: 'MarkList',
    }
  }
}
</script>

```