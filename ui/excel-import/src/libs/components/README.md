# EginBasemap 底图切换工具

> 参数说明

  >> props 参数说明

    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |
        | loadXmTdt  | 是否加载厦门天地图历史底图 |
        | show  | 是否显示组件 |

  >> data 参数说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | showOrHide  |  组件是否可见 |
      | basemapLayers  | 底图数据 |
      | defaultImage  | 默认图片 |
      | activeIndex | 当前选中位置 |
      | imgLayers  | 厦门历史影像 |
      | vetorLayers | 厦门历史矢量 |



  >> 方法 说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | initCRS  | 设置crs 4490 |
      | createLayers  | 创建底图 |
      | initLayers  | 初始化图层  |
      | changeMap  | 底图切换 |

> 组件代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 11:01:33
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-22 16:50:20
 * @ Description: 底图切换工具
 -->

<template>
 <div class="egin-basemap-layer" v-show="show">
   <div class="egin-basemap-layer-item" :class="{'active': activeIndex == key }" v-for="(item,key) in eginMap.basemapLayers" :key="key" @click="changeMap(item._leaflet_id, key)">
     <div class="img"><img :src="item.url ? item.url : defaultImage" :alt="item.name"/></div>
     <div class="name">{{item.name}}</div>
   </div>
 </div>
</template>

<script>
import * as L from "leaflet";
import * as Esri from "esri-leaflet";
import "proj4";
import "proj4leaflet";


import defaultImage from '../../assets/img_map.png'
export default {
  name: "EginBasemap",
  props: {
    eginMap: Object,
    loadXmTdt: {
      type: Boolean,
      default() {
        return false
      }
    },
    show: {
      type: Boolean,
      default() {
        return false
      }
    }
  },
  data() {
    return {
      showOrHide: false,
      basemapLayers: this.eginMap.basemapLayers,
      defaultImage: defaultImage,
      activeIndex: 0,
      imgLayers: [
        // {name: '2012厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2012/MapServer',token: 'KhTlmtGizEANsimdOfkRQq3OIGH9iqNIX185zLci6ttGelMXSwd8xKTJq4CpAcSNWZCpbLRxu8P5DJuuZPCdJw..'},
        {name: '2013厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2013/MapServer',token: 'KhTlmtGizEANsimdOfkRQsuJQg281bfxM8NO1dUY5-xgzyE3yMXWoB13uuZbaePR9YjhIgRmdVaO16DWUJQ0fA..'},
        {name: '2014厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2014/MapServer',token: 'KhTlmtGizEANsimdOfkRQi0ridpq68lVjRQD65FXxDyDUo9RmMLpYT51B_3RYFw7o_wI7uQ6wWQ27DAZHGi2uA..'},
        {name: '2015厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2015/MapServer',token: 'KhTlmtGizEANsimdOfkRQoDxUbJfXA6VI4xvQlgRjCrxJSI4BjFKZUsDovfQoROTpPTkkfh6r33We9CVJgxYSw..'},
        {name: '2016厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2016/MapServer',token: 'KhTlmtGizEANsimdOfkRQmRinbEn_YyeDj-qDSV0YwYNkPBmQ--7fIVoVqkYqCPwGacBQsg-skfIbZhiXFBYAA..'},
        {name: '2017厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2017/MapServer',token: 'KhTlmtGizEANsimdOfkRQna9aPia93yQlcne1xjc8bCpNAtA5Yq7PAG2KPPEnksiOJ8JEGIAede1lUXHpE7H2Q..'},
        {name: '2018厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2018/MapServer',token: 'KhTlmtGizEANsimdOfkRQixA4mmyzgtWNc19FKxn1LzM36kSNKFWDx9se7GLnxPe670-FWKjp1iWQ2Jnto_unw..'},
        {name: '2019厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/DOM2019/MapServer',token: 'KhTlmtGizEANsimdOfkRQk92vRFIHvGsM8BTS11OU9VUMzLFDlqVsSJ0b0qKKOBpl3bVWy0VkXBq2A-vm-tMMg..'},
        {name: '2020厦门影像',url: 'http://222.76.242.138/arcgis/rest/services/CGCS_DOMMAP/MapServer',token: 'KhTlmtGizEANsimdOfkRQlOfF2AR4wy2NQzfAFl-DVj6t0LIpw7hyP1GBsvAeRavVeWBB1BnQ6rp8LWe7G0tGg..'},
      ],
      vetorLayers: [
        {name: '2018厦门晕渲',url: 'http://222.76.242.138/arcgis/rest/services/CGCS_DEMMAP/MapServer',token: 'KhTlmtGizEANsimdOfkRQqwSVed4dhDk3Bkvt5Ixy2wJb8uG-beMmvJueoO2zcsBH_5DngxB8qRf-HN50tsddw..'},
        {name: '2019厦门矢量',url: 'http://222.76.242.138/arcgis/rest/services/CGCS_XMMAP/MapServer',token: 'KhTlmtGizEANsimdOfkRQnfvgB7xTmc6WlWX5U5bBDEk4z6tAjyR17bLHiCgmMntZQPmDiGlD-kKuOep53nxQQ..'},
      ],
      crs: null,
    }
  },
  mounted() {
    this.crs = this.initCRS()
    this.initLayers()
    this.eginMap.setCRS(this.crs)
  },
  methods: {
    
    initCRS() {
       const CRS_4490 = new L.Proj.CRS(
        "EPSG:4490",
        "+proj=longlat +ellps=GRS80 +no_defs",
        {
          resolutions: [
            0.703125,
            0.3515625,
            0.17578125,
            0.087890625,
            0.0439453125,
            0.02197265625,
            0.010986328125,
            0.0054931640625,
            0.00274658203125,
            0.001373291015625,
            6.866455078125e-4,
            3.4332275390625e-4,
            1.71661376953125e-4,
            8.58306884765625e-5,
            4.291534423828125e-5,
            2.1457672119140625e-5,
            1.0728836059570312e-5,
            5.364418029785156e-6,
            2.682209064925356e-6,
            1.3411045324626732e-6
          ],
          origin: [-180, 90]
        }
      );
      return CRS_4490
    },
    // 初始化图层
    initLayers() {
      let imageLayers = []
      let vetorLayers = []
      // let all = []
      // 影像
      this.imgLayers.forEach((item)=>{
        let layer = this.createLayers(item)
        layer.name = item.name
        imageLayers.push(layer)
      })
      // 矢量
      this.vetorLayers.forEach((item)=>{
        let layer = this.createLayers(item)
        layer.name = item.name
        vetorLayers.push(layer)
      })
      
      if(this.loadXmTdt){
        let tdtLayers =  imageLayers.concat(vetorLayers)
        // console.log(all, 'all')
        let all = this.basemapLayers.concat(tdtLayers)
        // this.basemapLayers.push.apply(all)
        this.basemapLayers = all
        this.eginMap.basemapLayers = all
        console.log(this.basemapLayers, 'basemapLayers')
      }
    },
    // 创建年度图层
    createLayers(item) {
      let tileMapLayer = Esri.dynamicMapLayer({
        url: item.url,
        maxZoom: 18,
        minZoom: 10,
        zoomOffset: 1,
        useCors: true,
        token: item.token
      });
      return tileMapLayer
    },
    //
    changeMap(id, key) {
      //
      let map = this.eginMap.map
      // 获取当前图层
      this.basemapLayers.forEach(group => {
        group._leaflet_id == id ? group.addTo(map) : map.removeLayer(group)
      })
      //
      this.activeIndex = key
    }
  }
}
</script>

<style scoped>
.egin-basemap-layer{
  position: absolute;
  bottom: 23px;
  height: 95px;
  background: #fff;
  opacity: 0.8;
  z-index: 999;
  right: 60px;
  display: flex;
}
.egin-basemap-layer img{
  height: 60px;
  width: 60px;
  padding: 5px;
}
.egin-basemap-layer-item {
  height: 95px;
  font-size: 12px;
  text-align: center;
  overflow: hidden;
  cursor: pointer;
  /* padding: 0px 5px; */
}
.egin-basemap-layer-item .name{
  font-size: 14px;
}
.active {
  border-bottom: 1px solid #666;
}
</style>
```

# EginLayerConfig.js 图层列表图层配置文件 (需要与 EginLegendConfig.js 共用)

``` 
getUserId() 获取当前登录用户
newDivIcon() 自定义Icon样式
getStyleConfig() 获取样式
ginLayerConfig.js 图层配置文件
```

## 图层配置文件 说明

``` javascript
// 水保措施
18: { // 图层ID(发布出来的图层ID)
    operate: { // 图形编辑时用到
        query: false, // 搜索
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
}
```

##  EginLegendConfig.js 配置说明

``` javascript
// 图层ID(发布出来的图层ID)
18: [
        // 每一项
        {
            layerId: 18, // 确保这个值与前面的key一致
            // 类型
            geometryType: 'Point',
            // 图例显示名称
            name: '',
            // 字段名
            field: 'ZT', // 如果无则填写 * 图例过滤使用
            // 字段值
            value: '正常', // 如果无则填写'' 图例过滤使用
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
```

# EginLayers 图层

> 参数说明

  >> props 参数说明

    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

        | showLegend  | 是否显示图例(暂无用通过勾选后就直接显示与隐藏) |

  >> data 参数说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | layers  |  图层列表数据(根据后台接口返回) |
      | defaultProps  | 树形结构 |
      | defaultProps.children  | 树形结构子项 |
      | defaultProps.label  | 树形结构显示名称 |
      | map | 地图实例 |
      | L | leaflet框架 |
      | allDataOfLayers | 图层数据 |

  >> 方法 说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | loadLayers  | 加载图层数据 |
      | initLayers  | 创建图层 并添加到地图 |
      | createFeatureLayers  | 加载图层 |
      | loadClusterFeature  | 加载聚合图层 |
      | loadWmsService  | 加载WMS图层 |
      | loadRestService  | 加载ArcgisMap图层 |
      | checkChange  | 图层勾选变更 |

  

> 组件代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 13:38:03
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 18:52:42
 * @ Description: 图层列表
 -->

<template>
    <div class="egin-layers" v-show="show">
        <tree :data="layers" show-checkbox ref="tree" node-key="id" @check-change="checkChange" :default-expanded-keys="[2, 3]" :default-checked-keys="[5]" :props="defaultProps">
        </tree>
    </div>
</template>

<script>
    import {
        Tree
    } from 'element-ui'
    import layers from '../data/layers.json'

    import EginLayer from '../../libs/EginLayer'
    import layerConfig from './EginLayerConfig'
    import {
        mapGetters
    } from 'vuex'

    export default {
        name: 'EginLayers',
        props: {
            eginMap: Object,
            show: {
                type: Boolean,
                default () {
                    return false
                }
            },
            //
            showLegend: {
                type: Boolean,
                default () {
                    return true
                }
            }
        },
        components: {
            Tree
        },
        computed: {
            ...mapGetters({
                getLayerCheckedIds: 'map/getLayerCheckedIds',
                graphicEditType: 'map/graphicEditType'
                // baseInfoComponent: 'business/getBaseInfoComponent',
                // detailInfoComponent: 'business/detailInfoComponent',
            })
        },
        data() {
            return {
                EginLayer: null,
                layers: layers,
                defaultProps: {
                    children: 'children',
                    label: 'name'
                },
                map: this.eginMap.map,
                L: this.eginMap.L,
                // 所有图层数据
                allDataOfLayers: []
            }
        },
        mounted() {
            // console.log("aaaa")
            this.EginLayer = new EginLayer({
                map: this.map
            })
            this.loadLayers(this.layers)
            this.initLayers()
        },
        methods: {
            // 添加图层到地图上
            loadLayers(layers) {
                // 图层数据处理 只获取有图层信息的数据
                layers.forEach(layer => {
                    // console.log(layer.pg.LAYER_TYPE,'element')
                    if (layer.pg.DIR_TYPE == 2) {
                        // console.log('layer.pg.DIR_TYPE',layer.pg.DIR_TYPE)
                        this.allDataOfLayers.push(layer)
                    } else {
                        // 有1与-1
                        this.loadLayers(layer.children)
                    }
                })
            },
            //
            // 创建图层 并添加到地图
            initLayers() {
                // console.log(this.allDataOfLayers,'allDataOfLayers')
                this.allDataOfLayers.forEach((item) => {
                    //
                    let pg = item.pg
                    // this.createFeatureLayers(pg)

                    let layer = this.createFeatureLayers(pg)
                    // console.log(layer, 'initLayers')
                    if (layer) {
                        // 暂时不需要添加到地图上(勾选的时候再添加)
                        // layer.addTo(this.map)
                        // 添加到图层数据里
                        this.eginMap.layers.push(layer)
                    }

                    // 切割测试用
                    if (layer.layerId == 19) {
                        layer.addTo(this.map)
                    }

                    return
                })
                // console.error(this.eginMap.layers,'layers')
            },
            // 加载图层
            createFeatureLayers(pg) {
                const layerUrl = pg.LAYER_URL
                // wms server测试链接 
                // const layerUrl = 'http://dev.eginsoft.cn:6080/arcgis/services/fireService/FireService/MapServer/WMSServer'
                const layerId = pg.LAYER_TABLE
                const canEdit = pg.EDIT_ATTRIBUTE == 1 ? true : false
                // const canSearch = pg
                // const layerType = pg.LAYER_TYPE
                let currentLayer = null
                let url = layerUrl + '/' + layerId

                let options = layerConfig(layerId, this.map)
                // 重置
                options.operate = {
                    query: canEdit, // 搜索
                    add: canEdit, // 新增
                    edit: canEdit, // 编辑
                    delete: canEdit, // 删除
                    cut: canEdit, // 切割
                    union: canEdit, // 合并
                    remove: canEdit, // 整形
                }
                const pane = this.map.createPane('pane_' + layerId)

                // 点击获取数据
                const onEachFeature = (feature, layer) => {

                    // 单击查看详情
                    layer.off('click').on('click', (evt) => {
                        //
                        this.eginMap.L.DomEvent.stopPropagation(evt)
                        //
                        if (this.graphicEditType == 'edit') {
                            // 如果事件是双击则禁用双击缩放编辑 图形
                            // this.map.doubleClickZoom.disable();
                            this.$store.dispatch('map/setEditLayer', layer)

                        } else if (this.graphicEditType == 'cut' || this.graphicEditType == 'remove') {
                            this.$store.dispatch('map/setEditLayer', layer)
                        }
                        // 如果处于编辑状态则直接返回
                        if (this.graphicEditType == '') {
                            this.eginMap.L.DomEvent.stopPropagation(evt)
                            // console.log(feature, 'feature',evt)
                            this.$store.dispatch('business/setBaseInfoComponent', {
                                layerId: layerId,
                                feature: feature,
                                evt: evt
                            })
                        }

                    })
                }

                // 这里不能写死需要后台传参
                if (layerId == 18 || layerId == 44) {
                    currentLayer = this.loadClusterFeature(url, options.pointToLayer, onEachFeature, options.filter, options.where, options.operate, layerId, pane)
                } else {
                    // loadRestService
                    currentLayer = this.loadRestService(url, options.style, options.pointToLayer, onEachFeature, options.filter, options.where, options.operate, layerId, pane)
                }
                // console.log(currentLayer.getPane(), 'currentLayer')
                // 添加属性
                currentLayer.layerName = pg.LAYER_NAME
                currentLayer.layerId = layerId

                return currentLayer
            },
            // 加载
            loadClusterFeature(url, pointToLayer, onEachFeature, filter, where, operate, layerId, pane) {
                //
                const options = {
                    url: url,
                    // spiderfyOnMaxZoom: true, // 最大层级后不显示扩展线
                    // removeOutsideVisibleBounds: true,
                    onEachFeature: onEachFeature,
                    pointToLayer: pointToLayer,
                    where: where,
                    filter: filter,
                    operate: operate,
                    layerId: layerId,
                    pane: pane,
                    disableClusteringAtZoom: 16, // 16级后取消聚合
                }
                const layer = this.EginLayer.addClusterLayer(options)
                return layer
            },
            // 加载wms服务
            loadWmsService(url, layerId) {
                let layer = this.L.tileLayer.wms(url, {
                    layers: [layerId], // 0,1,2,3,4,5,6,7
                    format: 'image/png',
                    transparent: true,
                    attribution: ""
                })
                return layer
            },
            // 服务
            loadRestService(url, style, pointToLayer, onEachFeature, filter, where, operate, layerId, pane) {
                //

                let options = {
                    url: url,
                    // 颜色配置
                    style: style,
                    where: where, //"MAP_DIVISION_ID='350213'",
                    filter: filter,
                    operate: operate,
                    layerId: layerId,
                    pointToLayer: pointToLayer,
                    onEachFeature: onEachFeature,
                    pane: pane,
                }
                // console.log(style,pointToLayer,onEachFeature,filter,where)
                let layer = this.EginLayer.addEsriFeatureLayer(options)

                return layer
            },

            // loadGeoJson 某个
            loadGeoJsonLayer(geojson) {
                //

                const style = function() {
                    return {

                    }
                }
                // 点样式
                const pointToLayer = function() {
                    // 
                }
                // 循环每个图形
                const onEachFeature = function() {

                }
                // 条件过滤
                const filter = function() {
                    // 
                }
                //
                const options = {
                    filter: filter,
                    onEachFeature: onEachFeature,
                    pointToLayer: pointToLayer,
                    style: style
                }
                this.EginLayer.addGeoJsonLayer(geojson, options)
            },
            // 选择变更
            checkChange: function(item, isChecked) {
                //
                // 获取到子节点的数据 且 包括当前子节点的父结点
                let layerId = item.pg.LAYER_TABLE
                // 
                let layerCheckedIds = this.getLayerCheckedIds
                // 判断item是否存在
                this.eginMap.layers.forEach((item) => {
                    // 判断地图中是否已添加
                    // const hasLayer =  this.map.hasLayer(item)
                    let removeIndex = layerCheckedIds.indexOf(item)
                    // console.log('removeIndex', removeIndex)
                    if (layerId == item.layerId) {
                        isChecked ? layerCheckedIds.push(item) : layerCheckedIds.splice(removeIndex, 1) // console.log("删除",  ) //layerCheckedIds.splice(index,1)
                        isChecked ? item.addTo(this.map) : this.map.removeLayer(item)
                    }
                })
                console.log(layerCheckedIds, 'layerCheckIds')
                // 修改值
                this.$store.dispatch('map/setLayerCheckedIds', layerCheckedIds)
            }
        },
    }
</script>

<style scoped>
    .egin-layers {
        z-index: 999;
        background: #fff;
        opacity: 0.9;
        top: 12px;
        position: absolute;
        right: 50px;
        /*测试用*/
        /* width: 300px;  */
        /* height: 300px; */
        border-radius: 5px;
        overflow: hidden;
        min-width: 220px;

    }
</style>
```

# EginLegend 图例

> 参数说明

  >> props 参数说明

    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

  >> data 参数说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | maxLevel  |  最大层级 |
      | legendConfig  | 图例配置文件 |
      | layerOpcity  | 图层透明度,默认为非透明 |
      | map | 地图实例 |
      | L | leaflet框架 |
      | checkedLayers | 已选中的图层 |
      | maxZIndex  |  最大z-index |
      | minZIndex  | 最小z-index |
      | LayerIndex  | 图层z-index |
      | layerWhere | 图层条件 |
      | legendData | 图例数据 |
      | clickLayerId | 当前点击的是哪个图层 |
      | showSlider | 滑块 |
      | checkedIndex | 切换顺序时存储点击勾选时父图层 |

  >> watch 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | layerOpcity  | 图层透明度,默认为非透明 |
      | getLayerCheckedIds  | 已勾选图层id |
      | layerWhere | 条件筛选 |
      | clickLayerId | 点击的layerId |
     

  >> computed 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | getLayerCheckedIds  | 已勾选图层id |
      | showEginLegend | 是否显示图例 |

     
  >> 方法 说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | _getPane  | 获取Pane |
      | _setPaneIndex  | 设置pane#index |
      | changeIndex  | 修改层级 |
      | getChangeIndex  | 获取图层位置 |
      | changeArrIndex  | 修改数组顺序 |
      | getStyle  | 获取样式 |
      | checkChange  | 树形勾选事件 |
      | getTreeItemStyle  | 获取树形item样式 |


> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-23 09:13:53
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-22 16:36:58
 * @ Description: 图例工具
 -->

<template>
  <div class="egin-legend" v-show="showEginLegend">
    <div class="layers" >
      <div class="layer-bar">
        <div class="layer-opcity" style="padding-right:10px;">
          <span class="title">透明度:</span>
          <!-- <input-number v-model="layerOpcity" :max="1" :min="0" :step="0.01" size="small" /> -->
          <div class="opcity" @click="showSlider = !showSlider">
            {{ parseInt(layerOpcity * 100) }}%
          </div>
        </div>

         <div>
          层级:
          <el-tooltip content="至顶" placement="top">
            <i
              class="iconfont icondajiantou-shang"
              @click="changeIndex(LayerIndex.Top)"
            ></i>
          </el-tooltip>
          <el-tooltip content="上一级" placement="top">
            <i
              class="iconfont iconjiantou-copy-copy"
              @click="changeIndex(LayerIndex.Up)"
            ></i>
          </el-tooltip>
           <el-tooltip content="下一级" placement="top">
          <i
            class="iconfont iconjiantou"
            @click="changeIndex(LayerIndex.Down)"
          ></i>
          </el-tooltip>
          <el-tooltip content="至底" placement="top">
          <i
            class="iconfont icondajiantou-xia"
            @click="changeIndex(LayerIndex.Bottom)"
          ></i>
          </el-tooltip>
        </div>
      </div>

      <!-- 滚动条 -->
      <transition name="fade">
        <slider
          v-show="showSlider"
          size="mini"
          style="width:100%;"
          v-model="layerOpcity"
          :max="1"
          :min="0"
          :step="0.01"
          :show-tooltip="false"
        ></slider>
      </transition>

      <div class="egin-legend-layer">
        <tree
          :data="legendData"
          show-checkbox
          node-key="id"
          ref="legendTree"
          default-expand-all
          @check="checkChange"
          :expand-on-click-node="false"
        >
          <div class="custom-tree-node" slot-scope="{ node, data }">
            <div>{{ node.label }}</div>
            <div class="custom-html" v-html="getTreeItemStyle(data)"></div>
          </div>
        </tree>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import legendConfig from './EginLegendConfig'
import LayerIndex from '../../libs/const/LayerIndex'
import { Tree, Slider, Tooltip } from 'element-ui'


export default {
  name: 'EginLegend',
  props: {
    show: {
      type: Boolean,
      default() {
        return false
      }
    },
    eginMap: {}
  },
  components: {
    Tree,
    Slider,
    ElTooltip: Tooltip
    // InputNumber
  },
  watch: {
    // 修改透明度(已勾选的透明度)
    layerOpcity: function(newVal) {
      //
      this.checkedLayers.map(layer => {
        // 获取当前图层所处pane
        let pane = this._getPane(layer.layerId)
        // 设置透明度
        this.L.DomUtil.setOpacity(pane, newVal)
      })
    },
    // 变化需要修改data
    getLayerCheckedIds: function(checkedLayersIds) {
      console.error('checkedLayersIds',checkedLayersIds)
      this.legendData = [] // 清空
      let keys = []
      //
      checkedLayersIds.map(layer => {
        const layerId = layer.layerId
        // 跳过标绘图层
        // if(['42','43','44','45'].indexOf(layerId) == -1) {
          keys.push(layerId)
          let tmp = {
            id: layerId, // 做key
            layerId: layerId,
            label: layer.layerName,
            layer: layer,
            children: legendConfig[layerId]
          }
          this.legendData.push(tmp)
        // }
       
      })
      // 默认选中当前状态
      this.$refs.legendTree.setCheckedKeys(keys)
      // 赋值已选中
      this.checkedLayers = checkedLayersIds
    },
    // 条件筛选
    layerWhere: {
      handler: function(newVal) {
        // console.log(newVal, 'layerWhere@@@@@@handler')
        let currentLayer = null

        // console.log('Watch', newVal)
        // 如果checkedLayers.length > 0
        this.getLayerCheckedIds.map(item => {
          if (item.layerId == this.clickLayerId) {
            currentLayer = item
          }
        })
        console.log(currentLayer, 'layer')
        // 如果图层不存在则返回
        if (!currentLayer) return

        // 当前点击图层的数据
        let conditions = newVal[this.clickLayerId]
        console.log(conditions, 'conditions')
        // 无条件即未选中
        if (typeof conditions == 'undefined' || conditions.length == 0) {
          return this.map.removeLayer(currentLayer)
        }
        // 勾选条件
        let checkWhere = ''
        // 需要保存当前
        let where = currentLayer.options.where
        //
        let index = where.indexOf(' and ')
        // 原始条件
        const orignWhere = index == -1 ? where : where.substring(0, index)

        // console.log(conditions,'condition')
        // 无条件即未选中
        // 已有的数据
        conditions.forEach(item => {
          // 判断条件是否为* 如果为* 则说明是无条件直接显示当前图层
          if (item.field != '*') {
            //
            checkWhere += `${item.field} = '${item.value}' or `
          } else {
            // 如果图层不存在则需要添加
            currentLayer.addTo(this.map)
            return
          }
        })
        // 删除最后一个条件 or
        let lastOrIndex = checkWhere.lastIndexOf(' or ')
        checkWhere = checkWhere.substring(0, lastOrIndex)
        // 连接条件
        const tmpWhere =
          checkWhere == ''
            ? orignWhere
            : orignWhere + ' and (' + checkWhere + ')'
        // console.log(tmpWhere,'tmpWhere')
        // 添加条件(对当前图层做筛选)
        currentLayer.options.where = tmpWhere
        // 先删除后面再添加(相当于刷新图层)
        this.map.removeLayer(currentLayer)
        // 如果图层不存在则需要添加
        currentLayer.addTo(this.map)
      },
      deep: true,
      immediate: true
    },
    clickLayerId: function(newVal, oldVal) {
      if (newVal || newVal == oldVal) {
        console.log('newVal', oldVal, newVal)
      }
    }
  },
  data() {
    return {
      // 图层最大index
      maxLevel: 0,
      legendConfig: legendConfig,
      layerOpcity: 1, // 默认为非透明
      map: null,
      L: null,
      // 已选中的图层
      checkedLayers: [],
      maxZIndex: 400,
      minZIndex: 400,
      LayerIndex: LayerIndex,
      // 图层条件
      layerWhere: [],
      legendData: [],
      // 当前点击的是哪个图层
      clickLayerId: null,
      // 滑块
      showSlider: false,
      // 切换顺序时存储点击勾选时父图层
      checkedIndex: -1
    }
  },
  computed: {
    ...mapGetters({
      getLayerCheckedIds: 'map/getLayerCheckedIds'
    }),
    // 是否显示
    showEginLegend() {
      return this.getLayerCheckedIds.length > 0
    }
  },
  mounted() {
    this.map = this.eginMap.map
    this.L = this.eginMap.L
    // 重置已选项
    this.checkedLayers = []
    // 重置条件
    this.layerWhere = []
    // console.log(this.legendConfig,'legendConfig')
  },
  methods: {
    // 获取Pane
    _getPane(layerId) {
      return this.map.getPane('pane_' + layerId)
    },
    // 设置pane#index
    _setPaneIndex(type) {
      // 获取所有图层的最大值 最小值
      // 如果是上一级则++
      // 如果是下一级则--
      // 如果是置顶则 max++
      // 如果是置底则 min--
      // 设置zindex
      // console.log(type,'type')
      this.checkedLayers.map(layer => {
        let pane = this._getPane(layer.layerId)
        let currentIndex = pane.style['z-index'] ? pane.style['z-index'] : 0
        switch (type) {
          case LayerIndex.Up:
            currentIndex++
            break
          case LayerIndex.Top:
            currentIndex = ++this.maxZIndex
            break
          case LayerIndex.Down:
            currentIndex--
            break
          case LayerIndex.Bottom:
            currentIndex = --this.minZIndex
            break
        }
        // console.log(pane, currentIndex, type, 'currentIndex')
        pane.style['z-index'] = currentIndex
      })
      let arr  = this.legendData
      // 修改图例列表顺序
      this.changeArrIndex(arr,type)
    },
    // 修改层级
    changeIndex(type) {
      this._setPaneIndex(type)
    },
    // 获取图层
    getChangeIndex() {
      this.checkedIndex = -1;
      this.legendData.forEach((item,key)=>{
        if(item.layerId == this.clickLayerId) {
          this.checkedIndex = key
        }
      })
      // this.changeIndex(type)
      console.log(this.checkedIndex, 'getChangeIndex#checkedIndex')
    },
    // 修改数组顺序
    changeArrIndex(arr,type) {
      //
      this.getChangeIndex()
      console.log(arr,'before')
      let key = this.checkedIndex

      console.error(key,'key',arr,'arr')
      if(type == LayerIndex.Top) {
        if (key == -1 || key==0) return;
        const delItems = arr.splice(key, 1)
        // 删除后的元素放到头部
        arr.unshift(...delItems)
        key = 0
      }else if(type == LayerIndex.Up) {
        if (key == -1 || key==0) return;
        // 拼接函数(索引位置, 要删除元素的数量, 元素)
        const delItems =  arr.splice(key, 1);
        arr.splice(key-1, 0, ...delItems )
        --key

      }else if(type == LayerIndex.Down){
        if (key == -1 || key == arr.length-1) return;
        const delItems =  arr.splice(key, 1);
        arr.splice(key+1, 0, ...delItems )
        ++key
      }else if(type == LayerIndex.Bottom){
        if (key == -1 || key == arr.length-1) return;
        let delItems = arr.splice(key, 1)
        // 删除后的元素放到头部
        arr.push(...delItems)
        key = arr.length -1
      }
      this.checkedIndex = key
      console.log(arr,'after')
      // 设置有效
      // 设置有效  // 不能使用 this.legendData = arr 无效还有 也不能用this.$set(this,'legendData',arr) 这个也无效
      this.legendData = [...arr]
    },
    // 获取样式
    getStyle(config) {
      return `margin-top:3px; border:2px solid ${config.style.value.color};background:${config.style.value.fillColor};opacity:${config.style.value.opacity}`
    },
    // 树形勾选事件
    checkChange(checkedNodes) {
      // 最近点击的图层编号
      if (typeof checkedNodes.layerId == 'undefined') {
        console.error('图层ID不能为空')
        return
      }
      const layerId = checkedNodes.layerId
      //
      this.clickLayerId = layerId
      // console.log(layerId, 'clickId')
      // 重置当前点击图层的条件 如果有数据则清除 如果没有则创建
      if (this.layerWhere[layerId] || !this.layerWhere[layerId]) {
        this.layerWhere[layerId] = []
      }
      // 设置透明度
      let pane = this._getPane(layerId)
      let opacity = pane.style['opacity'] ? pane.style['opacity'] : 1
      opacity = parseFloat(opacity)
      this.layerOpcity = opacity
      // 如果有 halfCheckedNodes 不为空 则这个是父级
      // checkedNodes 有可能是父级也有可能是子集
      let allCheckedNodes = this.$refs.legendTree.getCheckedNodes()
      let allHalfCheckedNodes = this.$refs.legendTree.getHalfCheckedNodes()
      // 已选中的图层
      let legendLayerChecked = []

      // 半选中状态 说明肯定是图层
      allHalfCheckedNodes.map(item => {
        legendLayerChecked.push(item.layer)
      })
      let tmpWhere = []
      // 全选中状态(不一定都是图层)
      allCheckedNodes.map(item => {
        // console.log(item,'item')
        // 图层
        if (typeof item.id != 'undefined') {
          legendLayerChecked.push(item.layer)
        } else {
          // 如果图层不是当前点击的图层则继续
          if (item.layerId != layerId) {
            console.log('return')
            return
          }
          // 条件
          const field = item.field
          // 判断当前条件下是否存在
          const value = item.value
          let tmp = {
            field: field,
            value: value
          }
          // 如果不存在且已勾选
          tmpWhere.push(tmp)
          // this.layerWhere[layerId].push(tmp)
        }
      })
      // this.layerWhere[layerId] = tmpWhere (无法监听到)
      this.$set(this.layerWhere, layerId, tmpWhere) // (可以监听到)
      this.checkedLayers = legendLayerChecked
      // console.log(this.layerWhere,'layerWhere')
    },
    // 获取树形item样式
    getTreeItemStyle(data) {
      // 如果这个存在
      let html = ''
      if (data.style) {
        console.log(data, 'data')
        switch (data.style.type) {
          case 'icon':
            html = `<div class="item-style ${data.style.value.icon.options.className}" ></div>`
            break
          case 'style':
            html = `<div class="item-style" style="${this.getStyle(
              data
            )}"></div>`
            break
          case 'img':
            html = `<div class="item-style"><img src="${data.style.value.icon.options.iconUrl}" width="${data.style.value.icon.options.iconSize[0]}" height="${data.style.value.icon.options.iconSize[0]}"/></div>`
            break
        }
      }
      html =
        typeof data.name == 'undefined' || data.name == ''
          ? html
          : html + `<div class="item-name">${data.name}</div>`
      return html
    }
  }
}
</script>

<style lang="scss" scoped>
.egin-legend {
  overflow: hidden;
  border-radius: 4px;
  // display: flex;
  // flex-direction: column;

  .layers {
    padding: 10px;
    .egin-legend-layer {
      max-height: 300px;
      overflow-y: scroll;
    }
    .layer-bar {
      padding-bottom: 6px;
      border-bottom: 1px solid rgb(190, 190, 190);

      .layer-opcity {
        display: flex;

        .opcity {
          width: 100px;
          border-radius: 2px;
          background-color: #bdbdbd;
        }
      }
    }
  }
  .custom-html {
    display: flex;
  }
  .top-header {
    color: #ffffff;
    // @include top-header-gradients;
    padding: 4px 0px;
  }
}
</style>
```
# EginPosition 显示鼠标当前位置坐标组件

> 参数说明

  >> props 参数说明

    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  >> data 参数说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | point  |  当前位置 |
      | point.x  | 经度 |
      | point.y  | 纬度 |

> 组件代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-29 09:04:58
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 17:49:43
 * @ Description: 显示当前坐标组件
 -->

<template>
    <div class="egin-position">
        坐标:({{point.x}}, {{point.y}})
    </div>
</template>

<script>
    export default {
        name: 'EginPosition',
        props: {
            eginMap: {
                type: Object,
                require: true
            }
        },
        data() {
            return {
                point: {
                    x: 0,
                    y: 0
                }
            }
        },
        mounted() {
            // this.map = this.eginMap.map
            // console.log(this.map,'map')
            this.eginMap.map.on('mousemove', (evt) => {
                let latlng = evt.latlng
                this.point = {
                    x: latlng.lng,
                    y: latlng.lat
                }
            })
        }
    }
</script>

<style scoped>
    .egin-position {
        left: 50px;
        bottom: 10px;
        z-index: 9999;
        font-size: 14px;
        color: #fff;
        padding: 5px;
        position: absolute;
    }
</style>
```

# EginStatic 统计组件(根据业务需要来处理)

> 组件代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 14:25:23
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:23:15
 * @ Description: 统计组件
 -->

<template>
    <div class="egin-static" v-show="show">
        <div>Static</div>
    </div>
</template>

<script>
    export default {
        name: 'EginStatic',
        props: {
            show: {
                type: Boolean,
                default () {
                    return false
                }
            }
        }
    }
</script>

<style scoped>
    .egin-static {
        width: 300px;
        position: absolute;
        /* left: 0; */
        height: 100%;
        background: #fff;
        z-index: 999;
        top: 0px;
    }
</style>
```

# EginTools 工具组件

> 参数说明

  >> props 参数说明

    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

        | show  | 是否显示组件 |

  >> compoments 说明

    >>>  
        |  名称   |  说明  |
        |  ----  | ----  |
        | Compare  | 对比组件 |
        | GraphicEdit  | 图形编辑组件 |
        | Location  | 定位组件 |
        | Mark  | 标绘组件 |
        | Measure  | 测量组件 |
        | Print  | 打印组件 |

  >> data 参数说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | currentTool  |  当前组件 |
      | tools  | 工具 |
      | tools.name  | 工具项名称 |
      | tools.icon | 工具项icon |
      | tools.type | 工具项类型(组件名称) |
    

  >> 方法 说明

    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | changeTools  | 工具切换 |
      | clearActivity  | 取消选中状态 |

> 组件代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 14:18:01
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:44:27
 * @ Description: 工具栏组件
 -->

<template>
    <div class="egin-tools-wrap" v-show="show">
        <div class="egin-tools">
            <div class="egin-tools-item" :class="{'activity': currentTool == item.type }" v-for="(item,key) in tools" :key="key" @click="changeTools(item)">
                <!--树形结构-->
                <div class="egin-tools-menu"><i class="egin-icon iconfont" :class="item.icon"></i>{{item.name}}</div>
            </div>
        </div>
        <!--动态组件(工具)-->
        <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap"></component>
        <location ref="location" :eginMap="eginMap"></location>
    </div>
</template>

<script>
    export default {
        name: 'EginLayers',
        props: {
            eginMap: Object,
            show: {
                type: Boolean,
                default () {
                    return false
                }
            }
        },
        components: {
            'Compare': () => import('./tools/Compare'),
            'GraphicSearch': () => import('./tools/GraphicSearch'),
            'GraphicEdit': () => import('./tools/GraphicEdit'),
            'Location': () => import('./tools/Location'),
            'Mark': () => import('./tools/Mark'),
            // 'MarkList': ()=> import('./tools/MarkList'),
            'Measure': () => import('./tools/Measure'),
            'Print': () => import('./tools/Print'),
        },
        data() {
            return {
                currentTool: '',
                tools: [{
                        name: '测量',
                        icon: 'iconRuler',
                        type: 'Measure'
                    },
                    {
                        name: '标绘',
                        icon: 'iconPenandruler_px',
                        type: 'Mark'
                    },
                    {
                        name: '对比',
                        icon: 'iconcomparearrows',
                        type: 'Compare'
                    },
                    {
                        name: '图形编辑',
                        icon: 'iconbianjishuru',
                        type: 'graphic-edit'
                    },
                    {
                        name: '图形搜索',
                        icon: 'iconchangehistory',
                        type: 'graphic-search'
                    },
                    {
                        name: '定位',
                        icon: 'iconsoph_location',
                        type: 'Location'
                    },
                    {
                        name: '打印',
                        icon: 'iconprint',
                        type: 'Print'
                    },
                    {
                        name: '下载',
                        icon: 'icondownload',
                        type: 'Download'
                    },
                ]
            }
        },
        mounted() {
            // 导出时需要用到
            this.eginMap.loadPrintWidget()
        },
        methods: {
            changeTools(item) {
                //
                if (this.$refs.tools) {
                    this.$refs.tools.showTools = !this.$refs.tools.showTools
                }
                this.currentTool = item.type
                // console.log(item)
                if (item.type == 'Print') {
                    // this.clearActivity()
                    this.eginMap.printMap()
                } else if (item.type == 'Download') {
                    // this.clearActivity()
                    this.eginMap.exportMap('map')
                } else if (item.type == 'Location') {
                    // this.clearActivity()
                    // console.log('location')
                    this.$refs.location.dialogVisible = true
                }

            },
            clearActivity() {
                this.currentTool = ''
            }
        }
    }
</script>

<style scoped>
    .egin-tools-wrap {
        display: flex;
        justify-content: center;
    }

    .egin-tools {
        z-index: 999;
        background: #fff;
        top: 12px;
        position: absolute;
        right: 50px;
        width: 150px;
        height: auto;
    }

    .egin-tools .egin-tools-item {
        padding: 10px;
        font-size: 12px;
        height: 40px;
        border-bottom: 1px solid #ccc;
        box-sizing: border-box;
        cursor: pointer;
    }

    .egin-tools .egin-tools-item .egin-tools-menu {
        height: 20px;
        line-height: 20px;
        display: flex;
        vertical-align: middle;
    }

    .egin-tools .egin-tools-item i {
        padding-right: 10px;
        font-size: 18px;
    }

    .tool {
        position: absolute;
        top: 10px;
        width: auto;
        z-index: 999;
        border-radius: 5px;
        background: rgba(255, 255, 255, 0.8);
    }

    .activity {
        background: #eeeeee;
    }

    .leaflet-control-easyPrint-button-export {
        display: none;
    }
</style>
```

# 调用事例代码

``` html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-16 13:38:06
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-22 10:49:15
 * @ Description: 地图测试组件
 -->

<template>
  <div>
    <div class="map-container" ref="mapContainer"></div>
    <EginBasemap v-if="eginMap" :eginMap="eginMap" :show="showEginBasemap" :loadXmTdt="true"/>
    <EginLayers v-if="eginMap" :eginMap="eginMap" :show="showEginLayers" />

    <EginTools :show="showEginTools" :eginMap="eginMap" style="top:60px;" />
    <EginStatic :show="showEginStatic" />
    <EginLegend :show="showEginLegend" v-if="eginMap" :eginMap="eginMap" />
    <!--当前原型上没有 暂时不显示-->
    <!-- <EginPosition v-if="eginMap" :eginMap="eginMap" /> -->
    <!--基本信息动态组件-->
    <component :is="baseInfo" ref="baseInfo"></component>
    <!--详情组件-->
    <component
      :is="detailInfo.componentsName"
      :data="detailInfo.data"
      ref="detailInfo"
    ></component>

    <!--切割弹窗列表-->
    <component
      ref="graphicDialog"
      :is="graphicEditComponent"
      :data="graphicEditData"
      :color-list="graphicEditColorList"
      field="OBJECTID"
    ></component>
    <!--标绘组件弹窗-->
    <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="markerConfirm"
      @remove="markerRemove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component>
    <!--标绘列表弹窗(暂时不显示)-->
    <!-- <mark-list
      :list="markerDataLists"
      @remove="markerListRemove"
      @location="markerLocation"
    ></mark-list> -->
    
    <!--提示-->
    <el-alert
      class="alert"
      :title="alertInfo.title"
      :type="alertInfo.type"
      center
      show-icon
      v-show="showAlert"
      @close="close"
    ></el-alert>
  </div>
</template>

<script>
// import EginMap from '../../libs//EginMap.js'
import EginMap from '../libs/EginMap'
import TiandiTu from '../libs/basemap/TiandiTu'
import EginPosition from '../libs/const/EginPosition'
import { Alert } from 'element-ui'
import ComponentNameMap from '../libs/const/ComponentName'
import { mapGetters } from 'vuex'
import GeometryType from '../libs/const/GeometryType'
// import { mapApi } from '@api/mapApi/map'

export default {
  name: 'EginMap',
  components: {
    ElAlert: Alert,
    EginBasemap: () => import('../libs/components/EginBasemap'),
    EginLayers: () => import('../libs/components/EginLayers'),
    EginTools: () => import('../libs/components/EginTools'),
    EginStatic: () => import('../libs/components/EginStatic'),
    EginLegend: () => import('../libs/components/EginLegend'),
    EginLineInfo: () => import('../libs/components/info/EginLineInfo'),
    EginPointInfo: () => import('../libs/components/info/EginPointInfo'),
    EginPolygonInfo: () => import('../libs/components/info/EginPolygonInfo'),
    Detail: () => import('../libs/components/info/Detail'),
    Detail1: () => import('../libs/components/info/Detail1'),
    Detail2: () => import('../libs/components/info/Detail2'),
    // Detail2: () => import('../libs/components/info/ProductDetail'),
    EginPosition: () => import('../libs/components/EginPosition'),
    UnionList: () => import('../libs/components/tools/graphicEdit/UnionList'),
    CutList: () => import('../libs/components/tools/graphicEdit/CutList'),
    // 标绘弹窗
    PointMarker: () => import('../libs/components/tools/marker/PointMarker'),
    PolylineMarker: () => import('../libs/components/tools/marker/PolylineMarker'),
    PolygonMarker: () => import('../libs/components/tools/marker/PolygonMarker'),
    TextMarker: () => import('../libs/components/tools/marker/TextMarker'),
    MarkList: () => import('../libs/components/tools/MarkList'),
    //
    
  },
  mounted() {
    this.loadMap()
  },
  computed: {
    ...mapGetters({
      baseInfoComponent: 'business/getBaseInfoComponent',
      detailInfo: 'business/getDetailInfoComponent',
      alertInfo: 'alert/alertInfo',
      getSelectDialog: 'map/getSelectDialog',
      markTool: 'map/markTool'
    })
  },
  watch: {
    baseInfoComponent(newVal) {
      console.log(newVal, 'newVal')
      let layerId = newVal.layerId
      this.baseInfo = ComponentNameMap[layerId] ? ComponentNameMap[layerId] : ''
      if (this.baseInfo != '') {
        // 数据请求
        // mapApi.getPortList(params).then(res => {
        //   console.log(res, 'res')
        // })

        setTimeout(() => {
          this.$refs.baseInfo.dialogVisible = true
        }, 300)
      }
      // console.log(newVal ,'baseInfoComponent',this.baseInfo)
    },
    detailInfo(newVal) {
      if (newVal != '') {
        setTimeout(() => {
          this.$refs.detailInfo.dialogVisible = true
        }, 300)
      }
    },
    //
    alertInfo: {
      handler: function(newVal) {
        console.log(newVal, 'alertInfo')
        this.showAlert = true
        setTimeout(() => {
          this.showAlert = false
        }, 2000)
      },
      deep: true
    },
    // 这里不能用deep
    getSelectDialog: function(newVal) {
      this.graphicEditComponent = newVal.componentName
      this.graphicEditData = newVal.dataList
      this.graphicEditColorList = newVal.colorList
      console.log('dataList###', newVal.dataList, newVal.colorList)
    },
    // markTool: function(newVal){
    //   console.log(newVal, 'newVal')
    // }
    markTool: {
      handler: function(newVal){
        console.error('newVal', newVal)
        //
        setTimeout(() => {
          // 
         this.$refs.markerDialog.dialogVisible = true 
        }, 300)
      },
      immediate: true,
      deep: true
    }
  },
  data() {
    return {
      // EginMap
      eginMap: null,
      showEginBasemap: false,
      showEginLayers: false,
      showTools: false,
      showEginTools: false,
      showEginStatic: false,
      showEginLegend: false,
      // 是否显示提示信息
      showAlert: false,
      // 默认为 EginLineInfo
      baseInfo: '',
      // detailInfo: '', // 详情弹窗组件
      graphicEditComponent: '',
      graphicEditData: [],
      graphicEditShowField: 'ID',
      graphicEditColorList: [],
      // 标绘数据
      markerDataLists: [],
      // 存放标绘的图层
      markerFeatureGroup: null,
      //
      pointLayerId: 'e6e25e52f6384d52a6e941cfe7ad551d', // 标绘点
      //
      poylineLayerId: 'b197aba58a934a8ba576632f44938f54', // 标绘线
      //
      polygonLayerId: '7fc9aff62ea640bc82de56ebd4d458b0', // 标绘面
      //
      textLayerId: 'b7c72b2ed74c448ba1664bcd22dc8878 ', // 文字标绘
    }
  },
  methods: {
    close() {
      this.showAlert = false
    },
    // 加载底图
    loadBasemap() {
      // 添加底图
      const tiandiTu = new TiandiTu({ key: '4b350b4f343fa22cdb2047e93b4d8712' })
      // 卫星图
      let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
      //  卫星注记图
      let tdtSatelliteAnnotion = tiandiTu.addTianDiTuSatelliteAnnotion()
      // 分组
      let tdtSatelliteGroup = this.eginMap.L.layerGroup([
        tdtSatellite,
        tdtSatelliteAnnotion
      ])
      // 添加 title 与 url 属性 为切换底图做准备
      tdtSatelliteGroup.name = '卫星图'
      tdtSatelliteGroup.url = ''
      // 添加到地图
      this.eginMap.map.addLayer(tdtSatelliteGroup)
      // 矢量图
      let tdtNormal = tiandiTu.addTianDiTuNormalMap()
      //  矢量注记图
      let tdttdtNormalAnnotion = tiandiTu.addTianDiTuNormalAnnotion()
      // 分组
      let tdtNormalGroup = this.eginMap.L.layerGroup([
        tdtNormal,
        tdttdtNormalAnnotion
      ])
      tdtNormalGroup.name = '矢量图'
      tdtNormalGroup.url = ''

      // 添加到底图里
      this.eginMap.basemapLayers.push(tdtSatelliteGroup)
      // 添加到底图里
      this.eginMap.basemapLayers.push(tdtNormalGroup)

      // let baseMap = {
      //   "卫星图": tdtSatelliteGroup,
      //   "矢量图": tdtNormalGroup
      // }
      // this.eginMap.L.control.layers(baseMap).addTo(this.eginMap.map)
    },
    // 加载缩放控件
    loadZoomWidget() {
      //
      let zoomPosition = {
        zoomInTitle: '放大',
        zoomOutTitle: '缩小',
        position: EginPosition.BottomRight
      }
      // 设置zoom控件的位置
      this.eginMap.setZoomPosition(zoomPosition)
    },
    // 加载图层列表
    loadLayersWidget() {
      const layersOptions = {
        iconfont: 'iconic_layers_px',
        position: EginPosition.TopRight
      }
      // 加载图层
      this.eginMap.loadLayersWidget(layersOptions, evt => {
        console.log(evt, 'loadLayersWidget')
        this.showEginLayers = !this.showEginLayers
      })
    },
    // 加载底图
    loadBasempWidget() {
      // 加载底图控件
      const basemapWidgetOptions = {
        iconfont: 'iconglobe',
        position: EginPosition.BottomRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadBasempWidget(basemapWidgetOptions, () => {
        // console.log(evt, 'loadBasempWidget')
        this.showEginBasemap = !this.showEginBasemap
        // console.log(this.showEginBasemap, 'showBasemapLayer')
      })
    },
    // 加载工具条
    loadToolsWidget() {
      // 加载底图控件
      const toolsWidgetOptions = {
        iconfont: 'icongongjuxiang',
        position: EginPosition.TopRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadToolsWidget(toolsWidgetOptions, evt => {
        console.log(evt, 'loadToolsWidget')
        this.showEginTools = !this.showEginTools
        console.log(this.showEginTools, 'showEginTools')
      })
    },
    // 加载统计条
    loadStaticWidget() {
      const loadWidgetOptions = {
        iconfont: 'iconstar',
        position: EginPosition.TopRight
      }
      // 加载底图并添加点击事件
      this.eginMap.loadStaticWidget(loadWidgetOptions, (evt) => {
        console.log(evt, 'loadStaticWidget')
        this.showEginStatic = !this.showEginStatic
        // console.log(this.showEginStatic, 'showEginStatic')
      })
    },
    // 加载所有
    loadWidgets() {
      // 加载控件
      this.loadZoomWidget()
      this.loadLayersWidget()
      this.loadBasempWidget()
      this.loadToolsWidget()
      this.loadStaticWidget()
      this.eginMap.loadPrintWidget()
      // 比例尺 
      // 原型上暂没有不显示
      // this.eginMap.loadScaleWidget()
    },
    // 加载地图
    loadMap() {
      // 加载地图
      const eginMap = new EginMap({ container: this.$refs.mapContainer })
      this.eginMap = eginMap
      eginMap.setView([24.48, 118.108]).setZoom(11)
      // 加载底图
      this.loadBasemap()
      this.loadWidgets()
      // this.eginMap.map.on('click',(evt)=>{
      //   console.log(evt, 'click')
      // })
    },

    initMarker() {
      const [featureGroup, pane] = this.eginMap.MapTools.initPane('markPane', 501)
      console.log('pane', pane)
      this.markerFeatureGroup = featureGroup
      featureGroup.addTo(this.eginMap.map)

    },

    // 获取业务图层ID
    getLayerBusinessId(markerName) {
      let layerId = ''
      switch(markerName){
       case 'PointMarker':  layerId = this.pointLayerId;break;
       case 'PolylineMarker': layerId = this.poylineLayerId;break;
       case 'PolygonMarker': layerId = this.polygonLayerId; break;
       case 'TextMarker': layerId = this.textLayerId; break;
      }
      return layerId
    },

    // 删除某个图层
    markerRemove(layer) {
      //
      let properties = layer.feature.properties
      // 需要判断是否已添加到库中 如果是则需要调用接口
      console.log(layer,'layer')
      //
      const layerId = this.getLayerBusinessId(properties.MARKER_NAME)
      let delParams = {
        layerId: layerId,
        causeField: 'OBJECTID',
        causeValue: properties.OBJECTID
      }
      console.log(delParams)
      //
      // mapApi.deleteFeatures(delParams).then(res => {
      //   console.log('union#delete', res)
      //   const { code } = res
      //   code == 200 ? layer.remove() : ''
      // })

      
    },
     // 确认
    markerConfirm(layer, data, selectedIcon) {
      // 删除
      // 如果data.id不存在则新增
      if (!data.OBJECTID) {
        data.GUID = this.eginMap.Tools.guid()
        // 把数据添加到列表中
        let tmpData = {
          data: data,
          layer: layer,
          geoJSON: layer.toGeoJSON()
        }
        console.log(tmpData)
        layer.options.data = data
      } else {
        // 不能这样 有可能会有修改
        // let OBJECTID = layer.options.OBJECTID
        layer.options.data = data
      }
      let geoJSON = layer.toGeoJSON()
      let feature = this.eginMap.Tools.geoJsonToArcgis(geoJSON)
      // console.log(geoJSON,'geoJSON', data)
      // return
      // return
      // 获取新增或更新的图层ID
      let layerId = this.getLayerBusinessId(data.MARKER_NAME)
      let isUpdate = false
      // 如果有OBJECTID则说明是更新
      isUpdate = !geoJSON.properties.OBJECTID ? true: false
      //
      // console.log(data,'data')
      feature.attributes = data 
      feature.attributes.USER_ID = ''

      // return false;

      // 添加 需要有图层
      let params = {
        layerId: layerId,
        features: JSON.stringify([feature])
      }
      // console.log('return' ,JSON.stringify([feature]))
      // return
      //
      if(isUpdate){
        console.log(params)
        // mapApi.addFeatures(params).then((res)=>{
        //   // console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   code == 200 ? layer.addTo(this.eginMap.map): ''
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }else{
        console.log(params)
        // mapApi.updateFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }
      layer.options.icon = selectedIcon
    },
     // 定位
    markerLocation(layer, geoJSON) {
      const type = geoJSON.geometry.type
      if (type == GeometryType.Polygon || type == GeometryType.LineString) {
        // get
        const bounds = layer.getBounds()
        this.eginMap.map.fitBounds(bounds)
      } else {
        const latlng = layer.getLatLng()
        this.eginMap.flyTo(latlng, 16)
      }
    },
    // 标记移除
    markerListRemove(guid) {
      //
      this.dataLists.map(item => {
        if (item.data.guid == guid) {
          // 删除图层中的数据
          this.removeLayer(item.layer)
        }
      })
      //
      this.removeDataList(guid)
      console.log(this.dataLists, 'removeList')
    },
  }
}
</script>

<style  scoped>
.map-container {
  /* height:  calc(100vh - 70px); */
  height: 100vh;
  width: 100%;
}
.alert {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 500px;
  height: 35px;
  opacity: 0.8;
  z-index: 999;
}
/* .map-container >>> .egin-tools {
  top: 80px;
}
.map-container >>> .leaflet-control-container .leaflet-top.leaflet-right{
  top: 80px;
} */
/* .map-container
  >>> .leaflet-control-container
  >>> .leaflet-top
  >>> .leaflet-right {
  border: 1px solid red !important;
  top: 77px !important;
} */
</style>
```
