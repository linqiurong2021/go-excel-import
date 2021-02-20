# 打印工具组件(未实现)

## 引入
  无

## 使用说明

  无

# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'Print',
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
    'Print': () => import('./tools/Print')
  },
  data() {
    return {
      currentTool: 'Print',
    }
  }
}
</script>

```