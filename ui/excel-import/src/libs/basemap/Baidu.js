/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 14:54:31
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 15:11:21
 * @ Description: 百度地图
 */

import L from 'leaflet'
import ChinaMapType from '../const/ChinaMapType'

class Baidu {
  // 初始化
  constructor (options) {
    // key  百度key
    this.options = options
  }
  options = {}
  // 已加载的底图
  layers =  []
  // 添加地图
  loadLayer(type, options) {
    //
    if (!this.key) {
      console.error("百度地图 Key 不能为空")
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
  addBaiDuNormalMap(options) {
    //
    return this.loadLayer(ChinaMapType.BaiDuNormalMap, options)
  }

  // 卫星地图
  addBaiDuSatelliteMap(options) {
    return this.loadLayer(ChinaMapType.BaiDuSatelliteMap, options)
  }
  // 卫星注记图
  addBaiDuSatelliteAnnotion(options) {
    return this.loadLayer(ChinaMapType.BaiDuSatelliteAnnotion, options)
  }
}

export default Baidu