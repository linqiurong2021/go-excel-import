/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 14:54:31
 * @ Modified by: linqiurong
 * @ Modified time: 2020-09-16 16:18:26
 * @ Description: 谷歌地图
 */

import L from 'leaflet'
import ChinaMapType from '../const/ChinaMapType'

class Google {
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
      console.error("Google地图 Key 不能为空")
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
  addMap(options) {
    //
    return this.loadLayer(ChinaMapType.GoogleNormalMap, options)
  }

  // 卫星地图
  addSatelliteMap(options) {
    return this.loadLayer(ChinaMapType.GoogleSatelliteMap, options)
  }
  // 卫星注记图
  addSatelliteAnnotion(options) {
    return this.loadLayer(ChinaMapType.GoogleSatelliteAnnotion, options)
  }
}

export default Google