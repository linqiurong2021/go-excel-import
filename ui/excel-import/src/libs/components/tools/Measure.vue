<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:13:45
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 16:28:54
 * @ Description: 测量工具栏
 -->

<template>
 <div class="measure-wrap" v-show="showTools">
   <div class="title">{{title}}</div>
   <div class="measure-tools">
     <div class="item measure-item" v-for="(item,key) in tools" :key="key" @click="measure(item.type)">{{item.name}}</div>
   </div>
   <div class="measure-tools item clear" @click="clear">清除</div>
   <div class="measure-tools item close" @click="showTools = false">关闭</div>
 </div>
</template>

<script>
import './tooltip.js'
 export default {
   name: 'Measure',
   props: {
     eginMap: {
       require: true
     }
   },
   data() {
     return {
       map: null,
       title: '测量',
       showTools: true,
       drawPolyline: null,
       drawPolygon: null,
       drawType: null,
       layers: [],
       tools: [
        {
          name: '长度测量',
          type: 'length',
        },
        {
          name: '面积测量',
          type: 'area',
        }
       ]
     }
   },
   watch: {
     // 关闭
     showTools: function(newVal){
       if(!newVal){
        this.drawPolyline.disable()
        this.drawPolygon.disable()
       }
     }
   },
   mounted() {
    this.map = this.eginMap.map
    //
    this.drawPolyline = this.eginMap.MapTools.initDrawPolyline()
    this.drawPolygon = this.eginMap.MapTools.initDrawPolygon()
    // 获取结果
    this.getResult()
   },
   methods: {
     measure( type ) {
       type == 'length' ? this.length() :  this.area()
       this.drawType = type
     },
     length() {
       // 禁用画面
       this.drawPolygon.disable()
       // 开启画线
       this.drawPolyline.enable()
     },
     area() {
       // 禁用画本
       this.drawPolyline.disable()
       // 开启画面
       this.drawPolygon.enable()
     },
     //
     getResult() {
      // console.log
      this.map.off('draw:created').on('draw:created', (e)=>{
        const layer = e.layer;
        // 加入到数组中  删除时需要用到
        this.layers.push(layer)
        layer.addTo(this.map)
        // 判断是length 还是 area 需要计算长度或面积
        //
        const text = this.eginMap.MapTools.getPopupContent(layer)
        // const text = this.drawType == 'length' ? `长度:` : `面积`
        layer.bindTooltip(text)
      })
     },
     clear(){
       this.layers.forEach(layer => {
         if(this.map.hasLayer(layer)){
           this.map.removeLayer(layer)
         }
       })
     }
   }
 }
</script>

<style scoped>
.measure-wrap{
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
}
.measure-wrap .title{
    background: #3F51B5;
    opacity: 0.8;
    height: 50px;
    line-height: 50px;
    vertical-align: middle;
    align-items: center;
    text-align: center;
    color: #fff;
    /* border-radius: 5px; */
    margin-right: 10px;
    width: 85px;
    font-size: 14px
}
.measure-tools{
  display: flex;
  background: #fff;
  /* background: rgba(255, 255, 255, 0.7); */
  border-radius: 5px;
  margin: 0px 5px;
}

.measure-wrap .item{
  cursor: pointer;
  
  margin: 0px 5px;
  border-radius: 5px;
  padding: 10px 10px;
  font-size: 14px;
}

</style>