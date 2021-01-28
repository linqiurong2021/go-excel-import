
const state = {
  searchParams: {}
}

// get
const getters = {
    // 拾取坐标
  getSearchParams: (state) => state.searchParams,
}

const mutations = {
  SET_SEARCH_PARAMS: (state, params) => {
    state.searchParams = params
  }
}

const actions = {
  setSearchParams({ commit }, params) {
    commit('SET_SEARCH_PARAMS', params)
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
