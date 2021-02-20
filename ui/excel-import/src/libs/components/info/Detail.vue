<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-28 15:55:56
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-12 11:10:31
 * @ Description: 自定义弹窗 详情
 -->

<template>
    <my-dialog title="详情" class="baseinfo-wrap"
    :visible="dialogVisible"
    width="80%"
    draggable
    @close="close">
      <div class="content">
        
        <el-tabs v-model="leftActiveName" tab-position="left" style="80%">
          <el-tab-pane label="基本信息" name="info">
            <el-tabs v-model="activeName" @tab-click="handleClick">
              <el-tab-pane label="基本属性" name="baseinfo">
                <div class="item number">编号： <el-input :value="data.number"/></div>
                <div class="item name">名称：<el-input :value="data.name"/> </div>
                <div class="item area">占地面积：<el-input :value="data.area"/> 公顷</div>
                <div class="item info">项目规模：<el-input :value="data.info"/></div>
           
              </el-tab-pane>
              <el-tab-pane label="核查处置" name="check">
                <div class="images">
                  <el-image :src="item.fileurl" v-for="(item,key) in data.imageList" :key="key" :preview-src-list="[item.fileurl]" />
                </div>
              </el-tab-pane>
            </el-tabs>

          </el-tab-pane>
          <el-tab-pane label="关联的项目" name="project">
            关联的项目内容
          </el-tab-pane>
        </el-tabs>
   </div>
       
  
       
    </my-dialog>

</template>

<script>
import { Image, Input, Tabs, TabPane } from "element-ui";
import MyDialog from '../../../libs/components/dialog/index'
export default {
  name: "Detail",
  props: {
    data: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  components: {
    MyDialog: MyDialog,
    ElImage: Image,
    ElInput: Input,
    ElTabs: Tabs,
    ElTabPane: TabPane
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
      leftActiveName: 'info',
      activeName: 'baseinfo',
    };
  },
  methods: {
    close() {
      this.dialogVisible = false;
    },
    handleClick() {
      console.log('handleClick')
    },
    iconClick(key) {
      this.index = key
    }
  },
};
</script>

<style scoped>
.baseinfo-wrap{
  right: 80px;
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