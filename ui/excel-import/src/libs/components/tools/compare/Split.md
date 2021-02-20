# 分屏组件

# 引入

::: tip
  注意引用的路径

  Split 组件(分屏比对组件) 需要与  SplitLayers 是分屏比对右上角的图层列表组件 共同使用
:::

```
import { Dialog } from "element-ui";
import EginMap from '../../../EginMap'
import EginPosition from '../../../const/EginPosition'
import TiandiTu from '../../../../libs/basemap/TiandiTu'
import SplitLayers from './SplitLayers'

```

# 使用说明

## 参数说明
  ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElDialog  | ElementUI 弹窗组件 |
      | Layers  | 分屏对比图层组件 |
  
  ### watch 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  | 弹窗是否可见 |d


  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  |  弹窗是否可见 |
      | left  | 左底图 |
      | right  | 右底图 |
      | maps | 左底图与右底图 |
      | showEginLayers  |  是否显示图层 |

  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | initMaps  | 初始化地图图 |
      | loadBasemap  |加载左、右底图 |
      | loadLayersWidget  | 加载左、右图层列表 |
      | close  | 关闭弹窗|


# 事例

```html

<template>
  <div class="compare-wrap" v-show="showTools">
    <Split  ref="splitDialog" />
  </div>
</template>

<script>

export default {
  name: "Compare",
  props: {
    eginMap: Object,
  },
  components: {
    Split: () => import("./compare/Split")
  }
};
</script>

```