<template>
  <div class="polls">
    <h2>
      Active polls
    </h2>
    <poll v-for="poll in polls" :key="poll.id" :poll="poll"></poll>
    <div class="poll create">
      <input type="text" v-model="newPoll.name" placeholder="Create a new poll...">
      <textarea v-if="newPoll.name" v-model="newPoll.description" placeholder="description" rows="8"></textarea>
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
      }
    },
    methods: {
      addPoll () {
        // todo should be an action in the store to be dispatched
        var store = this.$store
        var that = this
        API.post('api/polls', {
          name: this.newPoll.name,
          description: this.newPoll.description
        }).then((response) => {
          if (response.success) {
            store.commit('addPoll', response.poll)
            that.newPoll.name = ''
            that.newPoll.description = ''
          } else {
            store.commit('error', { message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(', code: 500 })
          }
        }).catch((reason) => {
          store.commit('error', { message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(', code: reason.code })
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
    position: absolute;
    bottom: 20px;
    right: 20px;
    margin: 5px;
  }
  .name {
    font-weight: bolder;
  }
  .created {
    position: absolute;
    right: 10px;
    top: 2px
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
