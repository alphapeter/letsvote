<template>
  <Modal disableButtons=true>
    <div class="message">
      <slot></slot>
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
  .message{
    position: relative;
    top: calc(50% - 1em/2)
  }
</style>
