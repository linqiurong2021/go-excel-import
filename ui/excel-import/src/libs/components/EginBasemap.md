# 底图切换组件

## 引入
::: tip
  注意引用的路径
:::

```
import * as L from "leaflet";
import * as Esri from "esri-leaflet";
import "proj4";
import "proj4leaflet";
import defaultImage from '../../assets/img_map.png'

```

## 使用说明

   ### props 参数说明

      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |
        | loadXmTdt  | 是否加载厦门天地图历史底图 |
        | show  | 是否显示组件 |

  ### data 参数说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | showOrHide  |  组件是否可见 |
      | basemapLayers  | 底图数据 |
      | defaultImage  | 默认图片 |
      | activeIndex | 当前选中位置 |
      | imgLayers  | 厦门历史影像 |
      | vetorLayers | 厦门历史矢量 |



  ### 方法 说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | initCRS  | 设置crs 4490 |
      | createLayers  | 创建底图 |
      | initLayers  | 初始化图层  |
      | changeMap  | 底图切换 |

# 事例

```html

<template>
  <div>
    <EginBasemap v-if="eginMap" :eginMap="eginMap" :show="true" :loadXmTdt="true"/>
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {

    EginBasemap: () => import('../libs/components/EginBasemap'),
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