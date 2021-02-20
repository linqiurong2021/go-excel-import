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