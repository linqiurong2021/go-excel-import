<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-17 14:18:01
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-27 16:45:14
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
    'GraphicEdit': () => import('./tools/GraphicEdit'),
    'GraphicSearch': () => import('./tools/GraphicSearch'),
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