<template>
  <div v-if="isLoggedIn" class="poll create">
    <input type="text" v-model="newPoll.name" placeholder="Create a new poll...">
    <textarea v-if="newPoll.name" v-model="newPoll.description" placeholder="description" rows="2"></textarea>
    <input v-if="newPoll.name" @click="addPoll" type="button" value="create"/>
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
    width: calc(100% - 400px);
    background-color: #FFF;
    margin-top: 20px;
    margin-bottom: 50px;
    border: 1px solid #e9e8e8;
    border-radius: 5px;
  }
  .poll.create input[type=text], textarea {
    width: calc(100% - 10px);
    margin: 5px;
    border: none;
  }

  .poll.create input[type=text] {
    font-size: 16px;
  }

  .poll.create input[type=button] {
    margin: 5px;
    float: right;
  }
</style>
