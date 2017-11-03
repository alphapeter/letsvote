<template>
  <div id="app">
    <pageHeader></pageHeader>
    <activePolls></activePolls>
    <historicPolls></historicPolls>
    <dialogs></dialogs>
  </div>
</template>

<script>
  import ActivePolls from './components/ActivePolls.vue'
  import HistoricPolls from './components/HistoricPolls.vue'
  import Dialogs from './components/Dialogs/Dialogs.vue'
  import PageHeader from './components/PageHeader.vue'
  import Poll from './components/Poll.vue'

  export default {
    name: 'app',
    components: {
      ActivePolls,
      HistoricPolls,
      Dialogs,
      PageHeader,
      Poll
    },
    data () {
      return {
      }
    },
    methods: {

    },
    created () {
      this.$store.dispatch('init')
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
