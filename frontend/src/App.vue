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
    created () {
      var user = null
      try {
        var a = this.$cookie.get('lets_vote.authenticated')
        var b = decodeURIComponent(a)
        user = JSON.parse(b)
      } catch (e) {
        user = null
      }
      this.$store.dispatch('init', user)
      this.initWebSocket()

      var lastTime = (new Date()).getTime()

      setInterval(function () {
        var currentTime = (new Date()).getTime()
        if (currentTime > (lastTime + 2000 * 2)) {
          window.location.reload()
        }
        lastTime = currentTime
      }, 2000)
    },
    destroyed () {

    },
    methods: {
      initWebSocket () {
        var socketProtocol = window.location.protocol.toLowerCase().indexOf('https') !== -1
          ? 'wss://'
          : 'ws://'
        var url = socketProtocol + window.location.host + '/tap'
        var ws = new WebSocket(url)
        var store = this.$store
        ws.onmessage = function (message) {
          var data = JSON.parse(message.data)
          if (data.event === 'heartbeat') {
            return
          }
          store.dispatch(data.event, data.payload)
          console.log(data)
        }
        return ws
      }
    }
  }
</script>

<style>
  #app {
    height: 100%;
  }
</style>
