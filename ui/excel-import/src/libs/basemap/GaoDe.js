/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 14:56:51
 * @ Modified by: linqiurong
 * @ Modified time: 2020-09-16 16:18:19
 * @ Description: 高德地图
 */

import L from 'leaflet'
import ChinaMapType from '../const/ChinaMapType'

class GaoDe {
  // 初始化
  constructor(options) {
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
      console.error("高德地图 Key 不能为空")
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
  addGaoDeNormalMap(options) {
    //
    return this.loadLayer(ChinaMapType.GaoDeNormalMap, options)
  }

  // 卫星地图
  addGaoDeSatelliteMap(options) {
    return this.loadLayer(ChinaMapType.GaoDeSatelliteMap, options)
  }
  // 卫星注记图
  addGaoDeSatelliteAnnotion(options) {
    return this.loadLayer(ChinaMapType.GaoDeSatelliteAnnotion, options)
  }
}

export default GaoDe