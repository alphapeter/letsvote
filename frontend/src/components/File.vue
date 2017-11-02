<template>
  <div class="file"
       :class="{'selected': selected, 'focused': focused}"
       tabindex="-1">
    <div class="fileName">{{file.name}}</div>
    <div class="fileModified">{{modified(file.modified)}}</div>
    <div class="fileSize" v-if="file.type == 'd'">&ltdir&gt</div>
    <div class="fileSize" v-else>{{size(file.size)}}</div>
  </div>
</template>
<script>
  export default{
    props: ['file'],
    computed: {
      selected () {
        return this.file.selected
      },
      focused () {
        return this.file.focused
      }
    },
    methods: {
      size (size) {
        if (size < 1000) {
          return size
        }
        if (size < 1000000) {
          return Math.round(size / 1000) + 'KB'
        }
        if (size < 1000000000) {
          return Math.round(size / 1000000) + 'MB'
        }

        if (size < 1000000000000) {
          return Math.round(size / 1000000000) + 'GB'
        }
        if (size < 1000000000000000) {
          return Math.round(size / 1000000000000) + 'TB'
        }
        return Math.round(size / 1000000000000000) + 'PB'
      },
      modified (date) {
        return new Date(date).toLocaleString()
      }
    },
    watch: {
      focused () {
        if (this.focused) {
          this.$el.focus()
        }
      }
    }
  }
</script>
<style>
  .file {
    cursor: default;
    position: relative;
    width: 100%;
  }

  .file:hover {
    color: white;
  }

  .file.focused {
    background: gray;
  }

  .file.selected {
    background: yellow;
    color: blue;
  }

  .file.selected.focused {
    background: green;
    color: blue;
  }

  .fileName {
    display: inline-block;
    overflow-x: hidden;
    word-break: break-all;
    width: calc(53% - 2px);
    height: 1em;
    margin: 0;
    white-space: nowrap;
  }

  .fileSize {
    display: inline-block;
    overflow-x: hidden;
    width: calc(12% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    white-space: nowrap;
  }

  .fileModified {
    display: inline-block;
    overflow-x: hidden;
    width: calc(33% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    overflow-y: hidden;
    white-space: nowrap;
  }
</style>
