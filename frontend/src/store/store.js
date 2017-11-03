import Vue from 'vue'
import Vuex from 'vuex'
import { API } from '../api.js'
import ui from './modules/uiState'

Vue.use(Vuex)

const state = {
  polls: []
}

const actions = {
  init: function ({commit}) {
    API.getPolls()
      .then((polls) => {
        commit('init', polls)
        commit('loadingComplete')
      })
  }
}

const getters = { }
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    init (state, polls) {
      this.state.polls = polls
    },
    selectPoll (state, message) {
      state.selectedPoll = state.polls[message.pollId]
    }
  },
  actions: actions,
  modules: {
    ui
  }
})
