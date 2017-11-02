<template>
  <Modal class="overlay">
    <span slot="title">Rename file/directory</span>
    <input id="directoryName"
           autofocus
           type="text"
           v-model="name"
           placeholder="name"/>
    <button slot="buttons"
            v-text="'OK'"
            @click="rename">
    </button>
  </Modal>
</template>

<script>
  import Modal from './Modal.vue'
  import { Rpc } from '../../rpc'
  import { mapGetters } from 'vuex'
  export default {
    components: {
      Modal
    },
    data () {
      return {
        name: '',
        oldName: '',
        keypress: null
      }
    },
    computed: {
      ...mapGetters([
        'selectedFiles'
      ])
    },
    methods: {
      rename () {
        this.$store.commit('renameWait')
        let currentState = this.$store.getters.currentState
        let path = currentState.selectedRoot + this.$store.getters.currentPathString + '/'
        Rpc.call('mv', [path + this.oldName, path + this.name])
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
            vm.rename()
          }
        }
      }
      this.name = this.selectedFiles[0]
      this.oldName = this.selectedFiles[0]
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
