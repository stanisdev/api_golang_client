<template>
  <div>
    <menu-navigation></menu-navigation>

    <loader v-bind:is_showing="loading"></loader>
    <transition name="fade">
    <div v-if="done">
      <notificationForm v-bind:mode="'edit'" v-bind:notification="notification"></notificationForm>
    </div>
    </transition>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import Loader from '@/components/Loader.vue'
import NotificationForm from '@/components/NotificationForm.vue'
import ApiService from '@/common/api.service'

export default {
  name: 'NotificationView',
  components: {
    MenuNavigation,
    Loader,
    NotificationForm
  },
  created () {
    this.fetchNotification()
  },
  methods: {
    fetchNotification () {
      const {id} = this.$route.params
      ApiService
        .setAuth().get
        .call(this, `/notification/get/${id}`)
        .then((data) => {
          this.done = true
          let exp = data.payload.expired + '000'
          exp = new Date(+exp)
          const mnth = exp.getMonth() + 1
          const day = exp.getDate()
          exp = exp.getFullYear() + '-' + mnth + '-' + day
          data.payload.expired = exp
          this.notification = data.payload
        })
        .catch(Symbol)
        .finally(() => {
          this.loading = false
          const ta = document.getElementsByTagName('textarea')[0]
          if (ta instanceof Object) {
            ta.style.height = '150px'
            ta.style['font-size'] = '14px'
            ta.style['font-family'] = "'Avenir', Helvetica, Arial, sans-serif"
          }
        })
    }
  },
  data () {
    return {
      loading: true,
      done: false,
      notification: {}
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
