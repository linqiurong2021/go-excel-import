<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 13:38:06
 * @ Modified by: linqiurong
 * @ Modified time: 2021-02-20 20:11:17
 * @ Description: 地图测试组件
 -->

<template>
  <div>
    <div class="map-container" ref="mapContainer"></div>
    <EginBasemap v-if="eginMap" :eginMap="eginMap" :show="showEginBasemap" :loadXmTdt="true"/>
    <EginLayers v-if="eginMap" :eginMap="eginMap" :show="showEginLayers" />

    <EginTools :show="showEginTools" :eginMap="eginMap" style="top:60px;" />
    <EginStatic :show="showEginStatic" />
    <EginLegend :show="showEginLegend" v-if="eginMap" :eginMap="eginMap" />
    <!--当前原型上没有 暂时不显示-->
    <!-- <EginPosition v-if="eginMap" :eginMap="eginMap" /> -->
    <!--基本信息动态组件-->
    <component :is="baseInfo" ref="baseInfo"></component>
    <!--详情组件-->
    <component
      :is="detailInfo.componentsName"
      :data="detailInfo.data"
      ref="detailInfo"
    ></component>

    <!--切割弹窗列表-->
    <component
      ref="graphicDialog"
      :is="graphicEditComponent"
      :data="graphicEditData"
      :color-list="graphicEditColorList"
      field="OBJECTID"
    ></component>
    <!--标绘组件弹窗-->
    <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="markerConfirm"
      @remove="markerRemove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component>
    <!--标绘列表弹窗(暂时不显示)-->
    <!-- <mark-list
      :list="markerDataLists"
      @remove="markerListRemove"
      @location="markerLocation"
    ></mark-list> -->
    
    <!--提示-->
    <el-alert
      class="alert"
      :title="alertInfo.title"
      :type="alertInfo.type"
      center
      show-icon
      v-show="showAlert"
      @close="close"
    ></el-alert>
  </div>
</template>

<script>
// import EginMap from '../../libs//EginMap.js'
import EginMap from '../libs/EginMap'
import TiandiTu from '../libs/basemap/TiandiTu'
import EginPosition from '../libs/const/EginPosition'
import { Alert } from 'element-ui'
import ComponentNameMap from '../libs/const/ComponentName'
import { mapGetters } from 'vuex'
import GeometryType from '../libs/const/GeometryType'
// import { mapApi } from '@api/mapApi/map'

export default {
  name: 'EginMap',
  components: {
    ElAlert: Alert,
    EginBasemap: () => import('../libs/components/EginBasemap'),
    EginLayers: () => import('../libs/components/EginLayers'),
    EginTools: () => import('../libs/components/EginTools'),
    EginStatic: () => import('../libs/components/EginStatic'),
    EginLegend: () => import('../libs/components/EginLegend'),
    EginLineInfo: () => import('../libs/components/info/EginLineInfo'),
    EginPointInfo: () => import('../libs/components/info/EginPointInfo'),
    EginPolygonInfo: () => import('../libs/components/info/EginPolygonInfo'),
    Detail: () => import('../libs/components/info/Detail'),
    Detail1: () => import('../libs/components/info/Detail1'),
    Detail2: () => import('../libs/components/info/Detail2'),
    // Detail2: () => import('../libs/components/info/ProductDetail'),
    EginPosition: () => import('../libs/components/EginPosition'),
    UnionList: () => import('../libs/components/tools/graphicEdit/UnionList'),
    CutList: () => import('../libs/components/tools/graphicEdit/CutList'),
    // 标绘弹窗
    PointMarker: () => import('../libs/components/tools/marker/PointMarker'),
    PolylineMarker: () => import('../libs/components/tools/marker/PolylineMarker'),
    PolygonMarker: () => import('../libs/components/tools/marker/PolygonMarker'),
    TextMarker: () => import('../libs/components/tools/marker/TextMarker'),
    MarkList: () => import('../libs/components/tools/MarkList'),
    //
    
  },
  mounted() {
    // 获取行政区化
    this.getMapDivisionLocation()
    this.loadMap()
  },
  computed: {
    ...mapGetters({
      baseInfoComponent: 'business/getBaseInfoComponent',
      detailInfo: 'business/getDetailInfoComponent',
      alertInfo: 'alert/alertInfo',
      getSelectDialog: 'map/getSelectDialog',
      markTool: 'map/markTool',
      locateTo: 'map/locateTo'
    })
  },
  watch: {
    baseInfoComponent(newVal) {
      console.log(newVal, 'newVal')
      let layerId = newVal.layerId
      this.baseInfo = ComponentNameMap[layerId] ? ComponentNameMap[layerId] : ''
      if (this.baseInfo != '') {
        // 数据请求
        // mapApi.getPortList(params).then(res => {
        //   console.log(res, 'res')
        // })

        setTimeout(() => {
          this.$refs.baseInfo.dialogVisible = true
        }, 300)
      }
      // console.log(newVal ,'baseInfoComponent',this.baseInfo)
    },
    detailInfo(newVal) {
      if (newVal != '') {
        setTimeout(() => {
          this.$refs.detailInfo.dialogVisible = true
        }, 300)
      }
    },
    //
    alertInfo: {
      handler: function(newVal) {
        console.log(newVal, 'alertInfo')
        this.showAlert = true
        setTimeout(() => {
          this.showAlert = false
        }, 2000)
      },
      deep: true
    },
    // 这里不能用deep
    getSelectDialog: function(newVal) {
      this.graphicEditComponent = newVal.componentName
      this.graphicEditData = newVal.dataList
      this.graphicEditColorList = newVal.colorList
      console.log('dataList###', newVal.dataList, newVal.colorList)
    },
    // markTool: function(newVal){
    //   console.log(newVal, 'newVal')
    // }
    markTool: {
      handler: function(newVal){
        console.error('newVal', newVal)
        //
        setTimeout(() => {
          // 
         this.$refs.markerDialog.dialogVisible = true 
        }, 300)
      },
      immediate: true,
      deep: true
    },
    // 定位
    locateTo: function(newVal) {
      if (newVal == null) return
      let to = newVal.to // 定位点
      let level = newVal.level // 缩放级别
      let pinToMap = newVal.pinToMap // 是否在地图上闪烁
      //定位
      this.eginMap.flyTo(to, level)
      // 以下闪烁 添加到地图
      let map = this.eginMap.map
      let flashTimes = newVal.flashTimes ? newVal.flashTimes : 3
      if (pinToMap) {
        let timer = null, times = 0, layer = null;
        timer = setInterval(()=>{
          if ( times > flashTimes * 2) clearInterval(timer)
         
          if (times % 2 == 0) {
            layer = this.eginMap.L.marker(to)
            map.addLayer(layer)
          }else {
            map.removeLayer(layer)
          }
          times++
        }, 500)
        // 添加到地图上
        // this.$store.dispatch('map/setLocateTo', null)
      }
    }
  },
  data() {
    return {
      // EginMap
      eginMap: null,
      showEginBasemap: false,
      showEginLayers: false,
      showTools: false,
      showEginTools: false,
      showEginStatic: false,
      showEginLegend: false,
      // 是否显示提示信息
      showAlert: false,
      // 默认为 EginLineInfo
      baseInfo: '',
      // detailInfo: '', // 详情弹窗组件
      graphicEditComponent: '',
      graphicEditData: [],
      graphicEditShowField: 'ID',
      graphicEditColorList: [],
      // 标绘数据
      markerDataLists: [],
      // 存放标绘的图层
      markerFeatureGroup: null,
      //
      pointLayerId: 'e6e25e52f6384d52a6e941cfe7ad551d', // 标绘点
      //
      poylineLayerId: 'b197aba58a934a8ba576632f44938f54', // 标绘线
      //
      polygonLayerId: '7fc9aff62ea640bc82de56ebd4d458b0', // 标绘面
      //
      textLayerId: 'b7c72b2ed74c448ba1664bcd22dc8878 ', // 文字标绘
    }
  },
  methods: {
     // 获取区化坐标
    getMapDivisionLocation() {
      localStorage.setItem('extend', [118.42,24.48])
      // mapApi.findRingsOfMapDivision().then((res)=>{
      //   let { code,data} = res
      //   if(code === 200) {
      //     console.log("extend", res)
      //     localStorage.setItem('extend', data[0].extend)
      //   }
      // })
    },
    // 定位区化位置
    flyToMapDivisionLocation() {
      var point = localStorage.getItem("extend");
      if (point) {
        point = point.split(",");
        let params = {
          to: [point[1],point[0]],
          level: 12,
          pinToMap: false,
        }
        console.log(params,'params')
        this.$store.dispatch('map/setLocateTo',params)
      }
    },
    close() {
      this.showAlert = false
    },
    // 加载底图
    loadBasemap() {
      // 添加底图
      const tiandiTu = new TiandiTu({ key: '4b350b4f343fa22cdb2047e93b4d8712' })
      // 卫星图
      let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
      //  卫星注记图
      let tdtSatelliteAnnotion = tiandiTu.addTianDiTuSatelliteAnnotion()
      // 分组
      let tdtSatelliteGroup = this.eginMap.L.layerGroup([
        tdtSatellite,
        tdtSatelliteAnnotion
      ])
      // 添加 title 与 url 属性 为切换底图做准备
      tdtSatelliteGroup.name = '卫星图'
      tdtSatelliteGroup.url = ''
      // 添加到地图
      this.eginMap.map.addLayer(tdtSatelliteGroup)
      // 矢量图
      let tdtNormal = tiandiTu.addTianDiTuNormalMap()
      //  矢量注记图
      let tdttdtNormalAnnotion = tiandiTu.addTianDiTuNormalAnnotion()
      // 分组
      let tdtNormalGroup = this.eginMap.L.layerGroup([
        tdtNormal,
        tdttdtNormalAnnotion
      ])
      tdtNormalGroup.name = '矢量图'
      tdtNormalGroup.url = ''

      // 添加到底图里
      this.eginMap.basemapLayers.push(tdtSatelliteGroup)
      // 添加到底图里
      this.eginMap.basemapLayers.push(tdtNormalGroup)

      // let baseMap = {
      //   "卫星图": tdtSatelliteGroup,
      //   "矢量图": tdtNormalGroup
      // }
      // this.eginMap.L.control.layers(baseMap).addTo(this.eginMap.map)
    },
    // 加载缩放控件
    loadZoomWidget() {
      //
      let zoomPosition = {
        zoomInTitle: '放大',
        zoomOutTitle: '缩小',
        position: EginPosition.BottomRight
      }
      // 设置zoom控件的位置
      this.eginMap.setZoomPosition(zoomPosition)
    },
    // 加载图层列表
    loadLayersWidget() {
      const layersOptions = {
        iconfont: 'iconic_layers_px',
        position: EginPosition.TopRight
      }
      // 加载图层
      this.eginMap.loadLayersWidget(layersOptions, evt => {
        console.log(evt, 'loadLayersWidget')
        this.showEginLayers = !this.showEginLayers
      })
    },
    // 加载底图
    loadBasempWidget() {
      // 加载底图控件
      const basemapWidgetOptions = {
        iconfont: 'iconglobe',
        position: EginPosition.BottomRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadBasempWidget(basemapWidgetOptions, () => {
        // console.log(evt, 'loadBasempWidget')
        this.showEginBasemap = !this.showEginBasemap
        // console.log(this.showEginBasemap, 'showBasemapLayer')
      })
    },
    // 加载工具条
    loadToolsWidget() {
      // 加载底图控件
      const toolsWidgetOptions = {
        iconfont: 'icongongjuxiang',
        position: EginPosition.TopRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadToolsWidget(toolsWidgetOptions, evt => {
        console.log(evt, 'loadToolsWidget')
        this.showEginTools = !this.showEginTools
        console.log(this.showEginTools, 'showEginTools')
      })
    },
    // 加载统计条
    loadStaticWidget() {
      const loadWidgetOptions = {
        iconfont: 'iconstar',
        position: EginPosition.TopRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadStaticWidget(loadWidgetOptions, (evt) => {
        console.log(evt, 'loadStaticWidget')
        this.showEginStatic = !this.showEginStatic
        // console.log(this.showEginStatic, 'showEginStatic')
      })
    },
    // 加载所有
    loadWidgets() {
      // 加载控件
      this.loadZoomWidget()
      this.loadLayersWidget()
      this.loadBasempWidget()
      this.loadToolsWidget()
      this.loadStaticWidget()
      this.eginMap.loadPrintWidget()
      // 比例尺 
      // 原型上暂没有不显示
      // this.eginMap.loadScaleWidget()
    },
    // 加载地图
    loadMap() {
      // 加载地图
      const eginMap = new EginMap({ container: this.$refs.mapContainer })
      this.eginMap = eginMap
      eginMap.setView([24.48, 118.108]).setZoom(11)
      // 加载底图
      this.loadBasemap()
      this.loadWidgets()
      // 定位当前区化
      this.flyToMapDivisionLocation()
      // this.eginMap.map.on('click',(evt)=>{
      //   console.log(evt, 'click')
      // })
    },

    initMarker() {
      const [featureGroup, pane] = this.eginMap.MapTools.initPane('markPane', 501)
      console.log('pane', pane)
      this.markerFeatureGroup = featureGroup
      featureGroup.addTo(this.eginMap.map)

    },

    // 获取业务图层ID
    getLayerBusinessId(markerName) {
      let layerId = ''
      switch(markerName){
       case 'PointMarker':  layerId = this.pointLayerId;break;
       case 'PolylineMarker': layerId = this.poylineLayerId;break;
       case 'PolygonMarker': layerId = this.polygonLayerId; break;
       case 'TextMarker': layerId = this.textLayerId; break;
      }
      return layerId
    },

    // 删除某个图层
    markerRemove(layer) {
      //
      let properties = layer.feature.properties
      // 需要判断是否已添加到库中 如果是则需要调用接口
      console.log(layer,'layer')
      //
      const layerId = this.getLayerBusinessId(properties.MARKER_NAME)
      let delParams = {
        layerId: layerId,
        causeField: 'OBJECTID',
        causeValue: properties.OBJECTID
      }
      console.log(delParams)
      //
      // mapApi.deleteFeatures(delParams).then(res => {
      //   console.log('union#delete', res)
      //   const { code } = res
      //   code == 200 ? layer.remove() : ''
      // })

      
    },
     // 确认
    markerConfirm(layer, data, selectedIcon) {
      // 删除
      // 如果data.id不存在则新增
      if (!data.OBJECTID) {
        data.GUID = this.eginMap.Tools.guid()
        // 把数据添加到列表中
        let tmpData = {
          data: data,
          layer: layer,
          geoJSON: layer.toGeoJSON()
        }
        console.log(tmpData)
        layer.options.data = data
      } else {
        // 不能这样 有可能会有修改
        // let OBJECTID = layer.options.OBJECTID
        layer.options.data = data
      }
      let geoJSON = layer.toGeoJSON()
      let feature = this.eginMap.Tools.geoJsonToArcgis(geoJSON)
      // console.log(geoJSON,'geoJSON', data)
      // return
      // return
      // 获取新增或更新的图层ID
      let layerId = this.getLayerBusinessId(data.MARKER_NAME)
      let isUpdate = false
      // 如果有OBJECTID则说明是更新
      isUpdate = !geoJSON.properties.OBJECTID ? true: false
      //
      // console.log(data,'data')
      feature.attributes = data 
      feature.attributes.USER_ID = ''

      // return false;

      // 添加 需要有图层
      let params = {
        layerId: layerId,
        features: JSON.stringify([feature])
      }
      // console.log('return' ,JSON.stringify([feature]))
      // return
      //
      if(isUpdate){
        console.log(params)
        // mapApi.addFeatures(params).then((res)=>{
        //   // console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   code == 200 ? layer.addTo(this.eginMap.map): ''
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }else{
        console.log(params)
        // mapApi.updateFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }
      layer.options.icon = selectedIcon
    },
     // 定位
    markerLocation(layer, geoJSON) {
      const type = geoJSON.geometry.type
      if (type == GeometryType.Polygon || type == GeometryType.LineString) {
        // get
        const bounds = layer.getBounds()
        this.eginMap.map.fitBounds(bounds)
      } else {
        const latlng = layer.getLatLng()
        this.eginMap.flyTo(latlng, 16)
      }
    },
    // 标记移除
    markerListRemove(guid) {
      //
      this.dataLists.map(item => {
        if (item.data.guid == guid) {
          // 删除图层中的数据
          this.removeLayer(item.layer)
        }
      })
      //
      this.removeDataList(guid)
      console.log(this.dataLists, 'removeList')
    },
  }
}
</script>

<style  scoped>
.map-container {
  /* height:  calc(100vh - 70px); */
  height: 100vh;
  width: 100%;
}
.alert {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 500px;
  height: 35px;
  opacity: 0.8;
  z-index: 999;
}
/* .map-container >>> .egin-tools {
  top: 80px;
}
.map-container >>> .leaflet-control-container .leaflet-top.leaflet-right{
  top: 80px;
} */
/* .map-container
  >>> .leaflet-control-container
  >>> .leaflet-top
  >>> .leaflet-right {
  border: 1px solid red !important;
  top: 77px !important;
} */
</style>