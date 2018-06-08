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
    width: calc(100% - 20em);
    left: 10em;
    margin-top: 120px;
    margin-left: 0;
    margin-right: 0;
    position: relative;
    clear: both;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .polls {
      width: 100%;
      top: 0;
      left: 0;
      margin-top: 20px;
    }
  }
</style>
