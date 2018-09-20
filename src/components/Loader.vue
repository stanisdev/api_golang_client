<template>
  <div>
    <i v-show="is_showing" class="el-icon-loading c-loader"></i>
  </div>
</template>

<script>

export default {
  name: 'Loader',
  props: ['is_showing'],
  data () {
    return {}
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .c-loader {
    font-size: 22px;
  }
</style>
