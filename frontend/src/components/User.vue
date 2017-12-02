<template>
    <div class="userInfo">

        <img class="profilePicture"
             v-bind:class="{voting: voting}"
             :src="profilePicture"
             :title="user.name"
             align="middle"/>
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
        return gravatar.profilePicture(this.user, 24)
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
    border-radius: 15px;
  }
  .voting {
    border: 5px solid red;
  }
  .username {
    text-overflow: ellipsis;
    line-height: 30px;
    vertical-align: middle;
  }

</style>
