<template>
  <div id="app">
    <div class="fileviews">

    </div>
    <dialogs></dialogs>

  </div>
</template>

<script>
  import Dialogs from './components/Dialogs/Dialogs.vue'

  export default {
    name: 'app',
    components: {
      Dialogs
    },
    data () {
      return {
        polls: []
      }
    },
    methods: {

    },
    created () {
      this.$store.dispatch('init')
      let vm = this
      this.eventListener = (e) => {
        if (vm.$store.state.ui.state !== '') {
          return
        }
        if (e.key === 'ArrowLeft' || e.key === 'ArrowRight') {
          console.log('keypress')
        }
      }
      window.addEventListener('keydown', this.eventListener)
    },
    destroyed () {
      window.removeEventListener('keydown', this.eventListener)
    }
  }
</script>

<style>
  #app {
    height: 100%;
  }
</style>
