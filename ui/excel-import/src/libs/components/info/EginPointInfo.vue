<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-28 18:36:48
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-12 11:07:45
 * @ Description: 自定义弹窗 点 信息
 -->

<template>
    <my-dialog title="基本信息(点)" class="baseinfo-wrap"
    :visible="dialogVisible"
    width="20%"
    draggable
    @close="close">
      <div class="content">
        <div class="item number">编号： {{temp.number}}</div>
        <div class="item name">名称： {{temp.name}}</div>
        <div class="item area">占地面积：{{temp.area}} 公顷</div>
        <div class="item info">项目规模：{{temp.info}}</div>
      </div>
      <div class="images">
        <el-image :src="item.fileurl" v-for="(item,key) in temp.imageList" :key="key" :preview-src-list="[item.fileurl]" />
      </div>
      <div class="btns">
        <div class="iconfont" :class="[item.icon, {'activity' : index == key}]" v-for="(item,key) in btnList" :key="key" @click="iconClick(key)"></div>
      </div>
    </my-dialog>

</template>

<script>
import { Image } from "element-ui";
import MyDialog from '../../../libs/components/dialog/index'
export default {
  name: "EginPointInfo",
  components: {
    MyDialog: MyDialog,
    ElImage: Image
  },
  data() {
    return {
      dialogVisible: true,
      index: 0,
      btnList: [
        {icon: 'iconbianjishuru'},
        {icon: 'icontrash_fill'},
        {icon: 'icontimeline'},
        {icon: 'icongengduo1'},
      ],
      temp: {
        number: '20200910_350200_0001',
        name: '翔安新城核心区东山片区市政道路',
        area: 2320,
        info: '市政道路城市支路9条，道路总长4262.079m市政道路城市支路9条，道路总长4262.079m市政道路城市支路9条，道路总长4262.079m',
        imageList: [{
          fileurl: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
        },{
          fileurl: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
        },{
          fileurl: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
        },{
          fileurl: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
        },{
          fileurl: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
        }],
      }
    };
  },
  methods: {
    close() {
      this.dialogVisible = false;
    },
    iconClick(key) {
      this.index = key
      if(key == this.btnList.length - 1){
        this.$store.dispatch('business/setDetailInfoComponent',{componentsName: 'Detail1', Data: this.temp })
      }
    }
  },
};
</script>

<style scoped>
.baseinfo-wrap{
  right: 80px;
  width: 20%;
  left: 75%;
  opacity: 0.8;
  border-radius: 5px;
  overflow: hidden;
}
.baseinfo-wrap >>> .my-dialog-body{
  margin: 0;
  padding: 30px 0px;
  font-weight: 300;
  line-height: 30px;
}
.baseinfo-wrap >>>.el-icon-circle-close{
  color:#ffffff
}
.title .iconfont {
  cursor: pointer;
}
.content .item {
  border-bottom: 1px solid #eeeeee;
  padding: 10px;
  width: 100%;
}

.content .item .info{
  min-height: 80px;
  max-height: 130px;
  overflow: hidden;
}
.images{
  display: flex;
  flex-wrap: wrap;
  margin: 20px 0px;
}
.images .el-image{
  width: 20%;
  margin: 5px;
  padding: 10px;
}

.btns {
  display: flex;
  justify-content: space-around;
}
.btns .iconfont{
  font-size: 24px;
  cursor: pointer;
  padding: 10px;
  color:  rgb(24, 144, 255);
  border: 1px solid rgb(24, 144, 255);
  border-radius: 5px;
}
.btns .activity{
  background: rgb(24, 144, 255);
  color: #ffffff;
}
</style>