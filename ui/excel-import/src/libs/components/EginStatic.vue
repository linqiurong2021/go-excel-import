<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 14:25:23
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:23:15
 * @ Description: 统计组件
 -->

<template>
 <div class="egin-static" v-show="show">
   <div>Static</div>
 </div>
</template>

<script>
export default {
  name: 'EginStatic',
  props: {
    show: {
      type: Boolean,
      default() {
        return false
      }
    }
  }
}
</script>

<style scoped>
.egin-static{
  width: 300px;
  position: absolute;
  /* left: 0; */
  height: 100%;
  background: #fff;
  z-index: 999;
  top: 0px;
}
</style>
