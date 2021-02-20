# 标绘工具组件

## 引入
::: tip
  注意引用的路径
:::

```
import GeometryType from '../../const/GeometryType'
// 工具提示
import './tooltip.js'
import { mapGetters } from 'vuex'

```

## 使用说明

  ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  ### computed 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | markTool  | 标绘工具 |
  

  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 标题 |
      | featureGroup  | 标绘的图形存放图层 |
      | pane  | 图层显示 |
      | showTools | 是否显示标绘工具 |
      | currentMarker  | 当前工具项 |
      | drawTools  | 标绘工具实例(点工具 线工具 面工具 文本工具) |
      | currentTool  | 当前标绘工具 |
      | currentLayer  | 当前图层 |


      | tmpLayers  | 临时图层(未确认的) |
      | layers  | 已添加的图层(图形)(已确认的) |
      | show  | 是否显示弹窗[点弹窗、线弹窗、面弹窗、文本弹窗] |
      | drawTools  | 标绘工具实例(点工具 线工具 面工具 文本工具) |
      | dataLists  | 存储后的数组 |
      | map  | 地图实例 |
      | tools  | 标绘工具项 |
      | tools.name  | 标绘工具项名称 |
      | tools.type  |  标绘工具项类型(组件名称)|
      | markerData  | 弹窗填写的数据 |
      | pointLayerId  |  标绘点对应的空间数据业务图层ID|
      | poylineLayerId  | 标绘线对应的空间数据业务图层ID |
      | polygonLayerId |  标绘面对应的空间数据业务图层ID|
      | textLayerId  | 标绘文本对应的空间数据业务图层ID |

  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | initTools  | 初始化工具 |
      | mark  | 标绘项  |
      | getResult  | 标绘完成 |
      | remove  | 删除 |
      | markLocation  | 定位(列表使用-目前列表已取消)  |
      | markRemove  | 确认(列表使用-目前列表已取消) |
      | removeDataList  | 删除列表数据-目前列表已取消  |
      | removeLayer  | 删除图形 |
      | confirm  | 确认  |
      | getBusinessId  | 获取当前图形业务ID |
      | clear  | 清除-目前已取消  |

# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'Mark',
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
    'Mark': () => import('./tools/Mark')
  },
  data() {
    return {
      currentTool: 'Mark',
    }
  }
}
</script>

```