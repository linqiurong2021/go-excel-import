
const state = {
  // 表单名称
  tableName: {},
  // 搜索参数
  searchParams: '',
  // 表单字段 字段类型 key 名称等配置
  tableConfigs: []
}

// get
const getters = {
  // 获取参数
  getSearchParams: (state) => state.searchParams,
  //
  getTableConfigs: (state) => state.tableConfigs,
  //
  getTableName: (state) => state.tableName
}

const mutations = {
  SET_SEARCH_PARAMS: (state, params) => {
    state.searchParams = params
  },
  SET_TABLE_CONFIGS: (state, configs) => {
    state.tableConfigs = configs
  },
  SET_TABLE_NAME: (state, tableName) => {
    state.tableName = tableName
  }
}

const actions = {
  setSearchParams({ commit }, params) {
    commit('SET_SEARCH_PARAMS', params)
  },
  setTableConfigs({ commit }, configs) {
    commit('SET_TABLE_CONFIGS', configs)
  },
  setTableName({ commit }, tableName) {
    commit('SET_TABLE_NAME', tableName)
  },
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
