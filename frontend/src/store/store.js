import Vue from 'vue'
import Vuex from 'vuex'
import { API } from '../api.js'
import ui from './modules/uiState'

Vue.use(Vuex)

const state = {
  polls: [],
  activeUsers: [],
  me: null
}

const getters = {}
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    initUser (state, user) {
      this.state.me = user
    },
    init (state, polls) {
      this.state.polls = polls
    },
    selectPoll (state, message) {
      state.selectedPoll = state.polls[message.poll_id]
    },
    'POLL_CREATED' (state, poll) {
      state.polls.push(poll)
    },
    'POLL_DELETED' (state, id) {
      state.polls = state.polls.filter(poll => poll.id !== id)
    },
    'OPTION_CREATED' (state, option) {
      var poll = state.polls.find(poll => poll.id === option.poll_id)
      poll.options.push(option)
    },
    'OPTION_DELETED' (state, payload) {
      var poll = state.polls.find(poll => poll.id === payload.poll_id)
      var index = poll.options.findIndex(o => o.id === payload.option_id)
      poll.options.splice(index, 1)
    },
    'USER_CONNECT' (state, user) {
      if (!state.activeUsers.some(u => u.id === user.id)) {
        state.activeUsers.push(user)
      }
    },
    'USER_DISCONNECT' (state, user) {
      state.activeUsers = state.activeUsers.filter(u => u.id !== user.id)
    },
    'CONNECTED_USERS' (state, users) {
      state.activeUsers = users
    }
  },
  actions: {
    init ({commit}, user) {
      commit('initUser', user)
      API.getPolls()
        .then((polls) => {
          commit('init', polls)
          commit('loadingComplete')
        })
    },
    'POLL_CREATED' ({commit}, poll) {
      commit('POLL_CREATED', poll)
    },
    'POLL_DELETED' ({commit}, id) {
      commit('POLL_DELETED', id)
    },
    'OPTION_CREATED' ({commit}, option) {
      commit('OPTION_CREATED', option)
    },
    'OPTION_DELETED' ({commit}, payload) {
      commit('OPTION_DELETED', payload)
    },
    'USER_CONNECT' ({commit}, user) {
      commit('USER_CONNECT', user)
    },
    'USER_DISCONNECT' ({commit}, user) {
      commit('USER_DISCONNECT', user)
    },
    'CONNECTED_USERS' ({commit}, users) {
      commit('CONNECTED_USERS', users)
    }
  },
  modules: {
    ui
  }
})
