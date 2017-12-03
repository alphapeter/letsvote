<template>
  <div class="poll">
    <div class="header">
      <h2 class="name">{{poll.name}}</h2>
      <div class="created" :title="createdBy"><span class="status">{{status}}</span>{{printDate}}</div>
      <div class="editPoll" v-if="hasPermissionToEdit">
        <input v-if="canDelete" class="deletePollButton" type="button" @click="deletePoll" value="Delete poll"/>
        <input v-if="poll.status < 5" class="activatePollButton" type="button" @click="activatePoll" value="Activate poll"/>
        <input v-else-if="poll.status < 8" class="activatePollButton" type="button" @click="endPoll" value="End poll"/>
        <div v-else-if="poll.status < 10"> Counting scores </div>
      </div>
    </div>
    <div v-if="false" class="description">{{poll.description}}</div>


    <poll-option v-for="option in options" :option="option" :poll="poll" :key="option.id" :totalScore="totalScore"></poll-option>

    <div v-if="statusCreated && isLoggedIn">
      <hr v-if="poll.options.length"/>
      <poll-option-add :poll="poll"></poll-option-add>
    </div>
  </div>
</template>

<script>
  import { API } from '../api.js'
  import PollOption from './Option.vue'
  import PollOptionAdd from './OptionAdd.vue'
  import { gravatar } from '../gravatar.js'
  export default {
    props: ['poll'],
    components: {
      PollOption,
      PollOptionAdd
    },
    computed: {
      options () {
        if (this.poll.status < 10) {
          this.poll.options.sort((a, b) => {
            if (a.created_at < b.created_at) {
              return -1
            }
          })
        } else {
          this.poll.options.sort((a, b) => {
            return b.score - a.score
          })
          var position = 1
          for (var i = 0; i < this.poll.options.length; i++) {
            var option = this.poll.options[i]
            this.$set(option, 'position', position)
            option.position = position
            if (i + 1 < this.poll.options.length && this.poll.options[i + 1].score !== option.score) {
              position++
            }
            console.log('hejs')
          }
        }
        return this.poll.options
      },
      totalScore () {
        return this.poll.options.reduce((acc, current) => acc + current.score, 0)
      },
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
      canDelete () {
        return this.hasPermissionToEdit && this.poll.status === 0
      },
      hasPermissionToEdit () {
        return this.$store.state.me && this.poll.created_by.id === this.$store.state.me.id
      },
      isLoggedIn () {
        return this.$store.state.me
      },
      statusCreated () {
        return this.poll.status === 0
      },
      statusActive () {
        return this.poll.status === 5
      },
      status () {
        let status = this.poll.status
        if (status < 5) {
          return 'Open'
        }
        if (status < 8) {
          return 'Voting'
        }
        if (status < 10) {
          return 'Counting'
        }
        return 'Ended'
      }
    },
    methods: {
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
      },
      activatePoll () {
        this.changeStatus('5')
      },
      endPoll () {
        this.changeStatus('8')
      },
      changeStatus (code) {
        var store = this.$store
        API.patch('api/polls/' + this.poll.id, { status: code })
          .then((response) => {
            if (!response.success) {
              store.commit('error', {
                message: response.reason || 'unexpected error :('
              })
            }
          }).catch((reason) => {
            console.log(reason)
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
  .name {
    margin: 1em;
  }
  .created {
    position: absolute;
    right: 0.5em;
    top: 0.8em;
    font-size: 0.8em;
    margin: 0.1em;
  }
  .description {
    font-style: italic;
    margin: 10px;
  }
  .poll {
    background-color: #FFF;
    margin-bottom: 20px;
    position: relative;
    border: 1px solid #e9e8e8;
    border-radius: 5px;
  }
  .editPoll {
    position: absolute;
    top: 50px;
    right: 5px;
  }
  .status {
    font-size: 0.5em;
    height: 1em;
    margin-right: 0.5em;
    border-radius: 0.2em;
    background-color: lightblue;
  }
  .header {
    width: 100%;
    border-bottom: 1px solid #e9ebee;
  }

  @media only screen
  and (max-device-width : 1024px) {
    .poll {
      margin-bottom: 1em;
    }
    .created {
      font-size: 20pt;
    }
  }
</style>
