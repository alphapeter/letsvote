<template>
    <div>
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

</style>
