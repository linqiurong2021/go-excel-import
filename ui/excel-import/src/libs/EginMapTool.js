/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 18:45:15
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-20 09:02:49
 * @ Description: 地图工具 与 地图相关
 */

import L from 'leaflet'
import * as Esri from 'esri-leaflet'
class EginMapTool {
  // 未使用ToolBar
  constructor(map, options = null) {
    this.map = map
    // if(!options){
    //   this.initOptions(this.map)
    // }
    console.log(options)
    this.initOptions(this.map)
    // this.options = options

    // 先添加再隐藏
  }
  //
  newFeatureGroup() {
    const layer = new L.FeatureGroup()
    return layer
  }
  //
  newMakerIcon (options) {
    return L.Icon.extend({
      options: options
    });
  }
  initOptions(map) {
    // 编辑图层 
    // var editableLayers = this.newFeatureGroup();
    // map.addLayer(editableLayers);
    var editableLayers = new L.FeatureGroup();
    map.addLayer(editableLayers);
    // 画图层
    var drawLayers = this.newFeatureGroup()
    map.addLayer(drawLayers)

    var MyCustomMarker = this.newMakerIcon({
      shadowUrl: null,
      iconAnchor: new L.Point(12, 12),
      iconSize: new L.Point(24, 24),
      iconUrl: 'link/to/image.png'
    })

   

    let options = {
      position: 'topleft',
      // format: {
      // 	numeric: {
      // 		delimiters: {
      // 			thousands: ',',
      // 			decimal: '.'
      // 		}
      // 	}
      // },
      draw: {
        toolbar: {
          // #TODO: this should be reorganized where actions are nested in actions
          // ex: actions.undo  or actions.cancel
          actions: {
            title: '取消操作',
            text: '取消'
          },
          finish: {
            title: '完成操作',
            text: '完成'
          },
          undo: {
            title: 'Delete last point drawn',
            text: '删除最后一个点'
          },
          buttons: {
            polyline: '线',
            polygon: '面',
            rectangle: '矩正',
            circle: '圆',
            marker: '点',
            circlemarker: '圆点'
          }
        },
        handlers: {
          circle: {
            tooltip: {
              start: '点击画圆'
            },
            radius: '半径'
          },
          circlemarker: {
            tooltip: {
              start: 'Click map to place circle marker.'
            }
          },
          marker: {
            tooltip: {
              start: 'Click map to place marker.'
            }
          },
          polygon: {
            tooltip: {
              start: '点击地图开始',
              cont: '点击地图继续',
              end: '点击起点结束',
            }
          },
          polyline: {
            error: '<strong>Error:</strong> shape edges cannot cross!',
            tooltip: {
              
              start: '点击地图开始画线',
              cont: '继续点击地图继续画线',
              end: '点击最后一个点结束画线'
            }
          },
          rectangle: {
            tooltip: {
              start: 'Click and drag to draw rectangle.'
            }
          },
          simpleshape: {
            tooltip: {
              end: 'Release mouse to finish drawing.'
            }
          }
        },
        polyline: {
          shapeOptions: {
            color: '#f357a1',
            weight: 10
          }
        },
        polygon: {
          allowIntersection: false, // Restricts shapes to simple polygons
          drawError: {
            color: '#e1e100', // Color the shape will turn when intersects
            message: '<strong>Oh snap!<strong> you can\'t draw that!' // Message that will show when intersect
          },
          shapeOptions: {
            color: '#bada55'
          }
        },
        circle: false, // Turns off this drawing tool
        rectangle: {
          shapeOptions: {
            clickable: false
          }
        },
        marker: {
          icon: new MyCustomMarker()
        }
      },
      edit: {
        featureGroup: drawLayers,
        toolbar: {
          actions: {
            save: {
              title: 'Save changes',
              text: 'Save'
            },
            cancel: {
              title: 'Cancel editing, discards all changes',
              text: 'Cancel'
            },
            clearAll: {
              title: 'Clear all layers',
              text: 'Clear All'
            }
          },
          buttons: {
            edit: 'Edit layers',
            editDisabled: 'No layers to edit',
            remove: 'Delete layers',
            removeDisabled: 'No layers to delete'
          }
        },
        
        handlers: {
          edit: {
            tooltip: {
              text: 'Drag handles or markers to edit features.',
              subtext: 'Click cancel to undo changes.'
            }
          },
          remove: {
            tooltip: {
              text: 'Click on a feature to remove.'
            }
          }
        }
      }
    }
    this.options = options
  }
  initDrawPolyline() {
    return new L.Draw.Polyline(this.map)
  }
  //
  initDrawPolygon() {
    // drawOptions
    return new L.Draw.Polygon(this.map)
  }
  // 初始化文本编辑
  initTextMarker(options) {
    return new L.Draw.Marker(this.map, options) 
  }
  //
  initMarker(options) {
    return new L.Draw.Marker(this.map, options)
  }
  // 圆
  initCircle(options){
    return new L.Draw.Circle(this.map, options)
  }
  // 矩形
  initRectangle(options) {
    return new L.Draw.Rectangle(this.map, options)
  }
  // 
  initPane(name, index) {
    //
    let pane = this.map.createPane(name)
    // 需要设置zindex 否则会在下面看不到
    let featureGroup = L.featureGroup([], { pane: pane })
    //
    featureGroup.setZIndex(index)
    featureGroup.addTo(this.map)
    console.error('initPane',featureGroup, pane, index)
    return [featureGroup, pane]
  }
  
  // Helper method to format LatLng object (x.xxxxxx, y.yyyyyy)
  strLatLng(latlng) {
    return "(" + this._round(latlng.lat, 6) + ", " + this._round(latlng.lng, 6) + ")";
  }
  // Truncate value based on number of decimals
  _round(num, len) {
    let result =  Math.round(num * (Math.pow(10, len))) / (Math.pow(10, len));
    return result
  }
  getArea(latlngs) {
    // 返回数值未做单位转换
    let area = L.GeometryUtil.geodesicArea(latlngs);
    return area
  }
  // 
  query(options) {
    let task = Esri.query(options)
    return task
  }
 
  getPopupContent(layer) {
    // Marker - add lat/long
    if (layer instanceof L.Marker || layer instanceof L.CircleMarker) {
      return this.strLatLng(layer.getLatLng());
      // Circle - lat/long, radius
    } else if (layer instanceof L.Circle) {
      var center = layer.getLatLng(),
        radius = layer.getRadius();
      return "中心点: " + this.strLatLng(center) + "<br />"
        + "半径: " + this._round(radius, 2) + " 米";
      // Rectangle/Polygon - area
    } else if (layer instanceof L.Polygon) {
      let latlngs = layer._defaultShape ? layer._defaultShape() : layer.getLatLngs(),
      area = this.getArea(latlngs)
      area = area > 10000 ? (area / 10000).toFixed(4) + '公顷' : area.toFixed(4) + '平方米'
      return "面积: " + area;
      // Polyline - distance
    } else if (layer instanceof L.Polyline) {
      let latlngs = layer._defaultShape ? layer._defaultShape() : layer.getLatLngs(),
        distance = 0;
      if (latlngs.length < 2) {
        return "Distance: N/A";
      } else {
        for (var i = 0; i < latlngs.length - 1; i++) {
          distance += latlngs[i].distanceTo(latlngs[i + 1]);
        }
        //
        // console.log(distance, 'distance')
        distance = distance > 1000 ? (distance / 1000).toFixed(4) + "千米" : distance.toFixed(4) + '米'
        return "长度: " + distance ;
      }
    }
  }

  editCircle(shape, editOptions) {
    const options = !editOptions ? editOptions : this.options
    return new L.Edit.Circle(shape, options)
  }
  
  // icon样式
  newDivIcon(iconName, size, html = '') {
    //
    let className = 'iconfont ' + iconName
    return L.divIcon({ className: className, iconSize: size, html: html})
  }
  // 
  newMarker(latlng, options){
    return L.marker(latlng, options)
  }

  initDelete(options) {
    console.log(options)
    return new L.EditToolbar.Delete(this.map, this.options)

  }
}
export default EginMapTool