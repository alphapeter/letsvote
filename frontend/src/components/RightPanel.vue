<template>
  <div class="rightpanel">
    <div class="content">
      <label class="label">Active users</label>
      <user v-for="user in activeUsers" :user="user" :key="user.id"></user>
    </div>

  </div>
</template>

<script>
import User from './User.vue'

export default {
  components: {
    User
  },
  computed: {
    activeUsers () {
      return this.$store.state.activeUsers.map(user => {
        return {
          name: user.id === 'anonymous' ? user.id : user.name,
          gravatar: user.gravatar,
          id: user.id
        }
      })
    }
  }
}
</script>

<style scoped>
  .rightpanel {
    top: 5em;
    right: 0;
    width: 10em;
    height: 100%;
    position: fixed;
  }
  .label {
    margin-left: 0.5em;
    font-weight: bold;
  }

  .content {
    margin-right: 0.4em;
    margin-left: 0.4em;
    box-shadow: 0px 8px 10px 0px rgba(0,0,0,0.15)
  }

  @media only screen
  and (max-device-width : 1023px) {
    .rightpanel {
      display: none;
    }
  }

</style>
