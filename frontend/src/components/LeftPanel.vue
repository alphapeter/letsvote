<template>
  <div class="leftpanel">
    <div class="content">
      <div>
        <user v-if="isLoggedIn" :user="user" :size="24"></user>
      </div>
      <div class="center">
              </div>
    </div>
    <div class="menu" >
      <div v-if="isAdmin && !isSmallScreen" class="menuitem">
        <A href="/admin">Administrate users</A>
      </div>

      <div class="menuitem" v-if="isLoggedIn">
        <a href="/auth/logout">
          logout
        </a>
      </div>

      <div class="menuitem" v-if="!isLoggedIn">
        <a href="/auth/login/office365">
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
    },
    isAdmin () {
      return this.$store.state.me && this.$store.state.me.is_admin
    },
    isSmallScreen () {
      if (window.matchMedia('(min-width: 1023px)').matches) {
        return false
      } else {
        return true
      }
    }
  }
}
</script>

<style scoped>
  .menu {
    margin-top: 2em;
    background-color: #7f7f7f;
    position: relative;
    border-top: 1px solid #ddd;
    margin-right: 0.4em;
  }

  .menuitem {
    background-color: #e9ebee;
    padding: 0.5em;
    border-bottom: 1px solid #ddd;
  }

  .menuitem:hover {
    background-color: #ffffff;

  }

  .menuitem a {
    font-size: 0.75em;
    text-decoration: none;
    color: black;
    width: 100%;
  }

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
