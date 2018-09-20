<template>
  <div>
    <menu-navigation></menu-navigation>

    <loader v-bind:is_showing="loading"></loader>
    <transition name="fade">
    <el-table
      ref="singleTable"
      v-if="done"
      :data="notifications"
      style="width: 100%;">

      <el-table-column
        prop="id"
        label="ID"
        width="100">
      </el-table-column>

      <el-table-column
        property="date"
        label="Date"
        width="170">
      </el-table-column>

      <el-table-column
        prop="text"
        label="Text">
      </el-table-column>

      <el-table-column
        width="175"
        label="Operations">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
        </template>
      </el-table-column>

    </el-table>
    </transition>

  </div>
</template>

<script>
import MenuNavigation from '@/components/MenuNavigation.vue'
import Loader from '@/components/Loader.vue'
import ApiService from '@/common/api.service'

export default {
  name: 'Notifications',
  components: {
    MenuNavigation,
    Loader
  },
  created () {
    this.fetchNotifications()
  },
  methods: {
    handleEdit (index, row) {
      this.$router.push({
        name: 'NotificationView',
        params: {
          id: row.id
        }
      })
    },
    handleDelete (index, row) {
      console.log(row.id)
    },
    fetchNotifications () {
      ApiService
        .setAuth().get
        .call(this, '/notifications')
        .then((data) => {
          this.done = true
          const {notifications} = data.payload
          this.notifications = notifications || []
        })
        .catch(Symbol)
        .finally(() => {
          this.loading = false
        })
    }
  },
  data () {
    return {
      done: false,
      loading: true,
      notifications: []
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>
