<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 13:38:03
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-27 16:39:09
 * @ Description: 图层列表
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
import layers from '../data/layers.json'
import EginLayer from '../../libs/EginLayer'
import layerConfig from './EginLayerConfig'
import { mapGetters } from 'vuex'


export default {
  name: 'EginLayers',
  props: {
    eginMap: Object,
    show: {
      type: Boolean,
      default() {
        return false
      }
    },
    // 图层配置列表
    layerConfig: {
      type: Function,
      default() {
        return ()=>{}
      }
    },
    // 图层数据(后台请求)
    layers: {
      type: Array,
      default() {
        return layers
      }
    },
    // 图例配置列表
    legendConfig: {
      type: Array,
      default() {
        return []
      }
    },
    //
    showLegend: {
      type: Boolean,
      default() {
        return true
      }
    }
  },
  components: {
    Tree
  },
  computed: {
    ...mapGetters({
      getLayerCheckedIds:  'map/getLayerCheckedIds',
      graphicEditType: 'map/graphicEditType'
    })
  },
  data() {
    return {
      EginLayer: null,
      // layers: layers,
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      map: this.eginMap.map,
      L: this.eginMap.L,
      // 所有图层数据
      allDataOfLayers: []
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

        // 切割测试用
        if(layer.layerId == 19){
          layer.addTo(this.map)
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
      const canEdit = pg.EDIT_ATTRIBUTE == 1 ? true: false
      // const canSearch = pg
      // const layerType = pg.LAYER_TYPE
      let currentLayer = null
      let url = layerUrl + '/' + layerId

      let options = layerConfig(layerId, this.map)
      // 重置
      options.operate = {
        query: canEdit,// 搜索
        add: canEdit, // 新增
        edit: canEdit, // 编辑
        delete: canEdit, // 删除
        cut: canEdit, // 切割
        union: canEdit, // 合并
        remove: canEdit, // 整形
      }
      const pane = this.map.createPane('pane_'+layerId)   
      

      // 点击获取数据
      const onEachFeature = (feature, layer)=>{
      
        // 单击查看详情
        layer.off('click').on('click', (evt)=>{
          //
          this.eginMap.L.DomEvent.stopPropagation(evt)
          //
          if(this.graphicEditType == 'edit'){
            // 如果事件是双击则禁用双击缩放编辑 图形
            // this.map.doubleClickZoom.disable();
            this.$store.dispatch('map/setEditLayer', layer)

          }else if(this.graphicEditType == 'cut' || this.graphicEditType == 'remove'){
            this.$store.dispatch('map/setEditLayer', layer)
          }
          // 如果处于编辑状态则直接返回
          if(this.graphicEditType == '') {
            this.eginMap.L.DomEvent.stopPropagation(evt)
            // console.log(feature, 'feature',evt)
            this.$store.dispatch('business/setBaseInfoComponent', {layerId: layerId, feature: feature, evt: evt}) 
          }
         
        })
      }

      // 这里不能写死需要后台传参
      if(layerId == 18 || layerId == 44) {
        currentLayer = this.loadClusterFeature(url,options.pointToLayer, onEachFeature ,options.filter,options.where,options.operate,layerId,pane)
      }else{
        // loadRestService
        currentLayer = this.loadRestService(url, options.style,options.pointToLayer, onEachFeature, options.filter, options.where, options.operate,layerId, pane)
      }
      // console.log(currentLayer.getPane(), 'currentLayer')
      // 添加属性
      currentLayer.layerName = pg.LAYER_NAME
      currentLayer.layerId = layerId
      
      return currentLayer
    },
    // 加载
    loadClusterFeature(url,pointToLayer,onEachFeature,filter,where,operate, layerId, pane){
      //
      const options = {
        url: url,
        // spiderfyOnMaxZoom: true, // 最大层级后不显示扩展线
        // removeOutsideVisibleBounds: true,
        onEachFeature: onEachFeature,
        pointToLayer: pointToLayer,
        where: where,
        filter: filter,
        operate: operate,
        layerId: layerId,
        pane:pane,
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
    loadRestService(url, style,pointToLayer,onEachFeature,filter,where, operate,layerId,pane){
      //

      let options = {
        url: url,
        // 颜色配置
        style: style,
        where: where, //"MAP_DIVISION_ID='350213'",
        filter: filter,
        operate: operate,
        layerId: layerId,
        pointToLayer: pointToLayer,
        onEachFeature: onEachFeature,
        pane:pane,
      }
      // console.log(style,pointToLayer,onEachFeature,filter,where)
      let layer = this.EginLayer.addEsriFeatureLayer(options)

      return layer
      // layer.addTo(this.map)
    //   if(!pg.dataInfo){
    //     pg.dataInfo = {
    //       iconfont: {},
    //       style: {}
    //     }
    //   }
    //   // 配置默认的图标信息
    //   if(!pg.dataInfo || !pg.dataInfo.iconfont) {
    //     const iconfont = {
    //       icon: 'iconic_location_on_px',
    //       size: [20,20]
    //     }
    //     pg.dataInfo.iconfont = iconfont
    //   }
    //   // 默认style
    //   if(!pg.dataInfo.style) {
    //     //
    //     const defaultStyle = {
    //       color: "#ff0000", opacity: 0.5
    //     }

    //     pg.dataInfo.style = defaultStyle
    //   }
    

      // console.error(pg.dataInfo,'dataInfo')
      // 获取点位数据坐标
      // const iconfont = pg.dataInfo.iconfont
      // const icon = new EginIcon(iconfont)
      // console.log(icon, 'icon')
      // url = 'http://dev.eginsoft.cn:6080/arcgis/rest/services/clone/ClusterService/MapServer/0'
      // let options = {
      //   url: url,
      //   pointToLayer:  (geojson, latlng) =>{
      //     // console.log(geojson, 'geojson')
      //     let earthquakeSymbol = 'https://esri.github.io/esri-leaflet/img/red-triangle.png';
      //     var icon = this.eginMap.L.icon({
      //       iconUrl: earthquakeSymbol,
      //       iconSize: [20, 20]
      //     });

      //     return this.eginMap.L.marker(latlng, {
      //       icon: icon
      //     });
      //   }
      // }
      // this.eginMap.Esri.featureLayer(options).addTo(this.map)


      // // 参数
      // let options = {
      //   url: url,
      //   pointToLayer: (geojson, latlng)=>{
      //     return this.eginMap.L.marker(latlng, {
      //       icon: icon
      //     });
      //   },
      //   filter: ()=>{
      //     //
      //     console.log("filter")
      //   }
      // }
      // let layer = this.EginLayer.addFeatureLayer(options)
      // console.log( this.EginLayer,'Layer')
      // // let layer = this.eginMap.Layer.featureLayer(options)
      // console.log(layer,'layer')
      // // let layer = this.eginMap.L.tileLayer(options)
      // // this.map.addLayer(layer)

      // // console.log(layer,'layer')
      // console.log(options,'options')
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
      let layerCheckedIds = this.getLayerCheckedIds
      // 判断item是否存在
      this.eginMap.layers.forEach((item)=>{
        // 判断地图中是否已添加
        // const hasLayer =  this.map.hasLayer(item)
        let removeIndex = layerCheckedIds.indexOf(item)
        // console.log('removeIndex', removeIndex)
        if(layerId == item.layerId){
          isChecked ? layerCheckedIds.push(item) :layerCheckedIds.splice(removeIndex,1) // console.log("删除",  ) //layerCheckedIds.splice(index,1)
          isChecked ? item.addTo(this.map): this.map.removeLayer(item)
        }
      })
      console.log(layerCheckedIds, 'layerCheckIds')
      // 修改值
      this.$store.dispatch('map/setLayerCheckedIds', layerCheckedIds)
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