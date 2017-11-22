<template>
    <div class="poll">
      <div class="header">
        <div class="name">{{poll.name}}</div>
        <div class="created" :title="createdBy">{{printDate}}</div>
      </div>
      <div class="description">{{poll.description}}</div>
      <hr/>
      <div class="options" v-for="(o, i) in poll.options">
        <div class="option">
          {{i+1}}.{{o.name}} ({{o.created_by.name}})
        </div>
      </div>
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
  export default {
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
        // todo should be an action in the store to be dispatched
        var store = this.$store
        var that = this
        API.post('api/polls/' + this.poll.id + '/options', {
          name: this.newOption.name,
          description: this.newOption.description
        }).then((response) => {
          if (response.success) {
            store.commit('addOption', response.option)
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
      }
    }
  }
</script>

<style scoped>
  .option.create input[type=text], textarea {
    width: calc(100% - 60px)
  }
</style>
