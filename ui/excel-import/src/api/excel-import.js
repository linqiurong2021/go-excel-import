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


export {
  getConfig,
  getSelectOptions,
  getListByPage,
  getDataByID,
  getFieldsType,
  getFieldsName
}