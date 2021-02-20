/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-28 17:25:05
 * @ Modified by: linqiurong
 * @ Modified time: 2020-10-12 11:04:14
 * @ Description: 业务数据
 */

// egin存储
const state = {
  // 
  baseInfoComponent: {},
  detailInfoComponent: {},
}

const getters = {
  // 
  getBaseInfoComponent: (state) => state.baseInfoComponent,
  getDetailInfoComponent: (state) => state.detailInfoComponent,
}

const mutations = {
  SET_BASEINFO_COMPONENT: (state, component) => {
    state.baseInfoComponent = component
  },
   SET_DETAILINFO_COMPONENT: (state, component) => {
     state.detailInfoComponent = component
  }
}

const actions = {
  setBaseInfoComponent({ commit }, component) {
    commit('SET_BASEINFO_COMPONENT', component)
  },
  setDetailInfoComponent({ commit }, component) {
    commit('SET_DETAILINFO_COMPONENT', component)
  },
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
