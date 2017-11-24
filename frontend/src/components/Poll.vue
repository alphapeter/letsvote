<template>
  <div class="poll">
    <div class="header">
      <span class="name">{{poll.name}}</span>
      <span class="created" :title="createdBy">{{printDate}}</span>
      <input v-if="createdByMe" class="deletePollButton" type="button" name="Delete poll" @click="deletePoll" value="Delete poll"/>
    </div>
    <div class="description">{{poll.description}}</div>
    <hr/>
    <div v-if="!poll.options.length">no options yet...</div>
    <poll-option v-for="option in poll.options" :option="option" :key="option.id"></poll-option>
    <div v-if="!poll.has_ended && isLoggedIn">
      <hr v-if="poll.options.length"/>
      <div class="option create">
        <img :src="profilePicture">
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
  import { gravatar } from '../gravatar.js'
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
      },
      profilePicture () {
        gravatar.profilePicture(this.poll.created_by)
      },
      createdByMe () {
        return this.$store.state.me && this.poll.created_by.id === this.$store.state.me.id
      },
      isLoggedIn () {
        return this.$store.state.me
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
