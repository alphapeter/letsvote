<template>
  <div class="header">
    <h1 class="headerTitle">Let's vote!</h1>
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
        var user = JSON.parse(a)
        return user
      } catch (e) {
        return false
      }
    },
    profilePicture () {
      return 'https://www.gravatar.com/avatar/' + this.user.gravatar + '?s=50'
    },
    isLoggedIn () {
      return this.user
    }
  }
}
</script>

<style scoped>
  .header {
    position: fixed;
    top: 0;
    width: 100%;
    text-align: center;
    background-color: #607D8B;
    height: 4em;
    z-index: 10;
    box-shadow: 0px 12px 12px 0px rgba(0,0,0,0.15);
    color: #fff;
  }
  h1 {
    margin: 0.4em;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .header {
      position: relative;
    }
  }
</style>
