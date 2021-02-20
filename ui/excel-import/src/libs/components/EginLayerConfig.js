/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-21 11:03:21
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 08:21:01
 * @ Description: 图层配置文件 需要与 图例配置文件联合使用
 */


import L from 'leaflet'
// import GeometryType from '../const/GeometryType'
// legendClick
import legendConfig from './EginLegendConfig'

// import { getUserId } from '@utils/session' 

function getUserId() {
  return ''
}
// icon样式
function newDivIcon(iconName, size, html = '') {
  //
  let className = 'iconfont ' + iconName
  return L.divIcon({ className: className, iconSize: size, html: html })
}
// 
const getStyleConfig = (feature, layerId) => {
  let layerStyleConfig = legendConfig[layerId]
  let style = {}
  layerStyleConfig.forEach((item) => {
    if (item.field == '*') {
      style = item.style.value;
    } else {
      // 字段处理
      let fieldValue = feature.properties[item.field]

      if (layerId == 19 && feature.properties.ID == 13330) {
        console.log(fieldValue, 'field', item.value, '===', fieldValue == item.value)
        // let field = feature.properties[item.field]
        // console.error(style, '1606#style', fieldValue,'#styleValue', item.style.value)
      }
      // 如果值相等
      if (fieldValue == item.value) {
        style = item.style.value;
      }
    }
  })
  style.pane = 'pane_' + layerId
  // console.error('style',style)
  return style
}

// 图层配置文件 
const EginLayerConfig = (layerId) => {

  let config = {
    // 水保措施
    18: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      // 线面样式
      style: (feature) => {
        // feature
        return getStyleConfig(feature, layerId)
        // console.error(config, 'config')
        // return style
      },
      // 点位样式
      pointToLayer: (feature, latlng) => {
        // let icon = {}
        //ZT 这里使用自定义Icon 也可直接返回 L.circleMaker(xx,xx)
        // switch (feature.properties.ZT) {
        //   case '正常': icon = normal; break; //
        //   case '异常': icon = unnormal; break; // 图例
        // }
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      // 只返回有ZT数据的
      filter: () => {
        // 有状态的数据才返回
        // feature
        // if (feature.properties.ZT == '正常') {
        //   return feature
        // }
      },
      // 事件点击
      onEachFeature: (feature, layer) => {
        layer.on('click', () => {
          console.log(feature, 'feature')
        })
      },
      // 条件筛选
      where: "MAP_DIVISION_ID='350213'"
    },
    // 防治责任范围
    19: {
      operate: {
        query: true,// 搜索
        add: true, // 新增
        edit: true, // 编辑
        delete: true, // 删除
        cut: true, // 切割
        union: true, // 合并
        remove: true, // 整形
      },
      style: (feature) => {
        // let config = getStyleConfig(feature, layerId)
        // console.error(config, 'config#pointToLayer')
        //feature
        // console.log(feature,'feature')
        //feature
        // console.log(feature,'feature')
        // switch() {
        // case '' : return {};break;
        // case '' : return {};break;
        // }
        // 
        // return {
        //   fillColor: "#53e033",//填充颜 #53e033
        //   opacity: 0.5,
        //   color: "#ff0000"
        // }
        return getStyleConfig(feature, layerId)
        // return style
        // return L.marker(latlng, style);
      },
      onEachFeature: (feature, layer) => {
        // feature
        layer.on('click', () => {
          const id = feature.properties.ID
          // showTemplate
          console.log(id, 'id', 'layerId', layerId)
        })
      },
      pointToLayer: (feature, latlng) => {
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: (feature) => {
        if (feature.properties.ZT == '异常') {
          return feature
        }
      },
      where: "MAP_DIVISION_ID='350213'"
    },
    39: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => {
        return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      pointToLayer: (feature, latlng) => {
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: (feature) => {
        if (feature.properties.MAP_DIVISION_ID == '350213') {
          return feature
        }
      },
      where: "MAP_DIVISION_ID='350213'"
    },
    40: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => {
        return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      pointToLayer: (feature, latlng) => {
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: () => { },
      where: "MAP_DIVISION_ID='350213'"
    },
    // 标绘点
    42: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => { // 点位不需要这个
        return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      pointToLayer: (feature, latlng) => {
        let divIcon = newDivIcon(feature.properties.ICON_NAME, [20, 20], '');
        return L.marker(latlng, { icon: divIcon});
      },
      filter: () => { },
      where: `USER_ID='${getUserId()}'` // 当前用户的ID
    },
    // 文字标绘
    43: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => { // 文字标绘也不需要
        return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      pointToLayer: (feature, latlng) => {
        // 文字标绘样式 处理
        let html = `<div style="width:200px;font-size:${feature.properties.FONT_SIZE}px;font-weight:${feature.properties.FONT_WEIGHT};color:${feature.properties.FONT_COLOR}">${feature.properties.NAME}</div>`;
        let divIcon = newDivIcon("iconsoph_location", [0, 0], html);
        //
        return L.marker(latlng, {
          icon: divIcon,
        })
      },
      filter: () => { },
      where: `USER_ID='${getUserId()}'` //
    },
    44: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => {

       return {
          color: feature.properties.LINE_COLOR,
          weight: feature.properties.LINE_WIDTH,
          dashArray: feature.properties.LINE_STYLE == 'dash' ? 10 : ''
        }
        // return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      // 面不需要这个
      pointToLayer: (feature, latlng) => {
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: () => { },
      where: `USER_ID='${getUserId()}'` //
    },
    // 标绘面
    45: {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => {

        return {
          color: feature.properties.BORDER_COLOR,
          weight: feature.properties.BORDER_WIDTH,
          fillColor: feature.properties.FILL_COLOR,
          fillOpacity: feature.properties.FILL_OPACITY,
        }
        // return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      //
      pointToLayer: (feature, latlng) => {
        //
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: () => { },
      where: `USER_ID='${getUserId()}'` //
    },

    // 43: {
    //   operate: {
    //     query: false,// 搜索
    //     add: false, // 新增
    //     edit: false, // 编辑
    //     delete: false, // 删除
    //     cut: false, // 切割
    //     union: false, // 合并
    //     remove: false, // 整形
    //   },
    //   style: (feature) => {
    //     return getStyleConfig(feature, layerId)
    //   },
    //   onEachFeature: () => { },
    //   pointToLayer: (feature, latlng) => {
    //     const style = getStyleConfig(feature, layerId)
    //     return L.marker(latlng, style);
    //   },
    //   filter: () => { },
    //   where: "MAP_DIVISION_ID='350213'"
    // },
    // 44: {
    //   operate: {
    //     query: false,// 搜索
    //     add: false, // 新增
    //     edit: false, // 编辑
    //     delete: false, // 删除
    //     cut: false, // 切割
    //     union: false, // 合并
    //     remove: false, // 整形
    //   },
    //   style: (feature) => {
    //     return getStyleConfig(feature, layerId)
    //   },
    //   onEachFeature: () => { },
    //   pointToLayer: (feature, latlng) => {
    //     const style = getStyleConfig(feature, layerId)
    //     return L.marker(latlng, style);
    //   },
    //   filter: () => { },
    //   where: "MAP_DIVISION_ID='350213'"
    // },
    'default': {
      operate: {
        query: false,// 搜索
        add: false, // 新增
        edit: false, // 编辑
        delete: false, // 删除
        cut: false, // 切割
        union: false, // 合并
        remove: false, // 整形
      },
      style: (feature) => {
        return getStyleConfig(feature, layerId)
      },
      onEachFeature: () => { },
      pointToLayer: (feature, latlng) => {
        const style = getStyleConfig(feature, layerId)
        return L.marker(latlng, style);
      },
      filter: () => { },
      where: "MAP_DIVISION_ID='350213'"
    }
  }

  return config[layerId] ? config[layerId] : config['default']
}

// 1 获取配置文件 
// 2 判断配置文件是否有数据
// 3 加载配置文件并读取每图层的配置
// 4

export default EginLayerConfig