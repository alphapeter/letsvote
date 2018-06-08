<template>
  <font-awesome :icon="icons.thumbsUp" class="thumb"/>
</template>

<script>
import FontAwesome from '@fortawesome/vue-fontawesome'
import faThumbsUp from '@fortawesome/fontawesome-free-solid/faThumbsUp'

export default {
  components: {
    FontAwesome
  },
  computed: {
    icons () {
      return {
        thumbsUp: faThumbsUp
      }
    }
  }
}
</script>

<style scoped>
  .thumb {
    position: absolute;
    left: 0;
    top: 0;
  }

</style>
