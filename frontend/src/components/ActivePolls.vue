<template>
  <div class="polls">
    <h2>
      Aktiva omröstningar
    </h2>
    <div class="poll" v-for="poll in polls" :key="poll.id">
      <div class="name">{{poll.name}}</div>
      <div class="description">{{poll.description}}</div>
      <div class="options" v-for="(o, i) in poll.options">
        <div>
          {{i+1}}.{{o.name}} ({{o.created_by.name}})
        </div>
      </div>

    </div>
    <div class="poll create">
      <input type="text" v-model="newPoll.name" placeholder="name">
      <textarea v-model="newPoll.description" placeholder="description" rows="8"></textarea>
      <input @click="addPoll" type="button" value="create"/>
      {{newPoll}}
    </div>
  </div>
</template>

<script>
  import { API } from '../api.js'
  export default {
    props: ['a', 'b'],
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
        API.post('api/polls', {
          name: this.newPoll.name,
          description: this.newPoll.description
        })
      }
    }
  }
</script>

<style scoped>
  .polls {
    margin: 10px;
    position: relative;
    clear: both;
  }
  .poll {
    width: 300px;
    height: 300px;
    background-color: blanchedalmond;
    border: 1px solid grey;
    border-radius: 5px;
    padding: 15px;
    margin: 20px;
    float: left;
    position: relative;
  }

  .poll.create input[type=text], textarea {
    width: calc(100% - 10px);
    margin: 5px;
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
  .description {
    font-style: italic;
    margin: 10px;
  }
</style>
