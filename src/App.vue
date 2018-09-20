<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
body {
  margin: 0px;
  padding: 0px;
  padding-top: 26px;
}
.page-container {
  box-sizing: border-box;
  height: 100%;
  width: 940px;
  margin: 0 auto;
}
</style>
