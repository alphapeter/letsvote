<template>
  <div class="option">
    <div class="name">
      <div class="position"
           v-bind:class="{winner: isWinner, second: isSecond, third: isThird}"
           v-if="poll.status >=10">
        {{option.position}}
      </div>

      <img class="profilePicture" :title="createdBy" :src="profilePicture"/>
      <span>
        {{option.name}}
      </span>

      <button v-if="canDelete" class="deleteOptionButton" @click="deleteOption">
        <font-awesome :icon="icons.trash"/>
      </button>
    </div>

    <div v-if="canVote" class="votes" >
      <div class="vote noselect" v-bind:class="{none: score == 0}" @click="vote(0)">0</div>
      <div class="vote noselect" v-bind:class="{third: score == 1, selected: score == 1 }" @click="vote(1)">1</div>
      <div class="vote noselect" v-bind:class="{second: score == 2, selected: score == 2}" @click="vote(2)">2</div>
      <div class="vote noselect" v-bind:class="{winner: score == 3, selected: score == 3}" @click="vote(3)">3</div>
    </div>
    <div v-if="poll.status >= 10" class="scoring">
      <div class="meter" v-bind:style="{ width: meter + '%' }">{{option.score}}</div>
    </div>
  </div>
</template>

<script>
import { API } from '../api.js'
import { gravatar } from '../gravatar.js'
import FontAwesome from '@fortawesome/vue-fontawesome'
import faTrashAlt from '@fortawesome/fontawesome-free-solid/faTrashAlt'
export default {
  data () {
    return {
      active: false
    }
  },
  components: {
    FontAwesome
  },
  props: ['option', 'poll', 'order', 'totalScore'],
  computed: {
    icons () {
      return {
        trash: faTrashAlt
      }
    },
    printDate () {
      var date = new Date(this.option.created_at)
      return date.toLocaleDateString()
    },
    createdBy () {
      return 'created by ' + this.option.created_by.name
    },
    profilePicture () {
      return gravatar.profilePicture(this.option.created_by)
    },
    me () {
      return this.$store.state.me ? this.$store.state.me : {}
    },
    isLoggedIn () {
      return this.$store.state.me
    },
    canDelete () {
      return ((this.createdByMe === this.me.id) || (this.$store.state.me && this.$store.state.me.is_admin)) && this.poll.status < 5
    },
    canVote () {
      return this.isActive && !this.createdByMe && this.isLoggedIn
    },
    createdByMe () {
      return this.me.id === this.option.created_by.id
    },
    isActive () {
      return this.poll.status === 5
    },
    score () {
      var votes = this.$store.state.votes
      if (!(votes && votes[this.poll.id])) {
        return 0
      }
      if (votes[this.poll.id].score_1 === this.option.id) {
        return 1
      }
      if (votes[this.poll.id].score_2 === this.option.id) {
        return 2
      }
      if (votes[this.poll.id].score_3 === this.option.id) {
        return 3
      }
      return 0
    },
    isWinner () {
      return this.option.position === 1
    },
    isSecond () {
      return this.option.position === 2
    },
    isThird () {
      return this.option.position === 3
    },
    meter () {
      return this.option.score / this.totalScore * 100
    }
  },
  methods: {
    deleteOption () {
      var store = this.$store
      API.delete('api/polls/' + this.option.pollId + '/options/' + this.option.id)
        .then((response) => {
          if (!response.success) {
            store.commit('error', {
              message: 'Could not be deleted',
              code: 500
            })
          }
        }).catch((reason) => {
          store.commit('error', {
            message: 'Could not be deleted',
            code: reason.code
          })
        })
    },
    vote (count) {
      if (this.score === count) {
        return
      }
      this.$store.dispatch('vote', {
        score: count,
        option_id: this.option.id,
        poll_id: this.poll.id
      })
    }
  }
}
</script>

<style scoped>

  .deleteOptionButton {
    display: none;
    border: 0;
    background: transparent;
    cursor: pointer;
  }
  .name:hover .deleteOptionButton {
    display: inline-block;

  }
  .deleteOptionButton:hover {
    color: #7f7f7f;
  }
  .option {
    margin: 0.5em;
    position: relative;
  }

  .name {
    max-width: 69%;
  }
  .votes {
    position: absolute;
    right: 0;
    top: 0.1em;
  }
  .vote {
    background-color: lightblue;
    padding: 0.2em;
    margin-left: 0.4em;
    float: left;
    cursor: pointer;
  }
  .vote.selected {
    font-weight: bolder;
  }
  .position {
    position: relative;
    width: 1.2em;
    height: 1.2em;
    font-size: 100%;
    border-radius: 0.6em;
    text-align: center;
    float: left;
    margin-right: 0.5em;
    vertical-align: middle;
    display: inline-block;
  }
  .winner {
    background-color: gold;
  }
  .second {
    background-color: silver;
  }
  .third {
    background-color: #cd7f32;
  }
  .none {
    background-color: black;
  }

  .scoring {
    height: 1em;
    width: 10em;
    background-color: #e9ebee;
    position: absolute;
    right: 0;
    top: 0.2em;
    max-width: 30%;
  }
  .meter {
    height: 100%;
    float: left;
    background-color: red;
    text-align: right;
  }
  .profilePicture {
    width: 1em;
    height: 1em;
    vertical-align: middle;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .scoring {
      width: 5em;
    }
  }

</style>
