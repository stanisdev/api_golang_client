<template>
  <div>
    <menu-navigation></menu-navigation>
    <notificationForm v-bind:mode="'create'" v-bind:notification="notification"></notificationForm>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import NotificationForm from '@/components/NotificationForm.vue'

export default {
  name: 'CreateNotification',
  components: {
    MenuNavigation,
    NotificationForm
  },
  data () {
    return {
      notification: {}
    }
  },
  methods: {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
