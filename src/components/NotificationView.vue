<template>
  <div>
    <menu-navigation></menu-navigation>

    <loader v-bind:is_showing="loading"></loader>
    <div v-if="done">
      <el-form ref="form" :model="notification" label-width="120px">
        <el-form-item label="Activity form">
          <el-input type="textarea" v-model="notification.text"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSave">Create</el-button>
          <el-button>Cancel</el-button>
        </el-form-item>
      </el-form>

    </div>
  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import Loader from '@/components/Loader.vue'
import ApiService from '@/common/api.service'

export default {
  name: 'NotificationView',
  components: {
    MenuNavigation,
    Loader
  },
  created () {
    this.fetchNotification()
  },
  methods: {
    fetchNotification () {
      const {id} = this.$route.params
      ApiService
        .setAuth().get
        .call(this, `/notifications/${id}`)
        .then((data) => {
          this.done = true
          this.notification = data.payload
        })
        .catch(Symbol)
        .finally(() => {
          this.loading = false
        })
    },
    onSave () {
      console.log('Saved')
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
