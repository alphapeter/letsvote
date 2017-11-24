<template>
  <div id="app">
    <pageHeader></pageHeader>
    <left-panel></left-panel>
    <right-panel></right-panel>
    <activePolls></activePolls>
    <dialogs></dialogs>
  </div>
</template>

<script>
  import ActivePolls from './components/ActivePolls.vue'
  import Dialogs from './components/Dialogs/Dialogs.vue'
  import PageHeader from './components/PageHeader.vue'
  import LeftPanel from './components/LeftPanel.vue'
  import RightPanel from './components/RightPanel.vue'

  export default {
    name: 'app',
    components: {
      ActivePolls,
      Dialogs,
      PageHeader,
      LeftPanel,
      RightPanel
    },
    data () {
      return {
      }
    },
    methods: {

    },
    created () {
      this.$store.dispatch('init')

      var url = 'ws://' + window.location.host + '/tap'
      var ws = new WebSocket(url)
      var store = this.$store
      ws.onmessage = function (message) {
        var event = JSON.parse(message.data)
        store.dispatch(event.event, event.payload)
        console.log(event)
      }
    },
    destroyed () {

    }
  }
</script>

<style>
  #app {
    height: 100%;
  }
</style>
