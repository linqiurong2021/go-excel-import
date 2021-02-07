import axios from  "../utils/request.js"

const getConfig = (params)=> {
  return axios.request({
    url:"/getConfig",
    method: 'get',
    params
  })
}

// getSelectOptions
const getSelectOptions = (data)=> {
  return axios.request({
    url:"/getSelectOptions",
    method: 'post',
    data
  })
}
// 获取表单列表分页
const getListByPage = ( data ) => {
  return axios.request({
    url:"/getListByPage",
    method: 'post',
    data
  })
}

// 获取详情
const getDataByID = ( params ) => {
  return axios.request({
    url:"/getDataByID",
    method: 'get',
    params
  })
}

// 获取详情
const getFieldsType = ( params ) => {
  return axios.request({
    url:"/getFieldsType",
    method: 'get',
    params
  })
}

// 获取详情
const getFieldsName = ( params ) => {
  return axios.request({
    url:"/getFieldsName",
    method: 'get',
    params
  })
}

// 获取字段
const getFields = ( params ) => {
  return axios.request({
    url:"/getFields",
    method: 'get',
    params
  })
}

// 更新数据
const updateBySysID = ( data ) => {
  return axios.request({
    url:"/updateBySysID",
    method: 'post',
    data
  })
}
// 删除数据
const deleteBySysIDs = ( data ) => {
  return axios.request({
    url:"/deleteBySysIDs",
    method: 'post',
    data
  })
}

// 更新数据
const createData = ( data ) => {
  return axios.request({
    url:"/createData",
    method: 'post',
    data
  })
}

// 获取已有模板列表
const getTemplates = ( data ) => {
  return axios.request({
    url:"/getTemplates",
    method: 'post',
    data
  })
}



// 导入数据
const importData = ( data ) => {
  return axios.request({
    url:"/import",
    method: 'post',
    data
  })
}

// 导出数据
const exportData = ( data ) => {
  return axios.request({
    url:"/export",
    method: 'post',
    data,
    responseType: 'blob'
  })
}
// 下载
const download = (data, fileName)=>{

  let url = window.URL.createObjectURL(data);
  let link = document.createElement('a');
  link.style.display = 'none';
  link.href = url;
  /**
    * 添加属性,并赋指定的值 el.setAttribute('download','zzz')
    * demo: <a href="abc.gif" download="zzz"> 
    * download属性的值即使当前要导出的文件的文件名
    * */
  link.download = fileName
  // link.setAttribute('download', fileName);
  link.click();
  // 释放创建的对象(创建的新的URL必须通过该方法释放)
  window.URL.revokeObjectURL(url);
}


export {
  getConfig,
  getSelectOptions,
  getListByPage,
  getDataByID,
  getFieldsType,
  getFieldsName,
  updateBySysID,
  deleteBySysIDs,
  getFields,
  createData,
  importData,
  getTemplates,
  exportData,
  download
}