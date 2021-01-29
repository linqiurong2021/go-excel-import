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


export {
  getConfig,
  getSelectOptions,
  getListByPage
}