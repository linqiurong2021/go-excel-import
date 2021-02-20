/**
 * @ Author: linqiurong
 * @ Create Time: 2020-09-21 15:54:24
 * @ Modified by: linqiurong
 * @ Modified time: 2021-01-14 09:47:29
 * @ Description: 地图数据
 */

// egin存储
const state = {
  // 图层列表已勾选的ID
  layerCheckedIds: [],
  //
  isEdit: false,
  // 图形编辑时的类型
  graphicEditType: '',
  // 图形编辑时的图形
  editLayer: '',
  // 切割或合并时选中的项
  selected: {
    selected: {},
    key: -1
  },
  // 切割或合并时的弹窗列表信息
  selectDialog: {
    componentName: '',
    dataList: []
  },
  // 确认切割
  confirmCut: {},
  // 确认合并
  confirmUnion: {},
  // 取消合并
  cancleUnion: false,
  // 取消切割
  cancleCut: false,
  // 标绘弹窗
  markTool: {},
  // 定位
  locateTo: {}
}

const getters = {
  // 获取已选择图层
  getLayerCheckedIds: (state) => state.layerCheckedIds,

  isEdit: (state) => state.isEdit,

  graphicEditType: (state) => state.graphicEditType,

  editLayer: (state) => state.editLayer,
  getSelected: (state) => state.selected,
  getSelectDialog: (state) => state.selectDialog,

  confirmCut: (state) => state.confirmCut,
  confirmUnion: (state) => state.confirmUnion,
  cancleUnion: (state) => state.cancleUnion,
  cancleCut: (state) => state.cancleCut,
  markTool: (state) => state.markTool,
  locateTo: (state) => state.locateTo
}

const mutations = {
  SET_LAYER_CHECKED_IDS: (state, ids) => {
    state.layerCheckedIds = ids
  },

  SET_TOOLS_EDIT: (state, edit) => {
    state.isEdit = edit
  },

  SET_GRAPHIC_EDIT_TYPE: (state, type) => {
    state.graphicEditType = type
  },

  SET_EDIT_LAYER: (state, layer) => {
    state.editLayer = layer
  },

  SET_SELECTED: (state, selected) => {
    state.selected = selected
  },
  SET_SELECTE_DIALOG: (state, selectDialog) => {
    state.selectDialog = selectDialog
  },

  SET_CONFIRM_CUT: (state, confirmCut) => {
    state.confirmCut = confirmCut
  },
  SET_CONFIRM_UNION: (state, confirmUnion) => {
    state.confirmUnion = confirmUnion
  },
  SET_CANCLE_UNION: (state, cancleUnion) => {
    state.cancleUnion = cancleUnion
  },
  SET_CANCLE_CUT: (state, cancleCut) => {
    state.cancleCut = cancleCut
  },
  SET_MARK_TOOL: (state, markTool) => {
    state.markTool = markTool
  },
  SET_LOCATE_TO: (state, to) => {
    state.locateTo = to
  }
}

const actions = {
  setLayerCheckedIds({ commit }, ids) {
    commit('SET_LAYER_CHECKED_IDS', ids)
  },
  //
  setToolsEdit({ commit }, edit) {
    commit('SET_TOOLS_EDIT', edit)
  },
  setGraphicEditType({ commit }, type) {
    commit('SET_GRAPHIC_EDIT_TYPE', type)
  },
  setEditLayer({ commit }, layer) {
    commit('SET_EDIT_LAYER', layer)
  },
  setSelected({ commit }, selected) {
    commit('SET_SELECTED', selected)
  },
  setSelecteDialog({ commit }, selectDialog) {
    commit('SET_SELECTE_DIALOG', selectDialog)
  },
  setConfirmCut({ commit }, confirmCut) {
    commit('SET_CONFIRM_CUT', confirmCut)
  },
  setConfirmUnion({ commit }, confirmUnion) {
    commit('SET_CONFIRM_UNION', confirmUnion)
  },
  setCancleUnion({ commit }, cancleUnion) {
    commit('SET_CANCLE_UNION', cancleUnion)
  },
  setCancleCut({ commit }, cancleCut) {
    commit('SET_CANCLE_CUT', cancleCut)
  },
  setMarkTool({ commit }, markTool) {
    commit('SET_MARK_TOOL', markTool)
  },
  setLocateTo({ commit }, to) {
    commit('SET_LOCATE_TO', to)
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
