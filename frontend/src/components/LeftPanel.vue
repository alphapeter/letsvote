<template>
  <div class="leftpanel">
    <div class="userInfo">
      <div>
        <img class="profilePicture"
             v-if="isLoggedIn"
             :src="profilePicture"
             :title="user.name"
             align="middle"/>
        <span class="username"
              v-if="isLoggedIn">{{user.name}}</span>
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
  export default {
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
      profilePicture () {
        return 'https://www.gravatar.com/avatar/' + this.user.gravatar + '?s=24'
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
