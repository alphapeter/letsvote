<template>
    <div class="poll">
        <div>{{poll.name}}</div>
        <div>{{poll.description}}</div>
    </div>
</template>

<script>
  export default {
    props: ['poll']
  }
</script>

<style scoped>
  .poll {
    width: 100%;
  }
</style>
