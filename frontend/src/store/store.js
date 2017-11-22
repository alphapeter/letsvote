import Vue from 'vue'
import Vuex from 'vuex'
import { API } from '../api.js'
import ui from './modules/uiState'

Vue.use(Vuex)

const state = {
  polls: []
}

const getters = {}
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    init (state, polls) {
      this.state.polls = polls
    },
    selectPoll (state, message) {
      state.selectedPoll = state.polls[message.pollId]
    },
    addPoll (state, poll) {
      state.polls.push(poll)
    },
    addOption (state, option) {
      var poll = state.polls.find(poll => poll.id === option.pollId)
      poll.options.push(option)
    }
  },
  actions: {
    init ({commit}) {
      API.getPolls()
        .then((polls) => {
          commit('init', polls)
          commit('loadingComplete')
        })
    },
    login ({commit}) {
      commit('login')
    }
  },
  modules: {
    ui
  }
})
