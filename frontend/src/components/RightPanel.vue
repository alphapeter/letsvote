<template>
  <div class="rightpanel">
    Active users
    <user v-for="user in activeUsers" :user="user" :key="user.id"></user>
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
    top: 100px;
    right: 0;
    width: 200px;
    height: 100%;
    position: fixed;
  }


</style>
