<template>
  <div class="poll">
    <div class="header">
      <span class="name">{{poll.name}}</span>
      <span class="created" :title="createdBy">{{printDate}}</span>
      <input class="deletePollButton" type="button" name="Delete poll" @click="deletePoll" value="Delete poll"/>
    </div>
    <div class="description">{{poll.description}}</div>
    <hr/>
    <poll-option v-for="option in poll.options" :option="option" :key="option.id"></poll-option>
    <div v-if="!poll.has_ended">
      <hr/>
      <div  class="option create">
        <img src="https://www.gravatar.com/avatar/dfsdf?s=24">
        <input type="text" placeholder="Create a new option..." style="border: none" v-model="newOption.name"/>
        <textarea v-if="newOption.name" v-model="newOption.description" placeholder="description" rows="2"></textarea>
        <input v-if="newOption.name" type="button" name="Add" @click="addOption()" value="Add"/>
      </div>
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
        newOption: {
          name: '',
          description: ''
        }
      }
    },
    computed: {
      printDate () {
        var date = new Date(this.poll.created_at)
        return date.toLocaleDateString()
      },
      createdBy () {
        return 'created by ' + this.poll.created_by.name
      }
    },
    methods: {
      addOption () {
        var store = this.$store
        var that = this
        API.post('api/polls/' + this.poll.id + '/options', {
          name: this.newOption.name,
          description: this.newOption.description
        }).then((response) => {
          if (response.success) {
            that.newOption.name = ''
            that.newOption.description = ''
          } else {
            store.commit('error', {
              message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(',
              code: 500
            })
          }
        }).catch((reason) => {
          store.commit('error', {
            message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(',
            code: reason.code
          })
        })
      },
      deletePoll () {
        var store = this.$store
        API.delete('api/polls/' + this.poll.id)
          .then((response) => {
            if (!response.success) {
              store.commit('error', {
                message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(',
                code: 500
              })
            }
          }).catch((reason) => {
            store.commit('error', {
              message: 'Ooops, something went terribly wrong. Bad code monkey! We could not add your poll :(',
              code: reason.code
            })
          })
      }
    }
  }
</script>

<style scoped>
  .option.create input[type=text], textarea {
    width: calc(100% - 60px)
  }
  .deletePollButton {
    float: right;
    right: 5px;
  }
</style>
