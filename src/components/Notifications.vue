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
        <template slot-scope="scope">
          <span>{{ cropString(scope.row.message, 38) }}</span>
        </template>
      </el-table-column>

      <el-table-column
        prop="company"
        label="Publisher"
        width="240">
        <template slot-scope="scope">
          <span>{{ cropString(scope.row.company, 22) }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="Destination"
        width="100">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <p v-bind:style="{fontFamily: 'Avenir, Helvetica, Arial, sans-serif'}">{{ scope.row.link }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">Link</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>

      <el-table-column
        width="17"
        label="">
      </el-table-column>

      <el-table-column
        property="expired"
        label="Expired"
        width="120">
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

    <div v-if="totalCount > 10" class="block pgn-ntfs">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page.sync="currentPage"
        :page-sizes="[10, 15, 25, 50]"
        :page-size="100"
        layout="sizes, prev, pager, next"
        :total="totalCount">
      </el-pagination>
    </div>

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
  mounted () {
    setTimeout(() => {
      document.querySelectorAll('[class*=el-select-dropdown]').forEach((elem) => {
        if (elem.tagName.toLowerCase() === 'li') {
          elem.style['font-family'] = 'Avenir, Helvetica, Arial, sans-serif'
        }
      })
    }, 400)
  },
  methods: {
    cropString (string, length) {
      if (typeof string !== 'string') {
        return ''
      }
      if (string.length <= length) {
        return string
      }
      return string.substr(0, length) + '...'
    },
    handleSizeChange (val) {
      this.perPage = val
      this.currentPage = 1
      this.fetchNotifications()
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.fetchNotifications()
    },
    handleEdit (index, row) {
      this.$router.push({
        name: 'NotificationView',
        params: {
          id: row.id
        }
      })
    },
    handleDelete (index, row) {
      setTimeout(() => {
        document.querySelectorAll('[class=el-message-box]').forEach((elem) => {
          elem.style['font-family'] = 'Avenir, Helvetica, Arial, sans-serif'
        })
      }, 25)
      this.$msgbox({
        title: 'Warning',
        message: 'This will permanently delete the notification. Continue?',
        type: 'warning',
        showCancelButton: true,
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        beforeClose: (action, instance, done) => {
          if (action === 'confirm') {
            instance.confirmButtonLoading = true
            instance.confirmButtonText = 'Loading...'

            const {id} = row
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
              .finally(() => {
                done()
                setTimeout(() => {
                  instance.confirmButtonLoading = false
                }, 300)
              })
          } else {
            done()
          }
        }
      }).then(Symbol).catch(Symbol)
    },
    fetchNotifications () {
      const publisherId = parseInt(this.$route.query.pub)
      const tasks = []
      const offset = (this.currentPage - 1) * this.perPage
      let ntfUrl = '/notification/list?limit=' + this.perPage + '&offset=' + offset
      let ntfCountUrl = '/notification/count'
      if (!isNaN(publisherId)) {
        ntfUrl += '&pub=' + publisherId
        ntfCountUrl += '?pub=' + publisherId
      }
      this.done = false

      tasks.push(ApiService
        .setAuth().get
        .call(this, ntfUrl))
      tasks.push(ApiService
        .setAuth().get
        .call(this, ntfCountUrl))
      Promise.all(tasks)
        .then(([ntfs, count]) => {
          this.done = true
          this.notifications = ntfs.payload || []
          this.totalCount = +count.payload || 0
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
      notifications: [],
      currentPage: 1,
      totalCount: 0,
      perPage: 10
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .pgn-ntfs {
    margin-top: 37px;
    margin-bottom: 27px;
  }
</style>
