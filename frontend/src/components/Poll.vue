<template>
  <div class="poll">
    <div class="header">
      <h2 class="name">{{poll.name}}</h2>
      <div class="created" :title="createdBy"><span class="status">{{status}}</span>{{printDate}}</div>
      <div v-if="poll.status < 10 && poll.status > 8"> Counting scores </div>

      <div class="editPoll" v-if="hasPermissionToEdit">

        <span @click="showMenu = !showMenu">
          <font-awesome :icon="icons.bars"/>
        </span>
        <div class="menu" v-if="showMenu">
          <span title="Back status"
                          v-if="poll.status >= 5"
                          @click="backStatus">
            <font-awesome :icon="icons.back"/>
          </span>
          <span title="Start poll"
                v-if="poll.status < 5"
                @click="activatePoll">
            <font-awesome :icon="icons.start"/>
          </span>
          <span title="End poll"
                v-else-if="poll.status < 8"
                @click="endPoll">
            <font-awesome :icon="icons.finish"/>
          </span>

          <span title="delete poll"
                class="deletePoll"
                v-if="canDelete"
                @click="confirmDelete">
            <font-awesome :icon="icons.trash"/>
          </span>
          <div v-if="showConfirmDelete"
               class="confirmDelete">
            Delete?
            <div>
              <span @click="showConfirmDelete = false">No!</span>
              <span @click="deletePoll">Yes</span>
            </div>
          </div>

        </div>
      </div>
    </div>
    <div v-if="false" class="description">{{poll.description}}</div>

    <poll-option v-for="option in options" :option="option" :poll="poll" :key="option.id" :maxScore="maxScore"></poll-option>

    <div v-if="statusCreated && isLoggedIn">
      <hr v-if="poll.options.length"/>
      <poll-option-add :poll="poll"></poll-option-add>
    </div>

    <div class="voterInfo"
         v-if="poll.status >=5">
         #voters: {{voterCount}}
    </div>
  </div>
</template>

<script>
import { API } from '../api.js'
import PollOption from './Option.vue'
import PollOptionAdd from './OptionAdd.vue'
import { gravatar } from '../gravatar.js'
import FontAwesome from '@fortawesome/vue-fontawesome'
import faTrashAlt from '@fortawesome/fontawesome-free-solid/faTrashAlt'
import faBars from '@fortawesome/fontawesome-free-solid/faBars'
import faFlagCheckered from '@fortawesome/fontawesome-free-solid/faFlagCheckered'
import faPlay from '@fortawesome/fontawesome-free-solid/faPlay'
import faStepBackward from '@fortawesome/fontawesome-free-solid/faStepBackward'
export default {
  props: ['poll'],
  components: {
    PollOption,
    PollOptionAdd,
    FontAwesome
  },
  data () {
    return {
      showMenu: false,
      showConfirmDelete: false
    }
  },
  computed: {
    icons () {
      return {
        trash: faTrashAlt,
        bars: faBars,
        finish: faFlagCheckered,
        start: faPlay,
        back: faStepBackward
      }
    },
    voterCount () {
      return this.$store.state.voters[this.poll.id] && this.$store.state.voters[this.poll.id].length
    },
    options () {
      var options = this.poll.options
      if (this.poll.status < 10) {
        options.sort((a, b) => {
          if (a.created_at < b.created_at) {
            return -1
          }
        })
      } else {
        options.sort((a, b) => {
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
        }
      }
      return options
    },
    maxScore () {
      let number = Math.max(...this.poll.options.map(o => o.score))
      return number
    },
    printDate () {
      var date = new Date(this.poll.created_at)
      return date.toLocaleDateString()
    },
    createdBy () {
      return 'created by ' + this.poll.created_by.name
    },
    profilePicture () {
      return gravatar.profilePicture(this.poll.created_by)
    },
    canDelete () {
      return this.hasPermissionToEdit
    },
    hasPermissionToEdit () {
      return (this.$store.state.me && this.poll.created_by.id === this.$store.state.me.id) || (this.$store.state.me && this.$store.state.me.is_admin)
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
    confirmDelete () {
      this.showConfirmDelete = true
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
    },
    activatePoll () {
      this.changeStatus('5')
    },
    endPoll () {
      this.changeStatus('8')
    },
    backStatus () {
      if (this.poll.status <= 5) {
        this.changeStatus('0')
      } else if (this.poll.status <= 10) {
        this.changeStatus('5')
      } else {
        this.changeStatus('5')
      }
    },
    changeStatus (code) {
      var store = this.$store
      this.showMenu = false
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
    box-shadow: 0px 8px 10px 0px rgba(0,0,0,0.15)
  }
  .editPoll {
    position: absolute;
    top: 2em;
    right: 0.5em;
    background-color: #448AFF;
    padding: .4em;
    border-radius: 1em;
    height: 1.2em;
    width: 1.2em;
    color: #FFFFFF;
    text-align: center;
    box-shadow: 2px 4px 4px 0px rgba(0,0,0,0.15)
  }

  .editPoll span {
    cursor: pointer;
  }
  .editPoll span:hover {
    color: #212121;
  }
  .status {
    font-size: 0.8em;
    height: 1em;
    margin-right: 0.5em;
    border-radius: 0.2em;
    background-color: #CFD8DC;
    color: #212121;
  }
  .header {
    width: 100%;
    border-bottom: 1px solid #BDBDBD;
  }
  .menu {
    position: absolute;
    right: 0em;
    top: 2em;
    background: #fff;
    border: 1px solid #BDBDBD;
    z-index: 2;
    color: #757575;
    box-shadow: 2px 4px 4px 0px rgba(0,0,0,0.15);
    display: flex;
    border-radius: 0.2em;
  }

  .menu > span {
    padding: 0.4em;
    margin: 0.5em;
    background-color: #448AFF;
    color: #fff;
    border-radius: 0.3em;
    display: inline-block;
    box-shadow: 2px 4px 4px 0px rgba(0,0,0,0.15);
  }

  .menu > span:hover {
    color: #212121;
  }

  .confirmDelete {
    position: absolute;
    top: 2em;
    background: #ffffff;
    right: 0;
    border: 1px solid red;
  }
  .voterInfo {
    margin-top: 2em;
    border-top: 1px solid #BDBDBD;
    text-align: right;
    padding: 0.5em;
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
