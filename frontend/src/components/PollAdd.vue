<template>
  <div v-if="isLoggedIn" class="poll create">
    <input type="text" v-model="newPoll.name" placeholder="Create a new poll...">
    <textarea v-if="newPoll.name" v-model="newPoll.description" placeholder="description" rows="2"></textarea>
    <div v-if="newPoll.name" class="buttonPlaceHolder">
      <input @click="addPoll" type="button" value="create"/>
    </div>
  </div>
</template>

<script>
  import { API } from '../api.js'
  import PollOption from './Option.vue'
  export default {
    components: {
      PollOption
    },
    props: ['poll'],
    data () {
      return {
        newPoll: {
          name: '',
          description: ''
        }
      }
    },
    computed: {
      isLoggedIn () {
        return this.$store.state.me
      }
    },
    methods: {
      addPoll () {
        var store = this.$store
        var that = this
        API.post('api/polls', {
          name: that.newPoll.name,
          description: this.newPoll.description
        }).then((response) => {
          if (!response.success) {
            store.commit('error', {
              message: response.reason
            })
          } else {
            that.newPoll.name = ''
            that.newPoll.description = ''
          }
        }).catch((reason) => {
          store.commit('error', { message: 'Fail', code: reason.code })
        })
      }
    }
  }
</script>

<style scoped>
  .poll {
    width: 100%;
    background-color: #FFF;
    margin-top: 1em;
    margin-bottom: 1em;
    border: 1px solid #e9e8e8;
    border-radius: 0.2em;
  }
  .poll.create input[type=text], textarea {
    width: calc(100% - 1em);
    margin: 0.2em;
    border: none;
  }

  .poll.create input[type=text] {
    font-size: 1.2em;
  }

  .poll.create input[type=button] {
    margin: 0.2em;
  }
  .buttonPlaceHolder {
    width: 100%;
    text-align: right;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .poll {
      margin-bottom: 2em;
    }
    .poll input[type=text], textarea, input[type=button] {
      margin: 0;
      font-size: 1em;
      width: 100%;
    }

    .buttonPlaceHolder {
      text-align: center;
      width: 100%;
    }

    .poll input[type=button] {
      width: 95%
    }
    .poll.create input[type=text] {
      width: calc(100% - 2em);
    }

  }
</style>
