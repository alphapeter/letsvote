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
      <div class="meter" v-bind:style="{ width: meter + '%' }"><span class="score">{{option.score}}</span></div>
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
  props: ['option', 'poll', 'order', 'maxScore'],
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
      return this.option.score / this.maxScore * 100
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
    display: flex;
    justify-content: space-between;
  }

  .name {
    max-width: 69%;
    display: inline-block;
    vertical-align: middle;
    margin-top: 0.2em;
  }
  .votes {
    display: inline-block;
    vertical-align: middle;
   }

  .vote {
    background-color: #448AFF;
    color: #fff;
    padding: 0.2em;
    margin-left: 0.1em;
    cursor: pointer;
    display: inline-block;
    border: 1px solid transparent;
    box-shadow: 2px 4px 4px 0px rgba(0,0,0,0.15)
  }
  .vote:hover {
    color: #212121;
  }
  .vote.selected {
    font-weight: bolder;
  }
  .position {
    position: relative;
    width: 1em;
    height: 1em;
    font-size: 100%;
    border-radius: 0.6em;
    text-align: center;
    margin-right: 0.5em;
    vertical-align: middle;
    display: inline-block;
    padding: 0.1em;
  }
  .winner {
    background: radial-gradient(ellipse farthest-corner at right bottom, #FEDB37 0%, #FDB931 8%, #9f7928 30%, #8A6E2F 40%, transparent 80%),
    radial-gradient(ellipse farthest-corner at left top, #FFFFFF 0%, #FFFFAC 8%, #D1B464 25%, #5d4a1f 62.5%, #5d4a1f 100%);
    border: 1px solid #8A6E2F;
    color: #5D4A1F;
    text-shadow: 0px 1px 0px rgba(255,255,255,.3), 0px -1px 0px rgba(0,0,0,.7);

  }
  .second {
    background: radial-gradient(ellipse farthest-corner at right bottom, #eeeeee 0%, #ededed 8%, #dddddd 30%, #cbcbcb 40%, transparent 80%),
    radial-gradient(ellipse farthest-corner at left top, #FFFFFF 0%, #fffbf4 8%, #eeeeee 25%, #a1a1a1 62.5%, #efefef 100%);
    border: 1px solid #a1a1a1;
    color: #a1a1a1;
    text-shadow: 0px 1px 0px rgba(255,255,255,.3), 0px -1px 0px rgba(0,0,0,.7);
  }

  .third {
    background: radial-gradient(ellipse farthest-corner at right bottom, #ff9036 0%, #d4702b 8%, #a14521 30%, #a14521 40%, transparent 80%),
    radial-gradient(ellipse farthest-corner at left top, #fff6f2 0%, #ffdeca 8%, #ca7345 25%, #6e2a1e 62.5%, #ca7345 100%);
    border: 1px solid #6E2A1E;
    color: #6E2A1E;
    text-shadow: 0px 1px 0px rgba(255,255,255,.3), 0px -1px 0px rgba(0,0,0,.7);
  }
  .none {
    background-color: black;
    border: 1px solid #888888;
    text-shadow: 0px 1px 0px rgba(255,255,255,.3), 0px -1px 0px rgba(0,0,0,.7);
    color: #888888;
  }

  .scoring {
    height: 1em;
    width: 10em;
    background-color: #CFD8DC;
    position: absolute;
    right: 0;
    top: 0.2em;
    max-width: 30%;
    border-radius: 0.5em;
    overflow: hidden;
    margin-top: 0.1em;
  }
  .meter {
    height: 100%;
    background-color: #455A64;
    color: #fff;
  }
  .meter .score {
    margin-right: 0.2em;
    height: 100%;
    float: right;
    text-align: right;
  }
  .profilePicture {
    width: 1em;
    height: 1em;
    vertical-align: middle;
    border-radius: 0.5em;
  }

  @media only screen
  and (max-device-width : 1023px) {
    .scoring {
      width: 5em;
    }
  }

</style>
