<template>
  <div class="leftpanel">
    <div class="content">
      <div>
        <user v-if="isLoggedIn" :user="user" :size="24"></user>
      </div>
      <div class="center">
        <a href="/auth/logout"
           v-if="isLoggedIn">
          logout
        </a>

        <a href="/auth/login/office365"
           v-if="!isLoggedIn">
          login
        </a>
      </div>
  </div>
  </div>
</template>

<script>
  import User from './User.vue'
  export default {
    components: {
      User
    },
    methods: {
      login () {
        window.location.href = '/auth/login/office365'
      }
    },
    computed: {
      user () {
        return this.$store.state.me
      },
      isLoggedIn () {
        return this.user
      }
    }
  }
</script>

<style scoped>
  .leftpanel {
    top: 5em;
    left: 0;
    width: 10em;
    position: fixed;
  }
  .center {
    text-align: center;
    padding-top: 0.1em;
  }
  .content {
    margin-left: 0.4em;
    margin-right: 0.4em;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .leftpanel {
      position: relative;
      width: 100%;
      text-align: center;
      top: 0;
    }
  }

</style>
