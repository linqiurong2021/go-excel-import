
# CutList 组件是图形切割后的图形列表

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | data  | 切割后的数据列表 |
        | field  | 列表显示字段名称 |
        | colorList  | 颜色列表 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | Table  | ElementUI 表格组件 |
      | Button  | ElementUI 按钮组件 |
      | TableColumn  | ElementUI 表格组件 |
      | MyDialog  | 自定义弹窗组件 |

  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  切割后的数据列表 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | visible  |  弹窗是否可见 |
      | tableData  | 表格数据列表 |
      | confirmData  | 选择保留后的数据 |


  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | handleCurrentChange  | 表格行点击事件 |
      | close  | 关闭 |
      | cancleCut  | 取消切割 |
      | confirmCut  | 确认切割 |

> 组件代码

```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-30 11:13:12
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-15 08:27:11
 * @ Description: 自定义弹窗 切割图形列表组件
 -->

<template>
 <div class="layer-dialog">
   <my-dialog title="切割列表" :visible="visible" draggable @close="close">
      <div>请选择一个图斑继承未切割前的图斑业务属性：</div>
      <el-table :data="tableData" highlight-current-row @current-change="handleCurrentChange">
        <!--根据具体业务修改ID-->
        <el-table-column :property="field" label="图斑编号"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancleCut">
          取消
        </el-button>
        <el-button
          type="primary"
          @click="confirmCut"
        >
          确定
        </el-button>
      </div>
   </my-dialog>
 </div>
</template>

<script>
import MyDialog from '../../dialog/index'
import { Table, Button, TableColumn} from 'element-ui'
export default {
  name: 'CutList',
  props: {
    data: {
      type: Array,
      required: true
    },
    field: {
      type: String,
      require: true
    },
    colorList: {
      type: Array,
      default() {
        return []
      }
    }
  },
  components: {
    ElTable: Table,
    ElButton: Button, 
    ElTableColumn: TableColumn,
    MyDialog 
  },
  data() {
    return {
      visible: true,
      tableData: [],
      confirmData: {}
    }
  },
  watch:{
    data: {
      handler: function(newVal){
        console.log(newVal,'data###watch')
        let tmpDataList = []
        newVal.forEach((item)=>{
          tmpDataList.push(item.feature.properties)
        })
        
        // this.$set(this,this.tableData, tmpDataList)
        this.tableData = tmpDataList
        this.visible = true
        // console.log(this.tableData,'tableData',tmpDataList)
      },
      immediate: true
    }
  },
  methods: {
    handleCurrentChange(currentRow) {
      if(!currentRow) return
      // console.log(this.tableData,'handleCurrentChange###tableData',currentRow)
      this.tableData.map((item, key)=>{
        // console.log(item,'item')
        if(item[this.field] == currentRow[this.field]) {
          //
          const selected = {
            selected: this.data[key],
            key: key
          }
          
          this.confirmData = selected
          this.$store.dispatch('map/setSelected',this.confirmData)
        }
        console.log('item@handleCurrentChange',item)
      })
      // console.log('dispatch')
    },
    close() {
      this.$store.dispatch('map/setCancleCut', true)
      // 重置颜色 取消后则删除
      this.visible = false
    },
    cancleCut() {
      
      this.close()
      this.$store.dispatch('map/setCancleCut', true)
    },
    confirmCut() {
      //
      const confirmData = {
        selected: this.confirmData.selected,
        selectList: this.data,
        key: this.confirmData.key
      }
      this.$store.dispatch('map/setConfirmCut', confirmData)
      // console.log('confirmCut',)
      this.close()
    }
  }
}
</script>

<style scoped>

</style>

```


# UnionList 组件是图形合并选择的图形列表

> 参数说明
  >> props 参数说明
    >>>  
        |  字段   |  说明  |
        |  ----  | ----  |
        | data  | 切割后的数据列表 |
        | field  | 列表显示字段名称 |
        | colorList  | 颜色列表 |


  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | Table  | ElementUI 表格组件 |
      | Button  | ElementUI 按钮组件 |
      | TableColumn  | ElementUI 表格组件 |
      | MyDialog  | 自定义弹窗组件 |

  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | data  |  切割后的数据列表 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | visible  |  弹窗是否可见 |
      | tableData  | 表格数据列表 |
      | confirmData  | 选择保留后的数据 |


  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | handleCurrentChange  | 表格行点击事件 |
      | close  | 关闭 |
      | cancleUnion  | 取消合并 |
      | confirmUnion  | 确认合并 |

> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-30 11:17:06
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 17:04:38
 * @ Description: 自定义弹窗 合并图形列表组件
 -->

<template>
 <div class="layer-dialog">
   <my-dialog title="合并列表" :visible="visible" draggable @close="close">
      <div>请选择一个图斑继承未切割前的图斑业务属性：</div>
      <el-table :data="tableData" highlight-current-row @current-change="handleCurrentChange">
        <el-table-column :property="field" label="名称"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancleUnion">
          取消
        </el-button>
        <el-button
          type="primary"
          @click="confirmUnion"
        >
          确定
        </el-button>
      </div>
   </my-dialog>
 </div>
</template>

<script>

import MyDialog from '../../dialog/index'
import { Table, Button, TableColumn} from 'element-ui'
export default {
  name: 'UnionList',
  props: {
    data: {
      type: Array,
      required: true
    },
    field: {
      type: String,
      require: true
    },
    colorList: {
      type: Array,
      default() {
        return []
      }
    }
  },
  components: {
    ElTable: Table,
    ElButton: Button, 
    ElTableColumn: TableColumn,
    MyDialog 
  },
  data() {
    return {
      visible: true,
      tableData: [],
      confirmData: {}
    }
  },
  watch:{
    data: {
      handler: function(newVal){
     
        let tmpDataList = []
        newVal.forEach((item)=>{
          tmpDataList.push(item.feature.properties)
        })
        this.tableData = tmpDataList
        this.visible = true
        // console.log(this.tableData,'data###watch')
      },
      immediate: true
    }
  },
  methods: {
    close() {
      this.$store.dispatch('map/setCancleUnion', true)
      // 重置颜色 取消后则删除
      this.visible = false
    },
    // 修改颜色
    handleCurrentChange(currentRow, oldRow) {
      console.log(currentRow, oldRow)
      if(!currentRow) return
      // console.log(this.tableData,'handleCurrentChange###tableData',currentRow)
      this.tableData.map((item, key)=>{
        // console.log(item,'item')
        if(item[this.field] == currentRow[this.field]) {
          const selected = {
            selected: this.data[key],
            key: key
          }
          this.confirmData = selected
          this.$store.dispatch('map/setSelected',this.confirmData)
        }
      })

    },
    cancleUnion() {
      this.close()
    },
    confirmUnion() {
       //
      const confirmData = {
        selected: this.confirmData.selected,
        selectList: this.data,
        key: this.confirmData.key
      }
      this.$store.dispatch('map/setConfirmUnion', confirmData)
      // console.log('confirmCut',)
      this.close()
    }
  }
}
</script>

<style scoped>
.my-dialog{
  opacity: 0.9;
}
</style>
```


# 组件调用事例

<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:02
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-19 16:02:53
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