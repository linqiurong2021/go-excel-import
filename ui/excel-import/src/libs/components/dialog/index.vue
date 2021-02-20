<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-28 18:00:55
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:31:42
 * @ Description: 自定义弹窗组件
 -->

<template>
  <div v-show="visible" ref="dialog" class="my-dialog" :style="{width}">
    <div class="my-dialog-header"  @mousedown="dialogMouseDown" @mouseup="dialogMouseUp">
      <div>{{ title }}</div>
      <i class="iconfont icondelete" title="点击关闭" @click="close" />
    </div>
    <div class="my-dialog-body">
      <slot />
    </div>
    <div class="my-dialog-footer">
      <slot name="footer" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'MyDialog',
  props: {
    // 标题
    title: {
      type: String,
      default: '标题'
    },
    // 宽度
    width: {
      type: String,
      default: '35%'
    },
    // 显示弹窗
    visible: {
      type: Boolean,
      default: false
    },
    // 是否可拖动
    draggable: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    // 点击关闭
    close() {
      this.$emit('close')
    },
    // 鼠标按下时
    dialogMouseDown(event) {
      if (!this.draggable) return

      const Drag = this.$refs.dialog

      const ev = event || window.event

      event.stopPropagation()

      const disX = ev.clientX - Drag.offsetLeft
      const disY = ev.clientY - Drag.offsetTop
      
      document.onmousemove = function(event) {
        const ev2 = event || window.event
        Drag.style.left = ev2.clientX - disX + 'px'
        Drag.style.top = ev2.clientY - disY + 'px'
        Drag.style.cursor = 'move'
      }
    },
    // 鼠标放开时
    dialogMouseUp() {
      if (!this.draggable) return
      document.onmousemove = null
      this.$refs.dialog.style.cursor = 'default'
    }
  }
}
</script>

<style lang="scss" scoped>
.my-dialog {
  width: 50%;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  border-radius: 2px;
  z-index: 999;
  background-color: #ffffff;

  .icondelete{
    font-size: 20px;
    cursor: pointer;
  }

  >>> .el-form-item__label{
    color: #ffffff;
  }

  >>> input,textarea,.el-tree{
    color: #ffffff;
    background-color: transparent !important;
  }

  >>> .el-tree{
    background: transparent !important;
  }

  >>> .el-tree-node__label{
    color: #c0c4cc !important;
  }

  >>> .el-tree-node__content:hover {
    background-color: transparent;
  }

  >>> .el-input-group__append{
    color: #000 !important;
  }

  >>> .el-checkbox__label{
    color: #c0c4cc !important;
  }

  .my-dialog-header {
    color: #333333;
    font-size: 18px;
    padding: 20px 20px 10px;
    display: flex;
    background-color: #ffffff;
    justify-content: space-between;
    align-items: center;

    .close {
      cursor: pointer;
    }
  }

  .my-dialog-body {
    color: #333333;
    padding: 30px 20px;
  }

  .my-dialog-footer {
    display: flex;
    justify-content: flex-end;
    padding: 10px 20px 20px;
  }
}
</style>
