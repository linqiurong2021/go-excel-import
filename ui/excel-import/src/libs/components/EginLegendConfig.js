/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-21 15:26:52
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-23 16:46:12
 * @ Description: 图例工具配置文件
 */

//
import L from 'leaflet'
// 正常
const normal = L.divIcon({ className: 'normal iconfont icon-shengmingzhouqi', iconSize: [20, 20] })
// 异常
// const unnormal = L.divIcon({ className: 'unnormal iconfont icon-peizhi', iconSize: [20, 20] })
const earthquakeSymbol = 'https://esri.github.io/esri-leaflet/img/red-triangle.png';
const unnormal = L.icon({
  iconUrl: earthquakeSymbol,
  iconSize: [20, 20]
});

const eginLegendConfig = {
  18:[
    // 每一项
    {
      layerId: 18, // 确保这个值与前面的key一致
      // 类型
      geometryType: 'Point',
      // 图例显示名称
      name: '',
      // 字段名
      field: 'ZT',
      // 字段值
      value: '正常',
      // 样式
      style: {
        // 类型
        type: 'icon',
        // 样式值
        value: {
          icon: normal
        }
      }
    },
    {
      layerId: 18,
      // 类型
      geometryType: 'Point',
      // 图例显示名称
      name: '',
      // 字段名
      field: 'ZT',
      // 字段值
      value: '异常',
      // 样式
      style: {
        // 类型
        type: 'img', // 图片
        // 样式值
        value: {
          icon: unnormal
        }
      }
    }
  ],
  19: [
    // 每一项
    {
      layerId: 19,
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '正常',
      // 字段名
      field: 'ID',
      // 字段值
      value: '13330',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#53e033",//填充颜 #53e033
          opacity: 0.5,
          color: "#ff0000"
        }
      }
    },
    {
      layerId: 19,
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '异常',
      // 字段名
      field: 'ID',
      // 字段值
      value: '12574',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#53e033",//填充颜 #53e033
          opacity: 0.5,
          color: "#00ff00"
        }
      }
    }
  ],
  39: [
    // 每一项
    {
      layerId: 39,
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#53e033",//填充颜 #53e033
          opacity: 0.5,
          color: "#ff0000"
        }
      }
    }
  ],
  40: [
    {
      layerId: 40,
      // 类型
      geometryType: 'Point',
      // 图例显示名称
      name: '正常',
      // 字段名
      field: 'ZT',
      // 字段值
      value: '正常',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#53e033",//填充颜 #53e033
          opacity: 0.5,
          color: "#ff0000"
        }
      }
    },
    {
      layerId: 40,
      // 类型
      geometryType: 'Point',
      // 图例显示名称
      name: '异常',
      // 字段名
      field: 'ZT',
      // 字段值
      value: '异常',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#53e033",//填充颜 #53e033
          opacity: 0.5,
          color: "#0000ff"
        }
      }
    }
  ],
  42: [
    // 每一项
    {
      layerId: 42,
      // 类型
      geometryType: 'Point',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "",//填充颜 #53e033
          opacity: '',
          color: ""
        }
      }
    }
  ],
  43: [
    // 每一项
    {
      layerId: 43,
      // 类型
      geometryType: 'LineString',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "",//填充颜 #53e033
          opacity: '',
          color: ""
        }
      }
    }
  ],
  44: [
    // 每一项
    {
      layerId: 44,
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "",//填充颜 #53e033
          opacity: '',
          color: ""
        }
      }
    }
  ],
  45: [
    // 每一项
    {
      layerId: 45,
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "",//填充颜 #53e033
          opacity: '',
          color: ""
        }
      }
    }
  ],
  'default': [
    // 每一项
    {
      layerId: 'default',
      // 类型
      geometryType: 'Polygon',
      // 名称
      name: '',
      // 字段名
      field: '*',
      // 字段值
      value: '',
      // 样式
      style: {
        // 类型
        type: 'style',
        // 样式值
        value: {
          fillColor: "#ff0000",//填充颜 #53e033
          opacity: "0.9",
          color: "#ff0000"
        }
      }
    }
  ],
}

export default eginLegendConfig