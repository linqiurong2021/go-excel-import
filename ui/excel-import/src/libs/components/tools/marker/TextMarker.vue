<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-27 11:34:29
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 11:24:07
 * @ Description: 文本标绘工具组件
 -->

<template>
  <my-dialog
    title="文本标绘"
    :visible="dialogVisible"
    :modal="false"
    :close-on-click-modal="false"
    width="25%"
    draggable
    @close="close"
  >

    <el-form ref="markerText" :rules="rules" :model="temp" label-width="80px">
      <el-form-item class="item" label="名称" prop="NAME">
        <el-input  v-model="temp.NAME" placeholder="请输入名称" size="small"  clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="备注" prop="NOTE">
        <el-input type="note" v-model="temp.NOTE" placeholder="请输入备注" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="字体大小" prop="FONT_SIZE">
        <el-input type="note" v-model="temp.FONT_SIZE" placeholder="请输入字体大小" size="small" clearable></el-input>
      </el-form-item>
      <el-form-item class="item" label="字体颜色" prop="FONT_COLOR">
         <el-color-picker v-model="temp.FONT_COLOR" :predefine="predefineColors"/>
      </el-form-item>
       <el-form-item class="item" label="字体粗细" prop="FONT_WEIGHT">
         <el-select v-model="temp.FONT_WEIGHT" placeholder="请选择">
          <el-option
            v-for="item in fontWeightOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </el-form-item>
      
      <el-form-item class="item" label="是否公开" prop="IS_PUBLIC">
        <el-radio  v-model="temp.IS_PUBLIC" label="1">是</el-radio>
        <el-radio  v-model="temp.IS_PUBLIC" label="0">否</el-radio>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="confirm">确 定</el-button>
      <el-button @click="remove" type="danger">删除</el-button>
    </span>
  </my-dialog>
</template>

<script>
import {
  Button,
  Input,
  Radio,
  Select,
  Option,
  ColorPicker,
  Form,
  FormItem
} from "element-ui";
import MyDialog from '../../dialog/index'
export default {
  name: "PointMarker",
  props: {
    eginMap: {
      type: Object,
      require: true,
    },
    // 显示详情数据
    data: {
      type: Object,
      default() {
        return {}
      }
    },
    // 当前图层
    layer: {
      type: Object,
      require: true,
    },
  },
  components: {
    MyDialog: MyDialog,
    ElInput: Input,
    ElButton: Button,
    ElRadio: Radio,
    ElSelect: Select,
    ElColorPicker: ColorPicker,
    ElOption: Option,
    ElForm: Form,
    ElFormItem: FormItem
  },
  watch: {
    data: {
      handler: function(newVal) {
        this.temp = newVal
      },
      deep: true,
      immediate: true
    }
  },
  data() {
    return {
      // 校验规则
      rules: {
        NAME: [
          {
            required: true,
            message: '名称不能为空',
            trigger: 'blur'
          }
        ],
      },
      // 弹窗
      dialogVisible: false,
      // 默认颜色
      predefineColors: [
        "#ff4500",
        "#ff8c00",
        "#ffd700",
        "#90ee90",
        "#00ced1",
        "#1e90ff",
        "#c71585",
      ],
      // 字体样式
      fontWeightOptions: [
        {
          value: "normal",
          label: "正常",
        },
        {
          value: "bold",
          label: "加粗",
        },
        {
          value: "bolder",
          label: "更粗",
        },
      ],
      map: null,
      temp: {
        NAME: "",
        NOTE: "",
        FONT_SIZE: 12,
        FONT_WEIGHT: "normal",
        FONT_COLOR: "#000000",
        IS_PUBLIC: "0",
        USER_ID: ''
      },
    };
  },
  //
  mounted() {
    this.map = this.eginMap.map;
  },
  methods: {
    resetTemp() {
      this.temp = {
        NAME: "",
        NOTE: "",
        FONT_SIZE: 12,
        FONT_WEIGHT: "normal",
        FONT_COLOR: "#000000",
        IS_PUBLIC: "0",
        USER_ID: ''
      };
    },
    getStyle(iconSize) {
      return `"width:${iconSize[0]}";height:${iconSize[1]}`;
    },
    // DiVICON
    setIcon() {
      let html = `<div style="width:200px;font-size:${this.temp.FONT_SIZE}px;font-weight:${this.temp.FONT_WEIGHT};color:${this.temp.FONT_COLOR}">${this.temp.NAME}</div>`;
      //
      let divIcon = this.eginMap.MapTools.newDivIcon("iconic_location_on_px", [0, 0], html);
      //
      console.log(divIcon)
      const latlng = this.layer.getLatLng();
      let textMarker = this.eginMap.MapTools.newMarker(latlng, {
        icon: divIcon,
      });
      //
      textMarker.addTo(this.map);
      
      // 删除点位数据
      return textMarker;
    },
    closeDialog() {
      this.dialogVisible = false;
    },
    remove() {
      this.closeDialog();
      this.$emit("remove", this.layer);
    },
    confirm() {
      this.$refs['markerText'].validate(valid => {
        if (valid) {
          this.closeDialog();
          this.temp.MARKER_NAME = 'TextMarker'
          this.setIcon()
          this.$emit("confirm", this.layer, this.temp, null);
        }
      })
    },
    //
    close() {
      this.closeDialog();
      // this.$emit("remove", this.layer);
    },
  },
};
</script>

<style scoped>
.item{
  font-size: 14px;
  border-bottom: 1px solid #cccccc;
}
.item >>> .el-input  input {
  border: none
}
</style>