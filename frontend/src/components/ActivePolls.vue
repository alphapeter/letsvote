<template>
  <div class="polls">
    <poll-add></poll-add>
    <poll v-for="poll in polls" :key="poll.id" :poll="poll"></poll>
  </div>
</template>

<script>
  import Poll from './Poll.vue'
  import PollAdd from './PollAdd'
  export default {
    components: {
      Poll,
      PollAdd
    },
    computed: {
      polls () {
        return this.$store.state.polls
      },
      isLoggedIn () {
        return this.$store.state.me
      }
    }
  }
</script>

<style scoped>
  .polls {
    position: relative;
    left: 200px;
    margin-top: 120px;
    margin-left: 10px;
    margin-right: 10px;
    position: relative;
    clear: both;
  }
</style>
