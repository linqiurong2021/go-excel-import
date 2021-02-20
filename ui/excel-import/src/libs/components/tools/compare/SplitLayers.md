# 分屏组件图层数据

# 引入

::: tip
  注意引用的路径

  SplitLayers 组件 需要与  Split组件  共同使用
:::

```
import { Tree } from 'element-ui' 
import layers from '../../../data/layers.json'
import EginLayer from '../../../../libs/EginLayer'
import layerConfig from '../../EginLayerConfig'

```

# 使用说明

## 参数说明

  ### props 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
      | show  | 是否显示 |
     
  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | Tree  | Tree ElementUI组件 |

  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | EginLayer  |  弹窗是否可见 |
      | layers  | 图层数据(一般是请求后台的数据) |
      | defaultProps.children  | Tree的子集 |
      | defaultProps.label | Tree 显示内容的字段名称 |
      | L  |  Leaflet |
      | allDataOfLayers  | 所有图层数据 |
      | layerCheckedIds | 已勾选的图层ID |

  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | loadLayers  | 添加图层到地图上 |
      | initLayers  |创建图层 并添加到地图 |
      | createFeatureLayers  | 加载图层 |
      | loadClusterFeature  | 加载聚合点服务|
      | loadWmsService  | 加载wms服务 |
      | loadRestService  | 加载Rest服务|
      | loadGeoJsonLayer  | 加载GeoJSON数据图层 |
      | checkChange  | 树形事件切换 |


# 事例

```html
<template>
  <div class="left" ref="left">
      <Layers v-if="map"  :eginMap="map" :show="showEginLayers" />
  </div>
</template>

<script>
import SplitLayers from './SplitLayers'
export default {
  name: "Split",
  props: {
    eginMap: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  components: {
    Layers: SplitLayers
  },
  data() {
    return {
      map: this.eginMap.map,
      // 是否显示图层
      showEginLayers: true
    };
  }
};
</script>
```