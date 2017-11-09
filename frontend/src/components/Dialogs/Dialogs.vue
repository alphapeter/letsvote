<template>
  <wait-dialog v-if="$store.state.ui.state === 'loading'">loading...</wait-dialog>
  <error-dialog v-else-if="$store.state.ui.state === 'error'"></error-dialog>
  <login-dialog v-else-if="$store.state.ui.state === 'login'"></login-dialog>
</template>

<script>
  import WaitDialog from './WaitDialog.vue'
  import ErrorDialog from './ErrorDialog.vue'
  import LoginDialog from './LoginDialog.vue'

  export default {
    components: {
      WaitDialog,
      ErrorDialog,
      LoginDialog
    }
  }
</script>
