# 当前位置组件

## 引入
  无

## 使用说明

  ### props 参数说明

      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  ### data 参数说明

    
      |  字段   |  说明  |
      |  ----  | ----  |
      | point  |  当前位置 |
      | point.x  | 经度 |
      | point.y  | 纬度 |



# 事例

```html

<template>
  <div>
       <EginPosition :show="true" v-if="eginMap" :eginMap="eginMap" />
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {

    EginPosition: () => import('../libs/components/EginPosition'),
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