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