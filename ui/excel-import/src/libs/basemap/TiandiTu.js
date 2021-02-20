/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 14:51:35
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:14:27
 * @ Description: 天地图
 */

import L from 'leaflet'
import "leaflet.chinatmsproviders"
import ChinaMapType from '../const/ChinaMapType'

class TiandiTu {
  // 初始化
  constructor (options) {
    if (!options.key) {
      console.error("天地图 Key 不能为空")
      return
    }
    // key 天地图key
    this.options = options
  }
  options = {}
  // 已加载的底图
  layers = []
  // 加载地图
  loadLayer(type, options) {
    const key = !options || !options.key ? this.options.key : options.key
   
    const layer =  L.tileLayer.chinaProvider(
      type,
      {
        key: key,
        maxZoom: !options || !options.maxZoom ? 18: options.maxZoom,
        minZoom: !options || !options.minZoom ? 4: options.minZoom,
      })
    
    this.layers.push(layer)

    return layer

  }
  //
  addTianDiTuNormalMap(options) {
    return this.loadLayer(ChinaMapType.TianDiTuNormalMap, options)
  }
  // 天地图矢量
  addTianDiTuNormalAnnotion(options) {
    return this.loadLayer(ChinaMapType.TianDiTuNormalAnnotion, options)
  }
  // 卫星地图
  addTianDiTuSatelliteMap(options) {
    return this.loadLayer(ChinaMapType.TianDiTuSatelliteMap, options)
  }
  // 卫星注记图
  addTianDiTuSatelliteAnnotion(options) {
    return this.loadLayer(ChinaMapType.TianDiTuSatelliteAnnotion, options)
  }
  // 地形图
  addTianDiTuTerrainMap(options) {
    return this.loadLayer(ChinaMapType.TianDiTuTerrainMap, options)
  }
  // 地形图
  addTianDiTuTerrainAnnotion(options) {
    return this.loadLayer(ChinaMapType.TianDiTuTerrainAnnotion, options)
  }
}

export default TiandiTu