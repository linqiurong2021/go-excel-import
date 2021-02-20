/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 14:56:51
 * @ Modified by: linqiurong
 * @ Modified time: 2020-09-16 16:18:22
 * @ Description: 智图地图
 */

import L from 'leaflet'
import ChinaMapType from '../const/ChinaMapType'

class Geoq {
  // 初始化
  constructor (options) {
    // key  百度key
    this.options = options
  }
  options = {}
  // 已加载的底图
  layers = []
  // 添加地图
  loadLayer(type, options) {
    //
    if (!this.key) {
      console.error("智图 Key 不能为空")
      return
    }
    const layer = L.tileLayer.chinaProvider(
      type,
      {
        key: this.key,
        maxZoom: options.maxZoom ? options.maxZoom : 18,
        minZoom: options.minZoom ? options.minZoom : 4,
      }).addTo(this.eginMap)

    this.layers.push(layer)

    return layer
  }
  //
  addGeoqNormalMap(options) {
    //
    return this.loadLayer(ChinaMapType.GeoqNormalMap, options)
  }
  
  addGeoqNormalGray(options) {
    return this.loadLayer(ChinaMapType.GeoqNormalGray, options)
  }
  // 
  addGeoqNormalHydro(options) {
    return this.loadLayer(ChinaMapType.GeoqNormalHydro, options)
  }
  // 
  addGeoqNormalWarm(options) {
    return this.loadLayer(ChinaMapType.GeoqNormalWarm, options)
  }
  // 
  addGeoqNormalPurplishBlue(options) {
    return this.loadLayer(ChinaMapType.GeoqNormalPurplishBlue, options)
  }
}
export default Geoq
