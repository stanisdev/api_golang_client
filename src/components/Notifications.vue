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
        width="80">
      </el-table-column>

      <el-table-column
        prop="message"
        label="Message">
      </el-table-column>

      <el-table-column
        label="Destination"
        width="100">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <p>{{ scope.row.link }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">Link</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>

      <el-table-column
        width="50"
        label="">
      </el-table-column>

      <el-table-column
        property="expired"
        label="Expired"
        width="150">
      </el-table-column>

      <el-table-column
        width="155"
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
      const {id} = row
      this.$confirm('This will permanently delete the notification. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        ApiService
          .setAuth().get
          .call(this, '/notification/delete/' + id)
          .then((data) => {
            this.notifications = this.notifications.filter((ntf) => +ntf.id !== +id)
            this.$message({
              message: 'Notification was removed successfully',
              type: 'success'
            })
          })
          .catch(Symbol)
      }).catch(Symbol)
    },
    fetchNotifications () {
      ApiService
        .setAuth().get
        .call(this, '/notification/list')
        .then((data) => {
          this.done = true
          this.notifications = data.payload || []
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
</style>
