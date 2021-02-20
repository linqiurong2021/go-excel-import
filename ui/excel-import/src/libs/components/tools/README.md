
# Compare 比对工具组件

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | Split  | 分屏比对组件 |
      | Wrap  | 卷帘比对组件 |

  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | compare  | 显示工具 |


> 组件代码

```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:34
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:41:40
 * @ Description: 对比工具栏
 -->

<template>
  <div class="compare-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="compare-tools">
      <div
        class="item compare-item"
        v-for="(item, key) in tools"
        :key="key"
        @click="compare(item.type)"
      >
        {{ item.name }}
      </div>
    </div>
    <div class="compare-tools item close" @click="showTools = false">关闭</div>
    <!--弹窗-->
    <!--动态组件有问题 延迟加载问题-->
    <!-- <component
      ref="compareDialog"
      class="compare-dialog"
      :is="currentCompare"
    ></component> -->
    <Split  ref="splitDialog" />
    <Wrap  ref="wrapDialog"/>
  </div>
</template>

<script>

export default {
  name: "Compare",
  props: {
    eginMap: Object,
  },
  components: {
    Split: () => import("./compare/Split"),
    Wrap: () => import("./compare/Wrap")
  },
  data() {
    return {
      title: "对比",
      showTools: true,
      currentCompare: "",
      tools: [
        {
          name: "分屏对比",
          type: "split",
        },
        {
          name: "卷帘对比",
          type: "wrap",
        },
      ],
    };
  },
  methods: {
    //
    compare(type) {
      type == 'split' ?  this.$refs.splitDialog.dialogVisible = true : this.$refs.wrapDialog.dialogVisible = true
    }
  },
};
</script>

<style scoped>
.compare-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
}
.compare-wrap .title {
  background: #3f51b5;
  height: 50px;
  line-height: 50px;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  opacity: 0.8;
  /* border-radius: 5px; */
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
}
.compare-tools {
  display: flex;
  background: rgba(255, 255, 255, 0.7);
  border-radius: 5px;
  margin: 0px 5px;

}

.compare-wrap .item {
  cursor: pointer;
  /* background: rgba(255, 255, 255, 0.5); */
  margin: 0px 5px;
  border-radius: 5px;
  padding: 10px 10px;
  font-size: 14px;

}
</style>
```

# GraphicSearch 搜索工具组件

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |
        | searchUrl  | 搜索链接 |
        | where  | 搜索条件 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | currentToolName  | 当前工具名称 |
      | drawTool  | 画图工具 |
      | map  | 地图 |
      | pane  | 空间数据存放图层 |
      | searchLayerId  | 搜索图层ID(没有传searchUrl时使用) |
      | searchOptions.url  | 搜索图层链接(没有传searchUrl时使用) |
      | map  | 地图 |
      | pane  | 空间数据存放图层 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |



  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | initDrawTools  | 初始化工具 |
      | getDrawResult  | 获取画完结果 |
      | search  | 搜索 |
      | searchResult  | 搜索结果 |
      | graphicSearch  | 图形搜索按钮 |
      | drawToolsDisable  | 禁用工具 |
      | close  | 关闭 |
      | clear  | 清除 |

> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:13:54
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-27 16:49:19
 * @ Description: 图形搜索工具栏
 -->

<template>
  <div class="graphic-search-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="graphic-search-tools">
      <div
        class="item graphic-search-item"
        :class="{ activity: currentToolName == item.type }"
        v-for="(item, key) in tools"
        :key="key"
        @click="graphicSearch(item.type, key)"
      >
        {{ item.name }}
      </div>
    </div>
    <div class="graphic-search-tools item clear" @click="clear">清除</div>
    <div class="graphic-search-tools item close" @click="close">关闭</div>
  </div>
</template>

<script>
// 获取空间搜索条件
// import { getWhere } from '@utils/map'
export default {
  name: 'GraphicSearch',
  props: {
    eginMap: {
      type: Object,
      required: true
    },
    searchUrl: {
      type: String,
      default() {
        return ''
      }
    },
    where: {
      type: String,
      default() {
        return ''
      }
    }
  },
  data() {
    return {
      title: '图形搜索',
      showTools: true,
      // 当前工具名称
      currentToolName: '',
      // 画图工具
      drawTool: null,
      //
      map: null,
      pane: null,
      drawTools: [],
      // 搜索图层ID
      searchLayerId: 19,
      // 搜索图层链接
      searchOptions: {
        url: ''
      },
      tools: [
        {
          name: '线搜索',
          type: 'line'
        },
        {
          name: '圆搜索',
          type: 'circle'
        },
        {
          name: '矩形搜索',
          type: 'rectangle'
        },
        {
          name: '多边形搜索',
          type: 'polygon'
        }
      ]
    }
  },
  mounted() {
    this.initDrawTools()
    this.map = this.eginMap.map
    const [featureGroup, pane] = this.eginMap.MapTools.initPane(
      'searchPane',
      502
    )
    this.pane = pane
    this.featureGroup = featureGroup
    // console.error(featureGroup, pane)
    // this.searchOptions.url = `http://dev.eginsoft.cn:6080/arcgis/rest/services/shuibao2/MapServer/${this.searchLayerId}`
    this.searchOptions.url = !this.searchUrl ? this.searchUrl : `http://dev.eginsoft.cn:6080/arcgis/rest/services/shuibao2/MapServer/${this.searchLayerId}`
    // 获取画完结果
    this.getDrawResult()
  },
  methods: {
    // 初始化工具
    initDrawTools() {
      const mapTools = this.eginMap.MapTools
      // 初始化工具
      const line = mapTools.initDrawPolyline()
      const circle = mapTools.initCircle()
      const rectangle = mapTools.initRectangle()
      const polygon = mapTools.initDrawPolygon()
      //
      this.drawTools = [line, circle, rectangle, polygon]
    },
    // 获取画完结果
    getDrawResult() {
      //
      this.map.off('draw:created').on('draw:created', evt => {
        console.log(evt, 'evt', evt.layer)
        // 只能传点线面或有效的geojson的点线面
        let layerGeoJSON = evt.layer.toGeoJSON()
        // console.log(layerGeoJSON,'layerGeoJSON')
        this.search(layerGeoJSON, this.currentToolName, evt.layer) // 第三个参数是用于圆选
      })
    },
    // 搜索
    search(geometry, type, layer) {
      // // 测试圆搜索 需要arcgis 10.3+以后才支持
      // let searchOptions = {
      //   url: 'http://192.168.110.201:6080/arcgis/rest/services/Clone/searchTest/FeatureServer/0',
      // }
      // const task = this.eginMap.MapTools.query(searchOptions)
      const task = this.eginMap.MapTools.query(this.searchOptions)
      // 如果是线则用相交的  如果是面则在面里
      if (type == 'line') {
        task.intersects(geometry)
      } else if (type == 'circle') {
        const center = layer._latlng
        const distance = parseInt(layer._mRadius)
        // 需要arcgis server 10.3+以后才支持
        task.nearby(center, distance)
        // task.contains(geometry)
      } else {
        task.within(geometry)
      }
      // 条件
      task.where(this.where) // 测试圆选的时候需要删除这个条件
      // 字段
      task.fields('*')
      //
      // task.token()// 如果有token则需要加这个
      task.run((error, featureCollection) => {
        if (error) {
          console.error('空间搜索失败:', error)
          return
        }
        // 搜索结果
        this.searchResult(featureCollection)
        console.log(featureCollection, '搜索结果')
      }) //
    },
    // 图形搜索结果
    searchResult(featureCollection) {
      // debug
      console.log('RESULT##########')
      this.eginMap.L.geoJSON(featureCollection, {
        pane: this.pane, // 加入pane 提升z-index
        style: {
          color: '#ff0000'
        },
        onEachFeature: (feature, layer) => {
          // 为修改数据获取图层ID
          // layer.options.layerId = this.searchLayerId
          console.error(layer, 'searchResult', feature)
          //
          this.featureGroup.addLayer(layer)
        }
      })
      // 取消选中
      this.currentToolName = ''
      // 画完后禁用
      this.drawToolsDisable()
    },
    // 图形搜索按钮
    graphicSearch(type, key) {
      // console.log(type)
      this.currentToolName = type
      this.drawToolsDisable(key)
    },
    // 禁用工具
    drawToolsDisable(key = -1) {
      this.drawTools.map((tool, index) => {
        if (index != key) {
          tool.disable()
        } else {
          this.drawTool = tool
          tool.enable()
        }
      })
    },
    close() {
      this.clear()
      this.showTools = false
    },
    clear() {
      this.currentToolName = ''
      this.featureGroup.clearLayers()
      this.drawToolsDisable()
      // console.log('clear')
    }
  }
}
</script>

<style lang="scss" scoped>
.graphic-search-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 4px;
  padding: 0px;
}
.graphic-search-wrap .title {
  background: #3f51b5;
  height: 50px;
  line-height: 50px;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
  // @include top-header-gradients;
}
.graphic-search-tools {
  display: flex;
  background: #fff;
  border-radius: 10px;
  margin: 0px 5px;
}

.graphic-search-wrap .item {
  cursor: pointer;
  background: #fff;
  margin: 0px 5px;
  border-radius: 10px;
  padding: 10px 10px;
  font-size: 14px;
}

.item.activity {
  background: #8a9598;
  color: #fff;
}
</style>
```


# GraphicEdit 图形编辑工具组件

> 参数说明
  >> props 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | eginMap  | 地图实例 |
  >> computed 参数说明
    >>> 
      |  字段   |  说明  |
      |  ----  | ----  |
      | editLayer  | 编辑的图形 |
      | graphicEditType  | 编辑的图形类型 |
      | getSelected  | 当前选择的图形(主要用于切割与合并) |
      | confirmCut  | 确认切割 |
      | confirmUnion  | 确认合并 |
      | cancleCut  | 取消切割 |
      | cancleUnion  | 取消合并 |
    
  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 工具栏名称(有颜色区分段) |
      | showTools  | 是否显示工具栏 |
      | layer  | 当前图形 |
      | layers  | 一个图形中有多个图形组合拆分后的图形存储 |
      | map  | 地图实例 |
      | featureGroup  | 切割或合并后的图形存放图层 |
      | drawTool  | 画图工具(主要是线用于做合并与切割还有整形) |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型(组件名称) |
      | pane  | 图形存放的图层 |
      | errorMessage  | 错误信息 |
      | showError | 是否显示错误信息 |
      | colorList  | 颜色(切割与合并时会用到) |
      | searchLayerId  | 搜索图层ID(合并时用到) |
      | businessLayerId  | 空间数据保存时用到 |
      | searchOptions  | 搜索url |
      | layersGroups  | 编辑时有可能同时编辑多个图形 |
     
  >> watch 参数说明
    >>> 
      |  字段   |  说明  |
      |  ----  | ----  |
      | editLayer  | 编辑的图形 |
      | graphicEditType  | 编辑的图形类型 |
      | getSelected  | 当前选择的图形(主要用于切割与合并) |
      | confirmCut  | 确认切割 |
      | confirmUnion  | 确认合并 |
      | cancleCut  | 取消切割 |
      | cancleUnion  | 取消合并 |


  >> 方法 说明
    >>>
      |  名称   |  说明  |
      |  ----  | ----  |
      | setTooltip  | 设置工具提示信息 |
      | authCheck  | 编辑权限判断 |
      | graphicEdit  | 选择哪个工具操作 |
      | cut  | 切割操作 |
      | remove  | 整形操作 |
      | union  | 合并操作 |
      | searchResult  | 合并时空间搜索列表 |
      | edit  | 编辑() |
      | confirm  | 编辑后确认 |
      | clear  | 清除 |
      | clearToolsStatus  | 清除工具状态 |

> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:02
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 08:35:34
 * @ Description: 图形编辑工具栏
 -->

<template>
  <div class="graphic-edit-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="graphic-edit-tools">
      <div
        class="item graphic-edit-item"
        :class="{ activity: item.type == graphicEditType }"
        v-for="(item, key) in tools"
        :key="key"
        @click="graphicEdit(item.type)"
      >
        {{ item.name }}
      </div>
    </div>
    <div class="graphic-edit-tools item clear" @click="clear">清除</div>
    <div class="graphic-edit-tools item confirm" @click="confirm">确定</div>
    <div class="graphic-edit-tools item close" @click="close">关闭</div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import clipPolygonUtil from '../../ClipPolygon'
import L from 'leaflet'
import { feature } from '@turf/turf';
export default {
  name: 'GraphicEdit',
  props: {
    eginMap: {
      type: Object,
      required: true
    }
  },
  computed: {
    ...mapGetters({
      editLayer: 'map/editLayer',
      graphicEditType: 'map/graphicEditType',
      getSelected: 'map/getSelected',
      confirmCut: 'map/confirmCut',
      confirmUnion: 'map/confirmUnion',
      cancleCut: 'map/cancleCut',
      cancleUnion: 'map/cancleUnion'
    })
  },
  data() {
    return {
      title: '图形编辑',
      showTools: true,
      // 当前
      layer: null,
      // 一个图形中有多个图形组合拆分后的图形存储
      layers: [],
      map: null,
      featureGroup: null,
      // 画图工具
      drawTool: null,
      tools: [
        {
          name: '编辑',
          type: 'edit'
        },
        {
          name: '切割',
          type: 'cut'
        },
        {
          name: '合并',
          type: 'union'
        },
        {
          name: '整形',
          type: 'remove'
        }
      ],
      pane: '',
      // 错误信息
      errorMessage: '',
      // 是否显示错误信息
      showError: false,
      // 颜色
      colorList: ['#00FFFF', '#00FF00', '#0000FF', '#FFFF000', '#FF00FF'],
      // 搜索图层ID
      searchLayerId: 19,
      // 业务图层ID
      businessLayerId: 'a191c15515da4016b71681394e5a8479',
      // 需要合并的图层链接
      searchOptions: {
        url: ''
      },
      // 同时编辑多个图形
      layersGroups: []
    }
  },
  watch: {
    // 不能用deep 会死循环
    getSelected: function(newVal) {
      // 需要修改选中的样式
      const selectedKey = newVal.key
      const selected = newVal.selected
      let graphicEditType = this.graphicEditType
      // 设置当前选择的样式
      selected.setStyle({ color: '#ff0000' })
      //
      const tmpLayers = this.featureGroup.getLayers()

      tmpLayers.forEach((item, key) => {
        if (graphicEditType == 'cut') {
          // 是否需要设置index ? 另一个切割线看不到
          key === selectedKey
            ? ''
            : item.setStyle({ color: this.colorList[key] })
        } else if (this.graphicEditType == 'union') {
          key !== selectedKey ? item.setStyle({ color: '#00ff00' }) : ''
        }
      })
    },
    // 图层编辑
    editLayer(layer) {
      // console.error(layer,'editLayer')
      // 图层
      if (layer) {
        //
        let graphicEditType = this.graphicEditType
        const hasAuth = this.authCheck(graphicEditType, layer.options.operate)
        this.layer = layer
        // console.log(this.layer, layer, this.graphicEditType,'watch')
        if (hasAuth) {
          switch (graphicEditType) {
            case 'edit':
              this.edit()
              break
            case 'cut':
              this.setTooltip('切割')
              this.cut(graphicEditType)
              break
            // case 'union' : this.union();break; // 不走这个流程有单独的
            case 'remove':
              this.setTooltip('整形')
              this.cut(graphicEditType)
              console.log('remove')
              break
          }
        }
      }
      // this.layer = null
    },
    showTools(newVal) {
      if (newVal) {
        this.clearToolsStatus()
      }
    },
    // 切割弹窗确认
    confirmCut(newVal) {
      //
      let dataList = newVal.selectList
      let key = newVal.key
      let selected = newVal.selected
      console.log(selected, 'selected')
      //  this.layer = selected
      console.log(newVal, 'newVal')
      if (!dataList[key]) {
        this.$store.dispatch('alert/setAlertInfo', {
          title: '未选选择,请选选择'
        })
        // console.error("选中的列表不存在")
        return
      }
      // console.log('dataList',selected)
      // dataList[key]
      // console.log(this.layer,'layer')
      // console.log(dataList,'dataList')
      dataList[key].feature.properties = this.layer.feature.properties
      // console.log(key,'key')
      dataList.map((layer, index) => {
        if (index == key) {
          // console.log(layer.toGeoJSON(),'SaveGeoJSON')
          const features = this.eginMap.Tools.geoJsonToArcgis(layer.toGeoJSON())
          // console.log(features,'features')
          const params = {
            layerId: this.businessLayerId,
            features: JSON.stringify([features])
          }
          console.log(params.features, 'features')
          // mapApi.updateFeatures(params).then(res => {
          //   // console.log("cut#update", res)
          //   let { code } = res
          //   const type = code == 200 ? 'success' : 'error'
          //   this.$store.dispatch('alert/setAlertInfo', {
          //     title: res.msg,
          //     type: type
          //   })
          // })
          // console.log(this.eginMap.Tools.geoJsonToArcgis(layer.toGeoJSON()),'ArcgisGeometry#Update')
        }
      })
      // 清除图层数据
      this.featureGroup.clearLayers()
    },
    // 合并弹窗确认
    confirmUnion(newVal) {
      //
      let dataList = newVal.selectList
      let key = newVal.key
      // let selected = newVal.selected
      // console.log(selected,'selected')
      // this.layer = selected
      if (!dataList[key]) {
        this.$store.dispatch('alert/setAlertInfo', {
          title: '未选选择,请选选择'
        })
        // console.error("选中的列表不存在")
        return
      }
      // console.log(dataList,'dataList')
      // 数据处理
      // const layerId = selected.options.layerId
      //
      let objectIds = []
      // 保存列表中的空间数据
      let unionResult = []
      // 保存列表中选中的数据
      let selectedGeoJOSN = {}
      dataList.map((layer, index) => {
        // 需要判断是单个还是多个
        console.log(layer.toGeoJSON(), 'layer.toGeoJSON()')
        let layerGeoJSON = layer.toGeoJSON()
        // let polygon = []
        // 如果是多图形的
        if (layerGeoJSON.geometry.type == 'MultiPolygon') {
          //
          layer.feature.geometry.coordinates.forEach(item => {
            unionResult.push(item[0])
          })
        } else {
          unionResult.push(layer.feature.geometry.coordinates[0])
        }

        const objectId = layer.feature.properties.OBJECTID
        //
        index == key
          ? (selectedGeoJOSN = layer.toGeoJSON())
          : objectIds.push(objectId)
      })
      // console.log('unionResult,',unionResult)
      selectedGeoJOSN.geometry.coordinates = unionResult
      //
      let feature = this.eginMap.Tools.geoJsonToArcgis(selectedGeoJOSN)
      let params = {
        layerId: this.businessLayerId,
        features: JSON.stringify([feature])
      }
      console.log(params.features, 'features')
      // 更新空间库数据
      // mapApi.updateFeatures(params).then(res => {
      //   console.log('union#Update', res)
      //   // const { code } = res
      //   // 更新操作提示
      //   let { code } = res
      //   const type = code == 200 ? 'success' : 'error'
      //   this.$store.dispatch('alert/setAlertInfo', {
      //     title: res.msg,
      //     type: type
      //   })
      //   if (code == 200) {
      //     let delParams = {
      //       layerId: this.businessLayerId,
      //       causeFiled: 'OBJECTID',
      //       causeValue: objectIds.join(',')
      //     }
      //     // 更新
      //     mapApi.deleteFeatures(delParams).then(res => {
      //       let { code } = res
      //       // 删除操作提示
      //       const type = code == 200 ? 'success' : 'error'
      //       this.$store.dispatch('alert/setAlertInfo', {
      //         title: res.msg,
      //         type: type
      //       })
      //       // 清除图层数据
      //       this.featureGroup.clearLayers()
      //     })
          
      //   }
      // })
    },
    cancleUnion(newVal) {
      newVal ? this.clear() : ''
    },
    cancleCut(newVal) {
      newVal ? this.clear() : ''
    }
  },
  mounted() {
    
    this.map = this.eginMap.map
    const [featureGroup, pane] = this.eginMap.MapTools.initPane('editPane', 501)
    this.pane = pane
    this.featureGroup = featureGroup
    // 实例化画线工具
    this.drawTool = this.eginMap.MapTools.initDrawPolyline()
    //
    console.log(this.drawTool, 'drawTools','####')
    // 搜索图层
    this.searchOptions.url = `http://dev.eginsoft.cn:6080/arcgis/rest/services/shuibao2/MapServer/${this.searchLayerId}`
  },
  methods: {
    //
    setTooltip(name) {
      L.drawLocal.draw.handlers.polyline.tooltip = {
        start: `点击地图开始${name}`,
        cont: `点击地图继续${name}`,
        end: `点击最后一个点完成${name}`
      }
    },
    // 操作判断
    authCheck(
      graphicEditType,
      operate = { edit: false, cut: false, union: false, remove: false }
    ) {
      //
      let errMsg = ''

      switch (graphicEditType) {
        case 'edit':
          errMsg = !operate.edit ? '图层不支持编辑' : ''
          break
        case 'cut':
          errMsg = !operate.cut ? '图层不支持切割' : ''
          break
        case 'union':
          errMsg = !operate.union ? '图层不支持合并' : ''
          break
        case 'remove':
          errMsg = !operate.remove ? '图层不支持整形' : ''
          break
      }
      if (errMsg != '') {
        //
        // this.drawTool.disable()
        this.$store.dispatch('alert/setAlertInfo', { title: errMsg })

        return false
      }
      console.log(errMsg, 'err')
      return true
    },
    graphicEdit(type) {
      if (type == 'union') {
        this.setTooltip('合并操作')
        this.union()
      }

      this.$store.dispatch('map/setGraphicEditType', type)
    },
    // 切割
    cut(type) {
      console.error('type', type, this.drawTool)
      if (this.drawTool == null) {
        return
      }
      // if(this.drawTool == null) {
      //   this.drawTool = this.eginMap.MapTools.initDrawPolyline()
      //   // this.eginMap.L.drawLocal.draw.handlers.polyline.tooltip = {
      //   //   start: '点击地图开始裁剪',
      //   //   cont: '点击地图开始裁剪',
      //   //   end: '双击地图或点击最后一个点完成裁剪'
      //   // }
      // }
      // 开启画线模式
      this.drawTool.enable()
      //
      // console.error(type,'type')
      //
      this.map.off('draw:created').on('draw:created', evt => {
        // 动态组件名称
        const componentName = type == 'cut' ? 'CutList' : 'RemoveList'

        console.error('draw:created')
        // 线的geoJSON
        let lineGeoJSON = evt.layer.toGeoJSON()

        let setlectedGraphicFeature = this.layer.feature
        try {
          let tmpLayers = []

          var clipPolygons = clipPolygonUtil.polygonClipByLine(
            setlectedGraphicFeature,
            lineGeoJSON
          )
          // console.log(clipPolygons, '切割后的图形')
          if (type == 'cut') {
            // 添加切割后的图形到地图上
            clipPolygons.features.map((polygon, key) => {
              //
              let selectGrapihcProperties = JSON.parse(
                JSON.stringify(setlectedGraphicFeature.properties)
              )
              // 修改ID
              selectGrapihcProperties.ID =
                selectGrapihcProperties.ID + '_' + key
              // console.log(key)
              this.eginMap.L.geoJSON(polygon, {
                pane: this.pane, // 加入pane 提升z-index
                style: {
                  color: this.colorList[key]
                },
                onEachFeature: (feature, layer) => {
                  layer.feature.properties = selectGrapihcProperties
                  // 再次分配
                  tmpLayers.push(layer)
                  // layer.addTo(this.map)
                  this.featureGroup.addLayer(layer)
                }
              })
            })

            // 弹窗需要选择某个为主图 另一个为副图
            this.$store.dispatch('map/setSelecteDialog', {
              componentName: componentName,
              dataList: tmpLayers,
              colorList: this.colorList
            }) //
          } else if (type == 'remove') {
            //
            this.remove(clipPolygons)
          }
        } catch (error) {
          const title = `${error.state}:${error.message}`
          this.$store.dispatch('alert/setAlertInfo', { title: title })
        }
        this.clearToolsStatus()
      })
      this.map.off('draw:created', () => {
        console.log('取消监听')
      })
    },
    // 对切割后的图形做处理
    remove(clipPolygons) {
      //
      // console.log(clipPolygons,'clipPolygons')
      let max = 0,
        index = 0

      // 获取最大面积的那个
      clipPolygons.features.map((item, key) => {
        //
        this.eginMap.L.geoJSON(item, {
          onEachFeature: (feature, layer) => {
            //
            let latlngs = layer.getLatLngs()
            console.log(latlngs, 'latlngs')
            // 获取面积
            const area = this.eginMap.MapTools.getArea(latlngs[0])

            max = max > area ? max : area
            index = max > area ? index : key
            console.error(area, 'area', index, 'index')
          }
        })
      })
      console.error(max, 'area', index, 'index', feature)

      // return '';
      // 删除选中的 并调用接口 删除空间数据
      // this.map.removeLayer(this.getSelected)
      this.featureGroup.removeLayer(this.getSelected)
      //
      // 保留最大面积的那个 删除选中的与小的那个
      clipPolygons.features.map((item, key) => {
        //
        if (key === index) {
          this.eginMap.L.geoJSON(item, {
            pane: this.pane, // 加入pane 提升z-index
            style: {
              color: this.colorList[0]
            },
            onEachFeature: (feature, layer) => {
              //
              console.log( feature, 'feature######', layer.toGeoJSON())
              // layer.addTo(this.map)
              this.featureGroup.addLayer(layer)
            }
          })
          //
          const features = this.eginMap.Tools.geoJsonToArcgis(item)
          // console.log(features,'feature######remove')
          // 字段
          features.attributes = this.layer.feature.properties

          // 这里应该还要一个弹窗
          this.$confirm('是否确认整形,确认后数据不可恢复?', '整形提醒', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          })
            .then(() => {
              console.log('confirm', 'aaaa')
              // 更新操作
              // a191c15515da4016b71681394e5a8479
              let params = {
                layerId: this.businessLayerId, // 测试ID
                features: JSON.stringify([features])
              }
              //
              // console.log(this.layer,'layer')
              console.log(params['features'], 'params')
              // mapApi.updateFeatures(params).then(res => {
              //   console.log('res', res)
              //   let { code } = res
              //   const type = code == 200 ? 'success' : 'error'
              //   this.$store.dispatch('alert/setAlertInfo', {
              //     title: res.msg,
              //     type: type
              //   })
              // })
            })
            .catch(() => {
              console.log('cancle')
              this.clear()
            })

          // console.log("ArcgisGeometry#Update", this.eginMap.Tools.geoJsonToArcgis(item))
          // console.log(item, 'item##GEOJSON需要保留的数据')
        }
        //  else{
        //   // 这里应该不会删除操作
        //   console.log("ArcgisGeometry#Delete", this.eginMap.Tools.geoJsonToArcgis(item))
        //  }
      })

      // console.log(max, index,'result')
    },
    // 合并操作
    union() {
      //
      // 判断是否否有合并功能
      this.drawTool.enable()
      //
      this.map.off('draw:created').on('draw:created', evt => {
        // 线经过的面
        console.log(evt, 'evt')
        const layer = evt.layer
        const lineGeoJSON = layer.toGeoJSON()
        //
        const task = this.eginMap.MapTools.query(this.searchOptions)
        task.intersects(lineGeoJSON)
        // 条件
        task.where("MAP_DIVISION_ID='350213'") // 测试圆选的时候需要删除这个条件
        // 字段
        task.fields('*')
        //
        // task.token()// 如果有token则需要加这个
        task.run((error, featureCollection) => {
          if (error) {
            console.error('空间搜索失败:', error)
            return
          }
          // 搜索结果
          this.searchResult(featureCollection)
          console.log(featureCollection, '搜索结果')
        }) //
      })
      console.error('union')
    },
    // 弹窗
    searchResult(featureCollection) {
      //
      let tmpLayers = []
      this.eginMap.L.geoJSON(featureCollection, {
        pane: this.pane, // 加入pane 提升z-index
        style: {
          color: '#ff0000'
        },
        onEachFeature: (feature, layer) => {
          // 为修改数据获取图层ID
          layer.options.layerId = this.searchLayerId
          this.featureGroup.addLayer(layer)
          tmpLayers.push(layer)
        }
      })
      const componentName = 'UnionList'
      // 弹窗并显示
      // console.log(this.featureGroup,'featureGroup')
      this.$store.dispatch('map/setSelecteDialog', {
        componentName: componentName,
        dataList: tmpLayers,
        colorList: this.colorList
      }) //
    },
    // 编辑
    edit() {
      const geoJSON = this.layer.toGeoJSON()
      console.log('end return 无效')
      this.layersGroups = []
      if (geoJSON.geometry.type == 'MultiPolygon') {
        let tmpLayers = []
        geoJSON.geometry.coordinates.forEach(item => {
          const polygon = {
            type: 'Polygon',
            coordinates: item
          }
          // 拆分成多个图形再同时开启编辑状态
          this.eginMap.L.geoJSON(polygon, {
            onEachFeature: (feature, layer) => {
              this.layersGroups.push(layer)
              this.featureGroup.addLayer(layer)
              layer.editing.enable()
              // 再次分配
              tmpLayers.push(layer)
            }
          })
        })
      } else {
        // 需要判断是否是 MultiPolygon
        this.layer.editing.enable()
        this.layersGroups.push(this.layer)
      }

      // this.featureGroup.editing.enable()
    },
    confirm() {
      // 编辑后的GeoJSON数据
      let returnGeoJSON = null
      if (this.graphicEditType == 'edit') {
        // let tmpLayers = []
        //
        if (this.layersGroups.length == 0) {
          //
          // tmpLayers.push(this.layer)

          this.layer.editing.disable()
          // 这里可能需要调用更新空间数据接口 还需要转换成arcgis 存储的格式
          returnGeoJSON = this.layer.toGeoJSON()
        } else if (this.layersGroups.length > 0) {
          let coordinates = []
          //
          this.layersGroups.forEach(layer => {
            // 禁用编辑
            layer.editing.disable()
            const geoJSON = layer.toGeoJSON()
            coordinates.push(geoJSON.geometry.coordinates[0])
            // tmpLayers.push(layer)
          })
          console.log(coordinates, 'features')

          // 重新更新编辑后的JSON数据
          let geoJSON = this.layer.toGeoJSON()
          geoJSON.geometry.coordinates = coordinates
          returnGeoJSON = geoJSON
        }
      }

      // return
      // 清除
      this.clearToolsStatus()
      // console.log(this.layer, 'this.layer')
      const features = this.eginMap.Tools.geoJsonToArcgis(returnGeoJSON)
      console.log('ArcgisGeometry#Update', features)

      //
      let params = {
        layerId: this.businessLayerId,
        features: JSON.stringify([features])
      }
      console.log(params.features, 'features')
      // mapApi.updateFeatures(params).then(res => {
      //   let { code } = res
      //   const type = code == 200 ? 'success' : 'error'
      //   this.$store.dispatch('alert/setAlertInfo', {
      //     title: res.msg,
      //     type: type
      //   })
      //   // console.log('edit###update', res)
      // })

      // console.log('ArcgisGeometry#Update', this.eginMap.Tools.geoJsonToArcgis(returnGeoJSON))
    },
    clear() {
      // 清除
      this.featureGroup.clearLayers()
      this.clearToolsStatus()
      // console.log('clear')
      this.$store.dispatch('map/setGraphicEditType', '')
    },
    close() {
      this.clear()
      this.showTools = false
    },
    // 清除工具上的状态
    clearToolsStatus() {
      //
       // 清除编辑状态
      this.layersGroups.forEach((layer)=>{
        // 禁用
        layer.editing.disable()
      })
      // 禁用
      this.drawTool.disable()
      this.$store.dispatch('map/setEditLayer', null)
      // this.$store.dispatch('map/setGraphicEditType', '')
      // this.graphicEditType = ''
    }
  }
}
</script>

<style lang="scss" scoped>
.graphic-edit-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 4px;
  padding: 0px;
}
.graphic-edit-wrap .title {
  height: 50px;
  background: #3f51b5;
  opacity: 0.8;
  line-height: 50px;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
}
.graphic-edit-tools {
  display: flex;
  background: #fff;
  border-radius: 10px;
  margin: 0px 5px;
}

.graphic-edit-wrap .item {
  cursor: pointer;
  background: #fff;
  margin: 0px 5px;
  border-radius: 10px;
  padding: 10px 10px;
  font-size: 14px;
}

.item.activity {
  background: #8a9598;
  color: #fff;
}
</style>
```


# Location 定位组件


> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | ElButton  | ElementUI按钮组件 |
      | ElInput  | ElementUI Input组件 |
      | ElForm  | ElementUI表单组件 |
      | ElFormItem  | ElementUI表单组件 |
      | MyDialog  | 自定义弹窗组件 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | rules  | 表单校验规则 |
      | dialogVisible  | 弹窗是否可见 |
      | temp  | 临时变量 |
      | temp.longitude  | 经度 |
      | temp.latitude  | 纬度 |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | handleClose  | 关闭 |
      | cancle  | 取消  |
      | confirm  | 确认 |

> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:08
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:50:49
 * @ Description: 自定义 定位组件
 -->

<template>
  <my-dialog title="定位" class="location-wrap" :visible="dialogVisible" :modal="false" :close-on-click-modal="false"
      width="25%"
      draggable
      @close="handleClose">
      <el-form ref="location" :rules="rules" :model="temp" label-width="80px">
        <el-form-item class="item" label="经度" prop="longitude">
          <el-input  v-model="temp.longitude" placeholder="请输入经度" size="small"  clearable></el-input>
        </el-form-item>
         <el-form-item class="item" label="经度" prop="latitude">
          <el-input  v-model="temp.latitude" placeholder="请输入纬度" size="small"  clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="confirm">确 定</el-button>
      </span>
  </my-dialog>
</template>

<script>
import {  Button, Input, Form,FormItem } from 'element-ui'
import Dialog from '../../../libs/components/dialog/index'
 export default {
   name: 'Location',
   components: {
     ElButton: Button,
     ElInput: Input,
     MyDialog: Dialog,
     ElForm: Form,
     ElFormItem: FormItem
   },
   props: {
     eginMap: {
       require: true
     }
   },
   data() {
     return {
        // 校验规则
        rules: {
          longitude: [
            {
              required: true,
              message: '经度不能为空',
              trigger: 'blur'
            }
          ],
          latitude: [
            {
              required: true,
              message: '纬度不能为空',
              trigger: 'blur'
            }
          ],
        },
       dialogVisible: false,
       temp: {
        longitude: '',
        latitude: ''
       },
     }
   },
   methods: {
     handleClose() {
       this.dialogVisible = false
     },
     cancle() {
       this.dialogVisible = false
       this.latitude = ''
       this.longitude = ''
     },
     confirm() {
       //
      // console.error(this.latitude, this.longitude)
      //
      this.$refs['location'].validate(valid => {
        if (valid) {
          // 定位成功 需要添加当前位置信息
          //
          // 点击后删除
          //
          const divIcon = this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [30,30])
          const latlng = [this.temp.latitude, this.temp.longitude]
          const layer = this.eginMap.MapTools.newMarker(latlng,{icon: divIcon}).bindTooltip('点击删除').openTooltip();
          layer.on('click', ()=>{
            layer.remove()
          })  
          layer.addTo(this.eginMap.map)
          this.eginMap.panTo([this.temp.latitude, this.temp.longitude])
          this.dialogVisible = false
        }
      })
      
      //  console.log('确定')
     }
   }
 }
</script>

<style scoped>
.location-wrap{
  opacity: 0.95;
}
.item{
  font-size: 14px;
  border-bottom: 1px solid #cccccc;
}
.item >>> .el-input  input {
  border: none
}
</style>

```

# Mark 标绘工具组件

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | eginMap  | 地图实例 |

  >> computed 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | markTool  | 标绘工具 |
  

  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | title  | 标题 |
      | featureGroup  | 标绘的图形存放图层 |
      | pane  | 图层显示 |
      | showTools | 是否显示标绘工具 |
      | currentMarker  | 当前工具项 |
      | drawTools  | 标绘工具实例(点工具 线工具 面工具 文本工具) |
      | currentTool  | 当前标绘工具 |
      | currentLayer  | 当前图层 |


      | tmpLayers  | 临时图层(未确认的) |
      | layers  | 已添加的图层(图形)(已确认的) |
      | show  | 是否显示弹窗[点弹窗、线弹窗、面弹窗、文本弹窗] |
      | drawTools  | 标绘工具实例(点工具 线工具 面工具 文本工具) |
      | dataLists  | 存储后的数组 |
      | map  | 地图实例 |
      | tools  | 标绘工具项 |
      | tools.name  | 标绘工具项名称 |
      | tools.type  |  标绘工具项类型(组件名称)|
      | markerData  | 弹窗填写的数据 |
      | pointLayerId  |  标绘点对应的空间数据业务图层ID|
      | poylineLayerId  | 标绘线对应的空间数据业务图层ID |
      | polygonLayerId |  标绘面对应的空间数据业务图层ID|
      | textLayerId  | 标绘文本对应的空间数据业务图层ID |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | initTools  | 初始化工具 |
      | mark  | 标绘项  |
      | getResult  | 标绘完成 |
      | remove  | 删除 |
      | markLocation  | 定位(列表使用-目前列表已取消)  |
      | markRemove  | 确认(列表使用-目前列表已取消) |
      | removeDataList  | 删除列表数据-目前列表已取消  |
      | removeLayer  | 删除图形 |
      | confirm  | 确认  |
      | getBusinessId  | 获取当前图形业务ID |
      | clear  | 清除-目前已取消  |
     
> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:21
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 14:12:59
 * @ Description: 标绘工具栏
 -->

<template>
  <div class="mark-wrap" v-show="showTools">
    <div class="title">{{ title }}</div>
    <div class="mark-tools">
      <div
        class="item mark-item"
        v-for="(item, key) in tools"
        :key="key"
        @click="mark(item.type, key)"
      >
        {{ item.name }}
      </div>
    </div>
    <!-- 20201016 与潘核实取消清除功能--->
    <!-- <div class="item clear" @click="clear">清除</div> -->
    <div class="mark-tools item close" @click="showTools = false">关闭</div>
    <!--弹窗-->
    <!-- <component
      ref="markerDialog"
      class="marker-dialog"
      :is="markTool.componentName"
      :eginMap="eginMap"
      @confirm="confirm"
      @remove="remove"
      :layer="markTool.currentLayer"
      :data="markTool.markerData"
    ></component> -->
    <!-- <mark-list
      :list="dataLists"
      @remove="markRemove"
      @location="markLocation"
    ></mark-list> -->
  </div>
</template>

<script>
import GeometryType from '../../const/GeometryType'

// import { mapApi } from '@api/mapApi/map'

// 工具提示
import './tooltip.js'
import { mapGetters } from 'vuex'
export default {
  name: 'Mark',
  props: {
    eginMap: Object
  },
  components: {
    // PointMarker: () => import('./marker/PointMarker'),
    // PolylineMarker: () => import('./marker/PolylineMarker'),
    // PolygonMarker: () => import('./marker/PolygonMarker'),
    // TextMarker: () => import('./marker/TextMarker'),
    // MarkList: () => import('./MarkList')
  },
  computed: {
    ...mapGetters({
      markTool: 'map/markTool'
    })
  },
  data() {
    return {
      title: '标绘',
      // 存放标绘的图层
      featureGroup: null,
      // 图层显示
      pane: null,
      //
      showTools: true,
      // 当前选中是哪个组件
      currentMarker: 'PointMarker',
      // 点工具 线工具 面工具 文本工具
      drawTools: [],
      // 当前工具
      currentTool: [],
      // 当前操作的图层
      currentLayer: null,
      // 临时图层(未确认的)
      tmpLayers: [],
      // 已添加的图层(图形)(已确认的)
      layers: [],
      // 是否显示弹窗
      show: [false, false, false, false],
      // 存储后的数组
      dataLists: [],
      map: null,
      tools: [
        {
          name: '点标绘',
          type: 'PointMarker'
        },
        {
          name: '线标绘',
          type: 'PolylineMarker'
        },
        {
          name: '面标绘',
          type: 'PolygonMarker'
        },
        {
          name: '文字标绘',
          type: 'TextMarker'
        }
      ],
      // 标记的数据
      markerData: {},
      businessLayerId: 'a191c15515da4016b71681394e5a8479',
      //
      pointLayerId: 'e6e25e52f6384d52a6e941cfe7ad551d', // 标绘点
      //
      poylineLayerId: 'b197aba58a934a8ba576632f44938f54', // 标绘线
      //
      polygonLayerId: '7fc9aff62ea640bc82de56ebd4d458b0', // 标绘面
      //
      textLayerId: 'b7c72b2ed74c448ba1664bcd22dc8878 ', // 文字标绘
      // 7fc9aff62ea640bc82de56ebd4d458b0 标绘面
      // b197aba58a934a8ba576632f44938f54 标绘线
      // e6e25e52f6384d52a6e941cfe7ad551d 标绘点
      // b7c72b2ed74c448ba1664bcd22dc8878 文字标绘
    }
  },
  //
  mounted() {
    this.map = this.eginMap.map

    const [featureGroup, pane] = this.eginMap.MapTools.initPane('markPane', 501)
    this.pane = pane
    this.featureGroup = featureGroup
    featureGroup.addTo(this.map)

    this.initTools()
    // 监听新增结果
    this.getResult()
  },
  methods: {
    //
    initTools() {
      //
      const markerOptions = {
        //
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }

      const textMarkerOptions = {
        icon: this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [
          20,
          20
        ])
      }
      const drawPoint = this.eginMap.MapTools.initMarker(markerOptions)
      const drawPolyline = this.eginMap.MapTools.initDrawPolyline()
      const drawPolygon = this.eginMap.MapTools.initDrawPolygon()
      const drawText = this.eginMap.MapTools.initTextMarker(textMarkerOptions)
      const drawTools = [drawPoint, drawPolyline, drawPolygon, drawText]

      this.drawTools = drawTools
    },
    mark(type, key) {
      // 当前组件
      this.currentMarker = type
      //
      this.drawTools.map((tool, index) => {
        index == key ? tool.enable() : tool.disable()
        if (index == key) {
          // 开启哪个工具
          this.currentTool = tool
        }
      })
    },
    getResult() {
      this.map.off('draw:created').on('draw:created', e => {
        const layer = e.layer
        // 当前图层
        this.currentLayer = layer
        // 加入到数组中  删除时需要用到
        // this.tmpLayers.push(layer)
        // 先添加后地图上  如果选择删除后 再删除
        // layer.addTo(this.map)

        this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker , currentLayer: layer , markerData: {} })
        console.error('getResult#showDialog')
        layer.on('click', evt => {
          //
          let data = evt.target.options.data
          // 弹窗相应弹窗
          this.currentMarker = data.MARKER_NAME
          this.markerData = data
          // 
          this.$store.dispatch('map/setMarkTool', {componentName: this.currentMarker, currentLayer: layer , markerData: data })
          //
          // setTimeout(() => {
          //   this.$refs.markerDialog.dialogVisible = true
          // }, 300)
          console.log(evt, data)
        })
        // 不在这里做数据添加 需要在确认时做数据添加 点击时需要获取数据框中的数据
        // this.featureGroup.addLayer(layer)

        // setTimeout(() => {
        //   this.$refs.markerDialog.dialogVisible = true
        // }, 300)
        // console.log(this.show, 'show')
      })
    },
    // 删除
    remove(layer) {
      // 当前layer在datalist的位置再删除
      this.removeLayer(layer)
    },
    // 定位
    markLocation(layer, geoJSON) {
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
    markRemove(guid) {
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
    // 从dataList中删除
    removeDataList(guid) {
      //
      this.dataLists.map((item, key) => {
        if (item.data.guid == guid) {
          this.dataLists.splice(key, 1)
        }
      })
    },
    // 删除某个图层
    removeLayer(layer) {
      //
      const hasLayer = this.featureGroup.hasLayer(layer)
      //
      if (hasLayer) {
        console.log(layer.options, 'options')
        //
        // 需要判断是否已添加到库中 如果是则需要调用接口
        if(!layer.options.data.OBJECTID){
          //
          let delParams = {
            layerId: '',
            causeField: 'OBJECTID',
            causeValue: layer.options.data.OBJECTID
          }
          console.log(delParams)
          //
          // mapApi.deleteFeatures(delParams).then(res => {
          //   console.log('union#delete', res)
          // })
        }
       
        
        this.removeDataList(layer.options.data.guid)
        this.featureGroup.removeLayer(layer)
      }
      // 重置
      this.markerData = {}
    },

    // 确认
    confirm(layer, data, selectedIcon) {
      // 删除
      // this.tmpLayers.pop()
      // 如果data.id不存在则新增
      data.MARKER_NAME = this.currentMarker
      if (!data.GUID) {
        data.GUID = this.eginMap.Tools.guid()
        // 把数据添加到列表中
        let tmpData = {
          data: data,
          layer: layer,
          geoJSON: layer.toGeoJSON()
        }
        this.dataLists.push(tmpData)
        //
        layer.options.data = data
        // 添加到保存
        this.layers.push(layer)
        //
        this.featureGroup.addLayer(layer)
      } else {
        //
        layer.options.data = data
      }
      let geoJSON = layer.toGeoJSON()
      let feature = this.eginMap.Tools.geoJsonToArcgis(geoJSON)
      // console.log(geoJSON)
      // return
      // 获取新增或更新的图层ID
      let layerId = this.getBusinessId(geoJSON.geometry.type)
      let isUpdate = false
      // 如果有OBJECTID则说明是更新
      isUpdate = !geoJSON.properties.OBJECTID ? true: false
      //
      feature.attributes = data 
      feature.attributes.USER_ID = ''


      // 添加 需要有图层
      let params = {
        layerId: layerId,
        features: JSON.stringify([feature])
      }
      console.log('return' , console.log(params))
      // return
      //
      if(isUpdate){
        // mapApi.addFeatures(params).then((res)=>{
        //   console.log(res,'res')
        //   const {code} = res
        //   const type = code == 200 ? 'success' : 'error'
        //   this.$store.dispatch('alert/setAlertInfo', {
        //     title: res.msg,
        //     type: type
        //   })
        // })
      }else{
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
      
      this.currentLayer.options.icon = selectedIcon
      this.show = false
    },
    // 获取业务图层ID
    getBusinessId(type) {
      let layerId = ''
      switch(type){
       case 'Point':  layerId = this.pointLayerId;break;
       case 'LineString': layerId = this.polygonLayerId;break;
       case 'Polygon': layerId = this.polygonLayerId; break;
      }
      return layerId
    },
    clear() {
      // 清除未点击确定的数据
      this.featureGroup.clearLayers()
      //  console.log('clear')
      // this.tmpLayers.map((layer)=>{
      //   this.map.remove(layer)
      // })
    }
  }
}
</script>

<style lang="scss" scoped>
.mark-wrap {
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255,255,255,0.8);
  border-radius: 4px;
  padding: 0px;
}
.mark-wrap .title {
  height: 50px;
  line-height: 50px;
  opacity: 0.8;
  background: #3f51b5;
  vertical-align: middle;
  align-items: center;
  text-align: center;
  color: #fff;
  margin-right: 10px;
  width: 85px;
  font-size: 14px;
  // @include top-header-gradients;
}
.mark-tools {
  display: flex;
  background: #ffffff;
  border-radius: 10px;
  margin: 0px 5px;
}

.mark-wrap .item {
  cursor: pointer;
  background: #ffffff;
  margin: 0px 5px;
  border-radius: 10px;
  padding: 10px 10px;
  font-size: 14px;
}
</style>
```
# MarkList 标绘数据列表

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | list  | 标绘数据列表 |

  >> component 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | MyDialog  | 自定义弹窗 |
  

  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  | 是否显示标绘列表 |
      | markList  | 标绘数据列表 |
      | geometryType  | 标绘数据项类型(点、线、面、文本) |


  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | getIcon  | 获取标绘数据项的图标(icon) |
      | close  | 关闭标绘列表  |
      | location  | 定位 |
      | remove  | 移除 |
     
> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-10-14 09:00:04
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-14 11:18:21
 * @ Description: 标绘后的数据列表
 -->

<template>
 <my-dialog title="标记列表" class="mark-list-wrap"
    :visible="dialogVisible"
    width="10%"
    draggable
    @close="close">
      <div class="content list" v-if="markList.length > 0">
         <div class="list-item" v-for="(item) in markList" :key="item.data.guid"> 
            <div @click="location(item.layer, item.geoJSON)"><i class="item iconfont" :class="getIcon(item.geoJSON.geometry.type)"></i> {{item.data.name}}</div>
            <div class="item"><i class="iconfont icondelete" @click="remove(item.data.guid)"></i></div>
         </div>
      </div>
      <div v-else>
         未有标记记录
      </div>
   </my-dialog>
</template>

<script>
import MyDialog from '../dialog/index'
import GeometryType from '../../const/GeometryType'
export default {
   name: 'MarkList',
   components: {
      MyDialog
   },
   props: {
      list: {
         type: Array,
         default() {
            return []
         }
      }
   },
   data() {
      return {
         dialogVisible: true,
         markList: [],
         geometryType: GeometryType
      }
   },
   mounted() {
      this.markList = this.list
   },
   methods: {
      getIcon(type) {
         let icon = ''
         switch(type){
            case GeometryType.Polygon: icon='icontuxing';break;
            case GeometryType.LineString: icon='icontimeline';break;
            case GeometryType.Point: icon='iconsoph_location';break;
            default: {
               icon = 'iconword' ;break;
            }
         }
         return icon
      },
      //
      close() {
         this.dialogVisible = false
      },
      location(layer, geoJSON) {
         this.$emit('location', layer, geoJSON)
      },
      remove(guid) {
         this.$emit('remove', guid)
      }
   }
}
</script>

<style scoped>
.mark-list-wrap{
   opacity: 0.9;
   border-radius: 5px;
   overflow: hidden;
   min-width: 170px;
}
.mark-list-wrap .list{
   min-height: 100px;
   max-height: 60vh;
   overflow: scroll
}
.mark-list-wrap .list .list-item{
   line-height: 30px;
   display: flex;
   justify-content: space-between;
   border-bottom: 1px solid #e0dbdb
}
.list-item .iconfont{
   font-size: 20px;
   cursor: pointer;
}
.list-item .item{
   cursor: pointer;
}
</style>
```

# Measure 测量工具组件
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
      | map  | 地图实例 |
      | title  | 标题 |
      | drawPolyline  | 画线工具 |
      | drawPolygon  | 画面工具 |
      | drawType  | 画的类型 |
      | layers  | 存储图形 |
      | tools  | 工具项 |
      | tools.name  | 工具项名称 |
      | tools.type  | 工具项类型 |

  >> watch 参数说明
    >>>  
      |  字段   |  说明  |
      |  ----  | ----  |
      | showTools  | 是否显示弹窗 |


  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | measure  | 工具点击 |
      | length  | 长度测量  |
      | area  | 面积测量 |
      | getResult  | 图形结果 |
      | clear  | 清除 |

> 组件代码

```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:13:45
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 13:53:57
 * @ Description: 测量工具栏
 -->

<template>
 <div class="measure-wrap" v-show="showTools">
   <div class="title">{{title}}</div>
   <div class="measure-tools">
     <div class="item measure-item" v-for="(item,key) in tools" :key="key" @click="measure(item.type)">{{item.name}}</div>
   </div>
   <div class="measure-tools item clear" @click="clear">清除</div>
   <div class="measure-tools item close" @click="showTools = false">关闭</div>
 </div>
</template>

<script>
import './tooltip.js'
 export default {
   name: 'Measure',
   props: {
     eginMap: {
       require: true
     }
   },
   data() {
     return {
       map: null,
       title: '测量',
       showTools: true,
       drawPolyline: null,
       drawPolygon: null,
       delete: null,
       drawType: null,
       layers: [],
       tools: [
        {
          name: '长度测量',
          type: 'length',
        },
        {
          name: '面积测量',
          type: 'area',
        }
       ]
     }
   },
   watch: {
     // 关闭
     showTools: function(newVal){
       if(!newVal){
        this.drawPolyline.disable()
        this.drawPolygon.disable()
       }
     }
   },
   mounted() {
    this.map = this.eginMap.map
    //
    this.drawPolyline = this.eginMap.MapTools.initDrawPolyline()
    this.drawPolygon = this.eginMap.MapTools.initDrawPolygon()
    // 获取结果
    this.getResult()
   },
   methods: {
     measure( type ) {
       type == 'length' ? this.length() :  this.area()
       this.drawType = type
     },
     length() {
       // 禁用画面
       this.drawPolygon.disable()
       // 开启画线
       this.drawPolyline.enable()
     },
     area() {
       // 禁用画本
       this.drawPolyline.disable()
       // 开启画面
       this.drawPolygon.enable()
     },
     //
     getResult() {
      // console.log
      this.map.off('draw:created').on('draw:created', (e)=>{
        const layer = e.layer;
        // 加入到数组中  删除时需要用到
        this.layers.push(layer)
        layer.addTo(this.map)
        // 判断是length 还是 area 需要计算长度或面积
        //
        const text = this.eginMap.MapTools.getPopupContent(layer)
        // const text = this.drawType == 'length' ? `长度:` : `面积`
        layer.bindTooltip(text)
      })
     },
     clear(){
       this.layers.forEach(layer => {
         if(this.map.hasLayer(layer)){
           this.map.removeLayer(layer)
         }
       })
     }
   }
 }
</script>

<style scoped>
.measure-wrap{
  display: flex;
  vertical-align: middle;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
}
.measure-wrap .title{
    background: #3F51B5;
    opacity: 0.8;
    height: 50px;
    line-height: 50px;
    vertical-align: middle;
    align-items: center;
    text-align: center;
    color: #fff;
    /* border-radius: 5px; */
    margin-right: 10px;
    width: 85px;
    font-size: 14px
}
.measure-tools{
  display: flex;
  background: #fff;
  /* background: rgba(255, 255, 255, 0.7); */
  border-radius: 5px;
  margin: 0px 5px;
}

.measure-wrap .item{
  cursor: pointer;
  
  margin: 0px 5px;
  border-radius: 5px;
  padding: 10px 10px;
  font-size: 14px;
}

</style>
```

# Print 打印工具组件(暂无 目前使用window.print())




# 以上各组件调用事例:

```html
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
  <component ref="tools" class="tool" :is="currentTool" :eginMap="eginMap" ></component>
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
      default() {
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
      tools: [
        {
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
      if(this.$refs.tools){
        this.$refs.tools.showTools = !this.$refs.tools.showTools
      }
      this.currentTool = item.type
      // console.log(item)
      if(item.type == 'Print') {
        // this.clearActivity()
        this.eginMap.printMap()
      }else if(item.type == 'Download'){
        // this.clearActivity()
        this.eginMap.exportMap('map')
      }else if(item.type == 'Location'){
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
.egin-tools-wrap{
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
.egin-tools .egin-tools-item{
  padding: 10px;
  font-size: 12px;
  height: 40px;
  border-bottom:1px solid #ccc;
  box-sizing: border-box;
  cursor: pointer;
}
.egin-tools .egin-tools-item .egin-tools-menu{
  height: 20px;
  line-height: 20px;
  display: flex;
  vertical-align: middle;
}
.egin-tools .egin-tools-item i{
  padding-right: 10px;
  font-size: 18px;
}

.tool{
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
.leaflet-control-easyPrint-button-export{
  display: none;
}
</style>

```