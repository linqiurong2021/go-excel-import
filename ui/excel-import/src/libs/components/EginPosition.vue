<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-29 09:04:58
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 17:49:43
 * @ Description: 显示当前坐标组件
 -->

<template>
 <div class="egin-position">
   坐标:({{point.x}}, {{point.y}})
 </div>
</template>

<script>
export default {
  name: 'EginPosition',
  props: {
    eginMap: {
      type: Object,
      require: true
    }
  },
  data() {
    return {
      point: {
        x: 0,
        y: 0
      }
    }
  },
  mounted() {
    // this.map = this.eginMap.map
    // console.log(this.map,'map')
    this.eginMap.map.on('mousemove',(evt)=>{
      let latlng = evt.latlng
      this.point = {
        x: latlng.lng,
        y: latlng.lat
      }
    })
  }
}
</script>

<style scoped>
.egin-position{
  left: 50px;
  bottom: 10px;
  z-index: 9999;
  font-size: 14px;
  color: #fff;
  padding: 5px;
  position: absolute;
}
</style>
