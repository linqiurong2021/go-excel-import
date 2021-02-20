<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-23 09:13:53
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-27 16:39:26
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
    eginMap: {},
    // 图例配置
    legendConfig: {
      type: Function,
      default() {
        return ()=>{}
      }
    }
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
      // legendConfig: legendConfig,
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