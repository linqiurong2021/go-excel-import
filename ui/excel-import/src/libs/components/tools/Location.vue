<!--
 * @ Author: linqiurong
 * @ Create Time: 2020-09-25 09:14:08
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-17 10:50:49
 * @ Description: 自定义 定位组件
 -->

<template>
  <my-dialog title="定位" class="location-wrap" :visible="dialogVisible" :modal="false" :close-on-click-modal="false"
      width="25%"
      draggable
      @close="handleClose">
      <el-form ref="location" :rules="rules" :model="temp" label-width="80px">
        <el-form-item class="item" label="经度" prop="longitude">
          <el-input  v-model="temp.longitude" placeholder="请输入经度" size="small"  clearable></el-input>
        </el-form-item>
         <el-form-item class="item" label="经度" prop="latitude">
          <el-input  v-model="temp.latitude" placeholder="请输入纬度" size="small"  clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="confirm">确 定</el-button>
      </span>
  </my-dialog>
</template>

<script>
import {  Button, Input, Form,FormItem } from 'element-ui'
import Dialog from '../../../libs/components/dialog/index'
 export default {
   name: 'Location',
   components: {
     ElButton: Button,
     ElInput: Input,
     MyDialog: Dialog,
     ElForm: Form,
     ElFormItem: FormItem
   },
   props: {
     eginMap: {
       require: true
     }
   },
   data() {
     return {
        // 校验规则
        rules: {
          longitude: [
            {
              required: true,
              message: '经度不能为空',
              trigger: 'blur'
            }
          ],
          latitude: [
            {
              required: true,
              message: '纬度不能为空',
              trigger: 'blur'
            }
          ],
        },
       dialogVisible: false,
       temp: {
        longitude: '',
        latitude: ''
       },
     }
   },
   methods: {
     handleClose() {
       this.dialogVisible = false
     },
     cancle() {
       this.dialogVisible = false
       this.latitude = ''
       this.longitude = ''
     },
     confirm() {
       //
      // console.error(this.latitude, this.longitude)
      //
      this.$refs['location'].validate(valid => {
        if (valid) {
          // 定位成功 需要添加当前位置信息
          //
          // 点击后删除
          //
          const divIcon = this.eginMap.MapTools.newDivIcon('iconic_location_on_px', [30,30])
          const latlng = [this.temp.latitude, this.temp.longitude]
          const layer = this.eginMap.MapTools.newMarker(latlng,{icon: divIcon}).bindTooltip('点击删除').openTooltip();
          layer.on('click', ()=>{
            layer.remove()
          })  
          layer.addTo(this.eginMap.map)
          this.eginMap.panTo([this.temp.latitude, this.temp.longitude])
          this.dialogVisible = false
        }
      })
      
      //  console.log('确定')
     }
   }
 }
</script>

<style scoped>
.location-wrap{
  opacity: 0.95;
}
.item{
  font-size: 14px;
  border-bottom: 1px solid #cccccc;
}
.item >>> .el-input  input {
  border: none
}
</style>