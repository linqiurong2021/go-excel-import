<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-10-14 09:00:04
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-29 10:24:07
 * @ Description: 标绘后的数据列表
 -->

<template>
 <my-dialog title="标记列表" class="mark-list-wrap"
    :visible="dialogVisible"
    width="10%"
    draggable
    @close="close">
      <div class="content list" v-if="markList.length > 0">
         <div class="list-item" v-for="(item) in markList" :key="item.data.guid"> 
            <div @click="location(item.layer, item.geoJSON)"><i class="item iconfont" :class="getIcon(item.geoJSON.geometry.type)"></i> {{item.data.name}}</div>
            <div class="item"><i class="iconfont icondelete" @click="remove(item.data.guid)"></i></div>
         </div>
      </div>
      <div v-else>
         未有标记记录
      </div>
   </my-dialog>
</template>

<script>
import MyDialog from '../dialog/index'
import GeometryType from '../../const/GeometryType'
export default {
   name: 'MarkList',
   components: {
      MyDialog
   },
   props: {
      list: {
         type: Array,
         default() {
            return []
         }
      }
   },
   data() {
      return {
         dialogVisible: true,
         markList: [],
         geometryType: GeometryType
      }
   },
   mounted() {
      this.markList = this.list
   },
   methods: {
      getIcon(type) {
         let icon = ''
         switch(type){
            case GeometryType.Polygon: icon='icontuxing';break;
            case GeometryType.LineString: icon='icontimeline';break;
            case GeometryType.Point: icon='iconsoph_location';break;
            default: {
               icon = 'iconword' ;break;
            }
         }
         return icon
      },
      //
      close() {
         this.dialogVisible = false
      },
      location(layer, geoJSON) {
         this.$emit('location', layer, geoJSON)
      },
      remove(guid) {
         this.$emit('remove', guid)
      }
   }
}
</script>

<style scoped>
.mark-list-wrap{
   opacity: 0.9;
   border-radius: 5px;
   overflow: hidden;
   min-width: 170px;
}
.mark-list-wrap .list{
   min-height: 100px;
   max-height: 60vh;
   overflow: scroll
}
.mark-list-wrap .list .list-item{
   line-height: 30px;
   display: flex;
   justify-content: space-between;
   border-bottom: 1px solid #e0dbdb
}
.list-item .iconfont{
   font-size: 20px;
   cursor: pointer;
}
.list-item .item{
   cursor: pointer;
}
</style>