<template>
  <div class="rightpanel">
    Active users
    <div v-for="user in activeUsers">{{user.name}}</div>
  </div>
</template>

<script>
  export default {
    methods: {
      login () {
      }
    },
    computed: {
      activeUsers () {
        return this.$store.state.activeUsers.map(user => {
          return {
            name: user.id === 'anonymous' ? user.id : user.name,
            gravatar: user.gravatar
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
