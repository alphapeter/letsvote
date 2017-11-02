<template>
  <wait-dialog v-if="$store.state.ui.state === 'loading'">loading...</wait-dialog>
  <error-dialog v-else-if="$store.state.ui.state === 'error'"></error-dialog>
</template>

<script>
  import WaitDialog from './WaitDialog.vue'
  import ErrorDialog from './ErrorDialog.vue'

  export default {
    components: {
      WaitDialog,
      ErrorDialog
    }
  }
</script>
