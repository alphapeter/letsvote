<template>
  <Modal>
    <div class="message">
      {{$store.state.ui.error.message}}
    </div>
  </Modal>
</template>

<script>
import Modal from './Modal.vue'
export default {
  components: {
    Modal
  }
}
</script>

<style scoped>
  .message {
    margin-top: 1em;
    top: calc(50% - 2em / 2);
    color: red;
    background-color: lightgray;
  }
  button {
  : relative;
    bottom: 1em;
  }
</style>
