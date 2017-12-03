<template>
    <div class="userInfo">

        <img class="profilePicture"
             v-bind:class="{voting: voting}"
             :src="profilePicture"
             :title="user.name"/>
        <span class="username">{{user.name}}</span>
      </div>
</template>

<script>
  import { gravatar } from '../gravatar.js'
  import {EventBus} from '../EventBus'

  export default {
    data () {
      return {
        voting: false
      }
    },
    props: ['user', 'picturesize'],
    computed: {
      profilePicture () {
        return gravatar.profilePicture(this.user, 100)
      }
    },
    created () {
      var that = this
      EventBus.$on('USER_VOTED', (info) => {
        if (this.$store.state.me && info.user_id === this.$store.state.me.id) {
          return
        }
        if (info.user_id === this.user.id) {
          that.voting = !that.voting
        }
      })
    }
  }
</script>

<style scoped>
  .profilePicture {
    border-radius: 0.5em;
    vertical-align: middle;
    height: 1em;
    width: 1em;
  }
  .voting {
    border: 5px solid red;
  }
  .username {
    text-overflow: ellipsis;
    line-height: 100%;
    vertical-align: middle;
  }
  .userInfo {
    border-bottom: 1px solid #ddd;
  }

</style>
