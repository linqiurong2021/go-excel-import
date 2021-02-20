<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:13:54
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-28 11:13:38
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
        return 'http://dev.eginsoft.cn:6080/arcgis/rest/services/shuibao2/MapServer/19'
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
    console.log(this.searchOptions,'searchOptions')
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