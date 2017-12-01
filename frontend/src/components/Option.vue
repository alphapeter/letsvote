<template>
  <div class="option">
    <img :src="profilePicture"/>
    <span v-on:mouseenter="active = true"
          v-on:mouseleave="active = false">
      {{option.name}}
        <span v-if="active">edit</span>
    </span>

    <input v-if="canDelete" class="deleteOptionButton" type="button" name="delete" value="delete option" @click="deleteOption"/>
    <div v-if="canVote" class="votes" >
      <div class="vote" v-bind:class="{selected: score == 0}" @click="vote(0)">0</div>
      <div class="vote" v-bind:class="{selected: score == 1}" @click="vote(1)">1</div>
      <div class="vote" v-bind:class="{selected: score == 2}" @click="vote(2)">2</div>
      <div class="vote" v-bind:class="{selected: score == 3}" @click="vote(3)">3</div>
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
    props: ['option', 'poll'],
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
      canDelete () {
        return this.createdByMe === this.me.id && !this.isActive
      },
      canVote () {
        return this.isActive && !this.createdByMe
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
        this.$store.dispatch('vote', {
          score: count,
          option_id: this.option.id,
          poll_id: this.poll.id
        })
        console.log('vote' + count)
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
    background-color: gray;
    padding: 5px;
    margin-left: 2px;
    float: left;
    cursor: pointer;
  }
  .vote.selected {
    background-color: green;
  }
</style>
