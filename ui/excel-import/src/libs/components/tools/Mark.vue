<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:21
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 16:08:51
 * @ Description: 标绘工具栏
 -->

<template>
  <div class="mark-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="mark-tools">
      <div
        class="item mark-item"
        v-for="(item, key) in tools"
        :key="key"
        @click="mark(item.type, key)"
      >
        {{ item.name }}
      </div>
    </div>
    <!-- 20201016 与潘核实取消清除功能--->
    <!-- <div class="item clear" @click="clear">清除</div> -->
    <div class="mark-tools item close" @click="showTools = false">关闭</div>
    <!--弹窗-->
    <!-- <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="confirm"
      @remove="remove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component> -->
    <!-- <mark-list
      :list="dataLists"
      @remove="markRemove"
      @location="markLocation"
    ></mark-list> -->
  </div>
</template>

<script>
import GeometryType from '../../const/GeometryType'

// import { mapApi } from '@api/mapApi/map'

// 工具提示
import './tooltip.js'
import { mapGetters } from 'vuex'
export default {
  name: 'Mark',
  props: {
    eginMap: Object
  },
  computed: {
    ...mapGetters({
      markTool: 'map/markTool'
    })
  },
  data() {
    return {
      title: '标绘',
      // 存放标绘的图层
      featureGroup: null,
      // 图层显示
      pane: null,
      //
      showTools: true,
      // 当前选中是哪个组件
      currentMarker: 'PointMarker',
      // 点工具 线工具 面工具 文本工具
      drawTools: [],
      // 当前工具
      currentTool: [],
      // 当前操作的图层
      currentLayer: null,
      // 临时图层(未确认的)
      tmpLayers: [],
      // 已添加的图层(图形)(已确认的)
      layers: [],
      // 是否显示弹窗
      show: [false, false, false, false],
      // 存储后的数组
      dataLists: [],
      map: null,
      tools: [
        {
          name: '点标绘',
          type: 'PointMarker'
        },
        {
          name: '线标绘',
          type: 'PolylineMarker'
        },
        {
          name: '面标绘',
          type: 'PolygonMarker'
        },
        {
          name: '文字标绘',
          type: 'TextMarker'
        }
      ],
      // 标记的数据
      markerData: {},
      businessLayerId: 'a191c15515da4016b71681394e5a8479',
      //
      pointLayerId: 'e6e25e52f6384d52a6e941cfe7ad551d', // 标绘点
      //
      poylineLayerId: 'b197aba58a934a8ba576632f44938f54', // 标绘线
      //
      polygonLayerId: '7fc9aff62ea640bc82de56ebd4d458b0', // 标绘面
      //
      textLayerId: 'b7c72b2ed74c448ba1664bcd22dc8878 ', // 文字标绘
      // 7fc9aff62ea640bc82de56ebd4d458b0 标绘面
      // b197aba58a934a8ba576632f44938f54 标绘线
      // e6e25e52f6384d52a6e941cfe7ad551d 标绘点
      // b7c72b2ed74c448ba1664bcd22dc8878 文字标绘
    }
  },
  //
  mounted() {
    this.map = this.eginMap.map

    const [featureGroup, pane] = this.eginMap.MapTools.initPane('markPane', 501)
    this.pane = pane
    this.featureGroup = featureGroup
    featureGroup.addTo(this.map)

    this.initTools()
    // 监听新增结果
    this.getResult()
  },
  methods: {
    //
    initTools() {
      //
      const markerOptions = {
        //
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }

      const textMarkerOptions = {
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }
      const drawPoint = this.eginMap.MapTools.initMarker(markerOptions)
      const drawPolyline = this.eginMap.MapTools.initDrawPolyline()
      const drawPolygon = this.eginMap.MapTools.initDrawPolygon()
      const drawText = this.eginMap.MapTools.initTextMarker(textMarkerOptions)
      const drawTools = [drawPoint, drawPolyline, drawPolygon, drawText]

      this.drawTools = drawTools
    },
    mark(type, key) {
      // 当前组件
      this.currentMarker = type
      //
      this.drawTools.map((tool, index) => {
        index == key ? tool.enable() : tool.disable()
        if (index == key) {
          // 开启哪个工具
          this.currentTool = tool
        }
      })
    },
    getResult() {
      this.map.off('draw:created').on('draw:created', e => {
        const layer = e.layer
        // 当前图层
        this.currentLayer = layer
        // 加入到数组中  删除时需要用到
        // this.tmpLayers.push(layer)
        // 先添加后地图上  如果选择删除后 再删除
        // layer.addTo(this.map)

        this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker , currentLayer: layer , markerData: {} })
        console.error('getResult#showDialog')
        layer.on('click', evt => {
          //
          let data = evt.target.options.data
          // 弹窗相应弹窗
          this.currentMarker = data.MARKER_NAME
          this.markerData = data
          // 
          this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker, currentLayer: layer , markerData: data })
          //
          // setTimeout(() => {
          //   this.$refs.markerDialog.dialogVisible = true
          // }, 300)
          console.log(evt, data)
        })
        // 不在这里做数据添加 需要在确认时做数据添加 点击时需要获取数据框中的数据
        // this.featureGroup.addLayer(layer)

        // setTimeout(() => {
        //   this.$refs.markerDialog.dialogVisible = true
        // }, 300)
        // console.log(this.show, 'show')
      })
    },
    // 删除
    remove(layer) {
      // 当前layer在datalist的位置再删除
      this.removeLayer(layer)
    },
    // 定位
    markLocation(layer, geoJSON) {
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
    markRemove(guid) {
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
    // 从dataList中删除
    removeDataList(guid) {
      //
      this.dataLists.map((item, key) => {
        if (item.data.guid == guid) {
          this.dataLists.splice(key, 1)
        }
      })
    },
    // 删除某个图层
    removeLayer(layer) {
      //
      const hasLayer = this.featureGroup.hasLayer(layer)
      //
      if (hasLayer) {
        console.log(layer.options, 'options')
        //
        // 需要判断是否已添加到库中 如果是则需要调用接口
        if(!layer.options.data.OBJECTID){
          //
          let delParams = {
            layerId: '',
            causeField: 'OBJECTID',
            causeValue: layer.options.data.OBJECTID
          }
          console.log(delParams)
          //
          // mapApi.deleteFeatures(delParams).then(res => {
          //   console.log('union#delete', res)
          // })
        }
       
        
        this.removeDataList(layer.options.data.guid)
        this.featureGroup.removeLayer(layer)
      }
      // 重置
      this.markerData = {}
    },

    // 确认
    confirm(layer, data, selectedIcon) {
      // 删除
      // this.tmpLayers.pop()
      // 如果data.id不存在则新增
      data.MARKER_NAME = this.currentMarker
      if (!data.GUID) {
        data.GUID = this.eginMap.Tools.guid()
        // 把数据添加到列表中
        let tmpData = {
          data: data,
          layer: layer,
          geoJSON: layer.toGeoJSON()
        }
        this.dataLists.push(tmpData)
        //
        layer.options.data = data
        // 添加到保存
        this.layers.push(layer)
        //
        this.featureGroup.addLayer(layer)
      } else {
        //
        layer.options.data = data
      }
      let geoJSON = layer.toGeoJSON()
      let feature = this.eginMap.Tools.geoJsonToArcgis(geoJSON)
      // console.log(geoJSON)
      // return
      // 获取新增或更新的图层ID
      let layerId = this.getBusinessId(geoJSON.geometry.type)
      let isUpdate = false
      // 如果有OBJECTID则说明是更新
      isUpdate = !geoJSON.properties.OBJECTID ? true: false
      //
      feature.attributes = data 
      feature.attributes.USER_ID = ''


      // 添加 需要有图层
      let params = {
        layerId: layerId,
        features: JSON.stringify([feature])
      }
      console.log('return' , console.log(params))
      // return
      //
      if(isUpdate){
        // mapApi.addFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }else{
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
      
      this.currentLayer.options.icon = selectedIcon
      this.show = false
    },
    // 获取业务图层ID
    getBusinessId(type) {
      let layerId = ''
      switch(type){
       case 'Point':  layerId = this.pointLayerId;break;
       case 'LineString': layerId = this.polygonLayerId;break;
       case 'Polygon': layerId = this.polygonLayerId; break;
      }
      return layerId
    },
    clear() {
      // 清除未点击确定的数据
      this.featureGroup.clearLayers()
      //  console.log('clear')
      // this.tmpLayers.map((layer)=>{
      //   this.map.remove(layer)
      // })
    }
  }
}
</script>

<style lang="scss" scoped>
.mark-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255,255,255,0.8);
  border-radius: 4px;
  padding: 0px;
}
.mark-wrap .title {
  height: 50px;
  line-height: 50px;
  opacity: 0.8;
  background: #3f51b5;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
  // @include top-header-gradients;
}
.mark-tools {
  display: flex;
  background: #ffffff;
  border-radius: 10px;
  margin: 0px 5px;
}

.mark-wrap .item {
  cursor: pointer;
  background: #ffffff;
  margin: 0px 5px;
  border-radius: 10px;
  padding: 10px 10px;
  font-size: 14px;
}
</style>