<template>
  <div class="option create">
    <img :src="profilePicture">
    <input type="text" placeholder="Create a new option..." style="border: none" v-model="newOption.name"/>
    <textarea v-if="newOption.name" v-model="newOption.description" placeholder="description" rows="2"></textarea>
    <input v-if="newOption.name" type="button" name="Add" @click="addOption()" value="Add"/>
  </div>
</template>

<script>
  import { API } from '../api.js'
  import { gravatar } from '../gravatar.js'
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
      profilePicture () {
        return gravatar.profilePicture(this.user)
      },
      user () {
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
      }
    }
  }
</script>

<style scoped>
  .option.create input[type=text], textarea {
    width: calc(100% - 60px)
  }
  .option {
    margin: 10px;
  }
</style>
