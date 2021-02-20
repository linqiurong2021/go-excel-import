# 图例组件

## 引入
::: tip
  注意引用的路径
:::

```
import { mapGetters } from 'vuex'
import legendConfig from './EginLegendConfig'
import LayerIndex from '../../libs/const/LayerIndex'
import { Tree, Slider, Tooltip } from 'element-ui'
```

## 使用说明

 ### props 参数说明

      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

  ### data 参数说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | maxLevel  |  最大层级 |
      | legendConfig  | 图例配置文件 |
      | layerOpcity  | 图层透明度,默认为非透明 |
      | map | 地图实例 |
      | L | leaflet框架 |
      | checkedLayers | 已选中的图层 |
      | maxZIndex  |  最大z-index |
      | minZIndex  | 最小z-index |
      | LayerIndex  | 图层z-index |
      | layerWhere | 图层条件 |
      | legendData | 图例数据 |
      | clickLayerId | 当前点击的是哪个图层 |
      | showSlider | 滑块 |
      | checkedIndex | 切换顺序时存储点击勾选时父图层 |

  ### watch 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | layerOpcity  | 图层透明度,默认为非透明 |
      | getLayerCheckedIds  | 已勾选图层id |
      | layerWhere | 条件筛选 |
      | clickLayerId | 点击的layerId |
     

  ### computed 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | getLayerCheckedIds  | 已勾选图层id |
      | showEginLegend | 是否显示图例 |

     
  ### 方法 说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | _getPane  | 获取Pane |
      | _setPaneIndex  | 设置pane#index |
      | changeIndex  | 修改层级 |
      | getChangeIndex  | 获取图层位置 |
      | changeArrIndex  | 修改数组顺序 |
      | getStyle  | 获取样式 |
      | checkChange  | 树形勾选事件 |
      | getTreeItemStyle  | 获取树形item样式 |


# 事例

```html

<template>
  <div>
       <EginLegend :show="true" v-if="eginMap" :eginMap="eginMap" />
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {

    EginLegend: () => import('../libs/components/EginLegend'),
  },
 
  data() {
    return {
      // EginMap
      eginMap: null,
    }
  }
}
</script>

<style  scoped>
```