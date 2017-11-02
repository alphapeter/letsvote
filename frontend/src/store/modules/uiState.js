import { EventBus } from '../../EventBus.js'
const state = {
  state: 'initializing',
  error: null,
  progress: {
    message: '',
    progress: 0,
    max: 0
  }
}

const mutations = {
  rename (state) {
    state.state = 'rename'
  },
  renameWait (state) {
    state.state = 'rename-wait'
  },
  mkdir (state) {
    state.state = 'mkdir'
  },
  mkdirWait (state) {
    state.state = 'mkdir-wait'
  },
  copyWait (state) {
    state.state = 'copy-wait'
  },
  moveWait (state) {
    state.state = 'move-wait'
  },
  deleteFile (state) {
    state.state = 'delete-file'
  },
  deleteFileWait (state) {
    state.state = 'delete-file-wait'
  },
  error (state, error) {
    state.state = 'error'
    state.error = error
  },
  commandFinished (state) {
    state.state = 'browse'
    EventBus.$emit('commandFinished')
  },
  commandCanceled (state) {
    state.state = 'browse'
  },
  startProgress (state, data) {
    state.progress.progress = 0
    state.progress.max = data.max
    state.progress.message = '...'
  },
  progress (state, data) {
    state.progress.message = data.message
    state.progress.progress = data.progress
  }
}

export default {
  state,
  mutations
}
