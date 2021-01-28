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
const getTableListByPage = ( data ) => {
  return axios.request({
    url:"/getTableListByPage",
    method: 'post',
    data
  })
}


export {
  getConfig,
  getSelectOptions,
  getTableListByPage
}