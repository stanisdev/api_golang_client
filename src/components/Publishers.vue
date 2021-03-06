<template>
  <div>
    <menu-navigation></menu-navigation>

    <loader v-bind:is_showing="loading"></loader>
    <transition name="fade">
    <el-table
      ref="singleTable"
      v-if="done"
      :data="publishers"
      style="width: 100%;">

      <el-table-column
        prop="id"
        label="ID"
        width="80">
      </el-table-column>

      <el-table-column
        prop="name"
        label="Name">
        <template slot-scope="scope">
          <a class="a-class" @click="handleClickPublisher(scope.row.id)">{{ scope.row.name }}</a>
        </template>
      </el-table-column>

      <el-table-column
        prop="notifications_count"
        label="Notifications count"
        width="190">
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
  name: 'Publishers',
  components: {
    MenuNavigation,
    Loader
  },
  created () {
    this.fetchPublishers()
  },
  methods: {
    handleClickPublisher (id) {
      this.$router.push({
        name: 'Notifications',
        query: {
          pub: id
        }
      })
    },
    handleEdit (index, row) {
      this.$router.push({
        name: 'PublisherView',
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
        message: 'This will permanently delete the publisher, but the notifications will not be deleted. Continue?',
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
              .call(this, '/publisher/delete/' + id)
              .then((data) => {
                this.publishers = this.publishers.filter((publ) => +publ.id !== +id)
                this.$message({
                  message: 'Publisher was removed successfully',
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
    handleSizeChange (val) {
      this.perPage = val
      this.currentPage = 1
      this.fetchPublishers()
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.fetchPublishers()
    },
    fetchPublishers () {
      const tasks = []
      const offset = (this.currentPage - 1) * this.perPage
      const publUrl = '/publisher/list?limit=' + this.perPage + '&offset=' + offset
      this.done = false

      tasks.push(ApiService
        .setAuth().get
        .call(this, publUrl))
      tasks.push(ApiService
        .setAuth().get
        .call(this, '/publisher/count'))

      Promise.all(tasks)
        .then(([publishers, count]) => {
          this.done = true
          this.publishers = publishers.payload || []
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
      currentPage: 1,
      totalCount: 0,
      publishers: [],
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
  .a-class {
    color: #409eff;
    text-decoration: none;
    cursor: pointer;
  }
</style>
