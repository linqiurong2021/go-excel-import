# 工具组件

## 引入
  无

## 使用说明

 
  ### props 参数说明

      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

  ### compoments 说明

      
        |  名称   |  说明  |
        |  ----  | ----  |
        | Compare  | 对比组件 |
        | GraphicEdit  | 图形编辑组件 |
        | Location  | 定位组件 |
        | Mark  | 标绘组件 |
        | Measure  | 测量组件 |
        | Print  | 打印组件 |

  ### data 参数说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | currentTool  |  当前组件 |
      | tools  | 工具 |
      | tools.name  | 工具项名称 |
      | tools.icon | 工具项icon |
      | tools.type | 工具项类型(组件名称) |
    

  ### 方法 说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | changeTools  | 工具切换 |
      | clearActivity  | 取消选中状态 |



# 事例

```html

<template>
  <div>
       <EginTools :show="true" :eginMap="eginMap" style="top:60px;" />
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {

    EginTools: () => import('../libs/components/EginTools'),
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