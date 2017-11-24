<template>
  <div class="polls">
    <h2>
      Active polls
    </h2>
    <poll v-for="poll in polls" :key="poll.id" :poll="poll"></poll>
    <div v-if="isLoggedIn" class="poll create">
      <input type="text" v-model="newPoll.name" placeholder="Create a new poll...">
      <textarea v-if="newPoll.name" v-model="newPoll.description" placeholder="description" rows="2"></textarea>
      <input v-if="newPoll.name" @click="addPoll" type="button" value="create"/>
    </div>
  </div>
</template>

<script>
  import { API } from '../api.js'
  import Poll from './Poll.vue'
  export default {
    components: {
      Poll
    },
    data () {
      return {
        newPoll: {
          name: '',
          description: ''
        }
      }
    },
    computed: {
      polls () {
        return this.$store.state.polls
      },
      isLoggedIn () {
        return this.$store.state.me
      }
    },
    methods: {
      addPoll () {
        // todo should be an action in the store to be dispatched
        var store = this.$store
        var that = this
        API.post('api/polls', {
          name: that.newPoll.name,
          description: this.newPoll.description
        }).then((response) => {
          if (!response.success) {
            store.commit('error', {
              message: 'Fail! :(',
              code: 500
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

<style>
  .polls {
    position: relative;
    left: 200px;
    margin-top: 100px;
    margin-left: 10px;
    margin-right: 10px;
    position: relative;
    clear: both;
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
  .name {
    font-weight: bolder;
  }
  .created {
    float: right;
    margin: 2px;
  }
  .description {
    font-style: italic;
    margin: 10px;
  }
  .poll {
    width: calc(100% - 400px);
    background-color: #FFF;
    margin-bottom: 20px;
    float: left;
    position: relative;
    border: 1px solid #e9e8e8;
    border-radius: 5px;
  }
</style>
