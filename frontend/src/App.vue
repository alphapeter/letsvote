<template>
  <div id="app">
    <div class="fileviews">
      <file-view id="left" @click.native="selectView('left', $event)"></file-view>
      <file-view id="right" @click.native="selectView('right', $event)"></file-view>
    </div>
    <actions></actions>
    <dialogs></dialogs>

  </div>
</template>

<script>
  import Actions from './components/Actions.vue'
  import FileView from './components/FileView.vue'
  import Header from './components/Header.vue'
  import Dialogs from './components/Dialogs/Dialogs.vue'

  export default {
    name: 'app',
    components: {
      Actions,
      FileView,
      'appHeader': Header,
      Dialogs
    },
    data () {
      return {
        roots: [],
        selectedView: null
      }
    },
    methods: {
      selectView (viewId, event) {
        if (event.target.nodeName === 'SELECT') {
          return
        }
        this.$store.commit('selectView', viewId)
      }
    },
    created () {
      this.$store.dispatch('init')
      let vm = this
      this.eventListener = (e) => {
        if (vm.$store.state.ui.state !== 'browse') {
          return
        }
        if (e.key === 'ArrowLeft' || e.key === 'ArrowRight') {
          vm.$store.commit('toggleView')
          e.preventDefault()
        }
      }
      window.addEventListener('keydown', this.eventListener)
    },
    destroyed () {
      if (this.disableButtons) {
        return
      }
      window.removeEventListener('keydown', this.eventListener)
    }
  }
</script>

<style>
  #app {
    height: 100%;
  }

  .fileviews {
    position: relative;
    width: 100%;
    height: calc(100% - 61px);
  }

</style>
