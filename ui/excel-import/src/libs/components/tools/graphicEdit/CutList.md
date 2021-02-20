# 切割列表组件

## 引入

::: tip
  注意引用的路径
:::

```
import MyDialog from '../../dialog/index'
import { Table, Button, TableColumn} from 'element-ui'

```

## 使用说明

  ### props 参数说明
      
        |  字段   |  说明  |
        |  ----  | ----  |
        | data  | 切割后的数据列表 |
        | field  | 列表显示字段名称 |
        | colorList  | 颜色列表 |


  ### components 组件说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | Table  | ElementUI 表格组件 |
      | Button  | ElementUI 按钮组件 |
      | TableColumn  | ElementUI 表格组件 |
      | MyDialog  | 自定义弹窗组件 |

  
  ### watch 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  切割后的数据列表 |


  ### data 参数说明
  
      |  字段   |  说明  |
      |  ----  | ----  |
      | visible  |  弹窗是否可见 |
      | tableData  | 表格数据列表 |
      | confirmData  | 选择保留后的数据 |


  ### 方法 说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | handleCurrentChange  | 表格行点击事件 |
      | close  | 关闭 |
      | cancleCut  | 取消切割 |
      | confirmCut  | 确认切割 |



# 事例

```html
<template>
  <div>
    <cut-list 
      ref="graphicDialog"
      :data="graphicEditData"
      :color-list="graphicEditColorList"
      field="OBJECTID">
    </cut-list>
  
  </div>
</template>

<script>

export default {
  name: 'EginMap',
  components: {
    CutList: () => import('../libs/components/tools/graphicEdit/CutList'), 
  },
  data() {
    return {
      graphicEditData: [],
      graphicEditColorList: []
    }
  }
}
</script>
```