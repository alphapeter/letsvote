<template>
  <Modal class="overlay">
    <span slot="title">Delete {{fileCount}} file(s)</span>
    <button slot="buttons"
            v-text="'OK'"
            @click="deleteFile">
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
        fileCount: 0,
        keypress: null
      }
    },
    computed: {
      ...mapGetters([
        'selectedFiles',
        'currentPathString',
        'currentState'
      ])
    },
    methods: {
      deleteFile () {
        this.$store.commit('deleteFileWait')
        let currentState = this.currentState
        let path = currentState.selectedRoot + this.currentPathString
        let vm = this
        vm.$store.commit('startProgress', {
          max: this.selectedFiles.length
        })
        let fileIndex = 0

        function del (index) {
          let fileName = vm.selectedFiles.splice(0, 1)
          vm.$store.commit('progress', {
            message: fileName,
            progress: fileIndex
          })
          Rpc.call('rm', [path + '/' + fileName])
            .then(response => {
              if (response.error) {
                vm.$store.commit('error', response.error)
              } else if (vm.selectedFiles.length === 0) {
                vm.$store.commit('commandFinished')
              } else {
                fileIndex++
                del()
              }
            })
        }

        del()
      }
    },
    created () {
      this.fileCount = this.selectedFiles.length
      var vm = this
      this.keypress = function (e) {
        if (e.key === 'Enter' && e.target.nodeName !== 'BUTTON') {
          vm.deleteFile()
        }
      }
      window.addEventListener('keyup', this.keypress)
    },
    destroyed () {
      window.removeEventListener('keyup', this.keypress)
    }
  }
</script>

<style>

</style>
