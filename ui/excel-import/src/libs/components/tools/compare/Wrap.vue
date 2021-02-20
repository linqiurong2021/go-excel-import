<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 18:28:30
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:32:53
 * @ Description: 卷帘对比组件 leftLayer 与 rightLayer 只能是 layer 不能是featureGroup
 -->

<template>
  <my-dialog
    class="wrap-wrap"
    title="卷帘对比"
    :visible="dialogVisible"
    :modal="false"
    width="90%"
    @close="close"
    draggable
  >
    <div class="content-container">
      <div class="left">
        <div class="top item">
          <div class="title">
            卷帘模式:
          </div>
          <div class="content">
            <el-radio v-model="type" label="0">垂直分割</el-radio>
            <el-radio v-model="type" label="1">水平分割</el-radio>
          </div>
        </div>
        <div class="middle item">
          <div class="title">
            选择遮罩层：
          </div>
          <div class="content layer-content">
            <el-tree
              :data="rightLayersData"
              show-checkbox
              ref="right"
              node-key="id"
              @check-change="rightChange"
              :props="defaultProps">
            </el-tree>
          </div>
        </div>
        <div class="bottom item">
          <div class="title">
            选择对比层:
          </div>
          <div class="content layer-content">
            <el-tree
              :data="leftLayersData"
              show-checkbox
              ref="left"
              node-key="id"
              @check-change="leftChange"
              :props="defaultProps">
            </el-tree>
          </div>
        </div>
      </div>
      <div class="right">
        <div class="map" ref="mapContainer"></div>
      </div>
    </div>
  </my-dialog>
</template>

<script>
import { Radio,Tree } from "element-ui"
import MyDialog from '../../dialog/index'
import TiandiTu from '../../../../libs/basemap/TiandiTu'
import * as L from 'leaflet'
import "leaflet-side-by-side"
import "leaflet/dist/leaflet.css"

export default {
  name: "Wrap",
  components: {
    MyDialog: MyDialog,
    // ElDialog: Dialog,
    ElRadio: Radio,
    ElTree: Tree
  },
  watch: {
    dialogVisible: function(newVal){
      console.log(newVal,'newVal')
     
      if(newVal && !this.eginMap){
        this.$nextTick(()=>{
           this.initMap()
        })
      }
    },
  },
  data() {
    return {
      dialogVisible: false,
      type: '0',
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      EginLayer: null,
      // 对比图层列表数据
      leftLayersData: [],
      // 遮罩图层列表数据
      rightLayersData: [],
      map: null,
      control: {},
      selectedLayers : [],
      leftLayer: null,
      rightLayer: null
    };
  },
  methods: {
    // 
    initLayers() {
      // 一组
      var osmLayer = L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png')
      osmLayer.name = 'OSM'
      osmLayer.id = 0
      var otherLayer = L.tileLayer('https://stamen-tiles-{s}.a.ssl.fastly.net/watercolor/{z}/{x}/{y}.png')
      otherLayer.name = '其它'
      otherLayer.id = 1
      const tiandiTu = new TiandiTu({key: '4b350b4f343fa22cdb2047e93b4d8712'})
      // 卫星图
      let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
      tdtSatellite.name = '天地图卫星'
      tdtSatellite.id = 0
      //
      let tdtNormal = tiandiTu.addTianDiTuNormalMap()
      tdtNormal.name = '天地图矢量'
      tdtNormal.id = 1

      // this.selectedLayers = [rightLayer, leftLayer, tdtSatellite, tdtNormal ]
      this.leftLayersData = [osmLayer, otherLayer]
      //
      this.rightLayersData = [tdtSatellite, tdtNormal]
      
      osmLayer.addTo(this.map)
      tdtSatellite.addTo(this.map)
      this.leftLayer = osmLayer
      this.rightLayer = tdtSatellite
      
      this.control = L.control.sideBySide(this.leftLayer, this.rightLayer).addTo(this.map)

      this.$refs.right.setCheckedKeys([0])
      this.$refs.left.setCheckedKeys([1])
    },
    initMap() {
      // 有可能无法做数据的卷帘 只能用底图类的(目前卷帘功能未完成)
      var map = L.map(this.$refs.mapContainer).setView([24.48, 118.108], 13)
      this.map = map
      this.initLayers()
    },
    change(index, type) {
      // 如果不存在则直接返回
      // 先取消全选 再选择当前
      if(type == 'right'){ // right 是右
        if(!this.rightLayersData[index]) return
        const layer = this.rightLayersData[index]
        layer.addTo(this.map)
        // 先删除后赋值
        this.map.removeLayer(this.rightLayer)
        this.rightLayer = layer
        // 清除其它选中项
        this.$refs.right.setCheckedKeys([index])
        this.control.setRightLayers(this.rightLayer);
      }else{
        if(!this.leftLayersData[index]) return

        this.map.removeLayer(this.leftLayer)
        const layer = this.leftLayersData[index]
        layer.addTo(this.map)
        this.leftLayer = layer
        this.$refs.left.setCheckedKeys([index])
        this.control.setLeftLayers(this.leftLayer)
      }
    },
    // 
    leftChange(obj,isChecked) {
      // console.log(a,b,c)
      if(isChecked) {
        this.change(obj.id,'left')
      }
    },
    // 
    rightChange(obj,isChecked) {
     if(isChecked) {
        this.change(obj.id,'right')
      }
    },
    close() {
      this.dialogVisible = false
    },
  },
};
</script>

<style scoped>

.content-container{
  display: flex;
  background: #ffffff !important
}
.item{
  margin: 5px;
  padding: 5px;
}
.left{
  width: 20%;
  background: #cccccc;
  height: 600px;
}
.left .content{
  display: flex;
  margin-top: 10px;
}
.right{
  width: 85%;
  background: #666666;
  height: 600px;
}
.right .map {
  height: 100%;
  width: 100%;
}
.title{
  padding: 10px 0px;
  border-bottom: 2px solid #fff;
  color: rgb(0, 118, 255)
}
.layer-content{
  min-height: 180px;
  max-height: 210px;
  border: 1px solid #fff;

  overflow: scroll;
}
.el-tree{
  height: 100%;
  width: 100%;
}
</style>