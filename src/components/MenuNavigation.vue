<template>
  <div class="top-nav-menu">
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="1-notifications">Notifications</el-menu-item>
      <el-submenu index="2">
        <template slot="title">Workspace</template>
        <el-menu-item index="2-create-notification">Create notification</el-menu-item>
      </el-submenu>
      <el-submenu v-bind:style="{ 'float': 'right' }" index="3">
        <template slot="title">Settings</template>
        <el-menu-item index="3-change-password">Change password</el-menu-item>
        <el-menu-item index="3-logout">Logout</el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
const routes = {
  'create-notification': 'CreateNotification',
  notifications: 'Notifications',
  'change-password': 'ChangePassword'
}

const urls = {
  '/notifications': '1-notifications',
  '/notifications/create': '2-create-notification'
}

/* eslint-disable */
export default {
  name: 'MenuNavigation',
  data () {
    let url = window.location.href.split('#')[1];
    url = urls[url]
    if (!url) {
      url = '1-notifications'
    }
    return {
      activeIndex: url
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      let path = key.split(/^\d+\-/ig)[1]
      if (typeof path === 'string' && routes.hasOwnProperty(path)) {
        path = routes[path]
      } else {
        path = routes.notifications
      }
      this.$router.push({ name: path })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .top-nav-menu {
    margin-bottom: 37px;
  }
</style>
