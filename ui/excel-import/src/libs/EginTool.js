/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 13:47:56
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-15 10:17:30
 * @ Description: EginMap 工具类 不需要引用 map
 */


// GeometryType
import GeometryType from "./const/GeometryType"
// FeatureType
import FeatureType from "./const/FeatureType"

import * as Esri from 'esri-leaflet'
// import print from 'leaflet.browser.print'

class EginTool {
  // y,x  转 x,y
  latLngToLngLat(point) {
    return {
      x: point[1],
      y: point[0]
    }
  }
  // geoJson转arcgis
  geoJsonToArcgis(geoJson) {
    let geometry = []
    if (geoJson.type == FeatureType.FeatureCollection) {
      // 获取数据
      let features = geoJson.features
      // item
      features.forEach((item) => {
        geometry.push(this._getGeometry(item))
      })
    } else if (geoJson.type == FeatureType.Feature){
      geometry = this._getGeometry(geoJson)
    }
    return geometry
  }
  // arcgis 转 geoJSON
  arcgisToGeoJSON(geometry) {
    return Esri.Util.arcgisToGeoJSON(geometry)
  }
  // guid
  guid() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
      var r = Math.random() * 16 | 0,
        v = c == 'x' ? r : (r & 0x3 | 0x8);
      return v.toString(16);
    })
  }
  // 获取空间数据
  // 获取空间数据
  _getGeometry(feature) {
    // 
    let geometry = feature.geometry
    // 空间坐标数据
    let coordinates = geometry.coordinates
    // arcgis 数据格式
    // 面格式
    let arcgisFeature = {
      attributes: feature.properties, // 如果存储到arcgis里 还需要有非空的字段
      geometry: {
        rings: coordinates
      }
    }
    // 点位坐标转换
    if (GeometryType.Point === geometry.type) {
      arcgisFeature.geometry = {
        x: coordinates[0],
        y: coordinates[1]
      }
    }
    // 线格式
    if (GeometryType.LineString === geometry.type) {
      arcgisFeature.geometry = {
        paths: coordinates
      }
    }
    return arcgisFeature
  }
}

export default EginTool