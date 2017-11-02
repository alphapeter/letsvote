<template>
  <Modal :disableButtons="true">
    <slot>
    </slot>
    <span>{{$store.state.ui.progress.message}}</span>
    <div class="progressBar">
      <div class="progress" :style="{width: width}">{{progress}}</div>
    </div>
  </Modal>
</template>

<script>
  import Modal from './Modal.vue'
  export default {
    components: {
      Modal
    },
    computed: {
      progress () {
        return this.$store.state.ui.progress.progress + '/' + this.$store.state.ui.progress.max
      },
      width () {
        return (100 / this.$store.state.ui.progress.max) * this.$store.state.ui.progress.progress + '%'
      }
    }
  }
</script>

<style>
  .progressBar {
    position: relative;
    top: 1em;
    background-color: white;
    width: 40em;
    border: none;
  }

  .progress {
    background-color: yellow;
    color: blue;
    height: 100%;
    text-align: right;
  }
</style>
