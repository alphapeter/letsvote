import Vue from 'vue'
import Vuex from 'vuex'
import { Rpc } from '../rpc.js'
import ui from './modules/uiState'

Vue.use(Vuex)

const state = {
  activeView: 'left',
  roots: [],
  views: {
    left: {
      selectedRoot: '',
      files: [],
      path: []
    },
    right: {
      selectedRoot: '',
      files: [],
      path: []
    }
  }
}

const actions = {
  init: function ({commit}) {
    Rpc.getRoots()
      .then((response) => {
        commit('init', response.result)
      })
  }
}

const otherStateId = (id) => {
  return id === 'left'
    ? 'right'
    : 'left'
}

const getPathString = path => {
  return path.reduce((acc, p) => {
    return acc + '/' + p
  }, '')
}

const getters = {
  currentState (state) {
    return state.views[state.activeView]
  },
  currentPathString (state) {
    let path = state.views[state.activeView].path
    return getPathString(path)
  },
  otherState (state) {
    return state.views[otherStateId(state.activeView)]
  },
  otherPathString (state) {
    let otherState = state.views[otherStateId(state.activeView)]
    return getPathString(otherState.path)
  },
  otherStateId (state) {
    return otherStateId(state.activeView)
  },
  selectedFiles (state) {
    let selectedFiles = state.views[state.activeView]
      .files
      .filter(file => file.selected)
      .map(file => file.name)

    if (selectedFiles.length) {
      return selectedFiles
    }
    var focusedFile = state.views[state.activeView]
      .files
      .find(file => file.focused)
    return focusedFile ? [focusedFile.name] : []
  }
}
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    init (state, roots) {
      state.roots = roots
      state.ui.state = 'browse'
      state.loading = false
      if (roots.length > 0) {
        state.views['left'].selectedRoot = state.roots[0]
        state.views['right'].selectedRoot = roots.length > 1
          ? state.roots[1]
          : state.roots[0]
      }
    },
    selectRoot (state, message) {
      var viewState = state.views[message.stateId]
      viewState.selectedRoot = message.value
      viewState.path = []
    },
    selectView (state, viewId) {
      state.activeView = viewId
    },
    toggleView (state) {
      state.activeView = otherStateId(state.activeView)
    },
    changePath (state, message) {
      let viewState = state.views[message.stateId]
      viewState.path.push(message.value)
    },
    changePathToParent (state, message) {
      let viewState = state.views[message.stateId]
      viewState.path.pop()
    },
    setPath (state, message) {
      let viewState = state.views[message.stateId]
      viewState.path.splice(message.value, viewState.path.length)
    }
  },
  actions: actions,
  modules: {
    ui
  }
})
