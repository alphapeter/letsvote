<template>
  <wait-dialog v-if="$store.state.ui.state === 'initializing'">loading...</wait-dialog>
  <wait-dialog v-else-if="$store.state.ui.state === 'rename-wait'">renaming file</wait-dialog>
  <wait-dialog v-else-if="$store.state.ui.state === 'mkdir-wait'">creating new directory</wait-dialog>

  <progress-dialog v-else-if="$store.state.ui.state === 'copy-wait'">copying files...</progress-dialog>
  <progress-dialog v-else-if="$store.state.ui.state === 'move-wait'">moving files...</progress-dialog>
  <progress-dialog v-else-if="$store.state.ui.state === 'delete-file-wait'">deleting files...</progress-dialog>

  <rename-dialog v-else-if="$store.state.ui.state === 'rename'">renaming file</rename-dialog>
  <mkdir-dialog v-else-if="$store.state.ui.state === 'mkdir'"></mkdir-dialog>
  <error-dialog v-else-if="$store.state.ui.state === 'error'"></error-dialog>
  <delete-dialog v-else-if="$store.state.ui.state === 'delete-file'"></delete-dialog>
</template>

<script>
  import WaitDialog from './WaitDialog.vue'
  import RenameDialog from './RenameDialog.vue'
  import MkdirDialog from './MkdirDialog.vue'
  import ErrorDialog from './ErrorDialog.vue'
  import DeleteDialog from './DeleteDialog.vue'
  import ProgressDialog from './ProgressDialog.vue'

  export default {
    components: {
      WaitDialog,
      RenameDialog,
      MkdirDialog,
      DeleteDialog,
      ErrorDialog,
      ProgressDialog
    }
  }
</script>
