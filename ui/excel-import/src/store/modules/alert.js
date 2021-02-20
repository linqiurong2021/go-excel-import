/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-30 09:58:15
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-09 17:28:52
 * @ Description: 弹窗信息 类似 toast 
 */

// egin存储
const state = {
  // 
  alertInfo: {},
}

const getters = {
  // 
  alertInfo: (state) => state.alertInfo,
}

const mutations = {
  SET_ALERT_INFO: (state, { title = '', type = 'error'}) => {
    state.alertInfo = {
      title: title,
      type: type
    }
  }
}

const actions = {
  setAlertInfo({ commit }, info) {
    console.log('info',info)
    commit('SET_ALERT_INFO', info)
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
