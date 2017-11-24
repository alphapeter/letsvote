<template>
  <div class="leftpanel">
    <div class="userInfo">
      <div>
        <user v-if="isLoggedIn" :user="user" :size="24"></user>
      </div>
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
        try {
          var a = this.$cookie.get('lets_vote.authenticated')
          var b = decodeURIComponent(a)
          var user = JSON.parse(b)
          return user
        } catch (e) {
          return false
        }
      },
      isLoggedIn () {
        return this.user
      }
    }
  }
</script>

<style scoped>
  .leftpanel {
    top: 100px;
    left: 0;
    width: 200px;
    height: 100%;
    position: fixed;
  }
  .profilePicture {
    border-radius: 15px;
  }
  .username {
    text-overflow: ellipsis;
    line-height: 30px;
    vertical-align: middle;
  }

</style>
