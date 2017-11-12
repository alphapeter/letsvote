<template>
  <div class="polls">
    <h2>
      Previous polls
    </h2>
    <div class="poll" v-for="poll in polls" :key="poll.id">
      <div class="name">{{poll.name}}</div>
      Vinnare: Herr Brun
    </div>
  </div>
</template>

<script>
  export default {
    computed: {
      polls () {
        return this.$store.state.polls
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
    height: 50px;
    background-color: beige;
    border: 1px solid grey;
    border-radius: 5px;
    padding: 15px;
    margin: 20px;
    float: left;
  }
</style>
