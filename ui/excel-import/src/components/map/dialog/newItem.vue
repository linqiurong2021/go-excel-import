<template>
  <el-dialog
  title="提示"
  :visible.sync="show"
  width="30%"
  :before-close="cancel">
  <div>
    <el-input v-model="inputData" :placeholder="name" clearable></el-input>
  </div>
  <span slot="footer" class="dialog-footer">
    <el-button @click="cancel">取 消</el-button>
    <el-button type="primary" @click="save">确 定</el-button>
  </span>
</el-dialog>
</template>

<script>
import { Button,  Input, Form, Dialog } from "element-ui"
export default {
  name: "newItem",
  props: {
    itemKey: {
      type: String,
      default: ""
    },
    name: {
      type: String,
      default: ""
    }
  },
  components: {
    ElButton: Button,
    ElInput: Input,
    ElForm: Form,
    ElDialog: Dialog
  },
  data() {
    return {
      inputData: "",
      show: false
    }
  },
  methods: {
    save() {
      if(this.inputData == "") {
        this.show = false
        return
      }
      let data = {
        key: this.itemKey,
        value: this.inputData
      }
      this.$emit('getItem',data)
      this.show = false
    },
    cancel() {
      this.show = false
      this.$emit('cancel', this.inputData)
    }
  }
}
</script>

<style scoped>

</style>