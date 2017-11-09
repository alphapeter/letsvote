const state = {
  state: 'loading',
  error: null
}

const mutations = {
  loadingComplete (state) {
    state.state = ''
  },
  loading (state) {
    state.state = 'loading'
  },
  login (state) {
    state.state = 'login'
  },
  error (state, error) {
    state.state = 'error'
    state.error = error
  },
  commandCanceled (state) {
    state.state = ''
  }
}

export default {
  state,
  mutations
}
