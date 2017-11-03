<template>
  <div class="header">
    <h1 class="headerTitle">Let's vote!</h1>
    <input class="loginButton" type="button" value="Login"/>
  </div>
</template>

<script>
  export default {
  }
</script>

<style scoped>
  .header {
    width: 100%;
    text-align: center;
    background-color: lightskyblue;
    position: relative;
  }
  .header h1 {
    padding: 10px;
  }
  .loginButton {
    position: absolute;
    right: 10px;
    top: 15px;
  }


</style>
