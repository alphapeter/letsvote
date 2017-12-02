<template>
  <div class="option">
    <div class="position"
         v-bind:class="{winner: isWinner, second: isSecond, third: isThird}"
         v-if="poll.status >=10">
      {{option.position}}
    </div>

    <img :src="profilePicture"/>
    <span v-on:mouseenter="active = true"
          v-on:mouseleave="active = false">
      {{option.name}}
      <span v-if="active">edit</span>
    </span>
    <input v-if="canDelete" class="deleteOptionButton" type="button" name="delete" value="delete option" @click="deleteOption"/>
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
  export default {
    data () {
      return {
        active: false
      }
    },
    props: ['option', 'poll', 'order', 'totalScore'],
    computed: {
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
        return this.createdByMe === this.me.id && !this.isActive
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

<style>
  .deleteOptionButton {
    float: right;
  }
  .option {
    margin: 10px;
    position: relative;
  }
  .votes {
    position: absolute;
    right: 10px;
    top: 2px;
  }
  .vote {
    background-color: lightblue;
    padding: 5px;
    margin-left: 2px;
    float: left;
    cursor: pointer;
  }
  .vote.selected {
    font-weight: bolder;
  }
  .position {
    position: relative;
    width: 22px;
    height: 22px;
    border-radius: 15px;
    text-align: center;
    float: left;
    margin-right: 15px;
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
  .noselect {
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
  }
  .scoring {
    width: 100px;
    height: 22px;
    background-color: #e9ebee;
    position: absolute;
    right: 0px;
    top: 2px;
  }
  .meter {
    height: 100%;
    float: left;
    background-color: red;
  }
</style>
