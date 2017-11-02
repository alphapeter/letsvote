<template>
  <Modal class="overlay">
    <span slot="title">New directory</span>
    <input id="directoryName"
           autofocus
           type="text"
           v-model="name"
           placeholder="name"/>
    <button slot="buttons"
            v-text="'OK'"
            @click="mkdir">
    </button>
  </Modal>
</template>

<script>
  import Modal from './Modal.vue'
  import { Rpc } from '../../rpc'
  export default {
    components: {
      Modal
    },
    data () {
      return {
        name: '',
        keypress: null
      }
    },
    methods: {
      mkdir () {
        this.$store.commit('mkdirWait')
        let currentState = this.$store.getters.currentState
        let path = this.$store.getters.currentPathString
        Rpc.call('mkdir', [currentState.selectedRoot + path, this.name])
          .then((response) => {
            if (response.error) {
              this.$store.commit('error', response.error)
            } else {
              this.$store.commit('commandFinished')
            }
          })
      }
    },
    created () {
      var vm = this
      this.keypress = function (e) {
        if (e.key === 'Enter' && e.target.nodeName !== 'BUTTON') {
          if (vm.name.length) {
            vm.mkdir()
          }
        }
      }
      window.addEventListener('keyup', this.keypress)
    },
    mounted () {
      document.getElementById('directoryName').focus()
    },
    destroyed () {
      window.removeEventListener('keyup', this.keypress)
    }
  }
</script>

<style>

</style>
