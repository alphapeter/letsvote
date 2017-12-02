import Vue from 'vue'
import Vuex from 'vuex'
import { API } from '../api.js'
import ui from './modules/uiState'
import { EventBus } from '../EventBus.js'

Vue.use(Vuex)

const state = {
  polls: [],
  activeUsers: [],
  me: {},
  votes: {}
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
    vote (state, vote) {
      Vue.set(state.votes, vote.poll_id, vote)
    },
    'POLL_CREATED' (state, poll) {
      state.polls.push(poll)
    },
    'POLL_DELETED' (state, id) {
      state.polls = state.polls.filter(poll => poll.id !== id)
    },
    'POLL_UPDATED' (state, update) {
      let poll = state.polls.find(p => p.id === update.id)
      for (var property in update) {
        if (update.hasOwnProperty(property)) {
          poll[property] = update[property]
        }
      }
    },
    'INIT_VOTES' (state, votes) {
      Vue.set(state, 'votes', votes)
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
    'OPTION_UPDATED' (state, update) {
      let poll = state.polls.find(poll => poll.id === update.poll_id)
      var option = poll.options.find(o => o.id === update.option_id)
      for (var property in update) {
        if (update.hasOwnProperty(property)) {
          option[property] = update[property]
        }
      }
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
    init ({commit, state}, user) {
      commit('initUser', user)
      API.get('/api/polls')
        .then((polls) => {
          commit('init', polls)
          commit('loadingComplete')
        })
      if (state.me) {
        API.get('/api/votes')
          .then((votes) => {
            var map = {}
            for (var vote of votes) {
              map[vote.poll_id] = vote
            }
            commit('INIT_VOTES', map)
          })
      }
    },
    vote ({commit, state}, vote) {
      let votes = state.votes[vote.poll_id]
      if (votes === undefined) {
        votes = { poll_id: vote.poll_id }
      }

      let levels = ['score_1', 'score_2', 'score_3']
      for (let l of levels) {
        if (votes[l] && votes[l] === vote.option_id) {
          Vue.set(votes, l, null)
        }
      }
      if (vote.score !== 0) {
        var v = votes[levels[vote.score - 1]]
        var v_ = votes[levels[vote.score - 2]]
        for (var i = vote.score - 1; i > 0; i--) {
          v && Vue.set(votes, levels[i - 1], v)
          v = v_
          v_ = votes[levels[i - 2]]
        }
        Vue.set(votes, levels[vote.score - 1], vote.option_id)
      }
      API.post('/api/polls/' + vote.poll_id + '/vote', votes)
        .then(result => {
          commit('vote', votes)
        })
    },
    'POLL_CREATED' ({commit}, poll) {
      commit('POLL_CREATED', poll)
    },
    'POLL_DELETED' ({commit}, id) {
      commit('POLL_DELETED', id)
    },
    'POLL_UPDATED' ({commit}, id) {
      commit('POLL_UPDATED', id)
    },
    'OPTION_CREATED' ({commit}, option) {
      commit('OPTION_CREATED', option)
    },
    'OPTION_DELETED' ({commit}, payload) {
      commit('OPTION_DELETED', payload)
    },
    'OPTION_UPDATED' ({commit}, payload) {
      commit('OPTION_UPDATED', payload)
    },
    'USER_CONNECT' ({commit}, user) {
      commit('USER_CONNECT', user)
    },
    'USER_DISCONNECT' ({commit}, user) {
      commit('USER_DISCONNECT', user)
    },
    'CONNECTED_USERS' ({commit}, users) {
      commit('CONNECTED_USERS', users)
    },
    'USER_VOTED' (_, info) {
      EventBus.$emit('USER_VOTED', info)
    }
  },
  modules: {
    ui
  }
})
