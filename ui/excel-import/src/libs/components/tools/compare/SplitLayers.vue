<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-28 10:05:12
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-29 09:22:21
 * @ Description: 
 -->

<template>
 <div class="egin-layers" v-show="show">
   <tree
    :data="layers"
    show-checkbox
    ref="tree"
    node-key="id"
    @check-change="checkChange"
    :default-expanded-keys="[2, 3]"
    :default-checked-keys="[5]"
    :props="defaultProps">
  </tree>
 </div>
</template>

<script>
import { Tree } from 'element-ui'
import layers from '../../../data/layers.json'
import EginLayer from '../../../../libs/EginLayer'
import layerConfig from '../../EginLayerConfig'

export default {
  name: 'SplitLayers',
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
    Tree
  },
  data() {
    return {
      EginLayer: null,
      layers: layers,
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      map: this.eginMap.map,
      L: this.eginMap.L,
      // 所有图层数据
      allDataOfLayers: [],
      // 已勾选的图层ID
      layerCheckedIds: []
    }
  },
  mounted() {
    // console.log("aaaa")
    this.EginLayer = new EginLayer({map: this.map})
    this.loadLayers(this.layers)
    this.initLayers()
  },
  methods: {
    // 添加图层到地图上
    loadLayers(layers) {
      // 图层数据处理 只获取有图层信息的数据
      layers.forEach(layer => {
        // console.log(layer.pg.LAYER_TYPE,'element')
        if(layer.pg.DIR_TYPE == 2) {
          // console.log('layer.pg.DIR_TYPE',layer.pg.DIR_TYPE)
          this.allDataOfLayers.push(layer)
        }else{
          // 有1与-1
          this.loadLayers(layer.children)
        }
      })
    },
    //
    // 创建图层 并添加到地图
    initLayers() {
      // console.log(this.allDataOfLayers,'allDataOfLayers')
      this.allDataOfLayers.forEach((item)=>{
        //
        let pg = item.pg
        // this.createFeatureLayers(pg)
      
        let layer = this.createFeatureLayers(pg)
        // console.log(layer, 'initLayers')
        if(layer){
          // 暂时不需要添加到地图上(勾选的时候再添加)
          // layer.addTo(this.map)
          // 添加到图层数据里
          this.eginMap.layers.push(layer)
        }
       
        return
      })
      // console.error(this.eginMap.layers,'layers')
    },
    // 加载图层
    createFeatureLayers(pg){
      const layerUrl = pg.LAYER_URL
      // wms server测试链接 
      // const layerUrl = 'http://dev.eginsoft.cn:6080/arcgis/services/fireService/FireService/MapServer/WMSServer'
      const layerId = pg.LAYER_TABLE
      // const layerType = pg.LAYER_TYPE
      let currentLayer = null
      let url = layerUrl+'/'+layerId

      let options = layerConfig(layerId, this.map)
      const pane = this.map.createPane('pane_'+layerId)   
      // 这里不能写死需要后台传参
      if(layerId == 18 || layerId == 44) {
        currentLayer = this.loadClusterFeature(url,options.pointToLayer, options.onEachFeature,options.filter,options.where, pane)
      }else{
        // loadRestService
        currentLayer = this.loadRestService(url, options.style,options.pointToLayer, options.onEachFeature, options.filter, options.where, pane)
      }
      // console.log(currentLayer.getPane(), 'currentLayer')
      // 添加属性
      currentLayer.layerName = pg.LAYER_NAME
      currentLayer.layerId = layerId
      
      return currentLayer
    },
    // 加载
    loadClusterFeature(url,pointToLayer,onEachFeature,filter,where){
      //
      const options = {
        url: url,
        // spiderfyOnMaxZoom: true, // 最大层级后不显示扩展线
        // removeOutsideVisibleBounds: true,
        onEachFeature: onEachFeature,
        pointToLayer: pointToLayer,
        where: where,
        filter: filter,
        disableClusteringAtZoom: 16,// 16级后取消聚合
      }
      const layer = this.EginLayer.addClusterLayer(options)
      return layer
    },
    // 加载wms服务
    loadWmsService(url, layerId) {
      let layer =  this.L.tileLayer.wms(url, {
          layers: [layerId], // 0,1,2,3,4,5,6,7
          format: 'image/png',
          transparent: true,
          attribution: ""
      })
      return layer
    },
    // 服务
    loadRestService(url, style,pointToLayer,onEachFeature,filter,where){
      //

      let options = {
        url: url,
        // 颜色配置
        style: style,
        where: where, //"MAP_DIVISION_ID='350213'",
        filter: filter,
        pointToLayer: pointToLayer,
        onEachFeature: onEachFeature,
      }
      // console.log(style,pointToLayer,onEachFeature,filter,where)
      let layer = this.EginLayer.addEsriFeatureLayer(options)

      return layer
    },
    
    // loadGeoJson 某个
    loadGeoJsonLayer(geojson) {
      //

      const style = function() {
        return {

        }
      }
      // 点样式
      const pointToLayer = function() {
        // 
      }
      // 循环每个图形
      const onEachFeature = function() {

      }
      // 条件过滤
      const filter = function() {
        // 
      }
      //
      const options = {
        filter: filter,
        onEachFeature: onEachFeature,
        pointToLayer: pointToLayer,
        style: style
      }
      this.EginLayer.addGeoJsonLayer(geojson,options)
    },
    // 选择变更
    checkChange: function(item,isChecked){
      //
      // 获取到子节点的数据 且 包括当前子节点的父结点
      let layerId = item.pg.LAYER_TABLE
      // 
      let tmpLayerCheckedIds = this.layerCheckedIds
      // 判断item是否存在
      this.eginMap.layers.forEach((item)=>{
        // 判断地图中是否已添加
        // const hasLayer =  this.map.hasLayer(item)
        let removeIndex = tmpLayerCheckedIds.indexOf(item)
        // console.log('removeIndex', removeIndex)
        if(layerId == item.layerId){
          isChecked ? tmpLayerCheckedIds.push(item) :tmpLayerCheckedIds.splice(removeIndex,1) // console.log("删除",  ) //layerCheckedIds.splice(index,1)
          isChecked ? item.addTo(this.map): this.map.removeLayer(item)
        }
      })
      this.layerCheckedIds = tmpLayerCheckedIds
    }
  },
}
</script>

<style scoped>
.egin-layers {
  z-index: 999;
  background: #fff;
  opacity: 0.9;
  top: 12px;
  position: absolute;
  right: 50px;
  /*测试用*/
  /* width: 300px;  */
  /* height: 300px; */
  border-radius: 5px;
  overflow: hidden;
  min-width: 220px;
  
}


</style>