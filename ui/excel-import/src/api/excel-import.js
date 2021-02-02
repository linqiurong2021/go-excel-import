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
  createData
}