<template>
    <div class="option">
      <div class="header">
        <div class="name">{{poll.name}}</div>
        <div class="created" :title="createdBy">{{printDate}}</div>
      </div>
    </div>
</template>

<script>
  export default {
    props: ['option'],
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
      createOption () {
        console.log('hej')
      }
    }
  }
</script>

<style scoped>

</style>
