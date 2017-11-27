<template>
  <div class="option">
    <img :src="profilePicture"/>
    <span v-on:mouseenter="active = true"
          v-on:mouseleave="active = false">
      {{option.name}}
        <span v-if="active">edit</span>
    </span>

    <input v-if="createdByMe" class="deleteOptionButton" type="button" name="delete" value="delete option" @click="deleteOption"/>
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
    props: ['option'],
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
      createdByMe () {
        return this.$store.state.me && this.option.created_by.id === this.$store.state.me.id
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
      m () {
        this.active = true
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
  }
</style>
