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
        let polls = this.$store.state.polls
        return polls.sort((a, b) => {
          if (a.created_at > b.created_at) {
            return -1
          }
          return 1
        })
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
