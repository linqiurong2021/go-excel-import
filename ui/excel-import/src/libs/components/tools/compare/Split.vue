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