

# Split 组件(分屏比对组件) 需要与  SplitLayers 是分屏比对右上角的图层列表组件 共同使用

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
      | ElDialog  | ElementUI 弹窗组件 |
      | Layers  | 分屏对比图层组件 |
  
  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  | 弹窗是否可见 |d


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  |  弹窗是否可见 |
      | left  | 左底图 |
      | right  | 右底图 |
      | maps | 左底图与右底图 |
      | showEginLayers  |  是否显示图层 |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | initMaps  | 初始化地图图 |
      | loadBasemap  |加载左、右底图 |
      | loadLayersWidget  | 加载左、右图层列表 |
      | close  | 关闭弹窗|

> 组件代码
```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 18:26:56
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:15:55
 * @ Description: 分屏对比组件
 -->

<template>
  <el-dialog
    class="split-dialog-wrap"
    title="分屏对比"
    :visible="dialogVisible"
    :modal="false"
    :close-on-click-modal="false"
    width="90%"
    :before-close="close"
  >
  <div class="split-wrap">
    <div class="left" ref="left">
       <Layers v-if="maps[0]"  :eginMap="maps[0]" :show="showEginLayers[0]" />
    </div>
    <div class="right" ref="right">
       <Layers v-if="maps[1]"  :eginMap="maps[1]" :show="showEginLayers[1]" />
    </div>
  </div>
 
  </el-dialog>
</template>

<script>
import { Dialog } from "element-ui";
import EginMap from '../../../EginMap'
import EginPosition from '../../../const/EginPosition'
import TiandiTu from '../../../../libs/basemap/TiandiTu'
import SplitLayers from './SplitLayers'
export default {
  name: "Split",
  components: {
    ElDialog: Dialog,
    Layers: SplitLayers
  },
  data() {
    return {
      dialogVisible: false,
      left: null,
      right: null,
      maps: [],
      // 是否显示图层
      showEginLayers: [false, false] // left right
    };
  },
  mounted() {
    //
    // this._initMaps()
  },
  watch: {
    dialogVisible: function(newVal){
      // 加 this.maps.length 防止重复生成地图
      if(newVal && this.maps.length == 0){
        // 需要加nextTick
        this.$nextTick(()=>{
          this.initMaps()
        })
      }
    }
  },
  methods: {
    initMaps() {
      // 加载地图
      const left = new EginMap({container: this.$refs.left})
      const right = new EginMap({container: this.$refs.right})
      // console.log(left, right)
      this.maps = [left, right]
      this.loadBasemap()
      // 同步
      left.splitSync([left.map, right.map])
      // 加载图层
      this.loadLayersWidget(this.maps)

    },
    // 加载底图
    loadBasemap() {
       // 添加底图
      const tiandiTu = new TiandiTu({key: '4b350b4f343fa22cdb2047e93b4d8712'})

      this.maps.map((egin)=>{
        //
        egin.setView([24.48,118.108]).setZoom(11)
        // 卫星图
        let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
        //  卫星注记图
        let tdtSatelliteAnnotion = tiandiTu.addTianDiTuSatelliteAnnotion()
        // // 分组
        let tdtSatelliteGroup = egin.L.layerGroup([tdtSatellite,tdtSatelliteAnnotion])
        egin.map.addLayer(tdtSatelliteGroup)
      })
    },
    // 加载图层ICON
    loadLayersWidget(eginMaps) {
       // 加载图层数据
      const layersOptions = {
        iconfont: 'iconic_layers_px',
        position: EginPosition.TopRight
      }
      eginMaps.map((eginMap,key)=>{
        // 加载图层
        eginMap.loadLayersWidget(layersOptions, ()=>{
          // 
          // this.showEginLayers[key] = !this.showEginLayers[key] // 不能直接用这个 这样使用无效
          //
          this.$set(this.showEginLayers,key, !this.showEginLayers[key])
        })
      })
    },
    close() {
      this.dialogVisible = false
    },
  },
};
</script>

<style scoped>
.split-dialog-wrap{
  top: -50px
}
.split-wrap{
  height: 60vh;
  display: flex;
  justify-content: space-between;
}
.split-wrap .left{
  width: 50%;
}
.split-wrap .right{
  width: 50%;
  background: #cccccc;
}

</style>
```




# Wrap 组件(卷帘比对组件)
> 参数说明
  >> components 组件说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | MyDialog  | 自定义弹窗组件 |
      | ElRadio  | ElementUI Radio组件 |
      | ElTree  | ElementUI Tree组件 |


  >> watch 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  | 弹窗是否可见 |


  >> data 参数说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | dialogVisible  |  弹窗是否可见 |
      | type  | 左底图 |
      | defaultProps  | 树形结构参数 |
      | defaultProps.children | 树形结构参数显示下级字段 |
      | defaultProps.label | 树形结构参数显示名称字段 |
      | EginLayer  |  是否显示图层 |
      | leftLayersData  | 对比图层列表数据 |
      | rightLayersData | 遮罩图层列表数据 |
      | map | 地图实例 |
      | leftLayer  |  当前左图 |
      | rightLayer  |  当前右图 |
      | control  |  卷帘组件 |

  >> 方法 说明
    >>>
      |  字段   |  说明  |
      |  ----  | ----  |
      | initLayers  | 初始化图层 |
      | initMap  | 初始化地图 |
      | change  | 树形结构check事件 |
      | leftChange  | 左图底图变更|
      | rightChange  | 右图底图变更 |
      | close  | 关闭弹窗|

> 组件代码

```html
<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 18:28:30
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:32:53
 * @ Description: 卷帘对比组件 leftLayer 与 rightLayer 只能是 layer 不能是featureGroup
 -->

<template>
  <my-dialog
    class="wrap-wrap"
    title="卷帘对比"
    :visible="dialogVisible"
    :modal="false"
    width="90%"
    @close="close"
    draggable
  >
    <div class="content-container">
      <div class="left">
        <div class="top item">
          <div class="title">
            卷帘模式:
          </div>
          <div class="content">
            <el-radio v-model="type" label="0">垂直分割</el-radio>
            <el-radio v-model="type" label="1">水平分割</el-radio>
          </div>
        </div>
        <div class="middle item">
          <div class="title">
            选择遮罩层：
          </div>
          <div class="content layer-content">
            <el-tree
              :data="rightLayersData"
              show-checkbox
              ref="right"
              node-key="id"
              @check-change="rightChange"
              :props="defaultProps">
            </el-tree>
          </div>
        </div>
        <div class="bottom item">
          <div class="title">
            选择对比层:
          </div>
          <div class="content layer-content">
            <el-tree
              :data="leftLayersData"
              show-checkbox
              ref="left"
              node-key="id"
              @check-change="leftChange"
              :props="defaultProps">
            </el-tree>
          </div>
        </div>
      </div>
      <div class="right">
        <div class="map" ref="mapContainer"></div>
      </div>
    </div>
  </my-dialog>
</template>

<script>
import { Radio,Tree } from "element-ui"
import MyDialog from '../../dialog/index'
import TiandiTu from '../../../../libs/basemap/TiandiTu'
import * as L from 'leaflet'
import "leaflet-side-by-side"
import "leaflet/dist/leaflet.css"

export default {
  name: "Wrap",
  components: {
    MyDialog: MyDialog,
    // ElDialog: Dialog,
    ElRadio: Radio,
    ElTree: Tree
  },
  watch: {
    dialogVisible: function(newVal){
      console.log(newVal,'newVal')
     
      if(newVal && !this.eginMap){
        this.$nextTick(()=>{
           this.initMap()
        })
      }
    },
  },
  data() {
    return {
      dialogVisible: false,
      type: '0',
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      EginLayer: null,
      // 对比图层列表数据
      leftLayersData: [],
      // 遮罩图层列表数据
      rightLayersData: [],
      map: null,
      control: {},
      selectedLayers : [],
      leftLayer: null,
      rightLayer: null
    };
  },
  methods: {
    // 
    initLayers() {
      // 一组
      var osmLayer = L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png')
      osmLayer.name = 'OSM'
      osmLayer.id = 0
      var otherLayer = L.tileLayer('https://stamen-tiles-{s}.a.ssl.fastly.net/watercolor/{z}/{x}/{y}.png')
      otherLayer.name = '其它'
      otherLayer.id = 1
      const tiandiTu = new TiandiTu({key: '4b350b4f343fa22cdb2047e93b4d8712'})
      // 卫星图
      let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
      tdtSatellite.name = '天地图卫星'
      tdtSatellite.id = 0
      //
      let tdtNormal = tiandiTu.addTianDiTuNormalMap()
      tdtNormal.name = '天地图矢量'
      tdtNormal.id = 1

      // this.selectedLayers = [rightLayer, leftLayer, tdtSatellite, tdtNormal ]
      this.leftLayersData = [osmLayer, otherLayer]
      //
      this.rightLayersData = [tdtSatellite, tdtNormal]
      
      osmLayer.addTo(this.map)
      tdtSatellite.addTo(this.map)
      this.leftLayer = osmLayer
      this.rightLayer = tdtSatellite
      
      this.control = L.control.sideBySide(this.leftLayer, this.rightLayer).addTo(this.map)

      this.$refs.right.setCheckedKeys([0])
      this.$refs.left.setCheckedKeys([1])
    },
    initMap() {
      // 有可能无法做数据的卷帘 只能用底图类的(目前卷帘功能未完成)
      var map = L.map(this.$refs.mapContainer).setView([24.48, 118.108], 13)
      this.map = map
      this.initLayers()
    },
    change(index, type) {
      // 如果不存在则直接返回
      // 先取消全选 再选择当前
      if(type == 'right'){ // right 是右
        if(!this.rightLayersData[index]) return
        const layer = this.rightLayersData[index]
        layer.addTo(this.map)
        // 先删除后赋值
        this.map.removeLayer(this.rightLayer)
        this.rightLayer = layer
        // 清除其它选中项
        this.$refs.right.setCheckedKeys([index])
        this.control.setRightLayers(this.rightLayer);
      }else{
        if(!this.leftLayersData[index]) return

        this.map.removeLayer(this.leftLayer)
        const layer = this.leftLayersData[index]
        layer.addTo(this.map)
        this.leftLayer = layer
        this.$refs.left.setCheckedKeys([index])
        this.control.setLeftLayers(this.leftLayer)
      }
    },
    // 
    leftChange(obj,isChecked) {
      // console.log(a,b,c)
      if(isChecked) {
        this.change(obj.id,'left')
      }
    },
    // 
    rightChange(obj,isChecked) {
     if(isChecked) {
        this.change(obj.id,'right')
      }
    },
    close() {
      this.dialogVisible = false
    },
  },
};
</script>

<style scoped>

.content-container{
  display: flex;
  background: #ffffff !important
}
.item{
  margin: 5px;
  padding: 5px;
}
.left{
  width: 20%;
  background: #cccccc;
  height: 600px;
}
.left .content{
  display: flex;
  margin-top: 10px;
}
.right{
  width: 85%;
  background: #666666;
  height: 600px;
}
.right .map {
  height: 100%;
  width: 100%;
}
.title{
  padding: 10px 0px;
  border-bottom: 2px solid #fff;
  color: rgb(0, 118, 255)
}
.layer-content{
  min-height: 180px;
  max-height: 210px;
  border: 1px solid #fff;

  overflow: scroll;
}
.el-tree{
  height: 100%;
  width: 100%;
}
</style>
```



#  代码调用事例

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
