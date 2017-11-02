<template>
  <div class="overlay">
    <div class="dialog">
      <slot name="title"></slot>
      <slot>

      </slot>
      <div v-if="!disableButtons" class="buttons">
        <button v-text="'Cancel'" @click="cancel"></button>
        <slot name="buttons"></slot>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        eventListener: null
      }
    },
    props: ['disableButtons'],
    methods: {
      cancel () {
        this.$store.commit('commandCanceled')
      }
    },
    created () {
      if (this.disableButtons) {
        return
      }

      let vm = this
      this.eventListener = (e) => {
        if (e.key === 'Escape') {
          vm.cancel()
        }
      }
      window.addEventListener('keyup', this.eventListener)
    },
    destroyed () {
      if (this.disableButtons) {
        return
      }
      window.removeEventListener('keyup', this.eventListener)
    }
  }
</script>

<style scoped>
  .overlay {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
  }

  .dialog {
    position: fixed;
    top: calc(50% - 10em / 2);
    left: calc(50% - 40em / 2);
    border: 6px double white;
    width: 40em;
    height: 6em;
    background-color: gray;
    margin-left: auto;
    margin-right: auto;
    text-align: center;
    color: white;
    padding: 1em;
  }

  .buttons {
    position: absolute;
    width: calc(100% - 2em);
    bottom: 1em;
    margin: auto;
    text-align: center;
  }

  input {
    position: relative;
    top: 1em;
    background-color: white;
    width: 40em;
    border: none;
  }
</style>
