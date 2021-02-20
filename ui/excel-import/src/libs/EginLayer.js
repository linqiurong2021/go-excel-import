/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 15:13:17
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:34:08
 * @ Description: 图层工具
 */

import L from 'leaflet'
import 'leaflet.markercluster'
import * as Cluster from "esri-leaflet-cluster"
import "leaflet.markercluster/dist/MarkerCluster.css"
import "leaflet.markercluster/dist/MarkerCluster.Default.css"
import * as Esri from 'esri-leaflet'

// 添加
class EginLayer {
  map = null
  options = {}
  layerGroups = []
  // 初始化
  constructor (options) {
    // console.log(options,'options')
    if(!options.map) {
      console.error("EginLayer 参数 map 不能为空")
      return
    }
    this.map = options.map
  }
  // 设置图层透明度
  setOpacity(id, opacity) {
    //
    let layers = L.control.layers
    layers.forEach((item)=>{
      console.log(item)
    })
    console.log(id, opacity)
  }
  // 分组
  layerGroup(id, group) {
    //
    let hasGroup = this.getLayerGroup(id)
    // 如果已有组则合并
    if (!hasGroup){
      group = hasGroup.contact(group)
    }else{
      this.layerGroups.push(group)
    } 
  }
  // 
  getLayerGroup(id) {
    let group = null
    this.layerGroup.forEach((item) => {
      if (item.id == id) {
       return group = this.map.remove(item)
      }
    })
    return group
  }
  // 通过ID
  removeGroup(id) {
    let group = this.getLayerGroup(id)
    if(!group){
      this.map.removeLayer(group)
    }
    return group
  }
  // 添加geoJson
  addGeoJsonLayer(geoJSON, options = {}) {
    // 
    L.geoJSON(geoJSON, {
      // 样式
      style: options.style,
      // 点位置 自定义样式
      pointToLayer: options.pointToLayer,
      // 绑定事件
      onEachFeature: options.onEachFeature,
      // 过滤
      filter: options.filter

    }).addTo(this.eginMap.map);
  }
  // 添加arcgis图层
  addEsriFeatureLayer(options= {}) {
    //
    return Esri.featureLayer(options)
  }
  // 添加到地图上
  addEsriFeatureLayerToMap(options = {}) {
    return Esri.featureLayer(options).addTo(this.map)
  }
  addTileLayer(options) {
    return L.addTileLayer(options).addTo(this.map)
  }

  // 添加聚合图层
  addClusterLayer(options) {
    // console.log(this.esri,'esri')
    return Cluster.featureLayer(options);
  }
  // 添加聚合图层
  addClusterLayerToMap(options) {
    // console.log(this.esri,'esri')
    return Cluster.featureLayer(options).addTo(this.map);
  }
  // 默认底图
  setBaseLayer(layerName) {
    L.basemapLayer(layerName).addTo(this.map)
  }
  // 设置RetinalLayer
  setRetinaLayer(layerName) {
    L.basemapLayer(layerName, {
      detectRetina: true
    }).addTo(this.map);
  }
  //  Vector 有问题
  setVectorLayer(layerName) {
    L.Vector.basemap(layerName).addTo(this.map)
  }
}
//
export default EginLayer