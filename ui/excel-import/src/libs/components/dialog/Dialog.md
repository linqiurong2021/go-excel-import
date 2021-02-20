# 弹窗组件

## 引入

  无

## 使用说明

  ### props 参数说明
    
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 标题 |
      | width  | 弹窗宽度 |
      | visible  | 是否可见(显示与隐藏) |
      | draggable  | 是否可拖拽 |


  ### methods 方法说明
    
      |  方法名   |  说明  |
      |  ----  | ----  |
      | close  | 关闭弹窗 |
      | dialogMouseDown  | 鼠标按下时(拖拽) |
      | dialogMouseUp  | 鼠标放开时(拖拽) |


# 事例

```html

<template>
    <my-dialog title="详情" class="baseinfo-wrap"
    :visible="dialogVisible"
    width="80%"
    draggable
    @close="close">
      <div class="content">
        
     </div>
      
    </my-dialog>

</template>

<script>
import MyDialog from '../../../libs/components/dialog/index'
export default {
  name: "Detail",
  props: {
    data: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  components: {
    MyDialog: MyDialog
  },
  data() {
    return {
      dialogVisible: true,
    };
  },
  methods: {
    close() {
      this.dialogVisible = false;
    },
    handleClick() {
      console.log('handleClick')
    },
  },
};
</script>
}
</script>

```