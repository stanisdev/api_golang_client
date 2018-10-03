<template>
  <div>
    <menu-navigation></menu-navigation>
    <span>User Profile</span>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'

export default {
  name: 'UserProfile',
  components: {
    MenuNavigation
  },
  data () {
    return {}
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
