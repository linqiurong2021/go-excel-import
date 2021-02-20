# 定位工具组件

## 引入
::: tip
  注意引用的路径
:::

```
import {  Button, Input, Form,FormItem } from 'element-ui'
import Dialog from '../../../libs/components/dialog/index'

```

## 使用说明

 ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElButton  | ElementUI按钮组件 |
      | ElInput  | ElementUI Input组件 |
      | ElForm  | ElementUI表单组件 |
      | ElFormItem  | ElementUI表单组件 |
      | MyDialog  | 自定义弹窗组件 |


  ### data 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  | 表单校验规则 |
      | dialogVisible  | 弹窗是否可见 |
      | temp  | 临时变量 |
      | temp.longitude  | 经度 |
      | temp.latitude  | 纬度 |

  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | handleClose  | 关闭 |
      | cancle  | 取消  |
      | confirm  | 确认 |

# 事例

```html

<template>
 <div class="egin-tools-wrap" v-show="show">
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
 </div>
</template>

<script>
export default {
  name: 'Location',
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
    'Location': () => import('./tools/Location')
  },
  data() {
    return {
      currentTool: 'Location',
    }
  }
}
</script>

```