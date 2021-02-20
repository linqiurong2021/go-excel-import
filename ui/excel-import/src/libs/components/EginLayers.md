# 图层列表组件

## 引入
::: tip
  注意引用的路径
:::

```
import { Tree } from 'element-ui'
import layers from '../data/layers.json'
import EginLayer from '../../libs/EginLayer'
import layerConfig from './EginLayerConfig'
import { mapGetters } from 'vuex'

```

## 使用说明

 ### props 参数说明

      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

        | showLegend  | 是否显示图例(暂无用通过勾选后就直接显示与隐藏) |

  ### data 参数说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | layers  |  图层列表数据(根据后台接口返回) |
      | defaultProps  | 树形结构 |
      | defaultProps.children  | 树形结构子项 |
      | defaultProps.label  | 树形结构显示名称 |
      | map | 地图实例 |
      | L | leaflet框架 |
      | allDataOfLayers | 图层数据 |

  ### 方法 说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | loadLayers  | 加载图层数据 |
      | initLayers  | 创建图层 并添加到地图 |
      | createFeatureLayers  | 加载图层 |
      | loadClusterFeature  | 加载聚合图层 |
      | loadWmsService  | 加载WMS图层 |
      | loadRestService  | 加载ArcgisMap图层 |
      | checkChange  | 图层勾选变更 |

# 事例

```html
<template>
  <div>
     <EginLayers v-if="eginMap" :eginMap="eginMap" :show="true" />
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {

    EginLayers: () => import('../libs/components/EginLayers'),
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