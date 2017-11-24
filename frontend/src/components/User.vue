<template>
    <div class="userInfo">

        <img class="profilePicture"
             :src="profilePicture"
             :title="user.name"
             align="middle"/>
        <span class="username">{{user.name}}</span>
      </div>
</template>

<script>
  import { gravatar } from '../gravatar.js'
  export default {
    props: ['user', 'picturesize'],
    computed: {
      profilePicture () {
        return gravatar.profilePicture(this.user, 24)
      }
    }
  }
</script>

<style scoped>
  .profilePicture {
    border-radius: 15px;
  }
  .username {
    text-overflow: ellipsis;
    line-height: 30px;
    vertical-align: middle;
  }

</style>
