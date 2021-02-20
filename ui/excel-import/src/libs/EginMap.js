/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 16:03:51
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-23 18:01:12
 * @ Description: 地图工具 EginMap 入口文件
 */

import L from 'leaflet'
import "leaflet/dist/leaflet.css"
import * as Esri from 'esri-leaflet'
// import EginLayer from './EginLayer';
import EginTools from './EginTool'
import EginPosition from "./const/EginPosition"
import "./css/eginmap.css"
// import EginBasemap from './basemap/EginBasemap';
require('leaflet-easyprint')
require("leaflet-draw")
// 卷帘
require('leaflet-side-by-side')
import "leaflet-draw/dist/leaflet.draw.css"

// 默认Icon
// import icon from 'leaflet/dist/images/marker-icon.png';

// import iconShadow from 'leaflet/dist/images/marker-shadow.png';
import EginMapTool from './EginMapTool'

class EginMap {
  // leaflet
  L = L
  // esri-leaflet
  Esri = Esri
  //
  map = null
  // 底图
  basemapLayers = []
  // 图层
  layers = []
  // 
  options = {}
  

  // 打印
  printTool = {}
  // 初始化
  constructor(options) {
    // container 参数不能为空
    if (!options.container) {
      console.error('map container should not be null')
      return
    }
    this.options = options
    // 设置默认点位图标
    this.setDefaultIcon()

    this.map = L.map(options.container, {
      zoomControl: false,
      attributionControl: false,
      // drawControl: true // 是不否显示画图工具
    })
    // 地图工具
    this.MapTools = new EginMapTool(this.map, {})
    // 图层操作
    // this.Layer = new EginLayer({ map: this.map })
    // this.Basemap = new EginBasemap({ tdt_key: 'tdt_key'})
    
    // this.Basemap.TiandiTu
    // console.log(this.map, 'map')
  }
  setDefaultIcon() {
    let DefaultIcon = L.icon({
      iconUrl: require('leaflet/dist/images/marker-icon.png'),
      shadowUrl: require('leaflet/dist/images/marker-shadow.png')
    });
    L.Marker.prototype.options.icon = DefaultIcon;
  }
  setView(latlng) {
    this.map.setView(latlng)
    return this
  }
  setZoom(zoom) {
    
    this.map.setZoom(zoom)
    return this
  }
  // 定位有动画
  flyTo( latlng, zoom = 14 ) {
    this.map.flyTo(latlng, zoom)
  }
  // 定位无动画
  panTo(latlng) {
    this.map.panTo(latlng)
  }
  // 设置缩放的位置
  //
  // zoomInTitle , zoomOutTitle, position
  setZoomPosition(options) {
    this.L.control.zoom(options).addTo(this.map)
  }
  // 加载底图切换控件
  // widget选中与不选中
  widgetActive(evt) {
    let className = evt.target.className
    let classList = className.split(' ')
    const clickI = classList.indexOf('egin-widget')
    // 如果没有egin-widget 则直接返回
    if ( clickI == -1) { // 点击到的是 i标签
      className = evt.target.parentNode.className
      console.log(className, 'cla')
      classList = className.split(' ')
    }
    //
    let index = classList.indexOf('active')
    if (index >= 0) {
      // 删除
      classList.splice(index, 1)
      console.log(classList, 'classList', index, 'index')
      // evt.target.className = classList.join(' ')
    } else {
      // 添加
      classList.push('active')
    }
    //
    clickI == -1 ? evt.target.parentNode.className = classList.join(' ') : evt.target.className = classList.join(' ')
  }

  loadBasempWidget(options = {}, changeEvent) {
    // 默认
    if (!options || !options.iconfont) options.iconfont = 'iconglobe' 
    if (!options || !options.position) options.position = EginPosition.BottomRight
    var widget = L.control({ position: options.position })
    // 加载底图列表
    widget.onAdd = () => {
      var div = L.DomUtil.create('div', 'egin-widget egin-basemap-widget leaflet-bar')
      // console.log(div,'div')
      div.innerHTML += `<i class="icon iconfont ${options.iconfont}"></i>`

      L.DomEvent.on(div, 'click', (evt) => {
        // 如果有 active 则删除 没有则添加
        this.widgetActive(evt)
        L.DomEvent.stopPropagation(evt)
        // 调用事件
        changeEvent(evt)
      })
      return div;
    }
    widget.addTo(this.map)
    // return widget
  }
  // 加载图层
  // options = {
  //   iconfont: '',
  //   position: ''
  // }
  loadLayersWidget(options = {}, changeEvent) {
    // console.log(options, changeEvent)
    if (!options || !options.iconfont) options.iconfont = 'iconglobe'
    if (!options || !options.position) options.position = EginPosition.TopRight
    var widget = L.control({ position: options.position })
    // 加载底图列表
    widget.onAdd = () => {
      var div = L.DomUtil.create('div', 'egin-widget egin-layers-widget leaflet-bar')
      // console.log(div,'div')
      div.innerHTML += `<i class="egin-icon iconfont ${options.iconfont}"></i>`

      L.DomEvent.on(div, 'click', (evt) => {
        this.widgetActive(evt)
        L.DomEvent.stopPropagation(evt)
        // 调用事件
        changeEvent(evt)
      })
      return div;
    }
    widget.addTo(this.map)
  }
  // 加载工具箱
  // options = {
  //   iconfont: '',
  //   position: ''
  // }
  loadToolsWidget(options = {}, changeEvent) {
    console.log(options, changeEvent)
    if (!options || !options.iconfont) options.iconfont = 'iconglobe'
    if (!options || !options.position) options.position = EginPosition.TopRight
    var widget = L.control({ position: options.position })
    // 加载底图列表
    widget.onAdd = () => {
      var div = L.DomUtil.create('div', 'egin-widget egin-tools-widget leaflet-bar')
      // console.log(div,'div')
      div.innerHTML += `<i class="egin-icon iconfont ${options.iconfont}"></i>`

      L.DomEvent.on(div, 'click', (evt) => {
        this.widgetActive(evt)
        L.DomEvent.stopPropagation(evt)
        // 调用事件
        changeEvent(evt)
      })
      return div;
    }
    widget.addTo(this.map)
  }
  // 加载工具箱
  // options = {
  //   iconfont: '',
  //   position: ''
  // }
  loadStaticWidget(options = {}, changeEvent) {
    // console.log(options, changeEvent)
    if (!options || !options.iconfont) options.iconfont = 'iconstar'
    if (!options || !options.position) options.position = EginPosition.TopRight
    var widget = L.control({ position: options.position })
    // 加载底图列表
    widget.onAdd = () => {
      var div = L.DomUtil.create('div', 'egin-widget egin-static-widget leaflet-bar')
      // console.log(div,'div')
      div.innerHTML += `<i class="egin-icon iconfont ${options.iconfont}"></i>`
        
      L.DomEvent.on(div, 'click', (evt) => {
        this.widgetActive(evt)
        L.DomEvent.stopPropagation(evt)
        // 调用事件
        changeEvent(evt)
      })
      return div;
    }
    widget.addTo(this.map)
  }
  // 加载比例尺
  loadScaleWidget() {
    L.control.scale().addTo(this.map);
  }
  // 显示当前位置坐标信息
  
  // 添加打印(有跨域问题需要解决)
  loadPrintWidget() {
    let printPlugin = L.easyPrint({
      sizeModes: ['Current', 'A4Landscape', 'A4Portrait'],
      filename: 'myMap',
      exportOnly: true,
      hidden: true, // 不显示控件
    }).addTo(this.map)

    this.printTool = printPlugin
  }

  
  // 导出地图 (使用前请确保 loadPrintWidget 已添加)
  exportMap(filename){
    this.printTool.printMap('CurrentSize', filename)
  }
  // 打印地图
  printMap() {
    window.print()
  }
  setCRS(crs) {
    this.map.crs = crs
  }
  // 分屏比对拖动与Zoom时同步
  splitSync(maps) {
    // 监听
    maps.map(function (t) {
      t.on({
        drag: () => {
          maps.map((map) => {
            map.setView(t.getCenter(), t.getZoom())
          })
        }, zoom: () => {
          maps.map((map) => {
            map.setView(t.getCenter(), t.getZoom())
          })
        }
      })
    })
  }
  // 卷帘
  wrap(leftLayer, rightLayer) {
    return L.control.sideBySide(leftLayer, rightLayer)
  }


  // 工具类
  Tools = new EginTools()
  
  // 图层操作类
  Layer = null
  // 底图操作
  Basemap = null
}

export default EginMap
