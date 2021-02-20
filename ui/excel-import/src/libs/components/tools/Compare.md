# 比较工具组件

## 引入

  使用动态引入

## 使用说明

  ### props 参数说明
      
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |

  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | Split  | 分屏比对组件(动态引入) |
      | Wrap  | 卷帘比对组件(动态引入) |

  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |

  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | compare  | 显示工具 |


# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'EginLayers',
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
    'Compare': () => import('./tools/Compare')
  },
  data() {
    return {
      currentTool: 'Compare',
    }
  }
}
</script>

```