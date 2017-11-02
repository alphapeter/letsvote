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
  error (state, error) {
    state.state = 'error'
    state.error = error
  }
}

export default {
  state,
  mutations
}
