<template>
    <div class="option">
      <span>{{option.name}}</span>
      <input class="deleteOptionButton" type="button" name="delete" value="delete option" @click="deleteOption"/>
    </div>
</template>

<script>
  import { API } from '../api.js'

  export default {
    props: ['option'],
    computed: {
      printDate () {
        var date = new Date(this.option.created_at)
        return date.toLocaleDateString()
      },
      createdBy () {
        return 'created by ' + this.option.created_by.name
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
