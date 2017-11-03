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
          {{i}}.{{o.name}} ({{o.created_by.name}})
        </div>
      </div>
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
  }
  .poll {
    width: 300px;
    height: 300px;
    background-color: blanchedalmond;
    border: 1px solid grey;
    border-radius: 5px;
    padding: 15px;
    margin: 20px;
  }
  .name {
    font-weight: bolder;
  }
  .description {
    font-style: italic;
    margin: 10px;
  }
</style>
