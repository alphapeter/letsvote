<template>
    <div class="userInfo" v-bind:class="{voting: voting}">
        <img class="profilePicture"
             :src="profilePicture"
             :title="user.name"/>
      <span class="username" @click="test()">{{user.name}}</span>
      <transition-group name="fade" tag="p">
        <thumb v-for="t in thumbs" :key="t"/>
      </transition-group>
    </div>
</template>

<script>
  import { gravatar } from '../gravatar.js'
  import {EventBus} from '../EventBus'
  import Thumb from './Thumb.vue'

  export default {
    data () {
      return {
        voting: false,
        thumbs: []
      }
    },
    props: ['user', 'picturesize'],
    computed: {
      profilePicture () {
        return gravatar.profilePicture(this.user, 100)
      }
    },
    components: {
      Thumb
    },
    created () {
      var that = this
      EventBus.$on('USER_VOTED', (info) => {
        if (this.$store.state.me && info.user_id === this.$store.state.me.id) {
          return
        }
        if (info.user_id === this.user.id) {
          that.thumbs.push(Date.now())
          setTimeout(() => {
            that.thumbs.shift()
          }, 0)
          that.voting = true
          setTimeout(() => {
            that.voting = false
          }, 500)
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
    background-color: #efefef;
  }
  .username {
    text-overflow: ellipsis;
    line-height: 100%;
    vertical-align: middle;
  }
  .userInfo {
    border-bottom: 1px solid #ddd;
    position: relative;
  }

  .fade-leave-active {
    transition: opacity 1.2s;
  }
  .fade-leave-to {
    opacity: 0;
    transition: all 1.2s;
    transform: scale(1.2) translate(5em);
    color: #0074D9;
  }

</style>
